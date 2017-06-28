* [English](README.md)
* [中文](README_ZH.md)

[![GoDoc](https://godoc.org/github.com/imroc/req?status.svg)](https://godoc.org/github.com/imroc/req)
[![yang-f](https://avatars1.githubusercontent.com/u/4221375?v=3&s=460)](https://github.com/yang-f)

👩这是一个Golang实现的简易框架
==============================

你通过它可以实现一个简单的restful工程。

快速开始:
------------------------------
* 执行命令
    ```
    mkdir demo && cd demo
    go get gopkg.in/alecthomas/kingpin.v2
    go get github.com/yang-f/beauty
    ```
* 将$GOPATH/bin 加入到$PATH 环境变量
* 执行命令
    ```
    beauty
    ```
* 这时显示如下
    ```
        usage: beauty [<flags>] <command> [<args> ...]
        
        A command-line tools of beauty.
        
        Flags:
          --help  Show context-sensitive help (also try --help-long and --help-man).
        
        Commands:
          help [<command>...]
            Show help.
        
          demo
            Demo of web server.
        
          generate <name>
            Generate a new app.
    ```
* 测试beauty
    ```
    beauty demo
    ```
* 这时terminal显示
    ```golang
    2017/05/04 16:21:05 start server on port :8080
    ```
* 通过浏览器访问 127.0.0.1:8080
    ```golang
    {"description":"this is json"}
    ```

* 访问 127.0.0.1:8080/demo1
    ```golang
    {"status":403,"description":"token not found.","code":"AUTH_FAILED"}
    ```

* 访问 127.0.0.1:8080/demo2
    ```golang
    {"description":"this is json"}
    ```
    
* 访问 127.0.0.1:8080/demo3
    ```golang
    {"status":403,"description":"token not found.","code":"AUTH_FAILED"}
    ```
    
* 恭喜你，运行成功。

如何使用:
-------------------------------
* 生成app
    ```
    beauty generate app的名字
    ```
* 生成的app列表
    >GOPATH
    >>src
    >>>yourAppName
    >>>>controllers 
    >>>>>adminController.go
    >>>>>
    >>>>>controller_test.go
    >>>>
    >>>>models
    >>>>
    >>>>routers
    >>>>>routers.go
    >>>>
    >>>>tpl
    >>>>
    >>>>utils
    >>>>
    >>>>main.go
* 关于路由
    ```golang
    type Route struct {
        Name        string //日志中显示的路由名称
        Method      string //GET PUT POST DELETE ...
        Pattern     string //对用的访问路径
        HandlerFunc decorates.Handler //处理当前路由的Controller或者handler
        ContentType string //返回的数据类型"application/json;charset=utf-8" 或者 "text/html" 等等 ...
    }
    ```
    * 例子
    ```golang
        import (
            . "github.com/yang-f/beauty/decorates"

        )
        router.BRoutes = router.Routes{
            router.Route{
                "nothing",
                "GET",
                "/",
                controllers.Config,
                contenttype.JSON,
            },//正常路由
            router.Route{
                "authDemo",
                "GET",
                "/demo1",
                Handler(controllers.Config).
                    Auth(),
                contenttype.JSON,
            },//需要验证用户信息
            router.Route{
                "verifyDemo",
                "GET",
                "/demo2",
                Handler(controllers.Config).
                    Verify(),
                contenttype.JSON,
            },//需要过滤非法提交信息，比如SQL注入
            router.Route{
                "verifyAndAuthDemo",
                "GET",
                "/demo3",
                Handler(controllers.Config).
                    Auth().
                    Verify(),
                contenttype.JSON,
            },//都需要，也可以自己扩展，参考decorate实现。
        }
    ```
* token生成
    ```golang
    tokenString, err := token.Generate(fmt.Sprintf("%v|%v", user_id, user_pass))

    ```
    
* 小例子
    ```golang
    package main

    import (
        "net/http"
        "github.com/yang-f/beauty/consts/contenttype"
        "github.com/yang-f/beauty/utils/log"
        "github.com/yang-f/beauty/router"
        "github.com/yang-f/beauty/settings"
        "github.com/yang-f/beauty/controllers"
        . "github.com/yang-f/beauty/decorates"

    )

    func main() {

        log.Printf("start server on port %s", settings.Listen)

        router.BRoutes = router.Routes{
            router.Route{
                "nothing",
                "GET",
                "/",
                controllers.Config,
                contenttype.JSON,
            },
            router.Route{
                "authDemo",
                "GET",
                "/demo1",
                Handler(controllers.Config).
                    Auth(),
                contenttype.JSON,
            },
            router.Route{
                "verifyDemo",
                "GET",
                "/demo2",
                Handler(controllers.Config).
                    Verify(),
                contenttype.JSON,
            },
            router.Route{
                "verifyAndAuthDemo",
                "GET",
                "/demo3",
                Handler(controllers.Config).
                    Auth().
                    Verify(),
                contenttype.JSON,
            },
        }

        settings.Listen = ":8080"//服务运行端口

        settings.Domain = "yourdomain.com"//部署服务的域名

        settings.LogFile = "/your/path/yourname.log"//日志所在文件

        settings.DefaultOrigin = "http://defaultorigin.com"//默认的请求来源

        settings.HmacSampleSecret = "whatever"//令牌生产需要的字符串

        router := router.NewRouter()

        log.Fatal(http.ListenAndServe(settings.Listen, router))

    }
    ```

支持特性:
--------------------------
* 参数校验
    * 这是一个统一的参数校验，主要是针对SQL注入，有了它，或许就不用一个一个参数做校验了。
    * 使用方式
    ```golang
    router.Route{
        "authDemo",
        "GET",
        "/demo1",
        Handler(controllers.Config).
            Verify(),
        contenttype.JSON,
    },//需要验证用户信息
    ```
* 令牌 
    ```golang
    settings.HmacSampleSecret = "whatever"

    token, err := token.Generate(origin)
    
    origin, err := token.Valid(token)
    ```
* 数据库操作（基于mymysql）
    ```golang
    db.Query(sql, params...)
    ```
* cors跨域
    * 静态文件
    ```golang
    router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", decorates.CorsHeader2(http.FileServer(http.Dir("/your/static/path")))))
    ```
    * 其他: 
        * 默认是支持跨域操作的

* 日志
    * 使用
    ```golang
    settings.LogFile = "/you/log/path/beauty.log"

    log.Printf(msg, params...)
    ```
    * 每隔12小时且日志大于50M自动归档
* 会话
    ```golang
    currentUser := sessions.CurrentUser(r *http.Request)
    ```
* 错误处理以及http状态管理
    ```golang
    func XxxxController(w http.ResponseWriter, r *http.Request) *models.APPError {
        xxx,err := someOperation()
        if err != nil{
            return &models.APPError {err, Message, Code, Status}
        }
        ...
        return nil
    }
    ```

* 工具
    * Response
    * Rand
    * MD5
    * Post

* 测试
    * go test -v -bench=".*"
    * go run controllers/controller_test.go
    * ...

其他:
-------------------------------------------------------
    
* sql
    ```golang
    create database yourdatabase;
    use yourdatabase;
    create table if not exists user
    (
        user_id int primary key not null  auto_increment,
        user_name varchar(64),
        user_pass varchar(64),
        user_mobile varchar(32),
        user_type enum('user', 'admin', 'test') not null,
        add_time timestamp not null default CURRENT_TIMESTAMP
    );

    insert into user (user_name, user_pass) values('admin', 'admin');
    ``` 
* 你需要编辑一个配置文件'/srv/filestore/settings/latest.json' 格式如下：
    ```golang
    {
        "mysql_host":"127.0.0.1:3306",
        "mysql_user":"root",
        "mysql_pass":"root",
        "mysql_database":"yourdatabase"
    }
    ```


贡献代码:
---------------------------------

1. Fork代码!
2. 创建新的分支: `git checkout -b my-new-feature`
3. 提交更改: `git commit -m 'Add some feature'`
4. push到分支: `git push origin my-new-feature`
5. 发起一个pull request :D

将要实现:
----------------------------------

* [x] 命令行工具
* [ ] 继续提升文档质量
* [ ] 权限控制
* [ ] 增加测试覆盖率
