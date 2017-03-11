package decorates

import (
	"github.com/yang-f/beauty/utils/log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf(
			"start:\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
		)
		inner.ServeHTTP(w, r)

		log.Printf(
			"end:\t%s",
			time.Since(start),
		)
	})
}
