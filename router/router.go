package router

import (
	"net/http"
	"github.com/yang-f/beauty/decorates"
	"github.com/gorilla/mux"
)

var router *mux.Router

func NewRouter() *mux.Router {
	if router == nil {
		router = mux.NewRouter().StrictSlash(true)
	}
	for _, route := range BRoutes {
		var handler http.Handler = decorates.CorsHeader(route.HandlerFunc)
		if(route.Auth){
			handler = decorates.Auth(handler)
		}
		handler = decorates.ContentType(handler,route.ContentType)
		handler = decorates.Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
		router.
			Methods("OPTIONS").
			Path(route.Pattern).
			Name("cors").
			Handler(decorates.CorsHeader(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				return
			})))
	}

	return router
}
