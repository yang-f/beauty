// MIT License

// Copyright (c) 2017 FLYING

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yang-f/beauty/decorates"
	"github.com/yang-f/beauty/models"
)

var router *Router

type Router struct {
	*mux.Router
}

func New() *Router {
	if router == nil {
		router = &Router{
			mux.NewRouter().StrictSlash(true),
		}
	}
	for _, route := range BRoutes {
		handler := route.Handler.
			CorsHeader().
			ContentType(route.ContentType).
			Logger(route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
		router.
			Methods("OPTIONS").
			Path(route.Pattern).
			Name("cors").
			Handler(
				decorates.Handler(
					func(w http.ResponseWriter, r *http.Request) *models.APPError {
						return nil
					},
				).CorsHeader(),
			)
	}

	return router
}
