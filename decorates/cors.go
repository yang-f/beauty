package decorates

import (
	"github.com/yang-f/beauty/models"
	"net/http"
)

func CorsHeader(inner Handler) Handler {
	return Handler(func(w http.ResponseWriter, r *http.Request) *models.APPError {
		origin := r.Header.Get("Origin")
		if origin == "" {
			origin = "http://ar.huluuu.com"
		}
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Method", "POST, OPTIONS, GET, HEAD, PUT, PATCH, DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-HTTP-Method-Override,accept-charset,accept-encoding , Content-Type, Accept, Cookie")
		inner.ServeHTTP(w, r)
		return nil
	})
}

func CorsHeader2(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			origin = "http://ar.huluuu.com"
		}
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Method", "POST, OPTIONS, GET, HEAD, PUT, PATCH, DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-HTTP-Method-Override,accept-charset,accept-encoding , Content-Type, Accept, Cookie")
		inner.ServeHTTP(w, r)
	})
}
