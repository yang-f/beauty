package controllers

import (
	"net/http"
)

func Config(w http.ResponseWriter, r *http.Request) {
	b := []byte(`{"description":"this is json"}`)
	w.Write(b)
}