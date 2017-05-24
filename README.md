* [English](README.md)
* [中文](README_ZH.md)

A simple framwork with golang.
==============================

You can build a simple restful project with it.

quick start:
------------------------------
* run cmd
    ```
    mkdir demo && cd demo
    go get github.com/yang-f/beauty
    ```
* add $GOPATH/bin to your $PATH
* run cmd
    ```
    beauty
    ```
* then
    ```golang
    2017/05/04 16:21:05 start server on port :8080
    ```
* visit 127.0.0.1:8080
    ```golang
    {"description":"this is json"}
    ```

* visit 127.0.0.1:8080/demo1
    ```golang
    {"status":403,"description":"token not found.","code":"AUTH_FAILED"}
    ```

* visit 127.0.0.1:8080/demo2
    ```golang
    {"description":"this is json"}
    ```
    
* visit 127.0.0.1:8080/demo3
    ```golang
    {"status":403,"description":"token not found.","code":"AUTH_FAILED"}
    ```

How to use:
-------------------------------

* about Route
    ```golang
    type Route struct {
        Name        string //Show route name in log file
        Method      string //GET PUT POST DELETE ...
        Pattern     string //Path
        HandlerFunc decorates.Handler //Controller or handler
        ContentType string //"application/json;charset=utf-8" or "text/html" etc...
    }
    ```
    * demo
    ```golang
        router.BRoutes = router.Routes{
            router.Route{
                "nothing",
                "GET",
                "/",
                controllers.Config,
                "application/json;charset=utf-8",
            },//normal route
            router.Route{
                "authDemo",
                "GET",
                "/demo1",
                Handler(controllers.Config).
                    Auth(),
                "application/json;charset=utf-8",
            },//route need auth
            router.Route{
                "verifyDemo",
                "GET",
                "/demo2",
                Handler(controllers.Config).
                    Verify(),
                "application/json;charset=utf-8",
            },//route need verify, such as sql injection
            router.Route{
                "verifyAndAuthDemo",
                "GET",
                "/demo3",
                Handler(controllers.Config).
                    Auth().
                    Verify(),
                "application/json;charset=utf-8",
            },//both
        }
    ```
* token generate
    ```golang
    tokenString, err := token.Generate(fmt.Sprintf("%v|%v", user_id, user_pass))

    ```
    
* demo
    ```golang
    package main

    import (
        "net/http"
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
                "application/json;charset=utf-8",
            },
            router.Route{
                "authDemo",
                "GET",
                "/demo1",
                Handler(controllers.Config).
                    Auth(),
                "application/json;charset=utf-8",
            },
            router.Route{
                "verifyDemo",
                "GET",
                "/demo2",
                Handler(controllers.Config).
                    Verify(),
                "application/json;charset=utf-8",
            },
            router.Route{
                "verifyAndAuthDemo",
                "GET",
                "/demo3",
                Handler(controllers.Config).
                    Auth().
                    Verify(),
                "application/json;charset=utf-8",
            },
        }

        settings.Listen = ":8080"

        settings.Domain = "yourdomain.com"

        settings.LogFile = "/your/path/yourname.log"

        settings.DefaultOrigin = "http://defaultorigin.com"

        settings.HmacSampleSecret = "whatever"

        router := router.NewRouter()

        log.Fatal(http.ListenAndServe(settings.Listen, router))

    }
    ```

Support:
--------------------------

* token 
    ```golang
    settings.HmacSampleSecret = "whatever"

    token, err := token.Generate(origin)
    
    origin, err := token.Valid(token)
    ```
* db
    ```golang
    db.Query(sql, params...)
    ```
* cors
    * static file server
    ```golang
    router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", decorates.CorsHeader2(http.FileServer(http.Dir("/your/static/path")))))
    ```
    * api etc: 
        * default is cors

* log
    * use
    ```golang
    settings.LogFile = "/you/log/path/beauty.log"

    log.Printf(msg, params...)
    ```
    * auto archive
* sessions
    ```golang
    currentUser := sessions.CurrentUser(r *http.Request)
    ```
* error handler
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

* utils
    * Response
    * Rand
    * MD5
    * Post

* test
    * go test -v -bench=".*"
    * go run controllers/controller_test.go
    * ...

Etc:
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
* you need set a json file in '/srv/filestore/settings/latest.json' fomart like this
    ```golang
    {
        "mysql_host":"127.0.0.1:3306",
        "mysql_user":"root",
        "mysql_pass":"root",
        "mysql_database":"yourdatabase"
    }
    ```


Contributing:
---------------------------------

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -m 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D

TODO:
----------------------------------

* [ ] Cmd tools
* [ ] Improve document
* [ ] Role review
