package router

import (
	"net/http"
	"github.com/yang-f/beauty/controllers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Auth		bool
	HandlerFunc http.HandlerFunc
	ContentType string
}

type Routes []Route

var BRoutes = Routes{
	Route{
		"msg",
		"GET",
		"/",
		false,
		controllers.Config,
		"application/json;charset=utf-8",
	},
}
