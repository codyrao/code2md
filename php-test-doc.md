
# php-test - Code Documentation

## 项目概览

**生成时间**: 2026-01-26T00:49:03+08:00
**项目路径**: E:\cody\goproject\qcoder\mcp-server\code2md\test_projects\php-test
**主要语言**: 

### 语言统计


- **php**: 4 个文件


### 项目摘要

| 指标 | 数量 |
|------|------|
| 总文件数 | 4 |
| 总代码行数 | 154 |
| 总注释行数 | 0 |
| 函数数量 | 20 |
| 类数量 | 3 |
| 接口数量 | 0 |
| 变量数量 | 3 |
| 常量数量 | 0 |

## 项目结构

├── Database.php
├── User.php
├── UserService.php
└── index.php


## 文件清单


### Database.php

- **语言**: php
- **包**: 
- **代码行数**: 54
- **元素数量**: 12




#### 代码元素


**Class**: `Database`

- **完整名称**: Database
- **文件位置**: 行 3-61
- **可见性**: public








- **子元素**: connection host database username password connect disconnect getConnection query execute 

**Method**: `connect`

- **完整名称**: Database::connect
- **文件位置**: 行 19-29
- **可见性**: public









**Method**: `disconnect`

- **完整名称**: Database::disconnect
- **文件位置**: 行 31-34
- **可见性**: public









**Method**: `getConnection`

- **完整名称**: Database::getConnection
- **文件位置**: 行 36-39
- **可见性**: public









**Method**: `query`

- **完整名称**: Database::query
- **文件位置**: 行 41-50
- **可见性**: public




- **参数**:
  - `sql`: mixed
  - `params`: array (可选)






**Method**: `execute`

- **完整名称**: Database::execute
- **文件位置**: 行 52-60
- **可见性**: public




- **参数**:
  - `sql`: mixed
  - `params`: array (可选)






**Function**: `__construct`

- **完整名称**: __construct
- **文件位置**: 行 11-17
- **可见性**: public




- **参数**:
  - `host`: mixed
  - `database`: mixed
  - `username`: mixed
  - `password`: mixed






**Function**: `connect`

- **完整名称**: connect
- **文件位置**: 行 19-29
- **可见性**: public









**Function**: `disconnect`

- **完整名称**: disconnect
- **文件位置**: 行 31-34
- **可见性**: public









**Function**: `getConnection`

- **完整名称**: getConnection
- **文件位置**: 行 36-39
- **可见性**: public









**Function**: `query`

- **完整名称**: query
- **文件位置**: 行 41-50
- **可见性**: public




- **参数**:
  - `sql`: mixed
  - `params`: array (可选)






**Function**: `execute`

- **完整名称**: execute
- **文件位置**: 行 52-60
- **可见性**: public




- **参数**:
  - `sql`: mixed
  - `params`: array (可选)









### User.php

- **语言**: php
- **包**: 
- **代码行数**: 41
- **元素数量**: 14




#### 代码元素


**Class**: `User`

- **完整名称**: User
- **文件位置**: 行 3-49
- **可见性**: public








- **子元素**: id name email getId getName getEmail setName setEmail toArray 

**Method**: `getId`

- **完整名称**: User::getId
- **文件位置**: 行 16-19
- **可见性**: public









**Method**: `getName`

- **完整名称**: User::getName
- **文件位置**: 行 21-24
- **可见性**: public









**Method**: `getEmail`

- **完整名称**: User::getEmail
- **文件位置**: 行 26-29
- **可见性**: public









**Method**: `setName`

- **完整名称**: User::setName
- **文件位置**: 行 31-34
- **可见性**: public




- **参数**:
  - `name`: mixed






**Method**: `setEmail`

- **完整名称**: User::setEmail
- **文件位置**: 行 36-39
- **可见性**: public




- **参数**:
  - `email`: mixed






**Method**: `toArray`

- **完整名称**: User::toArray
- **文件位置**: 行 41-48
- **可见性**: public









**Function**: `__construct`

- **完整名称**: __construct
- **文件位置**: 行 9-14
- **可见性**: public




- **参数**:
  - `id`: mixed
  - `name`: mixed
  - `email`: mixed






**Function**: `getId`

- **完整名称**: getId
- **文件位置**: 行 16-19
- **可见性**: public









**Function**: `getName`

- **完整名称**: getName
- **文件位置**: 行 21-24
- **可见性**: public









**Function**: `getEmail`

- **完整名称**: getEmail
- **文件位置**: 行 26-29
- **可见性**: public









**Function**: `setName`

- **完整名称**: setName
- **文件位置**: 行 31-34
- **可见性**: public




- **参数**:
  - `name`: mixed






**Function**: `setEmail`

- **完整名称**: setEmail
- **文件位置**: 行 36-39
- **可见性**: public




- **参数**:
  - `email`: mixed






**Function**: `toArray`

- **完整名称**: toArray
- **文件位置**: 行 41-48
- **可见性**: public












### UserService.php

- **语言**: php
- **包**: 
- **代码行数**: 46
- **元素数量**: 14




#### 代码元素


**Class**: `UserService`

- **完整名称**: UserService
- **文件位置**: 行 3-54
- **可见性**: public








- **子元素**: database createUser getUserById getUserByEmail getAllUsers updateUser deleteUser 

**Method**: `createUser`

- **完整名称**: UserService::createUser
- **文件位置**: 行 12-17
- **可见性**: public




- **参数**:
  - `name`: mixed
  - `email`: mixed






**Method**: `getUserById`

- **完整名称**: UserService::getUserById
- **文件位置**: 行 19-25
- **可见性**: public




- **参数**:
  - `id`: mixed






**Method**: `getUserByEmail`

- **完整名称**: UserService::getUserByEmail
- **文件位置**: 行 27-33
- **可见性**: public




- **参数**:
  - `email`: mixed






**Method**: `getAllUsers`

- **完整名称**: UserService::getAllUsers
- **文件位置**: 行 35-39
- **可见性**: public









**Method**: `updateUser`

- **完整名称**: UserService::updateUser
- **文件位置**: 行 41-46
- **可见性**: public




- **参数**:
  - `id`: mixed
  - `name`: mixed
  - `email`: mixed






**Method**: `deleteUser`

- **完整名称**: UserService::deleteUser
- **文件位置**: 行 48-53
- **可见性**: public




- **参数**:
  - `id`: mixed






**Function**: `__construct`

- **完整名称**: __construct
- **文件位置**: 行 7-10
- **可见性**: public




- **参数**:
  - `Database $database`: mixed






**Function**: `createUser`

- **完整名称**: createUser
- **文件位置**: 行 12-17
- **可见性**: public




- **参数**:
  - `name`: mixed
  - `email`: mixed






**Function**: `getUserById`

- **完整名称**: getUserById
- **文件位置**: 行 19-25
- **可见性**: public




- **参数**:
  - `id`: mixed






**Function**: `getUserByEmail`

- **完整名称**: getUserByEmail
- **文件位置**: 行 27-33
- **可见性**: public




- **参数**:
  - `email`: mixed






**Function**: `getAllUsers`

- **完整名称**: getAllUsers
- **文件位置**: 行 35-39
- **可见性**: public









**Function**: `updateUser`

- **完整名称**: updateUser
- **文件位置**: 行 41-46
- **可见性**: public




- **参数**:
  - `id`: mixed
  - `name`: mixed
  - `email`: mixed






**Function**: `deleteUser`

- **完整名称**: deleteUser
- **文件位置**: 行 48-53
- **可见性**: public




- **参数**:
  - `id`: mixed









### index.php

- **语言**: php
- **包**: 
- **代码行数**: 13
- **元素数量**: 3




#### 代码元素


**Variable**: `database`

- **完整名称**: database
- **文件位置**: 行 7-7
- **可见性**: public

- **类型**: `mixed`




- **值**: `new Database('localhost', 'testdb', 'root', 'password')`


**Variable**: `userService`

- **完整名称**: userService
- **文件位置**: 行 10-10
- **可见性**: public

- **类型**: `mixed`




- **值**: `new UserService($database)`


**Variable**: `users`

- **完整名称**: users
- **文件位置**: 行 15-15
- **可见性**: public

- **类型**: `mixed`




- **值**: `$userService->getAllUsers()`






## API 参考

### 公共接口



#### Class: Database








---




#### Method: connect








---




#### Method: disconnect








---




#### Method: getConnection








---




#### Method: query





**参数**:

- `sql`: mixed

- `params`: array (可选)





---




#### Method: execute





**参数**:

- `sql`: mixed

- `params`: array (可选)





---




#### Function: __construct





**参数**:

- `host`: mixed

- `database`: mixed

- `username`: mixed

- `password`: mixed





---




#### Function: connect








---




#### Function: disconnect








---




#### Function: getConnection








---




#### Function: query





**参数**:

- `sql`: mixed

- `params`: array (可选)





---




#### Function: execute





**参数**:

- `sql`: mixed

- `params`: array (可选)





---




#### Class: User








---




#### Method: getId








---




#### Method: getName








---




#### Method: getEmail








---




#### Method: setName





**参数**:

- `name`: mixed





---




#### Method: setEmail





**参数**:

- `email`: mixed





---




#### Method: toArray








---




#### Function: __construct





**参数**:

- `id`: mixed

- `name`: mixed

- `email`: mixed





---




#### Function: getId








---




#### Function: getName








---




#### Function: getEmail








---




#### Function: setName





**参数**:

- `name`: mixed





---




#### Function: setEmail





**参数**:

- `email`: mixed





---




#### Function: toArray








---




#### Class: UserService








---




#### Method: createUser





**参数**:

- `name`: mixed

- `email`: mixed





---




#### Method: getUserById





**参数**:

- `id`: mixed





---




#### Method: getUserByEmail





**参数**:

- `email`: mixed





---




#### Method: getAllUsers








---




#### Method: updateUser





**参数**:

- `id`: mixed

- `name`: mixed

- `email`: mixed





---




#### Method: deleteUser





**参数**:

- `id`: mixed





---




#### Function: __construct





**参数**:

- `Database $database`: mixed





---




#### Function: createUser





**参数**:

- `name`: mixed

- `email`: mixed





---




#### Function: getUserById





**参数**:

- `id`: mixed





---




#### Function: getUserByEmail





**参数**:

- `email`: mixed





---




#### Function: getAllUsers








---




#### Function: updateUser





**参数**:

- `id`: mixed

- `name`: mixed

- `email`: mixed





---




#### Function: deleteUser





**参数**:

- `id`: mixed





---




#### Variable: database








---




#### Variable: userService








---




#### Variable: users








---




## 调用关系图


### Database (class)

- **文件**: E:\cody\goproject\qcoder\mcp-server\code2md\test_projects\php-test\Database.php:3
- **调用深度**: 0




