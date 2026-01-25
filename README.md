# Code2MD - 代码文档生成器

## 简介

Code2MD 是一个用 Go 编写的命令行工具，用于解析项目源代码并自动生成全面的 Markdown 格式的代码文档。

### 主要功能

- **多语言支持**: 支持 Go、JavaScript、TypeScript、Java、Python、C++、PHP 等主流编程语言
- **自动解析**: 自动提取函数、类、变量、常量、接口等代码元素
- **调用关系分析**: 分析函数调用关系和继承关系
- **完整文档**: 生成包含项目概览、API 参考、调用图谱的详细文档
- **灵活配置**: 支持自定义输出路径、排除模式、语言过滤等

## 安装

### 方式一：使用 Go 安装

```bash
go install github.com/code2md/code2md@latest
```

### 方式二：使用 Docker

```bash
docker pull codyrao/code2md:latest
```

### 方式三：从源码编译

```bash
git clone https://github.com/codyrao/code2md.git
cd code2md
go build -o code2md ./cmd/main.go
```

## 使用方法

### 基本用法

```bash
# 分析当前目录并生成 code.md
code2md

# 分析指定目录
code2md -p /path/to/project

# 指定输出文件名
code2md -o my-project-docs.md
```

### 完整参数

```bash
code2md [OPTIONS]

Options:
  -p, --path string      项目路径 (默认: ./)
  -o, --output string    输出 markdown 文件名 (默认: code.md)
  -l, --language string  要包含的编程语言 (默认: 全部)
  -e, --exclude string   要排除的模式 (例如: 'vendor/*', '*.test.go')
  -v, --verbose          启用详细输出
  -H, --hidden           包含隐藏文件和目录
  --no-structure         跳过目录树结构
  --no-code              跳过完整代码部分
  --no-api               跳过 API 文档部分
  --help, -h             显示帮助信息
  --version              显示版本信息
```

### 使用示例

```bash
# 详细模式分析
code2md -p ./myproject -v

# 只分析 Go 和 Python 文件
code2md -p ./myproject -l go -l python

# 排除测试文件和 vendor 目录
code2md -p ./myproject -e '**/*_test.go' -e 'vendor/*'

# 指定输出位置
code2d -p ./myproject -o ./docs/api.md
```

## Docker 使用方法

### 基本命令

```bash
# 分析当前目录
docker run -v $(pwd):/app codyrao/code2md -p /app

# 分析指定目录并输出到宿主机
docker run -v /path/to/project:/app -v $(pwd):/output codyrao/code2md -p /app -o /output/code.md

# 使用详细模式
docker run -v $(pwd):/app codyrao/code2md -p /app -v
```

### 参数说明

- `-v $(pwd):/app`: 将当前目录挂载到容器的 /app 目录
- `-v $(pwd):/output`: 将输出目录挂载到容器的 /output 目录（可选）
- `-p /app`: 指定要分析的项目路径
- `-o /output/code.md`: 指定输出文件路径

### 示例

```bash
# 在项目根目录下执行
docker run -v $(pwd):/app codyrao/code2md

# 分析其他目录
docker run -v /home/user/myproject:/project codyrao/code2md -p /project -o /project/docs.md

# 带详细输出
docker run -v $(pwd):/app codyrao/code2md -p /app -v -o docs.md
```

## 生成文档示例

Code2MD 生成的 markdown 文档包含以下部分：

### 1. 项目概览

```markdown
# ProjectName - Code Documentation

## 项目概览

**生成时间**: 2024-01-25T12:00:00Z
**项目路径**: /path/to/project
**主要语言**: go

### 语言统计

- **go**: 15 个文件
- **python**: 8 个文件

### 项目摘要

| 指标 | 数量 |
|------|------|
| 总文件数 | 23 |
| 总代码行数 | 5000 |
| 函数数量 | 150 |
| 类数量 | 30 |
```

### 2. 项目结构

```markdown
## 项目结构

```
project/
├── src/
│   ├── main.go
│   ├── config/
│   └── utils/
├── tests/
└── docs/
```
```

### 3. 代码元素文档

```markdown
### src/main.go

**Function**: `main`

- **完整名称**: main
- **文件位置**: 行 10-25
- **可见性**: public
- **描述**: 程序入口点

**Parameters**:
- `args`: []string

**Calls**: init(), parseConfig()
```

### 4. API 参考

```markdown
## API 参考

### Function: initialize

```go
func Initialize(config *Config) error
```

初始化系统配置...

**Parameters**:
- `config`: *Config

**Returns**:
- error
```

## 支持的语言

| 语言 | 扩展名 | 解析器 |
|------|--------|--------|
| Go | .go | go/ast |
| JavaScript | .js, .jsx, .mjs | 正则表达式解析 |
| TypeScript | .ts, .tsx | 正则表达式解析 |
| Java | .java | 正则表达式解析 |
| Python | .py, .pyw | 正则表达式解析 |
| C++ | .cpp, .cc, .cxx, .hpp, .h | 正则表达式解析 |
| PHP | .php, .phtml | 正则表达式解析 |

## 排除的目录

默认情况下，Code2MD 会排除以下目录：
- `node_modules`
- `vendor`
- `.git`
- `.svn`
- `__pycache__`
- `.idea`
- `.vscode`
- `dist`
- `build`
- `target`
- `out`
- `obj`
- `bin`

## 高级用法

### 在 CI/CD 中使用

```yaml
# GitHub Actions 示例
- name: Generate Code Documentation
  run: |
    docker run -v ${{ github.workspace }}:/app codyrao/code2md -p /app
    git add code.md
    git commit -m "Update code documentation"
```

### 忽略特定文件

创建 `.code2mdignore` 文件：

```
# 忽略测试文件
*_test.go

# 忽略生成的文件
*.pb.go
*.generated.go

# 忽略特定目录
vendor/*
```

## 贡献

欢迎贡献代码！请提交 Issue 或 Pull Request。

## 许可证

MIT License
