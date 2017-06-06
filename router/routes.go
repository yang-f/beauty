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
	"github.com/yang-f/beauty/consts/contenttype"
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
		contenttype.JSON,
	},
	Route{
		"authDemo",
		"GET",
		"/demo1",
		Handler(Config).
			Auth(),
		contenttype.JSON,
	},
	Route{
		"verifyDemo",
		"GET",
		"/demo2",
		Handler(Config).
			Verify(),
		contenttype.JSON,
	},
	Route{
		"verifyAndAuthDemo",
		"GET",
		"/demo3",
		Handler(Config).
			Auth().
			Verify(),
		contenttype.JSON,
	},
}
