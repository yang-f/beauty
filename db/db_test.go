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

package db_test

import (
	"errors"
	"testing"

	"github.com/yang-f/beauty/db"
	"github.com/yang-f/beauty/settings"
)

func init() {
	db.Query("drop database beauty_test")
	db.Query("create database beauty_test")
	settings.Local["mysql_database"] = "beauty_test"
	db.Query(`create table if not exists user
        (
            user_id int primary key not null  auto_increment,
            user_name varchar(64),
            user_pass varchar(64),
            user_mobile varchar(32),
            user_type enum('user', 'admin', 'test') not null,
            add_time timestamp not null default CURRENT_TIMESTAMP
        )`)
	db.Query(`insert into user (user_name, user_pass, user_mobile, add_time) values('test', 'test', '12345678901', '2017-01-12 12:15:38')`)
}

type tester struct {
	sql string
	err error
}

func TestQuery(t *testing.T) {
	var tests = []tester{
		{"select * from whatever", errors.New(`Received #1146 error from MySQL server: "Table 'beauty_test.whatever' doesn't exist"`)},
		{"select * from user", nil},
	}

	for _, test := range tests {
		_, _, err := db.Query(test.sql)
		if err != nil {
			if err.Error() != test.err.Error() {
				t.Errorf("returned wrong err code: got %v want %v",
					err.Error(), test.err)
			}
			continue
		}
		if err != test.err {
			t.Errorf("returned wrong err code: got %v want %v",
				err.Error(), test.err)
		}
	}

}
