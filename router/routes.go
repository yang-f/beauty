package router

import (
	. "github.com/yang-f/beauty/controllers"
	. "github.com/yang-f/beauty/decorates"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc Handler
	ContentType string
}

type Routes []Route

var BRoutes = Routes{
	Route{
		"nothing",
		"GET",
		"/",
		Config,
		"application/json;charset=utf-8",
	},
	Route{
		"authDemo",
		"GET",
		"/demo1",
		Handler(Config).
			Auth(),
		"application/json;charset=utf-8",
	},
	Route{
		"verifyDemo",
		"GET",
		"/demo2",
		Handler(Config).
			Verify(),
		"application/json;charset=utf-8",
	},
	Route{
		"verifyAndAuthDemo",
		"GET",
		"/demo3",
		Handler(Config).
			Auth().
			Verify(),
		"application/json;charset=utf-8",
	},
}
