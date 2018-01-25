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

package sessions

import (
	"net/http"
	"strings"

	"github.com/yang-f/beauty/db"
	"github.com/yang-f/beauty/models"
	"github.com/yang-f/beauty/utils/token"
)

// CurrentUser for session
func CurrentUser(r *http.Request) *models.User {
	tokenString := ""
	cookie, _ := r.Cookie("token")
	if cookie != nil {
		tokenString = cookie.Value
	}
	if tokenString == "" {
		if r.Header != nil {
			if authorization := r.Header["Authorization"]; len(authorization) > 0 {
				tokenString = authorization[0]
			}
		}
	}
	key, err := token.Valid(tokenString)
	if err != nil {
		return &models.User{}
	}
	if !strings.Contains(key, "|") {
		return &models.User{}
	}
	keys := strings.Split(key, "|")
	rows, res, err := db.QueryNonLogging("select * from user where user_id = '%v' and user_pass = '%v'", keys[0], keys[1])

	if err != nil {
		return &models.User{}
	}

	if len(rows) == 0 {
		return &models.User{}
	}
	row := rows[0]
	user := models.User{
		User_id:   row.Int(res.Map("user_id")),
		User_name: row.Str(res.Map("user_name")),
		User_type: row.Str(res.Map("user_type")),
		Add_time:  row.Str(res.Map("add_time"))}

	return &user
}
