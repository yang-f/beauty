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

package db

import (
	"github.com/yang-f/beauty/settings"
	"github.com/yang-f/beauty/utils/log"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/thrsafe"
)

func Query(sql string, params ...interface{}) ([]mysql.Row, mysql.Result, error) {
	raws, res, err := query(sql, true, params...)
	return raws, res, err
}

func QueryNonLogging(sql string, params ...interface{}) ([]mysql.Row, mysql.Result, error) {
	raws, res, err := query(sql, false, params...)
	return raws, res, err
}

func query(sql string, logging bool, params ...interface{}) ([]mysql.Row, mysql.Result, error) {
	db := mysql.New("tcp", "", settings.Local["mysql_host"], settings.Local["mysql_user"], settings.Local["mysql_pass"], settings.Local["mysql_database"])
	if err := db.Connect(); err != nil {
		log.Println(err)
		return nil, nil, err
	}
	defer db.Close()
	if logging {
		go log.Printf("query:"+sql, params...)
	}
	raws, res, err := db.Query(sql, params...)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	return raws, res, err
}
