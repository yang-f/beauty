A simple framwork with golang.
==============================

You can build a simple restful project with it.


How to use:
-------------------------------

```golang
package main

import (

    "net/http"

    "github.com/yang-f/beauty/utils/log"

    "github.com/yang-f/beauty/router"

    "github.com/yang-f/beauty/settings"

    "github.com/yang-f/beauty/controllers"

    "github.com/yang-f/beauty/decorates"

)

func main() {

    log.Printf("start server on port %s", settings.Listen)

    router.BRoutes = router.Routes{

	router.Route{

	    "getConfig",

	    "GET",

	    "/",

	    false,

	    controllers.XxxxController,

	    "application/json;charset=utf-8",

	},

    }

    settings.Listen = ":8080"

    settings.Domain = "yourdomain.com"

    settings.LogFile = "/your/path/yourname.log"

    settings.DefaultOrigin = "http://defaultorigin.com"

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
    * etc: 
        * set with Route

* log
    ```golang
    settings.LogFile = "/you/log/path/beauty.log"

    log.Printf(msg, params...)
    ```
* sessions
    ```golang
    currentUser := sessions.CurrentUser(r *http.Request)
    ```
* error handler
    ```golang
    func XxxxController(w http.ResponseWriter, r *http.Request) *models.APPError {
        xxx,err := someOprate()
        return &models.APPError {err, Message, Code, Status}
    }
    ```

* utils
    * Response
    * Rand
    * MD5
    * Post

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
