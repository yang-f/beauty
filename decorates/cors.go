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

package decorates

import (
	"net/http"

	"github.com/yang-f/beauty/models"
	"github.com/yang-f/beauty/router"
)

func (inner Handler) CorsHeader() Handler {
	return Handler(func(c *router.Context) *models.APPError {
		c.W.Header().Set("Access-Control-Allow-Origin", "*")
		c.W.Header().Set("Access-Control-Allow-Credentials", "true")
		c.W.Header().Add("Access-Control-Allow-Method", "POST, OPTIONS, GET, HEAD, PUT, PATCH, DELETE")
		c.W.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-HTTP-Method-Override,accept-charset,accept-encoding , Content-Type, Accept, Cookie")
		inner.ServeHTTP(c.W, c.R)
		return nil
	})
}

func CorsHeader2(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Method", "POST, OPTIONS, GET, HEAD, PUT, PATCH, DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-HTTP-Method-Override,accept-charset,accept-encoding , Content-Type, Accept, Cookie")
		inner.ServeHTTP(w, r)
	})
}
