
# gorm - Code Documentation

## 项目概览

**生成时间**: 2026-01-26T00:05:30+08:00
**项目路径**: E:\cody\goproject\qcoder\mcp-server\code2md\test_projects\gorm
**主要语言**: 

### 语言统计


- **go**: 66 个文件


### 项目摘要

| 指标 | 数量 |
|------|------|
| 总文件数 | 66 |
| 总代码行数 | 12141 |
| 总注释行数 | 0 |
| 函数数量 | 66 |
| 类数量 | 111 |
| 接口数量 | 62 |
| 变量数量 | 303 |
| 常量数量 | 77 |

## 项目结构

├── callbacks/
│   ├── associations.go
│   ├── callbacks.go
│   ├── callmethod.go
│   ├── create.go
│   ├── delete.go
│   ├── helper.go
│   ├── interfaces.go
│   ├── preload.go
│   ├── query.go
│   ├── raw.go
│   ├── row.go
│   ├── transaction.go
│   └── update.go
├── clause/
│   ├── association.go
│   ├── clause.go
│   ├── delete.go
│   ├── expression.go
│   ├── from.go
│   ├── group_by.go
│   ├── insert.go
│   ├── joins.go
│   ├── limit.go
│   ├── locking.go
│   ├── on_conflict.go
│   ├── order_by.go
│   ├── returning.go
│   ├── select.go
│   ├── set.go
│   ├── update.go
│   ├── values.go
│   ├── where.go
│   └── with.go
├── internal/
│   ├── lru/
│   │   └── lru.go
│   └── stmt_store/
│       └── stmt_store.go
├── logger/
│   ├── logger.go
│   ├── slog.go
│   └── sql.go
├── migrator/
│   ├── column_type.go
│   ├── index.go
│   ├── migrator.go
│   └── table_type.go
├── schema/
│   ├── constraint.go
│   ├── field.go
│   ├── index.go
│   ├── interfaces.go
│   ├── naming.go
│   ├── pool.go
│   ├── relationship.go
│   ├── schema.go
│   ├── serializer.go
│   └── utils.go
├── utils/
│   └── utils.go
├── association.go
├── callbacks.go
├── chainable_api.go
├── errors.go
├── finisher_api.go
├── generics.go
├── gorm.go
├── interfaces.go
├── migrator.go
├── model.go
├── prepare_stmt.go
├── scan.go
├── soft_delete.go
└── statement.go


## 文件清单


### association.go

- **语言**: go
- **包**: gorm
- **代码行数**: 572
- **元素数量**: 44


#### 导入


- `fmt`

- `reflect`

- `strings`

- `gorm.io/gorm/clause`

- `gorm.io/gorm/schema`

- `gorm.io/gorm/utils`




#### 代码元素


**Struct**: `Association`

- **完整名称**: Association
- **文件位置**: 行 14-19
- **可见性**: public








- **子元素**: DB Relationship Unscope Error 

**Method**: `Association`

- **完整名称**: Association
- **文件位置**: 行 21-40
- **可见性**: public


- **返回值**: `*Association`

- **参数**:
  - `column`: string






**Method**: `Unscoped`

- **完整名称**: Unscoped
- **文件位置**: 行 42-49
- **可见性**: public


- **返回值**: `*Association`






**Method**: `Find`

- **完整名称**: Find
- **文件位置**: 行 51-56
- **可见性**: public


- **返回值**: `error`

- **参数**:
  - `out`: interface{}
  - `conds`: ...interface{}






**Method**: `Append`

- **完整名称**: Append
- **文件位置**: 行 58-73
- **可见性**: public


- **返回值**: `error`

- **参数**:
  - `values`: ...interface{}






**Method**: `Replace`

- **完整名称**: Replace
- **文件位置**: 行 75-197
- **可见性**: public


- **返回值**: `error`

- **参数**:
  - `values`: ...interface{}






**Variable**: `oldBelongsToExpr`

- **完整名称**: oldBelongsToExpr
- **文件位置**: 行 82-82
- **可见性**: private

- **类型**: `clause.Expression`







**Variable**: `foreignFields`

- **完整名称**: foreignFields
- **文件位置**: 行 85-85
- **可见性**: private

- **类型**: `[]*schema.Field`







**Variable**: `primaryFields`

- **完整名称**: primaryFields
- **文件位置**: 行 127-127
- **可见性**: private

- **类型**: `[]*schema.Field`







**Variable**: `foreignKeys`

- **完整名称**: foreignKeys
- **文件位置**: 行 128-128
- **可见性**: private

- **类型**: `[]string`







**Variable**: `updateMap`

- **完整名称**: updateMap
- **文件位置**: 行 129-129
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...`


**Variable**: `relValues`

- **完整名称**: relValues
- **文件位置**: 行 130-130
- **可见性**: private

- **类型**: `interface{}`




- **值**: `schema.GetRelationsValues(...)`


**Variable**: `modelValue`

- **完整名称**: modelValue
- **文件位置**: 行 131-131
- **可见性**: private

- **类型**: `interface{}`




- **值**: `reflect.New(...).Interface(...)`


**Variable**: `tx`

- **完整名称**: tx
- **文件位置**: 行 132-132
- **可见性**: private

- **类型**: `interface{}`




- **值**: `association.DB.Model(...)`


**Variable**: `primaryFields`

- **完整名称**: primaryFields
- **文件位置**: 行 161-161
- **可见性**: private

- **类型**: `[]*schema.Field`







**Variable**: `relPrimaryFields`

- **完整名称**: relPrimaryFields
- **文件位置**: 行 161-161
- **可见性**: private

- **类型**: `[]*schema.Field`







**Variable**: `joinPrimaryKeys`

- **完整名称**: joinPrimaryKeys
- **文件位置**: 行 162-162
- **可见性**: private

- **类型**: `[]string`







**Variable**: `joinRelPrimaryKeys`

- **完整名称**: joinRelPrimaryKeys
- **文件位置**: 行 162-162
- **可见性**: private

- **类型**: `[]string`







**Variable**: `modelValue`

- **完整名称**: modelValue
- **文件位置**: 行 163-163
- **可见性**: private

- **类型**: `interface{}`




- **值**: `reflect.New(...).Interface(...)`


**Variable**: `tx`

- **完整名称**: tx
- **文件位置**: 行 164-164
- **可见性**: private

- **类型**: `interface{}`




- **值**: `association.DB.Model(...)`


**Method**: `Delete`

- **完整名称**: Delete
- **文件位置**: 行 199-365
- **可见性**: public


- **返回值**: `error`

- **参数**:
  - `values`: ...interface{}






**Variable**: `reflectValue`

- **完整名称**: reflectValue
- **文件位置**: 行 204-204
- **可见性**: private

- **类型**: `association.DB.Statement.ReflectValue`




- **值**: `association.DB.Statement.ReflectValue`


**Variable**: `rel`

- **完整名称**: rel
- **文件位置**: 行 205-205
- **可见性**: private

- **类型**: `association.Relationship`




- **值**: `association.Relationship`


**Variable**: `primaryFields`

- **完整名称**: primaryFields
- **文件位置**: 行 206-206
- **可见性**: private

- **类型**: `[]*schema.Field`







**Variable**: `foreignKeys`

- **完整名称**: foreignKeys
- **文件位置**: 行 207-207
- **可见性**: private

- **类型**: `[]string`







**Variable**: `updateAttrs`

- **完整名称**: updateAttrs
- **文件位置**: 行 208-208
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...`


**Variable**: `conds`

- **完整名称**: conds
- **文件位置**: 行 209-209
- **可见性**: private

- **类型**: `[]clause.Expression`







**Variable**: `foreignFields`

- **完整名称**: foreignFields
- **文件位置**: 行 240-240
- **可见性**: private

- **类型**: `[]*schema.Field`







**Variable**: `primaryFields`

- **完整名称**: primaryFields
- **文件位置**: 行 273-273
- **可见性**: private

- **类型**: `[]*schema.Field`







**Variable**: `relPrimaryFields`

- **完整名称**: relPrimaryFields
- **文件位置**: 行 273-273
- **可见性**: private

- **类型**: `[]*schema.Field`







**Variable**: `joinPrimaryKeys`

- **完整名称**: joinPrimaryKeys
- **文件位置**: 行 274-274
- **可见性**: private

- **类型**: `[]string`







**Variable**: `joinRelPrimaryKeys`

- **完整名称**: joinRelPrimaryKeys
- **文件位置**: 行 274-274
- **可见性**: private

- **类型**: `[]string`







**Variable**: `joinValue`

- **完整名称**: joinValue
- **文件位置**: 行 275-275
- **可见性**: private

- **类型**: `interface{}`




- **值**: `reflect.New(...).Interface(...)`


**Method**: `Clear`

- **完整名称**: Clear
- **文件位置**: 行 367-369
- **可见性**: public


- **返回值**: `error`






**Method**: `Count`

- **完整名称**: Count
- **文件位置**: 行 371-376
- **可见性**: public


- **返回值**: `int64`






**Struct**: `assignBack`

- **完整名称**: assignBack
- **文件位置**: 行 378-382
- **可见性**: private








- **子元素**: Source Index Dest 

**Method**: `saveAssociation`

- **完整名称**: saveAssociation
- **文件位置**: 行 384-606
- **可见性**: private




- **参数**:
  - `clear`: bool
  - `values`: ...interface{}






**Variable**: `reflectValue`

- **完整名称**: reflectValue
- **文件位置**: 行 386-386
- **可见性**: private

- **类型**: `association.DB.Statement.ReflectValue`




- **值**: `association.DB.Statement.ReflectValue`


**Variable**: `assignBacks`

- **完整名称**: assignBacks
- **文件位置**: 行 387-387
- **可见性**: private

- **类型**: `[]assignBack`







**Variable**: `fieldValue`

- **完整名称**: fieldValue
- **文件位置**: 行 416-416
- **可见性**: private

- **类型**: `reflect.Value`







**Method**: `buildCondition`

- **完整名称**: buildCondition
- **文件位置**: 行 608-636
- **可见性**: private


- **返回值**: `*DB`






**Variable**: `queryConds`

- **完整名称**: queryConds
- **文件位置**: 行 610-610
- **可见性**: private

- **类型**: `interface{}`




- **值**: `association.Relationship.ToQueryConditions(...)`


**Variable**: `modelValue`

- **完整名称**: modelValue
- **文件位置**: 行 611-611
- **可见性**: private

- **类型**: `interface{}`




- **值**: `reflect.New(...).Interface(...)`


**Variable**: `tx`

- **完整名称**: tx
- **文件位置**: 行 612-612
- **可见性**: private

- **类型**: `interface{}`




- **值**: `association.DB.Model(...)`





### callbacks\associations.go

- **语言**: go
- **包**: callbacks
- **代码行数**: 381
- **元素数量**: 14


#### 导入


- `reflect`

- `strings`

- `gorm.io/gorm`

- `gorm.io/gorm/clause`

- `gorm.io/gorm/schema`

- `gorm.io/gorm/utils`




#### 代码元素


**Function**: `SaveBeforeAssociations`

- **完整名称**: SaveBeforeAssociations
- **文件位置**: 行 13-108
- **可见性**: public


- **返回值**: `func(...)`

- **参数**:
  - `create`: bool






**Variable**: `rValLen`

- **完整名称**: rValLen
- **文件位置**: 行 43-43
- **可见性**: private

- **类型**: `interface{}`




- **值**: `db.Statement.ReflectValue.Len(...)`


**Variable**: `objs`

- **完整名称**: objs
- **文件位置**: 行 44-44
- **可见性**: private

- **类型**: `interface{}`




- **值**: `make(...)`


**Variable**: `fieldType`

- **完整名称**: fieldType
- **文件位置**: 行 45-45
- **可见性**: private

- **类型**: `rel.Field.FieldType`




- **值**: `rel.Field.FieldType`


**Variable**: `isPtr`

- **完整名称**: isPtr
- **文件位置**: 行 46-46
- **可见性**: private

- **类型**: `interface{}`




- **值**: `fieldType.Kind(...) == reflect.Ptr`


**Function**: `SaveAfterAssociations`

- **完整名称**: SaveAfterAssociations
- **文件位置**: 行 110-358
- **可见性**: public


- **返回值**: `func(...)`

- **参数**:
  - `create`: bool






**Variable**: `fieldType`

- **完整名称**: fieldType
- **文件位置**: 行 124-124
- **可见性**: private

- **类型**: `rel.Field.FieldType`




- **值**: `rel.Field.FieldType`


**Variable**: `isPtr`

- **完整名称**: isPtr
- **文件位置**: 行 125-125
- **可见性**: private

- **类型**: `interface{}`




- **值**: `fieldType.Kind(...) == reflect.Ptr`


**Variable**: `selects`

- **完整名称**: selects
- **文件位置**: 行 385-385
- **可见性**: private

- **类型**: `[]string`







**Variable**: `omits`

- **完整名称**: omits
- **文件位置**: 行 385-385
- **可见性**: private

- **类型**: `[]string`







**Variable**: `onConflict`

- **完整名称**: onConflict
- **文件位置**: 行 386-386
- **可见性**: private

- **类型**: `interface{}`




- **值**: `onConflictOption(...)`


**Variable**: `refName`

- **完整名称**: refName
- **文件位置**: 行 387-387
- **可见性**: private

- **类型**: `interface{}`




- **值**: `rel.Name + "."`


**Variable**: `values`

- **完整名称**: values
- **文件位置**: 行 388-388
- **可见性**: private

- **类型**: `interface{}`




- **值**: `rValues.Interface(...)`


**Variable**: `visitMapStoreKey`

- **完整名称**: visitMapStoreKey
- **文件位置**: 行 437-437
- **可见性**: private

- **类型**: `interface{}`




- **值**: `"gorm:saved_association_map"`





### callbacks\callbacks.go

- **语言**: go
- **包**: callbacks
- **代码行数**: 72
- **元素数量**: 6


#### 导入


- `gorm.io/gorm`




#### 代码元素


**Variable**: `createClauses`

- **完整名称**: createClauses
- **文件位置**: 行 8-8
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...`


**Variable**: `queryClauses`

- **完整名称**: queryClauses
- **文件位置**: 行 9-9
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...`


**Variable**: `updateClauses`

- **完整名称**: updateClauses
- **文件位置**: 行 10-10
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...`


**Variable**: `deleteClauses`

- **完整名称**: deleteClauses
- **文件位置**: 行 11-11
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...`


**Struct**: `Config`

- **完整名称**: Config
- **文件位置**: 行 14-20
- **可见性**: public








- **子元素**: LastInsertIDReversed CreateClauses QueryClauses UpdateClauses DeleteClauses 

**Function**: `RegisterDefaultCallbacks`

- **完整名称**: RegisterDefaultCallbacks
- **文件位置**: 行 22-83
- **可见性**: public




- **参数**:
  - `db`: *gorm.DB
  - `config`: *Config









### callbacks\callmethod.go

- **语言**: go
- **包**: callbacks
- **代码行数**: 29
- **元素数量**: 0


#### 导入


- `reflect`

- `gorm.io/gorm`





### callbacks\create.go

- **语言**: go
- **包**: callbacks
- **代码行数**: 360
- **元素数量**: 11


#### 导入


- `fmt`

- `reflect`

- `strings`

- `gorm.io/gorm`

- `gorm.io/gorm/clause`

- `gorm.io/gorm/schema`

- `gorm.io/gorm/utils`




#### 代码元素


**Function**: `BeforeCreate`

- **完整名称**: BeforeCreate
- **文件位置**: 行 15-34
- **可见性**: public
- **描述**: BeforeCreate before create hooks



- **参数**:
  - `db`: *gorm.DB






**Function**: `Create`

- **完整名称**: Create
- **文件位置**: 行 37-220
- **可见性**: public
- **描述**: Create create hook

- **返回值**: `func(...)`

- **参数**:
  - `config`: *Config






**Variable**: `pkField`

- **完整名称**: pkField
- **文件位置**: 行 128-128
- **可见性**: private

- **类型**: `*schema.Field`







**Variable**: `pkFieldName`

- **完整名称**: pkFieldName
- **文件位置**: 行 129-129
- **可见性**: private

- **类型**: `interface{}`




- **值**: `"@id"`


**Function**: `AfterCreate`

- **完整名称**: AfterCreate
- **文件位置**: 行 223-242
- **可见性**: public
- **描述**: AfterCreate after create hooks



- **参数**:
  - `db`: *gorm.DB






**Function**: `ConvertToCreateValues`

- **完整名称**: ConvertToCreateValues
- **文件位置**: 行 245-415
- **可见性**: public
- **描述**: ConvertToCreateValues convert to create values

- **返回值**: `clause.Values`

- **参数**:
  - `stmt`: *gorm.Statement






**Variable**: `selectColumns`

- **完整名称**: selectColumns
- **文件位置**: 行 259-259
- **可见性**: private

- **类型**: `interface{}`




- **值**: `stmt.SelectAndOmitColumns(...)`


**Variable**: `restricted`

- **完整名称**: restricted
- **文件位置**: 行 259-259
- **可见性**: private

- **类型**: `interface{}`




- **值**: `stmt.SelectAndOmitColumns(...)`


**Variable**: `_`

- **完整名称**: _
- **文件位置**: 行 260-260
- **可见性**: private

- **类型**: `interface{}`




- **值**: `stmt.Get(...)`


**Variable**: `updateTrackTime`

- **完整名称**: updateTrackTime
- **文件位置**: 行 260-260
- **可见性**: private

- **类型**: `interface{}`




- **值**: `stmt.Get(...)`


**Variable**: `isZero`

- **完整名称**: isZero
- **文件位置**: 行 261-261
- **可见性**: private

- **类型**: `bool`










### callbacks\delete.go

- **语言**: go
- **包**: callbacks
- **代码行数**: 161
- **元素数量**: 10


#### 导入


- `reflect`

- `strings`

- `gorm.io/gorm`

- `gorm.io/gorm/clause`

- `gorm.io/gorm/schema`

- `gorm.io/gorm/utils`




#### 代码元素


**Function**: `BeforeDelete`

- **完整名称**: BeforeDelete
- **文件位置**: 行 13-24
- **可见性**: public




- **参数**:
  - `db`: *gorm.DB






**Function**: `DeleteBeforeAssociations`

- **完整名称**: DeleteBeforeAssociations
- **文件位置**: 行 26-111
- **可见性**: public




- **参数**:
  - `db`: *gorm.DB






**Variable**: `queryConds`

- **完整名称**: queryConds
- **文件位置**: 行 80-80
- **可见性**: private

- **类型**: `interface{}`




- **值**: `make(...)`


**Variable**: `foreignFields`

- **完整名称**: foreignFields
- **文件位置**: 行 81-81
- **可见性**: private

- **类型**: `interface{}`




- **值**: `make(...)`


**Variable**: `relForeignKeys`

- **完整名称**: relForeignKeys
- **文件位置**: 行 82-82
- **可见性**: private

- **类型**: `interface{}`




- **值**: `make(...)`


**Variable**: `modelValue`

- **完整名称**: modelValue
- **文件位置**: 行 83-83
- **可见性**: private

- **类型**: `interface{}`




- **值**: `reflect.New(...).Interface(...)`


**Variable**: `table`

- **完整名称**: table
- **文件位置**: 行 84-84
- **可见性**: private

- **类型**: `rel.JoinTable.Table`




- **值**: `rel.JoinTable.Table`


**Variable**: `tx`

- **完整名称**: tx
- **文件位置**: 行 85-85
- **可见性**: private

- **类型**: `interface{}`




- **值**: `db.Session(...).Model(...).Table(...)`


**Function**: `Delete`

- **完整名称**: Delete
- **文件位置**: 行 113-183
- **可见性**: public


- **返回值**: `func(...)`

- **参数**:
  - `config`: *Config






**Function**: `AfterDelete`

- **完整名称**: AfterDelete
- **文件位置**: 行 185-195
- **可见性**: public




- **参数**:
  - `db`: *gorm.DB









### callbacks\helper.go

- **语言**: go
- **包**: callbacks
- **代码行数**: 125
- **元素数量**: 6


#### 导入


- `reflect`

- `sort`

- `gorm.io/gorm`

- `gorm.io/gorm/clause`




#### 代码元素


**Function**: `ConvertMapToValuesForCreate`

- **完整名称**: ConvertMapToValuesForCreate
- **文件位置**: 行 12-40
- **可见性**: public
- **描述**: ConvertMapToValuesForCreate convert map to values

- **返回值**: `clause.Values`

- **参数**:
  - `stmt`: *gorm.Statement
  - `mapValue`: map[string]interface{}






**Function**: `ConvertSliceOfMapToValuesForCreate`

- **完整名称**: ConvertSliceOfMapToValuesForCreate
- **文件位置**: 行 43-94
- **可见性**: public
- **描述**: ConvertSliceOfMapToValuesForCreate convert slice of map to values

- **返回值**: `clause.Values`

- **参数**:
  - `stmt`: *gorm.Statement
  - `mapValues`: []map[string]interface{}






**Variable**: `result`

- **完整名称**: result
- **文件位置**: 行 54-54
- **可见性**: private

- **类型**: `interface{}`




- **值**: `make(...)`


**Variable**: `selectColumns`

- **完整名称**: selectColumns
- **文件位置**: 行 55-55
- **可见性**: private

- **类型**: `interface{}`




- **值**: `stmt.SelectAndOmitColumns(...)`


**Variable**: `restricted`

- **完整名称**: restricted
- **文件位置**: 行 55-55
- **可见性**: private

- **类型**: `interface{}`




- **值**: `stmt.SelectAndOmitColumns(...)`


**Class**: `visitMap`

- **完整名称**: visitMap
- **文件位置**: 行 125-125
- **可见性**: private












### callbacks\interfaces.go

- **语言**: go
- **包**: callbacks
- **代码行数**: 29
- **元素数量**: 9


#### 导入


- `gorm.io/gorm`




#### 代码元素


**Interface**: `BeforeCreateInterface`

- **完整名称**: BeforeCreateInterface
- **文件位置**: 行 5-7
- **可见性**: public








- **子元素**: BeforeCreate 

**Interface**: `AfterCreateInterface`

- **完整名称**: AfterCreateInterface
- **文件位置**: 行 9-11
- **可见性**: public








- **子元素**: AfterCreate 

**Interface**: `BeforeUpdateInterface`

- **完整名称**: BeforeUpdateInterface
- **文件位置**: 行 13-15
- **可见性**: public








- **子元素**: BeforeUpdate 

**Interface**: `AfterUpdateInterface`

- **完整名称**: AfterUpdateInterface
- **文件位置**: 行 17-19
- **可见性**: public








- **子元素**: AfterUpdate 

**Interface**: `BeforeSaveInterface`

- **完整名称**: BeforeSaveInterface
- **文件位置**: 行 21-23
- **可见性**: public








- **子元素**: BeforeSave 

**Interface**: `AfterSaveInterface`

- **完整名称**: AfterSaveInterface
- **文件位置**: 行 25-27
- **可见性**: public








- **子元素**: AfterSave 

**Interface**: `BeforeDeleteInterface`

- **完整名称**: BeforeDeleteInterface
- **文件位置**: 行 29-31
- **可见性**: public








- **子元素**: BeforeDelete 

**Interface**: `AfterDeleteInterface`

- **完整名称**: AfterDeleteInterface
- **文件位置**: 行 33-35
- **可见性**: public








- **子元素**: AfterDelete 

**Interface**: `AfterFindInterface`

- **完整名称**: AfterFindInterface
- **文件位置**: 行 37-39
- **可见性**: public








- **子元素**: AfterFind 




### callbacks\preload.go

- **语言**: go
- **包**: callbacks
- **代码行数**: 283
- **元素数量**: 10


#### 导入


- `fmt`

- `reflect`

- `sort`

- `strings`

- `gorm.io/gorm`

- `gorm.io/gorm/clause`

- `gorm.io/gorm/schema`

- `gorm.io/gorm/utils`




#### 代码元素


**Variable**: `reflectValue`

- **完整名称**: reflectValue
- **文件位置**: 行 187-187
- **可见性**: private

- **类型**: `tx.Statement.ReflectValue`




- **值**: `tx.Statement.ReflectValue`


**Variable**: `relForeignKeys`

- **完整名称**: relForeignKeys
- **文件位置**: 行 188-188
- **可见性**: private

- **类型**: `[]string`







**Variable**: `relForeignFields`

- **完整名称**: relForeignFields
- **文件位置**: 行 189-189
- **可见性**: private

- **类型**: `[]*schema.Field`







**Variable**: `foreignFields`

- **完整名称**: foreignFields
- **文件位置**: 行 190-190
- **可见性**: private

- **类型**: `[]*schema.Field`







**Variable**: `foreignValues`

- **完整名称**: foreignValues
- **文件位置**: 行 191-191
- **可见性**: private

- **类型**: `[][]interface{}`







**Variable**: `identityMap`

- **完整名称**: identityMap
- **文件位置**: 行 192-192
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...`


**Variable**: `inlineConds`

- **完整名称**: inlineConds
- **文件位置**: 行 193-193
- **可见性**: private

- **类型**: `[]interface{}`







**Variable**: `joinForeignFields`

- **完整名称**: joinForeignFields
- **文件位置**: 行 198-198
- **可见性**: private

- **类型**: `interface{}`




- **值**: `make(...)`


**Variable**: `joinRelForeignFields`

- **完整名称**: joinRelForeignFields
- **文件位置**: 行 199-199
- **可见性**: private

- **类型**: `interface{}`




- **值**: `make(...)`


**Variable**: `joinForeignKeys`

- **完整名称**: joinForeignKeys
- **文件位置**: 行 200-200
- **可见性**: private

- **类型**: `interface{}`




- **值**: `make(...)`





### callbacks\query.go

- **语言**: go
- **包**: callbacks
- **代码行数**: 270
- **元素数量**: 7


#### 导入


- `fmt`

- `reflect`

- `strings`

- `gorm.io/gorm`

- `gorm.io/gorm/clause`

- `gorm.io/gorm/schema`

- `gorm.io/gorm/utils`




#### 代码元素


**Function**: `Query`

- **完整名称**: Query
- **文件位置**: 行 14-34
- **可见性**: public




- **参数**:
  - `db`: *gorm.DB






**Function**: `BuildQuerySQL`

- **完整名称**: BuildQuerySQL
- **文件位置**: 行 36-275
- **可见性**: public




- **参数**:
  - `db`: *gorm.DB






**Variable**: `conds`

- **完整名称**: conds
- **文件位置**: 行 48-48
- **可见性**: private

- **类型**: `[]clause.Expression`







**Variable**: `isRelations`

- **完整名称**: isRelations
- **文件位置**: 行 120-120
- **可见性**: private

- **类型**: `bool`







**Variable**: `relations`

- **完整名称**: relations
- **文件位置**: 行 121-121
- **可见性**: private

- **类型**: `[]*schema.Relationship`







**Function**: `Preload`

- **完整名称**: Preload
- **文件位置**: 行 277-296
- **可见性**: public




- **参数**:
  - `db`: *gorm.DB






**Function**: `AfterQuery`

- **完整名称**: AfterQuery
- **文件位置**: 行 298-314
- **可见性**: public




- **参数**:
  - `db`: *gorm.DB









### callbacks\raw.go

- **语言**: go
- **包**: callbacks
- **代码行数**: 18
- **元素数量**: 1


#### 导入


- `gorm.io/gorm`




#### 代码元素


**Function**: `RawExec`

- **完整名称**: RawExec
- **文件位置**: 行 7-22
- **可见性**: public




- **参数**:
  - `db`: *gorm.DB









### callbacks\row.go

- **语言**: go
- **包**: callbacks
- **代码行数**: 19
- **元素数量**: 1


#### 导入


- `gorm.io/gorm`




#### 代码元素


**Function**: `RowQuery`

- **完整名称**: RowQuery
- **文件位置**: 行 7-23
- **可见性**: public




- **参数**:
  - `db`: *gorm.DB









### callbacks\transaction.go

- **语言**: go
- **包**: callbacks
- **代码行数**: 28
- **元素数量**: 2


#### 导入


- `gorm.io/gorm`




#### 代码元素


**Function**: `BeginTransaction`

- **完整名称**: BeginTransaction
- **文件位置**: 行 7-18
- **可见性**: public




- **参数**:
  - `db`: *gorm.DB






**Function**: `CommitOrRollbackTransaction`

- **完整名称**: CommitOrRollbackTransaction
- **文件位置**: 行 20-32
- **可见性**: public




- **参数**:
  - `db`: *gorm.DB









### callbacks\update.go

- **语言**: go
- **包**: callbacks
- **代码行数**: 273
- **元素数量**: 10


#### 导入


- `reflect`

- `sort`

- `gorm.io/gorm`

- `gorm.io/gorm/clause`

- `gorm.io/gorm/schema`

- `gorm.io/gorm/utils`




#### 代码元素


**Function**: `SetupUpdateReflectValue`

- **完整名称**: SetupUpdateReflectValue
- **文件位置**: 行 13-30
- **可见性**: public




- **参数**:
  - `db`: *gorm.DB






**Function**: `BeforeUpdate`

- **完整名称**: BeforeUpdate
- **文件位置**: 行 33-53
- **可见性**: public
- **描述**: BeforeUpdate before update hooks



- **参数**:
  - `db`: *gorm.DB






**Function**: `Update`

- **完整名称**: Update
- **文件位置**: 行 56-114
- **可见性**: public
- **描述**: Update update hook

- **返回值**: `func(...)`

- **参数**:
  - `config`: *Config






**Function**: `AfterUpdate`

- **完整名称**: AfterUpdate
- **文件位置**: 行 117-137
- **可见性**: public
- **描述**: AfterUpdate after update hooks



- **参数**:
  - `db`: *gorm.DB






**Function**: `ConvertToAssignments`

- **完整名称**: ConvertToAssignments
- **文件位置**: 行 140-313
- **可见性**: public
- **描述**: ConvertToAssignments convert to update assignments

- **返回值**: `clause.Set`

- **参数**:
  - `stmt`: *gorm.Statement






**Variable**: `selectColumns`

- **完整名称**: selectColumns
- **文件位置**: 行 142-142
- **可见性**: private

- **类型**: `interface{}`




- **值**: `stmt.SelectAndOmitColumns(...)`


**Variable**: `restricted`

- **完整名称**: restricted
- **文件位置**: 行 142-142
- **可见性**: private

- **类型**: `interface{}`




- **值**: `stmt.SelectAndOmitColumns(...)`


**Variable**: `assignValue`

- **完整名称**: assignValue
- **文件位置**: 行 143-143
- **可见性**: private

- **类型**: `func(...)`







**Variable**: `isZero`

- **完整名称**: isZero
- **文件位置**: 行 175-175
- **可见性**: private

- **类型**: `bool`







**Variable**: `isDiffSchema`

- **完整名称**: isDiffSchema
- **文件位置**: 行 258-258
- **可见性**: private

- **类型**: `bool`










### callbacks.go

- **语言**: go
- **包**: gorm
- **代码行数**: 298
- **元素数量**: 30


#### 导入


- `context`

- `errors`

- `fmt`

- `reflect`

- `sort`

- `time`

- `gorm.io/gorm/schema`

- `gorm.io/gorm/utils`




#### 代码元素


**Struct**: `callbacks`

- **完整名称**: callbacks
- **文件位置**: 行 29-31
- **可见性**: private








- **子元素**: processors 

**Struct**: `processor`

- **完整名称**: processor
- **文件位置**: 行 33-38
- **可见性**: private








- **子元素**: db Clauses fns callbacks 

**Struct**: `callback`

- **完整名称**: callback
- **文件位置**: 行 40-49
- **可见性**: private








- **子元素**: name before after remove replace match handler processor 

**Method**: `Create`

- **完整名称**: Create
- **文件位置**: 行 51-53
- **可见性**: public


- **返回值**: `*processor`






**Method**: `Query`

- **完整名称**: Query
- **文件位置**: 行 55-57
- **可见性**: public


- **返回值**: `*processor`






**Method**: `Update`

- **完整名称**: Update
- **文件位置**: 行 59-61
- **可见性**: public


- **返回值**: `*processor`






**Method**: `Delete`

- **完整名称**: Delete
- **文件位置**: 行 63-65
- **可见性**: public


- **返回值**: `*processor`






**Method**: `Row`

- **完整名称**: Row
- **文件位置**: 行 67-69
- **可见性**: public


- **返回值**: `*processor`






**Method**: `Raw`

- **完整名称**: Raw
- **文件位置**: 行 71-73
- **可见性**: public


- **返回值**: `*processor`






**Method**: `Execute`

- **完整名称**: Execute
- **文件位置**: 行 75-159
- **可见性**: public


- **返回值**: `*DB`

- **参数**:
  - `db`: *DB






**Variable**: `curTime`

- **完整名称**: curTime
- **文件位置**: 行 82-82
- **可见性**: private

- **类型**: `interface{}`




- **值**: `time.Now(...)`


**Variable**: `stmt`

- **完整名称**: stmt
- **文件位置**: 行 83-83
- **可见性**: private

- **类型**: `db.Statement`




- **值**: `db.Statement`


**Variable**: `resetBuildClauses`

- **完整名称**: resetBuildClauses
- **文件位置**: 行 84-84
- **可见性**: private

- **类型**: `bool`







**Method**: `Get`

- **完整名称**: Get
- **文件位置**: 行 161-168
- **可见性**: public


- **返回值**: `func(...)`

- **参数**:
  - `name`: string






**Method**: `Before`

- **完整名称**: Before
- **文件位置**: 行 170-172
- **可见性**: public


- **返回值**: `*callback`

- **参数**:
  - `name`: string






**Method**: `After`

- **完整名称**: After
- **文件位置**: 行 174-176
- **可见性**: public


- **返回值**: `*callback`

- **参数**:
  - `name`: string






**Method**: `Match`

- **完整名称**: Match
- **文件位置**: 行 178-180
- **可见性**: public


- **返回值**: `*callback`

- **参数**:
  - `fc`: func(...)






**Method**: `Register`

- **完整名称**: Register
- **文件位置**: 行 182-184
- **可见性**: public


- **返回值**: `error`

- **参数**:
  - `name`: string
  - `fn`: func(...)






**Method**: `Remove`

- **完整名称**: Remove
- **文件位置**: 行 186-188
- **可见性**: public


- **返回值**: `error`

- **参数**:
  - `name`: string






**Method**: `Replace`

- **完整名称**: Replace
- **文件位置**: 行 190-192
- **可见性**: public


- **返回值**: `error`

- **参数**:
  - `name`: string
  - `fn`: func(...)






**Method**: `compile`

- **完整名称**: compile
- **文件位置**: 行 194-215
- **可见性**: private


- **返回值**: `error`






**Variable**: `callbacks`

- **完整名称**: callbacks
- **文件位置**: 行 195-195
- **可见性**: private

- **类型**: `[]*callback`







**Method**: `Before`

- **完整名称**: Before
- **文件位置**: 行 217-220
- **可见性**: public


- **返回值**: `*callback`

- **参数**:
  - `name`: string






**Method**: `After`

- **完整名称**: After
- **文件位置**: 行 222-225
- **可见性**: public


- **返回值**: `*callback`

- **参数**:
  - `name`: string






**Method**: `Register`

- **完整名称**: Register
- **文件位置**: 行 227-232
- **可见性**: public


- **返回值**: `error`

- **参数**:
  - `name`: string
  - `fn`: func(...)






**Method**: `Remove`

- **完整名称**: Remove
- **文件位置**: 行 234-240
- **可见性**: public


- **返回值**: `error`

- **参数**:
  - `name`: string






**Method**: `Replace`

- **完整名称**: Replace
- **文件位置**: 行 242-249
- **可见性**: public


- **返回值**: `error`

- **参数**:
  - `name`: string
  - `fn`: func(...)






**Variable**: `names`

- **完整名称**: names
- **文件位置**: 行 263-263
- **可见性**: private

- **类型**: `[]string`







**Variable**: `sorted`

- **完整名称**: sorted
- **文件位置**: 行 263-263
- **可见性**: private

- **类型**: `[]string`







**Variable**: `sortCallback`

- **完整名称**: sortCallback
- **文件位置**: 行 264-264
- **可见性**: private

- **类型**: `func(...)`










### chainable_api.go

- **语言**: go
- **包**: gorm
- **代码行数**: 269
- **元素数量**: 26


#### 导入


- `fmt`

- `regexp`

- `strings`

- `gorm.io/gorm/clause`

- `gorm.io/gorm/utils`




#### 代码元素


**Method**: `Model`

- **完整名称**: Model
- **文件位置**: 行 18-22
- **可见性**: public
- **描述**: Model specify the model you would like to run db operations

- **返回值**: `*DB`

- **参数**:
  - `value`: interface{}






**Method**: `Clauses`

- **完整名称**: Clauses
- **文件位置**: 行 38-56
- **可见性**: public
- **描述**: Clauses Add clauses

- **返回值**: `*DB`

- **参数**:
  - `conds`: ...clause.Expression






**Variable**: `whereConds`

- **完整名称**: whereConds
- **文件位置**: 行 40-40
- **可见性**: private

- **类型**: `[]interface{}`







**Variable**: `tableRegexp`

- **完整名称**: tableRegexp
- **文件位置**: 行 58-58
- **可见性**: private

- **类型**: `interface{}`




- **值**: `regexp.MustCompile(...)`


**Method**: `Table`

- **完整名称**: Table
- **文件位置**: 行 64-86
- **可见性**: public
- **描述**: Table specify the table you would like to run db operations

- **返回值**: `*DB`

- **参数**:
  - `name`: string
  - `args`: ...interface{}






**Method**: `Distinct`

- **完整名称**: Distinct
- **文件位置**: 行 94-101
- **可见性**: public
- **描述**: Distinct specify distinct fields that you want querying

- **返回值**: `*DB`

- **参数**:
  - `args`: ...interface{}






**Method**: `Select`

- **完整名称**: Select
- **文件位置**: 行 112-174
- **可见性**: public
- **描述**: Select specify fields that you want when querying, creating, updating

- **返回值**: `*DB`

- **参数**:
  - `query`: interface{}
  - `args`: ...interface{}






**Method**: `Omit`

- **完整名称**: Omit
- **文件位置**: 行 177-186
- **可见性**: public
- **描述**: Omit specify fields that you want to ignore when creating, updating and querying

- **返回值**: `*DB`

- **参数**:
  - `columns`: ...string






**Method**: `MapColumns`

- **完整名称**: MapColumns
- **文件位置**: 行 189-193
- **可见性**: public
- **描述**: MapColumns modify the column names in the query results to facilitate align to the corresponding structural fields

- **返回值**: `*DB`

- **参数**:
  - `m`: map[string]string






**Method**: `Where`

- **完整名称**: Where
- **文件位置**: 行 207-213
- **可见性**: public
- **描述**: Where add conditions

- **返回值**: `*DB`

- **参数**:
  - `query`: interface{}
  - `args`: ...interface{}






**Method**: `Not`

- **完整名称**: Not
- **文件位置**: 行 221-227
- **可见性**: public
- **描述**: Not add NOT conditions

- **返回值**: `*DB`

- **参数**:
  - `query`: interface{}
  - `args`: ...interface{}






**Method**: `Or`

- **完整名称**: Or
- **文件位置**: 行 235-241
- **可见性**: public
- **描述**: Or add OR conditions

- **返回值**: `*DB`

- **参数**:
  - `query`: interface{}
  - `args`: ...interface{}






**Method**: `Joins`

- **完整名称**: Joins
- **文件位置**: 行 248-250
- **可见性**: public
- **描述**: Joins specify Joins conditions

- **返回值**: `*DB`

- **参数**:
  - `query`: string
  - `args`: ...interface{}






**Method**: `InnerJoins`

- **完整名称**: InnerJoins
- **文件位置**: 行 254-256
- **可见性**: public
- **描述**: InnerJoins specify inner joins conditions

- **返回值**: `*DB`

- **参数**:
  - `query`: string
  - `args`: ...interface{}






**Method**: `Group`

- **完整名称**: Group
- **文件位置**: 行 283-291
- **可见性**: public
- **描述**: Group specify the group method on the find

- **返回值**: `*DB`

- **参数**:
  - `name`: string






**Method**: `Having`

- **完整名称**: Having
- **文件位置**: 行 297-303
- **可见性**: public
- **描述**: Having specify HAVING conditions for GROUP BY

- **返回值**: `*DB`

- **参数**:
  - `query`: interface{}
  - `args`: ...interface{}






**Method**: `Order`

- **完整名称**: Order
- **文件位置**: 行 313-333
- **可见性**: public
- **描述**: Order specify order when retrieving records from database

- **返回值**: `*DB`

- **参数**:
  - `value`: interface{}






**Method**: `Limit`

- **完整名称**: Limit
- **文件位置**: 行 343-347
- **可见性**: public
- **描述**: Limit specify the number of records to be retrieved

- **返回值**: `*DB`

- **参数**:
  - `limit`: int






**Method**: `Offset`

- **完整名称**: Offset
- **文件位置**: 行 357-361
- **可见性**: public
- **描述**: Offset specify the number of records to skip before starting to return the records

- **返回值**: `*DB`

- **参数**:
  - `offset`: int






**Method**: `Scopes`

- **完整名称**: Scopes
- **文件位置**: 行 376-380
- **可见性**: public
- **描述**: Scopes pass current database connection to arguments `func(DB) DB`, which could be used to add conditions dynamically

- **返回值**: `*DB`

- **参数**:
  - `funcs`: ...func(...)






**Method**: `executeScopes`

- **完整名称**: executeScopes
- **文件位置**: 行 382-389
- **可见性**: private


- **返回值**: `*DB`






**Method**: `Preload`

- **完整名称**: Preload
- **文件位置**: 行 395-402
- **可见性**: public
- **描述**: Preload preload associations with given conditions

- **返回值**: `*DB`

- **参数**:
  - `query`: string
  - `args`: ...interface{}






**Method**: `Attrs`

- **完整名称**: Attrs
- **文件位置**: 行 418-422
- **可见性**: public
- **描述**: Attrs provide attributes used in [FirstOrCreate] or [FirstOrInit]

- **返回值**: `*DB`

- **参数**:
  - `attrs`: ...interface{}






**Method**: `Assign`

- **完整名称**: Assign
- **文件位置**: 行 439-443
- **可见性**: public
- **描述**: Assign provide attributes used in [FirstOrCreate] or [FirstOrInit]

- **返回值**: `*DB`

- **参数**:
  - `attrs`: ...interface{}






**Method**: `Unscoped`

- **完整名称**: Unscoped
- **文件位置**: 行 455-459
- **可见性**: public
- **描述**: Unscoped disables the global scope of soft deletion in a query.

- **返回值**: `*DB`






**Method**: `Raw`

- **完整名称**: Raw
- **文件位置**: 行 461-471
- **可见性**: public


- **返回值**: `*DB`

- **参数**:
  - `sql`: string
  - `values`: ...interface{}









### clause\association.go

- **语言**: go
- **包**: clause
- **代码行数**: 24
- **元素数量**: 9




#### 代码元素


**Class**: `AssociationOpType`

- **完整名称**: AssociationOpType
- **文件位置**: 行 4-4
- **可见性**: public









**Constant**: `OpUnlink`

- **完整名称**: OpUnlink
- **文件位置**: 行 7-7
- **可见性**: public

- **类型**: `AssociationOpType`




- **值**: `iota`


**Constant**: `OpDelete`

- **完整名称**: OpDelete
- **文件位置**: 行 8-8
- **可见性**: public









**Constant**: `OpUpdate`

- **完整名称**: OpUpdate
- **文件位置**: 行 9-9
- **可见性**: public









**Constant**: `OpCreate`

- **完整名称**: OpCreate
- **文件位置**: 行 10-10
- **可见性**: public









**Struct**: `Association`

- **完整名称**: Association
- **文件位置**: 行 14-20
- **可见性**: public








- **子元素**: Association Type Conditions Set Values 

**Interface**: `AssociationAssigner`

- **完整名称**: AssociationAssigner
- **文件位置**: 行 23-25
- **可见性**: public








- **子元素**: AssociationAssignments 

**Method**: `Assignments`

- **完整名称**: Assignments
- **文件位置**: 行 28-30
- **可见性**: public
- **描述**: Assignments implements the Assigner interface so that AssociationOperation can be used as a Set method parameter

- **返回值**: `[]Assignment`






**Method**: `AssociationAssignments`

- **完整名称**: AssociationAssignments
- **文件位置**: 行 33-35
- **可见性**: public
- **描述**: AssociationAssignments implements the AssociationAssigner interface

- **返回值**: `[]Association`









### clause\clause.go

- **语言**: go
- **包**: clause
- **代码行数**: 68
- **元素数量**: 13




#### 代码元素


**Interface**: `Interface`

- **完整名称**: Interface
- **文件位置**: 行 4-8
- **可见性**: public








- **子元素**: Name Build MergeClause 

**Class**: `ClauseBuilder`

- **完整名称**: ClauseBuilder
- **文件位置**: 行 11-11
- **可见性**: public









**Interface**: `Writer`

- **完整名称**: Writer
- **文件位置**: 行 13-16
- **可见性**: public








- **子元素**: WriteByte WriteString 

**Interface**: `Builder`

- **完整名称**: Builder
- **文件位置**: 行 19-24
- **可见性**: public








- **子元素**: WriteQuoted AddVar AddError 

**Struct**: `Clause`

- **完整名称**: Clause
- **文件位置**: 行 27-34
- **可见性**: public








- **子元素**: Name BeforeExpression AfterNameExpression AfterExpression Expression Builder 

**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 37-63
- **可见性**: public
- **描述**: Build build clause



- **参数**:
  - `builder`: Builder






**Constant**: `PrimaryKey`

- **完整名称**: PrimaryKey
- **文件位置**: 行 66-66
- **可见性**: public

- **类型**: `string`




- **值**: `"~~~py~~~"`


**Constant**: `CurrentTable`

- **完整名称**: CurrentTable
- **文件位置**: 行 67-67
- **可见性**: public

- **类型**: `string`




- **值**: `"~~~ct~~~"`


**Constant**: `Associations`

- **完整名称**: Associations
- **文件位置**: 行 68-68
- **可见性**: public

- **类型**: `string`




- **值**: `"~~~as~~~"`


**Variable**: `currentTable`

- **完整名称**: currentTable
- **文件位置**: 行 72-72
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...`


**Variable**: `PrimaryColumn`

- **完整名称**: PrimaryColumn
- **文件位置**: 行 73-73
- **可见性**: public

- **类型**: `interface{}`




- **值**: `...`


**Struct**: `Column`

- **完整名称**: Column
- **文件位置**: 行 77-82
- **可见性**: public








- **子元素**: Table Name Alias Raw 

**Struct**: `Table`

- **完整名称**: Table
- **文件位置**: 行 85-89
- **可见性**: public








- **子元素**: Name Alias Raw 




### clause\delete.go

- **语言**: go
- **包**: clause
- **代码行数**: 18
- **元素数量**: 4




#### 代码元素


**Struct**: `Delete`

- **完整名称**: Delete
- **文件位置**: 行 3-5
- **可见性**: public








- **子元素**: Modifier 

**Method**: `Name`

- **完整名称**: Name
- **文件位置**: 行 7-9
- **可见性**: public


- **返回值**: `string`






**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 11-18
- **可见性**: public




- **参数**:
  - `builder`: Builder






**Method**: `MergeClause`

- **完整名称**: MergeClause
- **文件位置**: 行 20-23
- **可见性**: public




- **参数**:
  - `clause`: *Clause









### clause\expression.go

- **语言**: go
- **包**: clause
- **代码行数**: 310
- **元素数量**: 37


#### 导入


- `database/sql`

- `database/sql/driver`

- `go/ast`

- `reflect`




#### 代码元素


**Interface**: `Expression`

- **完整名称**: Expression
- **文件位置**: 行 11-13
- **可见性**: public








- **子元素**: Build 

**Interface**: `NegationExpressionBuilder`

- **完整名称**: NegationExpressionBuilder
- **文件位置**: 行 16-18
- **可见性**: public








- **子元素**: NegationBuild 

**Struct**: `Expr`

- **完整名称**: Expr
- **文件位置**: 行 21-25
- **可见性**: public








- **子元素**: SQL Vars WithoutParentheses 

**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 28-58
- **可见性**: public
- **描述**: Build build raw expression



- **参数**:
  - `builder`: Builder






**Variable**: `afterParenthesis`

- **完整名称**: afterParenthesis
- **文件位置**: 行 30-30
- **可见性**: private

- **类型**: `bool`







**Variable**: `idx`

- **完整名称**: idx
- **文件位置**: 行 31-31
- **可见性**: private

- **类型**: `int`







**Struct**: `NamedExpr`

- **完整名称**: NamedExpr
- **文件位置**: 行 61-64
- **可见性**: public








- **子元素**: SQL Vars 

**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 67-157
- **可见性**: public
- **描述**: Build build raw expression



- **参数**:
  - `builder`: Builder






**Variable**: `idx`

- **完整名称**: idx
- **文件位置**: 行 69-69
- **可见性**: private

- **类型**: `int`







**Variable**: `inName`

- **完整名称**: inName
- **文件位置**: 行 70-70
- **可见性**: private

- **类型**: `bool`







**Variable**: `afterParenthesis`

- **完整名称**: afterParenthesis
- **文件位置**: 行 71-71
- **可见性**: private

- **类型**: `bool`







**Variable**: `namedMap`

- **完整名称**: namedMap
- **文件位置**: 行 72-72
- **可见性**: private

- **类型**: `interface{}`




- **值**: `make(...)`


**Variable**: `appendFieldsToMap`

- **完整名称**: appendFieldsToMap
- **文件位置**: 行 84-84
- **可见性**: private

- **类型**: `func(...)`







**Struct**: `IN`

- **完整名称**: IN
- **文件位置**: 行 185-188
- **可见性**: public








- **子元素**: Column Values 

**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 190-209
- **可见性**: public




- **参数**:
  - `builder`: Builder






**Method**: `NegationBuild`

- **完整名称**: NegationBuild
- **文件位置**: 行 211-229
- **可见性**: public




- **参数**:
  - `builder`: Builder






**Struct**: `Eq`

- **完整名称**: Eq
- **文件位置**: 行 232-235
- **可见性**: public








- **子元素**: Column Value 

**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 237-263
- **可见性**: public




- **参数**:
  - `builder`: Builder






**Method**: `NegationBuild`

- **完整名称**: NegationBuild
- **文件位置**: 行 265-267
- **可见性**: public




- **参数**:
  - `builder`: Builder






**Class**: `Neq`

- **完整名称**: Neq
- **文件位置**: 行 270-270
- **可见性**: public









**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 272-294
- **可见性**: public




- **参数**:
  - `builder`: Builder






**Method**: `NegationBuild`

- **完整名称**: NegationBuild
- **文件位置**: 行 296-298
- **可见性**: public




- **参数**:
  - `builder`: Builder






**Class**: `Gt`

- **完整名称**: Gt
- **文件位置**: 行 301-301
- **可见性**: public









**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 303-307
- **可见性**: public




- **参数**:
  - `builder`: Builder






**Method**: `NegationBuild`

- **完整名称**: NegationBuild
- **文件位置**: 行 309-311
- **可见性**: public




- **参数**:
  - `builder`: Builder






**Class**: `Gte`

- **完整名称**: Gte
- **文件位置**: 行 314-314
- **可见性**: public









**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 316-320
- **可见性**: public




- **参数**:
  - `builder`: Builder






**Method**: `NegationBuild`

- **完整名称**: NegationBuild
- **文件位置**: 行 322-324
- **可见性**: public




- **参数**:
  - `builder`: Builder






**Class**: `Lt`

- **完整名称**: Lt
- **文件位置**: 行 327-327
- **可见性**: public









**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 329-333
- **可见性**: public




- **参数**:
  - `builder`: Builder






**Method**: `NegationBuild`

- **完整名称**: NegationBuild
- **文件位置**: 行 335-337
- **可见性**: public




- **参数**:
  - `builder`: Builder






**Class**: `Lte`

- **完整名称**: Lte
- **文件位置**: 行 340-340
- **可见性**: public









**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 342-346
- **可见性**: public




- **参数**:
  - `builder`: Builder






**Method**: `NegationBuild`

- **完整名称**: NegationBuild
- **文件位置**: 行 348-350
- **可见性**: public




- **参数**:
  - `builder`: Builder






**Class**: `Like`

- **完整名称**: Like
- **文件位置**: 行 353-353
- **可见性**: public









**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 355-359
- **可见性**: public




- **参数**:
  - `builder`: Builder






**Method**: `NegationBuild`

- **完整名称**: NegationBuild
- **文件位置**: 行 361-365
- **可见性**: public




- **参数**:
  - `builder`: Builder









### clause\from.go

- **语言**: go
- **包**: clause
- **代码行数**: 27
- **元素数量**: 4




#### 代码元素


**Struct**: `From`

- **完整名称**: From
- **文件位置**: 行 4-7
- **可见性**: public








- **子元素**: Tables Joins 

**Method**: `Name`

- **完整名称**: Name
- **文件位置**: 行 10-12
- **可见性**: public
- **描述**: Name from clause name

- **返回值**: `string`






**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 15-32
- **可见性**: public
- **描述**: Build build from clause



- **参数**:
  - `builder`: Builder






**Method**: `MergeClause`

- **完整名称**: MergeClause
- **文件位置**: 行 35-37
- **可见性**: public
- **描述**: MergeClause merge from clause



- **参数**:
  - `clause`: *Clause









### clause\group_by.go

- **语言**: go
- **包**: clause
- **代码行数**: 36
- **元素数量**: 4




#### 代码元素


**Struct**: `GroupBy`

- **完整名称**: GroupBy
- **文件位置**: 行 4-7
- **可见性**: public








- **子元素**: Columns Having 

**Method**: `Name`

- **完整名称**: Name
- **文件位置**: 行 10-12
- **可见性**: public
- **描述**: Name from clause name

- **返回值**: `string`






**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 15-28
- **可见性**: public
- **描述**: Build build group by clause



- **参数**:
  - `builder`: Builder






**Method**: `MergeClause`

- **完整名称**: MergeClause
- **文件位置**: 行 31-48
- **可见性**: public
- **描述**: MergeClause merge group by clause



- **参数**:
  - `clause`: *Clause









### clause\insert.go

- **语言**: go
- **包**: clause
- **代码行数**: 31
- **元素数量**: 4




#### 代码元素


**Struct**: `Insert`

- **完整名称**: Insert
- **文件位置**: 行 3-6
- **可见性**: public








- **子元素**: Table Modifier 

**Method**: `Name`

- **完整名称**: Name
- **文件位置**: 行 9-11
- **可见性**: public
- **描述**: Name insert clause name

- **返回值**: `string`






**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 14-26
- **可见性**: public
- **描述**: Build build insert clause



- **参数**:
  - `builder`: Builder






**Method**: `MergeClause`

- **完整名称**: MergeClause
- **文件位置**: 行 29-39
- **可见性**: public
- **描述**: MergeClause merge insert clause



- **参数**:
  - `clause`: *Clause









### clause\joins.go

- **语言**: go
- **包**: clause
- **代码行数**: 65
- **元素数量**: 13


#### 导入


- `gorm.io/gorm/utils`




#### 代码元素


**Class**: `JoinType`

- **完整名称**: JoinType
- **文件位置**: 行 5-5
- **可见性**: public









**Constant**: `CrossJoin`

- **完整名称**: CrossJoin
- **文件位置**: 行 8-8
- **可见性**: public

- **类型**: `JoinType`




- **值**: `"CROSS"`


**Constant**: `InnerJoin`

- **完整名称**: InnerJoin
- **文件位置**: 行 9-9
- **可见性**: public

- **类型**: `JoinType`




- **值**: `"INNER"`


**Constant**: `LeftJoin`

- **完整名称**: LeftJoin
- **文件位置**: 行 10-10
- **可见性**: public

- **类型**: `JoinType`




- **值**: `"LEFT"`


**Constant**: `RightJoin`

- **完整名称**: RightJoin
- **文件位置**: 行 11-11
- **可见性**: public

- **类型**: `JoinType`




- **值**: `"RIGHT"`


**Struct**: `JoinTarget`

- **完整名称**: JoinTarget
- **文件位置**: 行 14-19
- **可见性**: public








- **子元素**: Type Association Subquery Table 

**Function**: `Has`

- **完整名称**: Has
- **文件位置**: 行 21-23
- **可见性**: public


- **返回值**: `JoinTarget`

- **参数**:
  - `name`: string






**Method**: `Association`

- **完整名称**: Association
- **文件位置**: 行 25-27
- **可见性**: public


- **返回值**: `JoinTarget`

- **参数**:
  - `name`: string






**Method**: `AssociationFrom`

- **完整名称**: AssociationFrom
- **文件位置**: 行 29-31
- **可见性**: public


- **返回值**: `JoinTarget`

- **参数**:
  - `name`: string
  - `subquery`: Expression






**Method**: `As`

- **完整名称**: As
- **文件位置**: 行 33-36
- **可见性**: public


- **返回值**: `JoinTarget`

- **参数**:
  - `name`: string






**Struct**: `Join`

- **完整名称**: Join
- **文件位置**: 行 39-45
- **可见性**: public








- **子元素**: Type Table ON Using Expression 

**Function**: `JoinTable`

- **完整名称**: JoinTable
- **文件位置**: 行 47-51
- **可见性**: public


- **返回值**: `Table`

- **参数**:
  - `names`: ...string






**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 53-79
- **可见性**: public




- **参数**:
  - `builder`: Builder









### clause\limit.go

- **语言**: go
- **包**: clause
- **代码行数**: 35
- **元素数量**: 4




#### 代码元素


**Struct**: `Limit`

- **完整名称**: Limit
- **文件位置**: 行 4-7
- **可见性**: public








- **子元素**: Limit Offset 

**Method**: `Name`

- **完整名称**: Name
- **文件位置**: 行 10-12
- **可见性**: public
- **描述**: Name where clause name

- **返回值**: `string`






**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 15-27
- **可见性**: public
- **描述**: Build build where clause



- **参数**:
  - `builder`: Builder






**Method**: `MergeClause`

- **完整名称**: MergeClause
- **文件位置**: 行 30-46
- **可见性**: public
- **描述**: MergeClause merge order by clauses



- **参数**:
  - `clause`: *Clause









### clause\locking.go

- **语言**: go
- **包**: clause
- **代码行数**: 29
- **元素数量**: 8




#### 代码元素


**Constant**: `LockingStrengthUpdate`

- **完整名称**: LockingStrengthUpdate
- **文件位置**: 行 4-4
- **可见性**: public

- **类型**: `interface{}`




- **值**: `"UPDATE"`


**Constant**: `LockingStrengthShare`

- **完整名称**: LockingStrengthShare
- **文件位置**: 行 5-5
- **可见性**: public

- **类型**: `interface{}`




- **值**: `"SHARE"`


**Constant**: `LockingOptionsSkipLocked`

- **完整名称**: LockingOptionsSkipLocked
- **文件位置**: 行 6-6
- **可见性**: public

- **类型**: `interface{}`




- **值**: `"SKIP LOCKED"`


**Constant**: `LockingOptionsNoWait`

- **完整名称**: LockingOptionsNoWait
- **文件位置**: 行 7-7
- **可见性**: public

- **类型**: `interface{}`




- **值**: `"NOWAIT"`


**Struct**: `Locking`

- **完整名称**: Locking
- **文件位置**: 行 10-14
- **可见性**: public








- **子元素**: Strength Table Options 

**Method**: `Name`

- **完整名称**: Name
- **文件位置**: 行 17-19
- **可见性**: public
- **描述**: Name where clause name

- **返回值**: `string`






**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 22-33
- **可见性**: public
- **描述**: Build build where clause



- **参数**:
  - `builder`: Builder






**Method**: `MergeClause`

- **完整名称**: MergeClause
- **文件位置**: 行 36-38
- **可见性**: public
- **描述**: MergeClause merge order by clauses



- **参数**:
  - `clause`: *Clause









### clause\on_conflict.go

- **语言**: go
- **包**: clause
- **代码行数**: 50
- **元素数量**: 4




#### 代码元素


**Struct**: `OnConflict`

- **完整名称**: OnConflict
- **文件位置**: 行 3-11
- **可见性**: public








- **子元素**: Columns Where TargetWhere OnConstraint DoNothing DoUpdates UpdateAll 

**Method**: `Name`

- **完整名称**: Name
- **文件位置**: 行 13-15
- **可见性**: public


- **返回值**: `string`






**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 18-54
- **可见性**: public
- **描述**: Build build onConflict clause



- **参数**:
  - `builder`: Builder






**Method**: `MergeClause`

- **完整名称**: MergeClause
- **文件位置**: 行 57-59
- **可见性**: public
- **描述**: MergeClause merge onConflict clauses



- **参数**:
  - `clause`: *Clause









### clause\order_by.go

- **语言**: go
- **包**: clause
- **代码行数**: 43
- **元素数量**: 5




#### 代码元素


**Struct**: `OrderByColumn`

- **完整名称**: OrderByColumn
- **文件位置**: 行 3-7
- **可见性**: public








- **子元素**: Column Desc Reorder 

**Struct**: `OrderBy`

- **完整名称**: OrderBy
- **文件位置**: 行 9-12
- **可见性**: public








- **子元素**: Columns Expression 

**Method**: `Name`

- **完整名称**: Name
- **文件位置**: 行 15-17
- **可见性**: public
- **描述**: Name where clause name

- **返回值**: `string`






**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 20-35
- **可见性**: public
- **描述**: Build build where clause



- **参数**:
  - `builder`: Builder






**Method**: `MergeClause`

- **完整名称**: MergeClause
- **文件位置**: 行 38-54
- **可见性**: public
- **描述**: MergeClause merge order by clauses



- **参数**:
  - `clause`: *Clause









### clause\returning.go

- **语言**: go
- **包**: clause
- **代码行数**: 29
- **元素数量**: 4




#### 代码元素


**Struct**: `Returning`

- **完整名称**: Returning
- **文件位置**: 行 3-5
- **可见性**: public








- **子元素**: Columns 

**Method**: `Name`

- **完整名称**: Name
- **文件位置**: 行 8-10
- **可见性**: public
- **描述**: Name where clause name

- **返回值**: `string`






**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 13-25
- **可见性**: public
- **描述**: Build build where clause



- **参数**:
  - `builder`: Builder






**Method**: `MergeClause`

- **完整名称**: MergeClause
- **文件位置**: 行 28-37
- **可见性**: public
- **描述**: MergeClause merge order by clauses



- **参数**:
  - `clause`: *Clause









### clause\select.go

- **语言**: go
- **包**: clause
- **代码行数**: 49
- **元素数量**: 6




#### 代码元素


**Struct**: `Select`

- **完整名称**: Select
- **文件位置**: 行 4-8
- **可见性**: public








- **子元素**: Distinct Columns Expression 

**Method**: `Name`

- **完整名称**: Name
- **文件位置**: 行 10-12
- **可见性**: public


- **返回值**: `string`






**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 14-29
- **可见性**: public




- **参数**:
  - `builder`: Builder






**Method**: `MergeClause`

- **完整名称**: MergeClause
- **文件位置**: 行 31-45
- **可见性**: public




- **参数**:
  - `clause`: *Clause






**Struct**: `CommaExpression`

- **完整名称**: CommaExpression
- **文件位置**: 行 48-50
- **可见性**: public








- **子元素**: Exprs 

**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 52-59
- **可见性**: public




- **参数**:
  - `builder`: Builder









### clause\set.go

- **语言**: go
- **包**: clause
- **代码行数**: 55
- **元素数量**: 10


#### 导入


- `sort`




#### 代码元素


**Class**: `Set`

- **完整名称**: Set
- **文件位置**: 行 5-5
- **可见性**: public









**Struct**: `Assignment`

- **完整名称**: Assignment
- **文件位置**: 行 7-10
- **可见性**: public








- **子元素**: Column Value 

**Interface**: `Assigner`

- **完整名称**: Assigner
- **文件位置**: 行 13-15
- **可见性**: public








- **子元素**: Assignments 

**Method**: `Name`

- **完整名称**: Name
- **文件位置**: 行 17-19
- **可见性**: public


- **返回值**: `string`






**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 21-36
- **可见性**: public




- **参数**:
  - `builder`: Builder






**Method**: `MergeClause`

- **完整名称**: MergeClause
- **文件位置**: 行 39-43
- **可见性**: public
- **描述**: MergeClause merge assignments clauses



- **参数**:
  - `clause`: *Clause






**Method**: `Assignments`

- **完整名称**: Assignments
- **文件位置**: 行 46-46
- **可见性**: public
- **描述**: Assignments implements Assigner for Set.

- **返回值**: `[]Assignment`






**Function**: `Assignments`

- **完整名称**: Assignments
- **文件位置**: 行 48-60
- **可见性**: public


- **返回值**: `Set`

- **参数**:
  - `values`: map[string]interface{}






**Function**: `AssignmentColumns`

- **完整名称**: AssignmentColumns
- **文件位置**: 行 62-68
- **可见性**: public


- **返回值**: `Set`

- **参数**:
  - `values`: []string






**Method**: `Assignments`

- **完整名称**: Assignments
- **文件位置**: 行 71-71
- **可见性**: public
- **描述**: Assignments implements Assigner for a single Assignment.

- **返回值**: `[]Assignment`









### clause\update.go

- **语言**: go
- **包**: clause
- **代码行数**: 30
- **元素数量**: 4




#### 代码元素


**Struct**: `Update`

- **完整名称**: Update
- **文件位置**: 行 3-6
- **可见性**: public








- **子元素**: Modifier Table 

**Method**: `Name`

- **完整名称**: Name
- **文件位置**: 行 9-11
- **可见性**: public
- **描述**: Name update clause name

- **返回值**: `string`






**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 14-25
- **可见性**: public
- **描述**: Build build update clause



- **参数**:
  - `builder`: Builder






**Method**: `MergeClause`

- **完整名称**: MergeClause
- **文件位置**: 行 28-38
- **可见性**: public
- **描述**: MergeClause merge update clause



- **参数**:
  - `clause`: *Clause









### clause\values.go

- **语言**: go
- **包**: clause
- **代码行数**: 35
- **元素数量**: 4




#### 代码元素


**Struct**: `Values`

- **完整名称**: Values
- **文件位置**: 行 3-6
- **可见性**: public








- **子元素**: Columns Values 

**Method**: `Name`

- **完整名称**: Name
- **文件位置**: 行 9-11
- **可见性**: public
- **描述**: Name from clause name

- **返回值**: `string`






**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 14-39
- **可见性**: public
- **描述**: Build build from clause



- **参数**:
  - `builder`: Builder






**Method**: `MergeClause`

- **完整名称**: MergeClause
- **文件位置**: 行 42-45
- **可见性**: public
- **描述**: MergeClause merge values clauses



- **参数**:
  - `clause`: *Clause









### clause\where.go

- **语言**: go
- **包**: clause
- **代码行数**: 205
- **元素数量**: 15


#### 导入


- `strings`




#### 代码元素


**Constant**: `AndWithSpace`

- **完整名称**: AndWithSpace
- **文件位置**: 行 8-8
- **可见性**: public

- **类型**: `interface{}`




- **值**: `" AND "`


**Constant**: `OrWithSpace`

- **完整名称**: OrWithSpace
- **文件位置**: 行 9-9
- **可见性**: public

- **类型**: `interface{}`




- **值**: `" OR "`


**Struct**: `Where`

- **完整名称**: Where
- **文件位置**: 行 13-15
- **可见性**: public








- **子元素**: Exprs 

**Method**: `Name`

- **完整名称**: Name
- **文件位置**: 行 18-20
- **可见性**: public
- **描述**: Name where clause name

- **返回值**: `string`






**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 23-41
- **可见性**: public
- **描述**: Build build where clause



- **参数**:
  - `builder`: Builder






**Method**: `MergeClause`

- **完整名称**: MergeClause
- **文件位置**: 行 92-101
- **可见性**: public
- **描述**: MergeClause merge where clauses



- **参数**:
  - `clause`: *Clause






**Function**: `And`

- **完整名称**: And
- **文件位置**: 行 103-115
- **可见性**: public


- **返回值**: `Expression`

- **参数**:
  - `exprs`: ...Expression






**Struct**: `AndConditions`

- **完整名称**: AndConditions
- **文件位置**: 行 117-119
- **可见性**: public








- **子元素**: Exprs 

**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 121-129
- **可见性**: public




- **参数**:
  - `builder`: Builder






**Function**: `Or`

- **完整名称**: Or
- **文件位置**: 行 131-136
- **可见性**: public


- **返回值**: `Expression`

- **参数**:
  - `exprs`: ...Expression






**Struct**: `OrConditions`

- **完整名称**: OrConditions
- **文件位置**: 行 138-140
- **可见性**: public








- **子元素**: Exprs 

**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 142-150
- **可见性**: public




- **参数**:
  - `builder`: Builder






**Function**: `Not`

- **完整名称**: Not
- **文件位置**: 行 152-162
- **可见性**: public


- **返回值**: `Expression`

- **参数**:
  - `exprs`: ...Expression






**Struct**: `NotConditions`

- **完整名称**: NotConditions
- **文件位置**: 行 164-166
- **可见性**: public








- **子元素**: Exprs 

**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 168-245
- **可见性**: public




- **参数**:
  - `builder`: Builder









### clause\with.go

- **语言**: go
- **包**: clause
- **代码行数**: 2
- **元素数量**: 1




#### 代码元素


**Struct**: `With`

- **完整名称**: With
- **文件位置**: 行 3-3
- **可见性**: public












### errors.go

- **语言**: go
- **包**: gorm
- **代码行数**: 29
- **元素数量**: 22


#### 导入


- `errors`

- `gorm.io/gorm/logger`




#### 代码元素


**Variable**: `ErrRecordNotFound`

- **完整名称**: ErrRecordNotFound
- **文件位置**: 行 11-11
- **可见性**: public
- **描述**: ErrRecordNotFound record not found error
- **类型**: `logger.ErrRecordNotFound`




- **值**: `logger.ErrRecordNotFound`


**Variable**: `ErrInvalidTransaction`

- **完整名称**: ErrInvalidTransaction
- **文件位置**: 行 13-13
- **可见性**: public
- **描述**: ErrInvalidTransaction invalid transaction when you are trying to `Commit` or `Rollback`
- **类型**: `interface{}`




- **值**: `errors.New(...)`


**Variable**: `ErrNotImplemented`

- **完整名称**: ErrNotImplemented
- **文件位置**: 行 15-15
- **可见性**: public
- **描述**: ErrNotImplemented not implemented
- **类型**: `interface{}`




- **值**: `errors.New(...)`


**Variable**: `ErrMissingWhereClause`

- **完整名称**: ErrMissingWhereClause
- **文件位置**: 行 17-17
- **可见性**: public
- **描述**: ErrMissingWhereClause missing where clause
- **类型**: `interface{}`




- **值**: `errors.New(...)`


**Variable**: `ErrUnsupportedRelation`

- **完整名称**: ErrUnsupportedRelation
- **文件位置**: 行 19-19
- **可见性**: public
- **描述**: ErrUnsupportedRelation unsupported relations
- **类型**: `interface{}`




- **值**: `errors.New(...)`


**Variable**: `ErrPrimaryKeyRequired`

- **完整名称**: ErrPrimaryKeyRequired
- **文件位置**: 行 21-21
- **可见性**: public
- **描述**: ErrPrimaryKeyRequired primary keys required
- **类型**: `interface{}`




- **值**: `errors.New(...)`


**Variable**: `ErrModelValueRequired`

- **完整名称**: ErrModelValueRequired
- **文件位置**: 行 23-23
- **可见性**: public
- **描述**: ErrModelValueRequired model value required
- **类型**: `interface{}`




- **值**: `errors.New(...)`


**Variable**: `ErrModelAccessibleFieldsRequired`

- **完整名称**: ErrModelAccessibleFieldsRequired
- **文件位置**: 行 25-25
- **可见性**: public
- **描述**: ErrModelAccessibleFieldsRequired model accessible fields required
- **类型**: `interface{}`




- **值**: `errors.New(...)`


**Variable**: `ErrSubQueryRequired`

- **完整名称**: ErrSubQueryRequired
- **文件位置**: 行 27-27
- **可见性**: public
- **描述**: ErrSubQueryRequired sub query required
- **类型**: `interface{}`




- **值**: `errors.New(...)`


**Variable**: `ErrInvalidData`

- **完整名称**: ErrInvalidData
- **文件位置**: 行 29-29
- **可见性**: public
- **描述**: ErrInvalidData unsupported data
- **类型**: `interface{}`




- **值**: `errors.New(...)`


**Variable**: `ErrUnsupportedDriver`

- **完整名称**: ErrUnsupportedDriver
- **文件位置**: 行 31-31
- **可见性**: public
- **描述**: ErrUnsupportedDriver unsupported driver
- **类型**: `interface{}`




- **值**: `errors.New(...)`


**Variable**: `ErrRegistered`

- **完整名称**: ErrRegistered
- **文件位置**: 行 33-33
- **可见性**: public
- **描述**: ErrRegistered registered
- **类型**: `interface{}`




- **值**: `errors.New(...)`


**Variable**: `ErrInvalidField`

- **完整名称**: ErrInvalidField
- **文件位置**: 行 35-35
- **可见性**: public
- **描述**: ErrInvalidField invalid field
- **类型**: `interface{}`




- **值**: `errors.New(...)`


**Variable**: `ErrEmptySlice`

- **完整名称**: ErrEmptySlice
- **文件位置**: 行 37-37
- **可见性**: public
- **描述**: ErrEmptySlice empty slice found
- **类型**: `interface{}`




- **值**: `errors.New(...)`


**Variable**: `ErrDryRunModeUnsupported`

- **完整名称**: ErrDryRunModeUnsupported
- **文件位置**: 行 39-39
- **可见性**: public
- **描述**: ErrDryRunModeUnsupported dry run mode unsupported
- **类型**: `interface{}`




- **值**: `errors.New(...)`


**Variable**: `ErrInvalidDB`

- **完整名称**: ErrInvalidDB
- **文件位置**: 行 41-41
- **可见性**: public
- **描述**: ErrInvalidDB invalid db
- **类型**: `interface{}`




- **值**: `errors.New(...)`


**Variable**: `ErrInvalidValue`

- **完整名称**: ErrInvalidValue
- **文件位置**: 行 43-43
- **可见性**: public
- **描述**: ErrInvalidValue invalid value
- **类型**: `interface{}`




- **值**: `errors.New(...)`


**Variable**: `ErrInvalidValueOfLength`

- **完整名称**: ErrInvalidValueOfLength
- **文件位置**: 行 45-45
- **可见性**: public
- **描述**: ErrInvalidValueOfLength invalid values do not match length
- **类型**: `interface{}`




- **值**: `errors.New(...)`


**Variable**: `ErrPreloadNotAllowed`

- **完整名称**: ErrPreloadNotAllowed
- **文件位置**: 行 47-47
- **可见性**: public
- **描述**: ErrPreloadNotAllowed preload is not allowed when count is used
- **类型**: `interface{}`




- **值**: `errors.New(...)`


**Variable**: `ErrDuplicatedKey`

- **完整名称**: ErrDuplicatedKey
- **文件位置**: 行 49-49
- **可见性**: public
- **描述**: ErrDuplicatedKey occurs when there is a unique key constraint violation
- **类型**: `interface{}`




- **值**: `errors.New(...)`


**Variable**: `ErrForeignKeyViolated`

- **完整名称**: ErrForeignKeyViolated
- **文件位置**: 行 51-51
- **可见性**: public
- **描述**: ErrForeignKeyViolated occurs when there is a foreign key constraint violation
- **类型**: `interface{}`




- **值**: `errors.New(...)`


**Variable**: `ErrCheckConstraintViolated`

- **完整名称**: ErrCheckConstraintViolated
- **文件位置**: 行 53-53
- **可见性**: public
- **描述**: ErrCheckConstraintViolated occurs when there is a check constraint violation
- **类型**: `interface{}`




- **值**: `errors.New(...)`





### finisher_api.go

- **语言**: go
- **包**: gorm
- **代码行数**: 611
- **元素数量**: 43


#### 导入


- `context`

- `database/sql`

- `errors`

- `fmt`

- `hash/maphash`

- `reflect`

- `strings`

- `gorm.io/gorm/clause`

- `gorm.io/gorm/logger`

- `gorm.io/gorm/schema`

- `gorm.io/gorm/utils`




#### 代码元素


**Method**: `Create`

- **完整名称**: Create
- **文件位置**: 行 19-27
- **可见性**: public
- **描述**: Create inserts value, returning the inserted data's primary key in value's id

- **返回值**: `*DB`

- **参数**:
  - `value`: interface{}






**Method**: `CreateInBatches`

- **完整名称**: CreateInBatches
- **文件位置**: 行 30-72
- **可见性**: public
- **描述**: CreateInBatches inserts value in batches of batchSize

- **返回值**: `*DB`

- **参数**:
  - `value`: interface{}
  - `batchSize`: int






**Variable**: `rowsAffected`

- **完整名称**: rowsAffected
- **文件位置**: 行 35-35
- **可见性**: private

- **类型**: `int64`







**Method**: `Save`

- **完整名称**: Save
- **文件位置**: 行 75-117
- **可见性**: public
- **描述**: Save updates value in database. If value doesn't contain a matching primary key, value is inserted.

- **返回值**: `*DB`

- **参数**:
  - `value`: interface{}






**Method**: `First`

- **完整名称**: First
- **文件位置**: 行 120-132
- **可见性**: public
- **描述**: First finds the first record ordered by primary key, matching given conditions conds

- **返回值**: `*DB`

- **参数**:
  - `dest`: interface{}
  - `conds`: ...interface{}






**Method**: `Take`

- **完整名称**: Take
- **文件位置**: 行 135-145
- **可见性**: public
- **描述**: Take finds the first record returned by the database in no specified order, matching given conditions conds

- **返回值**: `*DB`

- **参数**:
  - `dest`: interface{}
  - `conds`: ...interface{}






**Method**: `Last`

- **完整名称**: Last
- **文件位置**: 行 148-161
- **可见性**: public
- **描述**: Last finds the last record ordered by primary key, matching given conditions conds

- **返回值**: `*DB`

- **参数**:
  - `dest`: interface{}
  - `conds`: ...interface{}






**Method**: `Find`

- **完整名称**: Find
- **文件位置**: 行 164-173
- **可见性**: public
- **描述**: Find finds all records matching given conditions conds

- **返回值**: `*DB`

- **参数**:
  - `dest`: interface{}
  - `conds`: ...interface{}






**Method**: `FindInBatches`

- **完整名称**: FindInBatches
- **文件位置**: 行 176-246
- **可见性**: public
- **描述**: FindInBatches finds all records in batches of batchSize

- **返回值**: `*DB`

- **参数**:
  - `dest`: interface{}
  - `batchSize`: int
  - `fc`: func(...)






**Variable**: `tx`

- **完整名称**: tx
- **文件位置**: 行 178-178
- **可见性**: private

- **类型**: `interface{}`




- **值**: `db.Order(...).Session(...)`


**Variable**: `queryDB`

- **完整名称**: queryDB
- **文件位置**: 行 181-181
- **可见性**: private

- **类型**: `tx`




- **值**: `tx`


**Variable**: `rowsAffected`

- **完整名称**: rowsAffected
- **文件位置**: 行 182-182
- **可见性**: private

- **类型**: `int64`







**Variable**: `batch`

- **完整名称**: batch
- **文件位置**: 行 183-183
- **可见性**: private

- **类型**: `int`







**Variable**: `totalSize`

- **完整名称**: totalSize
- **文件位置**: 行 187-187
- **可见性**: private

- **类型**: `int`







**Method**: `assignInterfacesToValue`

- **完整名称**: assignInterfacesToValue
- **文件位置**: 行 248-295
- **可见性**: private




- **参数**:
  - `values`: ...interface{}






**Method**: `FirstOrInit`

- **完整名称**: FirstOrInit
- **文件位置**: 行 309-332
- **可见性**: public
- **描述**: FirstOrInit finds the first matching record, otherwise if not found initializes a new instance with given conds.

- **返回值**: `*DB`

- **参数**:
  - `dest`: interface{}
  - `conds`: ...interface{}






**Method**: `FirstOrCreate`

- **完整名称**: FirstOrCreate
- **文件位置**: 行 348-400
- **可见性**: public
- **描述**: FirstOrCreate finds the first matching record, otherwise if not found creates a new instance with given conds.

- **返回值**: `*DB`

- **参数**:
  - `dest`: interface{}
  - `conds`: ...interface{}






**Method**: `Update`

- **完整名称**: Update
- **文件位置**: 行 403-407
- **可见性**: public
- **描述**: Update updates column with value using callbacks. Reference: https://gorm.io/docs/update.html#Update-Changed-Fields

- **返回值**: `*DB`

- **参数**:
  - `column`: string
  - `value`: interface{}






**Method**: `Updates`

- **完整名称**: Updates
- **文件位置**: 行 410-414
- **可见性**: public
- **描述**: Updates updates attributes using callbacks. values must be a struct or map. Reference: https://gorm.io/docs/update.html#Update-Changed-Fields

- **返回值**: `*DB`

- **参数**:
  - `values`: interface{}






**Method**: `UpdateColumn`

- **完整名称**: UpdateColumn
- **文件位置**: 行 416-421
- **可见性**: public


- **返回值**: `*DB`

- **参数**:
  - `column`: string
  - `value`: interface{}






**Method**: `UpdateColumns`

- **完整名称**: UpdateColumns
- **文件位置**: 行 423-428
- **可见性**: public


- **返回值**: `*DB`

- **参数**:
  - `values`: interface{}






**Method**: `Delete`

- **完整名称**: Delete
- **文件位置**: 行 433-442
- **可见性**: public
- **描述**: Delete deletes value matching given conditions. If value contains primary key it is included in the conditions. If

- **返回值**: `*DB`

- **参数**:
  - `value`: interface{}
  - `conds`: ...interface{}






**Method**: `Count`

- **完整名称**: Count
- **文件位置**: 行 444-504
- **可见性**: public


- **返回值**: `*DB`

- **参数**:
  - `count`: *int64






**Method**: `Row`

- **完整名称**: Row
- **文件位置**: 行 506-514
- **可见性**: public


- **返回值**: `*sql.Row`






**Method**: `Rows`

- **完整名称**: Rows
- **文件位置**: 行 516-524
- **可见性**: public


- **返回值**: `*sql.Rows`






**Method**: `Scan`

- **完整名称**: Scan
- **文件位置**: 行 527-550
- **可见性**: public
- **描述**: Scan scans selected value to the struct dest

- **返回值**: `*DB`

- **参数**:
  - `dest`: interface{}






**Method**: `Pluck`

- **完整名称**: Pluck
- **文件位置**: 行 556-575
- **可见性**: public
- **描述**: Pluck queries a single column from a model, returning in the slice dest. E.g.:

- **返回值**: `*DB`

- **参数**:
  - `column`: string
  - `dest`: interface{}






**Method**: `ScanRows`

- **完整名称**: ScanRows
- **文件位置**: 行 577-594
- **可见性**: public


- **返回值**: `error`

- **参数**:
  - `rows`: *sql.Rows
  - `dest`: interface{}






**Method**: `Connection`

- **完整名称**: Connection
- **文件位置**: 行 598-617
- **可见性**: public
- **描述**: Connection uses a db connection to execute an arbitrary number of commands in fc. When finished, the connection is

- **返回值**: `error`

- **参数**:
  - `fc`: func(...)






**Method**: `Transaction`

- **完整名称**: Transaction
- **文件位置**: 行 622-662
- **可见性**: public
- **描述**: Transaction start a transaction as a block, return error will rollback, otherwise to commit. Transaction executes an

- **返回值**: `error`

- **参数**:
  - `fc`: func(...)
  - `opts`: ...*sql.TxOptions






**Method**: `Begin`

- **完整名称**: Begin
- **文件位置**: 行 665-698
- **可见性**: public
- **描述**: Begin begins a transaction with any transaction options opts

- **返回值**: `*DB`

- **参数**:
  - `opts`: ...*sql.TxOptions






**Variable**: `tx`

- **完整名称**: tx
- **文件位置**: 行 668-668
- **可见性**: private
- **描述**: clone statement
- **类型**: `interface{}`




- **值**: `db.getInstance(...).Session(...)`


**Variable**: `opt`

- **完整名称**: opt
- **文件位置**: 行 669-669
- **可见性**: private

- **类型**: `*sql.TxOptions`







**Variable**: `err`

- **完整名称**: err
- **文件位置**: 行 670-670
- **可见性**: private

- **类型**: `error`







**Method**: `Commit`

- **完整名称**: Commit
- **文件位置**: 行 701-708
- **可见性**: public
- **描述**: Commit commits the changes in a transaction

- **返回值**: `*DB`






**Method**: `Rollback`

- **完整名称**: Rollback
- **文件位置**: 行 711-720
- **可见性**: public
- **描述**: Rollback rollbacks the changes in a transaction

- **返回值**: `*DB`






**Method**: `SavePoint`

- **完整名称**: SavePoint
- **文件位置**: 行 722-743
- **可见性**: public


- **返回值**: `*DB`

- **参数**:
  - `name`: string






**Variable**: `preparedStmtTx`

- **完整名称**: preparedStmtTx
- **文件位置**: 行 727-727
- **可见性**: private

- **类型**: `*PreparedStmtTX`







**Variable**: `isPreparedStmtTx`

- **完整名称**: isPreparedStmtTx
- **文件位置**: 行 728-728
- **可见性**: private

- **类型**: `bool`







**Method**: `RollbackTo`

- **完整名称**: RollbackTo
- **文件位置**: 行 745-766
- **可见性**: public


- **返回值**: `*DB`

- **参数**:
  - `name`: string






**Variable**: `preparedStmtTx`

- **完整名称**: preparedStmtTx
- **文件位置**: 行 750-750
- **可见性**: private

- **类型**: `*PreparedStmtTX`







**Variable**: `isPreparedStmtTx`

- **完整名称**: isPreparedStmtTx
- **文件位置**: 行 751-751
- **可见性**: private

- **类型**: `bool`







**Method**: `Exec`

- **完整名称**: Exec
- **文件位置**: 行 769-780
- **可见性**: public
- **描述**: Exec executes raw sql

- **返回值**: `*DB`

- **参数**:
  - `sql`: string
  - `values`: ...interface{}









### generics.go

- **语言**: go
- **包**: gorm
- **代码行数**: 753
- **元素数量**: 113


#### 导入


- `context`

- `database/sql`

- `errors`

- `fmt`

- `reflect`

- `sort`

- `strings`

- `gorm.io/gorm/clause`

- `gorm.io/gorm/logger`

- `gorm.io/gorm/schema`




#### 代码元素


**Struct**: `result`

- **完整名称**: result
- **文件位置**: 行 17-20
- **可见性**: private








- **子元素**: Result RowsAffected 

**Method**: `ModifyStatement`

- **完整名称**: ModifyStatement
- **文件位置**: 行 22-24
- **可见性**: public




- **参数**:
  - `stmt`: *Statement






**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 27-28
- **可见性**: public
- **描述**: Build implements clause.Expression interface








**Function**: `WithResult`

- **完整名称**: WithResult
- **文件位置**: 行 30-32
- **可见性**: public


- **返回值**: `*result`






**Interface**: `Interface`

- **完整名称**: Interface
- **文件位置**: 行 34-38
- **可见性**: public








- **子元素**: Raw Exec 

**Interface**: `CreateInterface`

- **完整名称**: CreateInterface
- **文件位置**: 行 40-69
- **可见性**: public








- **子元素**: Scopes Where Not Or Limit Offset Joins Preload Select Omit MapColumns Distinct Group Having Order Build Delete Update Updates Count Table Create CreateInBatches Set 

**Interface**: `ChainInterface`

- **完整名称**: ChainInterface
- **文件位置**: 行 71-97
- **可见性**: public








- **子元素**: Scopes Where Not Or Limit Offset Joins Preload Select Omit MapColumns Distinct Group Having Order Set Build Table Delete Update Updates Count 

**Interface**: `SetUpdateOnlyInterface`

- **完整名称**: SetUpdateOnlyInterface
- **文件位置**: 行 100-102
- **可见性**: public








- **子元素**: Update 

**Interface**: `SetCreateOrUpdateInterface`

- **完整名称**: SetCreateOrUpdateInterface
- **文件位置**: 行 105-108
- **可见性**: public








- **子元素**: Create Update 

**Interface**: `ExecInterface`

- **完整名称**: ExecInterface
- **文件位置**: 行 110-119
- **可见性**: public








- **子元素**: Scan First Last Take Find FindInBatches Row Rows 

**Interface**: `JoinBuilder`

- **完整名称**: JoinBuilder
- **文件位置**: 行 121-127
- **可见性**: public








- **子元素**: Select Omit Where Not Or 

**Interface**: `PreloadBuilder`

- **完整名称**: PreloadBuilder
- **文件位置**: 行 129-139
- **可见性**: public








- **子元素**: Select Omit Where Not Or Limit Offset Order LimitPerRecord 

**Class**: `op`

- **完整名称**: op
- **文件位置**: 行 141-141
- **可见性**: private









**Function**: `G`

- **完整名称**: G
- **文件位置**: 行 143-161
- **可见性**: public


- **返回值**: `interface{}`

- **参数**:
  - `db`: *DB
  - `opts`: ...clause.Expression






**Struct**: `g`

- **完整名称**: g
- **文件位置**: 行 163-167
- **可见性**: private








- **子元素**: db ops 

**Method**: `apply`

- **完整名称**: apply
- **文件位置**: 行 169-179
- **可见性**: private


- **返回值**: `*DB`

- **参数**:
  - `ctx`: context.Context






**Method**: `Raw`

- **完整名称**: Raw
- **文件位置**: 行 181-189
- **可见性**: public


- **返回值**: `interface{}`

- **参数**:
  - `sql`: string
  - `values`: ...interface{}






**Variable**: `r`

- **完整名称**: r
- **文件位置**: 行 185-185
- **可见性**: private

- **类型**: `T`







**Method**: `Exec`

- **完整名称**: Exec
- **文件位置**: 行 191-194
- **可见性**: public


- **返回值**: `error`

- **参数**:
  - `ctx`: context.Context
  - `sql`: string
  - `values`: ...interface{}






**Variable**: `r`

- **完整名称**: r
- **文件位置**: 行 192-192
- **可见性**: private

- **类型**: `T`







**Struct**: `createG`

- **完整名称**: createG
- **文件位置**: 行 196-198
- **可见性**: private









**Method**: `Table`

- **完整名称**: Table
- **文件位置**: 行 200-204
- **可见性**: public


- **返回值**: `interface{}`

- **参数**:
  - `name`: string
  - `args`: ...interface{}






**Method**: `Select`

- **完整名称**: Select
- **文件位置**: 行 206-210
- **可见性**: public


- **返回值**: `interface{}`

- **参数**:
  - `query`: string
  - `args`: ...interface{}






**Method**: `Omit`

- **完整名称**: Omit
- **文件位置**: 行 212-216
- **可见性**: public


- **返回值**: `interface{}`

- **参数**:
  - `columns`: ...string






**Method**: `Set`

- **完整名称**: Set
- **文件位置**: 行 218-220
- **可见性**: public


- **返回值**: `interface{}`

- **参数**:
  - `assignments`: ...clause.Assigner






**Method**: `Create`

- **完整名称**: Create
- **文件位置**: 行 222-224
- **可见性**: public


- **返回值**: `error`

- **参数**:
  - `ctx`: context.Context
  - `r`: *T






**Method**: `CreateInBatches`

- **完整名称**: CreateInBatches
- **文件位置**: 行 226-228
- **可见性**: public


- **返回值**: `error`

- **参数**:
  - `ctx`: context.Context
  - `r`: *[]T
  - `batchSize`: int






**Struct**: `chainG`

- **完整名称**: chainG
- **文件位置**: 行 230-232
- **可见性**: private









**Method**: `getInstance`

- **完整名称**: getInstance
- **文件位置**: 行 234-237
- **可见性**: private


- **返回值**: `*DB`






**Variable**: `r`

- **完整名称**: r
- **文件位置**: 行 235-235
- **可见性**: private

- **类型**: `T`







**Method**: `with`

- **完整名称**: with
- **文件位置**: 行 239-246
- **可见性**: private


- **返回值**: `interface{}`

- **参数**:
  - `v`: op






**Method**: `Table`

- **完整名称**: Table
- **文件位置**: 行 248-252
- **可见性**: public


- **返回值**: `interface{}`

- **参数**:
  - `name`: string
  - `args`: ...interface{}






**Method**: `Scopes`

- **完整名称**: Scopes
- **文件位置**: 行 254-261
- **可见性**: public


- **返回值**: `interface{}`

- **参数**:
  - `scopes`: ...func(...)






**Method**: `Where`

- **完整名称**: Where
- **文件位置**: 行 263-267
- **可见性**: public


- **返回值**: `interface{}`

- **参数**:
  - `query`: interface{}
  - `args`: ...interface{}






**Method**: `Not`

- **完整名称**: Not
- **文件位置**: 行 269-273
- **可见性**: public


- **返回值**: `interface{}`

- **参数**:
  - `query`: interface{}
  - `args`: ...interface{}






**Method**: `Or`

- **完整名称**: Or
- **文件位置**: 行 275-279
- **可见性**: public


- **返回值**: `interface{}`

- **参数**:
  - `query`: interface{}
  - `args`: ...interface{}






**Method**: `Limit`

- **完整名称**: Limit
- **文件位置**: 行 281-285
- **可见性**: public


- **返回值**: `interface{}`

- **参数**:
  - `offset`: int






**Method**: `Offset`

- **完整名称**: Offset
- **文件位置**: 行 287-291
- **可见性**: public


- **返回值**: `interface{}`

- **参数**:
  - `offset`: int






**Struct**: `joinBuilder`

- **完整名称**: joinBuilder
- **文件位置**: 行 293-295
- **可见性**: private








- **子元素**: db 

**Method**: `Where`

- **完整名称**: Where
- **文件位置**: 行 297-300
- **可见性**: public


- **返回值**: `JoinBuilder`

- **参数**:
  - `query`: interface{}
  - `args`: ...interface{}






**Method**: `Or`

- **完整名称**: Or
- **文件位置**: 行 302-305
- **可见性**: public


- **返回值**: `JoinBuilder`

- **参数**:
  - `query`: interface{}
  - `args`: ...interface{}






**Method**: `Not`

- **完整名称**: Not
- **文件位置**: 行 307-310
- **可见性**: public


- **返回值**: `JoinBuilder`

- **参数**:
  - `query`: interface{}
  - `args`: ...interface{}






**Method**: `Select`

- **完整名称**: Select
- **文件位置**: 行 312-315
- **可见性**: public


- **返回值**: `JoinBuilder`

- **参数**:
  - `columns`: ...string






**Method**: `Omit`

- **完整名称**: Omit
- **文件位置**: 行 317-320
- **可见性**: public


- **返回值**: `JoinBuilder`

- **参数**:
  - `columns`: ...string






**Struct**: `preloadBuilder`

- **完整名称**: preloadBuilder
- **文件位置**: 行 322-325
- **可见性**: private








- **子元素**: limitPerRecord db 

**Method**: `Where`

- **完整名称**: Where
- **文件位置**: 行 327-330
- **可见性**: public


- **返回值**: `PreloadBuilder`

- **参数**:
  - `query`: interface{}
  - `args`: ...interface{}






**Method**: `Or`

- **完整名称**: Or
- **文件位置**: 行 332-335
- **可见性**: public


- **返回值**: `PreloadBuilder`

- **参数**:
  - `query`: interface{}
  - `args`: ...interface{}






**Method**: `Not`

- **完整名称**: Not
- **文件位置**: 行 337-340
- **可见性**: public


- **返回值**: `PreloadBuilder`

- **参数**:
  - `query`: interface{}
  - `args`: ...interface{}






**Method**: `Select`

- **完整名称**: Select
- **文件位置**: 行 342-345
- **可见性**: public


- **返回值**: `PreloadBuilder`

- **参数**:
  - `columns`: ...string






**Method**: `Omit`

- **完整名称**: Omit
- **文件位置**: 行 347-350
- **可见性**: public


- **返回值**: `PreloadBuilder`

- **参数**:
  - `columns`: ...string






**Method**: `Limit`

- **完整名称**: Limit
- **文件位置**: 行 352-355
- **可见性**: public


- **返回值**: `PreloadBuilder`

- **参数**:
  - `limit`: int






**Method**: `Offset`

- **完整名称**: Offset
- **文件位置**: 行 357-360
- **可见性**: public


- **返回值**: `PreloadBuilder`

- **参数**:
  - `offset`: int






**Method**: `Order`

- **完整名称**: Order
- **文件位置**: 行 362-365
- **可见性**: public


- **返回值**: `PreloadBuilder`

- **参数**:
  - `value`: interface{}






**Method**: `LimitPerRecord`

- **完整名称**: LimitPerRecord
- **文件位置**: 行 367-370
- **可见性**: public


- **返回值**: `PreloadBuilder`

- **参数**:
  - `num`: int






**Method**: `Joins`

- **完整名称**: Joins
- **文件位置**: 行 372-429
- **可见性**: public


- **返回值**: `interface{}`

- **参数**:
  - `jt`: clause.JoinTarget
  - `on`: func(...)






**Method**: `Select`

- **完整名称**: Select
- **文件位置**: 行 431-435
- **可见性**: public


- **返回值**: `interface{}`

- **参数**:
  - `query`: string
  - `args`: ...interface{}






**Method**: `Omit`

- **完整名称**: Omit
- **文件位置**: 行 437-441
- **可见性**: public


- **返回值**: `interface{}`

- **参数**:
  - `columns`: ...string






**Method**: `MapColumns`

- **完整名称**: MapColumns
- **文件位置**: 行 443-447
- **可见性**: public


- **返回值**: `interface{}`

- **参数**:
  - `m`: map[string]string






**Method**: `Set`

- **完整名称**: Set
- **文件位置**: 行 449-451
- **可见性**: public


- **返回值**: `interface{}`

- **参数**:
  - `assignments`: ...clause.Assigner






**Method**: `Distinct`

- **完整名称**: Distinct
- **文件位置**: 行 453-457
- **可见性**: public


- **返回值**: `interface{}`

- **参数**:
  - `args`: ...interface{}






**Method**: `Group`

- **完整名称**: Group
- **文件位置**: 行 459-463
- **可见性**: public


- **返回值**: `interface{}`

- **参数**:
  - `name`: string






**Method**: `Having`

- **完整名称**: Having
- **文件位置**: 行 465-469
- **可见性**: public


- **返回值**: `interface{}`

- **参数**:
  - `query`: interface{}
  - `args`: ...interface{}






**Method**: `Order`

- **完整名称**: Order
- **文件位置**: 行 471-475
- **可见性**: public


- **返回值**: `interface{}`

- **参数**:
  - `value`: interface{}






**Method**: `Preload`

- **完整名称**: Preload
- **文件位置**: 行 477-558
- **可见性**: public


- **返回值**: `interface{}`

- **参数**:
  - `association`: string
  - `query`: func(...)






**Variable**: `ok`

- **完整名称**: ok
- **文件位置**: 行 492-492
- **可见性**: private

- **类型**: `bool`







**Method**: `Delete`

- **完整名称**: Delete
- **文件位置**: 行 560-564
- **可见性**: public


- **返回值**: `int`

- **参数**:
  - `ctx`: context.Context






**Method**: `Update`

- **完整名称**: Update
- **文件位置**: 行 566-570
- **可见性**: public


- **返回值**: `int`

- **参数**:
  - `ctx`: context.Context
  - `name`: string
  - `value`: any






**Variable**: `r`

- **完整名称**: r
- **文件位置**: 行 567-567
- **可见性**: private

- **类型**: `T`







**Method**: `Updates`

- **完整名称**: Updates
- **文件位置**: 行 572-575
- **可见性**: public


- **返回值**: `int`

- **参数**:
  - `ctx`: context.Context
  - `t`: T






**Method**: `Count`

- **完整名称**: Count
- **文件位置**: 行 577-581
- **可见性**: public


- **返回值**: `int64`

- **参数**:
  - `ctx`: context.Context
  - `column`: string






**Variable**: `r`

- **完整名称**: r
- **文件位置**: 行 578-578
- **可见性**: private

- **类型**: `T`







**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 583-618
- **可见性**: public




- **参数**:
  - `builder`: clause.Builder






**Variable**: `vars`

- **完整名称**: vars
- **文件位置**: 行 591-591
- **可见性**: private

- **类型**: `subdb.Statement.Vars`




- **值**: `subdb.Statement.Vars`


**Variable**: `sql`

- **完整名称**: sql
- **文件位置**: 行 592-592
- **可见性**: private

- **类型**: `interface{}`




- **值**: `subdb.Statement.SQL.String(...)`


**Struct**: `execG`

- **完整名称**: execG
- **文件位置**: 行 620-622
- **可见性**: private








- **子元素**: g 

**Method**: `First`

- **完整名称**: First
- **文件位置**: 行 624-628
- **可见性**: public


- **返回值**: `T`

- **参数**:
  - `ctx`: context.Context






**Variable**: `r`

- **完整名称**: r
- **文件位置**: 行 625-625
- **可见性**: private

- **类型**: `T`







**Method**: `Scan`

- **完整名称**: Scan
- **文件位置**: 行 630-634
- **可见性**: public


- **返回值**: `error`

- **参数**:
  - `ctx`: context.Context
  - `result`: interface{}






**Variable**: `r`

- **完整名称**: r
- **文件位置**: 行 631-631
- **可见性**: private

- **类型**: `T`







**Method**: `Last`

- **完整名称**: Last
- **文件位置**: 行 636-640
- **可见性**: public


- **返回值**: `T`

- **参数**:
  - `ctx`: context.Context






**Variable**: `r`

- **完整名称**: r
- **文件位置**: 行 637-637
- **可见性**: private

- **类型**: `T`







**Method**: `Take`

- **完整名称**: Take
- **文件位置**: 行 642-646
- **可见性**: public


- **返回值**: `T`

- **参数**:
  - `ctx`: context.Context






**Variable**: `r`

- **完整名称**: r
- **文件位置**: 行 643-643
- **可见性**: private

- **类型**: `T`







**Method**: `Find`

- **完整名称**: Find
- **文件位置**: 行 648-652
- **可见性**: public


- **返回值**: `[]T`

- **参数**:
  - `ctx`: context.Context






**Variable**: `r`

- **完整名称**: r
- **文件位置**: 行 649-649
- **可见性**: private

- **类型**: `[]T`







**Method**: `FindInBatches`

- **完整名称**: FindInBatches
- **文件位置**: 行 654-659
- **可见性**: public


- **返回值**: `error`

- **参数**:
  - `ctx`: context.Context
  - `batchSize`: int
  - `fc`: func(...)






**Variable**: `data`

- **完整名称**: data
- **文件位置**: 行 655-655
- **可见性**: private

- **类型**: `[]T`







**Method**: `Row`

- **完整名称**: Row
- **文件位置**: 行 661-664
- **可见性**: public


- **返回值**: `*sql.Row`

- **参数**:
  - `ctx`: context.Context






**Variable**: `r`

- **完整名称**: r
- **文件位置**: 行 662-662
- **可见性**: private

- **类型**: `T`







**Method**: `Rows`

- **完整名称**: Rows
- **文件位置**: 行 666-669
- **可见性**: public


- **返回值**: `*sql.Rows`

- **参数**:
  - `ctx`: context.Context






**Variable**: `r`

- **完整名称**: r
- **文件位置**: 行 667-667
- **可见性**: private

- **类型**: `T`







**Method**: `processSet`

- **完整名称**: processSet
- **文件位置**: 行 671-691
- **可见性**: private


- **返回值**: `interface{}`

- **参数**:
  - `items`: ...clause.Assigner






**Variable**: `assigns`

- **完整名称**: assigns
- **文件位置**: 行 673-673
- **可见性**: private

- **类型**: `[]clause.Assignment`







**Variable**: `assocOps`

- **完整名称**: assocOps
- **文件位置**: 行 674-674
- **可见性**: private

- **类型**: `[]clause.Association`







**Struct**: `setCreateOrUpdateG`

- **完整名称**: setCreateOrUpdateG
- **文件位置**: 行 695-699
- **可见性**: private








- **子元素**: c assigns assocOps 

**Method**: `Update`

- **完整名称**: Update
- **文件位置**: 行 701-717
- **可见性**: public


- **返回值**: `int`

- **参数**:
  - `ctx`: context.Context






**Variable**: `r`

- **完整名称**: r
- **文件位置**: 行 711-711
- **可见性**: private

- **类型**: `T`







**Method**: `Create`

- **完整名称**: Create
- **文件位置**: 行 719-738
- **可见性**: public


- **返回值**: `error`

- **参数**:
  - `ctx`: context.Context






**Variable**: `r`

- **完整名称**: r
- **文件位置**: 行 733-733
- **可见性**: private

- **类型**: `T`







**Method**: `executeAssociationOperation`

- **完整名称**: executeAssociationOperation
- **文件位置**: 行 741-753
- **可见性**: private
- **描述**: executeAssociationOperation executes an association operation

- **返回值**: `error`

- **参数**:
  - `ctx`: context.Context
  - `op`: clause.Association






**Variable**: `r`

- **完整名称**: r
- **文件位置**: 行 742-742
- **可见性**: private

- **类型**: `T`







**Method**: `handleAssociationCreate`

- **完整名称**: handleAssociationCreate
- **文件位置**: 行 755-769
- **可见性**: private


- **返回值**: `error`

- **参数**:
  - `ctx`: context.Context
  - `base`: *DB
  - `op`: clause.Association






**Method**: `handleAssociationForOwners`

- **完整名称**: handleAssociationForOwners
- **文件位置**: 行 772-789
- **可见性**: private
- **描述**: handleAssociationForOwners is a helper function that handles associations for all owners

- **返回值**: `error`

- **参数**:
  - `base`: *DB
  - `ctx`: context.Context
  - `handler`: func(...)
  - `associationName`: string






**Variable**: `owners`

- **完整名称**: owners
- **文件位置**: 行 773-773
- **可见性**: private

- **类型**: `[]T`







**Method**: `handleAssociation`

- **完整名称**: handleAssociation
- **文件位置**: 行 791-908
- **可见性**: private


- **返回值**: `error`

- **参数**:
  - `ctx`: context.Context
  - `base`: *DB
  - `op`: clause.Association






**Variable**: `rel`

- **完整名称**: rel
- **文件位置**: 行 798-798
- **可见性**: private

- **类型**: `assoc.Relationship`




- **值**: `assoc.Relationship`


**Variable**: `assocModel`

- **完整名称**: assocModel
- **文件位置**: 行 799-799
- **可见性**: private

- **类型**: `interface{}`




- **值**: `reflect.New(...).Interface(...)`


**Variable**: `fkNil`

- **完整名称**: fkNil
- **文件位置**: 行 800-800
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...`


**Variable**: `setMap`

- **完整名称**: setMap
- **文件位置**: 行 801-801
- **可见性**: private

- **类型**: `interface{}`




- **值**: `make(...)`


**Variable**: `ownerPKNames`

- **完整名称**: ownerPKNames
- **文件位置**: 行 802-802
- **可见性**: private

- **类型**: `[]string`







**Variable**: `ownerFKNames`

- **完整名称**: ownerFKNames
- **文件位置**: 行 803-803
- **可见性**: private

- **类型**: `[]string`







**Variable**: `primaryColumns`

- **完整名称**: primaryColumns
- **文件位置**: 行 804-804
- **可见性**: private

- **类型**: `[]any`







**Variable**: `foreignColumns`

- **完整名称**: foreignColumns
- **文件位置**: 行 805-805
- **可见性**: private

- **类型**: `[]any`










### gorm.go

- **语言**: go
- **包**: gorm
- **代码行数**: 405
- **元素数量**: 31


#### 导入


- `context`

- `database/sql`

- `fmt`

- `reflect`

- `sort`

- `sync`

- `time`

- `gorm.io/gorm/clause`

- `gorm.io/gorm/logger`

- `gorm.io/gorm/schema`




#### 代码元素


**Constant**: `preparedStmtDBKey`

- **完整名称**: preparedStmtDBKey
- **文件位置**: 行 18-18
- **可见性**: private

- **类型**: `interface{}`




- **值**: `"preparedStmt"`


**Struct**: `Config`

- **完整名称**: Config
- **文件位置**: 行 21-75
- **可见性**: public








- **子元素**: SkipDefaultTransaction DefaultTransactionTimeout DefaultContextTimeout NamingStrategy FullSaveAssociations Logger NowFunc DryRun PrepareStmt PrepareStmtMaxSize PrepareStmtTTL DisableAutomaticPing DisableForeignKeyConstraintWhenMigrating IgnoreRelationshipsWhenMigrating DisableNestedTransaction AllowGlobalUpdate QueryFields CreateBatchSize TranslateError PropagateUnscoped ClauseBuilders ConnPool Plugins callbacks cacheStore 

**Method**: `Apply`

- **完整名称**: Apply
- **文件位置**: 行 78-83
- **可见性**: public
- **描述**: Apply update config to new config

- **返回值**: `error`

- **参数**:
  - `config`: *Config






**Method**: `AfterInitialize`

- **完整名称**: AfterInitialize
- **文件位置**: 行 86-95
- **可见性**: public
- **描述**: AfterInitialize initialize plugins after db connected

- **返回值**: `error`

- **参数**:
  - `db`: *DB






**Interface**: `Option`

- **完整名称**: Option
- **文件位置**: 行 98-101
- **可见性**: public








- **子元素**: Apply AfterInitialize 

**Struct**: `DB`

- **完整名称**: DB
- **文件位置**: 行 104-110
- **可见性**: public








- **子元素**: Error RowsAffected Statement clone 

**Struct**: `Session`

- **完整名称**: Session
- **文件位置**: 行 113-129
- **可见性**: public








- **子元素**: DryRun PrepareStmt NewDB Initialized SkipHooks SkipDefaultTransaction DisableNestedTransaction AllowGlobalUpdate FullSaveAssociations PropagateUnscoped QueryFields Context Logger NowFunc CreateBatchSize 

**Function**: `Open`

- **完整名称**: Open
- **文件位置**: 行 132-247
- **可见性**: public
- **描述**: Open initialize db session based on dialector

- **返回值**: `*DB`

- **参数**:
  - `dialector`: Dialector
  - `opts`: ...Option






**Variable**: `skipAfterInitialize`

- **完整名称**: skipAfterInitialize
- **文件位置**: 行 149-149
- **可见性**: private

- **类型**: `bool`







**Method**: `Session`

- **完整名称**: Session
- **文件位置**: 行 250-349
- **可见性**: public
- **描述**: Session create new db session

- **返回值**: `*DB`

- **参数**:
  - `config`: *Session






**Variable**: `txConfig`

- **完整名称**: txConfig
- **文件位置**: 行 252-252
- **可见性**: private

- **类型**: `*db.Config`




- **值**: `...`


**Variable**: `tx`

- **完整名称**: tx
- **文件位置**: 行 253-253
- **可见性**: private

- **类型**: `interface{}`




- **值**: `&...`


**Variable**: `preparedStmt`

- **完整名称**: preparedStmt
- **文件位置**: 行 290-290
- **可见性**: private

- **类型**: `*PreparedStmtDB`







**Method**: `WithContext`

- **完整名称**: WithContext
- **文件位置**: 行 352-354
- **可见性**: public
- **描述**: WithContext change current instance db's context to ctx

- **返回值**: `*DB`

- **参数**:
  - `ctx`: context.Context






**Method**: `Debug`

- **完整名称**: Debug
- **文件位置**: 行 357-362
- **可见性**: public
- **描述**: Debug start debug mode

- **返回值**: `*DB`






**Method**: `Set`

- **完整名称**: Set
- **文件位置**: 行 365-369
- **可见性**: public
- **描述**: Set store value with key into current db instance's context

- **返回值**: `*DB`

- **参数**:
  - `key`: string
  - `value`: interface{}






**Method**: `Get`

- **完整名称**: Get
- **文件位置**: 行 372-374
- **可见性**: public
- **描述**: Get get value with key from current db instance's context

- **返回值**: `interface{}`

- **参数**:
  - `key`: string






**Method**: `InstanceSet`

- **完整名称**: InstanceSet
- **文件位置**: 行 377-381
- **可见性**: public
- **描述**: InstanceSet store value with key into current db instance's context

- **返回值**: `*DB`

- **参数**:
  - `key`: string
  - `value`: interface{}






**Method**: `InstanceGet`

- **完整名称**: InstanceGet
- **文件位置**: 行 384-386
- **可见性**: public
- **描述**: InstanceGet get value with key from current db instance's context

- **返回值**: `interface{}`

- **参数**:
  - `key`: string






**Method**: `Callback`

- **完整名称**: Callback
- **文件位置**: 行 389-391
- **可见性**: public
- **描述**: Callback returns callback manager

- **返回值**: `*callbacks`






**Method**: `AddError`

- **完整名称**: AddError
- **文件位置**: 行 394-409
- **可见性**: public
- **描述**: AddError add error to db

- **返回值**: `error`

- **参数**:
  - `err`: error






**Method**: `DB`

- **完整名称**: DB
- **文件位置**: 行 412-432
- **可见性**: public
- **描述**: DB returns `*sql.DB`

- **返回值**: `*sql.DB`






**Method**: `getInstance`

- **完整名称**: getInstance
- **文件位置**: 行 434-461
- **可见性**: private


- **返回值**: `*DB`






**Function**: `Expr`

- **完整名称**: Expr
- **文件位置**: 行 464-466
- **可见性**: public
- **描述**: Expr returns clause.Expr, which can be used to pass SQL expression as params

- **返回值**: `clause.Expr`

- **参数**:
  - `expr`: string
  - `args`: ...interface{}






**Method**: `SetupJoinTable`

- **完整名称**: SetupJoinTable
- **文件位置**: 行 469-517
- **可见性**: public
- **描述**: SetupJoinTable setup join table schema

- **返回值**: `error`

- **参数**:
  - `model`: interface{}
  - `field`: string
  - `joinTable`: interface{}






**Variable**: `tx`

- **完整名称**: tx
- **文件位置**: 行 471-471
- **可见性**: private

- **类型**: `interface{}`




- **值**: `db.getInstance(...)`


**Variable**: `stmt`

- **完整名称**: stmt
- **文件位置**: 行 472-472
- **可见性**: private

- **类型**: `tx.Statement`




- **值**: `tx.Statement`


**Variable**: `modelSchema`

- **完整名称**: modelSchema
- **文件位置**: 行 473-473
- **可见性**: private

- **类型**: `*schema.Schema`







**Variable**: `joinSchema`

- **完整名称**: joinSchema
- **文件位置**: 行 473-473
- **可见性**: private

- **类型**: `*schema.Schema`







**Method**: `Use`

- **完整名称**: Use
- **文件位置**: 行 520-530
- **可见性**: public
- **描述**: Use use plugin

- **返回值**: `error`

- **参数**:
  - `plugin`: Plugin






**Method**: `ToSQL`

- **完整名称**: ToSQL
- **文件位置**: 行 540-545
- **可见性**: public
- **描述**: ToSQL for generate SQL string.

- **返回值**: `string`

- **参数**:
  - `queryFn`: func(...)









### interfaces.go

- **语言**: go
- **包**: gorm
- **代码行数**: 66
- **元素数量**: 13


#### 导入


- `context`

- `database/sql`

- `gorm.io/gorm/clause`

- `gorm.io/gorm/schema`




#### 代码元素


**Interface**: `Dialector`

- **完整名称**: Dialector
- **文件位置**: 行 12-21
- **可见性**: public








- **子元素**: Name Initialize Migrator DataTypeOf DefaultValueOf BindVarTo QuoteTo Explain 

**Interface**: `Plugin`

- **完整名称**: Plugin
- **文件位置**: 行 24-27
- **可见性**: public








- **子元素**: Name Initialize 

**Interface**: `ParamsFilter`

- **完整名称**: ParamsFilter
- **文件位置**: 行 29-31
- **可见性**: public








- **子元素**: ParamsFilter 

**Interface**: `ConnPool`

- **完整名称**: ConnPool
- **文件位置**: 行 34-39
- **可见性**: public








- **子元素**: PrepareContext ExecContext QueryContext QueryRowContext 

**Interface**: `SavePointerDialectorInterface`

- **完整名称**: SavePointerDialectorInterface
- **文件位置**: 行 42-45
- **可见性**: public








- **子元素**: SavePoint RollbackTo 

**Interface**: `TxBeginner`

- **完整名称**: TxBeginner
- **文件位置**: 行 48-50
- **可见性**: public








- **子元素**: BeginTx 

**Interface**: `ConnPoolBeginner`

- **完整名称**: ConnPoolBeginner
- **文件位置**: 行 53-55
- **可见性**: public








- **子元素**: BeginTx 

**Interface**: `TxCommitter`

- **完整名称**: TxCommitter
- **文件位置**: 行 58-61
- **可见性**: public








- **子元素**: Commit Rollback 

**Interface**: `Tx`

- **完整名称**: Tx
- **文件位置**: 行 64-68
- **可见性**: public








- **子元素**: StmtContext 

**Interface**: `Valuer`

- **完整名称**: Valuer
- **文件位置**: 行 71-73
- **可见性**: public








- **子元素**: GormValue 

**Interface**: `GetDBConnector`

- **完整名称**: GetDBConnector
- **文件位置**: 行 76-78
- **可见性**: public








- **子元素**: GetDBConn 

**Interface**: `Rows`

- **完整名称**: Rows
- **文件位置**: 行 81-88
- **可见性**: public








- **子元素**: Columns ColumnTypes Next Scan Err Close 

**Interface**: `ErrorTranslator`

- **完整名称**: ErrorTranslator
- **文件位置**: 行 90-92
- **可见性**: public








- **子元素**: Translate 




### internal\lru\lru.go

- **语言**: go
- **包**: lru
- **代码行数**: 338
- **元素数量**: 42


#### 导入


- `sync`

- `time`




#### 代码元素


**Class**: `EvictCallback`

- **完整名称**: EvictCallback
- **文件位置**: 行 11-11
- **可见性**: public









**Struct**: `LRU`

- **完整名称**: LRU
- **文件位置**: 行 14-29
- **可见性**: public








- **子元素**: size evictList items onEvict mu ttl done buckets nextCleanupBucket 

**Struct**: `bucket`

- **完整名称**: bucket
- **文件位置**: 行 32-35
- **可见性**: private








- **子元素**: entries newestEntry 

**Constant**: `noEvictionTTL`

- **完整名称**: noEvictionTTL
- **文件位置**: 行 38-38
- **可见性**: private

- **类型**: `interface{}`




- **值**: `time.Hour * 24 * 365 * 10`


**Constant**: `numBuckets`

- **完整名称**: numBuckets
- **文件位置**: 行 42-42
- **可见性**: private

- **类型**: `interface{}`




- **值**: `100`


**Function**: `NewLRU`

- **完整名称**: NewLRU
- **文件位置**: 行 51-93
- **可见性**: public
- **描述**: NewLRU returns a new thread-safe cache with expirable entries.

- **返回值**: `*interface{}`

- **参数**:
  - `size`: int
  - `onEvict`: interface{}
  - `ttl`: time.Duration






**Method**: `Purge`

- **完整名称**: Purge
- **文件位置**: 行 97-112
- **可见性**: public
- **描述**: Purge clears the cache completely.








**Method**: `Add`

- **完整名称**: Add
- **文件位置**: 行 117-143
- **可见性**: public
- **描述**: Add adds a value to the cache. Returns true if an eviction occurred.

- **返回值**: `bool`

- **参数**:
  - `key`: K
  - `value`: V






**Method**: `Get`

- **完整名称**: Get
- **文件位置**: 行 146-159
- **可见性**: public
- **描述**: Get looks up a key's value from the cache.

- **返回值**: `V`

- **参数**:
  - `key`: K






**Variable**: `ent`

- **完整名称**: ent
- **文件位置**: 行 149-149
- **可见性**: private

- **类型**: `*interface{}`







**Method**: `Contains`

- **完整名称**: Contains
- **文件位置**: 行 163-168
- **可见性**: public
- **描述**: Contains checks if a key is in the cache, without updating the recent-ness

- **返回值**: `bool`

- **参数**:
  - `key`: K






**Method**: `Peek`

- **完整名称**: Peek
- **文件位置**: 行 172-184
- **可见性**: public
- **描述**: Peek returns the key value (or undefined if not found) without updating

- **返回值**: `V`

- **参数**:
  - `key`: K






**Variable**: `ent`

- **完整名称**: ent
- **文件位置**: 行 175-175
- **可见性**: private

- **类型**: `*interface{}`







**Method**: `Remove`

- **完整名称**: Remove
- **文件位置**: 行 188-196
- **可见性**: public
- **描述**: Remove removes the provided key from the cache, returning if the

- **返回值**: `bool`

- **参数**:
  - `key`: K






**Method**: `RemoveOldest`

- **完整名称**: RemoveOldest
- **文件位置**: 行 199-207
- **可见性**: public
- **描述**: RemoveOldest removes the oldest item from the cache.

- **返回值**: `K`






**Method**: `GetOldest`

- **完整名称**: GetOldest
- **文件位置**: 行 210-217
- **可见性**: public
- **描述**: GetOldest returns the oldest entry

- **返回值**: `K`






**Method**: `KeyValues`

- **完整名称**: KeyValues
- **文件位置**: 行 219-232
- **可见性**: public


- **返回值**: `map[K]V`






**Method**: `Keys`

- **完整名称**: Keys
- **文件位置**: 行 236-248
- **可见性**: public
- **描述**: Keys returns a slice of the keys in the cache, from oldest to newest.

- **返回值**: `[]K`






**Method**: `Values`

- **完整名称**: Values
- **文件位置**: 行 252-264
- **可见性**: public
- **描述**: Values returns a slice of the values in the cache, from oldest to newest.

- **返回值**: `[]V`






**Method**: `Len`

- **完整名称**: Len
- **文件位置**: 行 267-271
- **可见性**: public
- **描述**: Len returns the number of items in the cache.

- **返回值**: `int`






**Method**: `Resize`

- **完整名称**: Resize
- **文件位置**: 行 274-290
- **可见性**: public
- **描述**: Resize changes the cache size. Size of 0 means unlimited.

- **返回值**: `int`

- **参数**:
  - `size`: int






**Method**: `removeOldest`

- **完整名称**: removeOldest
- **文件位置**: 行 305-309
- **可见性**: private
- **描述**: removeOldest removes the oldest item from the cache. Has to be called with lock!








**Method**: `removeElement`

- **完整名称**: removeElement
- **文件位置**: 行 312-319
- **可见性**: private
- **描述**: removeElement is used to remove a given list element from the cache. Has to be called with lock!



- **参数**:
  - `e`: *interface{}






**Method**: `deleteExpired`

- **完整名称**: deleteExpired
- **文件位置**: 行 323-338
- **可见性**: private
- **描述**: deleteExpired deletes expired records from the oldest bucket, waiting for the newest entry








**Method**: `addToBucket`

- **完整名称**: addToBucket
- **文件位置**: 行 341-348
- **可见性**: private
- **描述**: addToBucket adds entry to expire bucket so that it will be cleaned up when the time comes. Has to be called with lock!



- **参数**:
  - `e`: *interface{}






**Method**: `removeFromBucket`

- **完整名称**: removeFromBucket
- **文件位置**: 行 351-353
- **可见性**: private
- **描述**: removeFromBucket removes the entry from its corresponding bucket. Has to be called with lock!



- **参数**:
  - `e`: *interface{}






**Method**: `Cap`

- **完整名称**: Cap
- **文件位置**: 行 356-358
- **可见性**: public
- **描述**: Cap returns the capacity of the cache

- **返回值**: `int`






**Struct**: `Entry`

- **完整名称**: Entry
- **文件位置**: 行 361-383
- **可见性**: public








- **子元素**: next prev list Key Value ExpiresAt ExpireBucket 

**Method**: `PrevEntry`

- **完整名称**: PrevEntry
- **文件位置**: 行 386-391
- **可见性**: public
- **描述**: PrevEntry returns the previous list element or nil.

- **返回值**: `*interface{}`






**Struct**: `LruList`

- **完整名称**: LruList
- **文件位置**: 行 395-398
- **可见性**: public








- **子元素**: root len 

**Method**: `Init`

- **完整名称**: Init
- **文件位置**: 行 401-406
- **可见性**: public
- **描述**: Init initializes or clears list l.

- **返回值**: `*interface{}`






**Function**: `NewList`

- **完整名称**: NewList
- **文件位置**: 行 409-409
- **可见性**: public
- **描述**: NewList returns an initialized list.

- **返回值**: `*interface{}`






**Method**: `Length`

- **完整名称**: Length
- **文件位置**: 行 413-413
- **可见性**: public
- **描述**: Length returns the number of elements of list l.

- **返回值**: `int`






**Method**: `Back`

- **完整名称**: Back
- **文件位置**: 行 416-421
- **可见性**: public
- **描述**: Back returns the last element of list l or nil if the list is empty.

- **返回值**: `*interface{}`






**Method**: `lazyInit`

- **完整名称**: lazyInit
- **文件位置**: 行 424-428
- **可见性**: private
- **描述**: lazyInit lazily initializes a zero List Value.








**Method**: `insert`

- **完整名称**: insert
- **文件位置**: 行 431-439
- **可见性**: private
- **描述**: insert inserts e after at, increments l.len, and returns e.

- **返回值**: `*interface{}`

- **参数**:
  - `e`: *interface{}
  - `at`: *interface{}






**Method**: `insertValue`

- **完整名称**: insertValue
- **文件位置**: 行 442-444
- **可见性**: private
- **描述**: insertValue is a convenience wrapper for insert(&Entry{Value: v, ExpiresAt: ExpiresAt}, at).

- **返回值**: `*interface{}`

- **参数**:
  - `k`: K
  - `v`: V
  - `expiresAt`: time.Time
  - `at`: *interface{}






**Method**: `Remove`

- **完整名称**: Remove
- **文件位置**: 行 447-456
- **可见性**: public
- **描述**: Remove removes e from its list, decrements l.len

- **返回值**: `V`

- **参数**:
  - `e`: *interface{}






**Method**: `move`

- **完整名称**: move
- **文件位置**: 行 459-470
- **可见性**: private
- **描述**: move moves e to next to at.



- **参数**:
  - `e`: *interface{}
  - `at`: *interface{}






**Method**: `PushFront`

- **完整名称**: PushFront
- **文件位置**: 行 473-476
- **可见性**: public
- **描述**: PushFront inserts a new element e with value v at the front of list l and returns e.

- **返回值**: `*interface{}`

- **参数**:
  - `k`: K
  - `v`: V






**Method**: `PushFrontExpirable`

- **完整名称**: PushFrontExpirable
- **文件位置**: 行 479-482
- **可见性**: public
- **描述**: PushFrontExpirable inserts a new expirable element e with Value v at the front of list l and returns e.

- **返回值**: `*interface{}`

- **参数**:
  - `k`: K
  - `v`: V
  - `expiresAt`: time.Time






**Method**: `MoveToFront`

- **完整名称**: MoveToFront
- **文件位置**: 行 487-493
- **可见性**: public
- **描述**: MoveToFront moves element e to the front of list l.



- **参数**:
  - `e`: *interface{}









### internal\stmt_store\stmt_store.go

- **语言**: go
- **包**: stmt_store
- **代码行数**: 87
- **元素数量**: 14


#### 导入


- `context`

- `database/sql`

- `math`

- `sync`

- `time`

- `gorm.io/gorm/internal/lru`




#### 代码元素


**Struct**: `Stmt`

- **完整名称**: Stmt
- **文件位置**: 行 13-18
- **可见性**: public








- **子元素**: Transaction prepared prepareErr 

**Method**: `Error`

- **完整名称**: Error
- **文件位置**: 行 20-22
- **可见性**: public


- **返回值**: `error`






**Method**: `Close`

- **完整名称**: Close
- **文件位置**: 行 24-31
- **可见性**: public


- **返回值**: `error`






**Interface**: `Store`

- **完整名称**: Store
- **文件位置**: 行 36-70
- **可见性**: public








- **子元素**: New Keys Get Set Delete 

**Constant**: `defaultMaxSize`

- **完整名称**: defaultMaxSize
- **文件位置**: 行 77-77
- **可见性**: private

- **类型**: `math.MaxInt`




- **值**: `math.MaxInt`


**Constant**: `defaultTTL`

- **完整名称**: defaultTTL
- **文件位置**: 行 80-80
- **可见性**: private
- **描述**: defaultTTL defines the default time-to-live (TTL) for each cache entry.
- **类型**: `interface{}`




- **值**: `time.Hour * 24`


**Function**: `New`

- **完整名称**: New
- **文件位置**: 行 98-113
- **可见性**: public
- **描述**: New creates and returns a new Store instance.

- **返回值**: `Store`

- **参数**:
  - `size`: int
  - `ttl`: time.Duration






**Struct**: `lruStore`

- **完整名称**: lruStore
- **文件位置**: 行 115-117
- **可见性**: private








- **子元素**: lru 

**Method**: `Keys`

- **完整名称**: Keys
- **文件位置**: 行 119-121
- **可见性**: public


- **返回值**: `[]string`






**Method**: `Get`

- **完整名称**: Get
- **文件位置**: 行 123-129
- **可见性**: public


- **返回值**: `*Stmt`

- **参数**:
  - `key`: string






**Method**: `Set`

- **完整名称**: Set
- **文件位置**: 行 131-133
- **可见性**: public




- **参数**:
  - `key`: string
  - `value`: *Stmt






**Method**: `Delete`

- **完整名称**: Delete
- **文件位置**: 行 135-137
- **可见性**: public




- **参数**:
  - `key`: string






**Interface**: `ConnPool`

- **完整名称**: ConnPool
- **文件位置**: 行 139-141
- **可见性**: public








- **子元素**: PrepareContext 

**Method**: `New`

- **完整名称**: New
- **文件位置**: 行 157-183
- **可见性**: public
- **描述**: New creates a new Stmt object for executing SQL queries.

- **返回值**: `*Stmt`

- **参数**:
  - `ctx`: context.Context
  - `key`: string
  - `isTransaction`: bool
  - `conn`: ConnPool
  - `locker`: sync.Locker









### logger\logger.go

- **语言**: go
- **包**: logger
- **代码行数**: 174
- **元素数量**: 43


#### 导入


- `context`

- `errors`

- `fmt`

- `io`

- `log`

- `os`

- `time`

- `gorm.io/gorm/utils`




#### 代码元素


**Variable**: `ErrRecordNotFound`

- **完整名称**: ErrRecordNotFound
- **文件位置**: 行 16-16
- **可见性**: public

- **类型**: `interface{}`




- **值**: `errors.New(...)`


**Constant**: `Reset`

- **完整名称**: Reset
- **文件位置**: 行 20-20
- **可见性**: public

- **类型**: `interface{}`




- **值**: `"\033[0m"`


**Constant**: `Red`

- **完整名称**: Red
- **文件位置**: 行 21-21
- **可见性**: public

- **类型**: `interface{}`




- **值**: `"\033[31m"`


**Constant**: `Green`

- **完整名称**: Green
- **文件位置**: 行 22-22
- **可见性**: public

- **类型**: `interface{}`




- **值**: `"\033[32m"`


**Constant**: `Yellow`

- **完整名称**: Yellow
- **文件位置**: 行 23-23
- **可见性**: public

- **类型**: `interface{}`




- **值**: `"\033[33m"`


**Constant**: `Blue`

- **完整名称**: Blue
- **文件位置**: 行 24-24
- **可见性**: public

- **类型**: `interface{}`




- **值**: `"\033[34m"`


**Constant**: `Magenta`

- **完整名称**: Magenta
- **文件位置**: 行 25-25
- **可见性**: public

- **类型**: `interface{}`




- **值**: `"\033[35m"`


**Constant**: `Cyan`

- **完整名称**: Cyan
- **文件位置**: 行 26-26
- **可见性**: public

- **类型**: `interface{}`




- **值**: `"\033[36m"`


**Constant**: `White`

- **完整名称**: White
- **文件位置**: 行 27-27
- **可见性**: public

- **类型**: `interface{}`




- **值**: `"\033[37m"`


**Constant**: `BlueBold`

- **完整名称**: BlueBold
- **文件位置**: 行 28-28
- **可见性**: public

- **类型**: `interface{}`




- **值**: `"\033[34;1m"`


**Constant**: `MagentaBold`

- **完整名称**: MagentaBold
- **文件位置**: 行 29-29
- **可见性**: public

- **类型**: `interface{}`




- **值**: `"\033[35;1m"`


**Constant**: `RedBold`

- **完整名称**: RedBold
- **文件位置**: 行 30-30
- **可见性**: public

- **类型**: `interface{}`




- **值**: `"\033[31;1m"`


**Constant**: `YellowBold`

- **完整名称**: YellowBold
- **文件位置**: 行 31-31
- **可见性**: public

- **类型**: `interface{}`




- **值**: `"\033[33;1m"`


**Class**: `LogLevel`

- **完整名称**: LogLevel
- **文件位置**: 行 35-35
- **可见性**: public









**Constant**: `Silent`

- **完整名称**: Silent
- **文件位置**: 行 39-39
- **可见性**: public
- **描述**: Silent silent log level
- **类型**: `LogLevel`




- **值**: `iota + 1`


**Constant**: `Error`

- **完整名称**: Error
- **文件位置**: 行 41-41
- **可见性**: public
- **描述**: Error error log level








**Constant**: `Warn`

- **完整名称**: Warn
- **文件位置**: 行 43-43
- **可见性**: public
- **描述**: Warn warn log level








**Constant**: `Info`

- **完整名称**: Info
- **文件位置**: 行 45-45
- **可见性**: public
- **描述**: Info info log level








**Interface**: `Writer`

- **完整名称**: Writer
- **文件位置**: 行 49-51
- **可见性**: public








- **子元素**: Printf 

**Struct**: `Config`

- **完整名称**: Config
- **文件位置**: 行 54-60
- **可见性**: public








- **子元素**: SlowThreshold Colorful IgnoreRecordNotFoundError ParameterizedQueries LogLevel 

**Interface**: `Interface`

- **完整名称**: Interface
- **文件位置**: 行 63-69
- **可见性**: public








- **子元素**: LogMode Info Warn Error Trace 

**Variable**: `Discard`

- **完整名称**: Discard
- **文件位置**: 行 73-73
- **可见性**: public
- **描述**: Discard logger will print any log to io.Discard
- **类型**: `interface{}`




- **值**: `New(...)`


**Variable**: `Default`

- **完整名称**: Default
- **文件位置**: 行 75-75
- **可见性**: public
- **描述**: Default Default logger
- **类型**: `interface{}`




- **值**: `New(...)`


**Variable**: `Recorder`

- **完整名称**: Recorder
- **文件位置**: 行 82-82
- **可见性**: public
- **描述**: Recorder logger records running SQL into a recorder instance
- **类型**: `interface{}`




- **值**: `...`


**Variable**: `RecorderParamsFilter`

- **完整名称**: RecorderParamsFilter
- **文件位置**: 行 85-85
- **可见性**: public
- **描述**: RecorderParamsFilter defaults to no-op, allows to be run-over by a different implementation
- **类型**: `interface{}`




- **值**: `...`


**Function**: `New`

- **完整名称**: New
- **文件位置**: 行 91-120
- **可见性**: public
- **描述**: New initialize logger

- **返回值**: `Interface`

- **参数**:
  - `writer`: Writer
  - `config`: Config






**Variable**: `infoStr`

- **完整名称**: infoStr
- **文件位置**: 行 93-93
- **可见性**: private

- **类型**: `interface{}`




- **值**: `"%s\n[info] "`


**Variable**: `warnStr`

- **完整名称**: warnStr
- **文件位置**: 行 94-94
- **可见性**: private

- **类型**: `interface{}`




- **值**: `"%s\n[warn] "`


**Variable**: `errStr`

- **完整名称**: errStr
- **文件位置**: 行 95-95
- **可见性**: private

- **类型**: `interface{}`




- **值**: `"%s\n[error] "`


**Variable**: `traceStr`

- **完整名称**: traceStr
- **文件位置**: 行 96-96
- **可见性**: private

- **类型**: `interface{}`




- **值**: `"%s\n[%.3fms] [rows:%v] %s"`


**Variable**: `traceWarnStr`

- **完整名称**: traceWarnStr
- **文件位置**: 行 97-97
- **可见性**: private

- **类型**: `interface{}`




- **值**: `"%s %s\n[%.3fms] [rows:%v] %s"`


**Variable**: `traceErrStr`

- **完整名称**: traceErrStr
- **文件位置**: 行 98-98
- **可见性**: private

- **类型**: `interface{}`




- **值**: `"%s %s\n[%.3fms] [rows:%v] %s"`


**Struct**: `logger`

- **完整名称**: logger
- **文件位置**: 行 122-127
- **可见性**: private








- **子元素**: infoStr warnStr errStr traceStr traceErrStr traceWarnStr 

**Method**: `LogMode`

- **完整名称**: LogMode
- **文件位置**: 行 130-134
- **可见性**: public
- **描述**: LogMode log mode

- **返回值**: `Interface`

- **参数**:
  - `level`: LogLevel






**Method**: `Info`

- **完整名称**: Info
- **文件位置**: 行 137-141
- **可见性**: public
- **描述**: Info print info



- **参数**:
  - `ctx`: context.Context
  - `msg`: string
  - `data`: ...interface{}






**Method**: `Warn`

- **完整名称**: Warn
- **文件位置**: 行 144-148
- **可见性**: public
- **描述**: Warn print warn messages



- **参数**:
  - `ctx`: context.Context
  - `msg`: string
  - `data`: ...interface{}






**Method**: `Error`

- **完整名称**: Error
- **文件位置**: 行 151-155
- **可见性**: public
- **描述**: Error print error messages



- **参数**:
  - `ctx`: context.Context
  - `msg`: string
  - `data`: ...interface{}






**Method**: `Trace`

- **完整名称**: Trace
- **文件位置**: 行 160-190
- **可见性**: public
- **描述**: Trace print sql message



- **参数**:
  - `ctx`: context.Context
  - `begin`: time.Time
  - `fc`: func(...)
  - `err`: error






**Method**: `ParamsFilter`

- **完整名称**: ParamsFilter
- **文件位置**: 行 193-198
- **可见性**: public
- **描述**: ParamsFilter filter params

- **返回值**: `string`

- **参数**:
  - `ctx`: context.Context
  - `sql`: string
  - `params`: ...interface{}






**Struct**: `traceRecorder`

- **完整名称**: traceRecorder
- **文件位置**: 行 200-206
- **可见性**: private








- **子元素**: BeginAt SQL RowsAffected Err 

**Method**: `New`

- **完整名称**: New
- **文件位置**: 行 209-211
- **可见性**: public
- **描述**: New trace recorder

- **返回值**: `*traceRecorder`






**Method**: `Trace`

- **完整名称**: Trace
- **文件位置**: 行 214-218
- **可见性**: public
- **描述**: Trace implement logger interface



- **参数**:
  - `ctx`: context.Context
  - `begin`: time.Time
  - `fc`: func(...)
  - `err`: error






**Method**: `ParamsFilter`

- **完整名称**: ParamsFilter
- **文件位置**: 行 220-225
- **可见性**: public


- **返回值**: `string`

- **参数**:
  - `ctx`: context.Context
  - `sql`: string
  - `params`: ...interface{}









### logger\slog.go

- **语言**: go
- **包**: logger
- **代码行数**: 95
- **元素数量**: 9


#### 导入


- `context`

- `errors`

- `fmt`

- `log/slog`

- `time`

- `gorm.io/gorm/utils`




#### 代码元素


**Struct**: `slogLogger`

- **完整名称**: slogLogger
- **文件位置**: 行 15-22
- **可见性**: private








- **子元素**: Logger LogLevel SlowThreshold Parameterized Colorful IgnoreRecordNotFoundError 

**Function**: `NewSlogLogger`

- **完整名称**: NewSlogLogger
- **文件位置**: 行 24-32
- **可见性**: public


- **返回值**: `Interface`

- **参数**:
  - `logger`: *slog.Logger
  - `config`: Config






**Method**: `LogMode`

- **完整名称**: LogMode
- **文件位置**: 行 34-38
- **可见性**: public


- **返回值**: `Interface`

- **参数**:
  - `level`: LogLevel






**Method**: `Info`

- **完整名称**: Info
- **文件位置**: 行 40-44
- **可见性**: public




- **参数**:
  - `ctx`: context.Context
  - `msg`: string
  - `data`: ...interface{}






**Method**: `Warn`

- **完整名称**: Warn
- **文件位置**: 行 46-50
- **可见性**: public




- **参数**:
  - `ctx`: context.Context
  - `msg`: string
  - `data`: ...interface{}






**Method**: `Error`

- **完整名称**: Error
- **文件位置**: 行 52-56
- **可见性**: public




- **参数**:
  - `ctx`: context.Context
  - `msg`: string
  - `data`: ...interface{}






**Method**: `Trace`

- **完整名称**: Trace
- **文件位置**: 行 58-94
- **可见性**: public




- **参数**:
  - `ctx`: context.Context
  - `begin`: time.Time
  - `fc`: func(...)
  - `err`: error






**Method**: `log`

- **完整名称**: log
- **文件位置**: 行 96-108
- **可见性**: private




- **参数**:
  - `ctx`: context.Context
  - `level`: slog.Level
  - `msg`: string
  - `args`: ...any






**Method**: `ParamsFilter`

- **完整名称**: ParamsFilter
- **文件位置**: 行 111-116
- **可见性**: public
- **描述**: ParamsFilter filter params

- **返回值**: `string`

- **参数**:
  - `ctx`: context.Context
  - `sql`: string
  - `params`: ...interface{}









### logger\sql.go

- **语言**: go
- **包**: logger
- **代码行数**: 161
- **元素数量**: 10


#### 导入


- `database/sql/driver`

- `fmt`

- `reflect`

- `regexp`

- `strconv`

- `strings`

- `time`

- `unicode`

- `gorm.io/gorm/utils`




#### 代码元素


**Constant**: `tmFmtWithMS`

- **完整名称**: tmFmtWithMS
- **文件位置**: 行 17-17
- **可见性**: private

- **类型**: `interface{}`




- **值**: `"2006-01-02 15:04:05.999"`


**Constant**: `tmFmtZero`

- **完整名称**: tmFmtZero
- **文件位置**: 行 18-18
- **可见性**: private

- **类型**: `interface{}`




- **值**: `"0000-00-00 00:00:00"`


**Constant**: `nullStr`

- **完整名称**: nullStr
- **文件位置**: 行 19-19
- **可见性**: private

- **类型**: `interface{}`




- **值**: `"NULL"`


**Variable**: `convertibleTypes`

- **完整名称**: convertibleTypes
- **文件位置**: 行 32-32
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...`


**Variable**: `numericPlaceholderRe`

- **完整名称**: numericPlaceholderRe
- **文件位置**: 行 35-35
- **可见性**: private

- **类型**: `interface{}`




- **值**: `regexp.MustCompile(...)`


**Function**: `ExplainSQL`

- **完整名称**: ExplainSQL
- **文件位置**: 行 51-181
- **可见性**: public
- **描述**: ExplainSQL generate SQL string with given parameters, the generated SQL is expected to be used in logger, execute it might introduce a SQL injection vulnerability

- **返回值**: `string`

- **参数**:
  - `sql`: string
  - `numericPlaceholder`: *regexp.Regexp
  - `escaper`: string
  - `avars`: ...interface{}






**Variable**: `convertParams`

- **完整名称**: convertParams
- **文件位置**: 行 53-53
- **可见性**: private

- **类型**: `func(...)`







**Variable**: `vars`

- **完整名称**: vars
- **文件位置**: 行 54-54
- **可见性**: private

- **类型**: `interface{}`




- **值**: `make(...)`


**Variable**: `idx`

- **完整名称**: idx
- **文件位置**: 行 149-149
- **可见性**: private

- **类型**: `int`







**Variable**: `newSQL`

- **完整名称**: newSQL
- **文件位置**: 行 150-150
- **可见性**: private

- **类型**: `strings.Builder`










### migrator\column_type.go

- **语言**: go
- **包**: migrator
- **代码行数**: 75
- **元素数量**: 13


#### 导入


- `database/sql`

- `reflect`




#### 代码元素


**Struct**: `ColumnType`

- **完整名称**: ColumnType
- **文件位置**: 行 9-24
- **可见性**: public








- **子元素**: SQLColumnType NameValue DataTypeValue ColumnTypeValue PrimaryKeyValue UniqueValue AutoIncrementValue LengthValue DecimalSizeValue ScaleValue NullableValue ScanTypeValue CommentValue DefaultValueValue 

**Method**: `Name`

- **完整名称**: Name
- **文件位置**: 行 27-32
- **可见性**: public
- **描述**: Name returns the name or alias of the column.

- **返回值**: `string`






**Method**: `DatabaseTypeName`

- **完整名称**: DatabaseTypeName
- **文件位置**: 行 40-45
- **可见性**: public
- **描述**: DatabaseTypeName returns the database system name of the column type. If an empty

- **返回值**: `string`






**Method**: `ColumnType`

- **完整名称**: ColumnType
- **文件位置**: 行 48-50
- **可见性**: public
- **描述**: ColumnType returns the database type of the column. like `varchar(16)`

- **返回值**: `string`






**Method**: `PrimaryKey`

- **完整名称**: PrimaryKey
- **文件位置**: 行 53-55
- **可见性**: public
- **描述**: PrimaryKey returns the column is primary key or not.

- **返回值**: `bool`






**Method**: `AutoIncrement`

- **完整名称**: AutoIncrement
- **文件位置**: 行 58-60
- **可见性**: public
- **描述**: AutoIncrement returns the column is auto increment or not.

- **返回值**: `bool`






**Method**: `Length`

- **完整名称**: Length
- **文件位置**: 行 63-68
- **可见性**: public
- **描述**: Length returns the column type length for variable length column types

- **返回值**: `int64`






**Method**: `DecimalSize`

- **完整名称**: DecimalSize
- **文件位置**: 行 71-76
- **可见性**: public
- **描述**: DecimalSize returns the scale and precision of a decimal type.

- **返回值**: `int64`






**Method**: `Nullable`

- **完整名称**: Nullable
- **文件位置**: 行 79-84
- **可见性**: public
- **描述**: Nullable reports whether the column may be null.

- **返回值**: `bool`






**Method**: `Unique`

- **完整名称**: Unique
- **文件位置**: 行 87-89
- **可见性**: public
- **描述**: Unique reports whether the column may be unique.

- **返回值**: `bool`






**Method**: `ScanType`

- **完整名称**: ScanType
- **文件位置**: 行 92-97
- **可见性**: public
- **描述**: ScanType returns a Go type suitable for scanning into using Rows.Scan.

- **返回值**: `reflect.Type`






**Method**: `Comment`

- **完整名称**: Comment
- **文件位置**: 行 100-102
- **可见性**: public
- **描述**: Comment returns the comment of current column.

- **返回值**: `string`






**Method**: `DefaultValue`

- **完整名称**: DefaultValue
- **文件位置**: 行 105-107
- **可见性**: public
- **描述**: DefaultValue returns the default value of current column.

- **返回值**: `string`









### migrator\index.go

- **语言**: go
- **包**: migrator
- **代码行数**: 28
- **元素数量**: 7


#### 导入


- `database/sql`




#### 代码元素


**Struct**: `Index`

- **完整名称**: Index
- **文件位置**: 行 6-13
- **可见性**: public








- **子元素**: TableName NameValue ColumnList PrimaryKeyValue UniqueValue OptionValue 

**Method**: `Table`

- **完整名称**: Table
- **文件位置**: 行 16-18
- **可见性**: public
- **描述**: Table return the table name of the index.

- **返回值**: `string`






**Method**: `Name`

- **完整名称**: Name
- **文件位置**: 行 21-23
- **可见性**: public
- **描述**: Name return the name of the index.

- **返回值**: `string`






**Method**: `Columns`

- **完整名称**: Columns
- **文件位置**: 行 26-28
- **可见性**: public
- **描述**: Columns return the columns of the index

- **返回值**: `[]string`






**Method**: `PrimaryKey`

- **完整名称**: PrimaryKey
- **文件位置**: 行 31-33
- **可见性**: public
- **描述**: PrimaryKey returns the index is primary key or not.

- **返回值**: `bool`






**Method**: `Unique`

- **完整名称**: Unique
- **文件位置**: 行 36-38
- **可见性**: public
- **描述**: Unique returns whether the index is unique or not.

- **返回值**: `bool`






**Method**: `Option`

- **完整名称**: Option
- **文件位置**: 行 41-43
- **可见性**: public
- **描述**: Option return the optional attribute of the index

- **返回值**: `string`









### migrator\migrator.go

- **语言**: go
- **包**: migrator
- **代码行数**: 812
- **元素数量**: 67


#### 导入


- `context`

- `database/sql`

- `errors`

- `fmt`

- `reflect`

- `regexp`

- `strconv`

- `strings`

- `time`

- `gorm.io/gorm`

- `gorm.io/gorm/clause`

- `gorm.io/gorm/logger`

- `gorm.io/gorm/schema`




#### 代码元素


**Variable**: `regFullDataType`

- **完整名称**: regFullDataType
- **文件位置**: 行 27-27
- **可见性**: private

- **类型**: `interface{}`




- **值**: `regexp.MustCompile(...)`


**Variable**: `_`

- **完整名称**: _
- **文件位置**: 行 31-31
- **可见性**: private

- **类型**: `gorm.Migrator`




- **值**: `(...)(...)`


**Struct**: `Migrator`

- **完整名称**: Migrator
- **文件位置**: 行 34-36
- **可见性**: public









**Struct**: `Config`

- **完整名称**: Config
- **文件位置**: 行 39-43
- **可见性**: public








- **子元素**: CreateIndexAfterCreateTable DB 

**Struct**: `printSQLLogger`

- **完整名称**: printSQLLogger
- **文件位置**: 行 45-47
- **可见性**: private









**Method**: `Trace`

- **完整名称**: Trace
- **文件位置**: 行 49-53
- **可见性**: public




- **参数**:
  - `ctx`: context.Context
  - `begin`: time.Time
  - `fc`: func(...)
  - `err`: error






**Interface**: `GormDataTypeInterface`

- **完整名称**: GormDataTypeInterface
- **文件位置**: 行 56-58
- **可见性**: public








- **子元素**: GormDBDataType 

**Method**: `RunWithValue`

- **完整名称**: RunWithValue
- **文件位置**: 行 61-75
- **可见性**: public
- **描述**: RunWithValue run migration with statement value

- **返回值**: `error`

- **参数**:
  - `value`: interface{}
  - `fc`: func(...)






**Method**: `DataTypeOf`

- **完整名称**: DataTypeOf
- **文件位置**: 行 78-87
- **可见性**: public
- **描述**: DataTypeOf return field's db data type

- **返回值**: `string`

- **参数**:
  - `field`: *schema.Field






**Method**: `FullDataTypeOf`

- **完整名称**: FullDataTypeOf
- **文件位置**: 行 90-108
- **可见性**: public
- **描述**: FullDataTypeOf returns field's db full data type

- **返回值**: `clause.Expr`

- **参数**:
  - `field`: *schema.Field






**Method**: `GetQueryAndExecTx`

- **完整名称**: GetQueryAndExecTx
- **文件位置**: 行 110-118
- **可见性**: public


- **返回值**: `*gorm.DB`






**Method**: `AutoMigrate`

- **完整名称**: AutoMigrate
- **文件位置**: 行 121-205
- **可见性**: public
- **描述**: AutoMigrate auto migrate values

- **返回值**: `error`

- **参数**:
  - `values`: ...interface{}






**Variable**: `parseIndexes`

- **完整名称**: parseIndexes
- **文件位置**: 行 140-140
- **可见性**: private

- **类型**: `interface{}`




- **值**: `stmt.Schema.ParseIndexes(...)`


**Variable**: `parseCheckConstraints`

- **完整名称**: parseCheckConstraints
- **文件位置**: 行 141-141
- **可见性**: private

- **类型**: `interface{}`




- **值**: `stmt.Schema.ParseCheckConstraints(...)`


**Variable**: `foundColumn`

- **完整名称**: foundColumn
- **文件位置**: 行 144-144
- **可见性**: private

- **类型**: `gorm.ColumnType`







**Method**: `GetTables`

- **完整名称**: GetTables
- **文件位置**: 行 208-212
- **可见性**: public
- **描述**: GetTables returns tables

- **返回值**: `[]string`






**Method**: `CreateTable`

- **完整名称**: CreateTable
- **文件位置**: 行 215-316
- **可见性**: public
- **描述**: CreateTable create table in database for values

- **返回值**: `error`

- **参数**:
  - `values`: ...interface{}






**Variable**: `createTableSQL`

- **完整名称**: createTableSQL
- **文件位置**: 行 225-225
- **可见性**: private

- **类型**: `interface{}`




- **值**: `"CREATE TABLE ? ("`


**Variable**: `values`

- **完整名称**: values
- **文件位置**: 行 226-226
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...`


**Variable**: `hasPrimaryKeyInDataType`

- **完整名称**: hasPrimaryKeyInDataType
- **文件位置**: 行 227-227
- **可见性**: private

- **类型**: `bool`







**Method**: `DropTable`

- **完整名称**: DropTable
- **文件位置**: 行 319-330
- **可见性**: public
- **描述**: DropTable drop table for values

- **返回值**: `error`

- **参数**:
  - `values`: ...interface{}






**Method**: `HasTable`

- **完整名称**: HasTable
- **文件位置**: 行 333-342
- **可见性**: public
- **描述**: HasTable returns table exists or not for value, value could be a struct or string

- **返回值**: `bool`

- **参数**:
  - `value`: interface{}






**Variable**: `count`

- **完整名称**: count
- **文件位置**: 行 334-334
- **可见性**: private

- **类型**: `int64`







**Method**: `RenameTable`

- **完整名称**: RenameTable
- **文件位置**: 行 345-370
- **可见性**: public
- **描述**: RenameTable rename table from oldName to newName

- **返回值**: `error`

- **参数**:
  - `oldName`: interface{}
  - `newName`: interface{}






**Variable**: `oldTable`

- **完整名称**: oldTable
- **文件位置**: 行 346-346
- **可见性**: private

- **类型**: `interface{}`







**Variable**: `newTable`

- **完整名称**: newTable
- **文件位置**: 行 346-346
- **可见性**: private

- **类型**: `interface{}`







**Method**: `AddColumn`

- **完整名称**: AddColumn
- **文件位置**: 行 373-393
- **可见性**: public
- **描述**: AddColumn create `name` column for value

- **返回值**: `error`

- **参数**:
  - `value`: interface{}
  - `name`: string






**Method**: `DropColumn`

- **完整名称**: DropColumn
- **文件位置**: 行 396-408
- **可见性**: public
- **描述**: DropColumn drop value's `name` column

- **返回值**: `error`

- **参数**:
  - `value`: interface{}
  - `name`: string






**Method**: `AlterColumn`

- **完整名称**: AlterColumn
- **文件位置**: 行 411-425
- **可见性**: public
- **描述**: AlterColumn alter value's `field` column' type based on schema definition

- **返回值**: `error`

- **参数**:
  - `value`: interface{}
  - `field`: string






**Method**: `HasColumn`

- **完整名称**: HasColumn
- **文件位置**: 行 428-446
- **可见性**: public
- **描述**: HasColumn check has column `field` for value or not

- **返回值**: `bool`

- **参数**:
  - `value`: interface{}
  - `field`: string






**Variable**: `count`

- **完整名称**: count
- **文件位置**: 行 429-429
- **可见性**: private

- **类型**: `int64`







**Method**: `RenameColumn`

- **完整名称**: RenameColumn
- **文件位置**: 行 449-466
- **可见性**: public
- **描述**: RenameColumn rename value's field name from oldName to newName

- **返回值**: `error`

- **参数**:
  - `value`: interface{}
  - `oldName`: string
  - `newName`: string






**Method**: `MigrateColumn`

- **完整名称**: MigrateColumn
- **文件位置**: 行 469-592
- **可见性**: public
- **描述**: MigrateColumn migrate column

- **返回值**: `error`

- **参数**:
  - `value`: interface{}
  - `field`: *schema.Field
  - `columnType`: gorm.ColumnType






**Variable**: `alterColumn`

- **完整名称**: alterColumn
- **文件位置**: 行 478-478
- **可见性**: private

- **类型**: `bool`







**Variable**: `isSameType`

- **完整名称**: isSameType
- **文件位置**: 行 479-479
- **可见性**: private

- **类型**: `interface{}`




- **值**: `fullDataType == realDataType`


**Method**: `MigrateColumnUnique`

- **完整名称**: MigrateColumnUnique
- **文件位置**: 行 594-612
- **可见性**: public


- **返回值**: `error`

- **参数**:
  - `value`: interface{}
  - `field`: *schema.Field
  - `columnType`: gorm.ColumnType






**Method**: `ColumnTypes`

- **完整名称**: ColumnTypes
- **文件位置**: 行 615-641
- **可见性**: public
- **描述**: ColumnTypes return columnTypes []gorm.ColumnType and execErr error

- **返回值**: `[]gorm.ColumnType`

- **参数**:
  - `value`: interface{}






**Variable**: `rawColumnTypes`

- **完整名称**: rawColumnTypes
- **文件位置**: 行 627-627
- **可见性**: private

- **类型**: `[]*sql.ColumnType`







**Method**: `CreateView`

- **完整名称**: CreateView
- **文件位置**: 行 655-676
- **可见性**: public
- **描述**: CreateView create view from Query in gorm.ViewOption.

- **返回值**: `error`

- **参数**:
  - `name`: string
  - `option`: gorm.ViewOption






**Method**: `DropView`

- **完整名称**: DropView
- **文件位置**: 行 679-681
- **可见性**: public
- **描述**: DropView drop view

- **返回值**: `error`

- **参数**:
  - `name`: string






**Method**: `GuessConstraintAndTable`

- **完整名称**: GuessConstraintAndTable
- **文件位置**: 行 686-696
- **可见性**: public
- **描述**: GuessConstraintAndTable guess statement's constraint and it's table based on name

- **返回值**: `*schema.Constraint`

- **参数**:
  - `stmt`: *gorm.Statement
  - `name`: string






**Method**: `GuessConstraintInterfaceAndTable`

- **完整名称**: GuessConstraintInterfaceAndTable
- **文件位置**: 行 700-754
- **可见性**: public
- **描述**: GuessConstraintInterfaceAndTable guess statement's constraint and it's table based on name

- **返回值**: `schema.ConstraintInterface`

- **参数**:
  - `stmt`: *gorm.Statement
  - `name`: string






**Method**: `CreateConstraint`

- **完整名称**: CreateConstraint
- **文件位置**: 行 757-770
- **可见性**: public
- **描述**: CreateConstraint create constraint

- **返回值**: `error`

- **参数**:
  - `value`: interface{}
  - `name`: string






**Method**: `DropConstraint`

- **完整名称**: DropConstraint
- **文件位置**: 行 773-781
- **可见性**: public
- **描述**: DropConstraint drop constraint

- **返回值**: `error`

- **参数**:
  - `value`: interface{}
  - `name`: string






**Method**: `HasConstraint`

- **完整名称**: HasConstraint
- **文件位置**: 行 784-800
- **可见性**: public
- **描述**: HasConstraint check has constraint or not

- **返回值**: `bool`

- **参数**:
  - `value`: interface{}
  - `name`: string






**Variable**: `count`

- **完整名称**: count
- **文件位置**: 行 785-785
- **可见性**: private

- **类型**: `int64`







**Method**: `BuildIndexOptions`

- **完整名称**: BuildIndexOptions
- **文件位置**: 行 803-822
- **可见性**: public
- **描述**: BuildIndexOptions build index options

- **返回值**: `[]interface{}`

- **参数**:
  - `opts`: []schema.IndexOption
  - `stmt`: *gorm.Statement






**Interface**: `BuildIndexOptionsInterface`

- **完整名称**: BuildIndexOptionsInterface
- **文件位置**: 行 825-827
- **可见性**: public








- **子元素**: BuildIndexOptions 

**Method**: `CreateIndex`

- **完整名称**: CreateIndex
- **文件位置**: 行 830-862
- **可见性**: public
- **描述**: CreateIndex create index `name`

- **返回值**: `error`

- **参数**:
  - `value`: interface{}
  - `name`: string






**Method**: `DropIndex`

- **完整名称**: DropIndex
- **文件位置**: 行 865-875
- **可见性**: public
- **描述**: DropIndex drop index `name`

- **返回值**: `error`

- **参数**:
  - `value`: interface{}
  - `name`: string






**Method**: `HasIndex`

- **完整名称**: HasIndex
- **文件位置**: 行 878-895
- **可见性**: public
- **描述**: HasIndex check has index `name` or not

- **返回值**: `bool`

- **参数**:
  - `value`: interface{}
  - `name`: string






**Variable**: `count`

- **完整名称**: count
- **文件位置**: 行 879-879
- **可见性**: private

- **类型**: `int64`







**Method**: `RenameIndex`

- **完整名称**: RenameIndex
- **文件位置**: 行 898-905
- **可见性**: public
- **描述**: RenameIndex rename index from oldName to newName

- **返回值**: `error`

- **参数**:
  - `value`: interface{}
  - `oldName`: string
  - `newName`: string






**Method**: `CurrentDatabase`

- **完整名称**: CurrentDatabase
- **文件位置**: 行 908-911
- **可见性**: public
- **描述**: CurrentDatabase returns current database name

- **返回值**: `string`






**Method**: `ReorderModels`

- **完整名称**: ReorderModels
- **文件位置**: 行 914-1015
- **可见性**: public
- **描述**: ReorderModels reorder models according to constraint dependencies

- **返回值**: `[]interface{}`

- **参数**:
  - `values`: []interface{}
  - `autoAdd`: bool






**Struct**: `Dependency`

- **完整名称**: Dependency
- **文件位置**: 行 915-918
- **可见性**: public








- **子元素**: Depends 

**Variable**: `modelNames`

- **完整名称**: modelNames
- **文件位置**: 行 921-921
- **可见性**: private

- **类型**: `[]string`







**Variable**: `orderedModelNames`

- **完整名称**: orderedModelNames
- **文件位置**: 行 921-921
- **可见性**: private

- **类型**: `[]string`







**Variable**: `orderedModelNamesMap`

- **完整名称**: orderedModelNamesMap
- **文件位置**: 行 922-922
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...`


**Variable**: `parsedSchemas`

- **完整名称**: parsedSchemas
- **文件位置**: 行 923-923
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...`


**Variable**: `valuesMap`

- **完整名称**: valuesMap
- **文件位置**: 行 924-924
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...`


**Variable**: `insertIntoOrderedList`

- **完整名称**: insertIntoOrderedList
- **文件位置**: 行 925-925
- **可见性**: private

- **类型**: `func(...)`







**Variable**: `parseDependence`

- **完整名称**: parseDependence
- **文件位置**: 行 926-926
- **可见性**: private

- **类型**: `func(...)`







**Method**: `CurrentTable`

- **完整名称**: CurrentTable
- **文件位置**: 行 1018-1023
- **可见性**: public
- **描述**: CurrentTable returns current statement's table expression

- **返回值**: `interface{}`

- **参数**:
  - `stmt`: *gorm.Statement






**Method**: `GetIndexes`

- **完整名称**: GetIndexes
- **文件位置**: 行 1026-1028
- **可见性**: public
- **描述**: GetIndexes return Indexes []gorm.Index and execErr error

- **返回值**: `[]gorm.Index`

- **参数**:
  - `dst`: interface{}






**Method**: `GetTypeAliases`

- **完整名称**: GetTypeAliases
- **文件位置**: 行 1031-1033
- **可见性**: public
- **描述**: GetTypeAliases return database type aliases

- **返回值**: `[]string`

- **参数**:
  - `databaseTypeName`: string






**Method**: `TableType`

- **完整名称**: TableType
- **文件位置**: 行 1036-1038
- **可见性**: public
- **描述**: TableType return tableType gorm.TableType and execErr error

- **返回值**: `gorm.TableType`

- **参数**:
  - `dst`: interface{}









### migrator\table_type.go

- **语言**: go
- **包**: migrator
- **代码行数**: 22
- **元素数量**: 5


#### 导入


- `database/sql`




#### 代码元素


**Struct**: `TableType`

- **完整名称**: TableType
- **文件位置**: 行 8-13
- **可见性**: public








- **子元素**: SchemaValue NameValue TypeValue CommentValue 

**Method**: `Schema`

- **完整名称**: Schema
- **文件位置**: 行 16-18
- **可见性**: public
- **描述**: Schema returns the schema of the table.

- **返回值**: `string`






**Method**: `Name`

- **完整名称**: Name
- **文件位置**: 行 21-23
- **可见性**: public
- **描述**: Name returns the name of the table.

- **返回值**: `string`






**Method**: `Type`

- **完整名称**: Type
- **文件位置**: 行 26-28
- **可见性**: public
- **描述**: Type returns the type of the table.

- **返回值**: `string`






**Method**: `Comment`

- **完整名称**: Comment
- **文件位置**: 行 31-33
- **可见性**: public
- **描述**: Comment returns the comment of current table.

- **返回值**: `string`









### migrator.go

- **语言**: go
- **包**: gorm
- **代码行数**: 79
- **元素数量**: 7


#### 导入


- `reflect`

- `gorm.io/gorm/clause`

- `gorm.io/gorm/schema`




#### 代码元素


**Method**: `Migrator`

- **完整名称**: Migrator
- **文件位置**: 行 11-20
- **可见性**: public
- **描述**: Migrator returns migrator

- **返回值**: `Migrator`






**Method**: `AutoMigrate`

- **完整名称**: AutoMigrate
- **文件位置**: 行 23-25
- **可见性**: public
- **描述**: AutoMigrate run auto migration for given models

- **返回值**: `error`

- **参数**:
  - `dst`: ...interface{}






**Struct**: `ViewOption`

- **完整名称**: ViewOption
- **文件位置**: 行 28-32
- **可见性**: public








- **子元素**: Replace CheckOption Query 

**Interface**: `ColumnType`

- **完整名称**: ColumnType
- **文件位置**: 行 35-48
- **可见性**: public








- **子元素**: Name DatabaseTypeName ColumnType PrimaryKey AutoIncrement Length DecimalSize Nullable Unique ScanType Comment DefaultValue 

**Interface**: `Index`

- **完整名称**: Index
- **文件位置**: 行 50-57
- **可见性**: public








- **子元素**: Table Name Columns PrimaryKey Unique Option 

**Interface**: `TableType`

- **完整名称**: TableType
- **文件位置**: 行 60-65
- **可见性**: public








- **子元素**: Schema Name Type Comment 

**Interface**: `Migrator`

- **完整名称**: Migrator
- **文件位置**: 行 68-111
- **可见性**: public








- **子元素**: AutoMigrate CurrentDatabase FullDataTypeOf GetTypeAliases CreateTable DropTable HasTable RenameTable GetTables TableType AddColumn DropColumn AlterColumn MigrateColumn MigrateColumnUnique HasColumn RenameColumn ColumnTypes CreateView DropView CreateConstraint DropConstraint HasConstraint CreateIndex DropIndex HasIndex RenameIndex GetIndexes 




### model.go

- **语言**: go
- **包**: gorm
- **代码行数**: 8
- **元素数量**: 1


#### 导入


- `time`




#### 代码元素


**Struct**: `Model`

- **完整名称**: Model
- **文件位置**: 行 11-16
- **可见性**: public








- **子元素**: ID CreatedAt UpdatedAt DeletedAt 




### prepare_stmt.go

- **语言**: go
- **包**: gorm
- **代码行数**: 165
- **元素数量**: 19


#### 导入


- `context`

- `database/sql`

- `database/sql/driver`

- `errors`

- `reflect`

- `sync`

- `time`

- `gorm.io/gorm/internal/stmt_store`




#### 代码元素


**Struct**: `PreparedStmtDB`

- **完整名称**: PreparedStmtDB
- **文件位置**: 行 15-19
- **可见性**: public








- **子元素**: Stmts Mux 

**Function**: `NewPreparedStmtDB`

- **完整名称**: NewPreparedStmtDB
- **文件位置**: 行 30-36
- **可见性**: public
- **描述**: NewPreparedStmtDB creates and initializes a new instance of PreparedStmtDB.

- **返回值**: `*PreparedStmtDB`

- **参数**:
  - `connPool`: ConnPool
  - `maxSize`: int
  - `ttl`: time.Duration






**Method**: `GetDBConn`

- **完整名称**: GetDBConn
- **文件位置**: 行 39-49
- **可见性**: public
- **描述**: GetDBConn returns the underlying *sql.DB connection

- **返回值**: `*sql.DB`






**Method**: `Close`

- **完整名称**: Close
- **文件位置**: 行 52-59
- **可见性**: public
- **描述**: Close closes all prepared statements in the store








**Method**: `Reset`

- **完整名称**: Reset
- **文件位置**: 行 62-64
- **可见性**: public
- **描述**: Reset Deprecated use Close instead








**Method**: `prepare`

- **完整名称**: prepare
- **文件位置**: 行 66-86
- **可见性**: private


- **返回值**: `*stmt_store.Stmt`

- **参数**:
  - `ctx`: context.Context
  - `conn`: ConnPool
  - `isTransaction`: bool
  - `query`: string






**Method**: `BeginTx`

- **完整名称**: BeginTx
- **文件位置**: 行 88-107
- **可见性**: public


- **返回值**: `ConnPool`

- **参数**:
  - `ctx`: context.Context
  - `opt`: *sql.TxOptions






**Method**: `ExecContext`

- **完整名称**: ExecContext
- **文件位置**: 行 109-118
- **可见性**: public


- **返回值**: `sql.Result`

- **参数**:
  - `ctx`: context.Context
  - `query`: string
  - `args`: ...interface{}






**Method**: `QueryContext`

- **完整名称**: QueryContext
- **文件位置**: 行 120-129
- **可见性**: public


- **返回值**: `*sql.Rows`

- **参数**:
  - `ctx`: context.Context
  - `query`: string
  - `args`: ...interface{}






**Method**: `QueryRowContext`

- **完整名称**: QueryRowContext
- **文件位置**: 行 131-137
- **可见性**: public


- **返回值**: `*sql.Row`

- **参数**:
  - `ctx`: context.Context
  - `query`: string
  - `args`: ...interface{}






**Method**: `Ping`

- **完整名称**: Ping
- **文件位置**: 行 139-145
- **可见性**: public


- **返回值**: `error`






**Struct**: `PreparedStmtTX`

- **完整名称**: PreparedStmtTX
- **文件位置**: 行 147-150
- **可见性**: public








- **子元素**: PreparedStmtDB 

**Method**: `GetDBConn`

- **完整名称**: GetDBConn
- **文件位置**: 行 152-154
- **可见性**: public


- **返回值**: `*sql.DB`






**Method**: `Commit`

- **完整名称**: Commit
- **文件位置**: 行 156-161
- **可见性**: public


- **返回值**: `error`






**Method**: `Rollback`

- **完整名称**: Rollback
- **文件位置**: 行 163-168
- **可见性**: public


- **返回值**: `error`






**Method**: `ExecContext`

- **完整名称**: ExecContext
- **文件位置**: 行 170-179
- **可见性**: public


- **返回值**: `sql.Result`

- **参数**:
  - `ctx`: context.Context
  - `query`: string
  - `args`: ...interface{}






**Method**: `QueryContext`

- **完整名称**: QueryContext
- **文件位置**: 行 181-190
- **可见性**: public


- **返回值**: `*sql.Rows`

- **参数**:
  - `ctx`: context.Context
  - `query`: string
  - `args`: ...interface{}






**Method**: `QueryRowContext`

- **完整名称**: QueryRowContext
- **文件位置**: 行 192-198
- **可见性**: public


- **返回值**: `*sql.Row`

- **参数**:
  - `ctx`: context.Context
  - `query`: string
  - `args`: ...interface{}






**Method**: `Ping`

- **完整名称**: Ping
- **文件位置**: 行 200-206
- **可见性**: public


- **返回值**: `error`









### scan.go

- **语言**: go
- **包**: gorm
- **代码行数**: 312
- **元素数量**: 22


#### 导入


- `database/sql`

- `database/sql/driver`

- `reflect`

- `strings`

- `time`

- `gorm.io/gorm/schema`

- `gorm.io/gorm/utils`




#### 代码元素


**Method**: `scanIntoStruct`

- **完整名称**: scanIntoStruct
- **文件位置**: 行 54-113
- **可见性**: private




- **参数**:
  - `rows`: Rows
  - `reflectValue`: reflect.Value
  - `values`: []interface{}
  - `fields`: []*schema.Field
  - `joinFields`: [][]*schema.Field






**Variable**: `isNilPtrValue`

- **完整名称**: isNilPtrValue
- **文件位置**: 行 78-78
- **可见性**: private

- **类型**: `bool`







**Variable**: `relValue`

- **完整名称**: relValue
- **文件位置**: 行 79-79
- **可见性**: private

- **类型**: `reflect.Value`







**Class**: `ScanMode`

- **完整名称**: ScanMode
- **文件位置**: 行 116-116
- **可见性**: public









**Constant**: `ScanInitialized`

- **完整名称**: ScanInitialized
- **文件位置**: 行 120-120
- **可见性**: public

- **类型**: `ScanMode`




- **值**: `1 << 0`


**Constant**: `ScanUpdate`

- **完整名称**: ScanUpdate
- **文件位置**: 行 121-121
- **可见性**: public

- **类型**: `ScanMode`




- **值**: `1 << 1`


**Constant**: `ScanOnConflictDoNothing`

- **完整名称**: ScanOnConflictDoNothing
- **文件位置**: 行 122-122
- **可见性**: public

- **类型**: `ScanMode`




- **值**: `1 << 2`


**Function**: `Scan`

- **完整名称**: Scan
- **文件位置**: 行 126-369
- **可见性**: public
- **描述**: Scan scan rows into db statement



- **参数**:
  - `rows`: Rows
  - `db`: *DB
  - `mode`: ScanMode






**Variable**: `columns`

- **完整名称**: columns
- **文件位置**: 行 128-128
- **可见性**: private

- **类型**: `interface{}`




- **值**: `rows.Columns(...)`


**Variable**: `_`

- **完整名称**: _
- **文件位置**: 行 128-128
- **可见性**: private

- **类型**: `interface{}`




- **值**: `rows.Columns(...)`


**Variable**: `values`

- **完整名称**: values
- **文件位置**: 行 129-129
- **可见性**: private

- **类型**: `interface{}`




- **值**: `make(...)`


**Variable**: `initialized`

- **完整名称**: initialized
- **文件位置**: 行 130-130
- **可见性**: private

- **类型**: `interface{}`




- **值**: `mode & ScanInitialized != 0`


**Variable**: `update`

- **完整名称**: update
- **文件位置**: 行 131-131
- **可见性**: private

- **类型**: `interface{}`




- **值**: `mode & ScanUpdate != 0`


**Variable**: `onConflictDonothing`

- **完整名称**: onConflictDonothing
- **文件位置**: 行 132-132
- **可见性**: private

- **类型**: `interface{}`




- **值**: `mode & ScanOnConflictDoNothing != 0`


**Variable**: `fields`

- **完整名称**: fields
- **文件位置**: 行 192-192
- **可见性**: private

- **类型**: `interface{}`




- **值**: `make(...)`


**Variable**: `joinFields`

- **完整名称**: joinFields
- **文件位置**: 行 193-193
- **可见性**: private

- **类型**: `[][]*schema.Field`







**Variable**: `sch`

- **完整名称**: sch
- **文件位置**: 行 194-194
- **可见性**: private

- **类型**: `db.Statement.Schema`




- **值**: `db.Statement.Schema`


**Variable**: `reflectValue`

- **完整名称**: reflectValue
- **文件位置**: 行 195-195
- **可见性**: private

- **类型**: `db.Statement.ReflectValue`




- **值**: `db.Statement.ReflectValue`


**Variable**: `val`

- **完整名称**: val
- **文件位置**: 行 278-278
- **可见性**: private

- **类型**: `interface{}`







**Variable**: `val`

- **完整名称**: val
- **文件位置**: 行 281-281
- **可见性**: private

- **类型**: `interface{}`







**Variable**: `elem`

- **完整名称**: elem
- **文件位置**: 行 291-291
- **可见性**: private

- **类型**: `reflect.Value`







**Variable**: `isArrayKind`

- **完整名称**: isArrayKind
- **文件位置**: 行 292-292
- **可见性**: private

- **类型**: `interface{}`




- **值**: `reflectValue.Kind(...) == reflect.Array`





### schema\constraint.go

- **语言**: go
- **包**: schema
- **代码行数**: 51
- **元素数量**: 9


#### 导入


- `regexp`

- `strings`

- `gorm.io/gorm/clause`




#### 代码元素


**Variable**: `regEnLetterAndMidline`

- **完整名称**: regEnLetterAndMidline
- **文件位置**: 行 11-11
- **可见性**: private

- **类型**: `interface{}`




- **值**: `regexp.MustCompile(...)`


**Struct**: `CheckConstraint`

- **完整名称**: CheckConstraint
- **文件位置**: 行 13-17
- **可见性**: public








- **子元素**: Name Constraint 

**Method**: `GetName`

- **完整名称**: GetName
- **文件位置**: 行 19-19
- **可见性**: public


- **返回值**: `string`






**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 21-23
- **可见性**: public


- **返回值**: `string`






**Method**: `ParseCheckConstraints`

- **完整名称**: ParseCheckConstraints
- **文件位置**: 行 26-43
- **可见性**: public
- **描述**: ParseCheckConstraints parse schema check constraints

- **返回值**: `map[string]CheckConstraint`






**Struct**: `UniqueConstraint`

- **完整名称**: UniqueConstraint
- **文件位置**: 行 45-48
- **可见性**: public








- **子元素**: Name Field 

**Method**: `GetName`

- **完整名称**: GetName
- **文件位置**: 行 50-50
- **可见性**: public


- **返回值**: `string`






**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 52-54
- **可见性**: public


- **返回值**: `string`






**Method**: `ParseUniqueConstraints`

- **完整名称**: ParseUniqueConstraints
- **文件位置**: 行 57-66
- **可见性**: public
- **描述**: ParseUniqueConstraints parse schema unique constraints

- **返回值**: `map[string]UniqueConstraint`









### schema\field.go

- **语言**: go
- **包**: schema
- **代码行数**: 913
- **元素数量**: 31


#### 导入


- `context`

- `database/sql`

- `database/sql/driver`

- `fmt`

- `reflect`

- `strconv`

- `strings`

- `sync`

- `time`

- `github.com/jinzhu/now`

- `gorm.io/gorm/clause`

- `gorm.io/gorm/utils`




#### 代码元素


**Variable**: `TimeReflectType`

- **完整名称**: TimeReflectType
- **文件位置**: 行 21-21
- **可见性**: public

- **类型**: `interface{}`




- **值**: `reflect.TypeOf(...)`


**Variable**: `TimePtrReflectType`

- **完整名称**: TimePtrReflectType
- **文件位置**: 行 22-22
- **可见性**: public

- **类型**: `interface{}`




- **值**: `reflect.TypeOf(...)`


**Variable**: `ByteReflectType`

- **完整名称**: ByteReflectType
- **文件位置**: 行 23-23
- **可见性**: public

- **类型**: `interface{}`




- **值**: `reflect.TypeOf(...)`


**Class**: `DataType`

- **完整名称**: DataType
- **文件位置**: 行 28-28
- **可见性**: public
- **描述**: DataType GORM data type








**Class**: `TimeType`

- **完整名称**: TimeType
- **文件位置**: 行 30-30
- **可见性**: public
- **描述**: TimeType GORM time type








**Constant**: `UnixTime`

- **完整名称**: UnixTime
- **文件位置**: 行 35-35
- **可见性**: public

- **类型**: `TimeType`




- **值**: `1`


**Constant**: `UnixSecond`

- **完整名称**: UnixSecond
- **文件位置**: 行 36-36
- **可见性**: public

- **类型**: `TimeType`




- **值**: `2`


**Constant**: `UnixMillisecond`

- **完整名称**: UnixMillisecond
- **文件位置**: 行 37-37
- **可见性**: public

- **类型**: `TimeType`




- **值**: `3`


**Constant**: `UnixNanosecond`

- **完整名称**: UnixNanosecond
- **文件位置**: 行 38-38
- **可见性**: public

- **类型**: `TimeType`




- **值**: `4`


**Constant**: `Bool`

- **完整名称**: Bool
- **文件位置**: 行 43-43
- **可见性**: public

- **类型**: `DataType`




- **值**: `"bool"`


**Constant**: `Int`

- **完整名称**: Int
- **文件位置**: 行 44-44
- **可见性**: public

- **类型**: `DataType`




- **值**: `"int"`


**Constant**: `Uint`

- **完整名称**: Uint
- **文件位置**: 行 45-45
- **可见性**: public

- **类型**: `DataType`




- **值**: `"uint"`


**Constant**: `Float`

- **完整名称**: Float
- **文件位置**: 行 46-46
- **可见性**: public

- **类型**: `DataType`




- **值**: `"float"`


**Constant**: `String`

- **完整名称**: String
- **文件位置**: 行 47-47
- **可见性**: public

- **类型**: `DataType`




- **值**: `"string"`


**Constant**: `Time`

- **完整名称**: Time
- **文件位置**: 行 48-48
- **可见性**: public

- **类型**: `DataType`




- **值**: `"time"`


**Constant**: `Bytes`

- **完整名称**: Bytes
- **文件位置**: 行 49-49
- **可见性**: public

- **类型**: `DataType`




- **值**: `"bytes"`


**Constant**: `DefaultAutoIncrementIncrement`

- **完整名称**: DefaultAutoIncrementIncrement
- **文件位置**: 行 52-52
- **可见性**: public

- **类型**: `int64`




- **值**: `1`


**Struct**: `Field`

- **完整名称**: Field
- **文件位置**: 行 55-99
- **可见性**: public








- **子元素**: Name DBName BindNames EmbeddedBindNames DataType GORMDataType PrimaryKey AutoIncrement AutoIncrementIncrement Creatable Updatable Readable AutoCreateTime AutoUpdateTime HasDefaultValue DefaultValue DefaultValueInterface NotNull Unique Comment Size Precision Scale IgnoreMigration FieldType IndirectFieldType StructField Tag TagSettings Schema EmbeddedSchema OwnerSchema ReflectValueOf ValueOf Set Serializer NewValuePool UniqueIndex 

**Method**: `BindName`

- **完整名称**: BindName
- **文件位置**: 行 101-103
- **可见性**: public


- **返回值**: `string`






**Method**: `ParseField`

- **完整名称**: ParseField
- **文件位置**: 行 106-448
- **可见性**: public
- **描述**: ParseField parses reflect.StructField to Field

- **返回值**: `*Field`

- **参数**:
  - `fieldStruct`: reflect.StructField






**Variable**: `err`

- **完整名称**: err
- **文件位置**: 行 108-108
- **可见性**: private

- **类型**: `error`







**Variable**: `tagSetting`

- **完整名称**: tagSetting
- **文件位置**: 行 109-109
- **可见性**: private

- **类型**: `interface{}`




- **值**: `ParseTagSetting(...)`


**Variable**: `getRealFieldValue`

- **完整名称**: getRealFieldValue
- **文件位置**: 行 149-149
- **可见性**: private

- **类型**: `func(...)`







**Variable**: `rv`

- **完整名称**: rv
- **文件位置**: 行 152-152
- **可见性**: private

- **类型**: `interface{}`




- **值**: `reflect.Indirect(...)`


**Variable**: `rvType`

- **完整名称**: rvType
- **文件位置**: 行 153-153
- **可见性**: private

- **类型**: `interface{}`




- **值**: `rv.Type(...)`


**Variable**: `err`

- **完整名称**: err
- **文件位置**: 行 394-394
- **可见性**: private

- **类型**: `error`







**Method**: `setupValuerAndSetter`

- **完整名称**: setupValuerAndSetter
- **文件位置**: 行 451-990
- **可见性**: private
- **描述**: create valuer, setter when parse struct



- **参数**:
  - `modelType`: reflect.Type






**Variable**: `oldFieldSetter`

- **完整名称**: oldFieldSetter
- **文件位置**: 行 952-952
- **可见性**: private

- **类型**: `field.Set`




- **值**: `field.Set`


**Variable**: `sameElemType`

- **完整名称**: sameElemType
- **文件位置**: 行 953-953
- **可见性**: private

- **类型**: `bool`







**Variable**: `sameType`

- **完整名称**: sameType
- **文件位置**: 行 954-954
- **可见性**: private

- **类型**: `interface{}`




- **值**: `field.FieldType == reflect.ValueOf(...).Type(...)`


**Method**: `setupNewValuePool`

- **完整名称**: setupNewValuePool
- **文件位置**: 行 992-1011
- **可见性**: private












### schema\index.go

- **语言**: go
- **包**: schema
- **代码行数**: 149
- **元素数量**: 12


#### 导入


- `fmt`

- `sort`

- `strconv`

- `strings`




#### 代码元素


**Struct**: `Index`

- **完整名称**: Index
- **文件位置**: 行 10-18
- **可见性**: public








- **子元素**: Name Class Type Where Comment Option Fields 

**Struct**: `IndexOption`

- **完整名称**: IndexOption
- **文件位置**: 行 20-27
- **可见性**: public








- **子元素**: Expression Sort Collate Length Priority 

**Method**: `ParseIndexes`

- **完整名称**: ParseIndexes
- **文件位置**: 行 30-78
- **可见性**: public
- **描述**: ParseIndexes parse schema indexes

- **返回值**: `[]*Index`






**Method**: `LookIndex`

- **完整名称**: LookIndex
- **文件位置**: 行 80-97
- **可见性**: public


- **返回值**: `*Index`

- **参数**:
  - `name`: string






**Variable**: `name`

- **完整名称**: name
- **文件位置**: 行 106-106
- **可见性**: private

- **类型**: `string`







**Variable**: `tag`

- **完整名称**: tag
- **文件位置**: 行 107-107
- **可见性**: private

- **类型**: `interface{}`




- **值**: `strings.Join(...)`


**Variable**: `idx`

- **完整名称**: idx
- **文件位置**: 行 108-108
- **可见性**: private

- **类型**: `interface{}`




- **值**: `strings.IndexByte(...)`


**Variable**: `tagSetting`

- **完整名称**: tagSetting
- **文件位置**: 行 109-109
- **可见性**: private

- **类型**: `interface{}`




- **值**: `strings.Join(...)`


**Variable**: `settings`

- **完整名称**: settings
- **文件位置**: 行 110-110
- **可见性**: private

- **类型**: `interface{}`




- **值**: `ParseTagSetting(...)`


**Variable**: `length`

- **完整名称**: length
- **文件位置**: 行 111-111
- **可见性**: private

- **类型**: `interface{}`




- **值**: `strconv.Atoi(...)`


**Variable**: `_`

- **完整名称**: _
- **文件位置**: 行 111-111
- **可见性**: private

- **类型**: `interface{}`




- **值**: `strconv.Atoi(...)`


**Constant**: `key`

- **完整名称**: key
- **文件位置**: 行 121-121
- **可见性**: private

- **类型**: `interface{}`




- **值**: `"COMPOSITE"`





### schema\interfaces.go

- **语言**: go
- **包**: schema
- **代码行数**: 27
- **元素数量**: 7


#### 导入


- `gorm.io/gorm/clause`




#### 代码元素


**Interface**: `ConstraintInterface`

- **完整名称**: ConstraintInterface
- **文件位置**: 行 8-11
- **可见性**: public








- **子元素**: GetName Build 

**Interface**: `GormDataTypeInterface`

- **完整名称**: GormDataTypeInterface
- **文件位置**: 行 14-16
- **可见性**: public








- **子元素**: GormDataType 

**Interface**: `FieldNewValuePool`

- **完整名称**: FieldNewValuePool
- **文件位置**: 行 19-22
- **可见性**: public








- **子元素**: Get Put 

**Interface**: `CreateClausesInterface`

- **完整名称**: CreateClausesInterface
- **文件位置**: 行 25-27
- **可见性**: public








- **子元素**: CreateClauses 

**Interface**: `QueryClausesInterface`

- **完整名称**: QueryClausesInterface
- **文件位置**: 行 30-32
- **可见性**: public








- **子元素**: QueryClauses 

**Interface**: `UpdateClausesInterface`

- **完整名称**: UpdateClausesInterface
- **文件位置**: 行 35-37
- **可见性**: public








- **子元素**: UpdateClauses 

**Interface**: `DeleteClausesInterface`

- **完整名称**: DeleteClausesInterface
- **文件位置**: 行 40-42
- **可见性**: public








- **子元素**: DeleteClauses 




### schema\naming.go

- **语言**: go
- **包**: schema
- **代码行数**: 151
- **元素数量**: 23


#### 导入


- `crypto/sha1`

- `encoding/hex`

- `regexp`

- `strings`

- `unicode/utf8`

- `github.com/jinzhu/inflection`

- `golang.org/x/text/cases`

- `golang.org/x/text/language`




#### 代码元素


**Interface**: `Namer`

- **完整名称**: Namer
- **文件位置**: 行 16-25
- **可见性**: public








- **子元素**: TableName SchemaName ColumnName JoinTableName RelationshipFKName CheckerName IndexName UniqueName 

**Interface**: `Replacer`

- **完整名称**: Replacer
- **文件位置**: 行 28-30
- **可见性**: public








- **子元素**: Replace 

**Variable**: `_`

- **完整名称**: _
- **文件位置**: 行 32-32
- **可见性**: private

- **类型**: `Namer`




- **值**: `(...)(...)`


**Struct**: `NamingStrategy`

- **完整名称**: NamingStrategy
- **文件位置**: 行 35-41
- **可见性**: public








- **子元素**: TablePrefix SingularTable NameReplacer NoLowerCase IdentifierMaxLength 

**Method**: `TableName`

- **完整名称**: TableName
- **文件位置**: 行 44-49
- **可见性**: public
- **描述**: TableName convert string to table name

- **返回值**: `string`

- **参数**:
  - `str`: string






**Method**: `SchemaName`

- **完整名称**: SchemaName
- **文件位置**: 行 52-59
- **可见性**: public
- **描述**: SchemaName generate schema name from table name, don't guarantee it is the reverse value of TableName

- **返回值**: `string`

- **参数**:
  - `table`: string






**Method**: `ColumnName`

- **完整名称**: ColumnName
- **文件位置**: 行 62-64
- **可见性**: public
- **描述**: ColumnName convert string to column name

- **返回值**: `string`

- **参数**:
  - `table`: string
  - `column`: string






**Method**: `JoinTableName`

- **完整名称**: JoinTableName
- **文件位置**: 行 67-76
- **可见性**: public
- **描述**: JoinTableName convert string to join table name

- **返回值**: `string`

- **参数**:
  - `str`: string






**Method**: `RelationshipFKName`

- **完整名称**: RelationshipFKName
- **文件位置**: 行 79-81
- **可见性**: public
- **描述**: RelationshipFKName generate fk name for relation

- **返回值**: `string`

- **参数**:
  - `rel`: Relationship






**Method**: `CheckerName`

- **完整名称**: CheckerName
- **文件位置**: 行 84-86
- **可见性**: public
- **描述**: CheckerName generate checker name

- **返回值**: `string`

- **参数**:
  - `table`: string
  - `column`: string






**Method**: `IndexName`

- **完整名称**: IndexName
- **文件位置**: 行 89-91
- **可见性**: public
- **描述**: IndexName generate index name

- **返回值**: `string`

- **参数**:
  - `table`: string
  - `column`: string






**Method**: `UniqueName`

- **完整名称**: UniqueName
- **文件位置**: 行 94-96
- **可见性**: public
- **描述**: UniqueName generate unique constraint name

- **返回值**: `string`

- **参数**:
  - `table`: string
  - `column`: string






**Method**: `formatName`

- **完整名称**: formatName
- **文件位置**: 行 98-115
- **可见性**: private


- **返回值**: `string`

- **参数**:
  - `prefix`: string
  - `table`: string
  - `name`: string






**Variable**: `commonInitialisms`

- **完整名称**: commonInitialisms
- **文件位置**: 行 119-119
- **可见性**: private
- **描述**: https://github.com/golang/lint/blob/master/lint.go#L770
- **类型**: `interface{}`




- **值**: `...`


**Variable**: `commonInitialismsReplacer`

- **完整名称**: commonInitialismsReplacer
- **文件位置**: 行 120-120
- **可见性**: private

- **类型**: `*strings.Replacer`







**Method**: `toDBName`

- **完整名称**: toDBName
- **文件位置**: 行 131-188
- **可见性**: private


- **返回值**: `string`

- **参数**:
  - `name`: string






**Variable**: `value`

- **完整名称**: value
- **文件位置**: 行 151-151
- **可见性**: private

- **类型**: `interface{}`




- **值**: `commonInitialismsReplacer.Replace(...)`


**Variable**: `buf`

- **完整名称**: buf
- **文件位置**: 行 152-152
- **可见性**: private

- **类型**: `strings.Builder`







**Variable**: `lastCase`

- **完整名称**: lastCase
- **文件位置**: 行 153-153
- **可见性**: private

- **类型**: `bool`







**Variable**: `nextCase`

- **完整名称**: nextCase
- **文件位置**: 行 153-153
- **可见性**: private

- **类型**: `bool`







**Variable**: `nextNumber`

- **完整名称**: nextNumber
- **文件位置**: 行 153-153
- **可见性**: private

- **类型**: `bool`







**Variable**: `curCase`

- **完整名称**: curCase
- **文件位置**: 行 154-154
- **可见性**: private

- **类型**: `interface{}`




- **值**: `... <= 'Z' && ... >= 'A'`


**Method**: `toSchemaName`

- **完整名称**: toSchemaName
- **文件位置**: 行 190-196
- **可见性**: private


- **返回值**: `string`

- **参数**:
  - `name`: string









### schema\pool.go

- **语言**: go
- **包**: schema
- **代码行数**: 16
- **元素数量**: 2


#### 导入


- `reflect`

- `sync`




#### 代码元素


**Variable**: `normalPool`

- **完整名称**: normalPool
- **文件位置**: 行 10-10
- **可见性**: private

- **类型**: `sync.Map`







**Variable**: `poolInitializer`

- **完整名称**: poolInitializer
- **文件位置**: 行 11-11
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...`





### schema\relationship.go

- **语言**: go
- **包**: schema
- **代码行数**: 658
- **元素数量**: 46


#### 导入


- `context`

- `fmt`

- `reflect`

- `strings`

- `sync`

- `github.com/jinzhu/inflection`

- `golang.org/x/text/cases`

- `golang.org/x/text/language`

- `gorm.io/gorm/clause`




#### 代码元素


**Class**: `RelationshipType`

- **完整名称**: RelationshipType
- **文件位置**: 行 18-18
- **可见性**: public









**Constant**: `HasOne`

- **完整名称**: HasOne
- **文件位置**: 行 21-21
- **可见性**: public

- **类型**: `RelationshipType`




- **值**: `"has_one"`


**Constant**: `HasMany`

- **完整名称**: HasMany
- **文件位置**: 行 22-22
- **可见性**: public

- **类型**: `RelationshipType`




- **值**: `"has_many"`


**Constant**: `BelongsTo`

- **完整名称**: BelongsTo
- **文件位置**: 行 23-23
- **可见性**: public

- **类型**: `RelationshipType`




- **值**: `"belongs_to"`


**Constant**: `Many2Many`

- **完整名称**: Many2Many
- **文件位置**: 行 24-24
- **可见性**: public

- **类型**: `RelationshipType`




- **值**: `"many_to_many"`


**Constant**: `has`

- **完整名称**: has
- **文件位置**: 行 25-25
- **可见性**: private

- **类型**: `RelationshipType`




- **值**: `"has"`


**Struct**: `Relationships`

- **完整名称**: Relationships
- **文件位置**: 行 28-38
- **可见性**: public








- **子元素**: HasOne BelongsTo HasMany Many2Many Relations EmbeddedRelations Mux 

**Struct**: `Relationship`

- **完整名称**: Relationship
- **文件位置**: 行 40-50
- **可见性**: public








- **子元素**: Name Type Field Polymorphic References Schema FieldSchema JoinTable foreignKeys primaryKeys 

**Struct**: `Polymorphic`

- **完整名称**: Polymorphic
- **文件位置**: 行 52-56
- **可见性**: public








- **子元素**: PolymorphicID PolymorphicType Value 

**Struct**: `Reference`

- **完整名称**: Reference
- **文件位置**: 行 58-63
- **可见性**: public








- **子元素**: PrimaryKey PrimaryValue ForeignKey OwnPrimaryKey 

**Method**: `parseRelation`

- **完整名称**: parseRelation
- **文件位置**: 行 65-131
- **可见性**: private


- **返回值**: `*Relationship`

- **参数**:
  - `field`: *Field






**Variable**: `err`

- **完整名称**: err
- **文件位置**: 行 67-67
- **可见性**: private

- **类型**: `error`







**Variable**: `fieldValue`

- **完整名称**: fieldValue
- **文件位置**: 行 68-68
- **可见性**: private

- **类型**: `interface{}`




- **值**: `reflect.New(...).Interface(...)`


**Variable**: `relation`

- **完整名称**: relation
- **文件位置**: 行 69-69
- **可见性**: private

- **类型**: `interface{}`




- **值**: `&...`


**Method**: `setRelation`

- **完整名称**: setRelation
- **文件位置**: 行 147-181
- **可见性**: private




- **参数**:
  - `relation`: *Relationship






**Method**: `buildPolymorphicRelation`

- **完整名称**: buildPolymorphicRelation
- **文件位置**: 行 195-269
- **可见性**: private
- **描述**: User has many Toys, its `Polymorphic` is `Owner`, Pet has one Toy, its `Polymorphic` is `Owner`



- **参数**:
  - `relation`: *Relationship
  - `field`: *Field






**Variable**: `typeName`

- **完整名称**: typeName
- **文件位置**: 行 203-203
- **可见性**: private

- **类型**: `interface{}`




- **值**: `polymorphic + "Type"`


**Variable**: `typeId`

- **完整名称**: typeId
- **文件位置**: 行 204-204
- **可见性**: private

- **类型**: `interface{}`




- **值**: `polymorphic + "ID"`


**Method**: `buildMany2ManyRelation`

- **完整名称**: buildMany2ManyRelation
- **文件位置**: 行 271-444
- **可见性**: private




- **参数**:
  - `relation`: *Relationship
  - `field`: *Field
  - `many2many`: string






**Variable**: `err`

- **完整名称**: err
- **文件位置**: 行 275-275
- **可见性**: private

- **类型**: `error`







**Variable**: `joinTableFields`

- **完整名称**: joinTableFields
- **文件位置**: 行 276-276
- **可见性**: private

- **类型**: `[]reflect.StructField`







**Variable**: `fieldsMap`

- **完整名称**: fieldsMap
- **文件位置**: 行 277-277
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...`


**Variable**: `ownFieldsMap`

- **完整名称**: ownFieldsMap
- **文件位置**: 行 278-278
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...`


**Variable**: `referFieldsMap`

- **完整名称**: referFieldsMap
- **文件位置**: 行 279-279
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...`


**Variable**: `joinForeignKeys`

- **完整名称**: joinForeignKeys
- **文件位置**: 行 280-280
- **可见性**: private

- **类型**: `interface{}`




- **值**: `toColumns(...)`


**Variable**: `joinReferences`

- **完整名称**: joinReferences
- **文件位置**: 行 281-281
- **可见性**: private

- **类型**: `interface{}`




- **值**: `toColumns(...)`


**Class**: `guessLevel`

- **完整名称**: guessLevel
- **文件位置**: 行 446-446
- **可见性**: private









**Constant**: `guessGuess`

- **完整名称**: guessGuess
- **文件位置**: 行 449-449
- **可见性**: private

- **类型**: `guessLevel`




- **值**: `iota`


**Constant**: `guessBelongs`

- **完整名称**: guessBelongs
- **文件位置**: 行 450-450
- **可见性**: private









**Constant**: `guessEmbeddedBelongs`

- **完整名称**: guessEmbeddedBelongs
- **文件位置**: 行 451-451
- **可见性**: private









**Constant**: `guessHas`

- **完整名称**: guessHas
- **文件位置**: 行 452-452
- **可见性**: private









**Constant**: `guessEmbeddedHas`

- **完整名称**: guessEmbeddedHas
- **文件位置**: 行 453-453
- **可见性**: private









**Method**: `guessRelation`

- **完整名称**: guessRelation
- **文件位置**: 行 456-622
- **可见性**: private




- **参数**:
  - `relation`: *Relationship
  - `field`: *Field
  - `cgl`: guessLevel






**Variable**: `primaryFields`

- **完整名称**: primaryFields
- **文件位置**: 行 458-458
- **可见性**: private

- **类型**: `[]*Field`







**Variable**: `foreignFields`

- **完整名称**: foreignFields
- **文件位置**: 行 458-458
- **可见性**: private

- **类型**: `[]*Field`







**Variable**: `primarySchema`

- **完整名称**: primarySchema
- **文件位置**: 行 459-459
- **可见性**: private

- **类型**: `schema`




- **值**: `schema`


**Variable**: `foreignSchema`

- **完整名称**: foreignSchema
- **文件位置**: 行 459-459
- **可见性**: private

- **类型**: `schema`




- **值**: `schema`


**Variable**: `gl`

- **完整名称**: gl
- **文件位置**: 行 460-460
- **可见性**: private

- **类型**: `cgl`




- **值**: `cgl`


**Struct**: `Constraint`

- **完整名称**: Constraint
- **文件位置**: 行 625-634
- **可见性**: public








- **子元素**: Name Field Schema ForeignKeys ReferenceSchema References OnDelete OnUpdate 

**Method**: `GetName`

- **完整名称**: GetName
- **文件位置**: 行 636-636
- **可见性**: public


- **返回值**: `string`






**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 638-659
- **可见性**: public


- **返回值**: `string`






**Method**: `ParseConstraint`

- **完整名称**: ParseConstraint
- **文件位置**: 行 661-726
- **可见性**: public


- **返回值**: `*Constraint`






**Variable**: `name`

- **完整名称**: name
- **文件位置**: 行 688-688
- **可见性**: private

- **类型**: `string`







**Variable**: `idx`

- **完整名称**: idx
- **文件位置**: 行 689-689
- **可见性**: private

- **类型**: `interface{}`




- **值**: `strings.IndexByte(...)`


**Variable**: `settings`

- **完整名称**: settings
- **文件位置**: 行 690-690
- **可见性**: private

- **类型**: `interface{}`




- **值**: `ParseTagSetting(...)`


**Method**: `ToQueryConditions`

- **完整名称**: ToQueryConditions
- **文件位置**: 行 728-773
- **可见性**: public


- **返回值**: `[]clause.Expression`

- **参数**:
  - `ctx`: context.Context
  - `reflectValue`: reflect.Value









### schema\schema.go

- **语言**: go
- **包**: schema
- **代码行数**: 331
- **元素数量**: 23


#### 导入


- `context`

- `errors`

- `fmt`

- `go/ast`

- `path`

- `reflect`

- `strings`

- `sync`

- `gorm.io/gorm/clause`

- `gorm.io/gorm/logger`




#### 代码元素


**Class**: `callbackType`

- **完整名称**: callbackType
- **文件位置**: 行 17-17
- **可见性**: private









**Constant**: `callbackTypeBeforeCreate`

- **完整名称**: callbackTypeBeforeCreate
- **文件位置**: 行 20-20
- **可见性**: private

- **类型**: `callbackType`




- **值**: `"BeforeCreate"`


**Constant**: `callbackTypeBeforeUpdate`

- **完整名称**: callbackTypeBeforeUpdate
- **文件位置**: 行 21-21
- **可见性**: private

- **类型**: `callbackType`




- **值**: `"BeforeUpdate"`


**Constant**: `callbackTypeAfterCreate`

- **完整名称**: callbackTypeAfterCreate
- **文件位置**: 行 22-22
- **可见性**: private

- **类型**: `callbackType`




- **值**: `"AfterCreate"`


**Constant**: `callbackTypeAfterUpdate`

- **完整名称**: callbackTypeAfterUpdate
- **文件位置**: 行 23-23
- **可见性**: private

- **类型**: `callbackType`




- **值**: `"AfterUpdate"`


**Constant**: `callbackTypeBeforeSave`

- **完整名称**: callbackTypeBeforeSave
- **文件位置**: 行 24-24
- **可见性**: private

- **类型**: `callbackType`




- **值**: `"BeforeSave"`


**Constant**: `callbackTypeAfterSave`

- **完整名称**: callbackTypeAfterSave
- **文件位置**: 行 25-25
- **可见性**: private

- **类型**: `callbackType`




- **值**: `"AfterSave"`


**Constant**: `callbackTypeBeforeDelete`

- **完整名称**: callbackTypeBeforeDelete
- **文件位置**: 行 26-26
- **可见性**: private

- **类型**: `callbackType`




- **值**: `"BeforeDelete"`


**Constant**: `callbackTypeAfterDelete`

- **完整名称**: callbackTypeAfterDelete
- **文件位置**: 行 27-27
- **可见性**: private

- **类型**: `callbackType`




- **值**: `"AfterDelete"`


**Constant**: `callbackTypeAfterFind`

- **完整名称**: callbackTypeAfterFind
- **文件位置**: 行 28-28
- **可见性**: private

- **类型**: `callbackType`




- **值**: `"AfterFind"`


**Variable**: `ErrUnsupportedDataType`

- **完整名称**: ErrUnsupportedDataType
- **文件位置**: 行 32-32
- **可见性**: public

- **类型**: `interface{}`




- **值**: `errors.New(...)`


**Struct**: `Schema`

- **完整名称**: Schema
- **文件位置**: 行 34-61
- **可见性**: public








- **子元素**: Name ModelType Table PrioritizedPrimaryField DBNames PrimaryFields PrimaryFieldDBNames Fields FieldsByName FieldsByBindName FieldsByDBName FieldsWithDefaultDBValue Relationships CreateClauses QueryClauses UpdateClauses DeleteClauses BeforeCreate AfterCreate BeforeUpdate AfterUpdate BeforeDelete AfterDelete BeforeSave AfterSave AfterFind err initialized namer cacheStore 

**Method**: `String`

- **完整名称**: String
- **文件位置**: 行 63-68
- **可见性**: public


- **返回值**: `string`






**Method**: `MakeSlice`

- **完整名称**: MakeSlice
- **文件位置**: 行 70-76
- **可见性**: public


- **返回值**: `reflect.Value`






**Method**: `LookUpField`

- **完整名称**: LookUpField
- **文件位置**: 行 78-96
- **可见性**: public


- **返回值**: `*Field`

- **参数**:
  - `name`: string






**Method**: `LookUpFieldByBindName`

- **完整名称**: LookUpFieldByBindName
- **文件位置**: 行 106-114
- **可见性**: public
- **描述**: LookUpFieldByBindName looks for the closest field in the embedded struct.

- **返回值**: `*Field`

- **参数**:
  - `bindNames`: []string
  - `name`: string






**Interface**: `Tabler`

- **完整名称**: Tabler
- **文件位置**: 行 116-118
- **可见性**: public








- **子元素**: TableName 

**Interface**: `TablerWithNamer`

- **完整名称**: TablerWithNamer
- **文件位置**: 行 120-122
- **可见性**: public








- **子元素**: TableName 

**Variable**: `callbackTypes`

- **完整名称**: callbackTypes
- **文件位置**: 行 124-124
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...`


**Function**: `Parse`

- **完整名称**: Parse
- **文件位置**: 行 133-135
- **可见性**: public
- **描述**: Parse get data type from dialector

- **返回值**: `*Schema`

- **参数**:
  - `dest`: interface{}
  - `cacheStore`: *sync.Map
  - `namer`: Namer






**Function**: `ParseWithSpecialTableName`

- **完整名称**: ParseWithSpecialTableName
- **文件位置**: 行 138-390
- **可见性**: public
- **描述**: ParseWithSpecialTableName get data type from dialector with extra schema table

- **返回值**: `*Schema`

- **参数**:
  - `dest`: interface{}
  - `cacheStore`: *sync.Map
  - `namer`: Namer
  - `specialTableName`: string






**Variable**: `schemaCacheKey`

- **完整名称**: schemaCacheKey
- **文件位置**: 行 167-167
- **可见性**: private

- **类型**: `interface{}`




- **值**: `modelType`


**Variable**: `tableName`

- **完整名称**: tableName
- **文件位置**: 行 180-180
- **可见性**: private

- **类型**: `string`










### schema\serializer.go

- **语言**: go
- **包**: schema
- **代码行数**: 149
- **元素数量**: 19


#### 导入


- `bytes`

- `context`

- `database/sql`

- `database/sql/driver`

- `encoding/gob`

- `encoding/json`

- `fmt`

- `math`

- `reflect`

- `strings`

- `sync`

- `time`




#### 代码元素


**Variable**: `serializerMap`

- **完整名称**: serializerMap
- **文件位置**: 行 18-18
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...`


**Function**: `RegisterSerializer`

- **完整名称**: RegisterSerializer
- **文件位置**: 行 21-23
- **可见性**: public
- **描述**: RegisterSerializer register serializer



- **参数**:
  - `name`: string
  - `serializer`: SerializerInterface






**Function**: `GetSerializer`

- **完整名称**: GetSerializer
- **文件位置**: 行 26-32
- **可见性**: public
- **描述**: GetSerializer get serializer

- **返回值**: `SerializerInterface`

- **参数**:
  - `name`: string






**Struct**: `serializer`

- **完整名称**: serializer
- **文件位置**: 行 41-49
- **可见性**: private








- **子元素**: Field Serializer SerializeValuer Destination Context value fieldValue 

**Method**: `Scan`

- **完整名称**: Scan
- **文件位置**: 行 52-55
- **可见性**: public
- **描述**: Scan implements sql.Scanner interface

- **返回值**: `error`

- **参数**:
  - `value`: interface{}






**Method**: `Value`

- **完整名称**: Value
- **文件位置**: 行 58-60
- **可见性**: public
- **描述**: Value implements driver.Valuer interface

- **返回值**: `driver.Value`






**Interface**: `SerializerInterface`

- **完整名称**: SerializerInterface
- **文件位置**: 行 63-66
- **可见性**: public








- **子元素**: Scan 

**Interface**: `SerializerValuerInterface`

- **完整名称**: SerializerValuerInterface
- **文件位置**: 行 69-71
- **可见性**: public








- **子元素**: Value 

**Struct**: `JSONSerializer`

- **完整名称**: JSONSerializer
- **文件位置**: 行 74-74
- **可见性**: public









**Method**: `Scan`

- **完整名称**: Scan
- **文件位置**: 行 77-101
- **可见性**: public
- **描述**: Scan implements serializer interface

- **返回值**: `error`

- **参数**:
  - `ctx`: context.Context
  - `field`: *Field
  - `dst`: reflect.Value
  - `dbValue`: interface{}






**Variable**: `bytes`

- **完整名称**: bytes
- **文件位置**: 行 81-81
- **可见性**: private

- **类型**: `[]byte`







**Method**: `Value`

- **完整名称**: Value
- **文件位置**: 行 104-113
- **可见性**: public
- **描述**: Value implements serializer interface

- **返回值**: `interface{}`

- **参数**:
  - `ctx`: context.Context
  - `field`: *Field
  - `dst`: reflect.Value
  - `fieldValue`: interface{}






**Struct**: `UnixSecondSerializer`

- **完整名称**: UnixSecondSerializer
- **文件位置**: 行 116-116
- **可见性**: public









**Method**: `Scan`

- **完整名称**: Scan
- **文件位置**: 行 119-126
- **可见性**: public
- **描述**: Scan implements serializer interface

- **返回值**: `error`

- **参数**:
  - `ctx`: context.Context
  - `field`: *Field
  - `dst`: reflect.Value
  - `dbValue`: interface{}






**Method**: `Value`

- **完整名称**: Value
- **文件位置**: 行 129-158
- **可见性**: public
- **描述**: Value implements serializer interface

- **返回值**: `interface{}`

- **参数**:
  - `ctx`: context.Context
  - `field`: *Field
  - `dst`: reflect.Value
  - `fieldValue`: interface{}






**Struct**: `GobSerializer`

- **完整名称**: GobSerializer
- **文件位置**: 行 161-161
- **可见性**: public









**Method**: `Scan`

- **完整名称**: Scan
- **文件位置**: 行 164-182
- **可见性**: public
- **描述**: Scan implements serializer interface

- **返回值**: `error`

- **参数**:
  - `ctx`: context.Context
  - `field`: *Field
  - `dst`: reflect.Value
  - `dbValue`: interface{}






**Variable**: `bytesValue`

- **完整名称**: bytesValue
- **文件位置**: 行 168-168
- **可见性**: private

- **类型**: `[]byte`







**Method**: `Value`

- **完整名称**: Value
- **文件位置**: 行 185-189
- **可见性**: public
- **描述**: Value implements serializer interface

- **返回值**: `interface{}`

- **参数**:
  - `ctx`: context.Context
  - `field`: *Field
  - `dst`: reflect.Value
  - `fieldValue`: interface{}









### schema\utils.go

- **语言**: go
- **包**: schema
- **代码行数**: 184
- **元素数量**: 13


#### 导入


- `context`

- `fmt`

- `reflect`

- `regexp`

- `strings`

- `gorm.io/gorm/clause`

- `gorm.io/gorm/utils`




#### 代码元素


**Variable**: `embeddedCacheKey`

- **完整名称**: embeddedCacheKey
- **文件位置**: 行 14-14
- **可见性**: private

- **类型**: `interface{}`




- **值**: `"embedded_cache_store"`


**Function**: `ParseTagSetting`

- **完整名称**: ParseTagSetting
- **文件位置**: 行 16-43
- **可见性**: public


- **返回值**: `map[string]string`

- **参数**:
  - `str`: string
  - `sep`: string






**Variable**: `parsedNames`

- **完整名称**: parsedNames
- **文件位置**: 行 20-20
- **可见性**: private

- **类型**: `[]string`







**Function**: `GetRelationsValues`

- **完整名称**: GetRelationsValues
- **文件位置**: 行 70-105
- **可见性**: public
- **描述**: GetRelationsValues get relations's values from a reflect value

- **返回值**: `reflect.Value`

- **参数**:
  - `ctx`: context.Context
  - `reflectValue`: reflect.Value
  - `rels`: []*Relationship






**Function**: `GetIdentityFieldValuesMap`

- **完整名称**: GetIdentityFieldValuesMap
- **文件位置**: 行 108-179
- **可见性**: public
- **描述**: GetIdentityFieldValuesMap get identity map from fields

- **返回值**: `map[string][]reflect.Value`

- **参数**:
  - `ctx`: context.Context
  - `reflectValue`: reflect.Value
  - `fields`: []*Field






**Variable**: `results`

- **完整名称**: results
- **文件位置**: 行 110-110
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...`


**Variable**: `dataResults`

- **完整名称**: dataResults
- **文件位置**: 行 111-111
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...`


**Variable**: `loaded`

- **完整名称**: loaded
- **文件位置**: 行 112-112
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...`


**Variable**: `notZero`

- **完整名称**: notZero
- **文件位置**: 行 113-113
- **可见性**: private

- **类型**: `bool`







**Variable**: `zero`

- **完整名称**: zero
- **文件位置**: 行 113-113
- **可见性**: private

- **类型**: `bool`







**Function**: `GetIdentityFieldValuesMapFromValues`

- **完整名称**: GetIdentityFieldValuesMapFromValues
- **文件位置**: 行 182-194
- **可见性**: public
- **描述**: GetIdentityFieldValuesMapFromValues get identity map from fields

- **返回值**: `map[string][]reflect.Value`

- **参数**:
  - `ctx`: context.Context
  - `values`: []interface{}
  - `fields`: []*Field






**Function**: `ToQueryValues`

- **完整名称**: ToQueryValues
- **文件位置**: 行 197-217
- **可见性**: public
- **描述**: ToQueryValues to query values

- **返回值**: `interface{}`

- **参数**:
  - `table`: string
  - `foreignKeys`: []string
  - `foreignValues`: [][]interface{}






**Struct**: `embeddedNamer`

- **完整名称**: embeddedNamer
- **文件位置**: 行 219-222
- **可见性**: private








- **子元素**: Table 




### soft_delete.go

- **语言**: go
- **包**: gorm
- **代码行数**: 136
- **元素数量**: 23


#### 导入


- `database/sql`

- `database/sql/driver`

- `encoding/json`

- `reflect`

- `github.com/jinzhu/now`

- `gorm.io/gorm/clause`

- `gorm.io/gorm/schema`




#### 代码元素


**Class**: `DeletedAt`

- **完整名称**: DeletedAt
- **文件位置**: 行 14-14
- **可见性**: public









**Method**: `Scan`

- **完整名称**: Scan
- **文件位置**: 行 17-19
- **可见性**: public
- **描述**: Scan implements the Scanner interface.

- **返回值**: `error`

- **参数**:
  - `value`: interface{}






**Method**: `Value`

- **完整名称**: Value
- **文件位置**: 行 22-27
- **可见性**: public
- **描述**: Value implements the driver Valuer interface.

- **返回值**: `driver.Value`






**Method**: `MarshalJSON`

- **完整名称**: MarshalJSON
- **文件位置**: 行 29-34
- **可见性**: public


- **返回值**: `[]byte`






**Method**: `UnmarshalJSON`

- **完整名称**: UnmarshalJSON
- **文件位置**: 行 36-46
- **可见性**: public


- **返回值**: `error`

- **参数**:
  - `b`: []byte






**Method**: `QueryClauses`

- **完整名称**: QueryClauses
- **文件位置**: 行 48-50
- **可见性**: public


- **返回值**: `[]clause.Interface`

- **参数**:
  - `f`: *schema.Field






**Struct**: `SoftDeleteQueryClause`

- **完整名称**: SoftDeleteQueryClause
- **文件位置**: 行 61-64
- **可见性**: public








- **子元素**: ZeroValue Field 

**Method**: `Name`

- **完整名称**: Name
- **文件位置**: 行 66-68
- **可见性**: public


- **返回值**: `string`






**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 70-71
- **可见性**: public









**Method**: `MergeClause`

- **完整名称**: MergeClause
- **文件位置**: 行 73-74
- **可见性**: public









**Method**: `ModifyStatement`

- **完整名称**: ModifyStatement
- **文件位置**: 行 76-96
- **可见性**: public




- **参数**:
  - `stmt`: *Statement






**Method**: `UpdateClauses`

- **完整名称**: UpdateClauses
- **文件位置**: 行 98-100
- **可见性**: public


- **返回值**: `[]clause.Interface`

- **参数**:
  - `f`: *schema.Field






**Struct**: `SoftDeleteUpdateClause`

- **完整名称**: SoftDeleteUpdateClause
- **文件位置**: 行 102-105
- **可见性**: public








- **子元素**: ZeroValue Field 

**Method**: `Name`

- **完整名称**: Name
- **文件位置**: 行 107-109
- **可见性**: public


- **返回值**: `string`






**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 111-112
- **可见性**: public









**Method**: `MergeClause`

- **完整名称**: MergeClause
- **文件位置**: 行 114-115
- **可见性**: public









**Method**: `ModifyStatement`

- **完整名称**: ModifyStatement
- **文件位置**: 行 117-121
- **可见性**: public




- **参数**:
  - `stmt`: *Statement






**Method**: `DeleteClauses`

- **完整名称**: DeleteClauses
- **文件位置**: 行 123-125
- **可见性**: public


- **返回值**: `[]clause.Interface`

- **参数**:
  - `f`: *schema.Field






**Struct**: `SoftDeleteDeleteClause`

- **完整名称**: SoftDeleteDeleteClause
- **文件位置**: 行 127-130
- **可见性**: public








- **子元素**: ZeroValue Field 

**Method**: `Name`

- **完整名称**: Name
- **文件位置**: 行 132-134
- **可见性**: public


- **返回值**: `string`






**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 136-137
- **可见性**: public









**Method**: `MergeClause`

- **完整名称**: MergeClause
- **文件位置**: 行 139-140
- **可见性**: public









**Method**: `ModifyStatement`

- **完整名称**: ModifyStatement
- **文件位置**: 行 142-170
- **可见性**: public




- **参数**:
  - `stmt`: *Statement









### statement.go

- **语言**: go
- **包**: gorm
- **代码行数**: 674
- **元素数量**: 24


#### 导入


- `context`

- `database/sql`

- `database/sql/driver`

- `fmt`

- `reflect`

- `regexp`

- `sort`

- `strconv`

- `strings`

- `sync`

- `gorm.io/gorm/clause`

- `gorm.io/gorm/logger`

- `gorm.io/gorm/schema`

- `gorm.io/gorm/utils`




#### 代码元素


**Struct**: `Statement`

- **完整名称**: Statement
- **文件位置**: 行 22-51
- **可见性**: public








- **子元素**: TableExpr Table Model Unscoped Dest ReflectValue Clauses BuildClauses Distinct Selects Omits ColumnMapping Joins Preloads Settings ConnPool Schema Context RaiseErrorOnNotFound SkipHooks SQL Vars CurDestIndex attrs assigns scopes Result 

**Struct**: `join`

- **完整名称**: join
- **文件位置**: 行 53-62
- **可见性**: private








- **子元素**: Name Alias Conds On Selects Omits Expression JoinType 

**Interface**: `StatementModifier`

- **完整名称**: StatementModifier
- **文件位置**: 行 65-67
- **可见性**: public








- **子元素**: ModifyStatement 

**Method**: `WriteString`

- **完整名称**: WriteString
- **文件位置**: 行 70-72
- **可见性**: public
- **描述**: WriteString write string

- **返回值**: `int`

- **参数**:
  - `str`: string






**Method**: `WriteByte`

- **完整名称**: WriteByte
- **文件位置**: 行 75-77
- **可见性**: public
- **描述**: WriteByte write byte

- **返回值**: `error`

- **参数**:
  - `c`: byte






**Method**: `WriteQuoted`

- **完整名称**: WriteQuoted
- **文件位置**: 行 80-82
- **可见性**: public
- **描述**: WriteQuoted write quoted value



- **参数**:
  - `value`: interface{}






**Method**: `QuoteTo`

- **完整名称**: QuoteTo
- **文件位置**: 行 85-165
- **可见性**: public
- **描述**: QuoteTo write quoted value to writer



- **参数**:
  - `writer`: clause.Writer
  - `field`: interface{}






**Method**: `Quote`

- **完整名称**: Quote
- **文件位置**: 行 168-172
- **可见性**: public
- **描述**: Quote returns quoted value

- **返回值**: `string`

- **参数**:
  - `field`: interface{}






**Variable**: `builder`

- **完整名称**: builder
- **文件位置**: 行 169-169
- **可见性**: private

- **类型**: `strings.Builder`







**Method**: `AddVar`

- **完整名称**: AddVar
- **文件位置**: 行 175-269
- **可见性**: public
- **描述**: AddVar add var



- **参数**:
  - `writer`: clause.Writer
  - `vars`: ...interface{}






**Variable**: `vars`

- **完整名称**: vars
- **文件位置**: 行 219-219
- **可见性**: private

- **类型**: `subdb.Statement.Vars`




- **值**: `subdb.Statement.Vars`


**Variable**: `sql`

- **完整名称**: sql
- **文件位置**: 行 220-220
- **可见性**: private

- **类型**: `interface{}`




- **值**: `cv.Statement.SQL.String(...)`


**Method**: `AddClause`

- **完整名称**: AddClause
- **文件位置**: 行 272-282
- **可见性**: public
- **描述**: AddClause add clause



- **参数**:
  - `v`: clause.Interface






**Method**: `AddClauseIfNotExists`

- **完整名称**: AddClauseIfNotExists
- **文件位置**: 行 285-289
- **可见性**: public
- **描述**: AddClauseIfNotExists add clause if not exists



- **参数**:
  - `v`: clause.Interface






**Method**: `BuildCondition`

- **完整名称**: BuildCondition
- **文件位置**: 行 292-490
- **可见性**: public
- **描述**: BuildCondition build condition

- **返回值**: `[]clause.Expression`

- **参数**:
  - `query`: interface{}
  - `args`: ...interface{}






**Method**: `Build`

- **完整名称**: Build
- **文件位置**: 行 493-510
- **可见性**: public
- **描述**: Build build sql with clauses names



- **参数**:
  - `clauses`: ...string






**Variable**: `firstClauseWritten`

- **完整名称**: firstClauseWritten
- **文件位置**: 行 494-494
- **可见性**: private

- **类型**: `bool`







**Method**: `Parse`

- **完整名称**: Parse
- **文件位置**: 行 512-514
- **可见性**: public


- **返回值**: `error`

- **参数**:
  - `value`: interface{}






**Method**: `ParseWithSpecialTableName`

- **完整名称**: ParseWithSpecialTableName
- **文件位置**: 行 516-527
- **可见性**: public


- **返回值**: `error`

- **参数**:
  - `value`: interface{}
  - `specialTableName`: string






**Method**: `clone`

- **完整名称**: clone
- **文件位置**: 行 529-581
- **可见性**: private


- **返回值**: `*Statement`






**Method**: `SetColumn`

- **完整名称**: SetColumn
- **文件位置**: 行 587-640
- **可见性**: public
- **描述**: SetColumn set column's value



- **参数**:
  - `name`: string
  - `value`: interface{}
  - `fromCallbacks`: ...bool






**Method**: `Changed`

- **完整名称**: Changed
- **文件位置**: 行 643-696
- **可见性**: public
- **描述**: Changed check model changed or not when updating

- **返回值**: `bool`

- **参数**:
  - `fields`: ...string






**Variable**: `matchName`

- **完整名称**: matchName
- **文件位置**: 行 698-698
- **可见性**: private

- **类型**: `interface{}`




- **值**: `...(...)`


**Method**: `SelectAndOmitColumns`

- **完整名称**: SelectAndOmitColumns
- **文件位置**: 行 715-772
- **可见性**: public
- **描述**: SelectAndOmitColumns get select and omit columns, select -> true, omit -> false

- **返回值**: `map[string]bool`

- **参数**:
  - `requireCreate`: bool
  - `requireUpdate`: bool









### utils\utils.go

- **语言**: go
- **包**: utils
- **代码行数**: 154
- **元素数量**: 14


#### 导入


- `database/sql/driver`

- `fmt`

- `path/filepath`

- `reflect`

- `runtime`

- `strconv`

- `strings`

- `unicode`




#### 代码元素


**Variable**: `gormSourceDir`

- **完整名称**: gormSourceDir
- **文件位置**: 行 14-14
- **可见性**: private

- **类型**: `string`







**Function**: `CallerFrame`

- **完整名称**: CallerFrame
- **文件位置**: 行 38-53
- **可见性**: public
- **描述**: CallerFrame retrieves the first relevant stack frame outside of GORM's internal implementation files.

- **返回值**: `runtime.Frame`






**Function**: `FileWithLineNum`

- **完整名称**: FileWithLineNum
- **文件位置**: 行 56-63
- **可见性**: public
- **描述**: FileWithLineNum return the file name and line number of the current file

- **返回值**: `string`






**Function**: `IsInvalidDBNameChar`

- **完整名称**: IsInvalidDBNameChar
- **文件位置**: 行 65-67
- **可见性**: public


- **返回值**: `bool`

- **参数**:
  - `c`: rune






**Function**: `CheckTruth`

- **完整名称**: CheckTruth
- **文件位置**: 行 70-77
- **可见性**: public
- **描述**: CheckTruth check string true or not

- **返回值**: `bool`

- **参数**:
  - `vals`: ...string






**Function**: `ToStringKey`

- **完整名称**: ToStringKey
- **文件位置**: 行 79-104
- **可见性**: public


- **返回值**: `string`

- **参数**:
  - `values`: ...interface{}






**Function**: `Contains`

- **完整名称**: Contains
- **文件位置**: 行 106-113
- **可见性**: public


- **返回值**: `bool`

- **参数**:
  - `elems`: []string
  - `elem`: string






**Function**: `AssertEqual`

- **完整名称**: AssertEqual
- **文件位置**: 行 115-137
- **可见性**: public


- **返回值**: `bool`

- **参数**:
  - `x`: interface{}
  - `y`: interface{}






**Function**: `ToString`

- **完整名称**: ToString
- **文件位置**: 行 139-165
- **可见性**: public


- **返回值**: `string`

- **参数**:
  - `value`: interface{}






**Constant**: `nestedRelationSplit`

- **完整名称**: nestedRelationSplit
- **文件位置**: 行 167-167
- **可见性**: private

- **类型**: `interface{}`




- **值**: `"__"`


**Function**: `NestedRelationName`

- **完整名称**: NestedRelationName
- **文件位置**: 行 170-172
- **可见性**: public
- **描述**: NestedRelationName nested relationships like `Manager__Company`

- **返回值**: `string`

- **参数**:
  - `prefix`: string
  - `name`: string






**Function**: `SplitNestedRelationName`

- **完整名称**: SplitNestedRelationName
- **文件位置**: 行 175-177
- **可见性**: public
- **描述**: SplitNestedRelationName Split nested relationships to `[]string{"Manager","Company"}`

- **返回值**: `[]string`

- **参数**:
  - `name`: string






**Function**: `JoinNestedRelationNames`

- **完整名称**: JoinNestedRelationNames
- **文件位置**: 行 180-182
- **可见性**: public
- **描述**: JoinNestedRelationNames nested relationships like `Manager__Company`

- **返回值**: `string`

- **参数**:
  - `relationNames`: []string






**Function**: `RTrimSlice`

- **完整名称**: RTrimSlice
- **文件位置**: 行 185-193
- **可见性**: public
- **描述**: RTrimSlice Right trims the given slice by given length

- **返回值**: `[]T`

- **参数**:
  - `v`: []T
  - `trimLen`: int










## API 参考

### 公共接口



#### Struct: Association








---




#### Method: Association


```go
func (*DB) Association(column string) (*Association)
```





**参数**:

- `column`: string



**返回**: `*Association`


**调用**: fmt.Errorf() reflect.ValueOf() 


---




#### Method: Unscoped


```go
func (*Association) Unscoped() (*Association)
```






**返回**: `*Association`



---




#### Method: Find


```go
func (*Association) Find(out interface{}, conds ...interface{}) (error)
```





**参数**:

- `out`: interface{}

- `conds`: ...interface{}



**返回**: `error`


**调用**: association.buildCondition() 


---




#### Method: Append


```go
func (*Association) Append(values ...interface{}) (error)
```





**参数**:

- `values`: ...interface{}



**返回**: `error`


**调用**: expandValues() len() association.Replace() association.saveAssociation() 


---




#### Method: Replace


```go
func (*Association) Replace(values ...interface{}) (error)
```





**参数**:

- `values`: ...interface{}



**返回**: `error`


**调用**: expandValues() append() schema.GetIdentityFieldValuesMap() len() schema.ToQueryValues() association.saveAssociation() len() reflectValue.Kind() reflectValue.Len() reflectValue.Index() reflect.Zero() reflect.Zero() reflect.New() schema.GetRelationsValues() reflect.New() schema.GetIdentityFieldValuesMap() len() schema.ToQueryValues() len() tx.Not() append() append() tx.Where() schema.GetIdentityFieldValuesMap() len() schema.ToQueryValues() tx.Where() tx.Where() reflect.New() append() append() append() append() tx.Clauses() schema.GetIdentityFieldValuesMap() schema.ToQueryValues() len() tx.Where() schema.GetIdentityFieldValuesMapFromValues() schema.ToQueryValues() len() tx.Where() clause.Not() tx.Delete() 


---
































#### Method: Delete


```go
func (*Association) Delete(values ...interface{}) (error)
```





**参数**:

- `values`: ...interface{}



**返回**: `error`


**调用**: expandValues() append() append() append() associationDB.Model() reflect.New() schema.GetIdentityFieldValuesMap() schema.ToQueryValues() len() append() schema.GetIdentityFieldValuesMapFromValues() schema.ToQueryValues() append() tx.Clauses() append() schema.GetIdentityFieldValuesMap() len() schema.ToQueryValues() associationDB.Model() reflect.New() reflect.New() schema.GetIdentityFieldValuesMap() schema.ToQueryValues() len() append() schema.GetIdentityFieldValuesMapFromValues() schema.ToQueryValues() append() tx.Clauses() tx.Clauses() reflect.New() append() append() append() append() append() schema.GetIdentityFieldValuesMap() schema.ToQueryValues() len() append() schema.GetIdentityFieldValuesMapFromValues() schema.ToQueryValues() append() schema.GetIdentityFieldValuesMapFromValues() reflect.Indirect() make() len() fieldValue.Kind() reflect.Zero() fieldValue.Len() field.ValueOf() fieldValue.Index() utils.ToStringKey() reflect.Append() fieldValue.Index() validFieldValues.Interface() field.ValueOf() utils.ToStringKey() reflect.Zero() reflect.Zero() reflect.Zero() reflectValue.Kind() reflectValue.Len() cleanUpDeletedRelations() reflect.Indirect() reflectValue.Index() cleanUpDeletedRelations() 


---




























#### Method: Clear


```go
func (*Association) Clear() (error)
```






**返回**: `error`


**调用**: association.Replace() 


---




#### Method: Count


```go
func (*Association) Count() (int64)
```






**返回**: `int64`


**调用**: association.buildCondition() 


---






















#### Function: SaveBeforeAssociations


```go
func SaveBeforeAssociations(create bool) (func(...))
```





**参数**:

- `create`: bool



**返回**: `func(...)`


**调用**: db.AddError() elem.Interface() make() fieldType.Kind() reflect.PointerTo() reflect.MakeSlice() reflect.SliceOf() reflect.MakeSlice() reflect.SliceOf() reflect.Indirect() rv.Addr() append() reflect.Append() make() len() pf.ValueOf() append() utils.ToStringKey() len() len() reflect.Append() elems.Len() saveAssociations() elems.Len() setupReferences() elems.Index() rv.Kind() rv.Addr() saveAssociations() setupReferences() 


---












#### Function: SaveAfterAssociations


```go
func SaveAfterAssociations(create bool) (func(...))
```





**参数**:

- `create`: bool



**返回**: `func(...)`


**调用**: fieldType.Kind() reflect.PointerTo() reflect.MakeSlice() reflect.SliceOf() reflect.Indirect() rv.Kind() rv.Addr() db.AddError() db.AddError() reflect.Append() elems.Len() make() len() append() saveAssociations() f.Kind() f.Addr() make() len() db.AddError() db.AddError() append() saveAssociations() fieldType.Kind() reflect.PointerTo() reflect.MakeSlice() reflect.SliceOf() reflect.Indirect() f.Len() f.Index() db.AddError() db.AddError() make() len() pf.ValueOf() append() utils.ToStringKey() len() len() reflect.Append() reflect.Append() elem.Addr() reflect.Indirect() appendToElems() appendToElems() elems.Len() make() len() append() saveAssociations() fieldType.Kind() reflect.PointerTo() reflect.MakeSlice() reflect.SliceOf() reflect.MakeSlice() reflect.SliceOf() reflect.MakeSlice() reflect.SliceOf() reflect.PointerTo() reflect.New() db.AddError() db.AddError() db.AddError() reflect.Append() reflect.Indirect() f.Len() f.Index() elem.Addr() append() reflect.Append() make() len() pf.ValueOf() append() utils.ToStringKey() len() len() reflect.Append() reflect.Indirect() appendToElems() appendToElems() elems.Len() saveAssociations() appendToJoins() elems.Index() joins.Len() db.AddError() db.Session() joins.Interface() 


---




























#### Struct: Config








---




#### Function: RegisterDefaultCallbacks


```go
func RegisterDefaultCallbacks(db *gorm.DB, config *Config)
```





**参数**:

- `db`: *gorm.DB

- `config`: *Config




**调用**: len() len() len() len() db.Callback() createCallback.Match() createCallback.Register() createCallback.Register() SaveBeforeAssociations() createCallback.Register() Create() createCallback.Register() SaveAfterAssociations() createCallback.Register() createCallback.Match() db.Callback() queryCallback.Register() queryCallback.Register() queryCallback.Register() db.Callback() deleteCallback.Match() deleteCallback.Register() deleteCallback.Register() deleteCallback.Register() Delete() deleteCallback.Register() deleteCallback.Match() db.Callback() updateCallback.Match() updateCallback.Register() updateCallback.Register() updateCallback.Register() SaveBeforeAssociations() updateCallback.Register() Update() updateCallback.Register() SaveAfterAssociations() updateCallback.Register() updateCallback.Match() db.Callback() rowCallback.Register() db.Callback() rawCallback.Register() 


---




#### Function: BeforeCreate


```go
func BeforeCreate(db *gorm.DB)
```


BeforeCreate before create hooks



```
BeforeCreate before create hooks

```



**参数**:

- `db`: *gorm.DB




**调用**: callMethod() db.AddError() i.BeforeSave() db.AddError() i.BeforeCreate() 


---




#### Function: Create


```go
func Create(config *Config) (func(...))
```


Create create hook



```
Create create hook

```



**参数**:

- `config`: *Config



**返回**: `func(...)`


**调用**: utils.Contains() len() make() len() append() len() ConvertToCreateValues() hasReturning() len() db.AddError() db.AddError() rows.Close() gorm.Scan() db.AddError() result.RowsAffected() result.LastInsertId() db.AddError() int64() len() reflect.Indirect() pkField.ValueOf() db.AddError() pkField.Set() reflect.Indirect() pkField.ValueOf() db.AddError() pkField.Set() pkField.ValueOf() db.AddError() pkField.Set() 


---








#### Function: AfterCreate


```go
func AfterCreate(db *gorm.DB)
```


AfterCreate after create hooks



```
AfterCreate after create hooks

```



**参数**:

- `db`: *gorm.DB




**调用**: callMethod() db.AddError() i.AfterCreate() db.AddError() i.AfterSave() 


---




#### Function: ConvertToCreateValues


```go
func ConvertToCreateValues(stmt *gorm.Statement) (clause.Values)
```


ConvertToCreateValues convert to create values



```
ConvertToCreateValues convert to create values

```



**参数**:

- `stmt`: *gorm.Statement



**返回**: `clause.Values`


**调用**: ConvertMapToValuesForCreate() ConvertMapToValuesForCreate() ConvertSliceOfMapToValuesForCreate() ConvertSliceOfMapToValuesForCreate() stmt.SelectAndOmitColumns() stmt.Get() make() len() append() stmt.AddError() make() len() make() reflect.Indirect() rv.IsValid() stmt.AddError() fmt.Errorf() make() len() field.ValueOf() stmt.AddError() field.Set() stmt.AddError() field.Set() field.ValueOf() stmt.AddError() field.Set() field.ValueOf() field.ValueOf() len() make() append() append() stmt.DefaultValueOf() append() make() len() field.ValueOf() stmt.AddError() field.Set() stmt.AddError() field.Set() field.ValueOf() stmt.AddError() field.Set() field.ValueOf() field.ValueOf() append() append() stmt.AddError() len() stmt.SelectAndOmitColumns() make() len() strings.EqualFold() curTime.UnixNano() curTime.UnixMilli() curTime.Unix() append() append() append() clause.AssignmentColumns() len() len() append() stmt.AddClause() 


---














#### Function: BeforeDelete


```go
func BeforeDelete(db *gorm.DB)
```





**参数**:

- `db`: *gorm.DB




**调用**: callMethod() db.AddError() i.BeforeDelete() 


---




#### Function: DeleteBeforeAssociations


```go
func DeleteBeforeAssociations(db *gorm.DB)
```





**参数**:

- `db`: *gorm.DB




**调用**: rel.ToQueryConditions() reflect.New() db.Session() tx.Unscoped() len() make() len() append() strings.HasPrefix() append() strings.TrimPrefix() len() tx.Select() len() db.AddError() tx.Clauses() make() len() make() len() make() len() reflect.New() db.Session() append() append() append() schema.GetIdentityFieldValuesMap() schema.ToQueryValues() append() db.AddError() tx.Clauses() 


---
















#### Function: Delete


```go
func Delete(config *Config) (func(...))
```





**参数**:

- `config`: *Config



**返回**: `func(...)`


**调用**: utils.Contains() schema.GetIdentityFieldValuesMap() schema.ToQueryValues() len() schema.GetIdentityFieldValuesMap() reflect.ValueOf() schema.ToQueryValues() len() checkMissingWhereConditions() hasReturning() db.AddError() result.RowsAffected() db.AddError() gorm.Scan() db.AddError() rows.Close() 


---




#### Function: AfterDelete


```go
func AfterDelete(db *gorm.DB)
```





**参数**:

- `db`: *gorm.DB




**调用**: callMethod() db.AddError() i.AfterDelete() 


---




#### Function: ConvertMapToValuesForCreate


```go
func ConvertMapToValuesForCreate(stmt *gorm.Statement, mapValue map[string]interface{}) (clause.Values)
```


ConvertMapToValuesForCreate convert map to values



```
ConvertMapToValuesForCreate convert map to values

```



**参数**:

- `stmt`: *gorm.Statement

- `mapValue`: map[string]interface{}



**返回**: `clause.Values`


**调用**: make() len() stmt.SelectAndOmitColumns() make() len() append() sort.Strings() append() len() append() 


---




#### Function: ConvertSliceOfMapToValuesForCreate


```go
func ConvertSliceOfMapToValuesForCreate(stmt *gorm.Statement, mapValues []map[string]interface{}) (clause.Values)
```


ConvertSliceOfMapToValuesForCreate convert slice of map to values



```
ConvertSliceOfMapToValuesForCreate convert slice of map to values

```



**参数**:

- `stmt`: *gorm.Statement

- `mapValues`: []map[string]interface{}



**返回**: `clause.Values`


**调用**: make() len() len() stmt.AddError() make() len() stmt.SelectAndOmitColumns() make() len() append() sort.Strings() make() len() make() len() len() make() len() 


---












#### Interface: BeforeCreateInterface








---




#### Interface: AfterCreateInterface








---




#### Interface: BeforeUpdateInterface








---




#### Interface: AfterUpdateInterface








---




#### Interface: BeforeSaveInterface








---




#### Interface: AfterSaveInterface








---




#### Interface: BeforeDeleteInterface








---




#### Interface: AfterDeleteInterface








---




#### Interface: AfterFindInterface








---
























#### Function: Query


```go
func Query(db *gorm.DB)
```





**参数**:

- `db`: *gorm.DB




**调用**: BuildQuerySQL() db.AddError() db.AddError() rows.Close() gorm.Scan() 


---




#### Function: BuildQuerySQL


```go
func BuildQuerySQL(db *gorm.DB)
```





**参数**:

- `db`: *gorm.DB




**调用**: primaryField.ValueOf() append() len() len() make() len() len() make() len() append() stmt.Parse() make() len() len() len() len() len() make() len() append() strings.Split() len() make() len() append() columnStmt.SelectAndOmitColumns() append() utils.NestedRelationName() make() len() onStmt.AddClause() onStmt.AddClause() where.Build() strings.Replace() bindvar.String() append() utils.NestedRelationName() len() append() genJoinClause() append() append() 


---










#### Function: Preload


```go
func Preload(db *gorm.DB)
```





**参数**:

- `db`: *gorm.DB




**调用**: len() db.AddError() fmt.Errorf() make() len() append() preloadDB() db.AddError() preloadEntryPoint() 


---




#### Function: AfterQuery


```go
func AfterQuery(db *gorm.DB)
```





**参数**:

- `db`: *gorm.DB




**调用**: utils.RTrimSlice() len() callMethod() db.AddError() i.AfterFind() 


---




#### Function: RawExec


```go
func RawExec(db *gorm.DB)
```





**参数**:

- `db`: *gorm.DB




**调用**: db.AddError() result.RowsAffected() 


---




#### Function: RowQuery


```go
func RowQuery(db *gorm.DB)
```





**参数**:

- `db`: *gorm.DB




**调用**: BuildQuerySQL() db.Get() 


---




#### Function: BeginTransaction


```go
func BeginTransaction(db *gorm.DB)
```





**参数**:

- `db`: *gorm.DB




**调用**: db.Begin() db.InstanceSet() 


---




#### Function: CommitOrRollbackTransaction


```go
func CommitOrRollbackTransaction(db *gorm.DB)
```





**参数**:

- `db`: *gorm.DB




**调用**: db.InstanceGet() db.Rollback() db.Commit() 


---




#### Function: SetupUpdateReflectValue


```go
func SetupUpdateReflectValue(db *gorm.DB)
```





**参数**:

- `db`: *gorm.DB




**调用**: reflect.ValueOf() db.AddError() 


---




#### Function: BeforeUpdate


```go
func BeforeUpdate(db *gorm.DB)
```


BeforeUpdate before update hooks



```
BeforeUpdate before update hooks

```



**参数**:

- `db`: *gorm.DB




**调用**: callMethod() db.AddError() i.BeforeSave() db.AddError() i.BeforeUpdate() 


---




#### Function: Update


```go
func Update(config *Config) (func(...))
```


Update update hook



```
Update update hook

```



**参数**:

- `config`: *Config



**返回**: `func(...)`


**调用**: utils.Contains() ConvertToAssignments() len() delete() checkMissingWhereConditions() hasReturning() db.AddError() gorm.Scan() db.AddError() rows.Close() db.AddError() result.RowsAffected() 


---




#### Function: AfterUpdate


```go
func AfterUpdate(db *gorm.DB)
```


AfterUpdate after update hooks



```
AfterUpdate after update hooks

```



**参数**:

- `db`: *gorm.DB




**调用**: callMethod() db.AddError() i.AfterUpdate() db.AddError() i.AfterSave() 


---




#### Function: ConvertToAssignments


```go
func ConvertToAssignments(stmt *gorm.Statement) (clause.Set)
```


ConvertToAssignments convert to update assignments



```
ConvertToAssignments convert to update assignments

```



**参数**:

- `stmt`: *gorm.Statement



**返回**: `clause.Set`


**调用**: stmt.SelectAndOmitColumns() field.Set() field.Set() reflect.ValueOf() updatingValue.Kind() updatingValue.Elem() updatingValue.CanAddr() field.ValueOf() schema.GetIdentityFieldValuesMap() schema.ToQueryValues() stmt.AddClause() field.ValueOf() stmt.AddClause() updatingValue.Interface() make() len() make() len() append() sort.Strings() append() assignValue() assignValue() append() assignValue() append() now.UnixNano() append() now.UnixMilli() append() now.Unix() append() updatingValue.CanAddr() updatingStmt.Parse() updatingValue.Kind() make() len() updatingSchema.LookUpField() updatingValue.CanAddr() field.ValueOf() append() assignValue() field.ValueOf() stmt.AddClause() stmt.AddError() 


---




















#### Method: Create


```go
func (*callbacks) Create() (*processor)
```






**返回**: `*processor`



---




#### Method: Query


```go
func (*callbacks) Query() (*processor)
```






**返回**: `*processor`



---




#### Method: Update


```go
func (*callbacks) Update() (*processor)
```






**返回**: `*processor`



---




#### Method: Delete


```go
func (*callbacks) Delete() (*processor)
```






**返回**: `*processor`



---




#### Method: Row


```go
func (*callbacks) Row() (*processor)
```






**返回**: `*processor`



---




#### Method: Raw


```go
func (*callbacks) Raw() (*processor)
```






**返回**: `*processor`



---




#### Method: Execute


```go
func (*processor) Execute(db *DB) (*DB)
```





**参数**:

- `db`: *DB



**返回**: `*DB`


**调用**: len() db.executeScopes() time.Now() len() optimizer.ModifyStatement() context.WithTimeout() stmt.Parse() errors.Is() errors.Is() db.AddError() fmt.Errorf() db.AddError() reflect.ValueOf() reflect.New() db.AddError() f() filter.ParamsFilter() 


---










#### Method: Get


```go
func (*processor) Get(name string) (func(...))
```





**参数**:

- `name`: string



**返回**: `func(...)`


**调用**: len() 


---




#### Method: Before


```go
func (*processor) Before(name string) (*callback)
```





**参数**:

- `name`: string



**返回**: `*callback`



---




#### Method: After


```go
func (*processor) After(name string) (*callback)
```





**参数**:

- `name`: string



**返回**: `*callback`



---




#### Method: Match


```go
func (*processor) Match(fc func(...)) (*callback)
```





**参数**:

- `fc`: func(...)



**返回**: `*callback`



---




#### Method: Register


```go
func (*processor) Register(name string, fn func(...)) (error)
```





**参数**:

- `name`: string

- `fn`: func(...)



**返回**: `error`



---




#### Method: Remove


```go
func (*processor) Remove(name string) (error)
```





**参数**:

- `name`: string



**返回**: `error`



---




#### Method: Replace


```go
func (*processor) Replace(name string, fn func(...)) (error)
```





**参数**:

- `name`: string

- `fn`: func(...)



**返回**: `error`



---








#### Method: Before


```go
func (*callback) Before(name string) (*callback)
```





**参数**:

- `name`: string



**返回**: `*callback`



---




#### Method: After


```go
func (*callback) After(name string) (*callback)
```





**参数**:

- `name`: string



**返回**: `*callback`



---




#### Method: Register


```go
func (*callback) Register(name string, fn func(...)) (error)
```





**参数**:

- `name`: string

- `fn`: func(...)



**返回**: `error`


**调用**: append() 


---




#### Method: Remove


```go
func (*callback) Remove(name string) (error)
```





**参数**:

- `name`: string



**返回**: `error`


**调用**: context.Background() utils.FileWithLineNum() append() 


---




#### Method: Replace


```go
func (*callback) Replace(name string, fn func(...)) (error)
```





**参数**:

- `name`: string

- `fn`: func(...)



**返回**: `error`


**调用**: context.Background() utils.FileWithLineNum() append() 


---










#### Method: Model


```go
func (*DB) Model(value interface{}) (*DB)
```


Model specify the model you would like to run db operations



```
Model specify the model you would like to run db operations

	// update all users's name to `hello`
	db.Model(&User{}).Update("name", "hello")
	// if user's primary key is non-blank, will use it as condition, then will only update that user's name to `hello`
	db.Model(&user).Update("name", "hello")

```



**参数**:

- `value`: interface{}



**返回**: `*DB`


**调用**: db.getInstance() 


---




#### Method: Clauses


```go
func (*DB) Clauses(conds ...clause.Expression) (*DB)
```


Clauses Add clauses



```
Clauses Add clauses

This supports both standard clauses (clause.OrderBy, clause.Limit, clause.Where) and more
advanced techniques like specifying lock strength and optimizer hints. See the
[docs] for more depth.

	// add a simple limit clause
	db.Clauses(clause.Limit{Limit: 1}).Find(&User{})
	// tell the optimizer to use the `idx_user_name` index
	db.Clauses(hints.UseIndex("idx_user_name")).Find(&User{})
	// specify the lock strength to UPDATE
	db.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&users)

[docs]: https://gorm.io/docs/sql_builder.html#Clauses

```



**参数**:

- `conds`: ...clause.Expression



**返回**: `*DB`


**调用**: db.getInstance() optimizer.ModifyStatement() append() len() 


---








#### Method: Table


```go
func (*DB) Table(name string, args ...interface{}) (*DB)
```


Table specify the table you would like to run db operations



```
Table specify the table you would like to run db operations

	// Get a user
	db.Table("users").Take(&result)

```



**参数**:

- `name`: string

- `args`: ...interface{}



**返回**: `*DB`


**调用**: db.getInstance() strings.Contains() strings.Contains() len() tableRegexp.FindStringSubmatch() len() strings.Split() len() 


---




#### Method: Distinct


```go
func (*DB) Distinct(args ...interface{}) (*DB)
```


Distinct specify distinct fields that you want querying



```
Distinct specify distinct fields that you want querying

	// Select distinct names of users
	db.Distinct("name").Find(&results)
	// Select distinct name/age pairs from users
	db.Distinct("name", "age").Find(&results)

```



**参数**:

- `args`: ...interface{}



**返回**: `*DB`


**调用**: db.getInstance() len() tx.Select() 


---




#### Method: Select


```go
func (*DB) Select(query interface{}, args ...interface{}) (*DB)
```


Select specify fields that you want when querying, creating, updating



```
Select specify fields that you want when querying, creating, updating

Use Select when you only want a subset of the fields. By default, GORM will select all fields.
Select accepts both string arguments and arrays.

	// Select name and age of user using multiple arguments
	db.Select("name", "age").Find(&users)
	// Select name and age of user using an array
	db.Select([]string{"name", "age"}).Find(&users)

```



**参数**:

- `query`: interface{}

- `args`: ...interface{}



**返回**: `*DB`


**调用**: db.getInstance() append() append() tx.AddError() fmt.Errorf() strings.Count() len() len() strings.Count() len() append() append() tx.AddError() fmt.Errorf() 


---




#### Method: Omit


```go
func (*DB) Omit(columns ...string) (*DB)
```


Omit specify fields that you want to ignore when creating, updating and querying



```
Omit specify fields that you want to ignore when creating, updating and querying

```



**参数**:

- `columns`: ...string



**返回**: `*DB`


**调用**: db.getInstance() len() strings.ContainsRune() strings.FieldsFunc() 


---




#### Method: MapColumns


```go
func (*DB) MapColumns(m map[string]string) (*DB)
```


MapColumns modify the column names in the query results to facilitate align to the corresponding structural fields



```
MapColumns modify the column names in the query results to facilitate align to the corresponding structural fields

```



**参数**:

- `m`: map[string]string



**返回**: `*DB`


**调用**: db.getInstance() 


---




#### Method: Where


```go
func (*DB) Where(query interface{}, args ...interface{}) (*DB)
```


Where add conditions



```
Where add conditions

See the [docs] for details on the various formats that where clauses can take. By default, where clauses chain with AND.

	// Find the first user with name jinzhu
	db.Where("name = ?", "jinzhu").First(&user)
	// Find the first user with name jinzhu and age 20
	db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
	// Find the first user with name jinzhu and age not equal to 20
	db.Where("name = ?", "jinzhu").Where("age <> ?", "20").First(&user)

[docs]: https://gorm.io/docs/query.html#Conditions

```



**参数**:

- `query`: interface{}

- `args`: ...interface{}



**返回**: `*DB`


**调用**: db.getInstance() len() 


---




#### Method: Not


```go
func (*DB) Not(query interface{}, args ...interface{}) (*DB)
```


Not add NOT conditions



```
Not add NOT conditions

Not works similarly to where, and has the same syntax.

	// Find the first user with name not equal to jinzhu
	db.Not("name = ?", "jinzhu").First(&user)

```



**参数**:

- `query`: interface{}

- `args`: ...interface{}



**返回**: `*DB`


**调用**: db.getInstance() len() clause.Not() 


---




#### Method: Or


```go
func (*DB) Or(query interface{}, args ...interface{}) (*DB)
```


Or add OR conditions



```
Or add OR conditions

Or is used to chain together queries with an OR.

	// Find the first user with name equal to jinzhu or john
	db.Where("name = ?", "jinzhu").Or("name = ?", "john").First(&user)

```



**参数**:

- `query`: interface{}

- `args`: ...interface{}



**返回**: `*DB`


**调用**: db.getInstance() len() clause.Or() clause.And() 


---




#### Method: Joins


```go
func (*DB) Joins(query string, args ...interface{}) (*DB)
```


Joins specify Joins conditions



```
Joins specify Joins conditions

	db.Joins("Account").Find(&user)
	db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Find(&user)
	db.Joins("Account", DB.Select("id").Where("user_id = users.id AND name = ?", "someName").Model(&Account{}))

```



**参数**:

- `query`: string

- `args`: ...interface{}



**返回**: `*DB`


**调用**: joins() 


---




#### Method: InnerJoins


```go
func (*DB) InnerJoins(query string, args ...interface{}) (*DB)
```


InnerJoins specify inner joins conditions



```
InnerJoins specify inner joins conditions
db.InnerJoins("Account").Find(&user)

```



**参数**:

- `query`: string

- `args`: ...interface{}



**返回**: `*DB`


**调用**: joins() 


---




#### Method: Group


```go
func (*DB) Group(name string) (*DB)
```


Group specify the group method on the find



```
Group specify the group method on the find

	// Select the sum age of users with given names
	db.Model(&User{}).Select("name, sum(age) as total").Group("name").Find(&results)

```



**参数**:

- `name`: string



**返回**: `*DB`


**调用**: db.getInstance() strings.FieldsFunc() len() 


---




#### Method: Having


```go
func (*DB) Having(query interface{}, args ...interface{}) (*DB)
```


Having specify HAVING conditions for GROUP BY



```
Having specify HAVING conditions for GROUP BY

	// Select the sum age of users with name jinzhu
	db.Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "jinzhu").Find(&result)

```



**参数**:

- `query`: interface{}

- `args`: ...interface{}



**返回**: `*DB`


**调用**: db.getInstance() 


---




#### Method: Order


```go
func (*DB) Order(value interface{}) (*DB)
```


Order specify order when retrieving records from database



```
Order specify order when retrieving records from database

	db.Order("name DESC")
	db.Order(clause.OrderByColumn{Column: clause.Column{Name: "name"}, Desc: true})
	db.Order(clause.OrderBy{Columns: []clause.OrderByColumn{
		{Column: clause.Column{Name: "name"}, Desc: true},
		{Column: clause.Column{Name: "age"}, Desc: true},
	}})

```



**参数**:

- `value`: interface{}



**返回**: `*DB`


**调用**: db.getInstance() 


---




#### Method: Limit


```go
func (*DB) Limit(limit int) (*DB)
```


Limit specify the number of records to be retrieved



```
Limit specify the number of records to be retrieved

Limit conditions can be cancelled by using `Limit(-1)`.

	// retrieve 3 users
	db.Limit(3).Find(&users)
	// retrieve 3 users into users1, and all users into users2
	db.Limit(3).Find(&users1).Limit(-1).Find(&users2)

```



**参数**:

- `limit`: int



**返回**: `*DB`


**调用**: db.getInstance() 


---




#### Method: Offset


```go
func (*DB) Offset(offset int) (*DB)
```


Offset specify the number of records to skip before starting to return the records



```
Offset specify the number of records to skip before starting to return the records

Offset conditions can be cancelled by using `Offset(-1)`.

	// select the third user
	db.Offset(2).First(&user)
	// select the first user by cancelling an earlier chained offset
	db.Offset(5).Offset(-1).First(&user)

```



**参数**:

- `offset`: int



**返回**: `*DB`


**调用**: db.getInstance() 


---




#### Method: Scopes


```go
func (*DB) Scopes(funcs ...func(...)) (*DB)
```


Scopes pass current database connection to arguments `func(DB) DB`, which could be used to add conditions dynamically



```
Scopes pass current database connection to arguments `func(DB) DB`, which could be used to add conditions dynamically

	func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
	    return db.Where("amount > ?", 1000)
	}

	func OrderStatus(status []string) func (db *gorm.DB) *gorm.DB {
	    return func (db *gorm.DB) *gorm.DB {
	        return db.Scopes(AmountGreaterThan1000).Where("status in (?)", status)
	    }
	}

	db.Scopes(AmountGreaterThan1000, OrderStatus([]string{"paid", "shipped"})).Find(&orders)

```



**参数**:

- `funcs`: ...func(...)



**返回**: `*DB`


**调用**: db.getInstance() append() 


---






#### Method: Preload


```go
func (*DB) Preload(query string, args ...interface{}) (*DB)
```


Preload preload associations with given conditions



```
Preload preload associations with given conditions

	// get all users, and preload all non-cancelled orders
	db.Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)

```



**参数**:

- `query`: string

- `args`: ...interface{}



**返回**: `*DB`


**调用**: db.getInstance() 


---




#### Method: Attrs


```go
func (*DB) Attrs(attrs ...interface{}) (*DB)
```


Attrs provide attributes used in [FirstOrCreate] or [FirstOrInit]



```
Attrs provide attributes used in [FirstOrCreate] or [FirstOrInit]

Attrs only adds attributes if the record is not found.

	// assign an email if the record is not found
	db.Where(User{Name: "non_existing"}).Attrs(User{Email: "fake@fake.org"}).FirstOrInit(&user)
	// user -> User{Name: "non_existing", Email: "fake@fake.org"}

	// assign an email if the record is not found, otherwise ignore provided email
	db.Where(User{Name: "jinzhu"}).Attrs(User{Email: "fake@fake.org"}).FirstOrInit(&user)
	// user -> User{Name: "jinzhu", Age: 20}

[FirstOrCreate]: https://gorm.io/docs/advanced_query.html#FirstOrCreate
[FirstOrInit]: https://gorm.io/docs/advanced_query.html#FirstOrInit

```



**参数**:

- `attrs`: ...interface{}



**返回**: `*DB`


**调用**: db.getInstance() 


---




#### Method: Assign


```go
func (*DB) Assign(attrs ...interface{}) (*DB)
```


Assign provide attributes used in [FirstOrCreate] or [FirstOrInit]



```
Assign provide attributes used in [FirstOrCreate] or [FirstOrInit]

Assign adds attributes even if the record is found. If using FirstOrCreate, this means that
records will be updated even if they are found.

	// assign an email regardless of if the record is not found
	db.Where(User{Name: "non_existing"}).Assign(User{Email: "fake@fake.org"}).FirstOrInit(&user)
	// user -> User{Name: "non_existing", Email: "fake@fake.org"}

	// assign email regardless of if record is found
	db.Where(User{Name: "jinzhu"}).Assign(User{Email: "fake@fake.org"}).FirstOrInit(&user)
	// user -> User{Name: "jinzhu", Age: 20, Email: "fake@fake.org"}

[FirstOrCreate]: https://gorm.io/docs/advanced_query.html#FirstOrCreate
[FirstOrInit]: https://gorm.io/docs/advanced_query.html#FirstOrInit

```



**参数**:

- `attrs`: ...interface{}



**返回**: `*DB`


**调用**: db.getInstance() 


---




#### Method: Unscoped


```go
func (*DB) Unscoped() (*DB)
```


Unscoped disables the global scope of soft deletion in a query.



```
Unscoped disables the global scope of soft deletion in a query.
By default, GORM uses soft deletion, marking records as "deleted"
by setting a timestamp on a specific field (e.g., `deleted_at`).
Unscoped allows queries to include records marked as deleted,
overriding the soft deletion behavior.
Example:

	var users []User
	db.Unscoped().Find(&users)
	// Retrieves all users, including deleted ones.

```




**返回**: `*DB`


**调用**: db.getInstance() 


---




#### Method: Raw


```go
func (*DB) Raw(sql string, values ...interface{}) (*DB)
```





**参数**:

- `sql`: string

- `values`: ...interface{}



**返回**: `*DB`


**调用**: db.getInstance() strings.Contains() 


---




#### Class: AssociationOpType








---




#### Constant: OpUnlink








---




#### Constant: OpDelete








---




#### Constant: OpUpdate








---




#### Constant: OpCreate








---




#### Struct: Association








---




#### Interface: AssociationAssigner








---




#### Method: Assignments


```go
func (Association) Assignments() ([]Assignment)
```


Assignments implements the Assigner interface so that AssociationOperation can be used as a Set method parameter



```
Assignments implements the Assigner interface so that AssociationOperation can be used as a Set method parameter

```




**返回**: `[]Assignment`



---




#### Method: AssociationAssignments


```go
func (Association) AssociationAssignments() ([]Association)
```


AssociationAssignments implements the AssociationAssigner interface



```
AssociationAssignments implements the AssociationAssigner interface

```




**返回**: `[]Association`



---




#### Interface: Interface








---




#### Class: ClauseBuilder








---




#### Interface: Writer








---




#### Interface: Builder








---




#### Struct: Clause








---




#### Method: Build


```go
func (Clause) Build(builder Builder)
```


Build build clause



```
Build build clause

```



**参数**:

- `builder`: Builder




**调用**: c.Builder() builder.WriteByte() builder.WriteString() builder.WriteByte() builder.WriteByte() builder.WriteByte() 


---




#### Constant: PrimaryKey








---




#### Constant: CurrentTable








---




#### Constant: Associations








---






#### Variable: PrimaryColumn








---




#### Struct: Column








---




#### Struct: Table








---




#### Struct: Delete








---




#### Method: Name


```go
func (Delete) Name() (string)
```






**返回**: `string`



---




#### Method: Build


```go
func (Delete) Build(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: builder.WriteString() builder.WriteByte() builder.WriteString() 


---




#### Method: MergeClause


```go
func (Delete) MergeClause(clause *Clause)
```





**参数**:

- `clause`: *Clause





---




#### Interface: Expression








---




#### Interface: NegationExpressionBuilder








---




#### Struct: Expr








---




#### Method: Build


```go
func (Expr) Build(builder Builder)
```


Build build raw expression



```
Build build raw expression

```



**参数**:

- `builder`: Builder




**调用**: len() processValue() builder.AddVar() builder.WriteByte() len() builder.AddVar() 


---








#### Struct: NamedExpr








---




#### Method: Build


```go
func (NamedExpr) Build(builder Builder)
```


Build build raw expression



```
Build build raw expression

```



**参数**:

- `builder`: Builder




**调用**: make() len() reflect.Indirect() reflectValue.Kind() reflectValue.Type() modelType.NumField() modelType.Field() ast.IsExported() reflectValue.Field() appendFieldsToMap() reflectValue.Field() appendFieldsToMap() reflect.ValueOf() make() string() processValue() builder.AddVar() builder.WriteByte() builder.WriteString() string() builder.WriteByte() len() processValue() builder.AddVar() append() builder.WriteByte() string() builder.AddVar() builder.WriteByte() builder.WriteString() string() 


---














#### Struct: IN








---




#### Method: Build


```go
func (IN) Build(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: builder.WriteQuoted() len() builder.WriteString() builder.WriteString() builder.AddVar() builder.WriteString() builder.AddVar() builder.WriteByte() 


---




#### Method: NegationBuild


```go
func (IN) NegationBuild(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: builder.WriteQuoted() len() builder.WriteString() builder.WriteString() builder.AddVar() builder.WriteString() builder.AddVar() builder.WriteByte() 


---




#### Struct: Eq








---




#### Method: Build


```go
func (Eq) Build(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: builder.WriteQuoted() reflect.ValueOf() rv.Len() builder.WriteString() builder.WriteString() rv.Len() builder.WriteByte() builder.AddVar() rv.Index() builder.WriteByte() eqNil() builder.WriteString() builder.WriteString() builder.AddVar() 


---




#### Method: NegationBuild


```go
func (Eq) NegationBuild(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: Neq() 


---




#### Class: Neq








---




#### Method: Build


```go
func (Neq) Build(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: builder.WriteQuoted() builder.WriteString() reflect.ValueOf() rv.Len() builder.WriteByte() builder.AddVar() rv.Index() builder.WriteByte() eqNil() builder.WriteString() builder.WriteString() builder.AddVar() 


---




#### Method: NegationBuild


```go
func (Neq) NegationBuild(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: Eq() 


---




#### Class: Gt








---




#### Method: Build


```go
func (Gt) Build(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: builder.WriteQuoted() builder.WriteString() builder.AddVar() 


---




#### Method: NegationBuild


```go
func (Gt) NegationBuild(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: Lte() 


---




#### Class: Gte








---




#### Method: Build


```go
func (Gte) Build(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: builder.WriteQuoted() builder.WriteString() builder.AddVar() 


---




#### Method: NegationBuild


```go
func (Gte) NegationBuild(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: Lt() 


---




#### Class: Lt








---




#### Method: Build


```go
func (Lt) Build(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: builder.WriteQuoted() builder.WriteString() builder.AddVar() 


---




#### Method: NegationBuild


```go
func (Lt) NegationBuild(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: Gte() 


---




#### Class: Lte








---




#### Method: Build


```go
func (Lte) Build(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: builder.WriteQuoted() builder.WriteString() builder.AddVar() 


---




#### Method: NegationBuild


```go
func (Lte) NegationBuild(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: Gt() 


---




#### Class: Like








---




#### Method: Build


```go
func (Like) Build(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: builder.WriteQuoted() builder.WriteString() builder.AddVar() 


---




#### Method: NegationBuild


```go
func (Like) NegationBuild(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: builder.WriteQuoted() builder.WriteString() builder.AddVar() 


---




#### Struct: From








---




#### Method: Name


```go
func (From) Name() (string)
```


Name from clause name



```
Name from clause name

```




**返回**: `string`



---




#### Method: Build


```go
func (From) Build(builder Builder)
```


Build build from clause



```
Build build from clause

```



**参数**:

- `builder`: Builder




**调用**: len() builder.WriteByte() builder.WriteQuoted() builder.WriteQuoted() builder.WriteByte() join.Build() 


---




#### Method: MergeClause


```go
func (From) MergeClause(clause *Clause)
```


MergeClause merge from clause



```
MergeClause merge from clause

```



**参数**:

- `clause`: *Clause





---




#### Struct: GroupBy








---




#### Method: Name


```go
func (GroupBy) Name() (string)
```


Name from clause name



```
Name from clause name

```




**返回**: `string`



---




#### Method: Build


```go
func (GroupBy) Build(builder Builder)
```


Build build group by clause



```
Build build group by clause

```



**参数**:

- `builder`: Builder




**调用**: builder.WriteByte() builder.WriteQuoted() len() builder.WriteString() 


---




#### Method: MergeClause


```go
func (GroupBy) MergeClause(clause *Clause)
```


MergeClause merge group by clause



```
MergeClause merge group by clause

```



**参数**:

- `clause`: *Clause




**调用**: make() len() copy() append() make() len() copy() append() len() groupBy.Name() 


---




#### Struct: Insert








---




#### Method: Name


```go
func (Insert) Name() (string)
```


Name insert clause name



```
Name insert clause name

```




**返回**: `string`



---




#### Method: Build


```go
func (Insert) Build(builder Builder)
```


Build build insert clause



```
Build build insert clause

```



**参数**:

- `builder`: Builder




**调用**: builder.WriteString() builder.WriteByte() builder.WriteString() builder.WriteQuoted() builder.WriteQuoted() 


---




#### Method: MergeClause


```go
func (Insert) MergeClause(clause *Clause)
```


MergeClause merge insert clause



```
MergeClause merge insert clause

```



**参数**:

- `clause`: *Clause





---




#### Class: JoinType








---




#### Constant: CrossJoin








---




#### Constant: InnerJoin








---




#### Constant: LeftJoin








---




#### Constant: RightJoin








---




#### Struct: JoinTarget








---




#### Function: Has


```go
func Has(name string) (JoinTarget)
```





**参数**:

- `name`: string



**返回**: `JoinTarget`



---




#### Method: Association


```go
func (JoinType) Association(name string) (JoinTarget)
```





**参数**:

- `name`: string



**返回**: `JoinTarget`



---




#### Method: AssociationFrom


```go
func (JoinType) AssociationFrom(name string, subquery Expression) (JoinTarget)
```





**参数**:

- `name`: string

- `subquery`: Expression



**返回**: `JoinTarget`



---




#### Method: As


```go
func (JoinTarget) As(name string) (JoinTarget)
```





**参数**:

- `name`: string



**返回**: `JoinTarget`



---




#### Struct: Join








---




#### Function: JoinTable


```go
func JoinTable(names ...string) (Table)
```





**参数**:

- `names`: ...string



**返回**: `Table`


**调用**: utils.JoinNestedRelationNames() 


---




#### Method: Build


```go
func (Join) Build(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: builder.WriteString() string() builder.WriteByte() builder.WriteString() builder.WriteQuoted() len() builder.WriteString() len() builder.WriteString() builder.WriteByte() builder.WriteQuoted() builder.WriteByte() 


---




#### Struct: Limit








---




#### Method: Name


```go
func (Limit) Name() (string)
```


Name where clause name



```
Name where clause name

```




**返回**: `string`



---




#### Method: Build


```go
func (Limit) Build(builder Builder)
```


Build build where clause



```
Build build where clause

```



**参数**:

- `builder`: Builder




**调用**: builder.WriteString() builder.AddVar() builder.WriteByte() builder.WriteString() builder.AddVar() 


---




#### Method: MergeClause


```go
func (Limit) MergeClause(clause *Clause)
```


MergeClause merge order by clauses



```
MergeClause merge order by clauses

```



**参数**:

- `clause`: *Clause





---




#### Constant: LockingStrengthUpdate








---




#### Constant: LockingStrengthShare








---




#### Constant: LockingOptionsSkipLocked








---




#### Constant: LockingOptionsNoWait








---




#### Struct: Locking








---




#### Method: Name


```go
func (Locking) Name() (string)
```


Name where clause name



```
Name where clause name

```




**返回**: `string`



---




#### Method: Build


```go
func (Locking) Build(builder Builder)
```


Build build where clause



```
Build build where clause

```



**参数**:

- `builder`: Builder




**调用**: builder.WriteString() builder.WriteString() builder.WriteQuoted() builder.WriteByte() builder.WriteString() 


---




#### Method: MergeClause


```go
func (Locking) MergeClause(clause *Clause)
```


MergeClause merge order by clauses



```
MergeClause merge order by clauses

```



**参数**:

- `clause`: *Clause





---




#### Struct: OnConflict








---




#### Method: Name


```go
func (OnConflict) Name() (string)
```






**返回**: `string`



---




#### Method: Build


```go
func (OnConflict) Build(builder Builder)
```


Build build onConflict clause



```
Build build onConflict clause

```



**参数**:

- `builder`: Builder




**调用**: builder.WriteString() builder.WriteString() builder.WriteByte() len() builder.WriteByte() builder.WriteByte() builder.WriteQuoted() builder.WriteString() len() builder.WriteString() builder.WriteByte() builder.WriteString() builder.WriteString() len() builder.WriteString() builder.WriteByte() 


---




#### Method: MergeClause


```go
func (OnConflict) MergeClause(clause *Clause)
```


MergeClause merge onConflict clauses



```
MergeClause merge onConflict clauses

```



**参数**:

- `clause`: *Clause





---




#### Struct: OrderByColumn








---




#### Struct: OrderBy








---




#### Method: Name


```go
func (OrderBy) Name() (string)
```


Name where clause name



```
Name where clause name

```




**返回**: `string`



---




#### Method: Build


```go
func (OrderBy) Build(builder Builder)
```


Build build where clause



```
Build build where clause

```



**参数**:

- `builder`: Builder




**调用**: builder.WriteByte() builder.WriteQuoted() builder.WriteString() 


---




#### Method: MergeClause


```go
func (OrderBy) MergeClause(clause *Clause)
```


MergeClause merge order by clauses



```
MergeClause merge order by clauses

```



**参数**:

- `clause`: *Clause




**调用**: len() make() len() copy() append() 


---




#### Struct: Returning








---




#### Method: Name


```go
func (Returning) Name() (string)
```


Name where clause name



```
Name where clause name

```




**返回**: `string`



---




#### Method: Build


```go
func (Returning) Build(builder Builder)
```


Build build where clause



```
Build build where clause

```



**参数**:

- `builder`: Builder




**调用**: len() builder.WriteByte() builder.WriteQuoted() builder.WriteByte() 


---




#### Method: MergeClause


```go
func (Returning) MergeClause(clause *Clause)
```


MergeClause merge order by clauses



```
MergeClause merge order by clauses

```



**参数**:

- `clause`: *Clause




**调用**: len() append() 


---




#### Struct: Select








---




#### Method: Name


```go
func (Select) Name() (string)
```






**返回**: `string`



---




#### Method: Build


```go
func (Select) Build(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: len() builder.WriteString() builder.WriteByte() builder.WriteQuoted() builder.WriteByte() 


---




#### Method: MergeClause


```go
func (Select) MergeClause(clause *Clause)
```





**参数**:

- `clause`: *Clause





---




#### Struct: CommaExpression








---




#### Method: Build


```go
func (CommaExpression) Build(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: builder.WriteString() expr.Build() 


---




#### Class: Set








---




#### Struct: Assignment








---




#### Interface: Assigner








---




#### Method: Name


```go
func (Set) Name() (string)
```






**返回**: `string`



---




#### Method: Build


```go
func (Set) Build(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: len() builder.WriteByte() builder.WriteQuoted() builder.WriteByte() builder.AddVar() builder.WriteQuoted() builder.WriteByte() builder.WriteQuoted() 


---




#### Method: MergeClause


```go
func (Set) MergeClause(clause *Clause)
```


MergeClause merge assignments clauses



```
MergeClause merge assignments clauses

```



**参数**:

- `clause`: *Clause




**调用**: make() len() copy() Set() 


---




#### Method: Assignments


```go
func (Set) Assignments() ([]Assignment)
```


Assignments implements Assigner for Set.



```
Assignments implements Assigner for Set.

```




**返回**: `[]Assignment`



---




#### Function: Assignments


```go
func Assignments(values map[string]interface{}) (Set)
```





**参数**:

- `values`: map[string]interface{}



**返回**: `Set`


**调用**: make() len() append() sort.Strings() make() len() 


---




#### Function: AssignmentColumns


```go
func AssignmentColumns(values []string) (Set)
```





**参数**:

- `values`: []string



**返回**: `Set`


**调用**: make() len() 


---




#### Method: Assignments


```go
func (Assignment) Assignments() ([]Assignment)
```


Assignments implements Assigner for a single Assignment.



```
Assignments implements Assigner for a single Assignment.

```




**返回**: `[]Assignment`



---




#### Struct: Update








---




#### Method: Name


```go
func (Update) Name() (string)
```


Name update clause name



```
Name update clause name

```




**返回**: `string`



---




#### Method: Build


```go
func (Update) Build(builder Builder)
```


Build build update clause



```
Build build update clause

```



**参数**:

- `builder`: Builder




**调用**: builder.WriteString() builder.WriteByte() builder.WriteQuoted() builder.WriteQuoted() 


---




#### Method: MergeClause


```go
func (Update) MergeClause(clause *Clause)
```


MergeClause merge update clause



```
MergeClause merge update clause

```



**参数**:

- `clause`: *Clause





---




#### Struct: Values








---




#### Method: Name


```go
func (Values) Name() (string)
```


Name from clause name



```
Name from clause name

```




**返回**: `string`



---




#### Method: Build


```go
func (Values) Build(builder Builder)
```


Build build from clause



```
Build build from clause

```



**参数**:

- `builder`: Builder




**调用**: len() builder.WriteByte() builder.WriteByte() builder.WriteQuoted() builder.WriteByte() builder.WriteString() builder.WriteByte() builder.WriteByte() builder.AddVar() builder.WriteByte() builder.WriteString() 


---




#### Method: MergeClause


```go
func (Values) MergeClause(clause *Clause)
```


MergeClause merge values clauses



```
MergeClause merge values clauses

```



**参数**:

- `clause`: *Clause





---




#### Constant: AndWithSpace








---




#### Constant: OrWithSpace








---




#### Struct: Where








---




#### Method: Name


```go
func (Where) Name() (string)
```


Name where clause name



```
Name where clause name

```




**返回**: `string`



---




#### Method: Build


```go
func (Where) Build(builder Builder)
```


Build build where clause



```
Build build where clause

```



**参数**:

- `builder`: Builder




**调用**: len() len() buildExprs() 


---




#### Method: MergeClause


```go
func (Where) MergeClause(clause *Clause)
```


MergeClause merge where clauses



```
MergeClause merge where clauses

```



**参数**:

- `clause`: *Clause




**调用**: make() len() len() copy() copy() len() 


---




#### Function: And


```go
func And(exprs ...Expression) (Expression)
```





**参数**:

- `exprs`: ...Expression



**返回**: `Expression`


**调用**: len() len() 


---




#### Struct: AndConditions








---




#### Method: Build


```go
func (AndConditions) Build(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: len() builder.WriteByte() buildExprs() builder.WriteByte() buildExprs() 


---




#### Function: Or


```go
func Or(exprs ...Expression) (Expression)
```





**参数**:

- `exprs`: ...Expression



**返回**: `Expression`


**调用**: len() 


---




#### Struct: OrConditions








---




#### Method: Build


```go
func (OrConditions) Build(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: len() builder.WriteByte() buildExprs() builder.WriteByte() buildExprs() 


---




#### Function: Not


```go
func Not(exprs ...Expression) (Expression)
```





**参数**:

- `exprs`: ...Expression



**返回**: `Expression`


**调用**: len() len() 


---




#### Struct: NotConditions








---




#### Method: Build


```go
func (NotConditions) Build(builder Builder)
```





**参数**:

- `builder`: Builder




**调用**: len() builder.WriteByte() builder.WriteString() negationBuilder.NegationBuild() builder.WriteString() strings.ToUpper() strings.Contains() strings.Contains() builder.WriteByte() c.Build() builder.WriteByte() len() builder.WriteByte() builder.WriteString() len() builder.WriteByte() builder.WriteString() builder.WriteString() strings.ToUpper() strings.Contains() strings.Contains() builder.WriteByte() c.Build() builder.WriteByte() len() builder.WriteByte() 


---




#### Struct: With








---




#### Variable: ErrRecordNotFound


ErrRecordNotFound record not found error



```
ErrRecordNotFound record not found error

```






---




#### Variable: ErrInvalidTransaction


ErrInvalidTransaction invalid transaction when you are trying to `Commit` or `Rollback`



```
ErrInvalidTransaction invalid transaction when you are trying to `Commit` or `Rollback`

```






---




#### Variable: ErrNotImplemented


ErrNotImplemented not implemented



```
ErrNotImplemented not implemented

```






---




#### Variable: ErrMissingWhereClause


ErrMissingWhereClause missing where clause



```
ErrMissingWhereClause missing where clause

```






---




#### Variable: ErrUnsupportedRelation


ErrUnsupportedRelation unsupported relations



```
ErrUnsupportedRelation unsupported relations

```






---




#### Variable: ErrPrimaryKeyRequired


ErrPrimaryKeyRequired primary keys required



```
ErrPrimaryKeyRequired primary keys required

```






---




#### Variable: ErrModelValueRequired


ErrModelValueRequired model value required



```
ErrModelValueRequired model value required

```






---




#### Variable: ErrModelAccessibleFieldsRequired


ErrModelAccessibleFieldsRequired model accessible fields required



```
ErrModelAccessibleFieldsRequired model accessible fields required

```






---




#### Variable: ErrSubQueryRequired


ErrSubQueryRequired sub query required



```
ErrSubQueryRequired sub query required

```






---




#### Variable: ErrInvalidData


ErrInvalidData unsupported data



```
ErrInvalidData unsupported data

```






---




#### Variable: ErrUnsupportedDriver


ErrUnsupportedDriver unsupported driver



```
ErrUnsupportedDriver unsupported driver

```






---




#### Variable: ErrRegistered


ErrRegistered registered



```
ErrRegistered registered

```






---




#### Variable: ErrInvalidField


ErrInvalidField invalid field



```
ErrInvalidField invalid field

```






---




#### Variable: ErrEmptySlice


ErrEmptySlice empty slice found



```
ErrEmptySlice empty slice found

```






---




#### Variable: ErrDryRunModeUnsupported


ErrDryRunModeUnsupported dry run mode unsupported



```
ErrDryRunModeUnsupported dry run mode unsupported

```






---




#### Variable: ErrInvalidDB


ErrInvalidDB invalid db



```
ErrInvalidDB invalid db

```






---




#### Variable: ErrInvalidValue


ErrInvalidValue invalid value



```
ErrInvalidValue invalid value

```






---




#### Variable: ErrInvalidValueOfLength


ErrInvalidValueOfLength invalid values do not match length



```
ErrInvalidValueOfLength invalid values do not match length

```






---




#### Variable: ErrPreloadNotAllowed


ErrPreloadNotAllowed preload is not allowed when count is used



```
ErrPreloadNotAllowed preload is not allowed when count is used

```






---




#### Variable: ErrDuplicatedKey


ErrDuplicatedKey occurs when there is a unique key constraint violation



```
ErrDuplicatedKey occurs when there is a unique key constraint violation

```






---




#### Variable: ErrForeignKeyViolated


ErrForeignKeyViolated occurs when there is a foreign key constraint violation



```
ErrForeignKeyViolated occurs when there is a foreign key constraint violation

```






---




#### Variable: ErrCheckConstraintViolated


ErrCheckConstraintViolated occurs when there is a check constraint violation



```
ErrCheckConstraintViolated occurs when there is a check constraint violation

```






---




#### Method: Create


```go
func (*DB) Create(value interface{}) (*DB)
```


Create inserts value, returning the inserted data's primary key in value's id



```
Create inserts value, returning the inserted data's primary key in value's id

```



**参数**:

- `value`: interface{}



**返回**: `*DB`


**调用**: db.CreateInBatches() db.getInstance() 


---




#### Method: CreateInBatches


```go
func (*DB) CreateInBatches(value interface{}, batchSize int) (*DB)
```


CreateInBatches inserts value in batches of batchSize



```
CreateInBatches inserts value in batches of batchSize

```



**参数**:

- `value`: interface{}

- `batchSize`: int



**返回**: `*DB`


**调用**: reflect.Indirect() reflect.ValueOf() reflectValue.Kind() db.getInstance() reflectValue.Len() tx.getInstance() reflectValue.Slice() tx.AddError() callFc() tx.Session() tx.AddError() tx.Transaction() db.getInstance() 


---






#### Method: Save


```go
func (*DB) Save(value interface{}) (*DB)
```


Save updates value in database. If value doesn't contain a matching primary key, value is inserted.



```
Save updates value in database. If value doesn't contain a matching primary key, value is inserted.

```



**参数**:

- `value`: interface{}



**返回**: `*DB`


**调用**: db.getInstance() reflect.Indirect() reflect.ValueOf() reflectValue.Kind() reflectValue.Kind() reflect.Indirect() reflectValue.Kind() tx.Clauses() tx.Set() pf.ValueOf() len() append() tx.Session() tx.Session() 


---




#### Method: First


```go
func (*DB) First(dest interface{}, conds ...interface{}) (*DB)
```


First finds the first record ordered by primary key, matching given conditions conds



```
First finds the first record ordered by primary key, matching given conditions conds

```



**参数**:

- `dest`: interface{}

- `conds`: ...interface{}



**返回**: `*DB`


**调用**: db.Limit() len() len() 


---




#### Method: Take


```go
func (*DB) Take(dest interface{}, conds ...interface{}) (*DB)
```


Take finds the first record returned by the database in no specified order, matching given conditions conds



```
Take finds the first record returned by the database in no specified order, matching given conditions conds

```



**参数**:

- `dest`: interface{}

- `conds`: ...interface{}



**返回**: `*DB`


**调用**: db.Limit() len() len() 


---




#### Method: Last


```go
func (*DB) Last(dest interface{}, conds ...interface{}) (*DB)
```


Last finds the last record ordered by primary key, matching given conditions conds



```
Last finds the last record ordered by primary key, matching given conditions conds

```



**参数**:

- `dest`: interface{}

- `conds`: ...interface{}



**返回**: `*DB`


**调用**: db.Limit() len() len() 


---




#### Method: Find


```go
func (*DB) Find(dest interface{}, conds ...interface{}) (*DB)
```


Find finds all records matching given conditions conds



```
Find finds all records matching given conditions conds

```



**参数**:

- `dest`: interface{}

- `conds`: ...interface{}



**返回**: `*DB`


**调用**: db.getInstance() len() len() 


---




#### Method: FindInBatches


```go
func (*DB) FindInBatches(dest interface{}, batchSize int, fc func(...)) (*DB)
```


FindInBatches finds all records in batches of batchSize



```
FindInBatches finds all records in batches of batchSize

```



**参数**:

- `dest`: interface{}

- `batchSize`: int

- `fc`: func(...)



**返回**: `*DB`


**调用**: db.Order() tx.Offset() queryDB.Limit() result.Session() tx.AddError() fc() tx.AddError() int() int() reflect.Indirect() reflect.ValueOf() tx.AddError() resultsValue.Index() resultsValue.Len() tx.AddError() tx.Clauses() 


---
















#### Method: FirstOrInit


```go
func (*DB) FirstOrInit(dest interface{}, conds ...interface{}) (*DB)
```


FirstOrInit finds the first matching record, otherwise if not found initializes a new instance with given conds.



```
FirstOrInit finds the first matching record, otherwise if not found initializes a new instance with given conds.
Each conds must be a struct or map.

FirstOrInit never modifies the database. It is often used with Assign and Attrs.

	// assign an email if the record is not found
	db.Where(User{Name: "non_existing"}).Attrs(User{Email: "fake@fake.org"}).FirstOrInit(&user)
	// user -> User{Name: "non_existing", Email: "fake@fake.org"}

	// assign email regardless of if record is found
	db.Where(User{Name: "jinzhu"}).Assign(User{Email: "fake@fake.org"}).FirstOrInit(&user)
	// user -> User{Name: "jinzhu", Age: 20, Email: "fake@fake.org"}

```



**参数**:

- `dest`: interface{}

- `conds`: ...interface{}



**返回**: `*DB`


**调用**: db.Limit() queryTx.Find() tx.assignInterfacesToValue() len() tx.assignInterfacesToValue() len() tx.assignInterfacesToValue() 


---




#### Method: FirstOrCreate


```go
func (*DB) FirstOrCreate(dest interface{}, conds ...interface{}) (*DB)
```


FirstOrCreate finds the first matching record, otherwise if not found creates a new instance with given conds.



```
FirstOrCreate finds the first matching record, otherwise if not found creates a new instance with given conds.
Each conds must be a struct or map.

Using FirstOrCreate in conjunction with Assign will result in an update to the database even if the record exists.

	// assign an email if the record is not found
	result := db.Where(User{Name: "non_existing"}).Attrs(User{Email: "fake@fake.org"}).FirstOrCreate(&user)
	// user -> User{Name: "non_existing", Email: "fake@fake.org"}
	// result.RowsAffected -> 1

	// assign email regardless of if record is found
	result := db.Where(User{Name: "jinzhu"}).Assign(User{Email: "fake@fake.org"}).FirstOrCreate(&user)
	// user -> User{Name: "jinzhu", Age: 20, Email: "fake@fake.org"}
	// result.RowsAffected -> 1

```



**参数**:

- `dest`: interface{}

- `conds`: ...interface{}



**返回**: `*DB`


**调用**: db.getInstance() db.Session() queryTx.Find() result.assignInterfacesToValue() len() result.assignInterfacesToValue() len() result.assignInterfacesToValue() tx.Create() len() len() append() tx.Model() 


---




#### Method: Update


```go
func (*DB) Update(column string, value interface{}) (*DB)
```


Update updates column with value using callbacks. Reference: https://gorm.io/docs/update.html#Update-Changed-Fields



```
Update updates column with value using callbacks. Reference: https://gorm.io/docs/update.html#Update-Changed-Fields

```



**参数**:

- `column`: string

- `value`: interface{}



**返回**: `*DB`


**调用**: db.getInstance() 


---




#### Method: Updates


```go
func (*DB) Updates(values interface{}) (*DB)
```


Updates updates attributes using callbacks. values must be a struct or map. Reference: https://gorm.io/docs/update.html#Update-Changed-Fields



```
Updates updates attributes using callbacks. values must be a struct or map. Reference: https://gorm.io/docs/update.html#Update-Changed-Fields

```



**参数**:

- `values`: interface{}



**返回**: `*DB`


**调用**: db.getInstance() 


---




#### Method: UpdateColumn


```go
func (*DB) UpdateColumn(column string, value interface{}) (*DB)
```





**参数**:

- `column`: string

- `value`: interface{}



**返回**: `*DB`


**调用**: db.getInstance() 


---




#### Method: UpdateColumns


```go
func (*DB) UpdateColumns(values interface{}) (*DB)
```





**参数**:

- `values`: interface{}



**返回**: `*DB`


**调用**: db.getInstance() 


---




#### Method: Delete


```go
func (*DB) Delete(value interface{}, conds ...interface{}) (*DB)
```


Delete deletes value matching given conditions. If value contains primary key it is included in the conditions. If



```
Delete deletes value matching given conditions. If value contains primary key it is included in the conditions. If
value includes a deleted_at field, then Delete performs a soft delete instead by setting deleted_at with the current
time if null.

```



**参数**:

- `value`: interface{}

- `conds`: ...interface{}



**返回**: `*DB`


**调用**: db.getInstance() len() len() 


---




#### Method: Count


```go
func (*DB) Count(count *int64) (*DB)
```





**参数**:

- `count`: *int64



**返回**: `*DB`


**调用**: db.getInstance() delete() len() strings.HasPrefix() strings.TrimSpace() strings.ToLower() len() strings.FieldsFunc() len() len() strings.ToUpper() delete() 


---




#### Method: Row


```go
func (*DB) Row() (*sql.Row)
```






**返回**: `*sql.Row`


**调用**: db.getInstance() ErrDryRunModeUnsupported.Error() 


---




#### Method: Rows


```go
func (*DB) Rows() (*sql.Rows, error)
```






**返回**: `*sql.Rows`


**调用**: db.getInstance() 


---




#### Method: Scan


```go
func (*DB) Scan(dest interface{}) (*DB)
```


Scan scans selected value to the struct dest



```
Scan scans selected value to the struct dest

```



**参数**:

- `dest`: interface{}



**返回**: `*DB`


**调用**: db.getInstance() tx.Rows() rows.Next() tx.ScanRows() tx.AddError() rows.Err() tx.AddError() rows.Close() currentLogger.Trace() 


---




#### Method: Pluck


```go
func (*DB) Pluck(column string, dest interface{}) (*DB)
```


Pluck queries a single column from a model, returning in the slice dest. E.g.:



```
Pluck queries a single column from a model, returning in the slice dest. E.g.:

	var ages []int64
	db.Model(&users).Pluck("age", &ages)

```



**参数**:

- `column`: string

- `dest`: interface{}



**返回**: `*DB`


**调用**: db.getInstance() len() strings.FieldsFunc() len() 


---




#### Method: ScanRows


```go
func (*DB) ScanRows(rows *sql.Rows, dest interface{}) (error)
```





**参数**:

- `rows`: *sql.Rows

- `dest`: interface{}



**返回**: `error`


**调用**: db.getInstance() errors.Is() tx.AddError() reflect.ValueOf() elem.IsValid() reflect.New() Scan() 


---




#### Method: Connection


```go
func (*DB) Connection(fc func(...)) (error)
```


Connection uses a db connection to execute an arbitrary number of commands in fc. When finished, the connection is



```
Connection uses a db connection to execute an arbitrary number of commands in fc. When finished, the connection is
returned to the connection pool.

```



**参数**:

- `fc`: func(...)



**返回**: `error`


**调用**: db.getInstance() tx.DB() sqlDB.Conn() conn.Close() fc() 


---




#### Method: Transaction


```go
func (*DB) Transaction(fc func(...), opts ...*sql.TxOptions) (error)
```


Transaction start a transaction as a block, return error will rollback, otherwise to commit. Transaction executes an



```
Transaction start a transaction as a block, return error will rollback, otherwise to commit. Transaction executes an
arbitrary number of commands in fc within a transaction. On success the changes are committed; if an error occurs
they are rolled back.

```



**参数**:

- `fc`: func(...)

- `opts`: ...*sql.TxOptions



**返回**: `error`


**调用**: new() db.SavePoint() fmt.Sprintf() db.RollbackTo() fmt.Sprintf() fc() db.Session() db.Begin() tx.Rollback() fc() tx.Commit() 


---




#### Method: Begin


```go
func (*DB) Begin(opts ...*sql.TxOptions) (*DB)
```


Begin begins a transaction with any transaction options opts



```
Begin begins a transaction with any transaction options opts

```



**参数**:

- `opts`: ...*sql.TxOptions



**返回**: `*DB`


**调用**: db.getInstance() len() ctx.Deadline() context.WithTimeout() beginner.BeginTx() beginner.BeginTx() tx.AddError() 


---










#### Method: Commit


```go
func (*DB) Commit() (*DB)
```


Commit commits the changes in a transaction



```
Commit commits the changes in a transaction

```




**返回**: `*DB`


**调用**: reflect.ValueOf() db.AddError() committer.Commit() db.AddError() 


---




#### Method: Rollback


```go
func (*DB) Rollback() (*DB)
```


Rollback rollbacks the changes in a transaction



```
Rollback rollbacks the changes in a transaction

```




**返回**: `*DB`


**调用**: reflect.ValueOf() db.AddError() committer.Rollback() db.AddError() 


---




#### Method: SavePoint


```go
func (*DB) SavePoint(name string) (*DB)
```





**参数**:

- `name`: string



**返回**: `*DB`


**调用**: db.AddError() savePointer.SavePoint() db.AddError() 


---








#### Method: RollbackTo


```go
func (*DB) RollbackTo(name string) (*DB)
```





**参数**:

- `name`: string



**返回**: `*DB`


**调用**: db.AddError() savePointer.RollbackTo() db.AddError() 


---








#### Method: Exec


```go
func (*DB) Exec(sql string, values ...interface{}) (*DB)
```


Exec executes raw sql



```
Exec executes raw sql

```



**参数**:

- `sql`: string

- `values`: ...interface{}



**返回**: `*DB`


**调用**: db.getInstance() strings.Contains() 


---






#### Method: ModifyStatement


```go
func (*result) ModifyStatement(stmt *Statement)
```





**参数**:

- `stmt`: *Statement





---




#### Method: Build


```go
func (result) Build(clause.Builder)
```


Build implements clause.Expression interface



```
Build implements clause.Expression interface

```






---




#### Function: WithResult


```go
func WithResult() (*result)
```






**返回**: `*result`



---




#### Interface: Interface








---




#### Interface: CreateInterface








---




#### Interface: ChainInterface








---




#### Interface: SetUpdateOnlyInterface








---




#### Interface: SetCreateOrUpdateInterface








---




#### Interface: ExecInterface








---




#### Interface: JoinBuilder








---




#### Interface: PreloadBuilder








---






#### Function: G


```go
func G(db *DB, opts ...clause.Expression) (interface{})
```





**参数**:

- `db`: *DB

- `opts`: ...clause.Expression



**返回**: `interface{}`


**调用**: make() len() append() db.Clauses() 


---








#### Method: Raw


```go
func (*interface{}) Raw(sql string, values ...interface{}) (interface{})
```





**参数**:

- `sql`: string

- `values`: ...interface{}



**返回**: `interface{}`


**调用**: append() db.Model() 


---






#### Method: Exec


```go
func (*interface{}) Exec(ctx context.Context, sql string, values ...interface{}) (error)
```





**参数**:

- `ctx`: context.Context

- `sql`: string

- `values`: ...interface{}



**返回**: `error`


**调用**: c.apply() 


---








#### Method: Table


```go
func (interface{}) Table(name string, args ...interface{}) (interface{})
```





**参数**:

- `name`: string

- `args`: ...interface{}



**返回**: `interface{}`


**调用**: c.with() db.Table() 


---




#### Method: Select


```go
func (interface{}) Select(query string, args ...interface{}) (interface{})
```





**参数**:

- `query`: string

- `args`: ...interface{}



**返回**: `interface{}`


**调用**: c.with() db.Select() 


---




#### Method: Omit


```go
func (interface{}) Omit(columns ...string) (interface{})
```





**参数**:

- `columns`: ...string



**返回**: `interface{}`


**调用**: c.with() db.Omit() 


---




#### Method: Set


```go
func (interface{}) Set(assignments ...clause.Assigner) (interface{})
```





**参数**:

- `assignments`: ...clause.Assigner



**返回**: `interface{}`


**调用**: c.processSet() 


---




#### Method: Create


```go
func (interface{}) Create(ctx context.Context, r *T) (error)
```





**参数**:

- `ctx`: context.Context

- `r`: *T



**返回**: `error`



---




#### Method: CreateInBatches


```go
func (interface{}) CreateInBatches(ctx context.Context, r *[]T, batchSize int) (error)
```





**参数**:

- `ctx`: context.Context

- `r`: *[]T

- `batchSize`: int



**返回**: `error`



---












#### Method: Table


```go
func (interface{}) Table(name string, args ...interface{}) (interface{})
```





**参数**:

- `name`: string

- `args`: ...interface{}



**返回**: `interface{}`


**调用**: c.with() db.Table() 


---




#### Method: Scopes


```go
func (interface{}) Scopes(scopes ...func(...)) (interface{})
```





**参数**:

- `scopes`: ...func(...)



**返回**: `interface{}`


**调用**: c.with() fc() 


---




#### Method: Where


```go
func (interface{}) Where(query interface{}, args ...interface{}) (interface{})
```





**参数**:

- `query`: interface{}

- `args`: ...interface{}



**返回**: `interface{}`


**调用**: c.with() db.Where() 


---




#### Method: Not


```go
func (interface{}) Not(query interface{}, args ...interface{}) (interface{})
```





**参数**:

- `query`: interface{}

- `args`: ...interface{}



**返回**: `interface{}`


**调用**: c.with() db.Not() 


---




#### Method: Or


```go
func (interface{}) Or(query interface{}, args ...interface{}) (interface{})
```





**参数**:

- `query`: interface{}

- `args`: ...interface{}



**返回**: `interface{}`


**调用**: c.with() db.Or() 


---




#### Method: Limit


```go
func (interface{}) Limit(offset int) (interface{})
```





**参数**:

- `offset`: int



**返回**: `interface{}`


**调用**: c.with() db.Limit() 


---




#### Method: Offset


```go
func (interface{}) Offset(offset int) (interface{})
```





**参数**:

- `offset`: int



**返回**: `interface{}`


**调用**: c.with() db.Offset() 


---






#### Method: Where


```go
func (*joinBuilder) Where(query interface{}, args ...interface{}) (JoinBuilder)
```





**参数**:

- `query`: interface{}

- `args`: ...interface{}



**返回**: `JoinBuilder`



---




#### Method: Or


```go
func (*joinBuilder) Or(query interface{}, args ...interface{}) (JoinBuilder)
```





**参数**:

- `query`: interface{}

- `args`: ...interface{}



**返回**: `JoinBuilder`



---




#### Method: Not


```go
func (*joinBuilder) Not(query interface{}, args ...interface{}) (JoinBuilder)
```





**参数**:

- `query`: interface{}

- `args`: ...interface{}



**返回**: `JoinBuilder`



---




#### Method: Select


```go
func (*joinBuilder) Select(columns ...string) (JoinBuilder)
```





**参数**:

- `columns`: ...string



**返回**: `JoinBuilder`



---




#### Method: Omit


```go
func (*joinBuilder) Omit(columns ...string) (JoinBuilder)
```





**参数**:

- `columns`: ...string



**返回**: `JoinBuilder`



---






#### Method: Where


```go
func (*preloadBuilder) Where(query interface{}, args ...interface{}) (PreloadBuilder)
```





**参数**:

- `query`: interface{}

- `args`: ...interface{}



**返回**: `PreloadBuilder`



---




#### Method: Or


```go
func (*preloadBuilder) Or(query interface{}, args ...interface{}) (PreloadBuilder)
```





**参数**:

- `query`: interface{}

- `args`: ...interface{}



**返回**: `PreloadBuilder`



---




#### Method: Not


```go
func (*preloadBuilder) Not(query interface{}, args ...interface{}) (PreloadBuilder)
```





**参数**:

- `query`: interface{}

- `args`: ...interface{}



**返回**: `PreloadBuilder`



---




#### Method: Select


```go
func (*preloadBuilder) Select(columns ...string) (PreloadBuilder)
```





**参数**:

- `columns`: ...string



**返回**: `PreloadBuilder`



---




#### Method: Omit


```go
func (*preloadBuilder) Omit(columns ...string) (PreloadBuilder)
```





**参数**:

- `columns`: ...string



**返回**: `PreloadBuilder`



---




#### Method: Limit


```go
func (*preloadBuilder) Limit(limit int) (PreloadBuilder)
```





**参数**:

- `limit`: int



**返回**: `PreloadBuilder`



---




#### Method: Offset


```go
func (*preloadBuilder) Offset(offset int) (PreloadBuilder)
```





**参数**:

- `offset`: int



**返回**: `PreloadBuilder`



---




#### Method: Order


```go
func (*preloadBuilder) Order(value interface{}) (PreloadBuilder)
```





**参数**:

- `value`: interface{}



**返回**: `PreloadBuilder`



---




#### Method: LimitPerRecord


```go
func (*preloadBuilder) LimitPerRecord(num int) (PreloadBuilder)
```





**参数**:

- `num`: int



**返回**: `PreloadBuilder`



---




#### Method: Joins


```go
func (interface{}) Joins(jt clause.JoinTarget, on func(...)) (interface{})
```





**参数**:

- `jt`: clause.JoinTarget

- `on`: func(...)



**返回**: `interface{}`


**调用**: c.with() clause.JoinTable() strings.Split() db.Session() on() db.AddError() db.getInstance() len() len() fmt.Sprintf() append() append() sort.Slice() 


---




#### Method: Select


```go
func (interface{}) Select(query string, args ...interface{}) (interface{})
```





**参数**:

- `query`: string

- `args`: ...interface{}



**返回**: `interface{}`


**调用**: c.with() db.Select() 


---




#### Method: Omit


```go
func (interface{}) Omit(columns ...string) (interface{})
```





**参数**:

- `columns`: ...string



**返回**: `interface{}`


**调用**: c.with() db.Omit() 


---




#### Method: MapColumns


```go
func (interface{}) MapColumns(m map[string]string) (interface{})
```





**参数**:

- `m`: map[string]string



**返回**: `interface{}`


**调用**: c.with() db.MapColumns() 


---




#### Method: Set


```go
func (interface{}) Set(assignments ...clause.Assigner) (interface{})
```





**参数**:

- `assignments`: ...clause.Assigner



**返回**: `interface{}`


**调用**: c.processSet() 


---




#### Method: Distinct


```go
func (interface{}) Distinct(args ...interface{}) (interface{})
```





**参数**:

- `args`: ...interface{}



**返回**: `interface{}`


**调用**: c.with() db.Distinct() 


---




#### Method: Group


```go
func (interface{}) Group(name string) (interface{})
```





**参数**:

- `name`: string



**返回**: `interface{}`


**调用**: c.with() db.Group() 


---




#### Method: Having


```go
func (interface{}) Having(query interface{}, args ...interface{}) (interface{})
```





**参数**:

- `query`: interface{}

- `args`: ...interface{}



**返回**: `interface{}`


**调用**: c.with() db.Having() 


---




#### Method: Order


```go
func (interface{}) Order(value interface{}) (interface{})
```





**参数**:

- `value`: interface{}



**返回**: `interface{}`


**调用**: c.with() db.Order() 


---




#### Method: Preload


```go
func (interface{}) Preload(association string, query func(...)) (interface{})
```





**参数**:

- `association`: string

- `query`: func(...)



**返回**: `interface{}`


**调用**: c.with() db.Preload() tx.getInstance() query() db.AddError() strings.Split() len() db.AddError() fmt.Errorf() db.AddError() fmt.Errorf() tx.AddError() fmt.Errorf() append() len() append() len() append() append() append() append() append() 


---






#### Method: Delete


```go
func (interface{}) Delete(ctx context.Context) (int, error)
```





**参数**:

- `ctx`: context.Context



**返回**: `int`


**调用**: new() int() 


---




#### Method: Update


```go
func (interface{}) Update(ctx context.Context, name string, value any) (int, error)
```





**参数**:

- `ctx`: context.Context

- `name`: string

- `value`: any



**返回**: `int`


**调用**: int() 


---






#### Method: Updates


```go
func (interface{}) Updates(ctx context.Context, t T) (int, error)
```





**参数**:

- `ctx`: context.Context

- `t`: T



**返回**: `int`


**调用**: int() 


---




#### Method: Count


```go
func (interface{}) Count(ctx context.Context, column string) (int64, error)
```





**参数**:

- `ctx`: context.Context

- `column`: string



**返回**: `int64`



---






#### Method: Build


```go
func (interface{}) Build(builder clause.Builder)
```





**参数**:

- `builder`: clause.Builder




**调用**: c.getInstance() make() len() append() subdb.BindVarTo() strings.Replace() bindvar.String() strings.Contains() append() builder.WriteString() 


---










#### Method: First


```go
func (interface{}) First(ctx context.Context) (T, error)
```





**参数**:

- `ctx`: context.Context



**返回**: `T`



---






#### Method: Scan


```go
func (interface{}) Scan(ctx context.Context, result interface{}) (error)
```





**参数**:

- `ctx`: context.Context

- `result`: interface{}



**返回**: `error`



---






#### Method: Last


```go
func (interface{}) Last(ctx context.Context) (T, error)
```





**参数**:

- `ctx`: context.Context



**返回**: `T`



---






#### Method: Take


```go
func (interface{}) Take(ctx context.Context) (T, error)
```





**参数**:

- `ctx`: context.Context



**返回**: `T`



---






#### Method: Find


```go
func (interface{}) Find(ctx context.Context) ([]T, error)
```





**参数**:

- `ctx`: context.Context



**返回**: `[]T`



---






#### Method: FindInBatches


```go
func (interface{}) FindInBatches(ctx context.Context, batchSize int, fc func(...)) (error)
```





**参数**:

- `ctx`: context.Context

- `batchSize`: int

- `fc`: func(...)



**返回**: `error`


**调用**: fc() 


---






#### Method: Row


```go
func (interface{}) Row(ctx context.Context) (*sql.Row)
```





**参数**:

- `ctx`: context.Context



**返回**: `*sql.Row`



---






#### Method: Rows


```go
func (interface{}) Rows(ctx context.Context) (*sql.Rows, error)
```





**参数**:

- `ctx`: context.Context



**返回**: `*sql.Rows`



---














#### Method: Update


```go
func (interface{}) Update(ctx context.Context) (int, error)
```





**参数**:

- `ctx`: context.Context



**返回**: `int`


**调用**: s.executeAssociationOperation() len() clause.Set() int() 


---






#### Method: Create


```go
func (interface{}) Create(ctx context.Context) (error)
```





**参数**:

- `ctx`: context.Context



**返回**: `error`


**调用**: s.executeAssociationOperation() len() make() len() 


---




































#### Struct: Config








---




#### Method: Apply


```go
func (*Config) Apply(config *Config) (error)
```


Apply update config to new config



```
Apply update config to new config

```



**参数**:

- `config`: *Config



**返回**: `error`



---




#### Method: AfterInitialize


```go
func (*Config) AfterInitialize(db *DB) (error)
```


AfterInitialize initialize plugins after db connected



```
AfterInitialize initialize plugins after db connected

```



**参数**:

- `db`: *DB



**返回**: `error`


**调用**: plugin.Initialize() 


---




#### Interface: Option








---




#### Struct: DB








---




#### Struct: Session








---




#### Function: Open


```go
func Open(dialector Dialector, opts ...Option) (*DB, error)
```


Open initialize db session based on dialector



```
Open initialize db session based on dialector

```



**参数**:

- `dialector`: Dialector

- `opts`: ...Option



**返回**: `*DB`


**调用**: sort.Slice() len() append() opt.Apply() opt.AfterInitialize() d.Apply() time.Now() initializeCallbacks() db.DB() db.Close() context.Background() NewPreparedStmtDB() context.Background() pinger.Ping() context.Background() 


---






#### Method: Session


```go
func (*DB) Session(config *Session) (*DB)
```


Session create new db session



```
Session create new db session

```



**参数**:

- `config`: *Session



**返回**: `*DB`


**调用**: NewPreparedStmtDB() tx.getInstance() 


---










#### Method: WithContext


```go
func (*DB) WithContext(ctx context.Context) (*DB)
```


WithContext change current instance db's context to ctx



```
WithContext change current instance db's context to ctx

```



**参数**:

- `ctx`: context.Context



**返回**: `*DB`


**调用**: db.Session() 


---




#### Method: Debug


```go
func (*DB) Debug() (*DB)
```


Debug start debug mode



```
Debug start debug mode

```




**返回**: `*DB`


**调用**: db.getInstance() tx.Session() 


---




#### Method: Set


```go
func (*DB) Set(key string, value interface{}) (*DB)
```


Set store value with key into current db instance's context



```
Set store value with key into current db instance's context

```



**参数**:

- `key`: string

- `value`: interface{}



**返回**: `*DB`


**调用**: db.getInstance() 


---




#### Method: Get


```go
func (*DB) Get(key string) (interface{}, bool)
```


Get get value with key from current db instance's context



```
Get get value with key from current db instance's context

```



**参数**:

- `key`: string



**返回**: `interface{}`



---




#### Method: InstanceSet


```go
func (*DB) InstanceSet(key string, value interface{}) (*DB)
```


InstanceSet store value with key into current db instance's context



```
InstanceSet store value with key into current db instance's context

```



**参数**:

- `key`: string

- `value`: interface{}



**返回**: `*DB`


**调用**: db.getInstance() fmt.Sprintf() 


---




#### Method: InstanceGet


```go
func (*DB) InstanceGet(key string) (interface{}, bool)
```


InstanceGet get value with key from current db instance's context



```
InstanceGet get value with key from current db instance's context

```



**参数**:

- `key`: string



**返回**: `interface{}`


**调用**: fmt.Sprintf() 


---




#### Method: Callback


```go
func (*DB) Callback() (*callbacks)
```


Callback returns callback manager



```
Callback returns callback manager

```




**返回**: `*callbacks`



---




#### Method: AddError


```go
func (*DB) AddError(err error) (error)
```


AddError add error to db



```
AddError add error to db

```



**参数**:

- `err`: error



**返回**: `error`


**调用**: errTranslator.Translate() fmt.Errorf() 


---




#### Method: DB


```go
func (*DB) DB() (*sql.DB, error)
```


DB returns `*sql.DB`



```
DB returns `*sql.DB`

```




**返回**: `*sql.DB`


**调用**: reflect.ValueOf() dbConnector.GetDBConn() 


---






#### Function: Expr


```go
func Expr(expr string, args ...interface{}) (clause.Expr)
```


Expr returns clause.Expr, which can be used to pass SQL expression as params



```
Expr returns clause.Expr, which can be used to pass SQL expression as params

```



**参数**:

- `expr`: string

- `args`: ...interface{}



**返回**: `clause.Expr`



---




#### Method: SetupJoinTable


```go
func (*DB) SetupJoinTable(model interface{}, field string, joinTable interface{}) (error)
```


SetupJoinTable setup join table schema



```
SetupJoinTable setup join table schema

```



**参数**:

- `model`: interface{}

- `field`: string

- `joinTable`: interface{}



**返回**: `error`


**调用**: db.getInstance() stmt.Parse() stmt.Parse() fmt.Errorf() joinSchema.LookUpField() fmt.Errorf() 


---












#### Method: Use


```go
func (*DB) Use(plugin Plugin) (error)
```


Use use plugin



```
Use use plugin

```



**参数**:

- `plugin`: Plugin



**返回**: `error`


**调用**: plugin.Name() plugin.Initialize() 


---




#### Method: ToSQL


```go
func (*DB) ToSQL(queryFn func(...)) (string)
```


ToSQL for generate SQL string.



```
ToSQL for generate SQL string.

	db.ToSQL(func(tx *gorm.DB) *gorm.DB {
			return tx.Model(&User{}).Where(&User{Name: "foo", Age: 20})
				.Limit(10).Offset(5)
				.Order("name ASC")
				.First(&User{})
	})

```



**参数**:

- `queryFn`: func(...)



**返回**: `string`


**调用**: queryFn() db.Session() 


---




#### Interface: Dialector








---




#### Interface: Plugin








---




#### Interface: ParamsFilter








---




#### Interface: ConnPool








---




#### Interface: SavePointerDialectorInterface








---




#### Interface: TxBeginner








---




#### Interface: ConnPoolBeginner








---




#### Interface: TxCommitter








---




#### Interface: Tx








---




#### Interface: Valuer








---




#### Interface: GetDBConnector








---




#### Interface: Rows








---




#### Interface: ErrorTranslator








---




#### Class: EvictCallback








---




#### Struct: LRU








---










#### Function: NewLRU


```go
func NewLRU(size int, onEvict interface{}, ttl time.Duration) (*interface{})
```


NewLRU returns a new thread-safe cache with expirable entries.



```
NewLRU returns a new thread-safe cache with expirable entries.

Size parameter set to 0 makes cache of unlimited size, e.g. turns LRU mechanism off.

Providing 0 TTL turns expiring off.

Delete expired entries every 1/100th of ttl value. Goroutine which deletes expired entries runs indefinitely.

```



**参数**:

- `size`: int

- `onEvict`: interface{}

- `ttl`: time.Duration



**返回**: `*interface{}`


**调用**: make() make() make() make() time.NewTicker() ticker.Stop() res.deleteExpired() 


---




#### Method: Purge


```go
func (*interface{}) Purge()
```


Purge clears the cache completely.



```
Purge clears the cache completely.
onEvict is called for each evicted key.

```





**调用**: c.onEvict() delete() delete() 


---




#### Method: Add


```go
func (*interface{}) Add(key K, value V) (bool)
```


Add adds a value to the cache. Returns true if an eviction occurred.



```
Add adds a value to the cache. Returns true if an eviction occurred.
Returns false if there was no eviction: the item was already in the cache,
or the size was not exceeded.

```



**参数**:

- `key`: K

- `value`: V



**返回**: `bool`


**调用**: time.Now() c.removeFromBucket() now.Add() c.addToBucket() now.Add() c.addToBucket() c.removeOldest() 


---




#### Method: Get


```go
func (*interface{}) Get(key K) (V, bool)
```


Get looks up a key's value from the cache.



```
Get looks up a key's value from the cache.

```



**参数**:

- `key`: K



**返回**: `V`


**调用**: time.Now() 


---






#### Method: Contains


```go
func (*interface{}) Contains(key K) (bool)
```


Contains checks if a key is in the cache, without updating the recent-ness



```
Contains checks if a key is in the cache, without updating the recent-ness
or deleting it for being stale.

```



**参数**:

- `key`: K



**返回**: `bool`



---




#### Method: Peek


```go
func (*interface{}) Peek(key K) (V, bool)
```


Peek returns the key value (or undefined if not found) without updating



```
Peek returns the key value (or undefined if not found) without updating
the "recently used"-ness of the key.

```



**参数**:

- `key`: K



**返回**: `V`


**调用**: time.Now() 


---






#### Method: Remove


```go
func (*interface{}) Remove(key K) (bool)
```


Remove removes the provided key from the cache, returning if the



```
Remove removes the provided key from the cache, returning if the
key was contained.

```



**参数**:

- `key`: K



**返回**: `bool`


**调用**: c.removeElement() 


---




#### Method: RemoveOldest


```go
func (*interface{}) RemoveOldest() (K, V, bool)
```


RemoveOldest removes the oldest item from the cache.



```
RemoveOldest removes the oldest item from the cache.

```




**返回**: `K`


**调用**: c.removeElement() 


---




#### Method: GetOldest


```go
func (*interface{}) GetOldest() (K, V, bool)
```


GetOldest returns the oldest entry



```
GetOldest returns the oldest entry

```




**返回**: `K`



---




#### Method: KeyValues


```go
func (*interface{}) KeyValues() (map[K]V)
```






**返回**: `map[K]V`


**调用**: make() time.Now() ent.PrevEntry() now.After() 


---




#### Method: Keys


```go
func (*interface{}) Keys() ([]K)
```


Keys returns a slice of the keys in the cache, from oldest to newest.



```
Keys returns a slice of the keys in the cache, from oldest to newest.
Expired entries are filtered out.

```




**返回**: `[]K`


**调用**: make() len() time.Now() ent.PrevEntry() now.After() append() 


---




#### Method: Values


```go
func (*interface{}) Values() ([]V)
```


Values returns a slice of the values in the cache, from oldest to newest.



```
Values returns a slice of the values in the cache, from oldest to newest.
Expired entries are filtered out.

```




**返回**: `[]V`


**调用**: make() len() time.Now() ent.PrevEntry() now.After() append() 


---




#### Method: Len


```go
func (*interface{}) Len() (int)
```


Len returns the number of items in the cache.



```
Len returns the number of items in the cache.

```




**返回**: `int`



---




#### Method: Resize


```go
func (*interface{}) Resize(size int) (int)
```


Resize changes the cache size. Size of 0 means unlimited.



```
Resize changes the cache size. Size of 0 means unlimited.

```



**参数**:

- `size`: int



**返回**: `int`


**调用**: c.removeOldest() 


---














#### Method: Cap


```go
func (*interface{}) Cap() (int)
```


Cap returns the capacity of the cache



```
Cap returns the capacity of the cache

```




**返回**: `int`



---




#### Struct: Entry








---




#### Method: PrevEntry


```go
func (*interface{}) PrevEntry() (*interface{})
```


PrevEntry returns the previous list element or nil.



```
PrevEntry returns the previous list element or nil.

```




**返回**: `*interface{}`



---




#### Struct: LruList








---




#### Method: Init


```go
func (*interface{}) Init() (*interface{})
```


Init initializes or clears list l.



```
Init initializes or clears list l.

```




**返回**: `*interface{}`



---




#### Function: NewList


```go
func NewList() (*interface{})
```


NewList returns an initialized list.



```
NewList returns an initialized list.

```




**返回**: `*interface{}`


**调用**: new() 


---




#### Method: Length


```go
func (*interface{}) Length() (int)
```


Length returns the number of elements of list l.



```
Length returns the number of elements of list l.
The complexity is O(1).

```




**返回**: `int`



---




#### Method: Back


```go
func (*interface{}) Back() (*interface{})
```


Back returns the last element of list l or nil if the list is empty.



```
Back returns the last element of list l or nil if the list is empty.

```




**返回**: `*interface{}`



---










#### Method: Remove


```go
func (*interface{}) Remove(e *interface{}) (V)
```


Remove removes e from its list, decrements l.len



```
Remove removes e from its list, decrements l.len

```



**参数**:

- `e`: *interface{}



**返回**: `V`



---






#### Method: PushFront


```go
func (*interface{}) PushFront(k K, v V) (*interface{})
```


PushFront inserts a new element e with value v at the front of list l and returns e.



```
PushFront inserts a new element e with value v at the front of list l and returns e.

```



**参数**:

- `k`: K

- `v`: V



**返回**: `*interface{}`


**调用**: l.lazyInit() l.insertValue() 


---




#### Method: PushFrontExpirable


```go
func (*interface{}) PushFrontExpirable(k K, v V, expiresAt time.Time) (*interface{})
```


PushFrontExpirable inserts a new expirable element e with Value v at the front of list l and returns e.



```
PushFrontExpirable inserts a new expirable element e with Value v at the front of list l and returns e.

```



**参数**:

- `k`: K

- `v`: V

- `expiresAt`: time.Time



**返回**: `*interface{}`


**调用**: l.lazyInit() l.insertValue() 


---




#### Method: MoveToFront


```go
func (*interface{}) MoveToFront(e *interface{})
```


MoveToFront moves element e to the front of list l.



```
MoveToFront moves element e to the front of list l.
If e is not an element of l, the list is not modified.
The element must not be nil.

```



**参数**:

- `e`: *interface{}




**调用**: l.move() 


---




#### Struct: Stmt








---




#### Method: Error


```go
func (*Stmt) Error() (error)
```






**返回**: `error`



---




#### Method: Close


```go
func (*Stmt) Close() (error)
```






**返回**: `error`



---




#### Interface: Store








---








#### Function: New


```go
func New(size int, ttl time.Duration) (Store)
```


New creates and returns a new Store instance.



```
New creates and returns a new Store instance.

Parameters:
  - size: The maximum capacity of the cache. If the provided size is less than or equal to 0,
    it defaults to defaultMaxSize.
  - ttl: The time-to-live duration for each cache entry. If the provided ttl is less than or equal to 0,
    it defaults to defaultTTL.

This function defines an onEvicted callback that is invoked when a cache entry is evicted.
The callback ensures that if the evicted value (v) is not nil, its Close method is called asynchronously
to release associated resources.

Returns:
  - A Store instance implemented by lruStore, which internally uses an LRU cache with the specified size,
    eviction callback, and TTL.

```



**参数**:

- `size`: int

- `ttl`: time.Duration



**返回**: `Store`


**调用**: v.Close() 


---






#### Method: Keys


```go
func (*lruStore) Keys() ([]string)
```






**返回**: `[]string`



---




#### Method: Get


```go
func (*lruStore) Get(key string) (*Stmt, bool)
```





**参数**:

- `key`: string



**返回**: `*Stmt`



---




#### Method: Set


```go
func (*lruStore) Set(key string, value *Stmt)
```





**参数**:

- `key`: string

- `value`: *Stmt





---




#### Method: Delete


```go
func (*lruStore) Delete(key string)
```





**参数**:

- `key`: string





---




#### Interface: ConnPool








---




#### Method: New


```go
func (*lruStore) New(ctx context.Context, key string, isTransaction bool, conn ConnPool, locker sync.Locker) (*Stmt, error)
```


New creates a new Stmt object for executing SQL queries.



```
New creates a new Stmt object for executing SQL queries.
It caches the Stmt object for future use and handles preparation and error states.
Parameters:

	ctx: Context for the request, used to carry deadlines, cancellation signals, etc.
	key: The key representing the SQL query, used for caching and preparing the statement.
	isTransaction: Indicates whether this operation is part of a transaction, affecting cache strategy.
	conn: A connection pool that provides database connections.
	locker: A synchronization lock that is unlocked after initialization to avoid deadlocks.

Returns:

	*Stmt: A newly created statement object for executing SQL operations.
	error: An error if the statement preparation fails.

```



**参数**:

- `ctx`: context.Context

- `key`: string

- `isTransaction`: bool

- `conn`: ConnPool

- `locker`: sync.Locker



**返回**: `*Stmt`


**调用**: make() s.Set() locker.Unlock() close() conn.PrepareContext() s.Delete() 


---




#### Variable: ErrRecordNotFound








---




#### Constant: Reset








---




#### Constant: Red








---




#### Constant: Green








---




#### Constant: Yellow








---




#### Constant: Blue








---




#### Constant: Magenta








---




#### Constant: Cyan








---




#### Constant: White








---




#### Constant: BlueBold








---




#### Constant: MagentaBold








---




#### Constant: RedBold








---




#### Constant: YellowBold








---




#### Class: LogLevel








---




#### Constant: Silent


Silent silent log level



```
Silent silent log level

```






---




#### Constant: Error


Error error log level



```
Error error log level

```






---




#### Constant: Warn


Warn warn log level



```
Warn warn log level

```






---




#### Constant: Info


Info info log level



```
Info info log level

```






---




#### Interface: Writer








---




#### Struct: Config








---




#### Interface: Interface








---




#### Variable: Discard


Discard logger will print any log to io.Discard



```
Discard logger will print any log to io.Discard

```






---




#### Variable: Default


Default Default logger



```
Default Default logger

```






---




#### Variable: Recorder


Recorder logger records running SQL into a recorder instance



```
Recorder logger records running SQL into a recorder instance

```






---




#### Variable: RecorderParamsFilter


RecorderParamsFilter defaults to no-op, allows to be run-over by a different implementation



```
RecorderParamsFilter defaults to no-op, allows to be run-over by a different implementation

```






---




#### Function: New


```go
func New(writer Writer, config Config) (Interface)
```


New initialize logger



```
New initialize logger

```



**参数**:

- `writer`: Writer

- `config`: Config



**返回**: `Interface`



---


















#### Method: LogMode


```go
func (*logger) LogMode(level LogLevel) (Interface)
```


LogMode log mode



```
LogMode log mode

```



**参数**:

- `level`: LogLevel



**返回**: `Interface`



---




#### Method: Info


```go
func (*logger) Info(ctx context.Context, msg string, data ...interface{})
```


Info print info



```
Info print info

```



**参数**:

- `ctx`: context.Context

- `msg`: string

- `data`: ...interface{}




**调用**: l.Printf() append() utils.FileWithLineNum() 


---




#### Method: Warn


```go
func (*logger) Warn(ctx context.Context, msg string, data ...interface{})
```


Warn print warn messages



```
Warn print warn messages

```



**参数**:

- `ctx`: context.Context

- `msg`: string

- `data`: ...interface{}




**调用**: l.Printf() append() utils.FileWithLineNum() 


---




#### Method: Error


```go
func (*logger) Error(ctx context.Context, msg string, data ...interface{})
```


Error print error messages



```
Error print error messages

```



**参数**:

- `ctx`: context.Context

- `msg`: string

- `data`: ...interface{}




**调用**: l.Printf() append() utils.FileWithLineNum() 


---




#### Method: Trace


```go
func (*logger) Trace(ctx context.Context, begin time.Time, fc func(...), err error)
```


Trace print sql message



```
Trace print sql message

```



**参数**:

- `ctx`: context.Context

- `begin`: time.Time

- `fc`: func(...)

- `err`: error




**调用**: time.Since() errors.Is() fc() l.Printf() utils.FileWithLineNum() float64() elapsed.Nanoseconds() l.Printf() utils.FileWithLineNum() float64() elapsed.Nanoseconds() fc() fmt.Sprintf() l.Printf() utils.FileWithLineNum() float64() elapsed.Nanoseconds() l.Printf() utils.FileWithLineNum() float64() elapsed.Nanoseconds() fc() l.Printf() utils.FileWithLineNum() float64() elapsed.Nanoseconds() l.Printf() utils.FileWithLineNum() float64() elapsed.Nanoseconds() 


---




#### Method: ParamsFilter


```go
func (*logger) ParamsFilter(ctx context.Context, sql string, params ...interface{}) (string, []interface{})
```


ParamsFilter filter params



```
ParamsFilter filter params

```



**参数**:

- `ctx`: context.Context

- `sql`: string

- `params`: ...interface{}



**返回**: `string`



---






#### Method: New


```go
func (*traceRecorder) New() (*traceRecorder)
```


New trace recorder



```
New trace recorder

```




**返回**: `*traceRecorder`


**调用**: time.Now() 


---




#### Method: Trace


```go
func (*traceRecorder) Trace(ctx context.Context, begin time.Time, fc func(...), err error)
```


Trace implement logger interface



```
Trace implement logger interface

```



**参数**:

- `ctx`: context.Context

- `begin`: time.Time

- `fc`: func(...)

- `err`: error




**调用**: fc() 


---




#### Method: ParamsFilter


```go
func (*traceRecorder) ParamsFilter(ctx context.Context, sql string, params ...interface{}) (string, []interface{})
```





**参数**:

- `ctx`: context.Context

- `sql`: string

- `params`: ...interface{}



**返回**: `string`


**调用**: RecorderParamsFilter() 


---






#### Function: NewSlogLogger


```go
func NewSlogLogger(logger *slog.Logger, config Config) (Interface)
```





**参数**:

- `logger`: *slog.Logger

- `config`: Config



**返回**: `Interface`



---




#### Method: LogMode


```go
func (*slogLogger) LogMode(level LogLevel) (Interface)
```





**参数**:

- `level`: LogLevel



**返回**: `Interface`



---




#### Method: Info


```go
func (*slogLogger) Info(ctx context.Context, msg string, data ...interface{})
```





**参数**:

- `ctx`: context.Context

- `msg`: string

- `data`: ...interface{}




**调用**: l.log() slog.Any() 


---




#### Method: Warn


```go
func (*slogLogger) Warn(ctx context.Context, msg string, data ...interface{})
```





**参数**:

- `ctx`: context.Context

- `msg`: string

- `data`: ...interface{}




**调用**: l.log() slog.Any() 


---




#### Method: Error


```go
func (*slogLogger) Error(ctx context.Context, msg string, data ...interface{})
```





**参数**:

- `ctx`: context.Context

- `msg`: string

- `data`: ...interface{}




**调用**: l.log() slog.Any() 


---




#### Method: Trace


```go
func (*slogLogger) Trace(ctx context.Context, begin time.Time, fc func(...), err error)
```





**参数**:

- `ctx`: context.Context

- `begin`: time.Time

- `fc`: func(...)

- `err`: error




**调用**: time.Since() fc() slog.String() fmt.Sprintf() float64() elapsed.Nanoseconds() slog.String() append() slog.Int64() errors.Is() append() slog.String() err.Error() l.log() slog.GroupValue() l.log() slog.GroupValue() l.log() slog.GroupValue() 


---






#### Method: ParamsFilter


```go
func (*slogLogger) ParamsFilter(ctx context.Context, sql string, params ...interface{}) (string, []interface{})
```


ParamsFilter filter params



```
ParamsFilter filter params

```



**参数**:

- `ctx`: context.Context

- `sql`: string

- `params`: ...interface{}



**返回**: `string`



---














#### Function: ExplainSQL


```go
func ExplainSQL(sql string, numericPlaceholder *regexp.Regexp, escaper string, avars ...interface{}) (string)
```


ExplainSQL generate SQL string with given parameters, the generated SQL is expected to be used in logger, execute it might introduce a SQL injection vulnerability



```
ExplainSQL generate SQL string with given parameters, the generated SQL is expected to be used in logger, execute it might introduce a SQL injection vulnerability

```



**参数**:

- `sql`: string

- `numericPlaceholder`: *regexp.Regexp

- `escaper`: string

- `avars`: ...interface{}



**返回**: `string`


**调用**: make() len() strconv.FormatBool() v.IsZero() v.Format() v.IsZero() v.Format() reflect.ValueOf() reflectValue.IsValid() reflectValue.Kind() reflectValue.IsNil() reflectValue.Kind() v.Value() convertParams() reflect.ValueOf() reflectValue.Kind() fmt.Sprintf() reflectValue.Interface() fmt.Sprintf() reflectValue.Interface() fmt.Sprintf() reflectValue.Interface() strings.ReplaceAll() fmt.Sprintf() reflectValue.IsValid() reflectValue.Kind() reflectValue.IsNil() reflectValue.Kind() strings.ReplaceAll() fmt.Sprintf() string() isPrintable() strings.ReplaceAll() utils.ToString() strconv.FormatFloat() float64() strconv.FormatFloat() strings.ReplaceAll() reflect.ValueOf() rv.IsValid() rv.Kind() rv.IsNil() valuer.Value() convertParams() rv.Kind() rv.IsZero() convertParams() reflect.Indirect() isNumeric() rv.Kind() rv.CanInt() rv.CanUint() fmt.Sprintf() rv.Interface() fmt.Sprintf() rv.Interface() rv.Type() convertParams() rv.Convert() strings.ReplaceAll() fmt.Sprint() convertParams() len() newSQL.WriteString() newSQL.WriteByte() newSQL.String() numericPlaceholder.ReplaceAllString() numericPlaceholderRe.ReplaceAllStringFunc() len() strconv.Atoi() len() 


---












#### Struct: ColumnType








---




#### Method: Name


```go
func (ColumnType) Name() (string)
```


Name returns the name or alias of the column.



```
Name returns the name or alias of the column.

```




**返回**: `string`



---




#### Method: DatabaseTypeName


```go
func (ColumnType) DatabaseTypeName() (string)
```


DatabaseTypeName returns the database system name of the column type. If an empty



```
DatabaseTypeName returns the database system name of the column type. If an empty
string is returned, then the driver type name is not supported.
Consult your driver documentation for a list of driver data types. Length specifiers
are not included.
Common type names include "VARCHAR", "TEXT", "NVARCHAR", "DECIMAL", "BOOL",
"INT", and "BIGINT".

```




**返回**: `string`



---




#### Method: ColumnType


```go
func (ColumnType) ColumnType() (string, bool)
```


ColumnType returns the database type of the column. like `varchar(16)`



```
ColumnType returns the database type of the column. like `varchar(16)`

```




**返回**: `string`



---




#### Method: PrimaryKey


```go
func (ColumnType) PrimaryKey() (bool, bool)
```


PrimaryKey returns the column is primary key or not.



```
PrimaryKey returns the column is primary key or not.

```




**返回**: `bool`



---




#### Method: AutoIncrement


```go
func (ColumnType) AutoIncrement() (bool, bool)
```


AutoIncrement returns the column is auto increment or not.



```
AutoIncrement returns the column is auto increment or not.

```




**返回**: `bool`



---




#### Method: Length


```go
func (ColumnType) Length() (int64, bool)
```


Length returns the column type length for variable length column types



```
Length returns the column type length for variable length column types

```




**返回**: `int64`



---




#### Method: DecimalSize


```go
func (ColumnType) DecimalSize() (int64, int64, bool)
```


DecimalSize returns the scale and precision of a decimal type.



```
DecimalSize returns the scale and precision of a decimal type.

```




**返回**: `int64`



---




#### Method: Nullable


```go
func (ColumnType) Nullable() (bool, bool)
```


Nullable reports whether the column may be null.



```
Nullable reports whether the column may be null.

```




**返回**: `bool`



---




#### Method: Unique


```go
func (ColumnType) Unique() (bool, bool)
```


Unique reports whether the column may be unique.



```
Unique reports whether the column may be unique.

```




**返回**: `bool`



---




#### Method: ScanType


```go
func (ColumnType) ScanType() (reflect.Type)
```


ScanType returns a Go type suitable for scanning into using Rows.Scan.



```
ScanType returns a Go type suitable for scanning into using Rows.Scan.

```




**返回**: `reflect.Type`



---




#### Method: Comment


```go
func (ColumnType) Comment() (string, bool)
```


Comment returns the comment of current column.



```
Comment returns the comment of current column.

```




**返回**: `string`



---




#### Method: DefaultValue


```go
func (ColumnType) DefaultValue() (string, bool)
```


DefaultValue returns the default value of current column.



```
DefaultValue returns the default value of current column.

```




**返回**: `string`



---




#### Struct: Index








---




#### Method: Table


```go
func (Index) Table() (string)
```


Table return the table name of the index.



```
Table return the table name of the index.

```




**返回**: `string`



---




#### Method: Name


```go
func (Index) Name() (string)
```


Name return the name of the index.



```
Name return the name of the index.

```




**返回**: `string`



---




#### Method: Columns


```go
func (Index) Columns() ([]string)
```


Columns return the columns of the index



```
Columns return the columns of the index

```




**返回**: `[]string`



---




#### Method: PrimaryKey


```go
func (Index) PrimaryKey() (bool, bool)
```


PrimaryKey returns the index is primary key or not.



```
PrimaryKey returns the index is primary key or not.

```




**返回**: `bool`



---




#### Method: Unique


```go
func (Index) Unique() (bool, bool)
```


Unique returns whether the index is unique or not.



```
Unique returns whether the index is unique or not.

```




**返回**: `bool`



---




#### Method: Option


```go
func (Index) Option() (string)
```


Option return the optional attribute of the index



```
Option return the optional attribute of the index

```




**返回**: `string`



---








#### Struct: Migrator








---




#### Struct: Config








---






#### Method: Trace


```go
func (*printSQLLogger) Trace(ctx context.Context, begin time.Time, fc func(...), err error)
```





**参数**:

- `ctx`: context.Context

- `begin`: time.Time

- `fc`: func(...)

- `err`: error




**调用**: fc() fmt.Println() 


---




#### Interface: GormDataTypeInterface








---




#### Method: RunWithValue


```go
func (Migrator) RunWithValue(value interface{}, fc func(...)) (error)
```


RunWithValue run migration with statement value



```
RunWithValue run migration with statement value

```



**参数**:

- `value`: interface{}

- `fc`: func(...)



**返回**: `error`


**调用**: stmt.ParseWithSpecialTableName() fc() 


---




#### Method: DataTypeOf


```go
func (Migrator) DataTypeOf(field *schema.Field) (string)
```


DataTypeOf return field's db data type



```
DataTypeOf return field's db data type

```



**参数**:

- `field`: *schema.Field



**返回**: `string`


**调用**: reflect.New() fieldValue.Interface() dataTyper.GormDBDataType() 


---




#### Method: FullDataTypeOf


```go
func (Migrator) FullDataTypeOf(field *schema.Field) (clause.Expr)
```


FullDataTypeOf returns field's db full data type



```
FullDataTypeOf returns field's db full data type

```



**参数**:

- `field`: *schema.Field



**返回**: `clause.Expr`


**调用**: m.DataTypeOf() 


---




#### Method: GetQueryAndExecTx


```go
func (Migrator) GetQueryAndExecTx() (*gorm.DB)
```






**返回**: `*gorm.DB`



---




#### Method: AutoMigrate


```go
func (Migrator) AutoMigrate(values ...interface{}) (error)
```


AutoMigrate auto migrate values



```
AutoMigrate auto migrate values

```



**参数**:

- `values`: ...interface{}



**返回**: `error`


**调用**: m.ReorderModels() m.GetQueryAndExecTx() queryTx.Migrator() execTx.Migrator() m.RunWithValue() errors.New() queryTx.Migrator() columnType.Name() execTx.Migrator() execTx.Migrator() rel.ParseConstraint() queryTx.Migrator() execTx.Migrator() queryTx.Migrator() execTx.Migrator() queryTx.Migrator() execTx.Migrator() 


---










#### Method: GetTables


```go
func (Migrator) GetTables() ([]string, error)
```


GetTables returns tables



```
GetTables returns tables

```




**返回**: `[]string`


**调用**: m.CurrentDatabase() 


---




#### Method: CreateTable


```go
func (Migrator) CreateTable(values ...interface{}) (error)
```


CreateTable create table in database for values



```
CreateTable create table in database for values

```



**参数**:

- `values`: ...interface{}



**返回**: `error`


**调用**: m.ReorderModels() m.RunWithValue() errors.New() m.CurrentTable() strings.Contains() strings.ToUpper() m.DataTypeOf() append() len() make() len() append() append() tx.Migrator() fmt.Sprintf() append() tx.Migrator() rel.ParseConstraint() constraint.Build() append() append() stmt.Quote() append() strings.TrimSuffix() fmt.Sprint() tx.Exec() 


---










#### Method: DropTable


```go
func (Migrator) DropTable(values ...interface{}) (error)
```


DropTable drop table for values



```
DropTable drop table for values

```



**参数**:

- `values`: ...interface{}



**返回**: `error`


**调用**: m.ReorderModels() len() m.RunWithValue() tx.Exec() m.CurrentTable() 


---




#### Method: HasTable


```go
func (Migrator) HasTable(value interface{}) (bool)
```


HasTable returns table exists or not for value, value could be a struct or string



```
HasTable returns table exists or not for value, value could be a struct or string

```



**参数**:

- `value`: interface{}



**返回**: `bool`


**调用**: m.RunWithValue() 


---






#### Method: RenameTable


```go
func (Migrator) RenameTable(oldName, newName interface{}) (error)
```


RenameTable rename table from oldName to newName



```
RenameTable rename table from oldName to newName

```



**参数**:

- `oldName`: interface{}

- `newName`: interface{}



**返回**: `error`


**调用**: stmt.Parse() m.CurrentTable() stmt.Parse() m.CurrentTable() 


---








#### Method: AddColumn


```go
func (Migrator) AddColumn(value interface{}, name string) (error)
```


AddColumn create `name` column for value



```
AddColumn create `name` column for value

```



**参数**:

- `value`: interface{}

- `name`: string



**返回**: `error`


**调用**: m.RunWithValue() errors.New() fmt.Errorf() m.CurrentTable() 


---




#### Method: DropColumn


```go
func (Migrator) DropColumn(value interface{}, name string) (error)
```


DropColumn drop value's `name` column



```
DropColumn drop value's `name` column

```



**参数**:

- `value`: interface{}

- `name`: string



**返回**: `error`


**调用**: m.RunWithValue() m.CurrentTable() 


---




#### Method: AlterColumn


```go
func (Migrator) AlterColumn(value interface{}, field string) (error)
```


AlterColumn alter value's `field` column' type based on schema definition



```
AlterColumn alter value's `field` column' type based on schema definition

```



**参数**:

- `value`: interface{}

- `field`: string



**返回**: `error`


**调用**: m.RunWithValue() m.FullDataTypeOf() m.CurrentTable() fmt.Errorf() 


---




#### Method: HasColumn


```go
func (Migrator) HasColumn(value interface{}, field string) (bool)
```


HasColumn check has column `field` for value or not



```
HasColumn check has column `field` for value or not

```



**参数**:

- `value`: interface{}

- `field`: string



**返回**: `bool`


**调用**: m.RunWithValue() 


---






#### Method: RenameColumn


```go
func (Migrator) RenameColumn(value interface{}, oldName, newName string) (error)
```


RenameColumn rename value's field name from oldName to newName



```
RenameColumn rename value's field name from oldName to newName

```



**参数**:

- `value`: interface{}

- `oldName`: string

- `newName`: string



**返回**: `error`


**调用**: m.RunWithValue() m.CurrentTable() 


---




#### Method: MigrateColumn


```go
func (Migrator) MigrateColumn(value interface{}, field *schema.Field, columnType gorm.ColumnType) (error)
```


MigrateColumn migrate column



```
MigrateColumn migrate column

```



**参数**:

- `value`: interface{}

- `field`: *schema.Field

- `columnType`: gorm.ColumnType



**返回**: `error`


**调用**: strings.TrimSpace() strings.ToLower() strings.ToLower() columnType.DatabaseTypeName() strings.HasPrefix() strings.HasPrefix() columnType.Length() int64() regFullDataType.FindAllStringSubmatch() len() fmt.Sprint() regexp.MustCompile() columnType.DecimalSize() strings.HasPrefix() fmt.Sprintf() strings.HasPrefix() fmt.Sprintf() columnType.DecimalSize() int64() regexp.MustCompile() fmt.Sprintf() m.DataTypeOf() columnType.Nullable() strings.EqualFold() columnType.DefaultValue() strings.EqualFold() strings.TrimSuffix() strings.TrimSuffix() strconv.ParseBool() strconv.ParseBool() strings.Trim() columnType.Comment() 


---








#### Method: MigrateColumnUnique


```go
func (Migrator) MigrateColumnUnique(value interface{}, field *schema.Field, columnType gorm.ColumnType) (error)
```





**参数**:

- `value`: interface{}

- `field`: *schema.Field

- `columnType`: gorm.ColumnType



**返回**: `error`


**调用**: columnType.Unique() m.RunWithValue() 


---




#### Method: ColumnTypes


```go
func (Migrator) ColumnTypes(value interface{}) ([]gorm.ColumnType, error)
```


ColumnTypes return columnTypes []gorm.ColumnType and execErr error



```
ColumnTypes return columnTypes []gorm.ColumnType and execErr error

```



**参数**:

- `value`: interface{}



**返回**: `[]gorm.ColumnType`


**调用**: make() m.RunWithValue() rows.Close() rows.ColumnTypes() append() 


---






#### Method: CreateView


```go
func (Migrator) CreateView(name string, option gorm.ViewOption) (error)
```


CreateView create view from Query in gorm.ViewOption.



```
CreateView create view from Query in gorm.ViewOption.
Query in gorm.ViewOption is a [subquery]

	// CREATE VIEW `user_view` AS SELECT * FROM `users` WHERE age > 20
	q := DB.Model(&User{}).Where("age > ?", 20)
	DB.Debug().Migrator().CreateView("user_view", gorm.ViewOption{Query: q})

	// CREATE OR REPLACE VIEW `users_view` AS SELECT * FROM `users` WITH CHECK OPTION
	q := DB.Model(&User{})
	DB.Debug().Migrator().CreateView("user_view", gorm.ViewOption{Query: q, Replace: true, CheckOption: "WITH CHECK OPTION"})

[subquery]: https://gorm.io/docs/advanced_query.html#SubQuery

```



**参数**:

- `name`: string

- `option`: gorm.ViewOption



**返回**: `error`


**调用**: new() sql.WriteString() sql.WriteString() sql.WriteString() m.QuoteTo() sql.WriteString() sql.WriteString() sql.WriteString() m.Explain() sql.String() 


---




#### Method: DropView


```go
func (Migrator) DropView(name string) (error)
```


DropView drop view



```
DropView drop view

```



**参数**:

- `name`: string



**返回**: `error`



---




#### Method: GuessConstraintAndTable


```go
func (Migrator) GuessConstraintAndTable(stmt *gorm.Statement, name string) (*schema.Constraint, *schema.CheckConstraint, string)
```


GuessConstraintAndTable guess statement's constraint and it's table based on name



```
GuessConstraintAndTable guess statement's constraint and it's table based on name

Deprecated: use GuessConstraintInterfaceAndTable instead.

```



**参数**:

- `stmt`: *gorm.Statement

- `name`: string



**返回**: `*schema.Constraint`


**调用**: m.GuessConstraintInterfaceAndTable() 


---




#### Method: GuessConstraintInterfaceAndTable


```go
func (Migrator) GuessConstraintInterfaceAndTable(stmt *gorm.Statement, name string) (schema.ConstraintInterface, string)
```


GuessConstraintInterfaceAndTable guess statement's constraint and it's table based on name



```
GuessConstraintInterfaceAndTable guess statement's constraint and it's table based on name
nolint:cyclop

```



**参数**:

- `stmt`: *gorm.Statement

- `name`: string



**返回**: `schema.ConstraintInterface`


**调用**: rel.ParseConstraint() getTable() rel.ParseConstraint() getTable() 


---




#### Method: CreateConstraint


```go
func (Migrator) CreateConstraint(value interface{}, name string) (error)
```


CreateConstraint create constraint



```
CreateConstraint create constraint

```



**参数**:

- `value`: interface{}

- `name`: string



**返回**: `error`


**调用**: m.RunWithValue() m.GuessConstraintInterfaceAndTable() constraint.Build() append() 


---




#### Method: DropConstraint


```go
func (Migrator) DropConstraint(value interface{}, name string) (error)
```


DropConstraint drop constraint



```
DropConstraint drop constraint

```



**参数**:

- `value`: interface{}

- `name`: string



**返回**: `error`


**调用**: m.RunWithValue() m.GuessConstraintInterfaceAndTable() constraint.GetName() 


---




#### Method: HasConstraint


```go
func (Migrator) HasConstraint(value interface{}, name string) (bool)
```


HasConstraint check has constraint or not



```
HasConstraint check has constraint or not

```



**参数**:

- `value`: interface{}

- `name`: string



**返回**: `bool`


**调用**: m.RunWithValue() m.GuessConstraintInterfaceAndTable() constraint.GetName() 


---






#### Method: BuildIndexOptions


```go
func (Migrator) BuildIndexOptions(opts []schema.IndexOption, stmt *gorm.Statement) ([]interface{})
```


BuildIndexOptions build index options



```
BuildIndexOptions build index options

```



**参数**:

- `opts`: []schema.IndexOption

- `stmt`: *gorm.Statement



**返回**: `[]interface{}`


**调用**: stmt.Quote() fmt.Sprintf() append() 


---




#### Interface: BuildIndexOptionsInterface








---




#### Method: CreateIndex


```go
func (Migrator) CreateIndex(value interface{}, name string) (error)
```


CreateIndex create index `name`



```
CreateIndex create index `name`

```



**参数**:

- `value`: interface{}

- `name`: string



**返回**: `error`


**调用**: m.RunWithValue() errors.New() m.CurrentTable() fmt.Sprintf() fmt.Errorf() 


---




#### Method: DropIndex


```go
func (Migrator) DropIndex(value interface{}, name string) (error)
```


DropIndex drop index `name`



```
DropIndex drop index `name`

```



**参数**:

- `value`: interface{}

- `name`: string



**返回**: `error`


**调用**: m.RunWithValue() m.CurrentTable() 


---




#### Method: HasIndex


```go
func (Migrator) HasIndex(value interface{}, name string) (bool)
```


HasIndex check has index `name` or not



```
HasIndex check has index `name` or not

```



**参数**:

- `value`: interface{}

- `name`: string



**返回**: `bool`


**调用**: m.RunWithValue() 


---






#### Method: RenameIndex


```go
func (Migrator) RenameIndex(value interface{}, oldName, newName string) (error)
```


RenameIndex rename index from oldName to newName



```
RenameIndex rename index from oldName to newName

```



**参数**:

- `value`: interface{}

- `oldName`: string

- `newName`: string



**返回**: `error`


**调用**: m.RunWithValue() m.CurrentTable() 


---




#### Method: CurrentDatabase


```go
func (Migrator) CurrentDatabase() (string)
```


CurrentDatabase returns current database name



```
CurrentDatabase returns current database name

```




**返回**: `string`



---




#### Method: ReorderModels


```go
func (Migrator) ReorderModels(values []interface{}, autoAdd bool) ([]interface{})
```


ReorderModels reorder models according to constraint dependencies



```
ReorderModels reorder models according to constraint dependencies

```



**参数**:

- `values`: []interface{}

- `autoAdd`: bool



**返回**: `[]interface{}`


**调用**: dep.ParseWithSpecialTableName() context.Background() rel.ParseConstraint() append() append() reflect.New() parseDependence() parseDependence() reflect.New() append() insertIntoOrderedList() parseDependence() reflect.New() insertIntoOrderedList() append() append() parseDependence() insertIntoOrderedList() append() 


---




#### Struct: Dependency








---


















#### Method: CurrentTable


```go
func (Migrator) CurrentTable(stmt *gorm.Statement) (interface{})
```


CurrentTable returns current statement's table expression



```
CurrentTable returns current statement's table expression

```



**参数**:

- `stmt`: *gorm.Statement



**返回**: `interface{}`



---




#### Method: GetIndexes


```go
func (Migrator) GetIndexes(dst interface{}) ([]gorm.Index, error)
```


GetIndexes return Indexes []gorm.Index and execErr error



```
GetIndexes return Indexes []gorm.Index and execErr error

```



**参数**:

- `dst`: interface{}



**返回**: `[]gorm.Index`


**调用**: errors.New() 


---




#### Method: GetTypeAliases


```go
func (Migrator) GetTypeAliases(databaseTypeName string) ([]string)
```


GetTypeAliases return database type aliases



```
GetTypeAliases return database type aliases

```



**参数**:

- `databaseTypeName`: string



**返回**: `[]string`



---




#### Method: TableType


```go
func (Migrator) TableType(dst interface{}) (gorm.TableType, error)
```


TableType return tableType gorm.TableType and execErr error



```
TableType return tableType gorm.TableType and execErr error

```



**参数**:

- `dst`: interface{}



**返回**: `gorm.TableType`


**调用**: errors.New() 


---




#### Struct: TableType








---




#### Method: Schema


```go
func (TableType) Schema() (string)
```


Schema returns the schema of the table.



```
Schema returns the schema of the table.

```




**返回**: `string`



---




#### Method: Name


```go
func (TableType) Name() (string)
```


Name returns the name of the table.



```
Name returns the name of the table.

```




**返回**: `string`



---




#### Method: Type


```go
func (TableType) Type() (string)
```


Type returns the type of the table.



```
Type returns the type of the table.

```




**返回**: `string`



---




#### Method: Comment


```go
func (TableType) Comment() (string, bool)
```


Comment returns the comment of current table.



```
Comment returns the comment of current table.

```




**返回**: `string`



---




#### Method: Migrator


```go
func (*DB) Migrator() (Migrator)
```


Migrator returns migrator



```
Migrator returns migrator

```




**返回**: `Migrator`


**调用**: db.getInstance() len() tx.executeScopes() tx.Session() 


---




#### Method: AutoMigrate


```go
func (*DB) AutoMigrate(dst ...interface{}) (error)
```


AutoMigrate run auto migration for given models



```
AutoMigrate run auto migration for given models

```



**参数**:

- `dst`: ...interface{}



**返回**: `error`


**调用**: db.Migrator() 


---




#### Struct: ViewOption








---




#### Interface: ColumnType








---




#### Interface: Index








---




#### Interface: TableType








---




#### Interface: Migrator








---




#### Struct: Model








---




#### Struct: PreparedStmtDB








---




#### Function: NewPreparedStmtDB


```go
func NewPreparedStmtDB(connPool ConnPool, maxSize int, ttl time.Duration) (*PreparedStmtDB)
```


NewPreparedStmtDB creates and initializes a new instance of PreparedStmtDB.



```
NewPreparedStmtDB creates and initializes a new instance of PreparedStmtDB.

Parameters:
- connPool: A connection pool that implements the ConnPool interface, used for managing database connections.
- maxSize: The maximum number of prepared statements that can be stored in the statement store.
- ttl: The time-to-live duration for each prepared statement in the store. Statements older than this duration will be automatically removed.

Returns:
- A pointer to a PreparedStmtDB instance, which manages prepared statements using the provided connection pool and configuration.

```



**参数**:

- `connPool`: ConnPool

- `maxSize`: int

- `ttl`: time.Duration



**返回**: `*PreparedStmtDB`


**调用**: stmt_store.New() 


---




#### Method: GetDBConn


```go
func (*PreparedStmtDB) GetDBConn() (*sql.DB, error)
```


GetDBConn returns the underlying *sql.DB connection



```
GetDBConn returns the underlying *sql.DB connection

```




**返回**: `*sql.DB`


**调用**: dbConnector.GetDBConn() 


---




#### Method: Close


```go
func (*PreparedStmtDB) Close()
```


Close closes all prepared statements in the store



```
Close closes all prepared statements in the store

```






---




#### Method: Reset


```go
func (*PreparedStmtDB) Reset()
```


Reset Deprecated use Close instead



```
Reset Deprecated use Close instead

```





**调用**: db.Close() 


---






#### Method: BeginTx


```go
func (*PreparedStmtDB) BeginTx(ctx context.Context, opt *sql.TxOptions) (ConnPool, error)
```





**参数**:

- `ctx`: context.Context

- `opt`: *sql.TxOptions



**返回**: `ConnPool`


**调用**: beginner.BeginTx() beginner.BeginTx() 


---




#### Method: ExecContext


```go
func (*PreparedStmtDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
```





**参数**:

- `ctx`: context.Context

- `query`: string

- `args`: ...interface{}



**返回**: `sql.Result`


**调用**: db.prepare() stmt.ExecContext() errors.Is() 


---




#### Method: QueryContext


```go
func (*PreparedStmtDB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
```





**参数**:

- `ctx`: context.Context

- `query`: string

- `args`: ...interface{}



**返回**: `*sql.Rows`


**调用**: db.prepare() stmt.QueryContext() errors.Is() 


---




#### Method: QueryRowContext


```go
func (*PreparedStmtDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) (*sql.Row)
```





**参数**:

- `ctx`: context.Context

- `query`: string

- `args`: ...interface{}



**返回**: `*sql.Row`


**调用**: db.prepare() stmt.QueryRowContext() 


---




#### Method: Ping


```go
func (*PreparedStmtDB) Ping() (error)
```






**返回**: `error`


**调用**: db.GetDBConn() conn.Ping() 


---




#### Struct: PreparedStmtTX








---




#### Method: GetDBConn


```go
func (*PreparedStmtTX) GetDBConn() (*sql.DB, error)
```






**返回**: `*sql.DB`



---




#### Method: Commit


```go
func (*PreparedStmtTX) Commit() (error)
```






**返回**: `error`


**调用**: reflect.ValueOf() 


---




#### Method: Rollback


```go
func (*PreparedStmtTX) Rollback() (error)
```






**返回**: `error`


**调用**: reflect.ValueOf() 


---




#### Method: ExecContext


```go
func (*PreparedStmtTX) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
```





**参数**:

- `ctx`: context.Context

- `query`: string

- `args`: ...interface{}



**返回**: `sql.Result`


**调用**: errors.Is() 


---




#### Method: QueryContext


```go
func (*PreparedStmtTX) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
```





**参数**:

- `ctx`: context.Context

- `query`: string

- `args`: ...interface{}



**返回**: `*sql.Rows`


**调用**: errors.Is() 


---




#### Method: QueryRowContext


```go
func (*PreparedStmtTX) QueryRowContext(ctx context.Context, query string, args ...interface{}) (*sql.Row)
```





**参数**:

- `ctx`: context.Context

- `query`: string

- `args`: ...interface{}



**返回**: `*sql.Row`



---




#### Method: Ping


```go
func (*PreparedStmtTX) Ping() (error)
```






**返回**: `error`


**调用**: tx.GetDBConn() conn.Ping() 


---










#### Class: ScanMode








---




#### Constant: ScanInitialized








---




#### Constant: ScanUpdate








---




#### Constant: ScanOnConflictDoNothing








---




#### Function: Scan


```go
func Scan(rows Rows, db *DB, mode ScanMode)
```


Scan scan rows into db statement



```
Scan scan rows into db statement

```



**参数**:

- `rows`: Rows

- `db`: *DB

- `mode`: ScanMode




**调用**: rows.Columns() make() len() len() rows.Next() rows.ColumnTypes() prepareValues() db.AddError() rows.Scan() scanIntoMap() rows.ColumnTypes() rows.Next() prepareValues() db.AddError() rows.Scan() scanIntoMap() append() rows.Next() db.AddError() rows.Scan() make() len() reflectValue.Kind() reflectValue.Elem() reflectValue.Type() reflectValueType.Kind() reflectValueType.Elem() reflectValueType.Kind() reflectValueType.Elem() reflectValueType.Kind() schema.Parse() len() reflect.New() reflectValueType.Kind() make() len() sch.LookUpField() utils.SplitNestedRelationName() len() utils.JoinNestedRelationNames() len() append() strings.Split() len() len() make() append() append() len() make() len() append() reflectValue.Kind() reflectValue.Kind() reflectValue.Len() reflect.Zero() reflectValue.Type() reflectValue.Cap() reflect.MakeSlice() reflectValue.Type() reflectValue.SetLen() rows.Next() int() reflectValue.Len() reflectValue.Index() int() field.ValueOf() reflect.New() db.scanIntoStruct() elem.Elem() reflectValue.Len() int() reflectValue.Index() int() reflect.Append() rows.Next() reflectValue.Kind() reflect.Zero() reflectValue.Type() db.scanIntoStruct() db.AddError() rows.Scan() rows.Err() db.AddError() db.AddError() 


---


































#### Struct: CheckConstraint








---




#### Method: GetName


```go
func (*CheckConstraint) GetName() (string)
```






**返回**: `string`



---




#### Method: Build


```go
func (*CheckConstraint) Build() (string, []interface{})
```






**返回**: `string`



---




#### Method: ParseCheckConstraints


```go
func (*Schema) ParseCheckConstraints() (map[string]CheckConstraint)
```


ParseCheckConstraints parse schema check constraints



```
ParseCheckConstraints parse schema check constraints

```




**返回**: `map[string]CheckConstraint`


**调用**: strings.Split() len() regEnLetterAndMidline.MatchString() strings.Join() strings.Join() 


---




#### Struct: UniqueConstraint








---




#### Method: GetName


```go
func (*UniqueConstraint) GetName() (string)
```






**返回**: `string`



---




#### Method: Build


```go
func (*UniqueConstraint) Build() (string, []interface{})
```






**返回**: `string`



---




#### Method: ParseUniqueConstraints


```go
func (*Schema) ParseUniqueConstraints() (map[string]UniqueConstraint)
```


ParseUniqueConstraints parse schema unique constraints



```
ParseUniqueConstraints parse schema unique constraints

```




**返回**: `map[string]UniqueConstraint`


**调用**: make() 


---




#### Variable: TimeReflectType








---




#### Variable: TimePtrReflectType








---




#### Variable: ByteReflectType








---




#### Class: DataType


DataType GORM data type



```
DataType GORM data type

```






---




#### Class: TimeType


TimeType GORM time type



```
TimeType GORM time type

```






---




#### Constant: UnixTime








---




#### Constant: UnixSecond








---




#### Constant: UnixMillisecond








---




#### Constant: UnixNanosecond








---




#### Constant: Bool








---




#### Constant: Int








---




#### Constant: Uint








---




#### Constant: Float








---




#### Constant: String








---




#### Constant: Time








---




#### Constant: Bytes








---




#### Constant: DefaultAutoIncrementIncrement








---




#### Struct: Field








---




#### Method: BindName


```go
func (*Field) BindName() (string)
```






**返回**: `string`


**调用**: strings.Join() 


---




#### Method: ParseField


```go
func (*Schema) ParseField(fieldStruct reflect.StructField) (*Field)
```


ParseField parses reflect.StructField to Field



```
ParseField parses reflect.StructField to Field

```



**参数**:

- `fieldStruct`: reflect.StructField



**返回**: `*Field`


**调用**: ParseTagSetting() utils.CheckTruth() utils.CheckTruth() utils.CheckTruth() utils.CheckTruth() utils.CheckTruth() reflect.New() fieldValue.Interface() fieldValue.Interface() valuer.Value() reflect.ValueOf() reflect.ValueOf() reflect.Indirect() rv.Type() rv.Kind() rvType.ConvertibleTo() rvType.NumField() ParseTagSetting() rvType.Field() rvType.NumField() rvType.Field() newFieldType.Kind() newFieldType.Elem() reflect.New() reflect.Indirect() getRealFieldValue() fieldValue.IsValid() getRealFieldValue() fieldValue.Interface() GetSerializer() fmt.Errorf() strconv.ParseInt() strconv.Atoi() strconv.Atoi() strconv.Atoi() strings.TrimSpace() strings.Contains() strings.Contains() strings.ToLower() reflect.Indirect() strconv.ParseBool() fmt.Errorf() strconv.ParseInt() fmt.Errorf() strconv.ParseUint() fmt.Errorf() strconv.ParseFloat() fmt.Errorf() strings.Trim() strings.Trim() fieldValue.Interface() fieldValue.Type() fieldValue.Type() now.Parse() reflect.Indirect() fieldValue.Interface() DataType() dataTyper.GormDataType() utils.CheckTruth() strings.ToUpper() strings.ToUpper() utils.CheckTruth() strings.ToUpper() strings.ToUpper() DataType() strings.ToLower() DataType() reflect.Indirect() strings.ToLower() strings.TrimSpace() strings.ToLower() strings.Contains() strings.Contains() reflect.Indirect() cacheStore.Store() getOrParse() fieldValue.Interface() append() append() append() append() utils.CheckTruth() utils.CheckTruth() fmt.Errorf() 


---


























#### Struct: Index








---




#### Struct: IndexOption








---




#### Method: ParseIndexes


```go
func (*Schema) ParseIndexes() ([]*Index)
```


ParseIndexes parse schema indexes



```
ParseIndexes parse schema indexes

```




**返回**: `[]*Index`


**调用**: parseFieldIndexes() append() append() sort.Slice() len() 


---




#### Method: LookIndex


```go
func (*Schema) LookIndex(name string) (*Index)
```





**参数**:

- `name`: string



**返回**: `*Index`


**调用**: schema.ParseIndexes() 


---




















#### Interface: ConstraintInterface








---




#### Interface: GormDataTypeInterface








---




#### Interface: FieldNewValuePool








---




#### Interface: CreateClausesInterface








---




#### Interface: QueryClausesInterface








---




#### Interface: UpdateClausesInterface








---




#### Interface: DeleteClausesInterface








---




#### Interface: Namer








---




#### Interface: Replacer








---






#### Struct: NamingStrategy








---




#### Method: TableName


```go
func (NamingStrategy) TableName(str string) (string)
```


TableName convert string to table name



```
TableName convert string to table name

```



**参数**:

- `str`: string



**返回**: `string`


**调用**: ns.toDBName() inflection.Plural() ns.toDBName() 


---




#### Method: SchemaName


```go
func (NamingStrategy) SchemaName(table string) (string)
```


SchemaName generate schema name from table name, don't guarantee it is the reverse value of TableName



```
SchemaName generate schema name from table name, don't guarantee it is the reverse value of TableName

```



**参数**:

- `table`: string



**返回**: `string`


**调用**: strings.TrimPrefix() ns.toSchemaName() ns.toSchemaName() inflection.Singular() 


---




#### Method: ColumnName


```go
func (NamingStrategy) ColumnName(table, column string) (string)
```


ColumnName convert string to column name



```
ColumnName convert string to column name

```



**参数**:

- `table`: string

- `column`: string



**返回**: `string`


**调用**: ns.toDBName() 


---




#### Method: JoinTableName


```go
func (NamingStrategy) JoinTableName(str string) (string)
```


JoinTableName convert string to join table name



```
JoinTableName convert string to join table name

```



**参数**:

- `str`: string



**返回**: `string`


**调用**: strings.ToLower() ns.toDBName() inflection.Plural() ns.toDBName() 


---




#### Method: RelationshipFKName


```go
func (NamingStrategy) RelationshipFKName(rel Relationship) (string)
```


RelationshipFKName generate fk name for relation



```
RelationshipFKName generate fk name for relation

```



**参数**:

- `rel`: Relationship



**返回**: `string`


**调用**: ns.formatName() ns.toDBName() 


---




#### Method: CheckerName


```go
func (NamingStrategy) CheckerName(table, column string) (string)
```


CheckerName generate checker name



```
CheckerName generate checker name

```



**参数**:

- `table`: string

- `column`: string



**返回**: `string`


**调用**: ns.formatName() 


---




#### Method: IndexName


```go
func (NamingStrategy) IndexName(table, column string) (string)
```


IndexName generate index name



```
IndexName generate index name

```



**参数**:

- `table`: string

- `column`: string



**返回**: `string`


**调用**: ns.formatName() ns.toDBName() 


---




#### Method: UniqueName


```go
func (NamingStrategy) UniqueName(table, column string) (string)
```


UniqueName generate unique constraint name



```
UniqueName generate unique constraint name

```



**参数**:

- `table`: string

- `column`: string



**返回**: `string`


**调用**: ns.formatName() ns.toDBName() 


---






























#### Class: RelationshipType








---




#### Constant: HasOne








---




#### Constant: HasMany








---




#### Constant: BelongsTo








---




#### Constant: Many2Many








---






#### Struct: Relationships








---




#### Struct: Relationship








---




#### Struct: Polymorphic








---




#### Struct: Reference








---




























































#### Struct: Constraint








---




#### Method: GetName


```go
func (*Constraint) GetName() (string)
```






**返回**: `string`



---




#### Method: Build


```go
func (*Constraint) Build() (string, []interface{})
```






**返回**: `string`


**调用**: make() len() append() make() len() append() append() 


---




#### Method: ParseConstraint


```go
func (*Relationship) ParseConstraint() (*Constraint)
```






**返回**: `*Constraint`


**调用**: len() len() strings.IndexByte() ParseTagSetting() regEnLetterAndMidline.MatchString() append() append() 


---










#### Method: ToQueryConditions


```go
func (*Relationship) ToQueryConditions(ctx context.Context, reflectValue reflect.Value) ([]clause.Expression)
```





**参数**:

- `ctx`: context.Context

- `reflectValue`: reflect.Value



**返回**: `[]clause.Expression`


**调用**: append() append() append() append() append() append() append() append() append() GetIdentityFieldValuesMap() ToQueryValues() append() 


---
























#### Variable: ErrUnsupportedDataType








---




#### Struct: Schema








---




#### Method: String


```go
func (*Schema) String() (string)
```






**返回**: `string`


**调用**: fmt.Sprintf() fmt.Sprintf() 


---




#### Method: MakeSlice


```go
func (*Schema) MakeSlice() (reflect.Value)
```






**返回**: `reflect.Value`


**调用**: reflect.MakeSlice() reflect.SliceOf() reflect.PointerTo() reflect.New() slice.Type() results.Elem() 


---




#### Method: LookUpField


```go
func (*Schema) LookUpField(name string) (*Field)
```





**参数**:

- `name`: string



**返回**: `*Field`



---




#### Method: LookUpFieldByBindName


```go
func (*Schema) LookUpFieldByBindName(bindNames []string, name string) (*Field)
```


LookUpFieldByBindName looks for the closest field in the embedded struct.



```
LookUpFieldByBindName looks for the closest field in the embedded struct.

	type Struct struct {
		Embedded struct {
			ID string // is selected by LookUpFieldByBindName([]string{"Embedded", "ID"}, "ID")
		}
		ID string // is selected by LookUpFieldByBindName([]string{"ID"}, "ID")
	}

```



**参数**:

- `bindNames`: []string

- `name`: string



**返回**: `*Field`


**调用**: len() strings.Join() 


---




#### Interface: Tabler








---




#### Interface: TablerWithNamer








---






#### Function: Parse


```go
func Parse(dest interface{}, cacheStore *sync.Map, namer Namer) (*Schema, error)
```


Parse get data type from dialector



```
Parse get data type from dialector

```



**参数**:

- `dest`: interface{}

- `cacheStore`: *sync.Map

- `namer`: Namer



**返回**: `*Schema`


**调用**: ParseWithSpecialTableName() 


---




#### Function: ParseWithSpecialTableName


```go
func ParseWithSpecialTableName(dest interface{}, cacheStore *sync.Map, namer Namer, specialTableName string) (*Schema, error)
```


ParseWithSpecialTableName get data type from dialector with extra schema table



```
ParseWithSpecialTableName get data type from dialector with extra schema table

```



**参数**:

- `dest`: interface{}

- `cacheStore`: *sync.Map

- `namer`: Namer

- `specialTableName`: string



**返回**: `*Schema`


**调用**: fmt.Errorf() reflect.ValueOf() modelType.Kind() modelType.Elem() modelType.Kind() modelType.Kind() reflect.Indirect() reflect.ValueOf() modelType.Kind() modelType.Kind() modelType.Kind() modelType.Elem() modelType.Kind() modelType.PkgPath() fmt.Errorf() fmt.Errorf() modelType.PkgPath() modelType.Name() fmt.Sprintf() cacheStore.Load() reflect.New() modelValue.Interface() tabler.TableName() modelValue.Interface() tabler.TableName() namer.TableName() modelType.Name() modelType.Name() make() make() make() make() make() make() close() cacheStore.Load() modelType.NumField() modelType.Field() ast.IsExported() schema.ParseField() append() append() namer.ColumnName() field.BindName() len() len() append() append() append() field.setupValuerAndSetter() schema.LookUpField() schema.LookUpField() len() append() len() len() append() append() append() field.BindName() reflect.New() append() fc.CreateClauses() append() fc.QueryClauses() append() fc.UpdateClauses() append() fc.DeleteClauses() append() cacheStore.LoadOrStore() context.Background() cacheStore.Delete() modelValue.MethodByName() string() methodValue.IsValid() methodValue.Type() path.Dir() reflect.TypeOf() methodValue.Type() reflect.Indirect() reflect.ValueOf() string() context.Background() context.Background() schema.parseRelation() 


---










#### Function: RegisterSerializer


```go
func RegisterSerializer(name string, serializer SerializerInterface)
```


RegisterSerializer register serializer



```
RegisterSerializer register serializer

```



**参数**:

- `name`: string

- `serializer`: SerializerInterface




**调用**: serializerMap.Store() strings.ToLower() 


---




#### Function: GetSerializer


```go
func GetSerializer(name string) (SerializerInterface, bool)
```


GetSerializer get serializer



```
GetSerializer get serializer

```



**参数**:

- `name`: string



**返回**: `SerializerInterface`


**调用**: serializerMap.Load() strings.ToLower() 


---






#### Method: Scan


```go
func (*serializer) Scan(value interface{}) (error)
```


Scan implements sql.Scanner interface



```
Scan implements sql.Scanner interface

```



**参数**:

- `value`: interface{}



**返回**: `error`



---




#### Method: Value


```go
func (serializer) Value() (driver.Value, error)
```


Value implements driver.Valuer interface



```
Value implements driver.Valuer interface

```




**返回**: `driver.Value`



---




#### Interface: SerializerInterface








---




#### Interface: SerializerValuerInterface








---




#### Struct: JSONSerializer








---




#### Method: Scan


```go
func (JSONSerializer) Scan(ctx context.Context, field *Field, dst reflect.Value, dbValue interface{}) (error)
```


Scan implements serializer interface



```
Scan implements serializer interface

```



**参数**:

- `ctx`: context.Context

- `field`: *Field

- `dst`: reflect.Value

- `dbValue`: interface{}



**返回**: `error`


**调用**: reflect.New() json.Marshal() len() json.Unmarshal() fieldValue.Interface() field.ReflectValueOf() fieldValue.Elem() 


---






#### Method: Value


```go
func (JSONSerializer) Value(ctx context.Context, field *Field, dst reflect.Value, fieldValue interface{}) (interface{}, error)
```


Value implements serializer interface



```
Value implements serializer interface

```



**参数**:

- `ctx`: context.Context

- `field`: *Field

- `dst`: reflect.Value

- `fieldValue`: interface{}



**返回**: `interface{}`


**调用**: json.Marshal() string() string() 


---




#### Struct: UnixSecondSerializer








---




#### Method: Scan


```go
func (UnixSecondSerializer) Scan(ctx context.Context, field *Field, dst reflect.Value, dbValue interface{}) (error)
```


Scan implements serializer interface



```
Scan implements serializer interface

```



**参数**:

- `ctx`: context.Context

- `field`: *Field

- `dst`: reflect.Value

- `dbValue`: interface{}



**返回**: `error`


**调用**: t.Scan() field.Set() 


---




#### Method: Value


```go
func (UnixSecondSerializer) Value(ctx context.Context, field *Field, dst reflect.Value, fieldValue interface{}) (interface{}, error)
```


Value implements serializer interface



```
Value implements serializer interface

```



**参数**:

- `ctx`: context.Context

- `field`: *Field

- `dst`: reflect.Value

- `fieldValue`: interface{}



**返回**: `interface{}`


**调用**: reflect.ValueOf() time.Unix() rv.Int() rv.Uint() fmt.Errorf() time.Unix() int64() rv.IsZero() time.Unix() rv.Elem() rv.IsZero() rv.Elem() fmt.Errorf() time.Unix() int64() fmt.Errorf() 


---




#### Struct: GobSerializer








---




#### Method: Scan


```go
func (GobSerializer) Scan(ctx context.Context, field *Field, dst reflect.Value, dbValue interface{}) (error)
```


Scan implements serializer interface



```
Scan implements serializer interface

```



**参数**:

- `ctx`: context.Context

- `field`: *Field

- `dst`: reflect.Value

- `dbValue`: interface{}



**返回**: `error`


**调用**: reflect.New() fmt.Errorf() len() gob.NewDecoder() bytes.NewBuffer() decoder.Decode() fieldValue.Interface() field.ReflectValueOf() fieldValue.Elem() 


---






#### Method: Value


```go
func (GobSerializer) Value(ctx context.Context, field *Field, dst reflect.Value, fieldValue interface{}) (interface{}, error)
```


Value implements serializer interface



```
Value implements serializer interface

```



**参数**:

- `ctx`: context.Context

- `field`: *Field

- `dst`: reflect.Value

- `fieldValue`: interface{}



**返回**: `interface{}`


**调用**: new() gob.NewEncoder() buf.Bytes() 


---






#### Function: ParseTagSetting


```go
func ParseTagSetting(str string, sep string) (map[string]string)
```





**参数**:

- `str`: string

- `sep`: string



**返回**: `map[string]string`


**调用**: strings.Split() len() strings.HasSuffix() len() len() append() strings.Split() strings.TrimSpace() strings.ToUpper() len() strings.Join() strings.ReplaceAll() 


---






#### Function: GetRelationsValues


```go
func GetRelationsValues(ctx context.Context, reflectValue reflect.Value, rels []*Relationship) (reflect.Value)
```


GetRelationsValues get relations's values from a reflect value



```
GetRelationsValues get relations's values from a reflect value

```



**参数**:

- `ctx`: context.Context

- `reflectValue`: reflect.Value

- `rels`: []*Relationship



**返回**: `reflect.Value`


**调用**: reflect.MakeSlice() reflect.SliceOf() reflect.PointerTo() reflect.Indirect() result.Kind() reflect.Append() result.Addr() result.Len() result.Index() elem.Kind() reflect.Append() reflect.Append() elem.Addr() reflectValue.Kind() appendToResults() reflectValue.Len() appendToResults() reflectValue.Index() 


---




#### Function: GetIdentityFieldValuesMap


```go
func GetIdentityFieldValuesMap(ctx context.Context, reflectValue reflect.Value, fields []*Field) (map[string][]reflect.Value, [][]interface{})
```


GetIdentityFieldValuesMap get identity map from fields



```
GetIdentityFieldValuesMap get identity map from fields

```



**参数**:

- `ctx`: context.Context

- `reflectValue`: reflect.Value

- `fields`: []*Field



**返回**: `map[string][]reflect.Value`


**调用**: reflectValue.Kind() reflectValue.Kind() reflectValue.Elem() reflectValue.Kind() make() len() reflectValue.MapIndex() reflect.ValueOf() mapValue.IsZero() reflectValue.MapIndex() reflect.ValueOf() mapValue.Interface() utils.ToStringKey() make() len() field.ValueOf() utils.ToStringKey() reflectValue.Len() reflectValue.Index() elem.Interface() elem.Kind() elem.CanAddr() elem.Addr() make() len() field.ValueOf() utils.ToStringKey() append() append() 


---














#### Function: GetIdentityFieldValuesMapFromValues


```go
func GetIdentityFieldValuesMapFromValues(ctx context.Context, values []interface{}, fields []*Field) (map[string][]reflect.Value, [][]interface{})
```


GetIdentityFieldValuesMapFromValues get identity map from fields



```
GetIdentityFieldValuesMapFromValues get identity map from fields

```



**参数**:

- `ctx`: context.Context

- `values`: []interface{}

- `fields`: []*Field



**返回**: `map[string][]reflect.Value`


**调用**: GetIdentityFieldValuesMap() reflect.Indirect() reflect.ValueOf() append() append() 


---




#### Function: ToQueryValues


```go
func ToQueryValues(table string, foreignKeys []string, foreignValues [][]interface{}) (interface{}, []interface{})
```


ToQueryValues to query values



```
ToQueryValues to query values

```



**参数**:

- `table`: string

- `foreignKeys`: []string

- `foreignValues`: [][]interface{}



**返回**: `interface{}`


**调用**: make() len() len() make() len() 


---






#### Class: DeletedAt








---




#### Method: Scan


```go
func (*DeletedAt) Scan(value interface{}) (error)
```


Scan implements the Scanner interface.



```
Scan implements the Scanner interface.

```



**参数**:

- `value`: interface{}



**返回**: `error`



---




#### Method: Value


```go
func (DeletedAt) Value() (driver.Value, error)
```


Value implements the driver Valuer interface.



```
Value implements the driver Valuer interface.

```




**返回**: `driver.Value`



---




#### Method: MarshalJSON


```go
func (DeletedAt) MarshalJSON() ([]byte, error)
```






**返回**: `[]byte`


**调用**: json.Marshal() json.Marshal() 


---




#### Method: UnmarshalJSON


```go
func (*DeletedAt) UnmarshalJSON(b []byte) (error)
```





**参数**:

- `b`: []byte



**返回**: `error`


**调用**: string() json.Unmarshal() 


---




#### Method: QueryClauses


```go
func (DeletedAt) QueryClauses(f *schema.Field) ([]clause.Interface)
```





**参数**:

- `f`: *schema.Field



**返回**: `[]clause.Interface`


**调用**: parseZeroValueTag() 


---




#### Struct: SoftDeleteQueryClause








---




#### Method: Name


```go
func (SoftDeleteQueryClause) Name() (string)
```






**返回**: `string`



---




#### Method: Build


```go
func (SoftDeleteQueryClause) Build(clause.Builder)
```








---




#### Method: MergeClause


```go
func (SoftDeleteQueryClause) MergeClause(*clause.Clause)
```








---




#### Method: ModifyStatement


```go
func (SoftDeleteQueryClause) ModifyStatement(stmt *Statement)
```





**参数**:

- `stmt`: *Statement




**调用**: len() len() clause.And() stmt.AddClause() 


---




#### Method: UpdateClauses


```go
func (DeletedAt) UpdateClauses(f *schema.Field) ([]clause.Interface)
```





**参数**:

- `f`: *schema.Field



**返回**: `[]clause.Interface`


**调用**: parseZeroValueTag() 


---




#### Struct: SoftDeleteUpdateClause








---




#### Method: Name


```go
func (SoftDeleteUpdateClause) Name() (string)
```






**返回**: `string`



---




#### Method: Build


```go
func (SoftDeleteUpdateClause) Build(clause.Builder)
```








---




#### Method: MergeClause


```go
func (SoftDeleteUpdateClause) MergeClause(*clause.Clause)
```








---




#### Method: ModifyStatement


```go
func (SoftDeleteUpdateClause) ModifyStatement(stmt *Statement)
```





**参数**:

- `stmt`: *Statement




**调用**: SoftDeleteQueryClause() 


---




#### Method: DeleteClauses


```go
func (DeletedAt) DeleteClauses(f *schema.Field) ([]clause.Interface)
```





**参数**:

- `f`: *schema.Field



**返回**: `[]clause.Interface`


**调用**: parseZeroValueTag() 


---




#### Struct: SoftDeleteDeleteClause








---




#### Method: Name


```go
func (SoftDeleteDeleteClause) Name() (string)
```






**返回**: `string`



---




#### Method: Build


```go
func (SoftDeleteDeleteClause) Build(clause.Builder)
```








---




#### Method: MergeClause


```go
func (SoftDeleteDeleteClause) MergeClause(*clause.Clause)
```








---




#### Method: ModifyStatement


```go
func (SoftDeleteDeleteClause) ModifyStatement(stmt *Statement)
```





**参数**:

- `stmt`: *Statement




**调用**: stmt.AddClause() stmt.SetColumn() schema.GetIdentityFieldValuesMap() schema.ToQueryValues() len() stmt.AddClause() schema.GetIdentityFieldValuesMap() reflect.ValueOf() schema.ToQueryValues() len() stmt.AddClause() SoftDeleteQueryClause() stmt.AddClauseIfNotExists() stmt.Build() 


---




#### Struct: Statement








---






#### Interface: StatementModifier








---




#### Method: WriteString


```go
func (*Statement) WriteString(str string) (int, error)
```


WriteString write string



```
WriteString write string

```



**参数**:

- `str`: string



**返回**: `int`



---




#### Method: WriteByte


```go
func (*Statement) WriteByte(c byte) (error)
```


WriteByte write byte



```
WriteByte write byte

```



**参数**:

- `c`: byte



**返回**: `error`



---




#### Method: WriteQuoted


```go
func (*Statement) WriteQuoted(value interface{})
```


WriteQuoted write quoted value



```
WriteQuoted write quoted value

```



**参数**:

- `value`: interface{}




**调用**: stmt.QuoteTo() 


---




#### Method: QuoteTo


```go
func (*Statement) QuoteTo(writer clause.Writer, field interface{})
```


QuoteTo write quoted value to writer



```
QuoteTo write quoted value to writer

```



**参数**:

- `writer`: clause.Writer

- `field`: interface{}




**调用**: writer.WriteString() write() stmt.AddError() stmt.Parse() write() write() writer.WriteByte() write() write() write() writer.WriteByte() write() len() write() write() writer.WriteString() write() writer.WriteByte() writer.WriteByte() stmt.QuoteTo() writer.WriteByte() v.Build() writer.WriteByte() writer.WriteByte() writer.WriteByte() fmt.Sprint() 


---




#### Method: Quote


```go
func (*Statement) Quote(field interface{}) (string)
```


Quote returns quoted value



```
Quote returns quoted value

```



**参数**:

- `field`: interface{}



**返回**: `string`


**调用**: stmt.QuoteTo() builder.String() 


---






#### Method: AddVar


```go
func (*Statement) AddVar(writer clause.Writer, vars ...interface{})
```


AddVar add var



```
AddVar add var

```



**参数**:

- `writer`: clause.Writer

- `vars`: ...interface{}




**调用**: writer.WriteByte() append() stmt.QuoteTo() reflect.ValueOf() reflectValue.Kind() reflectValue.IsNil() stmt.AddVar() stmt.AddVar() v.GormValue() v.Name() v.MergeClause() c.Build() v.Build() append() append() len() writer.WriteByte() stmt.AddVar() writer.WriteByte() writer.WriteString() v.getInstance() cv.Session() make() len() append() cv.BindVarTo() strings.Replace() bindvar.String() strings.Contains() append() writer.WriteString() reflect.ValueOf() rv.Kind() rv.Len() writer.WriteString() rv.Type() reflect.TypeOf() uint8() append() writer.WriteByte() rv.Len() writer.WriteByte() stmt.AddVar() rv.Index() writer.WriteByte() append() 


---








#### Method: AddClause


```go
func (*Statement) AddClause(v clause.Interface)
```


AddClause add clause



```
AddClause add clause

```



**参数**:

- `v`: clause.Interface




**调用**: optimizer.ModifyStatement() v.Name() v.MergeClause() 


---




#### Method: AddClauseIfNotExists


```go
func (*Statement) AddClauseIfNotExists(v clause.Interface)
```


AddClauseIfNotExists add clause if not exists



```
AddClauseIfNotExists add clause if not exists

```



**参数**:

- `v`: clause.Interface




**调用**: v.Name() stmt.AddClause() 


---




#### Method: BuildCondition


```go
func (*Statement) BuildCondition(query interface{}, args ...interface{}) ([]clause.Expression)
```


BuildCondition build condition



```
BuildCondition build condition

```



**参数**:

- `query`: interface{}

- `args`: ...interface{}



**返回**: `[]clause.Expression`


**调用**: strconv.Atoi() len() len() len() strings.Contains() len() strings.Contains() strings.Contains() strings.TrimSpace() len() make() append() valuer.Value() append() append() v.executeScopes() len() len() clause.AndConditions() append() clause.And() append() append() make() len() append() sort.Strings() strings.Contains() append() make() len() append() sort.Strings() reflect.Indirect() reflect.ValueOf() strings.Contains() reflectValue.Kind() append() append() reflectValue.Len() make() reflectValue.Index() append() append() reflect.Indirect() reflect.ValueOf() reflectValue.Kind() reflectValue.Elem() schema.Parse() len() reflectValue.Kind() field.ValueOf() append() append() reflectValue.Len() field.ValueOf() reflectValue.Index() append() append() reflectValue.IsValid() stmt.AddError() len() len() reflectValue.Kind() reflectValue.Len() make() reflectValue.Index() len() append() clause.And() append() len() clause.And() 


---




#### Method: Build


```go
func (*Statement) Build(clauses ...string)
```


Build build sql with clauses names



```
Build build sql with clauses names

```



**参数**:

- `clauses`: ...string




**调用**: stmt.WriteByte() b() c.Build() 


---






#### Method: Parse


```go
func (*Statement) Parse(value interface{}) (error)
```





**参数**:

- `value`: interface{}



**返回**: `error`


**调用**: stmt.ParseWithSpecialTableName() 


---




#### Method: ParseWithSpecialTableName


```go
func (*Statement) ParseWithSpecialTableName(value interface{}, specialTableName string) (error)
```





**参数**:

- `value`: interface{}

- `specialTableName`: string



**返回**: `error`


**调用**: schema.ParseWithSpecialTableName() strings.Split() len() stmt.Quote() 


---






#### Method: SetColumn


```go
func (*Statement) SetColumn(name string, value interface{}, fromCallbacks ...bool)
```


SetColumn set column's value



```
SetColumn set column's value

	stmt.SetColumn("Name", "jinzhu") // Hooks Method
	stmt.SetColumn("Name", "jinzhu", true) // Callbacks Method

```



**参数**:

- `name`: string

- `value`: interface{}

- `fromCallbacks`: ...bool




**调用**: reflect.ValueOf() destValue.Kind() destValue.Elem() destValue.CanAddr() reflect.New() destValue.Type() destValueCanAddr.Elem() destValueCanAddr.Interface() destValueCanAddr.Elem() destValue.Kind() stmt.AddError() field.Set() stmt.AddError() len() stmt.AddError() field.Set() stmt.AddError() field.Set() stmt.AddError() stmt.AddError() field.Set() stmt.AddError() stmt.AddError() 


---




#### Method: Changed


```go
func (*Statement) Changed(fields ...string) (bool)
```


Changed check model changed or not when updating



```
Changed check model changed or not when updating

```



**参数**:

- `fields`: ...string



**返回**: `bool`


**调用**: modelValue.Kind() stmt.SelectAndOmitColumns() field.ValueOf() utils.AssertEqual() utils.AssertEqual() reflect.ValueOf() destValue.Kind() destValue.Elem() schema.Parse() descSchema.LookUpField() destField.ValueOf() utils.AssertEqual() utils.AssertEqual() len() changed() changed() 


---






#### Method: SelectAndOmitColumns


```go
func (*Statement) SelectAndOmitColumns(requireCreate, requireUpdate bool) (map[string]bool, bool)
```


SelectAndOmitColumns get select and omit columns, select -> true, omit -> false



```
SelectAndOmitColumns get select and omit columns, select -> true, omit -> false

```



**参数**:

- `requireCreate`: bool

- `requireUpdate`: bool



**返回**: `map[string]bool`


**调用**: matchName() processColumn() processColumn() len() 


---






#### Function: CallerFrame


```go
func CallerFrame() (runtime.Frame)
```


CallerFrame retrieves the first relevant stack frame outside of GORM's internal implementation files.



```
CallerFrame retrieves the first relevant stack frame outside of GORM's internal implementation files.
It skips:
  - GORM's core source files (identified by gormSourceDir prefix)
  - Exclude test files (*_test.go)
  - go-gorm/gen's Generated files (*.gen.go)

```




**返回**: `runtime.Frame`


**调用**: runtime.Callers() runtime.CallersFrames() frames.Next() strings.HasPrefix() strings.HasSuffix() strings.HasSuffix() 


---




#### Function: FileWithLineNum


```go
func FileWithLineNum() (string)
```


FileWithLineNum return the file name and line number of the current file



```
FileWithLineNum return the file name and line number of the current file

```




**返回**: `string`


**调用**: CallerFrame() string() strconv.AppendInt() append() int64() 


---




#### Function: IsInvalidDBNameChar


```go
func IsInvalidDBNameChar(c rune) (bool)
```





**参数**:

- `c`: rune



**返回**: `bool`


**调用**: unicode.IsLetter() unicode.IsNumber() 


---




#### Function: CheckTruth


```go
func CheckTruth(vals ...string) (bool)
```


CheckTruth check string true or not



```
CheckTruth check string true or not

```



**参数**:

- `vals`: ...string



**返回**: `bool`


**调用**: strings.EqualFold() 


---




#### Function: ToStringKey


```go
func ToStringKey(values ...interface{}) (string)
```





**参数**:

- `values`: ...interface{}



**返回**: `string`


**调用**: make() len() valuer.Value() string() strconv.FormatUint() uint64() reflect.ValueOf() vv.IsValid() vv.IsZero() fmt.Sprint() reflect.Indirect() strings.Join() 


---




#### Function: Contains


```go
func Contains(elems []string, elem string) (bool)
```





**参数**:

- `elems`: []string

- `elem`: string



**返回**: `bool`



---




#### Function: AssertEqual


```go
func AssertEqual(x, y interface{}) (bool)
```





**参数**:

- `x`: interface{}

- `y`: interface{}



**返回**: `bool`


**调用**: reflect.DeepEqual() reflect.ValueOf() reflect.ValueOf() xval.Kind() xval.IsNil() yval.Kind() yval.IsNil() valuer.Value() valuer.Value() reflect.DeepEqual() 


---




#### Function: ToString


```go
func ToString(value interface{}) (string)
```





**参数**:

- `value`: interface{}



**返回**: `string`


**调用**: strconv.FormatInt() int64() strconv.FormatInt() int64() strconv.FormatInt() int64() strconv.FormatInt() int64() strconv.FormatInt() strconv.FormatUint() uint64() strconv.FormatUint() uint64() strconv.FormatUint() uint64() strconv.FormatUint() uint64() strconv.FormatUint() 


---






#### Function: NestedRelationName


```go
func NestedRelationName(prefix, name string) (string)
```


NestedRelationName nested relationships like `Manager__Company`



```
NestedRelationName nested relationships like `Manager__Company`

```



**参数**:

- `prefix`: string

- `name`: string



**返回**: `string`



---




#### Function: SplitNestedRelationName


```go
func SplitNestedRelationName(name string) ([]string)
```


SplitNestedRelationName Split nested relationships to `[]string{"Manager","Company"}`



```
SplitNestedRelationName Split nested relationships to `[]string{"Manager","Company"}`

```



**参数**:

- `name`: string



**返回**: `[]string`


**调用**: strings.Split() 


---




#### Function: JoinNestedRelationNames


```go
func JoinNestedRelationNames(relationNames []string) (string)
```


JoinNestedRelationNames nested relationships like `Manager__Company`



```
JoinNestedRelationNames nested relationships like `Manager__Company`

```



**参数**:

- `relationNames`: []string



**返回**: `string`


**调用**: strings.Join() 


---




#### Function: RTrimSlice


```go
func RTrimSlice(v []T, trimLen int) ([]T)
```


RTrimSlice Right trims the given slice by given length



```
RTrimSlice Right trims the given slice by given length

```



**参数**:

- `v`: []T

- `trimLen`: int



**返回**: `[]T`


**调用**: len() len() 


---




## 调用关系图


### Association (struct)

- **文件**: E:\cody\goproject\qcoder\mcp-server\code2md\test_projects\gorm\association.go:14
- **调用深度**: 0




