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

	    controllers.Config,

	    "application/json;charset=utf-8",

	},

    }

    settings.Listen = ":8080"

    settings.Domain = "yourdomain.com"

    settings.LogFile = "/your/path/yourname.log"

    router := router.NewRouter()

    log.Fatal(http.ListenAndServe(settings.Listen, router))

}
```

Support:
--------------------------

* token 
    ```golang
    settings.HmacSampleSecret = "whatever"

    token, err := token.Generate(user_id)
    
    user_id, err := token.Valid(token)
    ```
* db
    ```golang
    db.Query(sql, params...)
    ```
* cors

* log
    ```golang
    settings.LogFile = "/you/log/path/beauty.log"
    log.Printf(msg, params...)
    ```
* sessions
    ```golang
    currentUser := sessions.CurrentUser(r *http.Request)
    ```
* utils
    * Response
    * Rand
    * MD5
    * Post

Etc:
    you need set a json file in '/srv/filestore/settings/latest.json' fomart this:
    ```json
    {
        "mysql_host":"127.0.0.1:3306",
        "mysql_user":"root",
        "mysql_pass":"root",
        "mysql_database":"beauty"
    }
    ```


Contributing:
---------------------------------

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -m 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D
