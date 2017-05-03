package controllers

import (
	"facemark/models"
	"net/http"
)

func Config(w http.ResponseWriter, r *http.Request) *models.APPError {
	b := []byte(`{"description":"this is json"}`)
	w.Write(b)
	return nil
}
