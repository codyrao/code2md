
# py-test - Code Documentation

## 项目概览

**生成时间**: 2026-01-25T23:24:27+08:00
**项目路径**: E:\cody\goproject\qcoder\mcp-server\code2md\test_projects\py-test
**主要语言**: 

### 语言统计


- **python**: 35 个文件


### 项目摘要

| 指标 | 数量 |
|------|------|
| 总文件数 | 35 |
| 总代码行数 | 7298 |
| 总注释行数 | 0 |
| 函数数量 | 20 |
| 类数量 | 41 |
| 接口数量 | 0 |
| 变量数量 | 107 |
| 常量数量 | 6 |

## 项目结构

├── docs/
│   ├── _static/
│   ├── deploying/
│   ├── patterns/
│   ├── tutorial/
│   └── conf.py
├── examples/
│   ├── celery/
│   │   ├── src/
│   │   │   └── task_app/
│   │   │       ├── templates/
│   │   │       ├── __init__.py
│   │   │       ├── tasks.py
│   │   │       └── views.py
│   │   └── make_celery.py
│   ├── javascript/
│   │   └── js_example/
│   │       ├── templates/
│   │       ├── __init__.py
│   │       └── views.py
│   └── tutorial/
│       └── flaskr/
│           ├── templates/
│           │   ├── auth/
│           │   └── blog/
│           ├── __init__.py
│           ├── auth.py
│           ├── blog.py
│           └── db.py
└── src/
    └── flask/
        ├── json/
        │   ├── __init__.py
        │   ├── provider.py
        │   └── tag.py
        ├── sansio/
        │   ├── app.py
        │   ├── blueprints.py
        │   └── scaffold.py
        ├── __init__.py
        ├── __main__.py
        ├── app.py
        ├── blueprints.py
        ├── cli.py
        ├── config.py
        ├── ctx.py
        ├── debughelpers.py
        ├── globals.py
        ├── helpers.py
        ├── logging.py
        ├── sessions.py
        ├── signals.py
        ├── templating.py
        ├── testing.py
        ├── typing.py
        ├── views.py
        └── wrappers.py


## 文件清单


### docs\conf.py

- **语言**: python
- **包**: 
- **代码行数**: 80
- **元素数量**: 24


#### 导入


- `packaging.version`

- `pallets_sphinx_themes.get_version`

- `pallets_sphinx_themes.ProjectLink`




#### 代码元素


**Function**: `github_link`

- **完整名称**: github_link
- **文件位置**: 行 72-72
- **可见性**: public




- **参数**:
  - `name`: Any
  - `rawtext`: Any
  - `text`: Any
  - `lineno`: Any
  - `inliner`: Any
  - `options`: Any (可选)
  - `content`: Any (可选)






**Function**: `setup`

- **完整名称**: setup
- **文件位置**: 行 100-102
- **可见性**: public




- **参数**:
  - `app`: Any






**Variable**: `project`

- **完整名称**: project
- **文件位置**: 行 7-7
- **可见性**: public

- **类型**: `str`




- **值**: `"Flask"`


**Variable**: `copyright`

- **完整名称**: copyright
- **文件位置**: 行 8-8
- **可见性**: public

- **类型**: `str`




- **值**: `"2010 Pallets"`


**Variable**: `author`

- **完整名称**: author
- **文件位置**: 行 9-9
- **可见性**: public

- **类型**: `str`




- **值**: `"Pallets"`


**Variable**: `default_role`

- **完整名称**: default_role
- **文件位置**: 行 14-14
- **可见性**: public

- **类型**: `str`




- **值**: `"code"`


**Variable**: `extensions`

- **完整名称**: extensions
- **文件位置**: 行 15-15
- **可见性**: public

- **类型**: `list`




- **值**: `[`


**Variable**: `autodoc_member_order`

- **完整名称**: autodoc_member_order
- **文件位置**: 行 23-23
- **可见性**: public

- **类型**: `str`




- **值**: `"bysource"`


**Variable**: `autodoc_typehints`

- **完整名称**: autodoc_typehints
- **文件位置**: 行 24-24
- **可见性**: public

- **类型**: `str`




- **值**: `"description"`


**Variable**: `autodoc_preserve_defaults`

- **完整名称**: autodoc_preserve_defaults
- **文件位置**: 行 25-25
- **可见性**: public

- **类型**: `bool`




- **值**: `True`


**Variable**: `extlinks`

- **完整名称**: extlinks
- **文件位置**: 行 26-26
- **可见性**: public

- **类型**: `set`




- **值**: `{`


**Variable**: `intersphinx_mapping`

- **完整名称**: intersphinx_mapping
- **文件位置**: 行 31-31
- **可见性**: public

- **类型**: `set`




- **值**: `{`


**Variable**: `html_theme`

- **完整名称**: html_theme
- **文件位置**: 行 44-44
- **可见性**: public

- **类型**: `str`




- **值**: `"flask"`


**Variable**: `html_theme_options`

- **完整名称**: html_theme_options
- **文件位置**: 行 45-45
- **可见性**: public

- **类型**: `dict`




- **值**: `{"index_sidebar_logo": False}`


**Variable**: `html_context`

- **完整名称**: html_context
- **文件位置**: 行 46-46
- **可见性**: public

- **类型**: `set`




- **值**: `{`


**Variable**: `html_sidebars`

- **完整名称**: html_sidebars
- **文件位置**: 行 55-55
- **可见性**: public

- **类型**: `set`




- **值**: `{`


**Variable**: `singlehtml_sidebars`

- **完整名称**: singlehtml_sidebars
- **文件位置**: 行 59-59
- **可见性**: public

- **类型**: `dict`




- **值**: `{"index": ["project.html", "localtoc.html", "ethicalads.html"]}`


**Variable**: `html_static_path`

- **完整名称**: html_static_path
- **文件位置**: 行 60-60
- **可见性**: public

- **类型**: `list`




- **值**: `["_static"]`


**Variable**: `html_favicon`

- **完整名称**: html_favicon
- **文件位置**: 行 61-61
- **可见性**: public

- **类型**: `str`




- **值**: `"_static/flask-icon.svg"`


**Variable**: `html_logo`

- **完整名称**: html_logo
- **文件位置**: 行 62-62
- **可见性**: public

- **类型**: `str`




- **值**: `"_static/flask-logo.svg"`


**Variable**: `html_title`

- **完整名称**: html_title
- **文件位置**: 行 63-63
- **可见性**: public

- **类型**: `Any`




- **值**: `f"Flask Documentation ({version})"`


**Variable**: `html_show_sourcelink`

- **完整名称**: html_show_sourcelink
- **文件位置**: 行 64-64
- **可见性**: public

- **类型**: `bool`




- **值**: `False`


**Variable**: `gettext_uuid`

- **完整名称**: gettext_uuid
- **文件位置**: 行 66-66
- **可见性**: public

- **类型**: `bool`




- **值**: `True`


**Variable**: `gettext_compact`

- **完整名称**: gettext_compact
- **文件位置**: 行 67-67
- **可见性**: public

- **类型**: `bool`




- **值**: `False`





### examples\celery\make_celery.py

- **语言**: python
- **包**: 
- **代码行数**: 3
- **元素数量**: 2


#### 导入


- `task_app.create_app`




#### 代码元素


**Variable**: `flask_app`

- **完整名称**: flask_app
- **文件位置**: 行 3-3
- **可见性**: public

- **类型**: `Any`




- **值**: `create_app()`


**Variable**: `celery_app`

- **完整名称**: celery_app
- **文件位置**: 行 4-4
- **可见性**: public

- **类型**: `Any`




- **值**: `flask_app.extensions["celery"]`





### examples\celery\src\task_app\__init__.py

- **语言**: python
- **包**: 
- **代码行数**: 31
- **元素数量**: 1


#### 导入


- `celery.Celery`

- `celery.Task`

- `flask.Flask`

- `flask.render_template`




#### 代码元素


**Class**: `FlaskTask`

- **完整名称**: FlaskTask
- **文件位置**: 行 29-33
- **可见性**: public




- **继承**: Task 



- **子元素**: celery_app __call__ 




### examples\celery\src\task_app\tasks.py

- **语言**: python
- **包**: 
- **代码行数**: 15
- **元素数量**: 0


#### 导入


- `time`

- `celery.shared_task`

- `celery.Task`





### examples\celery\src\task_app\views.py

- **语言**: python
- **包**: 
- **代码行数**: 28
- **元素数量**: 1


#### 导入


- `celery.result.AsyncResult`

- `flask.Blueprint`

- `flask.request`

- `..tasks`




#### 代码元素


**Variable**: `bp`

- **完整名称**: bp
- **文件位置**: 行 7-7
- **可见性**: public

- **类型**: `Any`




- **值**: `Blueprint("tasks", __name__, url_prefix="/tasks")`





### examples\javascript\js_example\__init__.py

- **语言**: python
- **包**: 
- **代码行数**: 3
- **元素数量**: 1


#### 导入


- `flask.Flask`

- `js_example.views  # noqa: E402`

- `js_example.F401`




#### 代码元素


**Variable**: `app`

- **完整名称**: app
- **文件位置**: 行 3-3
- **可见性**: public

- **类型**: `Any`




- **值**: `Flask(__name__)`





### examples\javascript\js_example\views.py

- **语言**: python
- **包**: 
- **代码行数**: 13
- **元素数量**: 2


#### 导入


- `flask.jsonify`

- `flask.render_template`

- `flask.request`

- `..app`




#### 代码元素


**Function**: `index`

- **完整名称**: index
- **文件位置**: 行 10-10
- **可见性**: public




- **参数**:
  - `js`: Any






**Function**: `add`

- **完整名称**: add
- **文件位置**: 行 15-15
- **可见性**: public












### examples\tutorial\flaskr\__init__.py

- **语言**: python
- **包**: 
- **代码行数**: 28
- **元素数量**: 1


#### 导入


- `os`

- `flask.Flask`




#### 代码元素


**Function**: `create_app`

- **完整名称**: create_app
- **文件位置**: 行 6-6
- **可见性**: public




- **参数**:
  - `test_config`: Any (可选)









### examples\tutorial\flaskr\auth.py

- **语言**: python
- **包**: 
- **代码行数**: 86
- **元素数量**: 6


#### 导入


- `functools`

- `flask.Blueprint`

- `flask.flash`

- `flask.g`

- `flask.redirect`

- `flask.render_template`

- `flask.request`

- `flask.session`

- `flask.url_for`

- `werkzeug.security.check_password_hash`

- `werkzeug.security.generate_password_hash`

- `.db.get_db`




#### 代码元素


**Function**: `login_required`

- **完整名称**: login_required
- **文件位置**: 行 19-19
- **可见性**: public




- **参数**:
  - `view`: Any






**Function**: `load_logged_in_user`

- **完整名称**: load_logged_in_user
- **文件位置**: 行 33-33
- **可见性**: public









**Function**: `register`

- **完整名称**: register
- **文件位置**: 行 47-47
- **可见性**: public









**Function**: `login`

- **完整名称**: login
- **文件位置**: 行 85-85
- **可见性**: public









**Function**: `logout`

- **完整名称**: logout
- **文件位置**: 行 113-113
- **可见性**: public









**Variable**: `bp`

- **完整名称**: bp
- **文件位置**: 行 16-16
- **可见性**: public

- **类型**: `Any`




- **值**: `Blueprint("auth", __name__, url_prefix="/auth")`





### examples\tutorial\flaskr\blog.py

- **语言**: python
- **包**: 
- **代码行数**: 100
- **元素数量**: 6


#### 导入


- `flask.Blueprint`

- `flask.flash`

- `flask.g`

- `flask.redirect`

- `flask.render_template`

- `flask.request`

- `flask.url_for`

- `werkzeug.exceptions.abort`

- `.auth.login_required`

- `.db.get_db`




#### 代码元素


**Function**: `index`

- **完整名称**: index
- **文件位置**: 行 17-17
- **可见性**: public









**Function**: `get_post`

- **完整名称**: get_post
- **文件位置**: 行 28-28
- **可见性**: public




- **参数**:
  - `id`: Any
  - `check_author`: Any (可选)






**Function**: `create`

- **完整名称**: create
- **文件位置**: 行 62-62
- **可见性**: public









**Function**: `update`

- **完整名称**: update
- **文件位置**: 行 88-88
- **可见性**: public




- **参数**:
  - `id`: Any






**Function**: `delete`

- **完整名称**: delete
- **文件位置**: 行 115-115
- **可见性**: public




- **参数**:
  - `id`: Any






**Variable**: `bp`

- **完整名称**: bp
- **文件位置**: 行 13-13
- **可见性**: public

- **类型**: `Any`




- **值**: `Blueprint("blog", __name__)`





### examples\tutorial\flaskr\db.py

- **语言**: python
- **包**: 
- **代码行数**: 40
- **元素数量**: 5


#### 导入


- `sqlite3`

- `datetime.datetime`

- `click`

- `flask.current_app`

- `flask.g`




#### 代码元素


**Function**: `get_db`

- **完整名称**: get_db
- **文件位置**: 行 9-9
- **可见性**: public









**Function**: `close_db`

- **完整名称**: close_db
- **文件位置**: 行 23-23
- **可见性**: public




- **参数**:
  - `e`: Any (可选)






**Function**: `init_db`

- **完整名称**: init_db
- **文件位置**: 行 33-33
- **可见性**: public









**Function**: `init_db_command`

- **完整名称**: init_db_command
- **文件位置**: 行 42-42
- **可见性**: public









**Function**: `init_app`

- **完整名称**: init_app
- **文件位置**: 行 51-51
- **可见性**: public




- **参数**:
  - `app`: Any









### src\flask\__init__.py

- **语言**: python
- **包**: 
- **代码行数**: 39
- **元素数量**: 0


#### 导入


- `..json` (as json)

- `.app.Flask` (as Flask)

- `.blueprints.Blueprint` (as Blueprint)

- `.config.Config` (as Config)

- `.ctx.after_this_request` (as after_this_request)

- `.ctx.copy_current_request_context` (as copy_current_request_context)

- `.ctx.has_app_context` (as has_app_context)

- `.ctx.has_request_context` (as has_request_context)

- `.globals.current_app` (as current_app)

- `.globals.g` (as g)

- `.globals.request` (as request)

- `.globals.session` (as session)

- `.helpers.abort` (as abort)

- `.helpers.flash` (as flash)

- `.helpers.get_flashed_messages` (as get_flashed_messages)

- `.helpers.get_template_attribute` (as get_template_attribute)

- `.helpers.make_response` (as make_response)

- `.helpers.redirect` (as redirect)

- `.helpers.send_file` (as send_file)

- `.helpers.send_from_directory` (as send_from_directory)

- `.helpers.stream_with_context` (as stream_with_context)

- `.helpers.url_for` (as url_for)

- `.json.jsonify` (as jsonify)

- `.signals.appcontext_popped` (as appcontext_popped)

- `.signals.appcontext_pushed` (as appcontext_pushed)

- `.signals.appcontext_tearing_down` (as appcontext_tearing_down)

- `.signals.before_render_template` (as before_render_template)

- `.signals.got_request_exception` (as got_request_exception)

- `.signals.message_flashed` (as message_flashed)

- `.signals.request_finished` (as request_finished)

- `.signals.request_started` (as request_started)

- `.signals.request_tearing_down` (as request_tearing_down)

- `.signals.template_rendered` (as template_rendered)

- `.templating.render_template` (as render_template)

- `.templating.render_template_string` (as render_template_string)

- `.templating.stream_template` (as stream_template)

- `.templating.stream_template_string` (as stream_template_string)

- `.wrappers.Request` (as Request)

- `.wrappers.Response` (as Response)





### src\flask\__main__.py

- **语言**: python
- **包**: 
- **代码行数**: 2
- **元素数量**: 0


#### 导入


- `.cli.main`





### src\flask\app.py

- **语言**: python
- **包**: 
- **代码行数**: 1219
- **元素数量**: 9


#### 导入


- `__future__.annotations`

- `.collections.abc` (as cabc)

- `inspect`

- `os`

- `sys`

- `.typing` (as t)

- `weakref`

- `datetime.timedelta`

- `functools.update_wrapper`

- `inspect.iscoroutinefunction`

- `itertools.chain`

- `types.TracebackType`

- `urllib.parse.quote` (as _url_quote)

- `click`

- `werkzeug.datastructures.Headers`

- `werkzeug.datastructures.ImmutableDict`

- `werkzeug.exceptions.BadRequestKeyError`

- `werkzeug.exceptions.HTTPException`

- `werkzeug.exceptions.InternalServerError`

- `werkzeug.routing.BuildError`

- `werkzeug.routing.MapAdapter`

- `werkzeug.routing.RequestRedirect`

- `werkzeug.routing.RoutingException`

- `werkzeug.routing.Rule`

- `werkzeug.serving.is_running_from_reloader`

- `werkzeug.wrappers.Response` (as BaseResponse)

- `werkzeug.wsgi.get_host`

- `..cli`

- `..typing` (as ft)

- `.ctx.AppContext`

- `.globals._cv_app`

- `.globals.app_ctx`

- `.globals.g`

- `.globals.request`

- `.globals.session`

- `.helpers.get_debug_flag`

- `.helpers.get_flashed_messages`

- `.helpers.get_load_dotenv`

- `.helpers.send_from_directory`

- `.sansio.app.App`

- `.sessions.SecureCookieSessionInterface`

- `.sessions.SessionInterface`

- `.signals.appcontext_tearing_down`

- `.signals.got_request_exception`

- `.signals.request_finished`

- `.signals.request_started`

- `.signals.request_tearing_down`

- `.templating.Environment`

- `.wrappers.Request`

- `.wrappers.Response`




#### 代码元素


**Class**: `Flask`

- **完整名称**: Flask
- **文件位置**: 行 105-105
- **可见性**: public




- **继承**: App 



- **子元素**: app default_config base_method iter_params param stacklevel import_name static_url_path static_folder static_host host_matching subdomain_matching template_folder instance_path instance_relative_config root_path self_ref endpoint host view_func value max_age path options auto_reload rv url_for get_flashed_messages config request session g subdomain server_name script_name url_scheme names orig_ctx fg sn_host port client handler exc_info propagate server_error req response methods result url_adapter blueprint_name _external method force_external _anchor len_rv status headers builder environ ctx error __init_subclass__ get_send_file_max_age send_static_file open_resource open_instance_resource create_jinja_environment create_url_adapter raise_routing_exception update_template_context make_shell_context run test_client test_cli_runner handle_http_exception handle_user_exception handle_exception log_exception dispatch_request full_dispatch_request finalize_request make_default_options_response ensure_sync async_to_sync url_for make_response preprocess_request process_response do_teardown_request do_teardown_appcontext app_context request_context test_request_context wsgi_app __call__ 

**Class**: `CustomClient`

- **完整名称**: CustomClient
- **文件位置**: 行 781-786
- **可见性**: public




- **继承**: FlaskClient 



- **子元素**: client handler exc_info propagate server_error req stacklevel rv response methods result url_adapter blueprint_name endpoint _external method url_scheme force_external _anchor len_rv status headers names builder environ app ctx error test_cli_runner handle_http_exception handle_user_exception handle_exception log_exception dispatch_request full_dispatch_request finalize_request make_default_options_response ensure_sync async_to_sync url_for make_response preprocess_request process_response do_teardown_request do_teardown_appcontext app_context request_context test_request_context wsgi_app __call__ 

**Variable**: `T_shell_context_processor`

- **完整名称**: T_shell_context_processor
- **文件位置**: 行 63-63
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar(`


**Variable**: `T_teardown`

- **完整名称**: T_teardown
- **文件位置**: 行 66-66
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T_teardown", bound=ft.TeardownCallable)`


**Variable**: `T_template_filter`

- **完整名称**: T_template_filter
- **文件位置**: 行 67-67
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T_template_filter", bound=ft.TemplateFilterCallable)`


**Variable**: `T_template_global`

- **完整名称**: T_template_global
- **文件位置**: 行 68-68
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T_template_global", bound=ft.TemplateGlobalCallable)`


**Variable**: `T_template_test`

- **完整名称**: T_template_test
- **文件位置**: 行 69-69
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T_template_test", bound=ft.TemplateTestCallable)`


**Variable**: `F`

- **完整名称**: F
- **文件位置**: 行 79-79
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("F", bound=t.Callable[..., t.Any])`


**Constant**: `F`

- **完整名称**: F
- **文件位置**: 行 79-79
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("F", bound=t.Callable[..., t.Any])`





### src\flask\blueprints.py

- **语言**: python
- **包**: 
- **代码行数**: 94
- **元素数量**: 1


#### 导入


- `__future__.annotations`

- `os`

- `.typing` (as t)

- `datetime.timedelta`

- `.cli.AppGroup`

- `.globals.current_app`

- `.helpers.send_from_directory`

- `.sansio.blueprints.Blueprint` (as SansioBlueprint)

- `.sansio.blueprints.BlueprintSetupState` (as BlueprintSetupState  # noqa)

- `.sansio.scaffold._sentinel`




#### 代码元素


**Class**: `Blueprint`

- **完整名称**: Blueprint
- **文件位置**: 行 15-15
- **可见性**: public




- **继承**: SansioBlueprint 



- **子元素**: value max_age path get_send_file_max_age send_static_file open_resource 




### src\flask\cli.py

- **语言**: python
- **包**: 
- **代码行数**: 823
- **元素数量**: 10


#### 导入


- `__future__.annotations`

- `ast`

- `.collections.abc` (as cabc)

- `importlib.metadata`

- `inspect`

- `os`

- `platform`

- `re`

- `sys`

- `traceback`

- `.typing` (as t)

- `functools.update_wrapper`

- `operator.itemgetter`

- `types.ModuleType`

- `click`

- `click.core.ParameterSource`

- `werkzeug.run_simple`

- `werkzeug.serving.is_running_from_reloader`

- `werkzeug.utils.import_string`

- `.globals.current_app`

- `.helpers.get_debug_flag`

- `.helpers.get_load_dotenv`




#### 代码元素


**Class**: `ScriptInfo`

- **完整名称**: ScriptInfo
- **文件位置**: 行 290-290
- **可见性**: public








- **子元素**: app import_name pass_script_info F wrap_for_ctx f info metavar help is_eager expose_value callback source _debug_option ctx param type rv create_app set_debug_flag load_dotenv_defaults fg err name value obj cert is_adhoc is_context items super_convert default debug reload debugger use_reloader use_debugger threaded ssl_context extra_files exclude_patterns banner startup interactive_hook rules ignored_methods host_matching has_domain rows row headers sorts widths template cli load_app with_appcontext decorator command group _set_app _set_debug _env_file_callback _load_plugin_commands get_command list_commands make_context parse_args _path_is_ancestor load_dotenv show_server_banner convert _validate_key app main 

**Class**: `FlaskGroup`

- **完整名称**: FlaskGroup
- **文件位置**: 行 528-528
- **可见性**: public




- **继承**: AppGroup 



- **子元素**: rv info app create_app set_debug_flag load_dotenv_defaults fg err name value obj cert is_adhoc is_context items super_convert type help is_eager callback expose_value default debug reload debugger use_reloader use_debugger threaded ssl_context extra_files exclude_patterns banner startup interactive_hook rules ignored_methods host_matching has_domain rows row headers sorts widths template cli _load_plugin_commands get_command list_commands make_context parse_args _path_is_ancestor load_dotenv show_server_banner convert _validate_key app main 

**Variable**: `version_option`

- **完整名称**: version_option
- **文件位置**: 行 283-283
- **可见性**: public

- **类型**: `Any`




- **值**: `click.Option(`


**Variable**: `pass_script_info`

- **完整名称**: pass_script_info
- **文件位置**: 行 375-375
- **可见性**: public

- **类型**: `Any`




- **值**: `click.make_pass_decorator(ScriptInfo, ensure=True)`


**Variable**: `F`

- **完整名称**: F
- **文件位置**: 行 377-377
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("F", bound=t.Callable[..., t.Any])`


**Variable**: `_app_option`

- **完整名称**: _app_option
- **文件位置**: 行 453-453
- **可见性**: public

- **类型**: `Any`




- **值**: `click.Option(`


**Variable**: `_debug_option`

- **完整名称**: _debug_option
- **文件位置**: 行 485-485
- **可见性**: public

- **类型**: `Any`




- **值**: `click.Option(`


**Variable**: `_env_file_option`

- **完整名称**: _env_file_option
- **文件位置**: 行 517-517
- **可见性**: public

- **类型**: `Any`




- **值**: `click.Option(`


**Variable**: `cli`

- **完整名称**: cli
- **文件位置**: 行 1110-1110
- **可见性**: public

- **类型**: `Any`




- **值**: `FlaskGroup(`


**Constant**: `F`

- **完整名称**: F
- **文件位置**: 行 377-377
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("F", bound=t.Callable[..., t.Any])`





### src\flask\config.py

- **语言**: python
- **包**: 
- **代码行数**: 281
- **元素数量**: 2


#### 导入


- `__future__.annotations`

- `errno`

- `json`

- `os`

- `types`

- `.typing` (as t)

- `werkzeug.utils.import_string`




#### 代码元素


**Variable**: `T`

- **完整名称**: T
- **文件位置**: 行 17-17
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T")`


**Constant**: `T`

- **完整名称**: T
- **文件位置**: 行 17-17
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T")`





### src\flask\ctx.py

- **语言**: python
- **包**: 
- **代码行数**: 381
- **元素数量**: 8


#### 导入


- `__future__.annotations`

- `contextvars`

- `.typing` (as t)

- `functools.update_wrapper`

- `types.TracebackType`

- `werkzeug.exceptions.HTTPException`

- `werkzeug.routing.MapAdapter`

- `..typing` (as ft)

- `.globals._cv_app`

- `.signals.appcontext_popped`

- `.signals.appcontext_pushed`




#### 代码元素


**Class**: `_AppCtxGlobals`

- **完整名称**: _AppCtxGlobals
- **文件位置**: 行 26-26
- **可见性**: public








- **子元素**: ctx F remote_addr request session si result stacklevel __getattr__ __setattr__ __delattr__ get pop setdefault __contains__ __iter__ after_this_request index add_header copy_current_request_context do_some_work wrapper has_request_context has_app_context from_environ has_request copy request session match_request push __enter__ __exit__ 

**Method**: `index`

- **完整名称**: _AppCtxGlobals.index
- **文件位置**: 行 128-128
- **可见性**: public









**Method**: `add_header`

- **完整名称**: _AppCtxGlobals.add_header
- **文件位置**: 行 130-130
- **可见性**: public




- **参数**:
  - `response`: Any






**Method**: `do_some_work`

- **完整名称**: _AppCtxGlobals.do_some_work
- **文件位置**: 行 182-182
- **可见性**: public









**Class**: `AppContext`

- **完整名称**: AppContext
- **文件位置**: 行 256-256
- **可见性**: public








- **子元素**: request session si result ctx stacklevel from_environ has_request copy request session match_request push pop __enter__ __exit__ __getattr__ 

**Variable**: `_sentinel`

- **完整名称**: _sentinel
- **文件位置**: 行 26-26
- **可见性**: public

- **类型**: `Any`




- **值**: `object()`


**Variable**: `F`

- **完整名称**: F
- **文件位置**: 行 150-150
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("F", bound=t.Callable[..., t.Any])`


**Constant**: `F`

- **完整名称**: F
- **文件位置**: 行 150-150
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("F", bound=t.Callable[..., t.Any])`





### src\flask\debughelpers.py

- **语言**: python
- **包**: 
- **代码行数**: 146
- **元素数量**: 3


#### 导入


- `__future__.annotations`

- `.typing` (as t)

- `jinja2.loaders.BaseLoader`

- `werkzeug.routing.RequestRedirect`

- `.blueprints.Blueprint`

- `.globals._cv_app`

- `.sansio.app.App`




#### 代码元素


**Class**: `UnexpectedUnicodeError`

- **完整名称**: UnexpectedUnicodeError
- **文件位置**: 行 14-14
- **可见性**: public




- **继承**: AssertionError UnicodeError 



- **子元素**: form_matches buf names exc oldcls info total_found blueprint src_info detail seems_fishy attach_enctype_error_multidict __getitem__ _dump_loader_info explain_template_loading_attempts 

**Class**: `DebugFilesKeyError`

- **完整名称**: DebugFilesKeyError
- **文件位置**: 行 20-20
- **可见性**: public




- **继承**: KeyError AssertionError 



- **子元素**: form_matches buf names exc oldcls info total_found blueprint src_info detail seems_fishy attach_enctype_error_multidict __getitem__ _dump_loader_info explain_template_loading_attempts 

**Class**: `FormDataRoutingRedirect`

- **完整名称**: FormDataRoutingRedirect
- **文件位置**: 行 47-47
- **可见性**: public




- **继承**: AssertionError 



- **子元素**: exc buf oldcls info total_found blueprint src_info detail seems_fishy attach_enctype_error_multidict __getitem__ _dump_loader_info explain_template_loading_attempts 




### src\flask\globals.py

- **语言**: python
- **包**: 
- **代码行数**: 56
- **元素数量**: 2


#### 导入


- `__future__.annotations`

- `.typing` (as t)

- `contextvars.ContextVar`

- `werkzeug.local.LocalProxy`




#### 代码元素


**Variable**: `_no_app_msg`

- **完整名称**: _no_app_msg
- **文件位置**: 行 33-33
- **可见性**: public

- **类型**: `str`




- **值**: `"""\`


**Variable**: `_no_req_msg`

- **完整名称**: _no_req_msg
- **文件位置**: 行 51-51
- **可见性**: public

- **类型**: `str`




- **值**: `"""\`





### src\flask\helpers.py

- **语言**: python
- **包**: 
- **代码行数**: 472
- **元素数量**: 0


#### 导入


- `__future__.annotations`

- `importlib.util`

- `os`

- `sys`

- `.typing` (as t)

- `datetime.datetime`

- `functools.cache`

- `functools.update_wrapper`

- `werkzeug.utils`

- `werkzeug.exceptions.abort` (as _wz_abort)

- `werkzeug.utils.redirect` (as _wz_redirect)

- `werkzeug.wrappers.Response` (as BaseResponse)

- `.globals._cv_app`

- `.globals.app_ctx`

- `.globals.current_app`

- `.globals.request`

- `.globals.session`

- `.signals.message_flashed`





### src\flask\json\__init__.py

- **语言**: python
- **包**: 
- **代码行数**: 122
- **元素数量**: 0


#### 导入


- `__future__.annotations`

- `.json` (as _json)

- `.typing` (as t)

- `..globals.current_app`

- `.provider._default`





### src\flask\json\provider.py

- **语言**: python
- **包**: 
- **代码行数**: 163
- **元素数量**: 2


#### 导入


- `__future__.annotations`

- `dataclasses`

- `decimal`

- `json`

- `.typing` (as t)

- `uuid`

- `weakref`

- `datetime.date`

- `werkzeug.http.http_date`




#### 代码元素


**Class**: `JSONProvider`

- **完整名称**: JSONProvider
- **文件位置**: 行 16-16
- **可见性**: public








- **子元素**: obj ensure_ascii sort_keys mimetype dumps dump loads load _prepare_response_obj response _default 

**Class**: `DefaultJSONProvider`

- **完整名称**: DefaultJSONProvider
- **文件位置**: 行 121-121
- **可见性**: public




- **继承**: JSONProvider 



- **子元素**: ensure_ascii sort_keys mimetype obj dumps loads response 




### src\flask\json\tag.py

- **语言**: python
- **包**: 
- **代码行数**: 221
- **元素数量**: 14


#### 导入


- `__future__.annotations`

- `.typing` (as t)

- `base64.b64decode`

- `base64.b64encode`

- `datetime.datetime`

- `uuid.UUID`

- `markupsafe.Markup`

- `werkzeug.http.http_date`

- `werkzeug.http.parse_date`

- `..json.dumps`

- `..json.loads`




#### 代码元素


**Class**: `TagOrderedDict`

- **完整名称**: TagOrderedDict
- **文件位置**: 行 26-30
- **可见性**: public




- **继承**: JSONTag 



- **子元素**: __slots__ key tag default_tags value check to_json to_python tag register untag _untag_scan dumps loads 

**Method**: `check`

- **完整名称**: TagOrderedDict.check
- **文件位置**: 行 31-33
- **可见性**: public




- **参数**:
  - `self`: Any
  - `value`: Any






**Method**: `to_json`

- **完整名称**: TagOrderedDict.to_json
- **文件位置**: 行 34-36
- **可见性**: public




- **参数**:
  - `self`: Any
  - `value`: Any






**Method**: `to_python`

- **完整名称**: TagOrderedDict.to_python
- **文件位置**: 行 37-39
- **可见性**: public




- **参数**:
  - `self`: Any
  - `value`: Any






**Class**: `JSONTag`

- **完整名称**: JSONTag
- **文件位置**: 行 57-57
- **可见性**: public








- **子元素**: __slots__ key tag default_tags value check to_json to_python tag register untag _untag_scan dumps loads 

**Class**: `TagDict`

- **完整名称**: TagDict
- **文件位置**: 行 90-90
- **可见性**: public




- **继承**: JSONTag 



- **子元素**: __slots__ key tag default_tags value check to_json to_python register tag untag _untag_scan dumps loads 

**Class**: `PassDict`

- **完整名称**: PassDict
- **文件位置**: 行 116-116
- **可见性**: public




- **继承**: JSONTag 



- **子元素**: __slots__ tag key default_tags value check to_json to_python register tag untag _untag_scan dumps loads 

**Class**: `TagTuple`

- **完整名称**: TagTuple
- **文件位置**: 行 130-130
- **可见性**: public




- **继承**: JSONTag 



- **子元素**: __slots__ key tag default_tags value check to_json to_python register tag untag _untag_scan dumps loads 

**Class**: `PassList`

- **完整名称**: PassList
- **文件位置**: 行 144-144
- **可见性**: public




- **继承**: JSONTag 



- **子元素**: __slots__ tag key default_tags value check to_json to_python register tag untag _untag_scan dumps loads 

**Class**: `TagBytes`

- **完整名称**: TagBytes
- **文件位置**: 行 156-156
- **可见性**: public




- **继承**: JSONTag 



- **子元素**: __slots__ key default_tags tag value check to_json to_python register tag untag _untag_scan dumps loads 

**Class**: `TagMarkup`

- **完整名称**: TagMarkup
- **文件位置**: 行 170-170
- **可见性**: public




- **继承**: JSONTag 



- **子元素**: __slots__ key default_tags tag value check to_json to_python register tag untag _untag_scan dumps loads 

**Class**: `TagUUID`

- **完整名称**: TagUUID
- **文件位置**: 行 188-188
- **可见性**: public




- **继承**: JSONTag 



- **子元素**: __slots__ key default_tags tag value check to_json to_python register tag untag _untag_scan dumps loads 

**Class**: `TagDateTime`

- **完整名称**: TagDateTime
- **文件位置**: 行 202-202
- **可见性**: public




- **继承**: JSONTag 



- **子元素**: __slots__ key default_tags tag value check to_json to_python register tag untag _untag_scan dumps loads 

**Class**: `TaggedJSONSerializer`

- **完整名称**: TaggedJSONSerializer
- **文件位置**: 行 216-216
- **可见性**: public








- **子元素**: __slots__ default_tags tag key value register tag untag _untag_scan dumps loads 




### src\flask\logging.py

- **语言**: python
- **包**: 
- **代码行数**: 53
- **元素数量**: 1


#### 导入


- `__future__.annotations`

- `logging`

- `sys`

- `.typing` (as t)

- `werkzeug.local.LocalProxy`

- `.globals.request`




#### 代码元素


**Variable**: `default_handler`

- **完整名称**: default_handler
- **文件位置**: 行 52-52
- **可见性**: public

- **类型**: `Any`




- **值**: `logging.StreamHandler(wsgi_errors_stream)  # type: ignore`





### src\flask\sansio\app.py

- **语言**: python
- **包**: 
- **代码行数**: 673
- **元素数量**: 9


#### 导入


- `__future__.annotations`

- `logging`

- `os`

- `sys`

- `.typing` (as t)

- `datetime.timedelta`

- `itertools.chain`

- `werkzeug.exceptions.Aborter`

- `werkzeug.exceptions.BadRequest`

- `werkzeug.exceptions.BadRequestKeyError`

- `werkzeug.routing.BuildError`

- `werkzeug.routing.Map`

- `werkzeug.routing.Rule`

- `werkzeug.sansio.response.Response`

- `werkzeug.utils.cached_property`

- `werkzeug.utils.redirect` (as _wz_redirect)

- `...typing` (as ft)

- `..config.Config`

- `..config.ConfigAttribute`

- `..ctx._AppCtxGlobals`

- `..helpers._split_blueprint_path`

- `..helpers.get_debug_flag`

- `..json.provider.DefaultJSONProvider`

- `..json.provider.JSONProvider`

- `..logging.create_logger`

- `..templating.DispatchingJinjaLoader`

- `..templating.Environment`

- `.scaffold._endpoint_from_view_func`

- `.scaffold.find_package`

- `.scaffold.Scaffold`

- `.scaffold.setupmethod`




#### 代码元素


**Class**: `App`

- **完整名称**: App
- **文件位置**: 行 56-56
- **可见性**: public




- **继承**: Scaffold 



- **子元素**: app aborter_class jinja_environment app_ctx_globals_class config_class testing secret_key permanent_session_lifetime get_converter url_rule_class url_map_class import_name static_folder static_url_path template_folder root_path instance_path defaults endpoint methods provide_automatic_options rule_obj old_func names handler_map handler trap_bad_request code Response rv error _check_setup_finished name logger jinja_env create_jinja_environment make_config make_aborter auto_find_instance_path create_global_jinja_loader select_jinja_autoescape debug register_blueprint iter_blueprints add_url_rule template_filter reverse_filter decorator add_template_filter template_test is_prime_test add_template_test template_global double add_template_global teardown_appcontext shell_context_processor _find_error_handler trap_http_exception redirect inject_url_defaults handle_url_build_error 

**Method**: `reverse_filter`

- **完整名称**: App.reverse_filter
- **文件位置**: 行 676-676
- **可见性**: public




- **参数**:
  - `s`: Any






**Method**: `is_prime_test`

- **完整名称**: App.is_prime_test
- **文件位置**: 行 726-726
- **可见性**: public




- **参数**:
  - `n`: Any






**Method**: `double`

- **完整名称**: App.double
- **文件位置**: 行 785-785
- **可见性**: public




- **参数**:
  - `n`: Any






**Variable**: `T_shell_context_processor`

- **完整名称**: T_shell_context_processor
- **文件位置**: 行 43-43
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar(`


**Variable**: `T_teardown`

- **完整名称**: T_teardown
- **文件位置**: 行 46-46
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T_teardown", bound=ft.TeardownCallable)`


**Variable**: `T_template_filter`

- **完整名称**: T_template_filter
- **文件位置**: 行 47-47
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T_template_filter", bound=ft.TemplateFilterCallable)`


**Variable**: `T_template_global`

- **完整名称**: T_template_global
- **文件位置**: 行 48-48
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T_template_global", bound=ft.TemplateGlobalCallable)`


**Variable**: `T_template_test`

- **完整名称**: T_template_test
- **文件位置**: 行 49-49
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T_template_test", bound=ft.TemplateTestCallable)`





### src\flask\sansio\blueprints.py

- **语言**: python
- **包**: 
- **代码行数**: 553
- **元素数量**: 13


#### 导入


- `__future__.annotations`

- `os`

- `.typing` (as t)

- `collections.defaultdict`

- `functools.update_wrapper`

- `...typing` (as ft)

- `.scaffold._endpoint_from_view_func`

- `.scaffold._sentinel`

- `.scaffold.Scaffold`

- `.scaffold.setupmethod`




#### 代码元素


**Class**: `BlueprintSetupState`

- **完整名称**: BlueprintSetupState
- **文件位置**: 行 31-31
- **可见性**: public








- **子元素**: subdomain url_prefix rule endpoint defaults _got_registered_once import_name static_folder static_url_path template_folder root_path url_defaults name_prefix self_name name bp_desc existing_at first_bp_registration first_name_registration state view_func cli_resolved_group bp_options bp_url_prefix bp_subdomain key value provide_automatic_options add_url_rule _check_setup_finished record record_once wrapper make_setup_state register_blueprint register _merge_blueprint_funcs extend app_template_filter decorator add_app_template_filter register_template_filter app_template_test add_app_template_test register_template_test app_template_global add_app_template_global register_template_global before_app_request after_app_request teardown_app_request app_context_processor app_errorhandler from_blueprint app_url_value_preprocessor app_url_defaults 

**Class**: `Blueprint`

- **完整名称**: Blueprint
- **文件位置**: 行 31-31
- **可见性**: public




- **继承**: Scaffold 



- **子元素**: subdomain url_prefix rule endpoint defaults _got_registered_once import_name static_folder static_url_path template_folder root_path url_defaults name_prefix self_name name bp_desc existing_at first_bp_registration first_name_registration state view_func cli_resolved_group bp_options bp_url_prefix bp_subdomain key value provide_automatic_options add_url_rule _check_setup_finished record record_once wrapper make_setup_state register_blueprint register _merge_blueprint_funcs extend app_template_filter decorator add_app_template_filter register_template_filter app_template_test add_app_template_test register_template_test app_template_global add_app_template_global register_template_global before_app_request after_app_request teardown_app_request app_context_processor app_errorhandler from_blueprint app_url_value_preprocessor app_url_defaults 

**Variable**: `DeferredSetupFunction`

- **完整名称**: DeferredSetupFunction
- **文件位置**: 行 17-17
- **可见性**: public

- **类型**: `Any`




- **值**: `t.Callable[["BlueprintSetupState"], None]`


**Variable**: `T_after_request`

- **完整名称**: T_after_request
- **文件位置**: 行 18-18
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T_after_request", bound=ft.AfterRequestCallable[t.Any])`


**Variable**: `T_before_request`

- **完整名称**: T_before_request
- **文件位置**: 行 19-19
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T_before_request", bound=ft.BeforeRequestCallable)`


**Variable**: `T_error_handler`

- **完整名称**: T_error_handler
- **文件位置**: 行 20-20
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T_error_handler", bound=ft.ErrorHandlerCallable)`


**Variable**: `T_teardown`

- **完整名称**: T_teardown
- **文件位置**: 行 21-21
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T_teardown", bound=ft.TeardownCallable)`


**Variable**: `T_template_context_processor`

- **完整名称**: T_template_context_processor
- **文件位置**: 行 22-22
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar(`


**Variable**: `T_template_filter`

- **完整名称**: T_template_filter
- **文件位置**: 行 25-25
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T_template_filter", bound=ft.TemplateFilterCallable)`


**Variable**: `T_template_global`

- **完整名称**: T_template_global
- **文件位置**: 行 26-26
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T_template_global", bound=ft.TemplateGlobalCallable)`


**Variable**: `T_template_test`

- **完整名称**: T_template_test
- **文件位置**: 行 27-27
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T_template_test", bound=ft.TemplateTestCallable)`


**Variable**: `T_url_defaults`

- **完整名称**: T_url_defaults
- **文件位置**: 行 28-28
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T_url_defaults", bound=ft.URLDefaultCallable)`


**Variable**: `T_url_value_preprocessor`

- **完整名称**: T_url_value_preprocessor
- **文件位置**: 行 29-29
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar(`





### src\flask\sansio\scaffold.py

- **语言**: python
- **包**: 
- **代码行数**: 533
- **元素数量**: 17


#### 导入


- `__future__.annotations`

- `importlib.util`

- `os`

- `pathlib`

- `sys`

- `.typing` (as t)

- `collections.defaultdict`

- `functools.update_wrapper`

- `jinja2.BaseLoader`

- `jinja2.FileSystemLoader`

- `werkzeug.exceptions.default_exceptions`

- `werkzeug.exceptions.HTTPException`

- `werkzeug.utils.cached_property`

- `...typing` (as ft)

- `..helpers.get_root_path`

- `..templating._default_template_ctx_processor`




#### 代码元素


**Class**: `Scaffold`

- **完整名称**: Scaffold
- **文件位置**: 行 49-49
- **可见性**: public








- **子元素**: root_path value basename endpoint exc_class root_spec package_spec package_path search_location py_prefix _check_setup_finished static_folder has_static_folder static_url_path jinja_loader _method_route get post put delete patch route index decorator add_url_rule endpoint example before_request load_user after_request teardown_request context_processor url_value_preprocessor url_defaults errorhandler page_not_found special_exception_handler register_error_handler _get_exc_class_and_code _endpoint_from_view_func _find_package_path find_package 

**Method**: `index`

- **完整名称**: Scaffold.index
- **文件位置**: 行 344-344
- **可见性**: public









**Method**: `example`

- **完整名称**: Scaffold.example
- **文件位置**: 行 446-446
- **可见性**: public









**Method**: `load_user`

- **完整名称**: Scaffold.load_user
- **文件位置**: 行 469-470
- **可见性**: public









**Method**: `page_not_found`

- **完整名称**: Scaffold.page_not_found
- **文件位置**: 行 607-607
- **可见性**: public




- **参数**:
  - `error`: Any






**Method**: `special_exception_handler`

- **完整名称**: Scaffold.special_exception_handler
- **文件位置**: 行 613-613
- **可见性**: public




- **参数**:
  - `error`: Any






**Variable**: `_sentinel`

- **完整名称**: _sentinel
- **文件位置**: 行 25-25
- **可见性**: public

- **类型**: `Any`




- **值**: `object()`


**Variable**: `F`

- **完整名称**: F
- **文件位置**: 行 27-27
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("F", bound=t.Callable[..., t.Any])`


**Variable**: `T_after_request`

- **完整名称**: T_after_request
- **文件位置**: 行 28-28
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T_after_request", bound=ft.AfterRequestCallable[t.Any])`


**Variable**: `T_before_request`

- **完整名称**: T_before_request
- **文件位置**: 行 29-29
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T_before_request", bound=ft.BeforeRequestCallable)`


**Variable**: `T_error_handler`

- **完整名称**: T_error_handler
- **文件位置**: 行 30-30
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T_error_handler", bound=ft.ErrorHandlerCallable)`


**Variable**: `T_teardown`

- **完整名称**: T_teardown
- **文件位置**: 行 31-31
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T_teardown", bound=ft.TeardownCallable)`


**Variable**: `T_template_context_processor`

- **完整名称**: T_template_context_processor
- **文件位置**: 行 32-32
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar(`


**Variable**: `T_url_defaults`

- **完整名称**: T_url_defaults
- **文件位置**: 行 35-35
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T_url_defaults", bound=ft.URLDefaultCallable)`


**Variable**: `T_url_value_preprocessor`

- **完整名称**: T_url_value_preprocessor
- **文件位置**: 行 36-36
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar(`


**Variable**: `T_route`

- **完整名称**: T_route
- **文件位置**: 行 39-39
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("T_route", bound=ft.RouteCallable)`


**Constant**: `F`

- **完整名称**: F
- **文件位置**: 行 27-27
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("F", bound=t.Callable[..., t.Any])`





### src\flask\sessions.py

- **语言**: python
- **包**: 
- **代码行数**: 283
- **元素数量**: 4


#### 导入


- `__future__.annotations`

- `.collections.abc` (as c)

- `hashlib`

- `.typing` (as t)

- `collections.abc.MutableMapping`

- `datetime.datetime`

- `datetime.timezone`

- `itsdangerous.BadSignature`

- `itsdangerous.URLSafeTimedSerializer`

- `werkzeug.datastructures.CallbackDict`

- `.json.tag.TaggedJSONSerializer`




#### 代码元素


**Class**: `NullSession`

- **完整名称**: NullSession
- **文件位置**: 行 94-94
- **可见性**: public




- **继承**: SecureCookieSession 



- **子元素**: __setitem__ app null_session_class pickle_based session_json_serializer salt digest_method key_derivation serializer session_class signer_kwargs s val max_age data name domain path secure partitioned samesite httponly expires _fail make_null_session is_null_session get_cookie_name get_cookie_domain get_cookie_path get_cookie_httponly get_cookie_secure get_cookie_samesite get_cookie_partitioned get_expiration_time should_set_cookie open_session save_session _lazy_sha1 get_signing_serializer 

**Class**: `SessionInterface`

- **完整名称**: SessionInterface
- **文件位置**: 行 111-111
- **可见性**: public








- **子元素**: app null_session_class pickle_based session_json_serializer salt digest_method key_derivation serializer session_class signer_kwargs s val max_age data name domain path secure partitioned samesite httponly expires make_null_session is_null_session get_cookie_name get_cookie_domain get_cookie_path get_cookie_httponly get_cookie_secure get_cookie_samesite get_cookie_partitioned get_expiration_time should_set_cookie open_session save_session _lazy_sha1 get_signing_serializer 

**Class**: `SecureCookieSessionInterface`

- **完整名称**: SecureCookieSessionInterface
- **文件位置**: 行 295-295
- **可见性**: public




- **继承**: SessionInterface 



- **子元素**: salt digest_method key_derivation serializer session_class signer_kwargs s val max_age data name domain path secure partitioned samesite httponly expires get_signing_serializer open_session save_session 

**Variable**: `session_json_serializer`

- **完整名称**: session_json_serializer
- **文件位置**: 行 287-287
- **可见性**: public

- **类型**: `Any`




- **值**: `TaggedJSONSerializer()`





### src\flask\signals.py

- **语言**: python
- **包**: 
- **代码行数**: 13
- **元素数量**: 11


#### 导入


- `__future__.annotations`

- `blinker.Namespace`




#### 代码元素


**Variable**: `_signals`

- **完整名称**: _signals
- **文件位置**: 行 6-6
- **可见性**: public

- **类型**: `Any`




- **值**: `Namespace()`


**Variable**: `template_rendered`

- **完整名称**: template_rendered
- **文件位置**: 行 8-8
- **可见性**: public

- **类型**: `Any`




- **值**: `_signals.signal("template-rendered")`


**Variable**: `before_render_template`

- **完整名称**: before_render_template
- **文件位置**: 行 9-9
- **可见性**: public

- **类型**: `Any`




- **值**: `_signals.signal("before-render-template")`


**Variable**: `request_started`

- **完整名称**: request_started
- **文件位置**: 行 10-10
- **可见性**: public

- **类型**: `Any`




- **值**: `_signals.signal("request-started")`


**Variable**: `request_finished`

- **完整名称**: request_finished
- **文件位置**: 行 11-11
- **可见性**: public

- **类型**: `Any`




- **值**: `_signals.signal("request-finished")`


**Variable**: `request_tearing_down`

- **完整名称**: request_tearing_down
- **文件位置**: 行 12-12
- **可见性**: public

- **类型**: `Any`




- **值**: `_signals.signal("request-tearing-down")`


**Variable**: `got_request_exception`

- **完整名称**: got_request_exception
- **文件位置**: 行 13-13
- **可见性**: public

- **类型**: `Any`




- **值**: `_signals.signal("got-request-exception")`


**Variable**: `appcontext_tearing_down`

- **完整名称**: appcontext_tearing_down
- **文件位置**: 行 14-14
- **可见性**: public

- **类型**: `Any`




- **值**: `_signals.signal("appcontext-tearing-down")`


**Variable**: `appcontext_pushed`

- **完整名称**: appcontext_pushed
- **文件位置**: 行 15-15
- **可见性**: public

- **类型**: `Any`




- **值**: `_signals.signal("appcontext-pushed")`


**Variable**: `appcontext_popped`

- **完整名称**: appcontext_popped
- **文件位置**: 行 16-16
- **可见性**: public

- **类型**: `Any`




- **值**: `_signals.signal("appcontext-popped")`


**Variable**: `message_flashed`

- **完整名称**: message_flashed
- **文件位置**: 行 17-17
- **可见性**: public

- **类型**: `Any`




- **值**: `_signals.signal("message-flashed")`





### src\flask\templating.py

- **语言**: python
- **包**: 
- **代码行数**: 163
- **元素数量**: 2


#### 导入


- `__future__.annotations`

- `.typing` (as t)

- `jinja2.BaseLoader`

- `jinja2.Environment` (as BaseEnvironment)

- `jinja2.Template`

- `jinja2.TemplateNotFound`

- `.ctx.AppContext`

- `.globals.app_ctx`

- `.helpers.stream_with_context`

- `.signals.before_render_template`

- `.signals.template_rendered`




#### 代码元素


**Class**: `Environment`

- **完整名称**: Environment
- **文件位置**: 行 32-32
- **可见性**: public




- **继承**: BaseEnvironment 



- **子元素**: attempts rv trv loader result app ctx template get_source _get_source_explained _get_source_fast _iter_loaders list_templates _render render_template render_template_string _stream generate stream_template stream_template_string 

**Class**: `DispatchingJinjaLoader`

- **完整名称**: DispatchingJinjaLoader
- **文件位置**: 行 45-45
- **可见性**: public




- **继承**: BaseLoader 



- **子元素**: attempts rv trv loader result app ctx template get_source _get_source_explained _get_source_fast _iter_loaders list_templates _render render_template render_template_string _stream generate stream_template stream_template_string 




### src\flask\testing.py

- **语言**: python
- **包**: 
- **代码行数**: 227
- **元素数量**: 3


#### 导入


- `__future__.annotations`

- `importlib.metadata`

- `.typing` (as t)

- `contextlib.contextmanager`

- `contextlib.ExitStack`

- `copy.copy`

- `types.TracebackType`

- `urllib.parse.urlsplit`

- `werkzeug.test`

- `click.testing.CliRunner`

- `click.testing.Result`

- `werkzeug.test.Client`

- `werkzeug.wrappers.Request` (as BaseRequest)

- `.cli.ScriptInfo`

- `.sessions.SessionMixin`




#### 代码元素


**Class**: `FlaskClient`

- **完整名称**: FlaskClient
- **文件位置**: 行 106-106
- **可见性**: public




- **继承**: Client 



- **子元素**: app ctx sess resp out builder request response buffered follow_redirects cli session_transaction _copy_environ _request_from_builder_args open __enter__ __exit__ invoke 

**Class**: `FlaskCliRunner`

- **完整名称**: FlaskCliRunner
- **文件位置**: 行 262-262
- **可见性**: public




- **继承**: CliRunner 



- **子元素**: cli invoke 

**Variable**: `_werkzeug_version`

- **完整名称**: _werkzeug_version
- **文件位置**: 行 97-97
- **可见性**: public

- **类型**: `str`




- **值**: `""`





### src\flask\typing.py

- **语言**: python
- **包**: 
- **代码行数**: 63
- **元素数量**: 19


#### 导入


- `__future__.annotations`

- `.collections.abc` (as cabc)

- `.typing` (as t)




#### 代码元素


**Variable**: `ResponseValue`

- **完整名称**: ResponseValue
- **文件位置**: 行 12-12
- **可见性**: public

- **类型**: `Any`




- **值**: `t.Union[`


**Variable**: `HeaderValue`

- **完整名称**: HeaderValue
- **文件位置**: 行 26-26
- **可见性**: public

- **类型**: `Any`




- **值**: `str | list[str] | tuple[str, ...]`


**Variable**: `HeadersValue`

- **完整名称**: HeadersValue
- **文件位置**: 行 29-29
- **可见性**: public

- **类型**: `Any`




- **值**: `t.Union[`


**Variable**: `ResponseReturnValue`

- **完整名称**: ResponseReturnValue
- **文件位置**: 行 36-36
- **可见性**: public

- **类型**: `Any`




- **值**: `t.Union[`


**Variable**: `ResponseClass`

- **完整名称**: ResponseClass
- **文件位置**: 行 47-47
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("ResponseClass", bound="Response")`


**Variable**: `AppOrBlueprintKey`

- **完整名称**: AppOrBlueprintKey
- **文件位置**: 行 49-49
- **可见性**: public

- **类型**: `Any`




- **值**: `str | None  # The App key is None, whereas blueprints are named`


**Variable**: `AfterRequestCallable`

- **完整名称**: AfterRequestCallable
- **文件位置**: 行 50-50
- **可见性**: public

- **类型**: `tuple`




- **值**: `(`


**Variable**: `BeforeFirstRequestCallable`

- **完整名称**: BeforeFirstRequestCallable
- **文件位置**: 行 54-54
- **可见性**: public

- **类型**: `Any`




- **值**: `t.Callable[[], None] | t.Callable[[], t.Awaitable[None]]`


**Variable**: `BeforeRequestCallable`

- **完整名称**: BeforeRequestCallable
- **文件位置**: 行 55-55
- **可见性**: public

- **类型**: `tuple`




- **值**: `(`


**Variable**: `ShellContextProcessorCallable`

- **完整名称**: ShellContextProcessorCallable
- **文件位置**: 行 59-59
- **可见性**: public

- **类型**: `Any`




- **值**: `t.Callable[[], dict[str, t.Any]]`


**Variable**: `TeardownCallable`

- **完整名称**: TeardownCallable
- **文件位置**: 行 60-60
- **可见性**: public

- **类型**: `tuple`




- **值**: `(`


**Variable**: `TemplateContextProcessorCallable`

- **完整名称**: TemplateContextProcessorCallable
- **文件位置**: 行 64-64
- **可见性**: public

- **类型**: `tuple`




- **值**: `(`


**Variable**: `TemplateFilterCallable`

- **完整名称**: TemplateFilterCallable
- **文件位置**: 行 67-67
- **可见性**: public

- **类型**: `Any`




- **值**: `t.Callable[..., t.Any]`


**Variable**: `TemplateGlobalCallable`

- **完整名称**: TemplateGlobalCallable
- **文件位置**: 行 68-68
- **可见性**: public

- **类型**: `Any`




- **值**: `t.Callable[..., t.Any]`


**Variable**: `TemplateTestCallable`

- **完整名称**: TemplateTestCallable
- **文件位置**: 行 69-69
- **可见性**: public

- **类型**: `Any`




- **值**: `t.Callable[..., bool]`


**Variable**: `URLDefaultCallable`

- **完整名称**: URLDefaultCallable
- **文件位置**: 行 70-70
- **可见性**: public

- **类型**: `Any`




- **值**: `t.Callable[[str, dict[str, t.Any]], None]`


**Variable**: `URLValuePreprocessorCallable`

- **完整名称**: URLValuePreprocessorCallable
- **文件位置**: 行 71-71
- **可见性**: public

- **类型**: `Any`




- **值**: `t.Callable[[str | None, dict[str, t.Any] | None], None]`


**Variable**: `ErrorHandlerCallable`

- **完整名称**: ErrorHandlerCallable
- **文件位置**: 行 79-79
- **可见性**: public

- **类型**: `tuple`




- **值**: `(`


**Variable**: `RouteCallable`

- **完整名称**: RouteCallable
- **文件位置**: 行 84-84
- **可见性**: public

- **类型**: `tuple`




- **值**: `(`





### src\flask\views.py

- **语言**: python
- **包**: 
- **代码行数**: 116
- **元素数量**: 17


#### 导入


- `__future__.annotations`

- `.typing` (as t)

- `..typing` (as ft)

- `.globals.current_app`

- `.globals.request`




#### 代码元素


**Class**: `View`

- **完整名称**: View
- **文件位置**: 行 13-13
- **可见性**: public








- **子元素**: init_every_request view methods meth dispatch_request as_view view get post __init_subclass__ 

**Method**: `dispatch_request`

- **完整名称**: View.dispatch_request
- **文件位置**: 行 29-30
- **可见性**: public




- **参数**:
  - `self`: Any
  - `name`: Any






**Method**: `get`

- **完整名称**: View.get
- **文件位置**: 行 153-153
- **可见性**: public




- **参数**:
  - `self`: Any






**Method**: `post`

- **完整名称**: View.post
- **文件位置**: 行 155-158
- **可见性**: public




- **参数**:
  - `self`: Any






**Class**: `Hello`

- **完整名称**: Hello
- **文件位置**: 行 25-28
- **可见性**: public




- **继承**: View 



- **子元素**: init_every_request view methods meth dispatch_request as_view view get post __init_subclass__ 

**Method**: `dispatch_request`

- **完整名称**: Hello.dispatch_request
- **文件位置**: 行 29-30
- **可见性**: public




- **参数**:
  - `self`: Any
  - `name`: Any






**Method**: `get`

- **完整名称**: Hello.get
- **文件位置**: 行 153-153
- **可见性**: public




- **参数**:
  - `self`: Any






**Method**: `post`

- **完整名称**: Hello.post
- **文件位置**: 行 155-158
- **可见性**: public




- **参数**:
  - `self`: Any






**Class**: `MethodView`

- **完整名称**: MethodView
- **文件位置**: 行 135-135
- **可见性**: public




- **继承**: View 



- **子元素**: methods meth get post __init_subclass__ dispatch_request 

**Method**: `get`

- **完整名称**: MethodView.get
- **文件位置**: 行 153-153
- **可见性**: public




- **参数**:
  - `self`: Any






**Method**: `post`

- **完整名称**: MethodView.post
- **文件位置**: 行 155-158
- **可见性**: public




- **参数**:
  - `self`: Any






**Class**: `CounterAPI`

- **完整名称**: CounterAPI
- **文件位置**: 行 150-154
- **可见性**: public




- **继承**: MethodView 



- **子元素**: methods meth get post __init_subclass__ dispatch_request 

**Method**: `get`

- **完整名称**: CounterAPI.get
- **文件位置**: 行 153-153
- **可见性**: public




- **参数**:
  - `self`: Any






**Method**: `post`

- **完整名称**: CounterAPI.post
- **文件位置**: 行 155-158
- **可见性**: public




- **参数**:
  - `self`: Any






**Variable**: `F`

- **完整名称**: F
- **文件位置**: 行 9-9
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("F", bound=t.Callable[..., t.Any])`


**Variable**: `http_method_funcs`

- **完整名称**: http_method_funcs
- **文件位置**: 行 11-11
- **可见性**: public

- **类型**: `Any`




- **值**: `frozenset(`


**Constant**: `F`

- **完整名称**: F
- **文件位置**: 行 9-9
- **可见性**: public

- **类型**: `Any`




- **值**: `t.TypeVar("F", bound=t.Callable[..., t.Any])`





### src\flask\wrappers.py

- **语言**: python
- **包**: 
- **代码行数**: 175
- **元素数量**: 2


#### 导入


- `__future__.annotations`

- `.typing` (as t)

- `werkzeug.exceptions.BadRequest`

- `werkzeug.exceptions.HTTPException`

- `werkzeug.wrappers.Request` (as RequestBase)

- `werkzeug.wrappers.Response` (as ResponseBase)

- `..json`

- `.globals.current_app`

- `.helpers._split_blueprint_path`




#### 代码元素


**Class**: `Request`

- **完整名称**: Request
- **文件位置**: 行 15-15
- **可见性**: public




- **继承**: RequestBase 



- **子元素**: endpoint name json_module autocorrect_location_header max_content_length max_form_memory_size max_form_parts endpoint blueprint blueprints _load_form_data on_json_loading_failed max_cookie_size 

**Class**: `Response`

- **完整名称**: Response
- **文件位置**: 行 219-219
- **可见性**: public




- **继承**: ResponseBase 



- **子元素**: json_module autocorrect_location_header max_cookie_size aaaaaaaaaaaaaaaa aaaaaaaaaaaaaaaa aaaaaaaaaaaaaaaa aaaaaaaaaaaaaaaa aaaaaaaaaaaaaaaa aaaaaaaaaaaaaaaa aaaaaaaaaaaaaaaa aaaaaaaaaaaaaaaa aaaaaaaaaaaaaaaa aaaaaaaaaaaaaaaa aaaaaaaaaaaaaaaa aaaaaaaaaaaaaaaa aaaaaaaaaaaaaaaa aaaaaaaaaaaaaaaa 





## API 参考

### 公共接口



#### Function: github_link





**参数**:

- `name`: Any

- `rawtext`: Any

- `text`: Any

- `lineno`: Any

- `inliner`: Any

- `options`: Any (可选)

- `content`: Any (可选)





---




#### Function: setup





**参数**:

- `app`: Any





---




#### Variable: project








---




#### Variable: copyright








---




#### Variable: author








---




#### Variable: default_role








---




#### Variable: extensions








---




#### Variable: autodoc_member_order








---




#### Variable: autodoc_typehints








---




#### Variable: autodoc_preserve_defaults








---




#### Variable: extlinks








---




#### Variable: intersphinx_mapping








---




#### Variable: html_theme








---




#### Variable: html_theme_options








---




#### Variable: html_context








---




#### Variable: html_sidebars








---




#### Variable: singlehtml_sidebars








---




#### Variable: html_static_path








---




#### Variable: html_favicon








---




#### Variable: html_logo








---




#### Variable: html_title








---




#### Variable: html_show_sourcelink








---




#### Variable: gettext_uuid








---




#### Variable: gettext_compact








---




#### Variable: flask_app








---




#### Variable: celery_app








---




#### Class: FlaskTask








---




#### Variable: bp








---




#### Variable: app








---




#### Function: index





**参数**:

- `js`: Any





---




#### Function: add








---




#### Function: create_app





**参数**:

- `test_config`: Any (可选)





---




#### Function: login_required





**参数**:

- `view`: Any





---




#### Function: load_logged_in_user








---




#### Function: register








---




#### Function: login








---




#### Function: logout








---




#### Variable: bp








---




#### Function: index








---




#### Function: get_post





**参数**:

- `id`: Any

- `check_author`: Any (可选)





---




#### Function: create








---




#### Function: update





**参数**:

- `id`: Any





---




#### Function: delete





**参数**:

- `id`: Any





---




#### Variable: bp








---




#### Function: get_db








---




#### Function: close_db





**参数**:

- `e`: Any (可选)





---




#### Function: init_db








---




#### Function: init_db_command








---




#### Function: init_app





**参数**:

- `app`: Any





---




#### Class: Flask








---




#### Class: CustomClient








---




#### Variable: T_shell_context_processor








---




#### Variable: T_teardown








---




#### Variable: T_template_filter








---




#### Variable: T_template_global








---




#### Variable: T_template_test








---




#### Variable: F








---




#### Constant: F








---




#### Class: Blueprint








---




#### Class: ScriptInfo








---




#### Class: FlaskGroup








---




#### Variable: version_option








---




#### Variable: pass_script_info








---




#### Variable: F








---




#### Variable: _app_option








---




#### Variable: _debug_option








---




#### Variable: _env_file_option








---




#### Variable: cli








---




#### Constant: F








---




#### Variable: T








---




#### Constant: T








---




#### Class: _AppCtxGlobals








---




#### Method: index








---




#### Method: add_header





**参数**:

- `response`: Any





---




#### Method: do_some_work








---




#### Class: AppContext








---




#### Variable: _sentinel








---




#### Variable: F








---




#### Constant: F








---




#### Class: UnexpectedUnicodeError








---




#### Class: DebugFilesKeyError








---




#### Class: FormDataRoutingRedirect








---




#### Variable: _no_app_msg








---




#### Variable: _no_req_msg








---




#### Class: JSONProvider








---




#### Class: DefaultJSONProvider








---




#### Class: TagOrderedDict








---




#### Method: check





**参数**:

- `self`: Any

- `value`: Any





---




#### Method: to_json





**参数**:

- `self`: Any

- `value`: Any





---




#### Method: to_python





**参数**:

- `self`: Any

- `value`: Any





---




#### Class: JSONTag








---




#### Class: TagDict








---




#### Class: PassDict








---




#### Class: TagTuple








---




#### Class: PassList








---




#### Class: TagBytes








---




#### Class: TagMarkup








---




#### Class: TagUUID








---




#### Class: TagDateTime








---




#### Class: TaggedJSONSerializer








---




#### Variable: default_handler








---




#### Class: App








---




#### Method: reverse_filter





**参数**:

- `s`: Any





---




#### Method: is_prime_test





**参数**:

- `n`: Any





---




#### Method: double





**参数**:

- `n`: Any





---




#### Variable: T_shell_context_processor








---




#### Variable: T_teardown








---




#### Variable: T_template_filter








---




#### Variable: T_template_global








---




#### Variable: T_template_test








---




#### Class: BlueprintSetupState








---




#### Class: Blueprint








---




#### Variable: DeferredSetupFunction








---




#### Variable: T_after_request








---




#### Variable: T_before_request








---




#### Variable: T_error_handler








---




#### Variable: T_teardown








---




#### Variable: T_template_context_processor








---




#### Variable: T_template_filter








---




#### Variable: T_template_global








---




#### Variable: T_template_test








---




#### Variable: T_url_defaults








---




#### Variable: T_url_value_preprocessor








---




#### Class: Scaffold








---




#### Method: index








---




#### Method: example








---




#### Method: load_user








---




#### Method: page_not_found





**参数**:

- `error`: Any





---




#### Method: special_exception_handler





**参数**:

- `error`: Any





---




#### Variable: _sentinel








---




#### Variable: F








---




#### Variable: T_after_request








---




#### Variable: T_before_request








---




#### Variable: T_error_handler








---




#### Variable: T_teardown








---




#### Variable: T_template_context_processor








---




#### Variable: T_url_defaults








---




#### Variable: T_url_value_preprocessor








---




#### Variable: T_route








---




#### Constant: F








---




#### Class: NullSession








---




#### Class: SessionInterface








---




#### Class: SecureCookieSessionInterface








---




#### Variable: session_json_serializer








---




#### Variable: _signals








---




#### Variable: template_rendered








---




#### Variable: before_render_template








---




#### Variable: request_started








---




#### Variable: request_finished








---




#### Variable: request_tearing_down








---




#### Variable: got_request_exception








---




#### Variable: appcontext_tearing_down








---




#### Variable: appcontext_pushed








---




#### Variable: appcontext_popped








---




#### Variable: message_flashed








---




#### Class: Environment








---




#### Class: DispatchingJinjaLoader








---




#### Class: FlaskClient








---




#### Class: FlaskCliRunner








---




#### Variable: _werkzeug_version








---




#### Variable: ResponseValue








---




#### Variable: HeaderValue








---




#### Variable: HeadersValue








---




#### Variable: ResponseReturnValue








---




#### Variable: ResponseClass








---




#### Variable: AppOrBlueprintKey








---




#### Variable: AfterRequestCallable








---




#### Variable: BeforeFirstRequestCallable








---




#### Variable: BeforeRequestCallable








---




#### Variable: ShellContextProcessorCallable








---




#### Variable: TeardownCallable








---




#### Variable: TemplateContextProcessorCallable








---




#### Variable: TemplateFilterCallable








---




#### Variable: TemplateGlobalCallable








---




#### Variable: TemplateTestCallable








---




#### Variable: URLDefaultCallable








---




#### Variable: URLValuePreprocessorCallable








---




#### Variable: ErrorHandlerCallable








---




#### Variable: RouteCallable








---




#### Class: View








---




#### Method: dispatch_request





**参数**:

- `self`: Any

- `name`: Any





---




#### Method: get





**参数**:

- `self`: Any





---




#### Method: post





**参数**:

- `self`: Any





---




#### Class: Hello








---




#### Method: dispatch_request





**参数**:

- `self`: Any

- `name`: Any





---




#### Method: get





**参数**:

- `self`: Any





---




#### Method: post





**参数**:

- `self`: Any





---




#### Class: MethodView








---




#### Method: get





**参数**:

- `self`: Any





---




#### Method: post





**参数**:

- `self`: Any





---




#### Class: CounterAPI








---




#### Method: get





**参数**:

- `self`: Any





---




#### Method: post





**参数**:

- `self`: Any





---




#### Variable: F








---




#### Variable: http_method_funcs








---




#### Constant: F








---




#### Class: Request








---




#### Class: Response








---




## 调用关系图


### github_link (function)

- **文件**: E:\cody\goproject\qcoder\mcp-server\code2md\test_projects\py-test\docs\conf.py:72
- **调用深度**: 0




