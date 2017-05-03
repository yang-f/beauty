package router

import (
	"github.com/yang-f/beauty/controllers"
	"github.com/yang-f/beauty/decorates"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Auth        bool
	HandlerFunc decorates.Handler
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
