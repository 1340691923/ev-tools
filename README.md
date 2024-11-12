# ev tools

 [中文](./README-cn.md) | English 

### ev tools is an ElasticView plug-in for managing elasticsearch versions 6, 7, and 8

* Supports conversion of sql to dsl syntax
* More convenient index rebuilding
* Task management
* Backup management
* Can download query content as excel file
* Can perform index creation, mapping creation, alias creation, index deletion and other operations
* Support versions `6.x`, `7.x`, `8.x`
* Supports similar functions to Navicat

```
Project start-up preparation:

Base start:
Make sure the ElasticView base program has been started

Environment configuration:
golang version >= 1.20
node version >= 20.14.0

install gowatch:
go install github.com/silenceper/gowatch@latest

Install auxiliary tools to the project root directory:
go build -o ev_plugin_zip .\cmd\ev_plugin_zip\main.go
go build -o ev_plugin_builder .\cmd\ev_plugin_builder\main.go

install pnpm:
npm install -g pnpm

```

### Project infrastructure

```
-backend                      Backend project directory
    -api                      Controller layer
    -dto                      Web request structure
    -migrate                  SQLite data table version control module
    -versions                 Stores each version of the data table structure upgrade rollback
    -model                    Data access layer
    -my_error                 Template custom exception package
    -response                 Template custom response package
    -router                   Backend routing definition module
    -vo                       Web response structure
-cmd
    -ev_plugin_builder        Toolkit Used to compile binary plug-ins for various operating systems (windows, linux, darwin)
    -ev_plugin_zip            Toolkit Used to package project source code into zip
-frontend Frontend            project directory
    -dist                     Frontend final package file
    -src
    -api                      Interface access layer
    -lang                     Language package
    -layouts                  Default layout
    -plugin_sdk               Plugin sdk
    -router                   Frontend routing
    -views                    Page file
```

### plugin.json：

```json
{
  "developer": "xiaowenlong", //Developer Name
  "plugin_alias": "ev-tools", //Plugin id is also called plugin alias
  "plugin_name":"ev工具箱", //plugin display name
  "frontend_debug": false, // Whether to enable front-end page debugging
  "version": "0.0.1",     //current version
  "main_go_file": "main.go", //  main.go file position
  "frontend_dev_port":7001, //Front-end project startup port
  "frontend_routes": [  // The front-end routing is consistent with routes
    {
      "path": "hello-world", 
      "name": "HelloWorld",
      "meta": {
        "title": "HelloWorld",
        "icon": "el-icon-coin"
      }
    },
    {
      "path": "db-test",
      "name": "db-test",
      "meta": {
        "title": "操作数据库",
        "icon": "el-icon-coin"
      }
    }
  ],
  "backend_routes": [   //Backend Routing
    {
      "path": "/api/HelloWorld",  //Interface access path
      "remark": "HelloWorld测试接口", 
      "needAuth": false       //Is authentication required?
    }
  ]
}




```   




#### Backend plugin development

```
Project root directory execution:
1. go mod tidy
2. gowatch

```

#### Front-end plugin development
```

Key point: Modify frontend_debug in plugin.json to true

1. cd frontend && pnpm i

2. npm run dev

```

#### Front-end plugin packaging
```

1. cd frontend && npm run build

```

#### Back-end plugin packaging

```
Key point: Modify frontend_debug in plugin.json to false

Project root directory execution:
./ev_plugin_builder

A compressed package will appear in the dist directory. After decompression, it will be the plugin binary for different operating systems.

```
### Backend plugin source package

```

Execute in the project root directory:
ev_plugin_zip -source=  //source = need package folder directory address

```

