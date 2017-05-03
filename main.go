package main

import (
	"github.com/yang-f/beauty/router"
	"github.com/yang-f/beauty/settings"
	"github.com/yang-f/beauty/utils/log"
	"net/http"
)

func main() {
	log.Printf("start server on port %s", settings.Listen)
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(settings.Listen, router))
}
