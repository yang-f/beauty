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

    settings.Domain = "youdomain.com"

    router := router.NewRouter()

    log.Fatal(http.ListenAndServe(settings.Listen, router))

}
```

Support:
--------------------------

* token 
    ```golang
    token, err := token.Generate(user_id)
    
    user_id, err := token.Valid(token)
    ```
* db
    * db.Query(sql, params...)
* cors

* log
    * settings.LogFile = "/you/log/path/beauty.log"
