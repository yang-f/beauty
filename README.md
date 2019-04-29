- [English](README.md)
- [中文](README_ZH.md)

[![GoDoc](https://godoc.org/github.com/yang-f/beauty?status.svg)](https://godoc.org/github.com/yang-f/beauty)

# A simple framwork written in golang.

You can build a simple restful project or a web application with it.
If you dosen't want to use mysql db, you can implement your own Auth decorates and session.You can use your own DAO or whatever you like.

## quick start:

- run cmd
  ```
  mkdir demo && cd demo
  go get gopkg.in/alecthomas/kingpin.v2
  go get github.com/yang-f/beauty
  ```
- add $GOPATH/bin to your $PATH

- run cmd beauty

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

- test beauty
  ```
  beauty demo
  ```
- then
  ```golang
  2017/05/04 16:21:05 start server on port :8080
  ```
- visit 127.0.0.1:8080

  ```golang
  {"description":"this is json"}
  ```

- visit 127.0.0.1:8080/demo1

  ```golang
  {"status":403,"description":"token not found.","code":"AUTH_FAILED"}
  ```

- visit 127.0.0.1:8080/demo2
  ```golang
  {"description":"this is json"}
  ```
- visit 127.0.0.1:8080/demo3
  ```golang
  {"status":403,"description":"token not found.","code":"AUTH_FAILED"}
  ```

## How to use:

- Generate new app
  ```
  beauty generate yourAppName
  ```
- dir list

  ```
  GOPATH/src/yourAppName
  ├── controllers
  │   ├── adminController.go
  │   └── controller_test.go
  ├── decorates
  ├   └── http.go
  ├── main.go
  ├── models
  ├── tpl
  └── utils
  ```

- about Route

  - demo

  ```golang
      r := router.New()

      r.GET("/", controllers.Config().ContentJSON())

      r.GET("/demo1", controllers.Config().ContentJSON().Auth())

      r.GET("/demo2", controllers.Config().ContentJSON().Verify())

      r.GET("/demo3", controllers.Config().ContentJSON().Auth().Verify())

  ```

- token generate

  ```golang
  tokenString, err := token.Generate(fmt.Sprintf("%v|%v", user_id, user_pass))

  ```

- demo

  ```golang
  package main

  import (
      "net/http"
      "github.com/yang-f/beauty/consts/contenttype"
      "github.com/yang-f/beauty/utils/log"
      "github.com/yang-f/beauty/router"
      "github.com/yang-f/beauty/settings"
      "github.com/yang-f/beauty/controllers"
      "github.com/yang-f/beauty/decorates"

  )

  func main() {

      log.Printf("start server on port %s", settings.Listen)

      settings.Listen = ":8080"

      settings.Domain = "yourdomain.com"

      settings.LogFile = "/your/path/yourname.log"

      settings.DefaultOrigin = "http://defaultorigin.com"

      settings.HmacSampleSecret = "whatever"

      r := router.New()

      r.GET("/", controllers.Config().ContentJSON())

      r.GET("/demo1", controllers.Config().ContentJSON().Auth())

      r.GET("/demo2", controllers.Config().ContentJSON().Verify())

      r.GET("/demo3", controllers.Config().ContentJSON().Auth().Verify())

      log.Fatal(http.ListenAndServe(settings.Listen, r))

  }
  ```

## Support:
