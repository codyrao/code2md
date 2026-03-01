
# js-test - Code Documentation

## 项目概览

**生成时间**: 2026-01-25T23:15:55+08:00
**项目路径**: E:\cody\goproject\qcoder\mcp-server\code2md\test_projects\js-test
**主要语言**: 

### 语言统计


- **javascript**: 48 个文件


### 项目摘要

| 指标 | 数量 |
|------|------|
| 总文件数 | 48 |
| 总代码行数 | 2362 |
| 总注释行数 | 0 |
| 函数数量 | 113 |
| 类数量 | 0 |
| 接口数量 | 0 |
| 变量数量 | 375 |
| 常量数量 | 3 |

## 项目结构

├── examples/
│   ├── auth/
│   │   ├── views/
│   │   └── index.js
│   ├── content-negotiation/
│   │   ├── db.js
│   │   ├── index.js
│   │   └── users.js
│   ├── cookie-sessions/
│   │   └── index.js
│   ├── cookies/
│   │   └── index.js
│   ├── downloads/
│   │   ├── files/
│   │   │   └── notes/
│   │   └── index.js
│   ├── ejs/
│   │   ├── views/
│   │   └── index.js
│   ├── error/
│   │   └── index.js
│   ├── error-pages/
│   │   ├── views/
│   │   └── index.js
│   ├── hello-world/
│   │   └── index.js
│   ├── markdown/
│   │   ├── views/
│   │   └── index.js
│   ├── multi-router/
│   │   ├── controllers/
│   │   │   ├── api_v1.js
│   │   │   └── api_v2.js
│   │   └── index.js
│   ├── mvc/
│   │   ├── controllers/
│   │   │   ├── main/
│   │   │   │   └── index.js
│   │   │   ├── pet/
│   │   │   │   ├── views/
│   │   │   │   └── index.js
│   │   │   ├── user/
│   │   │   │   ├── views/
│   │   │   │   └── index.js
│   │   │   └── user-pet/
│   │   │       └── index.js
│   │   ├── lib/
│   │   │   └── boot.js
│   │   ├── views/
│   │   ├── db.js
│   │   └── index.js
│   ├── online/
│   │   └── index.js
│   ├── params/
│   │   └── index.js
│   ├── resource/
│   │   └── index.js
│   ├── route-map/
│   │   └── index.js
│   ├── route-middleware/
│   │   └── index.js
│   ├── route-separation/
│   │   ├── views/
│   │   │   ├── posts/
│   │   │   └── users/
│   │   ├── index.js
│   │   ├── post.js
│   │   ├── site.js
│   │   └── user.js
│   ├── search/
│   │   └── index.js
│   ├── session/
│   │   ├── index.js
│   │   └── redis.js
│   ├── static-files/
│   │   └── index.js
│   ├── vhost/
│   │   └── index.js
│   ├── view-constructor/
│   │   ├── github-view.js
│   │   └── index.js
│   ├── view-locals/
│   │   ├── views/
│   │   ├── index.js
│   │   └── user.js
│   └── web-service/
│       └── index.js
├── lib/
│   ├── application.js
│   ├── express.js
│   ├── request.js
│   ├── response.js
│   ├── utils.js
│   └── view.js
└── index.js


## 文件清单


### examples\auth\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 88
- **元素数量**: 10




#### 代码元素


**Function**: `authenticate`

- **完整名称**: authenticate
- **文件位置**: 行 60-73
- **可见性**: public









**Function**: `restrict`

- **完整名称**: restrict
- **文件位置**: 行 75-82
- **可见性**: public









**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../..')`


**Variable**: `hash`

- **完整名称**: hash
- **文件位置**: 行 8-8
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('pbkdf2-password')()
var path = require('node:path')`


**Variable**: `session`

- **完整名称**: session
- **文件位置**: 行 10-10
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('express-session')`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 12-12
- **可见性**: private

- **类型**: `unknown`




- **值**: `module.exports = express()`


**Variable**: `err`

- **完整名称**: err
- **文件位置**: 行 31-31
- **可见性**: private

- **类型**: `unknown`




- **值**: `req.session.error`


**Variable**: `msg`

- **完整名称**: msg
- **文件位置**: 行 32-32
- **可见性**: private

- **类型**: `unknown`




- **值**: `req.session.success`


**Variable**: `users`

- **完整名称**: users
- **文件位置**: 行 43-43
- **可见性**: private

- **类型**: `object`




- **值**: `{
  tj: { name: 'tj' }
}`


**Variable**: `user`

- **完整名称**: user
- **文件位置**: 行 62-62
- **可见性**: private

- **类型**: `unknown`




- **值**: `users[name]`





### examples\content-negotiation\db.js

- **语言**: javascript
- **包**: 
- **代码行数**: 6
- **元素数量**: 1




#### 代码元素


**Variable**: `users`

- **完整名称**: users
- **文件位置**: 行 3-3
- **可见性**: private

- **类型**: `array`




- **值**: `[]`





### examples\content-negotiation\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 32
- **元素数量**: 8




#### 代码元素


**Function**: `html`

- **完整名称**: html
- **文件位置**: 行 11-15
- **可见性**: public









**Function**: `text`

- **完整名称**: text
- **文件位置**: 行 17-21
- **可见性**: public









**Function**: `json`

- **完整名称**: json
- **文件位置**: 行 23-25
- **可见性**: public









**Function**: `format`

- **完整名称**: format
- **文件位置**: 行 33-38
- **可见性**: public









**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 3-3
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../../')`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 4-4
- **可见性**: private

- **类型**: `unknown`




- **值**: `module.exports = express()`


**Variable**: `users`

- **完整名称**: users
- **文件位置**: 行 5-5
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('./db')`


**Variable**: `obj`

- **完整名称**: obj
- **文件位置**: 行 34-34
- **可见性**: private

- **类型**: `unknown`




- **值**: `require(path)`





### examples\content-negotiation\users.js

- **语言**: javascript
- **包**: 
- **代码行数**: 15
- **元素数量**: 1




#### 代码元素


**Variable**: `users`

- **完整名称**: users
- **文件位置**: 行 3-3
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('./db')`





### examples\cookie-sessions\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 13
- **元素数量**: 3




#### 代码元素


**Variable**: `cookieSession`

- **完整名称**: cookieSession
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('cookie-session')`


**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 8-8
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../../')`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 10-10
- **可见性**: private

- **类型**: `unknown`




- **值**: `module.exports = express()`





### examples\cookies\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 32
- **元素数量**: 5




#### 代码元素


**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../../')`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 8-8
- **可见性**: private

- **类型**: `unknown`




- **值**: `module.exports = express()`


**Variable**: `logger`

- **完整名称**: logger
- **文件位置**: 行 9-9
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('morgan')`


**Variable**: `cookieParser`

- **完整名称**: cookieParser
- **文件位置**: 行 10-10
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('cookie-parser')`


**Variable**: `minute`

- **完整名称**: minute
- **文件位置**: 行 40-40
- **可见性**: private

- **类型**: `number`




- **值**: `60000`





### examples\downloads\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 25
- **元素数量**: 4




#### 代码元素


**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../../')`


**Variable**: `path`

- **完整名称**: path
- **文件位置**: 行 8-8
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:path')`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 10-10
- **可见性**: private

- **类型**: `unknown`




- **值**: `module.exports = express()`


**Variable**: `FILES_DIR`

- **完整名称**: FILES_DIR
- **文件位置**: 行 13-13
- **可见性**: public

- **类型**: `unknown`




- **值**: `path.join(__dirname, 'files')

app.get('/', function(req, res){
  res.send('<ul>' +
    '<li>Download <a href="/files/notes/groceries.txt">notes/groceries.txt</a>.</li>' +
    '<li>Download <a href="/files/amazing.txt">amazing.txt</a>.</li>' +
    '<li>Download <a href="/files/missing.txt">missing.txt</a>.</li>' +
    '<li>Download <a href="/files/CCTV大赛上海分赛区.txt">CCTV大赛上海分赛区.txt</a>.</li>' +
    '</ul>')
})`





### examples\ejs\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 24
- **元素数量**: 4




#### 代码元素


**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../../')`


**Variable**: `path`

- **完整名称**: path
- **文件位置**: 行 8-8
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:path')`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 10-10
- **可见性**: private

- **类型**: `unknown`




- **值**: `module.exports = express()`


**Variable**: `users`

- **完整名称**: users
- **文件位置**: 行 39-39
- **可见性**: private

- **类型**: `array`




- **值**: `[
  { name: 'tobi', email: 'tobi@learnboost.com' },
  { name: 'loki', email: 'loki@learnboost.com' },
  { name: 'jane', email: 'jane@learnboost.com' }
]`





### examples\error\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 24
- **元素数量**: 5




#### 代码元素


**Function**: `error`

- **完整名称**: error
- **文件位置**: 行 20-27
- **可见性**: public









**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../../')`


**Variable**: `logger`

- **完整名称**: logger
- **文件位置**: 行 8-8
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('morgan')`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 9-9
- **可见性**: private

- **类型**: `unknown`




- **值**: `module.exports = express()`


**Variable**: `test`

- **完整名称**: test
- **文件位置**: 行 10-10
- **可见性**: private

- **类型**: `unknown`




- **值**: `app.get('env') === 'test'

if (!test) app.use(logger('dev'))`





### examples\error-pages\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 47
- **元素数量**: 9




#### 代码元素


**Function**: `html`

- **完整名称**: html
- **文件位置**: 行 67-69
- **可见性**: public









**Function**: `json`

- **完整名称**: json
- **文件位置**: 行 70-72
- **可见性**: public









**Function**: `default`

- **完整名称**: default
- **文件位置**: 行 73-75
- **可见性**: public









**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../../')`


**Variable**: `path`

- **完整名称**: path
- **文件位置**: 行 8-8
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:path')`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 9-9
- **可见性**: private

- **类型**: `unknown`




- **值**: `module.exports = express()`


**Variable**: `logger`

- **完整名称**: logger
- **文件位置**: 行 10-10
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('morgan')`


**Variable**: `silent`

- **完整名称**: silent
- **文件位置**: 行 11-11
- **可见性**: private

- **类型**: `unknown`




- **值**: `process.env.NODE_ENV === 'test'

// general config
app.set('views', path.join(__dirname, 'views'))`


**Variable**: `err`

- **完整名称**: err
- **文件位置**: 行 43-43
- **可见性**: private

- **类型**: `unknown`




- **值**: `new Error('not allowed!')`





### examples\hello-world\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 10
- **元素数量**: 2




#### 代码元素


**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 3-3
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../../')`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 5-5
- **可见性**: private

- **类型**: `unknown`




- **值**: `module.exports = express()

app.get('/', function(req, res){
  res.send('Hello World')`





### examples\markdown\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 28
- **元素数量**: 7




#### 代码元素


**Variable**: `escapeHtml`

- **完整名称**: escapeHtml
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('escape-html')`


**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 8-8
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../..')`


**Variable**: `fs`

- **完整名称**: fs
- **文件位置**: 行 9-9
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:fs')`


**Variable**: `marked`

- **完整名称**: marked
- **文件位置**: 行 10-10
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('marked')`


**Variable**: `path`

- **完整名称**: path
- **文件位置**: 行 11-11
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:path')`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 13-13
- **可见性**: private

- **类型**: `unknown`




- **值**: `module.exports = express()`


**Variable**: `html`

- **完整名称**: html
- **文件位置**: 行 20-20
- **可见性**: private

- **类型**: `unknown`




- **值**: `marked.parse(str).replace(/\{([^}]+)\}/g, function(_, name){
      return escapeHtml(options[name] || '')`





### examples\multi-router\controllers\api_v1.js

- **语言**: javascript
- **包**: 
- **代码行数**: 10
- **元素数量**: 2




#### 代码元素


**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 3-3
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../../..')`


**Variable**: `apiv1`

- **完整名称**: apiv1
- **文件位置**: 行 5-5
- **可见性**: private

- **类型**: `unknown`




- **值**: `express.Router()`





### examples\multi-router\controllers\api_v2.js

- **语言**: javascript
- **包**: 
- **代码行数**: 10
- **元素数量**: 2




#### 代码元素


**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 3-3
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../../..')`


**Variable**: `apiv2`

- **完整名称**: apiv2
- **文件位置**: 行 5-5
- **可见性**: private

- **类型**: `unknown`




- **值**: `express.Router()`





### examples\multi-router\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 12
- **元素数量**: 2




#### 代码元素


**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 3-3
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../..')`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 5-5
- **可见性**: private

- **类型**: `unknown`




- **值**: `module.exports = express()`





### examples\mvc\controllers\main\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 4
- **元素数量**: 0





### examples\mvc\controllers\pet\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 21
- **元素数量**: 3




#### 代码元素


**Variable**: `db`

- **完整名称**: db
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../../db')`


**Variable**: `pet`

- **完整名称**: pet
- **文件位置**: 行 12-12
- **可见性**: private

- **类型**: `unknown`




- **值**: `db.pets[req.params.pet_id]`


**Variable**: `body`

- **完整名称**: body
- **文件位置**: 行 27-27
- **可见性**: private

- **类型**: `unknown`




- **值**: `req.body`





### examples\mvc\controllers\user\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 27
- **元素数量**: 3




#### 代码元素


**Variable**: `db`

- **完整名称**: db
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../../db')`


**Variable**: `id`

- **完整名称**: id
- **文件位置**: 行 12-12
- **可见性**: private

- **类型**: `unknown`




- **值**: `req.params.user_id`


**Variable**: `body`

- **完整名称**: body
- **文件位置**: 行 37-37
- **可见性**: private

- **类型**: `unknown`




- **值**: `req.body`





### examples\mvc\controllers\user-pet\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 15
- **元素数量**: 5




#### 代码元素


**Variable**: `db`

- **完整名称**: db
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../../db')`


**Variable**: `id`

- **完整名称**: id
- **文件位置**: 行 13-13
- **可见性**: private

- **类型**: `unknown`




- **值**: `req.params.user_id`


**Variable**: `user`

- **完整名称**: user
- **文件位置**: 行 14-14
- **可见性**: private

- **类型**: `unknown`




- **值**: `db.users[id]`


**Variable**: `body`

- **完整名称**: body
- **文件位置**: 行 15-15
- **可见性**: private

- **类型**: `unknown`




- **值**: `req.body`


**Variable**: `pet`

- **完整名称**: pet
- **文件位置**: 行 17-17
- **可见性**: private

- **类型**: `object`




- **值**: `{ name: body.pet.name }`





### examples\mvc\db.js

- **语言**: javascript
- **包**: 
- **代码行数**: 10
- **元素数量**: 2




#### 代码元素


**Variable**: `pets`

- **完整名称**: pets
- **文件位置**: 行 5-5
- **可见性**: private

- **类型**: `unknown`




- **值**: `exports.pets = []`


**Variable**: `users`

- **完整名称**: users
- **文件位置**: 行 12-12
- **可见性**: private

- **类型**: `unknown`




- **值**: `exports.users = []`





### examples\mvc\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 47
- **元素数量**: 8




#### 代码元素


**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../..')`


**Variable**: `logger`

- **完整名称**: logger
- **文件位置**: 行 8-8
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('morgan')`


**Variable**: `path`

- **完整名称**: path
- **文件位置**: 行 9-9
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:path')`


**Variable**: `session`

- **完整名称**: session
- **文件位置**: 行 10-10
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('express-session')`


**Variable**: `methodOverride`

- **完整名称**: methodOverride
- **文件位置**: 行 11-11
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('method-override')`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 13-13
- **可见性**: private

- **类型**: `unknown`




- **值**: `module.exports = express()`


**Variable**: `sess`

- **完整名称**: sess
- **文件位置**: 行 26-26
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.req.session`


**Variable**: `msgs`

- **完整名称**: msgs
- **文件位置**: 行 54-54
- **可见性**: private

- **类型**: `unknown`




- **值**: `req.session.messages || []`





### examples\mvc\lib\boot.js

- **语言**: javascript
- **包**: 
- **代码行数**: 63
- **元素数量**: 10




#### 代码元素


**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../../..')`


**Variable**: `fs`

- **完整名称**: fs
- **文件位置**: 行 8-8
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:fs')`


**Variable**: `path`

- **完整名称**: path
- **文件位置**: 行 9-9
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:path')`


**Variable**: `dir`

- **完整名称**: dir
- **文件位置**: 行 12-12
- **可见性**: private

- **类型**: `unknown`




- **值**: `path.join(__dirname, '..', 'controllers')`


**Variable**: `verbose`

- **完整名称**: verbose
- **文件位置**: 行 13-13
- **可见性**: private

- **类型**: `unknown`




- **值**: `options.verbose`


**Variable**: `file`

- **完整名称**: file
- **文件位置**: 行 15-15
- **可见性**: private

- **类型**: `unknown`




- **值**: `path.join(dir, name)
    if (!fs.statSync(file).isDirectory()) return`


**Variable**: `obj`

- **完整名称**: obj
- **文件位置**: 行 18-18
- **可见性**: private

- **类型**: `unknown`




- **值**: `require(file)`


**Variable**: `name`

- **完整名称**: name
- **文件位置**: 行 19-19
- **可见性**: private

- **类型**: `unknown`




- **值**: `obj.name || name`


**Variable**: `prefix`

- **完整名称**: prefix
- **文件位置**: 行 20-20
- **可见性**: private

- **类型**: `unknown`




- **值**: `obj.prefix || ''`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 21-21
- **可见性**: private

- **类型**: `unknown`




- **值**: `express()`





### examples\online\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 26
- **元素数量**: 6




#### 代码元素


**Function**: `list`

- **完整名称**: list
- **文件位置**: 行 40-44
- **可见性**: public
- **描述**: GET users online.








**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 14-14
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../..')`


**Variable**: `online`

- **完整名称**: online
- **文件位置**: 行 15-15
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('online')`


**Variable**: `redis`

- **完整名称**: redis
- **文件位置**: 行 16-16
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('redis')`


**Variable**: `db`

- **完整名称**: db
- **文件位置**: 行 17-17
- **可见性**: private

- **类型**: `unknown`




- **值**: `redis.createClient()`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 25-25
- **可见性**: private

- **类型**: `unknown`




- **值**: `express()`





### examples\params\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 43
- **元素数量**: 6




#### 代码元素


**Variable**: `createError`

- **完整名称**: createError
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('http-errors')
var express = require('../../')`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 9-9
- **可见性**: private

- **类型**: `unknown`




- **值**: `module.exports = express()`


**Variable**: `users`

- **完整名称**: users
- **文件位置**: 行 13-13
- **可见性**: private

- **类型**: `array`




- **值**: `[
  { name: 'tj' }
  , { name: 'tobi' }
  , { name: 'loki' }
  , { name: 'jane' }
  , { name: 'bandit' }
]`


**Variable**: `from`

- **完整名称**: from
- **文件位置**: 行 64-64
- **可见性**: private

- **类型**: `unknown`




- **值**: `req.params.from`


**Variable**: `to`

- **完整名称**: to
- **文件位置**: 行 65-65
- **可见性**: private

- **类型**: `unknown`




- **值**: `req.params.to`


**Variable**: `names`

- **完整名称**: names
- **文件位置**: 行 66-66
- **可见性**: private

- **类型**: `unknown`




- **值**: `users.map(function(user){ return user.name`





### examples\resource\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 70
- **元素数量**: 15




#### 代码元素


**Function**: `index`

- **完整名称**: index
- **文件位置**: 行 42-44
- **可见性**: public









**Function**: `show`

- **完整名称**: show
- **文件位置**: 行 45-47
- **可见性**: public









**Function**: `destroy`

- **完整名称**: destroy
- **文件位置**: 行 48-52
- **可见性**: public









**Function**: `range`

- **完整名称**: range
- **文件位置**: 行 53-67
- **可见性**: public









**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../../')`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 9-9
- **可见性**: private

- **类型**: `unknown`




- **值**: `module.exports = express()`


**Variable**: `a`

- **完整名称**: a
- **文件位置**: 行 16-16
- **可见性**: private

- **类型**: `unknown`




- **值**: `parseInt(req.params.a, 10)`


**Variable**: `b`

- **完整名称**: b
- **文件位置**: 行 17-17
- **可见性**: private

- **类型**: `unknown`




- **值**: `parseInt(req.params.b, 10)`


**Variable**: `format`

- **完整名称**: format
- **文件位置**: 行 18-18
- **可见性**: private

- **类型**: `unknown`




- **值**: `req.params.format`


**Variable**: `id`

- **完整名称**: id
- **文件位置**: 行 23-23
- **可见性**: private

- **类型**: `unknown`




- **值**: `parseInt(req.params.id, 10)`


**Variable**: `users`

- **完整名称**: users
- **文件位置**: 行 30-30
- **可见性**: private

- **类型**: `array`




- **值**: `[
  { name: 'tj' }
  , { name: 'ciaran' }
  , { name: 'aaron' }
  , { name: 'guillermo' }
  , { name: 'simon' }
  , { name: 'tobi' }
]`


**Variable**: `User`

- **完整名称**: User
- **文件位置**: 行 41-41
- **可见性**: public

- **类型**: `object`




- **值**: `{
  index: function(req, res){
    res.send(users)`


**Variable**: `destroyed`

- **完整名称**: destroyed
- **文件位置**: 行 49-49
- **可见性**: private

- **类型**: `unknown`




- **值**: `id in users`


**Variable**: `range`

- **完整名称**: range
- **文件位置**: 行 54-54
- **可见性**: private

- **类型**: `unknown`




- **值**: `users.slice(a, b + 1)`


**Variable**: `html`

- **完整名称**: html
- **文件位置**: 行 61-61
- **可见性**: private

- **类型**: `string`




- **值**: `'<ul>' + range.map(function(user){
          return '<li>' + user.name + '</li>'`





### examples\route-map\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 57
- **元素数量**: 10




#### 代码元素


**Function**: `get`

- **完整名称**: get
- **文件位置**: 行 22-22
- **可见性**: public









**Function**: `list`

- **完整名称**: list
- **文件位置**: 行 32-34
- **可见性**: public









**Function**: `get`

- **完整名称**: get
- **文件位置**: 行 36-38
- **可见性**: public









**Function**: `delete`

- **完整名称**: delete
- **文件位置**: 行 40-42
- **可见性**: public









**Function**: `list`

- **完整名称**: list
- **文件位置**: 行 32-34
- **可见性**: public









**Function**: `delete`

- **完整名称**: delete
- **文件位置**: 行 40-42
- **可见性**: public









**Variable**: `escapeHtml`

- **完整名称**: escapeHtml
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('escape-html')
var express = require('../../lib/express')`


**Variable**: `verbose`

- **完整名称**: verbose
- **文件位置**: 行 10-10
- **可见性**: private

- **类型**: `unknown`




- **值**: `process.env.NODE_ENV !== 'test'

var app = module.exports = express()`


**Variable**: `users`

- **完整名称**: users
- **文件位置**: 行 31-31
- **可见性**: private

- **类型**: `object`




- **值**: `{
  list: function(req, res){
    res.send('user list')`


**Variable**: `pets`

- **完整名称**: pets
- **文件位置**: 行 45-45
- **可见性**: private

- **类型**: `object`




- **值**: `{
  list: function(req, res){
    res.send('user ' + escapeHtml(req.params.uid) + '\'s pets')
  },

  delete: function(req, res){
    res.send('delete ' + escapeHtml(req.params.uid) + '\'s pet ' + escapeHtml(req.params.pid))
  }
}`





### examples\route-middleware\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 53
- **元素数量**: 7




#### 代码元素


**Function**: `loadUser`

- **完整名称**: loadUser
- **文件位置**: 行 25-34
- **可见性**: public









**Function**: `andRestrictToSelf`

- **完整名称**: andRestrictToSelf
- **文件位置**: 行 36-48
- **可见性**: public









**Function**: `andRestrictTo`

- **完整名称**: andRestrictTo
- **文件位置**: 行 50-58
- **可见性**: public









**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../../lib/express')`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 9-9
- **可见性**: private

- **类型**: `unknown`




- **值**: `express()`


**Variable**: `users`

- **完整名称**: users
- **文件位置**: 行 19-19
- **可见性**: private

- **类型**: `array`




- **值**: `[
  { id: 0, name: 'tj', email: 'tj@vision-media.ca', role: 'member' }
  , { id: 1, name: 'ciaran', email: 'ciaranj@gmail.com', role: 'member' }
  , { id: 2, name: 'aaron', email: 'aaron.heckmann+github@gmail.com', role: 'admin' }
]`


**Variable**: `user`

- **完整名称**: user
- **文件位置**: 行 27-27
- **可见性**: private

- **类型**: `unknown`




- **值**: `users[req.params.id]`





### examples\route-separation\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 32
- **元素数量**: 9




#### 代码元素


**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../..')`


**Variable**: `path`

- **完整名称**: path
- **文件位置**: 行 8-8
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:path')`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 9-9
- **可见性**: private

- **类型**: `unknown`




- **值**: `express()`


**Variable**: `logger`

- **完整名称**: logger
- **文件位置**: 行 10-10
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('morgan')`


**Variable**: `cookieParser`

- **完整名称**: cookieParser
- **文件位置**: 行 11-11
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('cookie-parser')`


**Variable**: `methodOverride`

- **完整名称**: methodOverride
- **文件位置**: 行 12-12
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('method-override')`


**Variable**: `site`

- **完整名称**: site
- **文件位置**: 行 13-13
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('./site')`


**Variable**: `post`

- **完整名称**: post
- **文件位置**: 行 14-14
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('./post')`


**Variable**: `user`

- **完整名称**: user
- **文件位置**: 行 15-15
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('./user')`





### examples\route-separation\post.js

- **语言**: javascript
- **包**: 
- **代码行数**: 9
- **元素数量**: 1




#### 代码元素


**Variable**: `posts`

- **完整名称**: posts
- **文件位置**: 行 5-5
- **可见性**: private

- **类型**: `array`




- **值**: `[
  { title: 'Foo', body: 'some foo bar' },
  { title: 'Foo bar', body: 'more foo bar' },
  { title: 'Foo bar baz', body: 'more foo bar baz' }
]`





### examples\route-separation\site.js

- **语言**: javascript
- **包**: 
- **代码行数**: 4
- **元素数量**: 0





### examples\route-separation\user.js

- **语言**: javascript
- **包**: 
- **代码行数**: 37
- **元素数量**: 4




#### 代码元素


**Variable**: `users`

- **完整名称**: users
- **文件位置**: 行 5-5
- **可见性**: private

- **类型**: `array`




- **值**: `[
  { name: 'TJ', email: 'tj@vision-media.ca' },
  { name: 'Tobi', email: 'tobi@vision-media.ca' }
]`


**Variable**: `id`

- **完整名称**: id
- **文件位置**: 行 15-15
- **可见性**: private

- **类型**: `unknown`




- **值**: `req.params.id`


**Variable**: `err`

- **完整名称**: err
- **文件位置**: 行 20-20
- **可见性**: private

- **类型**: `unknown`




- **值**: `new Error('cannot find user ' + id)`


**Variable**: `user`

- **完整名称**: user
- **文件位置**: 行 43-43
- **可见性**: private

- **类型**: `unknown`




- **值**: `req.body.user`





### examples\search\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 26
- **元素数量**: 6




#### 代码元素


**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 14-14
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../..')`


**Variable**: `path`

- **完整名称**: path
- **文件位置**: 行 15-15
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:path')`


**Variable**: `redis`

- **完整名称**: redis
- **文件位置**: 行 16-16
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('redis')`


**Variable**: `db`

- **完整名称**: db
- **文件位置**: 行 18-18
- **可见性**: private

- **类型**: `unknown`




- **值**: `redis.createClient()`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 22-22
- **可见性**: private

- **类型**: `unknown`




- **值**: `express()`


**Variable**: `query`

- **完整名称**: query
- **文件位置**: 行 39-39
- **可见性**: private

- **类型**: `unknown`




- **值**: `req.params.query`





### examples\session\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 23
- **元素数量**: 4




#### 代码元素


**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 10-10
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../..')`


**Variable**: `session`

- **完整名称**: session
- **文件位置**: 行 11-11
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('express-session')`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 13-13
- **可见性**: private

- **类型**: `unknown`




- **值**: `express()`


**Variable**: `body`

- **完整名称**: body
- **文件位置**: 行 23-23
- **可见性**: private

- **类型**: `string`




- **值**: `''`





### examples\session\redis.js

- **语言**: javascript
- **包**: 
- **代码行数**: 25
- **元素数量**: 6




#### 代码元素


**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../..')`


**Variable**: `logger`

- **完整名称**: logger
- **文件位置**: 行 8-8
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('morgan')`


**Variable**: `session`

- **完整名称**: session
- **文件位置**: 行 9-9
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('express-session')`


**Variable**: `RedisStore`

- **完整名称**: RedisStore
- **文件位置**: 行 13-13
- **可见性**: public

- **类型**: `unknown`




- **值**: `require('connect-redis')(session)`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 15-15
- **可见性**: private

- **类型**: `unknown`




- **值**: `express()`


**Variable**: `body`

- **完整名称**: body
- **文件位置**: 行 28-28
- **可见性**: private

- **类型**: `string`




- **值**: `''`





### examples\static-files\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 15
- **元素数量**: 4




#### 代码元素


**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../..')`


**Variable**: `logger`

- **完整名称**: logger
- **文件位置**: 行 8-8
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('morgan')`


**Variable**: `path`

- **完整名称**: path
- **文件位置**: 行 9-9
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:path')`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 10-10
- **可见性**: private

- **类型**: `unknown`




- **值**: `express()`





### examples\vhost\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 28
- **元素数量**: 6




#### 代码元素


**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../..')`


**Variable**: `logger`

- **完整名称**: logger
- **文件位置**: 行 8-8
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('morgan')`


**Variable**: `vhost`

- **完整名称**: vhost
- **文件位置**: 行 9-9
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('vhost')`


**Variable**: `main`

- **完整名称**: main
- **文件位置**: 行 21-21
- **可见性**: private

- **类型**: `unknown`




- **值**: `express()`


**Variable**: `redirect`

- **完整名称**: redirect
- **文件位置**: 行 35-35
- **可见性**: private

- **类型**: `unknown`




- **值**: `express()`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 44-44
- **可见性**: private

- **类型**: `unknown`




- **值**: `module.exports = express()`





### examples\view-constructor\github-view.js

- **语言**: javascript
- **包**: 
- **代码行数**: 28
- **元素数量**: 7




#### 代码元素


**Function**: `GithubView`

- **完整名称**: GithubView
- **文件位置**: 行 23-30
- **可见性**: public
- **描述**: Render the view.








**Variable**: `https`

- **完整名称**: https
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:https')`


**Variable**: `path`

- **完整名称**: path
- **文件位置**: 行 8-8
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:path')`


**Variable**: `extname`

- **完整名称**: extname
- **文件位置**: 行 9-9
- **可见性**: private

- **类型**: `unknown`




- **值**: `path.extname`


**Variable**: `self`

- **完整名称**: self
- **文件位置**: 行 37-37
- **可见性**: private

- **类型**: `unknown`




- **值**: `this`


**Variable**: `opts`

- **完整名称**: opts
- **文件位置**: 行 38-38
- **可见性**: private

- **类型**: `object`




- **值**: `{
    host: 'raw.githubusercontent.com',
    port: 443,
    path: this.path,
    method: 'GET'
  }`


**Variable**: `buf`

- **完整名称**: buf
- **文件位置**: 行 46-46
- **可见性**: private

- **类型**: `string`




- **值**: `''`





### examples\view-constructor\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 28
- **元素数量**: 5




#### 代码元素


**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../../')`


**Variable**: `GithubView`

- **完整名称**: GithubView
- **文件位置**: 行 8-8
- **可见性**: public

- **类型**: `unknown`




- **值**: `require('./github-view')`


**Variable**: `md`

- **完整名称**: md
- **文件位置**: 行 9-9
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('marked').parse`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 11-11
- **可见性**: private

- **类型**: `unknown`




- **值**: `module.exports = express()`


**Variable**: `html`

- **完整名称**: html
- **文件位置**: 行 16-16
- **可见性**: private

- **类型**: `unknown`




- **值**: `md(str)`





### examples\view-locals\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 80
- **元素数量**: 9




#### 代码元素


**Function**: `ferrets`

- **完整名称**: ferrets
- **文件位置**: 行 17-19
- **可见性**: public









**Function**: `count`

- **完整名称**: count
- **文件位置**: 行 48-54
- **可见性**: public









**Function**: `users`

- **完整名称**: users
- **文件位置**: 行 56-62
- **可见性**: public









**Function**: `count2`

- **完整名称**: count2
- **文件位置**: 行 86-92
- **可见性**: public









**Function**: `users2`

- **完整名称**: users2
- **文件位置**: 行 94-100
- **可见性**: public









**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../..')`


**Variable**: `path`

- **完整名称**: path
- **文件位置**: 行 8-8
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:path')`


**Variable**: `User`

- **完整名称**: User
- **文件位置**: 行 9-9
- **可见性**: public

- **类型**: `unknown`




- **值**: `require('./user')`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 10-10
- **可见性**: private

- **类型**: `unknown`




- **值**: `express()`





### examples\view-locals\user.js

- **语言**: javascript
- **包**: 
- **代码行数**: 23
- **元素数量**: 2




#### 代码元素


**Function**: `User`

- **完整名称**: User
- **文件位置**: 行 7-11
- **可见性**: public









**Variable**: `users`

- **完整名称**: users
- **文件位置**: 行 30-30
- **可见性**: private

- **类型**: `array`




- **值**: `[]`





### examples\web-service\index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 55
- **元素数量**: 11




#### 代码元素


**Function**: `error`

- **完整名称**: error
- **文件位置**: 行 15-19
- **可见性**: public









**Variable**: `express`

- **完整名称**: express
- **文件位置**: 行 7-7
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('../../')`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 9-9
- **可见性**: private

- **类型**: `unknown`




- **值**: `module.exports = express()`


**Variable**: `err`

- **完整名称**: err
- **文件位置**: 行 16-16
- **可见性**: private

- **类型**: `unknown`




- **值**: `new Error(msg)`


**Variable**: `key`

- **完整名称**: key
- **文件位置**: 行 31-31
- **可见性**: private

- **类型**: `unknown`




- **值**: `req.query['api-key']`


**Variable**: `apiKeys`

- **完整名称**: apiKeys
- **文件位置**: 行 49-49
- **可见性**: private

- **类型**: `array`




- **值**: `['foo', 'bar', 'baz']`


**Variable**: `repos`

- **完整名称**: repos
- **文件位置**: 行 53-53
- **可见性**: private

- **类型**: `array`




- **值**: `[
  { name: 'express', url: 'https://github.com/expressjs/express' },
  { name: 'stylus', url: 'https://github.com/learnboost/stylus' },
  { name: 'cluster', url: 'https://github.com/learnboost/cluster' }
]`


**Variable**: `users`

- **完整名称**: users
- **文件位置**: 行 59-59
- **可见性**: private

- **类型**: `array`




- **值**: `[
  { name: 'tobi' }
  , { name: 'loki' }
  , { name: 'jane' }
]`


**Variable**: `userRepos`

- **完整名称**: userRepos
- **文件位置**: 行 65-65
- **可见性**: private

- **类型**: `object`




- **值**: `{
  tobi: [repos[0], repos[1]]
  , loki: [repos[1]]
  , jane: [repos[2]]
}`


**Variable**: `name`

- **完整名称**: name
- **文件位置**: 行 86-86
- **可见性**: private

- **类型**: `unknown`




- **值**: `req.params.name`


**Variable**: `user`

- **完整名称**: user
- **文件位置**: 行 87-87
- **可见性**: private

- **类型**: `unknown`




- **值**: `userRepos[name]`





### index.js

- **语言**: javascript
- **包**: 
- **代码行数**: 2
- **元素数量**: 0





### lib\application.js

- **语言**: javascript
- **包**: 
- **代码行数**: 262
- **元素数量**: 59




#### 代码元素


**Function**: `init`

- **完整名称**: init
- **文件位置**: 行 59-83
- **可见性**: public
- **描述**: Initialize application configuration.








**Function**: `getrouter`

- **完整名称**: getrouter
- **文件位置**: 行 72-81
- **可见性**: public
- **描述**: Initialize application configuration.








**Function**: `defaultConfiguration`

- **完整名称**: defaultConfiguration
- **文件位置**: 行 90-141
- **可见性**: public
- **描述**: Dispatch a req, res pair into the application. Starts pipeline processing.








**Function**: `onmount`

- **完整名称**: onmount
- **文件位置**: 行 109-122
- **可见性**: public
- **描述**: Dispatch a req, res pair into the application. Starts pipeline processing.








**Function**: `handle`

- **完整名称**: handle
- **文件位置**: 行 152-178
- **可见性**: public
- **描述**: Proxy `Router#use()` to add middleware to the app router.








**Function**: `use`

- **完整名称**: use
- **文件位置**: 行 190-244
- **可见性**: public
- **描述**: Proxy to the app `Router#route()`








**Function**: `mounted_app`

- **完整名称**: mounted_app
- **文件位置**: 行 230-237
- **可见性**: public
- **描述**: Proxy to the app `Router#route()`








**Function**: `route`

- **完整名称**: route
- **文件位置**: 行 256-258
- **可见性**: public
- **描述**: Register the given template engine callback `fn`








**Function**: `engine`

- **完整名称**: engine
- **文件位置**: 行 294-308
- **可见性**: public
- **描述**: Proxy to `Router#param()` with one added api feature. The _name_ parameter








**Function**: `param`

- **完整名称**: param
- **文件位置**: 行 322-334
- **可见性**: public
- **描述**: Assign `setting` to `val`, or return `setting`'s value.








**Function**: `set`

- **完整名称**: set
- **文件位置**: 行 351-383
- **可见性**: public
- **描述**: Return the app's absolute pathname








**Function**: `path`

- **完整名称**: path
- **文件位置**: 行 399-403
- **可见性**: public
- **描述**: Check if `setting` is enabled (truthy).








**Function**: `enabled`

- **完整名称**: enabled
- **文件位置**: 行 420-422
- **可见性**: public
- **描述**: Check if `setting` is disabled.








**Function**: `disabled`

- **完整名称**: disabled
- **文件位置**: 行 439-441
- **可见性**: public
- **描述**: Enable `setting`.








**Function**: `enable`

- **完整名称**: enable
- **文件位置**: 行 451-453
- **可见性**: public
- **描述**: Disable `setting`.








**Function**: `disable`

- **完整名称**: disable
- **文件位置**: 行 463-465
- **可见性**: public
- **描述**: Delegate `.VERB(...)` calls to `router.VERB(...)`.








**Function**: `all`

- **完整名称**: all
- **文件位置**: 行 494-503
- **可见性**: public
- **描述**: Render the given view `name` name with `options`








**Function**: `render`

- **完整名称**: render
- **文件位置**: 行 522-575
- **可见性**: public
- **描述**: Listen for connections.








**Function**: `listen`

- **完整名称**: listen
- **文件位置**: 行 598-606
- **可见性**: public
- **描述**: Log error using console.error.








**Function**: `logerror`

- **完整名称**: logerror
- **文件位置**: 行 615-618
- **可见性**: public
- **描述**: Try rendering a view.








**Function**: `tryRender`

- **完整名称**: tryRender
- **文件位置**: 行 625-631
- **可见性**: public









**Variable**: `finalhandler`

- **完整名称**: finalhandler
- **文件位置**: 行 16-16
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('finalhandler')`


**Variable**: `debug`

- **完整名称**: debug
- **文件位置**: 行 17-17
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('debug')('express:application')`


**Variable**: `View`

- **完整名称**: View
- **文件位置**: 行 18-18
- **可见性**: public

- **类型**: `unknown`




- **值**: `require('./view')`


**Variable**: `http`

- **完整名称**: http
- **文件位置**: 行 19-19
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:http')`


**Variable**: `methods`

- **完整名称**: methods
- **文件位置**: 行 20-20
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('./utils').methods`


**Variable**: `compileETag`

- **完整名称**: compileETag
- **文件位置**: 行 21-21
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('./utils').compileETag`


**Variable**: `compileQueryParser`

- **完整名称**: compileQueryParser
- **文件位置**: 行 22-22
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('./utils').compileQueryParser`


**Variable**: `compileTrust`

- **完整名称**: compileTrust
- **文件位置**: 行 23-23
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('./utils').compileTrust`


**Variable**: `resolve`

- **完整名称**: resolve
- **文件位置**: 行 24-24
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:path').resolve`


**Variable**: `once`

- **完整名称**: once
- **文件位置**: 行 25-25
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('once')
var Router = require('router')`


**Variable**: `slice`

- **完整名称**: slice
- **文件位置**: 行 33-33
- **可见性**: private

- **类型**: `unknown`




- **值**: `Array.prototype.slice`


**Variable**: `flatten`

- **完整名称**: flatten
- **文件位置**: 行 34-34
- **可见性**: private

- **类型**: `unknown`




- **值**: `Array.prototype.flat`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 40-40
- **可见性**: private

- **类型**: `unknown`




- **值**: `exports = module.exports = {}`


**Variable**: `trustProxyDefaultSymbol`

- **完整名称**: trustProxyDefaultSymbol
- **文件位置**: 行 47-47
- **可见性**: private

- **类型**: `string`




- **值**: `'@@symbol:trust_proxy_default'`


**Variable**: `router`

- **完整名称**: router
- **文件位置**: 行 60-60
- **可见性**: private

- **类型**: `null/undefined`




- **值**: `null`


**Variable**: `env`

- **完整名称**: env
- **文件位置**: 行 91-91
- **可见性**: private

- **类型**: `unknown`




- **值**: `process.env.NODE_ENV || 'development'`


**Variable**: `done`

- **完整名称**: done
- **文件位置**: 行 154-154
- **可见性**: private

- **类型**: `unknown`




- **值**: `callback || finalhandler(req, res, {
    env: this.get('env'),
    onerror: logerror.bind(this)
  })`


**Variable**: `offset`

- **完整名称**: offset
- **文件位置**: 行 191-191
- **可见性**: private

- **类型**: `number`




- **值**: `0`


**Variable**: `path`

- **完整名称**: path
- **文件位置**: 行 192-192
- **可见性**: private

- **类型**: `string`




- **值**: `'/'`


**Variable**: `arg`

- **完整名称**: arg
- **文件位置**: 行 197-197
- **可见性**: private

- **类型**: `unknown`




- **值**: `fn`


**Variable**: `fns`

- **完整名称**: fns
- **文件位置**: 行 210-210
- **可见性**: private

- **类型**: `unknown`




- **值**: `flatten.call(slice.call(arguments, offset), Infinity)`


**Variable**: `router`

- **完整名称**: router
- **文件位置**: 行 217-217
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.router`


**Variable**: `orig`

- **完整名称**: orig
- **文件位置**: 行 231-231
- **可见性**: private

- **类型**: `unknown`




- **值**: `req.app`


**Variable**: `extension`

- **完整名称**: extension
- **文件位置**: 行 300-300
- **可见性**: private

- **类型**: `unknown`




- **值**: `ext[0] !== '.'
    ? '.' + ext
    : ext`


**Variable**: `i`

- **完整名称**: i
- **文件位置**: 行 324-324
- **可见性**: private

- **类型**: `number`




- **值**: `0`


**Variable**: `route`

- **完整名称**: route
- **文件位置**: 行 478-478
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.route(path)`


**Variable**: `route`

- **完整名称**: route
- **文件位置**: 行 478-478
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.route(path)`


**Variable**: `args`

- **完整名称**: args
- **文件位置**: 行 496-496
- **可见性**: private

- **类型**: `unknown`




- **值**: `slice.call(arguments, 1)`


**Variable**: `i`

- **完整名称**: i
- **文件位置**: 行 324-324
- **可见性**: private

- **类型**: `number`




- **值**: `0`


**Variable**: `cache`

- **完整名称**: cache
- **文件位置**: 行 523-523
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.cache`


**Variable**: `done`

- **完整名称**: done
- **文件位置**: 行 154-154
- **可见性**: private

- **类型**: `unknown`




- **值**: `callback`


**Variable**: `engines`

- **完整名称**: engines
- **文件位置**: 行 525-525
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.engines`


**Variable**: `opts`

- **完整名称**: opts
- **文件位置**: 行 526-526
- **可见性**: private

- **类型**: `unknown`




- **值**: `options`


**Variable**: `renderOptions`

- **完整名称**: renderOptions
- **文件位置**: 行 536-536
- **可见性**: private

- **类型**: `object`




- **值**: `{ ...this.locals, ...opts._locals, ...opts }`


**Variable**: `View`

- **完整名称**: View
- **文件位置**: 行 550-550
- **可见性**: public

- **类型**: `unknown`




- **值**: `this.get('view')`


**Variable**: `dirs`

- **完整名称**: dirs
- **文件位置**: 行 559-559
- **可见性**: private

- **类型**: `unknown`




- **值**: `Array.isArray(view.root) && view.root.length > 1
        ? 'directories "' + view.root.slice(0, -1).join('", "') + '" or "' + view.root[view.root.length - 1] + '"'
        : 'directory "' + view.root + '"'
      var err = new Error('Failed to lookup view "' + name + '" in views ' + dirs)`


**Variable**: `http`

- **完整名称**: http
- **文件位置**: 行 586-586
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:http')
 *      , https = require('node:https')
 *      , express = require('express')
 *      , app = express()`


**Variable**: `server`

- **完整名称**: server
- **文件位置**: 行 599-599
- **可见性**: private

- **类型**: `unknown`




- **值**: `http.createServer(this)
  var args = slice.call(arguments)
  if (typeof args[args.length - 1] === 'function') {
    var done = args[args.length - 1] = once(args[args.length - 1])
    server.once('error', done)
  }
  return server.listen.apply(server, args)
}

/**
 * Log error using console.error.
 *
 * @param {Error} err
 * @private
 */

function logerror(err) {
  /* istanbul ignore next */
  if (this.get('env') !== 'test') console.error(err.stack || err.toString())`





### lib\express.js

- **语言**: javascript
- **包**: 
- **代码行数**: 34
- **元素数量**: 8




#### 代码元素


**Function**: `createApplication`

- **完整名称**: createApplication
- **文件位置**: 行 36-56
- **可见性**: public
- **描述**: Expose the prototypes.








**Variable**: `bodyParser`

- **完整名称**: bodyParser
- **文件位置**: 行 15-15
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('body-parser')
var EventEmitter = require('node:events').EventEmitter`


**Variable**: `mixin`

- **完整名称**: mixin
- **文件位置**: 行 17-17
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('merge-descriptors')`


**Variable**: `proto`

- **完整名称**: proto
- **文件位置**: 行 18-18
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('./application')`


**Variable**: `Router`

- **完整名称**: Router
- **文件位置**: 行 19-19
- **可见性**: public

- **类型**: `unknown`




- **值**: `require('router')`


**Variable**: `req`

- **完整名称**: req
- **文件位置**: 行 20-20
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('./request')`


**Variable**: `res`

- **完整名称**: res
- **文件位置**: 行 21-21
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('./response')`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 37-37
- **可见性**: private

- **类型**: `unknown`




- **值**: `function(req, res, next) {
    app.handle(req, res, next)`





### lib\request.js

- **语言**: javascript
- **包**: 
- **代码行数**: 154
- **元素数量**: 50




#### 代码元素


**Function**: `header`

- **完整名称**: header
- **文件位置**: 行 64-83
- **可见性**: public
- **描述**: Check if the given `type(s)` is acceptable, returning








**Function**: `range`

- **完整名称**: range
- **文件位置**: 行 214-218
- **可见性**: public
- **描述**: Parse the query string of `req.url`.








**Function**: `query`

- **完整名称**: query
- **文件位置**: 行 230-241
- **可见性**: public
- **描述**: Check if the incoming request contains the "Content-Type"








**Function**: `is`

- **完整名称**: is
- **文件位置**: 行 269-281
- **可见性**: public
- **描述**: Return the protocol string "http" or "https"








**Function**: `protocol`

- **完整名称**: protocol
- **文件位置**: 行 297-315
- **可见性**: public
- **描述**: Short-hand for:








**Function**: `secure`

- **完整名称**: secure
- **文件位置**: 行 326-328
- **可见性**: public
- **描述**: Return the remote address from the trusted proxy.








**Function**: `ip`

- **完整名称**: ip
- **文件位置**: 行 340-343
- **可见性**: public
- **描述**: When "trust proxy" is set, trusted proxy addresses + client.








**Function**: `ips`

- **完整名称**: ips
- **文件位置**: 行 357-366
- **可见性**: public
- **描述**: Return subdomains as an array.








**Function**: `subdomains`

- **完整名称**: subdomains
- **文件位置**: 行 383-394
- **可见性**: public
- **描述**: Short-hand for `url.parse(req.url).pathname`.








**Function**: `path`

- **完整名称**: path
- **文件位置**: 行 403-405
- **可见性**: public
- **描述**: Parse the "Host" header field to a host.








**Function**: `host`

- **完整名称**: host
- **文件位置**: 行 418-431
- **可见性**: public
- **描述**: Parse the "Host" header field to a hostname.








**Function**: `hostname`

- **完整名称**: hostname
- **文件位置**: 行 444-458
- **可见性**: public
- **描述**: Check if the request is fresh, aka








**Function**: `stale`

- **完整名称**: stale
- **文件位置**: 行 497-499
- **可见性**: public
- **描述**: Check if the request was an _XMLHttpRequest_.








**Function**: `xhr`

- **完整名称**: xhr
- **文件位置**: 行 508-511
- **可见性**: public
- **描述**: Helper function for creating a getter on an object.








**Function**: `defineGetter`

- **完整名称**: defineGetter
- **文件位置**: 行 521-527
- **可见性**: public









**Variable**: `accepts`

- **完整名称**: accepts
- **文件位置**: 行 16-16
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('accepts')`


**Variable**: `isIP`

- **完整名称**: isIP
- **文件位置**: 行 17-17
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:net').isIP`


**Variable**: `typeis`

- **完整名称**: typeis
- **文件位置**: 行 18-18
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('type-is')`


**Variable**: `http`

- **完整名称**: http
- **文件位置**: 行 19-19
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:http')`


**Variable**: `fresh`

- **完整名称**: fresh
- **文件位置**: 行 20-20
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('fresh')`


**Variable**: `parseRange`

- **完整名称**: parseRange
- **文件位置**: 行 21-21
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('range-parser')`


**Variable**: `parse`

- **完整名称**: parse
- **文件位置**: 行 22-22
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('parseurl')`


**Variable**: `proxyaddr`

- **完整名称**: proxyaddr
- **文件位置**: 行 23-23
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('proxy-addr')`


**Variable**: `req`

- **完整名称**: req
- **文件位置**: 行 30-30
- **可见性**: private

- **类型**: `unknown`




- **值**: `Object.create(http.IncomingMessage.prototype)

/**
 * Module exports.
 * @public
 */

module.exports = req

/**
 * Return request header.
 *
 * The `Referrer` header field is special-cased,
 * both `Referrer` and `Referer` are interchangeable.
 *
 * Examples:
 *
 *     req.get('Content-Type')`


**Variable**: `lc`

- **完整名称**: lc
- **文件位置**: 行 73-73
- **可见性**: private

- **类型**: `unknown`




- **值**: `name.toLowerCase()`


**Variable**: `accept`

- **完整名称**: accept
- **文件位置**: 行 128-128
- **可见性**: private

- **类型**: `unknown`




- **值**: `accepts(this)`


**Variable**: `accept`

- **完整名称**: accept
- **文件位置**: 行 128-128
- **可见性**: private

- **类型**: `unknown`




- **值**: `accepts(this)`


**Variable**: `range`

- **完整名称**: range
- **文件位置**: 行 215-215
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.get('Range')`


**Variable**: `queryparse`

- **完整名称**: queryparse
- **文件位置**: 行 231-231
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.app.get('query parser fn')`


**Variable**: `querystring`

- **完整名称**: querystring
- **文件位置**: 行 238-238
- **可见性**: private

- **类型**: `unknown`




- **值**: `parse(this).query`


**Variable**: `arr`

- **完整名称**: arr
- **文件位置**: 行 270-270
- **可见性**: private

- **类型**: `unknown`




- **值**: `types`


**Variable**: `i`

- **完整名称**: i
- **文件位置**: 行 275-275
- **可见性**: private

- **类型**: `number`




- **值**: `0`


**Variable**: `proto`

- **完整名称**: proto
- **文件位置**: 行 298-298
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.socket.encrypted
    ? 'https'
    : 'http'`


**Variable**: `trust`

- **完整名称**: trust
- **文件位置**: 行 301-301
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.app.get('trust proxy fn')`


**Variable**: `header`

- **完整名称**: header
- **文件位置**: 行 309-309
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.get('X-Forwarded-Proto') || proto
  var index = header.indexOf(',')

  return index !== -1
    ? header.substring(0, index).trim()
    : header.trim()
})`


**Variable**: `trust`

- **完整名称**: trust
- **文件位置**: 行 301-301
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.app.get('trust proxy fn')`


**Variable**: `trust`

- **完整名称**: trust
- **文件位置**: 行 301-301
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.app.get('trust proxy fn')`


**Variable**: `addrs`

- **完整名称**: addrs
- **文件位置**: 行 359-359
- **可见性**: private

- **类型**: `unknown`




- **值**: `proxyaddr.all(this, trust)`


**Variable**: `hostname`

- **完整名称**: hostname
- **文件位置**: 行 384-384
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.hostname`


**Variable**: `offset`

- **完整名称**: offset
- **文件位置**: 行 388-388
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.app.get('subdomain offset')`


**Variable**: `subdomains`

- **完整名称**: subdomains
- **文件位置**: 行 389-389
- **可见性**: private

- **类型**: `unknown`




- **值**: `!isIP(hostname)
    ? hostname.split('.').reverse()
    : [hostname]`


**Variable**: `trust`

- **完整名称**: trust
- **文件位置**: 行 301-301
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.app.get('trust proxy fn')`


**Variable**: `val`

- **完整名称**: val
- **文件位置**: 行 420-420
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.get('X-Forwarded-Host')`


**Variable**: `host`

- **完整名称**: host
- **文件位置**: 行 445-445
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.host`


**Variable**: `offset`

- **完整名称**: offset
- **文件位置**: 行 450-450
- **可见性**: private

- **类型**: `unknown`




- **值**: `host[0] === '['
    ? host.indexOf(']') + 1
    : 0`


**Variable**: `index`

- **完整名称**: index
- **文件位置**: 行 453-453
- **可见性**: private

- **类型**: `unknown`




- **值**: `host.indexOf(':', offset)`


**Variable**: `method`

- **完整名称**: method
- **文件位置**: 行 470-470
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.method`


**Variable**: `res`

- **完整名称**: res
- **文件位置**: 行 471-471
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.res
  var status = res.statusCode

  // GET or HEAD for weak freshness validation only
  if ('GET' !== method && 'HEAD' !== method) return false`


**Variable**: `val`

- **完整名称**: val
- **文件位置**: 行 509-509
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.get('X-Requested-With') || ''`


**Constant**: `accept`

- **完整名称**: accept
- **文件位置**: 行 172-172
- **可见性**: private

- **类型**: `unknown`




- **值**: `accepts(this)`





### lib\response.js

- **语言**: javascript
- **包**: 
- **代码行数**: 475
- **元素数量**: 107




#### 代码元素


**Function**: `status`

- **完整名称**: status
- **文件位置**: 行 64-1048
- **可见性**: public
- **描述**: Set Link header field with the given `links`.








**Function**: `send`

- **完整名称**: send
- **文件位置**: 行 125-218
- **可见性**: public
- **描述**: Send JSON response.








**Function**: `json`

- **完整名称**: json
- **文件位置**: 行 232-246
- **可见性**: public
- **描述**: Send JSON response with JSONP callback support.








**Function**: `jsonp`

- **完整名称**: jsonp
- **文件位置**: 行 260-304
- **可见性**: public









**Function**: `sendStatus`

- **完整名称**: sendStatus
- **文件位置**: 行 321-328
- **可见性**: public
- **描述**: Transfer the file at the given `path`.








**Function**: `sendFile`

- **完整名称**: sendFile
- **文件位置**: 行 371-413
- **可见性**: public
- **描述**: Transfer the file at the given `path` as an attachment.








**Function**: `download`

- **完整名称**: download
- **文件位置**: 行 433-1048
- **可见性**: public
- **描述**: Set _Content-Type_ response header with `type` through `mime.contentType()`








**Function**: `contentType`

- **完整名称**: contentType
- **文件位置**: 行 504-510
- **可见性**: public
- **描述**: Respond to the Acceptable formats using an `obj`








**Function**: `text`

- **完整名称**: text
- **文件位置**: 行 545-547
- **可见性**: public
- **描述**: Set _Content-Disposition_ header to _attachment_ with optional `filename`.








**Function**: `html`

- **完整名称**: html
- **文件位置**: 行 549-551
- **可见性**: public
- **描述**: Set _Content-Disposition_ header to _attachment_ with optional `filename`.








**Function**: `json`

- **完整名称**: json
- **文件位置**: 行 553-555
- **可见性**: public
- **描述**: Set _Content-Disposition_ header to _attachment_ with optional `filename`.








**Function**: `attachment`

- **完整名称**: attachment
- **文件位置**: 行 604-612
- **可见性**: public
- **描述**: Append additional header `field` with value `val`.








**Function**: `append`

- **完整名称**: append
- **文件位置**: 行 629-641
- **可见性**: public
- **描述**: Set header `field` to `val`, or pass








**Function**: `header`

- **完整名称**: header
- **文件位置**: 行 665-686
- **可见性**: public
- **描述**: Get value for header `field`.








**Function**: `clearCookie`

- **完整名称**: clearCookie
- **文件位置**: 行 709-716
- **可见性**: public
- **描述**: Set cookie `name` to `value`, with the given `options`.








**Function**: `location`

- **完整名称**: location
- **文件位置**: 行 794-796
- **可见性**: public
- **描述**: Redirect to the given `url` with optional response `status`








**Function**: `redirect`

- **完整名称**: redirect
- **文件位置**: 行 812-864
- **可见性**: public
- **描述**: Add `field` to Vary. If already present in the Vary set, then








**Function**: `text`

- **完整名称**: text
- **文件位置**: 行 545-547
- **可见性**: public
- **描述**: Set _Content-Disposition_ header to _attachment_ with optional `filename`.








**Function**: `html`

- **完整名称**: html
- **文件位置**: 行 549-551
- **可见性**: public
- **描述**: Set _Content-Disposition_ header to _attachment_ with optional `filename`.








**Function**: `default`

- **完整名称**: default
- **文件位置**: 行 850-852
- **可见性**: public
- **描述**: Add `field` to Vary. If already present in the Vary set, then








**Function**: `render`

- **完整名称**: render
- **文件位置**: 行 894-918
- **可见性**: public
- **描述**: Stringify JSON, like JSON.stringify, but v8 optimized, with the








**Function**: `sendfile`

- **完整名称**: sendfile
- **文件位置**: 行 921-1009
- **可见性**: public
- **描述**: Stringify JSON, like JSON.stringify, but v8 optimized, with the








**Function**: `onaborted`

- **完整名称**: onaborted
- **文件位置**: 行 926-933
- **可见性**: public
- **描述**: Stringify JSON, like JSON.stringify, but v8 optimized, with the








**Function**: `ondirectory`

- **完整名称**: ondirectory
- **文件位置**: 行 936-943
- **可见性**: public
- **描述**: Stringify JSON, like JSON.stringify, but v8 optimized, with the








**Function**: `onerror`

- **完整名称**: onerror
- **文件位置**: 行 946-950
- **可见性**: public
- **描述**: Stringify JSON, like JSON.stringify, but v8 optimized, with the








**Function**: `onend`

- **完整名称**: onend
- **文件位置**: 行 953-957
- **可见性**: public
- **描述**: Stringify JSON, like JSON.stringify, but v8 optimized, with the








**Function**: `onfile`

- **完整名称**: onfile
- **文件位置**: 行 960-962
- **可见性**: public
- **描述**: Stringify JSON, like JSON.stringify, but v8 optimized, with the








**Function**: `onfinish`

- **完整名称**: onfinish
- **文件位置**: 行 965-980
- **可见性**: public
- **描述**: Stringify JSON, like JSON.stringify, but v8 optimized, with the








**Function**: `onstream`

- **完整名称**: onstream
- **文件位置**: 行 983-985
- **可见性**: public
- **描述**: Stringify JSON, like JSON.stringify, but v8 optimized, with the








**Function**: `headers`

- **完整名称**: headers
- **文件位置**: 行 996-1004
- **可见性**: public
- **描述**: Stringify JSON, like JSON.stringify, but v8 optimized, with the








**Function**: `stringify`

- **完整名称**: stringify
- **文件位置**: 行 1023-1047
- **可见性**: public









**Variable**: `contentDisposition`

- **完整名称**: contentDisposition
- **文件位置**: 行 15-15
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('content-disposition')`


**Variable**: `createError`

- **完整名称**: createError
- **文件位置**: 行 16-16
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('http-errors')
var deprecate = require('depd')('express')`


**Variable**: `encodeUrl`

- **完整名称**: encodeUrl
- **文件位置**: 行 18-18
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('encodeurl')`


**Variable**: `escapeHtml`

- **完整名称**: escapeHtml
- **文件位置**: 行 19-19
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('escape-html')`


**Variable**: `http`

- **完整名称**: http
- **文件位置**: 行 20-20
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:http')`


**Variable**: `onFinished`

- **完整名称**: onFinished
- **文件位置**: 行 21-21
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('on-finished')`


**Variable**: `mime`

- **完整名称**: mime
- **文件位置**: 行 22-22
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('mime-types')
var path = require('node:path')`


**Variable**: `pathIsAbsolute`

- **完整名称**: pathIsAbsolute
- **文件位置**: 行 24-24
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:path').isAbsolute`


**Variable**: `statuses`

- **完整名称**: statuses
- **文件位置**: 行 25-25
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('statuses')
var sign = require('cookie-signature').sign`


**Variable**: `normalizeType`

- **完整名称**: normalizeType
- **文件位置**: 行 27-27
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('./utils').normalizeType`


**Variable**: `normalizeTypes`

- **完整名称**: normalizeTypes
- **文件位置**: 行 28-28
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('./utils').normalizeTypes`


**Variable**: `setCharset`

- **完整名称**: setCharset
- **文件位置**: 行 29-29
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('./utils').setCharset`


**Variable**: `cookie`

- **完整名称**: cookie
- **文件位置**: 行 30-30
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('cookie')`


**Variable**: `send`

- **完整名称**: send
- **文件位置**: 行 31-31
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('send')`


**Variable**: `extname`

- **完整名称**: extname
- **文件位置**: 行 32-32
- **可见性**: private

- **类型**: `unknown`




- **值**: `path.extname`


**Variable**: `resolve`

- **完整名称**: resolve
- **文件位置**: 行 33-33
- **可见性**: private

- **类型**: `unknown`




- **值**: `path.resolve`


**Variable**: `vary`

- **完整名称**: vary
- **文件位置**: 行 34-34
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('vary')`


**Variable**: `res`

- **完整名称**: res
- **文件位置**: 行 42-42
- **可见性**: private

- **类型**: `unknown`




- **值**: `Object.create(http.ServerResponse.prototype)

/**
 * Module exports.
 * @public
 */

module.exports = res

/**
 * Set the HTTP status code for the response.
 *
 * Expects an integer value between 100 and 999 inclusive.
 * Throws an error if the provided status code is not an integer or if it's outside the allowable range.
 *
 * @param {number} code - The HTTP status code to set.
 * @return {ServerResponse} - Returns itself for chaining methods.
 * @throws {TypeError} If `code` is not an integer.
 * @throws {RangeError} If `code` is outside the range 100 to 999.
 * @public
 */

res.status = function status(code) {
  // Check if the status code is not an integer
  if (!Number.isInteger(code)) {
    throw new TypeError(`Invalid status code: ${JSON.stringify(code)}. Status code must be an integer.`)`


**Variable**: `link`

- **完整名称**: link
- **文件位置**: 行 98-98
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.get('Link') || ''`


**Variable**: `chunk`

- **完整名称**: chunk
- **文件位置**: 行 126-126
- **可见性**: private

- **类型**: `unknown`




- **值**: `body`


**Variable**: `req`

- **完整名称**: req
- **文件位置**: 行 128-128
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.req`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 131-131
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.app`


**Variable**: `etagFn`

- **完整名称**: etagFn
- **文件位置**: 行 161-161
- **可见性**: private

- **类型**: `unknown`




- **值**: `app.get('etag fn')
  var generateETag = !this.get('ETag') && typeof etagFn === 'function'

  // populate Content-Length
  var len
  if (chunk !== undefined) {
    if (Buffer.isBuffer(chunk)) {
      // get length of Buffer
      len = chunk.length
    } else if (!generateETag && chunk.length < 1000) {
      // just calculate length when no ETag + small chunk
      len = Buffer.byteLength(chunk, encoding)
    } else {
      // convert chunk to Buffer and calculate
      chunk = Buffer.from(chunk, encoding)
      encoding = undefined`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 131-131
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.app`


**Variable**: `escape`

- **完整名称**: escape
- **文件位置**: 行 235-235
- **可见性**: private

- **类型**: `unknown`




- **值**: `app.get('json escape')
  var replacer = app.get('json replacer')`


**Variable**: `spaces`

- **完整名称**: spaces
- **文件位置**: 行 237-237
- **可见性**: private

- **类型**: `unknown`




- **值**: `app.get('json spaces')`


**Variable**: `body`

- **完整名称**: body
- **文件位置**: 行 238-238
- **可见性**: private

- **类型**: `unknown`




- **值**: `stringify(obj, replacer, spaces, escape)

  // content-type
  if (!this.get('Content-Type')) {
    this.set('Content-Type', 'application/json')`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 131-131
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.app`


**Variable**: `escape`

- **完整名称**: escape
- **文件位置**: 行 235-235
- **可见性**: private

- **类型**: `unknown`




- **值**: `app.get('json escape')
  var replacer = app.get('json replacer')`


**Variable**: `spaces`

- **完整名称**: spaces
- **文件位置**: 行 237-237
- **可见性**: private

- **类型**: `unknown`




- **值**: `app.get('json spaces')`


**Variable**: `body`

- **完整名称**: body
- **文件位置**: 行 266-266
- **可见性**: private

- **类型**: `unknown`




- **值**: `stringify(obj, replacer, spaces, escape)
  var callback = this.req.query[app.get('jsonp callback name')]`


**Variable**: `body`

- **完整名称**: body
- **文件位置**: 行 322-322
- **可见性**: private

- **类型**: `unknown`




- **值**: `statuses.message[statusCode] || String(statusCode)

  this.status(statusCode)`


**Variable**: `uid`

- **完整名称**: uid
- **文件位置**: 行 356-356
- **可见性**: private

- **类型**: `unknown`




- **值**: `req.params.uid
 *         , file = req.params.file`


**Variable**: `done`

- **完整名称**: done
- **文件位置**: 行 372-372
- **可见性**: private

- **类型**: `unknown`




- **值**: `callback`


**Variable**: `req`

- **完整名称**: req
- **文件位置**: 行 128-128
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.req`


**Variable**: `res`

- **完整名称**: res
- **文件位置**: 行 374-374
- **可见性**: private

- **类型**: `unknown`




- **值**: `this`


**Variable**: `next`

- **完整名称**: next
- **文件位置**: 行 375-375
- **可见性**: private

- **类型**: `unknown`




- **值**: `req.next`


**Variable**: `opts`

- **完整名称**: opts
- **文件位置**: 行 376-376
- **可见性**: private

- **类型**: `unknown`




- **值**: `options || {}`


**Variable**: `pathname`

- **完整名称**: pathname
- **文件位置**: 行 397-397
- **可见性**: private

- **类型**: `unknown`




- **值**: `encodeURI(path)`


**Variable**: `file`

- **完整名称**: file
- **文件位置**: 行 401-401
- **可见性**: private

- **类型**: `unknown`




- **值**: `send(req, pathname, opts)`


**Variable**: `done`

- **完整名称**: done
- **文件位置**: 行 372-372
- **可见性**: private

- **类型**: `unknown`




- **值**: `callback`


**Variable**: `name`

- **完整名称**: name
- **文件位置**: 行 435-435
- **可见性**: private

- **类型**: `unknown`




- **值**: `filename`


**Variable**: `opts`

- **完整名称**: opts
- **文件位置**: 行 436-436
- **可见性**: private

- **类型**: `unknown`




- **值**: `options || null

  // support function as second or third arg
  if (typeof filename === 'function') {
    done = filename`


**Variable**: `headers`

- **完整名称**: headers
- **文件位置**: 行 456-456
- **可见性**: private

- **类型**: `object`




- **值**: `{
    'Content-Disposition': contentDisposition(name || path)
  }`


**Variable**: `keys`

- **完整名称**: keys
- **文件位置**: 行 462-462
- **可见性**: private

- **类型**: `unknown`




- **值**: `Object.keys(opts.headers)
    for (var i = 0`


**Variable**: `key`

- **完整名称**: key
- **文件位置**: 行 464-464
- **可见性**: private

- **类型**: `unknown`




- **值**: `keys[i]
      if (key.toLowerCase() !== 'content-disposition') {
        headers[key] = opts.headers[key]
      }
    }
  }

  // merge user-provided options
  opts = Object.create(opts)
  opts.headers = headers

  // Resolve the full path for sendFile
  var fullPath = !opts.root
    ? resolve(path)
    : path

  // send file
  return this.sendFile(fullPath, opts, done)
}`


**Variable**: `ct`

- **完整名称**: ct
- **文件位置**: 行 505-505
- **可见性**: private

- **类型**: `unknown`




- **值**: `type.indexOf('/') === -1
    ? (mime.contentType(type) || 'application/octet-stream')
    : type`


**Variable**: `req`

- **完整名称**: req
- **文件位置**: 行 128-128
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.req`


**Variable**: `next`

- **完整名称**: next
- **文件位置**: 行 375-375
- **可见性**: private

- **类型**: `unknown`




- **值**: `req.next`


**Variable**: `keys`

- **完整名称**: keys
- **文件位置**: 行 573-573
- **可见性**: private

- **类型**: `unknown`




- **值**: `Object.keys(obj)
    .filter(function (v) { return v !== 'default' })

  var key = keys.length > 0
    ? req.accepts(keys)
    : false`


**Variable**: `prev`

- **完整名称**: prev
- **文件位置**: 行 630-630
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.get(field)`


**Variable**: `value`

- **完整名称**: value
- **文件位置**: 行 631-631
- **可见性**: private

- **类型**: `unknown`




- **值**: `val`


**Variable**: `value`

- **完整名称**: value
- **文件位置**: 行 667-667
- **可见性**: private

- **类型**: `unknown`




- **值**: `Array.isArray(val)
      ? val.map(String)
      : String(val)`


**Variable**: `opts`

- **完整名称**: opts
- **文件位置**: 行 743-743
- **可见性**: private

- **类型**: `object`




- **值**: `{ ...options }`


**Variable**: `secret`

- **完整名称**: secret
- **文件位置**: 行 744-744
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.req.secret`


**Variable**: `signed`

- **完整名称**: signed
- **文件位置**: 行 745-745
- **可见性**: private

- **类型**: `unknown`




- **值**: `opts.signed`


**Variable**: `val`

- **完整名称**: val
- **文件位置**: 行 751-751
- **可见性**: private

- **类型**: `unknown`




- **值**: `typeof value === 'object'
    ? 'j:' + JSON.stringify(value)
    : String(value)`


**Variable**: `maxAge`

- **完整名称**: maxAge
- **文件位置**: 行 760-760
- **可见性**: private

- **类型**: `unknown`




- **值**: `opts.maxAge - 0

    if (!isNaN(maxAge)) {
      opts.expires = new Date(Date.now() + maxAge)
      opts.maxAge = Math.floor(maxAge / 1000)
    }
  }

  if (opts.path == null) {
    opts.path = '/'`


**Variable**: `address`

- **完整名称**: address
- **文件位置**: 行 813-813
- **可见性**: private

- **类型**: `unknown`




- **值**: `url`


**Variable**: `status`

- **完整名称**: status
- **文件位置**: 行 815-815
- **可见性**: private

- **类型**: `number`




- **值**: `302`


**Variable**: `u`

- **完整名称**: u
- **文件位置**: 行 845-845
- **可见性**: private

- **类型**: `unknown`




- **值**: `escapeHtml(address)`


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 895-895
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.req.app`


**Variable**: `done`

- **完整名称**: done
- **文件位置**: 行 372-372
- **可见性**: private

- **类型**: `unknown`




- **值**: `callback`


**Variable**: `opts`

- **完整名称**: opts
- **文件位置**: 行 376-376
- **可见性**: private

- **类型**: `unknown`




- **值**: `options || {}`


**Variable**: `req`

- **完整名称**: req
- **文件位置**: 行 128-128
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.req`


**Variable**: `self`

- **完整名称**: self
- **文件位置**: 行 899-899
- **可见性**: private

- **类型**: `unknown`




- **值**: `this`


**Variable**: `done`

- **完整名称**: done
- **文件位置**: 行 922-922
- **可见性**: private

- **类型**: `boolean`




- **值**: `false`


**Variable**: `err`

- **完整名称**: err
- **文件位置**: 行 930-930
- **可见性**: private

- **类型**: `unknown`




- **值**: `new Error('Request aborted')`


**Variable**: `err`

- **完整名称**: err
- **文件位置**: 行 940-940
- **可见性**: private

- **类型**: `unknown`




- **值**: `new Error('EISDIR, read')`


**Variable**: `obj`

- **完整名称**: obj
- **文件位置**: 行 997-997
- **可见性**: private

- **类型**: `unknown`




- **值**: `options.headers`


**Variable**: `keys`

- **完整名称**: keys
- **文件位置**: 行 573-573
- **可见性**: private

- **类型**: `unknown`




- **值**: `Object.keys(obj)`


**Variable**: `i`

- **完整名称**: i
- **文件位置**: 行 463-463
- **可见性**: private

- **类型**: `number`




- **值**: `0`


**Variable**: `k`

- **完整名称**: k
- **文件位置**: 行 1001-1001
- **可见性**: private

- **类型**: `unknown`




- **值**: `keys[i]`


**Variable**: `json`

- **完整名称**: json
- **文件位置**: 行 1026-1026
- **可见性**: private

- **类型**: `unknown`




- **值**: `replacer || spaces
    ? JSON.stringify(value, replacer, spaces)
    : JSON.stringify(value)`


**Constant**: `type`

- **完整名称**: type
- **文件位置**: 行 137-137
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.get('Content-Type')`


**Constant**: `opts`

- **完整名称**: opts
- **文件位置**: 行 711-711
- **可见性**: private

- **类型**: `object`




- **值**: `{ path: '/', ...options, expires: new Date(1)}`





### lib\utils.js

- **语言**: javascript
- **包**: 
- **代码行数**: 120
- **元素数量**: 22




#### 代码元素


**Function**: `acceptParams`

- **完整名称**: acceptParams
- **文件位置**: 行 89-120
- **可见性**: public
- **描述**: Compile "etag" value to function.








**Function**: `compileQueryParser`

- **完整名称**: compileQueryParser
- **文件位置**: 行 162-184
- **可见性**: public
- **描述**: Compile "proxy trust" value to function.








**Function**: `setCharset`

- **完整名称**: setCharset
- **文件位置**: 行 225-238
- **可见性**: public
- **描述**: Create an ETag generator function, generating ETags with








**Function**: `createETagGenerator`

- **完整名称**: createETagGenerator
- **文件位置**: 行 249-257
- **可见性**: public
- **描述**: Parse an extended query string with qs.








**Function**: `generateETag`

- **完整名称**: generateETag
- **文件位置**: 行 250-256
- **可见性**: public
- **描述**: Parse an extended query string with qs.








**Function**: `parseExtendedQueryString`

- **完整名称**: parseExtendedQueryString
- **文件位置**: 行 267-271
- **可见性**: public









**Variable**: `contentType`

- **完整名称**: contentType
- **文件位置**: 行 16-16
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('content-type')`


**Variable**: `etag`

- **完整名称**: etag
- **文件位置**: 行 17-17
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('etag')`


**Variable**: `mime`

- **完整名称**: mime
- **文件位置**: 行 18-18
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('mime-types')
var proxyaddr = require('proxy-addr')`


**Variable**: `qs`

- **完整名称**: qs
- **文件位置**: 行 20-20
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('qs')`


**Variable**: `querystring`

- **完整名称**: querystring
- **文件位置**: 行 21-21
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:querystring')`


**Variable**: `length`

- **完整名称**: length
- **文件位置**: 行 90-90
- **可见性**: private

- **类型**: `unknown`




- **值**: `str.length`


**Variable**: `colonIndex`

- **完整名称**: colonIndex
- **文件位置**: 行 91-91
- **可见性**: private

- **类型**: `unknown`




- **值**: `str.indexOf('`


**Variable**: `index`

- **完整名称**: index
- **文件位置**: 行 92-92
- **可见性**: private

- **类型**: `unknown`




- **值**: `colonIndex === -1 ? length : colonIndex`


**Variable**: `ret`

- **完整名称**: ret
- **文件位置**: 行 93-93
- **可见性**: private

- **类型**: `object`




- **值**: `{ value: str.slice(0, index).trim(), quality: 1, params: {} }`


**Variable**: `splitIndex`

- **完整名称**: splitIndex
- **文件位置**: 行 96-96
- **可见性**: private

- **类型**: `unknown`




- **值**: `str.indexOf('=', index)`


**Variable**: `colonIndex`

- **完整名称**: colonIndex
- **文件位置**: 行 91-91
- **可见性**: private

- **类型**: `unknown`




- **值**: `str.indexOf('`


**Variable**: `endIndex`

- **完整名称**: endIndex
- **文件位置**: 行 100-100
- **可见性**: private

- **类型**: `unknown`




- **值**: `colonIndex === -1 ? length : colonIndex`


**Variable**: `key`

- **完整名称**: key
- **文件位置**: 行 107-107
- **可见性**: private

- **类型**: `unknown`




- **值**: `str.slice(index, splitIndex).trim()`


**Variable**: `value`

- **完整名称**: value
- **文件位置**: 行 108-108
- **可见性**: private

- **类型**: `unknown`




- **值**: `str.slice(splitIndex + 1, endIndex).trim()`


**Variable**: `parsed`

- **完整名称**: parsed
- **文件位置**: 行 231-231
- **可见性**: private

- **类型**: `unknown`




- **值**: `contentType.parse(type)`


**Variable**: `buf`

- **完整名称**: buf
- **文件位置**: 行 251-251
- **可见性**: private

- **类型**: `unknown`




- **值**: `!Buffer.isBuffer(body)
      ? Buffer.from(body, encoding)
      : body

    return etag(buf, options)
  }
}

/**
 * Parse an extended query string with qs.
 *
 * @param {String} str
 * @return {Object}
 * @private
 */

function parseExtendedQueryString(str) {
  return qs.parse(str, {
    allowPrototypes: true
  })`





### lib\view.js

- **语言**: javascript
- **包**: 
- **代码行数**: 90
- **元素数量**: 31




#### 代码元素


**Function**: `View`

- **完整名称**: View
- **文件位置**: 行 52-95
- **可见性**: public
- **描述**: Lookup view by the given `name`








**Function**: `lookup`

- **完整名称**: lookup
- **文件位置**: 行 104-123
- **可见性**: public
- **描述**: Render with the given options.








**Function**: `render`

- **完整名称**: render
- **文件位置**: 行 133-159
- **可见性**: public
- **描述**: Resolve the file within the given directory.








**Function**: `onRender`

- **完整名称**: onRender
- **文件位置**: 行 139-156
- **可见性**: public
- **描述**: Resolve the file within the given directory.








**Function**: `renderTick`

- **完整名称**: renderTick
- **文件位置**: 行 153-155
- **可见性**: public
- **描述**: Resolve the file within the given directory.








**Function**: `resolve`

- **完整名称**: resolve
- **文件位置**: 行 169-187
- **可见性**: public
- **描述**: Return a stat, maybe.








**Function**: `tryStat`

- **完整名称**: tryStat
- **文件位置**: 行 197-205
- **可见性**: public









**Variable**: `debug`

- **完整名称**: debug
- **文件位置**: 行 16-16
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('debug')('express:view')`


**Variable**: `path`

- **完整名称**: path
- **文件位置**: 行 17-17
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:path')`


**Variable**: `fs`

- **完整名称**: fs
- **文件位置**: 行 18-18
- **可见性**: private

- **类型**: `unknown`




- **值**: `require('node:fs')`


**Variable**: `dirname`

- **完整名称**: dirname
- **文件位置**: 行 25-25
- **可见性**: private

- **类型**: `unknown`




- **值**: `path.dirname`


**Variable**: `basename`

- **完整名称**: basename
- **文件位置**: 行 26-26
- **可见性**: private

- **类型**: `unknown`




- **值**: `path.basename`


**Variable**: `extname`

- **完整名称**: extname
- **文件位置**: 行 27-27
- **可见性**: private

- **类型**: `unknown`




- **值**: `path.extname`


**Variable**: `join`

- **完整名称**: join
- **文件位置**: 行 28-28
- **可见性**: private

- **类型**: `unknown`




- **值**: `path.join`


**Variable**: `resolve`

- **完整名称**: resolve
- **文件位置**: 行 29-29
- **可见性**: private

- **类型**: `unknown`




- **值**: `path.resolve`


**Variable**: `opts`

- **完整名称**: opts
- **文件位置**: 行 53-53
- **可见性**: private

- **类型**: `unknown`




- **值**: `options || {}`


**Variable**: `fileName`

- **完整名称**: fileName
- **文件位置**: 行 64-64
- **可见性**: private

- **类型**: `unknown`




- **值**: `name`


**Variable**: `mod`

- **完整名称**: mod
- **文件位置**: 行 77-77
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.ext.slice(1)
    debug('require "%s"', mod)

    // default engine export
    var fn = require(mod).__express

    if (typeof fn !== 'function') {
      throw new Error('Module "' + mod + '" does not provide a view engine.')
    }

    opts.engines[this.ext] = fn
  }

  // store loaded engine
  this.engine = opts.engines[this.ext]`


**Variable**: `roots`

- **完整名称**: roots
- **文件位置**: 行 106-106
- **可见性**: private

- **类型**: `array`




- **值**: `[].concat(this.root)`


**Variable**: `i`

- **完整名称**: i
- **文件位置**: 行 110-110
- **可见性**: private

- **类型**: `number`




- **值**: `0`


**Variable**: `root`

- **完整名称**: root
- **文件位置**: 行 111-111
- **可见性**: private

- **类型**: `unknown`




- **值**: `roots[i]`


**Variable**: `loc`

- **完整名称**: loc
- **文件位置**: 行 114-114
- **可见性**: private

- **类型**: `unknown`




- **值**: `resolve(root, name)`


**Variable**: `dir`

- **完整名称**: dir
- **文件位置**: 行 115-115
- **可见性**: private

- **类型**: `unknown`




- **值**: `dirname(loc)`


**Variable**: `file`

- **完整名称**: file
- **文件位置**: 行 116-116
- **可见性**: private

- **类型**: `unknown`




- **值**: `basename(loc)`


**Variable**: `sync`

- **完整名称**: sync
- **文件位置**: 行 134-134
- **可见性**: private

- **类型**: `boolean`




- **值**: `true`


**Variable**: `args`

- **完整名称**: args
- **文件位置**: 行 145-145
- **可见性**: private

- **类型**: `unknown`




- **值**: `new Array(arguments.length)`


**Variable**: `cntx`

- **完整名称**: cntx
- **文件位置**: 行 146-146
- **可见性**: private

- **类型**: `unknown`




- **值**: `this`


**Variable**: `i`

- **完整名称**: i
- **文件位置**: 行 110-110
- **可见性**: private

- **类型**: `number`




- **值**: `0`


**Variable**: `ext`

- **完整名称**: ext
- **文件位置**: 行 170-170
- **可见性**: private

- **类型**: `unknown`




- **值**: `this.ext`


**Variable**: `path`

- **完整名称**: path
- **文件位置**: 行 173-173
- **可见性**: private

- **类型**: `unknown`




- **值**: `join(dir, file)`


**Variable**: `stat`

- **完整名称**: stat
- **文件位置**: 行 174-174
- **可见性**: private

- **类型**: `unknown`




- **值**: `tryStat(path)`






## API 参考

### 公共接口



#### Function: authenticate







**调用**: authenticate() log() fn() hash() function() restrict() next() redirect() get() send() destroy() render() post() sendStatus() regenerate() listen() 


---




#### Function: restrict







**调用**: restrict() next() redirect() get() function() send() destroy() render() post() sendStatus() authenticate() regenerate() listen() log() 


---






















#### Function: html







**调用**: function() send() map() join() json() format() get() listen() log() 


---




#### Function: text







**调用**: function() send() map() join() json() format() get() listen() log() 


---




#### Function: json







**调用**: function() json() format() get() listen() log() 


---




#### Function: format







**调用**: format() function() get() listen() log() 


---




































#### Variable: FILES_DIR








---












#### Function: error







**调用**: error() status() send() get() function() Error() next() nextTick() use() listen() log() 


---












#### Function: html







**调用**: function() render() json() type() send() signature() next() use() status() listen() log() 


---




#### Function: json







**调用**: function() json() type() send() signature() next() use() status() render() listen() log() 


---




#### Function: default







**调用**: function() type() send() signature() next() use() status() render() listen() log() 


---












































































































#### Function: list


GET users online.



```
/**
 * GET users online.
 */
```





**调用**: list() map() function() join() get() last() next() send() listen() log() 


---


























#### Function: index







**调用**: function() send() slice() map() join() resource() get() listen() log() 


---




#### Function: show







**调用**: function() send() slice() map() join() resource() get() listen() log() 


---




#### Function: destroy







**调用**: function() send() slice() map() join() resource() get() listen() log() 


---




#### Function: range







**调用**: function() slice() send() map() join() resource() get() listen() log() 


---


















#### Variable: User








---










#### Function: get







**调用**: function() log() send() escapeHtml() map() listen() 


---




#### Function: list







**调用**: function() send() escapeHtml() map() listen() log() 


---




#### Function: get







**调用**: function() send() escapeHtml() map() listen() log() 


---




#### Function: delete







**调用**: function() send() escapeHtml() map() listen() log() 


---




#### Function: list







**调用**: function() send() escapeHtml() map() listen() log() 


---




#### Function: delete







**调用**: function() send() escapeHtml() map() listen() log() 


---












#### Function: loadUser







**调用**: loadUser() next() Error() andRestrictToSelf() andRestrictTo() function() use() get() redirect() send() delete() listen() log() 


---




#### Function: andRestrictToSelf







**调用**: andRestrictToSelf() next() Error() andRestrictTo() function() use() get() redirect() send() delete() listen() log() 


---




#### Function: andRestrictTo







**调用**: andRestrictTo() function() next() Error() use() get() redirect() send() delete() listen() log() 


---


































































#### Variable: RedisStore








---




























#### Function: GithubView


Render the view.



```
/**
 * Render the view.
 */
```





**调用**: GithubView() extname() set() function() request() setEncoding() on() engine() end() 


---


















#### Variable: GithubView








---










#### Function: ferrets







**调用**: ferrets() next() get() function() count() all() render() filter() users() count2() users2() use() listen() log() 


---




#### Function: count







**调用**: count() function() next() users() all() get() render() filter() count2() users2() use() listen() log() 


---




#### Function: users







**调用**: users() all() function() next() get() render() filter() count2() count() users2() use() listen() log() 


---




#### Function: count2







**调用**: count2() count() function() next() users2() all() filter() get() render() use() listen() log() 


---




#### Function: users2







**调用**: users2() all() function() next() filter() get() render() use() listen() log() 


---








#### Variable: User








---






#### Function: User







**调用**: User() function() nextTick() fn() push() 


---






#### Function: error







**调用**: error() Error() use() function() next() indexOf() get() send() status() listen() log() 


---
























#### Function: init


Initialize application configuration.



```
/**
 * Initialize application configuration.
 * @private
 */
```





**调用**: init() create() defaultConfiguration() defineProperty() getrouter() Router() enabled() enable() set() debug() on() onmount() setPrototypeOf() resolve() handle() finalhandler() get() bind() setHeader() use() isArray() call() TypeError() forEach() function() mounted_app() next() emit() route() engine() renderFile() Error() param() compileETag() compileQueryParser() compileTrust() parent() path() Boolean() disabled() disable() VERB() apply() all() render() View() slice() join() done() tryRender() application() express() createServer() listen() once() logerror() error() toString() catch() callback() 


---




#### Function: getrouter


Initialize application configuration.



```
/**
 * Initialize application configuration.
 * @private
 */
```





**调用**: getrouter() Router() enabled() defaultConfiguration() enable() set() defineProperty() debug() on() onmount() setPrototypeOf() create() resolve() handle() finalhandler() get() bind() setHeader() use() isArray() call() TypeError() forEach() function() mounted_app() next() emit() route() engine() renderFile() Error() param() compileETag() compileQueryParser() compileTrust() parent() path() Boolean() disabled() disable() VERB() apply() all() render() View() slice() join() done() tryRender() application() express() createServer() listen() once() logerror() error() toString() catch() callback() 


---




#### Function: defaultConfiguration


Dispatch a req, res pair into the application. Starts pipeline processing.



```
/**
 * Dispatch a req, res pair into the application. Starts pipeline processing.
 *
 * If no callback is provided, then default error handlers will respond
 * in the event of an error bubbling through the stack.
 *
 * @private
 */
```





**调用**: defaultConfiguration() enable() set() defineProperty() debug() on() onmount() setPrototypeOf() create() resolve() handle() finalhandler() get() bind() enabled() setHeader() use() isArray() call() TypeError() forEach() function() mounted_app() next() emit() route() engine() renderFile() Error() param() compileETag() compileQueryParser() compileTrust() parent() path() Boolean() disabled() disable() VERB() apply() all() render() View() slice() join() done() tryRender() application() express() createServer() listen() once() logerror() error() toString() catch() callback() 


---




#### Function: onmount


Dispatch a req, res pair into the application. Starts pipeline processing.



```
/**
 * Dispatch a req, res pair into the application. Starts pipeline processing.
 *
 * If no callback is provided, then default error handlers will respond
 * in the event of an error bubbling through the stack.
 *
 * @private
 */
```





**调用**: onmount() setPrototypeOf() create() set() resolve() enable() handle() finalhandler() get() bind() enabled() setHeader() use() isArray() call() TypeError() forEach() function() debug() mounted_app() next() emit() route() engine() renderFile() Error() param() compileETag() compileQueryParser() compileTrust() defineProperty() parent() path() Boolean() disabled() disable() VERB() apply() all() render() View() slice() join() done() tryRender() application() express() createServer() listen() once() logerror() error() toString() catch() callback() 


---




#### Function: handle


Proxy `Router#use()` to add middleware to the app router.



```
/**
 * Proxy `Router#use()` to add middleware to the app router.
 * See Router#use() documentation for details.
 *
 * If the _fn_ parameter is an express app, then it will be
 * mounted at the _route_ specified.
 *
 * @public
 */
```





**调用**: handle() finalhandler() get() bind() enabled() setHeader() setPrototypeOf() create() use() isArray() call() TypeError() forEach() function() debug() mounted_app() next() emit() route() engine() renderFile() Error() param() set() compileETag() compileQueryParser() compileTrust() defineProperty() parent() path() enable() Boolean() disabled() disable() VERB() apply() all() render() View() slice() join() done() tryRender() application() express() createServer() listen() once() logerror() error() toString() catch() callback() 


---




#### Function: use


Proxy to the app `Router#route()`



```
/**
 * Proxy to the app `Router#route()`
 * Returns a new `Route` instance for the _path_.
 *
 * Routes are isolated middleware stacks for specific paths.
 * See the Route api docs for details.
 *
 * @public
 */
```





**调用**: use() isArray() call() TypeError() forEach() function() debug() mounted_app() handle() setPrototypeOf() next() emit() route() engine() renderFile() Error() param() set() get() compileETag() compileQueryParser() compileTrust() defineProperty() parent() path() enabled() enable() Boolean() disabled() disable() VERB() apply() all() render() View() slice() join() done() tryRender() application() express() createServer() listen() once() logerror() error() toString() catch() callback() 


---




#### Function: mounted_app


Proxy to the app `Router#route()`



```
/**
 * Proxy to the app `Router#route()`
 * Returns a new `Route` instance for the _path_.
 *
 * Routes are isolated middleware stacks for specific paths.
 * See the Route api docs for details.
 *
 * @public
 */
```





**调用**: mounted_app() handle() function() setPrototypeOf() next() emit() route() engine() renderFile() Error() param() isArray() set() get() debug() compileETag() compileQueryParser() compileTrust() defineProperty() parent() path() enabled() enable() Boolean() disabled() disable() VERB() forEach() apply() call() all() render() View() slice() join() done() tryRender() application() express() createServer() listen() once() logerror() error() toString() catch() callback() 


---




#### Function: route


Register the given template engine callback `fn`



```
/**
 * Register the given template engine callback `fn`
 * as `ext`.
 *
 * By default will `require()` the engine based on the
 * file extension. For example if you try to render
 * a "foo.ejs" file Express will invoke the following internally:
 *
 *     app.engine('ejs', require('ejs').__express);
 *
 * For engines that do not provide `.__express` out of the box,
 * or if you wish to "map" a different extension to the template engine
 * you may use this method. For example mapping the EJS template engine to
 * ".html" files:
 *
 *     app.engine('html', require('ejs').renderFile);
 *
 * In this case EJS provides a `.renderFile()` method with
 * the same signature that Express expects: `(path, options, callback)`,
 * though note that it aliases this method as `ejs.__express` internally
 * so if you're using ".ejs" extensions you don't need to do anything.
 *
 * Some template engines do not follow this convention, the
 * [Consolidate.js](https://github.com/tj/consolidate.js)
 * library was created to map all of node's popular template
 * engines to follow this convention, thus allowing them to
 * work seamlessly within Express.
 *
 * @param {String} ext
 * @param {Function} fn
 * @return {app} for chaining
 * @public
 */
```





**调用**: route() engine() renderFile() Error() param() isArray() set() get() debug() compileETag() compileQueryParser() compileTrust() defineProperty() parent() path() enabled() enable() Boolean() disabled() disable() VERB() forEach() function() apply() call() all() render() View() slice() join() done() tryRender() application() express() createServer() listen() once() logerror() error() toString() catch() callback() 


---




#### Function: engine


Proxy to `Router#param()` with one added api feature. The _name_ parameter



```
/**
 * Proxy to `Router#param()` with one added api feature. The _name_ parameter
 * can be an array of names.
 *
 * See the Router#param() docs for more details.
 *
 * @param {String|Array} name
 * @param {Function} fn
 * @return {app} for chaining
 * @public
 */
```





**调用**: engine() Error() param() isArray() set() get() debug() compileETag() compileQueryParser() compileTrust() defineProperty() parent() path() enabled() enable() Boolean() disabled() disable() VERB() forEach() function() route() apply() call() all() render() View() slice() join() done() tryRender() application() express() createServer() listen() once() logerror() error() toString() catch() callback() 


---




#### Function: param


Assign `setting` to `val`, or return `setting`'s value.



```
/**
 * Assign `setting` to `val`, or return `setting`'s value.
 *
 *    app.set('foo', 'bar');
 *    app.set('foo');
 *    // => "bar"
 *
 * Mounted servers inherit their parent server's settings.
 *
 * @param {String} setting
 * @param {*} [val]
 * @return {Server} for chaining
 * @public
 */
```





**调用**: param() isArray() set() get() debug() compileETag() compileQueryParser() compileTrust() defineProperty() parent() path() enabled() enable() Boolean() disabled() disable() VERB() forEach() function() route() apply() call() all() render() View() slice() join() Error() done() tryRender() application() express() createServer() listen() once() logerror() error() toString() catch() callback() 


---




#### Function: set


Return the app's absolute pathname



```
/**
 * Return the app's absolute pathname
 * based on the parent(s) that have
 * mounted it.
 *
 * For example if the application was
 * mounted as "/admin", which itself
 * was mounted as "/blog" then the
 * return value would be "/blog/admin".
 *
 * @return {String}
 * @private
 */
```





**调用**: set() get() debug() compileETag() compileQueryParser() compileTrust() defineProperty() parent() path() enabled() enable() Boolean() disabled() disable() VERB() forEach() function() route() apply() call() all() render() View() isArray() slice() join() Error() done() tryRender() application() express() createServer() listen() once() logerror() error() toString() catch() callback() 


---




#### Function: path


Check if `setting` is enabled (truthy).



```
/**
 * Check if `setting` is enabled (truthy).
 *
 *    app.enabled('foo')
 *    // => false
 *
 *    app.enable('foo')
 *    app.enabled('foo')
 *    // => true
 *
 * @param {String} setting
 * @return {Boolean}
 * @public
 */
```





**调用**: path() enabled() enable() Boolean() set() disabled() disable() VERB() forEach() function() get() route() apply() call() all() render() View() isArray() slice() join() Error() done() tryRender() application() express() createServer() listen() once() logerror() error() toString() catch() callback() 


---




#### Function: enabled


Check if `setting` is disabled.



```
/**
 * Check if `setting` is disabled.
 *
 *    app.disabled('foo')
 *    // => true
 *
 *    app.enable('foo')
 *    app.disabled('foo')
 *    // => false
 *
 * @param {String} setting
 * @return {Boolean}
 * @public
 */
```





**调用**: enabled() Boolean() set() disabled() enable() disable() VERB() forEach() function() get() route() apply() call() all() render() View() isArray() slice() join() Error() done() tryRender() application() express() createServer() listen() once() logerror() error() toString() catch() callback() 


---




#### Function: disabled


Enable `setting`.



```
/**
 * Enable `setting`.
 *
 * @param {String} setting
 * @return {app} for chaining
 * @public
 */
```





**调用**: disabled() set() enable() disable() VERB() forEach() function() get() route() apply() call() all() render() enabled() View() isArray() slice() join() Error() done() tryRender() application() express() createServer() listen() once() logerror() error() toString() catch() callback() 


---




#### Function: enable


Disable `setting`.



```
/**
 * Disable `setting`.
 *
 * @param {String} setting
 * @return {app} for chaining
 * @public
 */
```





**调用**: enable() set() disable() VERB() forEach() function() get() route() apply() call() all() render() enabled() View() isArray() slice() join() Error() done() tryRender() application() express() createServer() listen() once() logerror() error() toString() catch() callback() 


---




#### Function: disable


Delegate `.VERB(...)` calls to `router.VERB(...)`.



```
/**
 * Delegate `.VERB(...)` calls to `router.VERB(...)`.
 */
```





**调用**: disable() set() VERB() forEach() function() get() route() apply() call() all() render() enabled() View() isArray() slice() join() Error() done() tryRender() application() express() createServer() listen() once() logerror() error() toString() catch() callback() 


---




#### Function: all


Render the given view `name` name with `options`



```
/**
 * Render the given view `name` name with `options`
 * and a callback accepting an error and the
 * rendered template string.
 *
 * Example:
 *
 *    app.render('email', { name: 'Tobi' }, function(err, html){
 *      // ...
 *    })
 *
 * @param {String} name
 * @param {Object|Function} options or fn
 * @param {Function} callback
 * @public
 */
```





**调用**: all() route() call() apply() render() function() enabled() get() View() isArray() slice() join() Error() done() tryRender() application() express() createServer() listen() once() logerror() error() toString() catch() callback() 


---




#### Function: render


Listen for connections.



```
/**
 * Listen for connections.
 *
 * A node `http.Server` is returned, with this
 * application (which is a `Function`) as its
 * callback. If you wish to create both an HTTP
 * and HTTPS server you may do so with the "http"
 * and "https" modules as shown here:
 *
 *    var http = require('node:http')
 *      , https = require('node:https')
 *      , express = require('express')
 *      , app = express();
 *
 *    http.createServer(app).listen(80);
 *    https.createServer({ ... }, app).listen(443);
 *
 * @return {http.Server}
 * @public
 */
```





**调用**: render() enabled() get() View() isArray() slice() join() Error() done() tryRender() application() express() createServer() listen() call() once() apply() logerror() error() toString() catch() callback() 


---




#### Function: listen


Log error using console.error.



```
/**
 * Log error using console.error.
 *
 * @param {Error} err
 * @private
 */
```





**调用**: listen() createServer() call() once() apply() logerror() get() error() toString() tryRender() render() catch() callback() 


---




#### Function: logerror


Try rendering a view.



```
/**
 * Try rendering a view.
 * @private
 */
```





**调用**: logerror() get() error() toString() tryRender() render() catch() callback() 


---




#### Function: tryRender







**调用**: tryRender() render() catch() callback() 


---








#### Variable: View








---


































































#### Variable: View








---










#### Function: createApplication


Expose the prototypes.



```
/**
 * Expose the prototypes.
 */
```





**调用**: createApplication() function() handle() mixin() create() init() 


---










#### Variable: Router








---










#### Function: header


Check if the given `type(s)` is acceptable, returning



```
/**
 * Check if the given `type(s)` is acceptable, returning
 * the best match when true, otherwise `false`, in which
 * case you should respond with 406 "Not Acceptable".
 *
 * The `type` value may be a single MIME type string
 * such as "application/json", an extension name
 * such as "json", an argument list such as `"json", "html", "text/plain"`,
 * or an array `["json", "html", "text/plain"]`. When a list
 * or array is given, the _best_ match, if any is returned.
 *
 * Examples:
 *
 *     // Accept: text/html
 *     req.accepts('html');
 *     // => "html"
 *
 *     // Accept: text/*, application/json
 *     req.accepts('html');
 *     // => "html"
 *     req.accepts('text/html');
 *     // => "text/html"
 *     req.accepts('json', 'text');
 *     // => "json"
 *     req.accepts('application/json');
 *     // => "application/json"
 *
 *     // Accept: text/*, application/json
 *     req.accepts('image/png');
 *     req.accepts('png');
 *     // => false
 *
 *     // Accept: text/*;q=.5, application/json
 *     req.accepts(['html', 'json']);
 *     req.accepts('html', 'json');
 *     // => "json"
 *
 * @param {String|Array} type(s)
 * @return {String|Array|Boolean}
 * @public
 */
```





**调用**: header() TypeError() toLowerCase() type() accepts() function() apply() argument() string() arguments() charsets() acceptsCharsets() charset() languages() required() range() get() parseRange() defineGetter() query() create() parse() queryparse() is() isArray() Array() typeis() protocol() trust() indexOf() substring() trim() secure() ip() proxyaddr() ips() all() order() reverse() pop() subdomains() isIP() split() slice() path() host() trimRight() hostname() fresh() stale() xhr() defineProperty() 


---




#### Function: range


Parse the query string of `req.url`.



```
/**
 * Parse the query string of `req.url`.
 *
 * This uses the "query parser" setting to parse the raw
 * string into an object.
 *
 * @return {String}
 * @api public
 */
```





**调用**: range() get() parseRange() defineGetter() query() create() parse() queryparse() is() isArray() Array() typeis() protocol() trust() indexOf() substring() trim() secure() ip() proxyaddr() ips() all() order() reverse() pop() subdomains() isIP() split() slice() path() host() trimRight() hostname() function() fresh() stale() xhr() toLowerCase() defineProperty() 


---




#### Function: query


Check if the incoming request contains the "Content-Type"



```
/**
 * Check if the incoming request contains the "Content-Type"
 * header field, and it contains the given mime `type`.
 *
 * Examples:
 *
 *      // With Content-Type: text/html; charset=utf-8
 *      req.is('html');
 *      req.is('text/html');
 *      req.is('text/*');
 *      // => true
 *
 *      // When Content-Type is application/json
 *      req.is('json');
 *      req.is('application/json');
 *      req.is('application/*');
 *      // => true
 *
 *      req.is('html');
 *      // => false
 *
 * @param {String|Array} types...
 * @return {String|false|null}
 * @public
 */
```





**调用**: query() get() create() parse() queryparse() is() isArray() Array() typeis() defineGetter() protocol() trust() indexOf() substring() trim() secure() ip() proxyaddr() ips() all() order() reverse() pop() subdomains() isIP() split() slice() path() host() trimRight() hostname() function() fresh() stale() xhr() toLowerCase() defineProperty() 


---




#### Function: is


Return the protocol string "http" or "https"



```
/**
 * Return the protocol string "http" or "https"
 * when requested with TLS. When the "trust proxy"
 * setting trusts the socket address, the
 * "X-Forwarded-Proto" header field will be trusted
 * and used if present.
 *
 * If you're running behind a reverse proxy that
 * supplies https for you this may be enabled.
 *
 * @return {String}
 * @public
 */
```





**调用**: is() isArray() Array() typeis() defineGetter() protocol() get() trust() indexOf() substring() trim() secure() ip() proxyaddr() ips() all() order() reverse() pop() subdomains() isIP() split() slice() parse() path() host() trimRight() hostname() function() fresh() stale() xhr() toLowerCase() defineProperty() 


---




#### Function: protocol


Short-hand for:



```
/**
 * Short-hand for:
 *
 *    req.protocol === 'https'
 *
 * @return {Boolean}
 * @public
 */
```





**调用**: protocol() get() trust() indexOf() substring() trim() defineGetter() secure() ip() proxyaddr() ips() all() order() reverse() pop() subdomains() isIP() split() slice() parse() path() host() trimRight() hostname() function() fresh() stale() xhr() toLowerCase() defineProperty() 


---




#### Function: secure


Return the remote address from the trusted proxy.



```
/**
 * Return the remote address from the trusted proxy.
 *
 * The is the remote address on the socket unless
 * "trust proxy" is set.
 *
 * @return {String}
 * @public
 */
```





**调用**: secure() defineGetter() ip() get() proxyaddr() ips() all() order() reverse() pop() subdomains() isIP() split() slice() parse() path() host() trust() indexOf() substring() trimRight() hostname() function() fresh() stale() xhr() toLowerCase() defineProperty() 


---




#### Function: ip


When "trust proxy" is set, trusted proxy addresses + client.



```
/**
 * When "trust proxy" is set, trusted proxy addresses + client.
 *
 * For example if the value were "client, proxy1, proxy2"
 * you would receive the array `["client", "proxy1", "proxy2"]`
 * where "proxy2" is the furthest down-stream and "proxy1" and
 * "proxy2" were trusted.
 *
 * @return {Array}
 * @public
 */
```





**调用**: ip() get() proxyaddr() defineGetter() ips() all() order() reverse() pop() subdomains() isIP() split() slice() parse() path() host() trust() indexOf() substring() trimRight() hostname() function() fresh() stale() xhr() toLowerCase() defineProperty() 


---




#### Function: ips


Return subdomains as an array.



```
/**
 * Return subdomains as an array.
 *
 * Subdomains are the dot-separated parts of the host before the main domain of
 * the app. By default, the domain of the app is assumed to be the last two
 * parts of the host. This can be changed by setting "subdomain offset".
 *
 * For example, if the domain is "tobi.ferrets.example.com":
 * If "subdomain offset" is not set, req.subdomains is `["ferrets", "tobi"]`.
 * If "subdomain offset" is 3, req.subdomains is `["tobi"]`.
 *
 * @return {Array}
 * @public
 */
```





**调用**: ips() get() all() order() reverse() pop() defineGetter() subdomains() isIP() split() slice() parse() path() host() trust() indexOf() substring() trimRight() hostname() function() fresh() stale() xhr() toLowerCase() defineProperty() 


---




#### Function: subdomains


Short-hand for `url.parse(req.url).pathname`.



```
/**
 * Short-hand for `url.parse(req.url).pathname`.
 *
 * @return {String}
 * @public
 */
```





**调用**: subdomains() get() isIP() split() reverse() slice() parse() defineGetter() path() host() trust() indexOf() substring() trimRight() hostname() function() fresh() stale() xhr() toLowerCase() defineProperty() 


---




#### Function: path


Parse the "Host" header field to a host.



```
/**
 * Parse the "Host" header field to a host.
 *
 * When the "trust proxy" setting trusts the socket
 * address, the "X-Forwarded-Host" header field will
 * be trusted.
 *
 * @return {String}
 * @public
 */
```





**调用**: path() parse() defineGetter() host() get() trust() indexOf() substring() trimRight() hostname() function() fresh() stale() xhr() toLowerCase() defineProperty() 


---




#### Function: host


Parse the "Host" header field to a hostname.



```
/**
 * Parse the "Host" header field to a hostname.
 *
 * When the "trust proxy" setting trusts the socket
 * address, the "X-Forwarded-Host" header field will
 * be trusted.
 *
 * @return {String}
 * @api public
 */
```





**调用**: host() get() trust() indexOf() substring() trimRight() defineGetter() hostname() function() fresh() stale() xhr() toLowerCase() defineProperty() 


---




#### Function: hostname


Check if the request is fresh, aka



```
/**
 * Check if the request is fresh, aka
 * Last-Modified or the ETag
 * still match.
 *
 * @return {Boolean}
 * @public
 */
```





**调用**: hostname() indexOf() substring() defineGetter() function() fresh() get() stale() xhr() toLowerCase() defineProperty() 


---




#### Function: stale


Check if the request was an _XMLHttpRequest_.



```
/**
 * Check if the request was an _XMLHttpRequest_.
 *
 * @return {Boolean}
 * @public
 */
```





**调用**: stale() defineGetter() xhr() get() toLowerCase() defineProperty() 


---




#### Function: xhr


Helper function for creating a getter on an object.



```
/**
 * Helper function for creating a getter on an object.
 *
 * @param {Object} obj
 * @param {String} name
 * @param {Function} getter
 * @private
 */
```





**调用**: xhr() get() toLowerCase() defineGetter() defineProperty() 


---




#### Function: defineGetter







**调用**: defineGetter() defineProperty() 


---










































































#### Function: status


Set Link header field with the given `links`.



```
/**
 * Set Link header field with the given `links`.
 *
 * Examples:
 *
 *    res.links({
 *      next: 'http://api.example.com/users?page=2',
 *      last: 'http://api.example.com/users?page=5',
 *      pages: [
 *        'http://api.example.com/users?page=1',
 *        'http://api.example.com/users?page=2'
 *      ]
 *    });
 *
 * @param {Object} links
 * @return {ServerResponse}
 * @public
 */
```





**调用**: status() isInteger() TypeError() stringify() RangeError() links() function() get() set() keys() map() isArray() join() send() from() setCharset() type() isView() json() isBuffer() byteLength() etagFn() removeHeader() end() jsonp() replace() sendStatus() String() callback() 0() sendFile() static() mayViewFilesFrom() pathIsAbsolute() encodeURI() enabled() sendfile() done() next() download() contentDisposition() toLowerCase() create() resolve() contentType() indexOf() format() filter() accepts() vary() normalizeType() default() createError() normalizeTypes() attachment() extname() append() concat() header() setHeader() getHeader() clearCookie() Date() cookie() now() Error() cookieParser() sign() isNaN() floor() serialize() location() encodeUrl() redirect() deprecate() escapeHtml() render() onaborted() ondirectory() onerror() onend() onfile() onfinish() setImmediate() onstream() on() onFinished() headers() pipe() charCodeAt() 


---




#### Function: send


Send JSON response.



```
/**
 * Send JSON response.
 *
 * Examples:
 *
 *     res.json(null);
 *     res.json({ user: 'tj' });
 *
 * @param {string|number|boolean|object} obj
 * @public
 */
```





**调用**: send() get() set() setCharset() type() isView() json() isBuffer() byteLength() from() etagFn() status() removeHeader() end() stringify() jsonp() isArray() replace() sendStatus() String() callback() 0() sendFile() static() function() mayViewFilesFrom() TypeError() pathIsAbsolute() encodeURI() enabled() sendfile() done() next() download() contentDisposition() keys() toLowerCase() create() resolve() contentType() indexOf() format() filter() accepts() vary() normalizeType() default() createError() normalizeTypes() map() attachment() extname() append() concat() header() setHeader() getHeader() clearCookie() Date() cookie() now() Error() cookieParser() sign() isNaN() floor() serialize() location() encodeUrl() redirect() deprecate() escapeHtml() render() onaborted() ondirectory() onerror() onend() onfile() onfinish() setImmediate() onstream() on() onFinished() headers() pipe() charCodeAt() 


---




#### Function: json


Send JSON response with JSONP callback support.



```
/**
 * Send JSON response with JSONP callback support.
 *
 * Examples:
 *
 *     res.jsonp(null);
 *     res.jsonp({ user: 'tj' });
 *
 * @param {string|number|boolean|object} obj
 * @public
 */
```





**调用**: json() get() stringify() set() send() jsonp() isArray() replace() sendStatus() String() status() type() callback() 0() sendFile() static() function() mayViewFilesFrom() TypeError() pathIsAbsolute() encodeURI() enabled() sendfile() done() next() download() contentDisposition() keys() toLowerCase() create() resolve() contentType() indexOf() format() filter() accepts() vary() normalizeType() default() createError() normalizeTypes() map() attachment() extname() append() concat() header() setHeader() getHeader() clearCookie() Date() cookie() now() Error() cookieParser() sign() isNaN() floor() serialize() location() encodeUrl() redirect() deprecate() escapeHtml() byteLength() end() render() onaborted() ondirectory() onerror() onend() onfile() onfinish() setImmediate() onstream() on() onFinished() headers() pipe() charCodeAt() 


---




#### Function: jsonp




```
/**/ is a specific security mitigation for "Rosetta Flash JSONP abuse"
    // the typeof check is just to reduce client error noise
    body = '/**/
```





**调用**: jsonp() get() stringify() set() isArray() replace() send() sendStatus() String() status() type() callback() 0() sendFile() static() function() mayViewFilesFrom() TypeError() pathIsAbsolute() encodeURI() enabled() sendfile() done() next() download() contentDisposition() keys() toLowerCase() create() resolve() contentType() indexOf() format() filter() accepts() vary() normalizeType() default() createError() normalizeTypes() map() attachment() extname() append() concat() header() setHeader() getHeader() clearCookie() Date() cookie() now() Error() cookieParser() sign() isNaN() floor() serialize() location() encodeUrl() redirect() deprecate() escapeHtml() byteLength() end() render() onaborted() ondirectory() onerror() onend() onfile() onfinish() setImmediate() onstream() on() onFinished() headers() pipe() charCodeAt() 


---




#### Function: sendStatus


Transfer the file at the given `path`.



```
/**
 * Transfer the file at the given `path`.
 *
 * Automatically sets the _Content-Type_ response header field.
 * The callback `callback(err)` is invoked when the transfer is complete
 * or when an error occurs. Be sure to check `res.headersSent`
 * if you wish to attempt responding, as the header and some data
 * may have already been transferred.
 *
 * Options:
 *
 *   - `maxAge`   defaulting to 0 (can be string converted by `ms`)
 *   - `root`     root directory for relative filenames
 *   - `headers`  object of headers to serve with file
 *   - `dotfiles` serve dotfiles, defaulting to false; can be `"allow"` to send them
 *
 * Other options are passed along to `send`.
 *
 * Examples:
 *
 *  The following example illustrates how `res.sendFile()` may
 *  be used as an alternative for the `static()` middleware for
 *  dynamic situations. The code backing `res.sendFile()` is actually
 *  the same code, so HTTP cache support etc is identical.
 *
 *     app.get('/user/:uid/photos/:file', function(req, res){
 *       var uid = req.params.uid
 *         , file = req.params.file;
 *
 *       req.user.mayViewFilesFrom(uid, function(yes){
 *         if (yes) {
 *           res.sendFile('/uploads/' + uid + '/' + file);
 *         } else {
 *           res.send(403, 'Sorry! you cant see that.');
 *         }
 *       });
 *     });
 *
 * @public
 */
```





**调用**: sendStatus() String() status() type() send() callback() 0() sendFile() static() get() function() mayViewFilesFrom() TypeError() pathIsAbsolute() encodeURI() enabled() sendfile() done() next() download() contentDisposition() keys() toLowerCase() create() resolve() contentType() indexOf() set() format() filter() accepts() vary() normalizeType() default() createError() normalizeTypes() map() attachment() extname() append() isArray() concat() header() setHeader() getHeader() clearCookie() Date() cookie() now() Error() cookieParser() stringify() sign() isNaN() floor() serialize() location() encodeUrl() redirect() deprecate() escapeHtml() byteLength() end() render() onaborted() ondirectory() onerror() onend() onfile() onfinish() setImmediate() onstream() on() onFinished() headers() pipe() replace() charCodeAt() 


---




#### Function: sendFile


Transfer the file at the given `path` as an attachment.



```
/**
 * Transfer the file at the given `path` as an attachment.
 *
 * Optionally providing an alternate attachment `filename`,
 * and optional callback `callback(err)`. The callback is invoked
 * when the data transfer is complete, or when an error has
 * occurred. Be sure to check `res.headersSent` if you plan to respond.
 *
 * Optionally providing an `options` object to use with `res.sendFile()`.
 * This function will set the `Content-Disposition` header, overriding
 * any `Content-Disposition` header passed as header options in order
 * to set the attachment and filename.
 *
 * This method uses `res.sendFile()`.
 *
 * @public
 */
```





**调用**: sendFile() TypeError() pathIsAbsolute() encodeURI() enabled() send() sendfile() function() done() next() callback() download() contentDisposition() keys() toLowerCase() create() resolve() contentType() type() indexOf() set() format() filter() accepts() vary() normalizeType() default() createError() normalizeTypes() map() attachment() extname() append() get() isArray() concat() header() String() setHeader() getHeader() clearCookie() Date() cookie() now() Error() cookieParser() stringify() sign() isNaN() floor() serialize() location() encodeUrl() redirect() deprecate() escapeHtml() status() byteLength() end() render() onaborted() ondirectory() onerror() onend() onfile() onfinish() setImmediate() onstream() on() onFinished() headers() pipe() replace() charCodeAt() 


---




#### Function: download


Set _Content-Type_ response header with `type` through `mime.contentType()`



```
/**
 * Set _Content-Type_ response header with `type` through `mime.contentType()`
 * when it does not contain "/", or set the Content-Type to `type` otherwise.
 * When no mapping is found though `mime.contentType()`, the type is set to
 * "application/octet-stream".
 *
 * Examples:
 *
 *     res.type('.html');
 *     res.type('html');
 *     res.type('json');
 *     res.type('application/json');
 *     res.type('png');
 *
 * @param {String} type
 * @return {ServerResponse} for chaining
 * @public
 */
```





**调用**: download() contentDisposition() keys() toLowerCase() create() resolve() sendFile() contentType() type() indexOf() set() format() function() send() next() filter() accepts() vary() normalizeType() default() createError() normalizeTypes() map() attachment() extname() append() get() isArray() concat() header() String() TypeError() setHeader() getHeader() clearCookie() Date() cookie() now() Error() cookieParser() stringify() sign() isNaN() floor() serialize() location() encodeUrl() redirect() deprecate() escapeHtml() status() byteLength() end() render() sendfile() onaborted() callback() ondirectory() onerror() onend() onfile() onfinish() setImmediate() onstream() on() onFinished() headers() pipe() replace() charCodeAt() 


---




#### Function: contentType


Respond to the Acceptable formats using an `obj`



```
/**
 * Respond to the Acceptable formats using an `obj`
 * of mime-type callbacks.
 *
 * This method uses `req.accepted`, an array of
 * acceptable types ordered by their quality values.
 * When "Accept" is not present the _first_ callback
 * is invoked, otherwise the first match is used. When
 * no match is performed the server responds with
 * 406 "Not Acceptable".
 *
 * Content-Type is set for you, however if you choose
 * you may alter this within the callback using `res.type()`
 * or `res.set('Content-Type', ...)`.
 *
 *    res.format({
 *      'text/plain': function(){
 *        res.send('hey');
 *      },
 *
 *      'text/html': function(){
 *        res.send('<p>hey</p>');
 *      },
 *
 *      'application/json': function () {
 *        res.send({ message: 'hey' });
 *      }
 *    });
 *
 * In addition to canonicalized MIME types you may
 * also use extnames mapped to these types:
 *
 *    res.format({
 *      text: function(){
 *        res.send('hey');
 *      },
 *
 *      html: function(){
 *        res.send('<p>hey</p>');
 *      },
 *
 *      json: function(){
 *        res.send({ message: 'hey' });
 *      }
 *    });
 *
 * By default Express passes an `Error`
 * with a `.status` of 406 to `next(err)`
 * if a match is not made. If you provide
 * a `.default` callback it will be invoked
 * instead.
 *
 * @param {Object} obj
 * @return {ServerResponse} for chaining
 * @public
 */
```





**调用**: contentType() indexOf() set() type() format() function() send() next() keys() filter() accepts() vary() normalizeType() default() createError() normalizeTypes() map() attachment() extname() contentDisposition() append() get() isArray() concat() header() String() toLowerCase() TypeError() setHeader() getHeader() clearCookie() Date() cookie() now() Error() cookieParser() stringify() sign() isNaN() floor() serialize() location() encodeUrl() redirect() deprecate() escapeHtml() status() byteLength() end() render() sendfile() onaborted() callback() ondirectory() onerror() onend() onfile() onfinish() setImmediate() onstream() on() onFinished() headers() pipe() replace() charCodeAt() 


---




#### Function: text


Set _Content-Disposition_ header to _attachment_ with optional `filename`.



```
/**
 * Set _Content-Disposition_ header to _attachment_ with optional `filename`.
 *
 * @param {String} filename
 * @return {ServerResponse}
 * @public
 */
```





**调用**: function() send() next() keys() filter() accepts() vary() set() normalizeType() default() createError() normalizeTypes() map() attachment() type() extname() contentDisposition() append() get() isArray() concat() header() contentType() String() toLowerCase() TypeError() setHeader() getHeader() clearCookie() Date() cookie() now() Error() cookieParser() stringify() sign() isNaN() floor() serialize() location() encodeUrl() redirect() deprecate() format() escapeHtml() status() byteLength() end() render() sendfile() onaborted() callback() ondirectory() onerror() onend() onfile() onfinish() setImmediate() onstream() on() onFinished() headers() pipe() replace() charCodeAt() 


---




#### Function: html


Set _Content-Disposition_ header to _attachment_ with optional `filename`.



```
/**
 * Set _Content-Disposition_ header to _attachment_ with optional `filename`.
 *
 * @param {String} filename
 * @return {ServerResponse}
 * @public
 */
```





**调用**: function() send() next() keys() filter() accepts() vary() set() normalizeType() default() createError() normalizeTypes() map() attachment() type() extname() contentDisposition() append() get() isArray() concat() header() contentType() String() toLowerCase() TypeError() setHeader() getHeader() clearCookie() Date() cookie() now() Error() cookieParser() stringify() sign() isNaN() floor() serialize() location() encodeUrl() redirect() deprecate() format() escapeHtml() status() byteLength() end() render() sendfile() onaborted() callback() ondirectory() onerror() onend() onfile() onfinish() setImmediate() onstream() on() onFinished() headers() pipe() replace() charCodeAt() 


---




#### Function: json


Set _Content-Disposition_ header to _attachment_ with optional `filename`.



```
/**
 * Set _Content-Disposition_ header to _attachment_ with optional `filename`.
 *
 * @param {String} filename
 * @return {ServerResponse}
 * @public
 */
```





**调用**: function() send() next() keys() filter() accepts() vary() set() normalizeType() default() createError() normalizeTypes() map() attachment() type() extname() contentDisposition() append() get() isArray() concat() header() contentType() String() toLowerCase() TypeError() setHeader() getHeader() clearCookie() Date() cookie() now() Error() cookieParser() stringify() sign() isNaN() floor() serialize() location() encodeUrl() redirect() deprecate() format() escapeHtml() status() byteLength() end() render() sendfile() onaborted() callback() ondirectory() onerror() onend() onfile() onfinish() setImmediate() onstream() on() onFinished() headers() pipe() replace() charCodeAt() 


---




#### Function: attachment


Append additional header `field` with value `val`.



```
/**
 * Append additional header `field` with value `val`.
 *
 * Example:
 *
 *    res.append('Link', ['<http://localhost/>', '<http://localhost:3000/>']);
 *    res.append('Set-Cookie', 'foo=bar; Path=/; HttpOnly');
 *    res.append('Warning', '199 Miscellaneous warning');
 *
 * @param {String} field
 * @param {String|Array} val
 * @return {ServerResponse} for chaining
 * @public
 */
```





**调用**: attachment() type() extname() set() contentDisposition() append() get() isArray() concat() header() contentType() map() String() toLowerCase() TypeError() setHeader() function() getHeader() clearCookie() Date() cookie() now() Error() cookieParser() stringify() sign() isNaN() floor() serialize() location() encodeUrl() redirect() deprecate() format() escapeHtml() status() byteLength() end() vary() render() next() send() sendfile() onaborted() callback() ondirectory() onerror() onend() onfile() onfinish() setImmediate() onstream() on() onFinished() headers() keys() pipe() replace() charCodeAt() 


---




#### Function: append


Set header `field` to `val`, or pass



```
/**
 * Set header `field` to `val`, or pass
 * an object of header fields.
 *
 * Examples:
 *
 *    res.set('Foo', ['bar', 'baz']);
 *    res.set('Accept', 'application/json');
 *    res.set({ Accept: 'text/plain', 'X-API-Key': 'tobi' });
 *
 * Aliased as `res.header()`.
 *
 * When the set header is "Content-Type", the type is expanded to include
 * the charset if not present using `mime.contentType()`.
 *
 * @param {String|Object} field
 * @param {String|Array} val
 * @return {ServerResponse} for chaining
 * @public
 */
```





**调用**: append() get() isArray() concat() set() header() contentType() map() String() toLowerCase() TypeError() setHeader() function() getHeader() clearCookie() Date() cookie() now() Error() cookieParser() stringify() sign() isNaN() floor() serialize() location() encodeUrl() redirect() deprecate() format() escapeHtml() status() byteLength() end() vary() render() next() send() sendfile() onaborted() callback() ondirectory() onerror() onend() onfile() onfinish() setImmediate() onstream() on() onFinished() headers() keys() pipe() replace() charCodeAt() 


---




#### Function: header


Get value for header `field`.



```
/**
 * Get value for header `field`.
 *
 * @param {String} field
 * @return {String}
 * @public
 */
```





**调用**: header() isArray() map() String() toLowerCase() TypeError() contentType() setHeader() set() function() getHeader() clearCookie() Date() cookie() now() Error() cookieParser() stringify() sign() isNaN() floor() append() serialize() location() encodeUrl() redirect() deprecate() get() format() escapeHtml() status() byteLength() end() vary() render() next() send() sendfile() onaborted() callback() ondirectory() onerror() onend() onfile() onfinish() setImmediate() onstream() on() onFinished() headers() keys() pipe() replace() charCodeAt() 


---




#### Function: clearCookie


Set cookie `name` to `value`, with the given `options`.



```
/**
 * Set cookie `name` to `value`, with the given `options`.
 *
 * Options:
 *
 *    - `maxAge`   max-age in milliseconds, converted to `expires`
 *    - `signed`   sign the cookie
 *    - `path`     defaults to "/"
 *
 * Examples:
 *
 *    // "Remember Me" for 15 minutes
 *    res.cookie('rememberme', '1', { expires: new Date(Date.now() + 900000), httpOnly: true });
 *
 *    // same as above
 *    res.cookie('rememberme', '1', { maxAge: 900000, httpOnly: true })
 *
 * @param {String} name
 * @param {String|Object} value
 * @param {Object} [options]
 * @return {ServerResponse} for chaining
 * @public
 */
```





**调用**: clearCookie() Date() cookie() now() function() Error() cookieParser() stringify() String() sign() isNaN() floor() append() serialize() location() set() encodeUrl() redirect() deprecate() get() format() escapeHtml() status() byteLength() end() vary() render() next() send() sendfile() onaborted() callback() ondirectory() onerror() onend() onfile() onfinish() setImmediate() onstream() on() onFinished() headers() keys() setHeader() pipe() replace() charCodeAt() 


---




#### Function: location


Redirect to the given `url` with optional response `status`



```
/**
 * Redirect to the given `url` with optional response `status`
 * defaulting to 302.
 *
 * Examples:
 *
 *    res.redirect('/foo/bar');
 *    res.redirect('http://example.com');
 *    res.redirect(301, 'http://example.com');
 *    res.redirect('../login'); // /blog/post/1 -> /blog/login
 *
 * @public
 */
```





**调用**: location() set() encodeUrl() redirect() deprecate() get() format() function() escapeHtml() status() byteLength() end() vary() render() next() send() sendfile() onaborted() Error() callback() ondirectory() onerror() onend() onfile() onfinish() setImmediate() onstream() on() onFinished() headers() keys() setHeader() pipe() stringify() replace() charCodeAt() 


---




#### Function: redirect


Add `field` to Vary. If already present in the Vary set, then



```
/**
 * Add `field` to Vary. If already present in the Vary set, then
 * this call is simply ignored.
 *
 * @param {Array|String} field
 * @return {ServerResponse} for chaining
 * @public
 */
```





**调用**: redirect() deprecate() location() get() format() function() escapeHtml() status() set() byteLength() end() vary() render() next() send() sendfile() onaborted() Error() callback() ondirectory() onerror() onend() onfile() onfinish() setImmediate() onstream() on() onFinished() headers() keys() setHeader() pipe() stringify() replace() charCodeAt() 


---




#### Function: text


Set _Content-Disposition_ header to _attachment_ with optional `filename`.



```
/**
 * Set _Content-Disposition_ header to _attachment_ with optional `filename`.
 *
 * @param {String} filename
 * @return {ServerResponse}
 * @public
 */
```





**调用**: function() send() next() keys() filter() accepts() vary() set() normalizeType() default() createError() normalizeTypes() map() attachment() type() extname() contentDisposition() append() get() isArray() concat() header() contentType() String() toLowerCase() TypeError() setHeader() getHeader() clearCookie() Date() cookie() now() Error() cookieParser() stringify() sign() isNaN() floor() serialize() location() encodeUrl() redirect() deprecate() format() escapeHtml() status() byteLength() end() render() sendfile() onaborted() callback() ondirectory() onerror() onend() onfile() onfinish() setImmediate() onstream() on() onFinished() headers() pipe() replace() charCodeAt() 


---




#### Function: html


Set _Content-Disposition_ header to _attachment_ with optional `filename`.



```
/**
 * Set _Content-Disposition_ header to _attachment_ with optional `filename`.
 *
 * @param {String} filename
 * @return {ServerResponse}
 * @public
 */
```





**调用**: function() send() next() keys() filter() accepts() vary() set() normalizeType() default() createError() normalizeTypes() map() attachment() type() extname() contentDisposition() append() get() isArray() concat() header() contentType() String() toLowerCase() TypeError() setHeader() getHeader() clearCookie() Date() cookie() now() Error() cookieParser() stringify() sign() isNaN() floor() serialize() location() encodeUrl() redirect() deprecate() format() escapeHtml() status() byteLength() end() render() sendfile() onaborted() callback() ondirectory() onerror() onend() onfile() onfinish() setImmediate() onstream() on() onFinished() headers() pipe() replace() charCodeAt() 


---




#### Function: default


Add `field` to Vary. If already present in the Vary set, then



```
/**
 * Add `field` to Vary. If already present in the Vary set, then
 * this call is simply ignored.
 *
 * @param {Array|String} field
 * @return {ServerResponse} for chaining
 * @public
 */
```





**调用**: function() status() set() byteLength() end() vary() render() next() send() sendfile() onaborted() Error() callback() ondirectory() onerror() onend() onfile() onfinish() setImmediate() onstream() on() onFinished() headers() keys() setHeader() pipe() stringify() replace() charCodeAt() 


---




#### Function: render


Stringify JSON, like JSON.stringify, but v8 optimized, with the



```
/**
 * Stringify JSON, like JSON.stringify, but v8 optimized, with the
 * ability to escape characters that can trigger HTML sniffing.
 *
 * @param {*} value
 * @param {function} replacer
 * @param {number} spaces
 * @param {boolean} escape
 * @returns {string}
 * @private
 */
```





**调用**: render() function() next() send() sendfile() onaborted() Error() callback() ondirectory() onerror() onend() onfile() onfinish() setImmediate() onstream() on() onFinished() headers() keys() setHeader() pipe() stringify() replace() charCodeAt() 


---




#### Function: sendfile


Stringify JSON, like JSON.stringify, but v8 optimized, with the



```
/**
 * Stringify JSON, like JSON.stringify, but v8 optimized, with the
 * ability to escape characters that can trigger HTML sniffing.
 *
 * @param {*} value
 * @param {function} replacer
 * @param {number} spaces
 * @param {boolean} escape
 * @returns {string}
 * @private
 */
```





**调用**: sendfile() onaborted() Error() callback() ondirectory() onerror() onend() onfile() onfinish() setImmediate() function() onstream() on() onFinished() headers() keys() setHeader() pipe() stringify() replace() charCodeAt() 


---




#### Function: onaborted


Stringify JSON, like JSON.stringify, but v8 optimized, with the



```
/**
 * Stringify JSON, like JSON.stringify, but v8 optimized, with the
 * ability to escape characters that can trigger HTML sniffing.
 *
 * @param {*} value
 * @param {function} replacer
 * @param {number} spaces
 * @param {boolean} escape
 * @returns {string}
 * @private
 */
```





**调用**: onaborted() Error() callback() ondirectory() onerror() onend() onfile() onfinish() setImmediate() function() onstream() on() onFinished() headers() keys() setHeader() pipe() stringify() replace() charCodeAt() 


---




#### Function: ondirectory


Stringify JSON, like JSON.stringify, but v8 optimized, with the



```
/**
 * Stringify JSON, like JSON.stringify, but v8 optimized, with the
 * ability to escape characters that can trigger HTML sniffing.
 *
 * @param {*} value
 * @param {function} replacer
 * @param {number} spaces
 * @param {boolean} escape
 * @returns {string}
 * @private
 */
```





**调用**: ondirectory() Error() callback() onerror() onend() onfile() onfinish() onaborted() setImmediate() function() onstream() on() onFinished() headers() keys() setHeader() pipe() stringify() replace() charCodeAt() 


---




#### Function: onerror


Stringify JSON, like JSON.stringify, but v8 optimized, with the



```
/**
 * Stringify JSON, like JSON.stringify, but v8 optimized, with the
 * ability to escape characters that can trigger HTML sniffing.
 *
 * @param {*} value
 * @param {function} replacer
 * @param {number} spaces
 * @param {boolean} escape
 * @returns {string}
 * @private
 */
```





**调用**: onerror() callback() onend() onfile() onfinish() onaborted() setImmediate() function() onstream() on() onFinished() headers() keys() setHeader() pipe() stringify() replace() charCodeAt() 


---




#### Function: onend


Stringify JSON, like JSON.stringify, but v8 optimized, with the



```
/**
 * Stringify JSON, like JSON.stringify, but v8 optimized, with the
 * ability to escape characters that can trigger HTML sniffing.
 *
 * @param {*} value
 * @param {function} replacer
 * @param {number} spaces
 * @param {boolean} escape
 * @returns {string}
 * @private
 */
```





**调用**: onend() callback() onfile() onfinish() onaborted() onerror() setImmediate() function() onstream() on() onFinished() headers() keys() setHeader() pipe() stringify() replace() charCodeAt() 


---




#### Function: onfile


Stringify JSON, like JSON.stringify, but v8 optimized, with the



```
/**
 * Stringify JSON, like JSON.stringify, but v8 optimized, with the
 * ability to escape characters that can trigger HTML sniffing.
 *
 * @param {*} value
 * @param {function} replacer
 * @param {number} spaces
 * @param {boolean} escape
 * @returns {string}
 * @private
 */
```





**调用**: onfile() onfinish() onaborted() onerror() setImmediate() function() callback() onstream() on() onFinished() headers() keys() setHeader() pipe() stringify() replace() charCodeAt() 


---




#### Function: onfinish


Stringify JSON, like JSON.stringify, but v8 optimized, with the



```
/**
 * Stringify JSON, like JSON.stringify, but v8 optimized, with the
 * ability to escape characters that can trigger HTML sniffing.
 *
 * @param {*} value
 * @param {function} replacer
 * @param {number} spaces
 * @param {boolean} escape
 * @returns {string}
 * @private
 */
```





**调用**: onfinish() onaborted() onerror() setImmediate() function() callback() onstream() on() onFinished() headers() keys() setHeader() pipe() stringify() replace() charCodeAt() 


---




#### Function: onstream


Stringify JSON, like JSON.stringify, but v8 optimized, with the



```
/**
 * Stringify JSON, like JSON.stringify, but v8 optimized, with the
 * ability to escape characters that can trigger HTML sniffing.
 *
 * @param {*} value
 * @param {function} replacer
 * @param {number} spaces
 * @param {boolean} escape
 * @returns {string}
 * @private
 */
```





**调用**: onstream() on() onFinished() headers() keys() setHeader() pipe() stringify() replace() function() charCodeAt() 


---




#### Function: headers


Stringify JSON, like JSON.stringify, but v8 optimized, with the



```
/**
 * Stringify JSON, like JSON.stringify, but v8 optimized, with the
 * ability to escape characters that can trigger HTML sniffing.
 *
 * @param {*} value
 * @param {function} replacer
 * @param {number} spaces
 * @param {boolean} escape
 * @returns {string}
 * @private
 */
```





**调用**: headers() keys() setHeader() pipe() stringify() replace() function() charCodeAt() 


---




#### Function: stringify







**调用**: stringify() replace() function() charCodeAt() 


---




























































































































































#### Function: acceptParams


Compile "etag" value to function.



```
/**
 * Compile "etag" value to function.
 *
 * @param  {Boolean|String|Function} val
 * @return {Function}
 * @api private
 */
```





**调用**: acceptParams() indexOf() slice() trim() lastIndexOf() parseFloat() function() TypeError() compileQueryParser() split() map() compile() setCharset() parse() format() createETagGenerator() generateETag() isBuffer() from() etag() parseExtendedQueryString() 


---




#### Function: compileQueryParser


Compile "proxy trust" value to function.



```
/**
 * Compile "proxy trust" value to function.
 *
 * @param  {Boolean|String|Number|Array|Function} val
 * @return {Function}
 * @api private
 */
```





**调用**: compileQueryParser() TypeError() function() split() map() trim() compile() setCharset() parse() format() createETagGenerator() generateETag() isBuffer() from() etag() parseExtendedQueryString() 


---




#### Function: setCharset


Create an ETag generator function, generating ETags with



```
/**
 * Create an ETag generator function, generating ETags with
 * the given options.
 *
 * @param {object} options
 * @return {function}
 * @private
 */
```





**调用**: setCharset() parse() format() createETagGenerator() generateETag() isBuffer() from() etag() parseExtendedQueryString() 


---




#### Function: createETagGenerator


Parse an extended query string with qs.



```
/**
 * Parse an extended query string with qs.
 *
 * @param {String} str
 * @return {Object}
 * @private
 */
```





**调用**: createETagGenerator() generateETag() isBuffer() from() etag() parseExtendedQueryString() parse() 


---




#### Function: generateETag


Parse an extended query string with qs.



```
/**
 * Parse an extended query string with qs.
 *
 * @param {String} str
 * @return {Object}
 * @private
 */
```





**调用**: generateETag() isBuffer() from() etag() parseExtendedQueryString() parse() 


---




#### Function: parseExtendedQueryString







**调用**: parseExtendedQueryString() parse() 


---




































#### Function: View


Lookup view by the given `name`



```
/**
 * Lookup view by the given `name`
 *
 * @param {string} name
 * @private
 */
```





**调用**: View() extname() Error() slice() debug() lookup() concat() resolve() dirname() basename() render() engine() onRender() apply() Array() nextTick() renderTick() join() tryStat() isFile() statSync() catch() 


---




#### Function: lookup


Render with the given options.



```
/**
 * Render with the given options.
 *
 * @param {object} options
 * @param {function} callback
 * @private
 */
```





**调用**: lookup() concat() debug() resolve() dirname() basename() render() engine() onRender() apply() Array() nextTick() renderTick() join() tryStat() isFile() statSync() catch() 


---




#### Function: render


Resolve the file within the given directory.



```
/**
 * Resolve the file within the given directory.
 *
 * @param {string} dir
 * @param {string} file
 * @private
 */
```





**调用**: render() debug() engine() onRender() apply() Array() nextTick() renderTick() resolve() join() tryStat() isFile() basename() statSync() catch() 


---




#### Function: onRender


Resolve the file within the given directory.



```
/**
 * Resolve the file within the given directory.
 *
 * @param {string} dir
 * @param {string} file
 * @private
 */
```





**调用**: onRender() apply() Array() nextTick() renderTick() resolve() join() tryStat() isFile() basename() debug() statSync() catch() 


---




#### Function: renderTick


Resolve the file within the given directory.



```
/**
 * Resolve the file within the given directory.
 *
 * @param {string} dir
 * @param {string} file
 * @private
 */
```





**调用**: renderTick() apply() resolve() join() tryStat() isFile() basename() debug() statSync() catch() 


---




#### Function: resolve


Return a stat, maybe.



```
/**
 * Return a stat, maybe.
 *
 * @param {string} path
 * @return {fs.Stats}
 * @private
 */
```





**调用**: resolve() join() tryStat() isFile() basename() debug() statSync() catch() 


---




#### Function: tryStat







**调用**: tryStat() debug() statSync() catch() 


---




















































## 调用关系图


### authenticate (function)

- **文件**: E:\cody\goproject\qcoder\mcp-server\code2md\test_projects\js-test\examples\auth\index.js:60
- **调用深度**: 0




