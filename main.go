package main

import (
	"net/http"
	"github.com/yang-f/beauty/utils/log"
	"github.com/yang-f/beauty/router"
	"github.com/yang-f/beauty/settings"
)

func main() {
	log.Printf("start server on port %s", settings.Listen)
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(settings.Listen, router))
}
