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

func (r *Router) GET(path string, handler decorates.Handler) {
	r.register("GET", path, handler)
}

func (r *Router) POST(path string, handler decorates.Handler) {
	r.register("POST", path, handler)
}

func (r *Router) PUT(path string, handler decorates.Handler) {
	r.register("PUT", path, handler)
}

func (r *Router) TRACE(path string, handler decorates.Handler) {
	r.register("TRACE", path, handler)
}

func (r *Router) HEAD(path string, handler decorates.Handler) {
	r.register("HEAD", path, handler)
}

func (r *Router) OPTIONS(path string, handler decorates.Handler) {
	r.register("OPTIONS", path, handler)
}

func (r *Router) LOCK(path string, handler decorates.Handler) {
	r.register("LOCK", path, handler)
}

func (r *Router) DELETE(path string, handler decorates.Handler) {
	r.register("DELETE", path, handler)
}

func (r *Router) register(method, path string, handler decorates.Handler) {
	route := &Route{
		Method:  method,
		Handler: handler,
		Pattern: path,
	}
	handler = route.Handler.
		CorsHeader().
		Logger()
	router.
		Methods(route.Method).
		Path(route.Pattern).
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

func New() *Router {
	if router == nil {
		router = &Router{
			mux.NewRouter().StrictSlash(true),
		}
	}
	return router
}
