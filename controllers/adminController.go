package controllers

import (
	"github.com/yang-f/beauty/models"
	"net/http"
)

func Config(w http.ResponseWriter, r *http.Request) *models.APPError {
	b := []byte(`{"description":"this is json"}`)
	w.Write(b)
	return nil
}
