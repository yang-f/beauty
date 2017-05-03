package decorates

import (
	"github.com/yang-f/beauty/models"
	"net/http"
)

func ContentType(inner Handler, contentType string) Handler {
	return Handler(func(w http.ResponseWriter, r *http.Request) *models.APPError {
		w.Header().Set("Content-Type", contentType)
		inner.ServeHTTP(w, r)
		return nil
	})
}
