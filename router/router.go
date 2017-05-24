package router

import (
	"github.com/gorilla/mux"
	. "github.com/yang-f/beauty/decorates"
	"github.com/yang-f/beauty/models"
	"net/http"
)

var router *mux.Router

func NewRouter() *mux.Router {
	if router == nil {
		router = mux.NewRouter().StrictSlash(true)
	}
	for _, route := range BRoutes {
		var handler Handler = CorsHeader(route.HandlerFunc)
		handler = ContentType(handler, route.ContentType)
		handler = Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
		router.
			Methods("OPTIONS").
			Path(route.Pattern).
			Name("cors").
			Handler(CorsHeader(Handler(func(w http.ResponseWriter, r *http.Request) *models.APPError {
				return nil
			})))
	}

	return router
}
