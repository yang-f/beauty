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
	"github.com/yang-f/beauty/db"
	"github.com/yang-f/beauty/models"
	"github.com/yang-f/beauty/utils/log"
	"github.com/yang-f/beauty/utils/token"
	"net/http"
	"strings"
)

func (inner Handler) Auth() Handler {
	return Handler(func(w http.ResponseWriter, r *http.Request) *models.APPError {
		tokenString := ""
		if r.Header != nil {
			if authorization := r.Header["Authorization"]; len(authorization) > 0 {
				tokenString = authorization[0]
			}
		}
		if tokenString == "" {
			cookie, err := r.Cookie("token")
			if err != nil {
				return &models.APPError{err, "token not found.", "AUTH_FAILED", 403}
			}
			tokenString = cookie.Value
		}
		key, err := token.Valid(tokenString)
		if err != nil {
			return &models.APPError{err, "bad token.", "AUTH_FAILED", 403}
		}
		if !strings.Contains(key, "|") {
			return &models.APPError{err, "user not found.", "NOT_FOUND", 404}
		}
		keys := strings.Split(key, "|")
		rows, _, err := db.QueryNonLogging("select * from user where user_id = '%v' and user_pass = '%v'", keys[0], keys[1])
		if err != nil {
			return &models.APPError{err, "can not connect database.", "DB_ERROR", 500}
		}
		if len(rows) == 0 {
			return &models.APPError{err, "user not found.", "NOT_FOUND", 404}
		}
		go log.Printf("user_id:%v", keys[0])
		inner.ServeHTTP(w, r)
		return nil
	})
}
