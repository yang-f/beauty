package router

import (
	"github.com/gorilla/mux"
	"github.com/yang-f/beauty/decorates"
	"github.com/yang-f/beauty/models"
	"github.com/yang-f/beauty/utils/log"
	"net/http"
)

var router *mux.Router

func NewRouter() *mux.Router {
	if router == nil {
		router = mux.NewRouter().StrictSlash(true)
	}
	for _, route := range routes {
		var handler decorates.Handler = decorates.CorsHeader(route.HandlerFunc)
		if route.Auth {
			handler = decorates.Auth(handler)
		}
		handler = decorates.ContentType(handler, route.ContentType)
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
			Handler(decorates.CorsHeader(decorates.Handler(func(w http.ResponseWriter, r *http.Request) *models.APPError {
				return nil
			})))
	}

	return router
}
