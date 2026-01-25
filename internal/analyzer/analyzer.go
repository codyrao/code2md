package analyzer

import (
	"code2md/models"
	"strings"
)

type Analyzer struct {
	callGraph    *models.CallGraph
	allElements  map[string]*models.CodeElement
	fileElements map[string][]*models.CodeElement
}

func NewAnalyzer() *Analyzer {
	return &Analyzer{
		callGraph: &models.CallGraph{
			Nodes: make(map[string]*models.CallNode),
			Edges: make([]models.CallEdge, 0),
		},
		allElements:  make(map[string]*models.CodeElement),
		fileElements: make(map[string][]*models.CodeElement),
	}
}

func (a *Analyzer) Analyze(project *models.ProjectInfo) *models.CallGraph {
	a.callGraph = &models.CallGraph{
		Nodes: make(map[string]*models.CallNode),
		Edges: make([]models.CallEdge, 0),
	}

	for _, file := range project.Files {
		for _, elem := range file.Elements {
			a.allElements[elem.ID] = elem
			a.fileElements[file.Path] = append(a.fileElements[file.Path], elem)
		}
	}

	for _, elem := range project.AllElements {
		if _, exists := a.callGraph.Nodes[elem.ID]; !exists {
			a.createNode(elem)
		}

		for _, call := range elem.Calls {
			targetID := a.findElementByName(call, elem.FilePath)
			if targetID != "" {
				a.addEdge(elem.ID, targetID, "calls")

				if node, exists := a.callGraph.Nodes[targetID]; exists {
					if node.File != elem.FilePath {
						elem.References = append(elem.References, targetID)
					}
				}
			}
		}
	}

	a.resolveInheritance(project)

	a.buildCallGraph()

	return a.callGraph
}

func (a *Analyzer) createNode(elem *models.CodeElement) {
	node := &models.CallNode{
		ID:       elem.ID,
		Name:     elem.Name,
		Type:     string(elem.Type),
		File:     elem.FilePath,
		Line:     elem.StartLine,
		Outgoing: make([]string, 0),
		Incoming: make([]string, 0),
	}
	a.callGraph.Nodes[elem.ID] = node
}

func (a *Analyzer) addEdge(fromID, toID, callType string) {
	if fromID == toID {
		return
	}

	edge := models.CallEdge{
		From:     fromID,
		To:       toID,
		Type:     "call",
		CallType: callType,
	}

	a.callGraph.Edges = append(a.callGraph.Edges, edge)

	if fromNode, exists := a.callGraph.Nodes[fromID]; exists {
		fromNode.Outgoing = append(fromNode.Outgoing, toID)
	}

	if toNode, exists := a.callGraph.Nodes[toID]; exists {
		toNode.Incoming = append(toNode.Incoming, fromID)
	}
}

func (a *Analyzer) findElementByName(name string, currentFile string) string {
	for _, elems := range a.fileElements {
		for _, elem := range elems {
			if elem.Name == name || elem.FullName == name || strings.HasSuffix(elem.FullName, "."+name) {
				if elem.FilePath == currentFile || isImported(elem, currentFile, a.fileElements) {
					return elem.ID
				}
			}
		}
	}
	return ""
}

func isImported(target *models.CodeElement, currentFile string, fileElements map[string][]*models.CodeElement) bool {
	currentElems := fileElements[currentFile]
	for _, elem := range currentElems {
		if elem.Type == models.ElementTypeImport {
			importPath := elem.Name
			if strings.Contains(importPath, target.Name) {
				return true
			}
		}
	}
	return false
}

func (a *Analyzer) resolveInheritance(project *models.ProjectInfo) {
	for _, elem := range project.AllElements {
		for _, baseClass := range elem.Extends {
			baseID := a.findElementByName(baseClass, elem.FilePath)
			if baseID != "" {
				a.addEdge(elem.ID, baseID, "extends")

				if baseElem, exists := a.allElements[baseID]; exists {
					baseElem.Children = append(baseElem.Children, elem.ID)
				}
			}
		}

		for _, interfaceName := range elem.Implements {
			interfaceID := a.findElementByName(interfaceName, elem.FilePath)
			if interfaceID != "" {
				a.addEdge(elem.ID, interfaceID, "implements")

				if interfaceElem, exists := a.allElements[interfaceID]; exists {
					interfaceElem.Children = append(interfaceElem.Children, elem.ID)
				}
			}
		}
	}
}

func (a *Analyzer) buildCallGraph() {
	visited := make(map[string]bool)

	for id := range a.callGraph.Nodes {
		if !visited[id] {
			a.dfsTraverse(id, visited, 0)
		}
	}
}

func (a *Analyzer) dfsTraverse(nodeID string, visited map[string]bool, depth int) {
	if visited[nodeID] {
		return
	}

	visited[nodeID] = true

	if node, exists := a.callGraph.Nodes[nodeID]; exists {
		node.Visited = true
		node.Depth = depth

		for _, neighborID := range node.Outgoing {
			if !visited[neighborID] {
				a.dfsTraverse(neighborID, visited, depth+1)
			}
		}
	}
}

func (a *Analyzer) GetCallDepth(elemID string) int {
	if node, exists := a.callGraph.Nodes[elemID]; exists {
		return node.Depth
	}
	return 0
}

func (a *Analyzer) GetIncomingCalls(elemID string) []string {
	if node, exists := a.callGraph.Nodes[elemID]; exists {
		return node.Incoming
	}
	return nil
}

func (a *Analyzer) GetOutgoingCalls(elemID string) []string {
	if node, exists := a.callGraph.Nodes[elemID]; exists {
		return node.Outgoing
	}
	return nil
}

func (a *Analyzer) FindPublicAPI(project *models.ProjectInfo) []*models.CodeElement {
	publicAPI := make([]*models.CodeElement, 0)

	for _, elem := range project.AllElements {
		if elem.Visibility == models.VisibilityPublic && elem.Type != models.ElementTypeVariable {
			publicAPI = append(publicAPI, elem)
		}
	}

	return publicAPI
}

func (a *Analyzer) FindEntryPoints(project *models.ProjectInfo) []*models.CodeElement {
	entryPoints := make([]*models.CodeElement, 0)

	for _, elem := range project.AllElements {
		if elem.Name == "main" || elem.Name == "init" || elem.Name == "__main__" {
			entryPoints = append(entryPoints, elem)
		}

		if elem.Type == models.ElementTypeFunction && len(elem.Calls) == 0 && elem.Visibility == models.VisibilityPublic {
			entryPoints = append(entryPoints, elem)
		}
	}

	return entryPoints
}

func (a *Analyzer) AnalyzeComplexity(elem *models.CodeElement) int {
	complexity := 1

	for _, call := range elem.Calls {
		if strings.Contains(call, "if") || strings.Contains(call, "else") ||
			strings.Contains(call, "while") || strings.Contains(call, "for") ||
			strings.Contains(call, "switch") || strings.Contains(call, "case") {
			complexity++
		}
	}

	return complexity
}

func (a *Analyzer) FindDependencies(project *models.ProjectInfo) map[string][]string {
	dependencies := make(map[string][]string)

	for _, file := range project.Files {
		deps := make([]string, 0)
		for _, elem := range file.Elements {
			for _, call := range elem.Calls {
				if !stringInSlice(call, deps) {
					deps = append(deps, call)
				}
			}
		}
		dependencies[file.Path] = deps
	}

	return dependencies
}

func stringInSlice(str string, list []string) bool {
	for _, s := range list {
		if s == str {
			return true
		}
	}
	return false
}
