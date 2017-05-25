package controllers_test

import (
	"github.com/yang-f/beauty/db"
	"github.com/yang-f/beauty/router"
	"github.com/yang-f/beauty/settings"
	"net/http"
	"net/http/httptest"
	"testing"
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
	Pattern  string
	Status   int
	Expected string
	Method   string
}

func TestConfig(t *testing.T) {
	var tests = []tester{
		{"/", 200, `{"description":"this is json"}`, "GET"},
		{"/demo1", 403, `{"status":403,"description":"token not found.","code":"AUTH_FAILED"}`, "GET"},
		{"/demo2", 200, `{"description":"this is json"}`, "GET"},
		{"/demo3", 403, `{"status":403,"description":"token not found.","code":"AUTH_FAILED"}`, "GET"},
	}

	for _, test := range tests {
		req, err := http.NewRequest(test.Method, test.Pattern, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		router := router.NewRouter()
		router.ServeHTTP(rr, req)
		if status := rr.Code; status != test.Status {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		if rr.Body.String() != test.Expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), test.Expected)
		}
	}
}

func Benchmark_config(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			b.Fatal(err)
		}
		rr := httptest.NewRecorder()

		router := router.NewRouter()
		router.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			b.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		expected := `{"description":"this is json"}`
		if rr.Body.String() != expected {
			b.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
}
