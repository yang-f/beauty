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
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/yang-f/beauty/models"
	"io/ioutil"
	"net/http"
	"regexp"
)

func (inner Handler) Verify() Handler {
	return Handler(func(w http.ResponseWriter, r *http.Request) *models.APPError {
		sqlInjection, err := regexp.Compile(`(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`)
		if err != nil {
			return &models.APPError{err, "insecure params.", "INSECURE_PARAMES", 403}
		}
		result := []byte{}

		if r.Body != nil {
			result, err = ioutil.ReadAll(r.Body)
			r.ParseForm()
			r.Body = ioutil.NopCloser(bytes.NewBuffer(result))
		}
		if err != nil {
			return &models.APPError{err, "insecure params.", "INSECURE_PARAMES", 403}
		}
		if sqlInjection.MatchString(string(result)) {
			return &models.APPError{err, "insecure params.", "INSECURE_PARAMES", 403}
		}
		if sqlInjection.MatchString(fmt.Sprintf("%v", mux.Vars(r))) {
			return &models.APPError{err, "insecure params.", "INSECURE_PARAMES", 403}
		}
		if sqlInjection.MatchString(fmt.Sprintf("%v", r.Form)) {
			return &models.APPError{err, "insecure params.", "INSECURE_PARAMES", 403}
		}

		inner.ServeHTTP(w, r)
		return nil
	})
}
