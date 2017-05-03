package decorates

import (
	"net/http"
)

func ContentType(inner http.Handler, contentType string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		inner.ServeHTTP(w, r)
	})
}
