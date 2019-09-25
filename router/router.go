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
	"github.com/gorilla/mux"
)

var router *Router

type Router struct {
	*mux.Router
}

func (r *Router) GET(path string, handler func(c *Context) error) {
	r.register("GET", path, handler)
}

func (r *Router) POST(path string, handler func(c *Context) error) {
	r.register("POST", path, handler)
}

func (r *Router) PUT(path string, handler func(c *Context) error) {
	r.register("PUT", path, handler)
}

func (r *Router) TRACE(path string, handler func(c *Context) error) {
	r.register("TRACE", path, handler)
}

func (r *Router) HEAD(path string, handler func(c *Context) error) {
	r.register("HEAD", path, handler)
}

func (r *Router) OPTIONS(path string, handler func(c *Context) error) {
	r.register("OPTIONS", path, handler)
}

func (r *Router) LOCK(path string, handler func(c *Context) error) {
	r.register("LOCK", path, handler)
}

func (r *Router) DELETE(path string, handler func(c *Context) error) {
	r.register("DELETE", path, handler)
}

func (r *Router) register(method, path string, handlerFunc func(c *Context) error) {
	h := handler(handlerFunc).logger().corsHeader()
	router.
		Methods(method).
		Path(path).
		Handler(h)

	router.
		Methods("OPTIONS").
		Path(path).
		Name("cors").
		Handler(
			handler(
				func(c *Context) error {
					return nil
				},
			).corsHeader(),
		)
}

func New() *Router {
	if router != nil {
		return router
	}
	router = &Router{
		mux.NewRouter().StrictSlash(true),
	}
	return router
}
