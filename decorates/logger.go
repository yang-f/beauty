package decorates

import (
	"github.com/yang-f/beauty/models"
	"github.com/yang-f/beauty/utils/log"
	"net/http"
	"time"
)

func Logger(inner Handler, name string) Handler {
	return Handler(func(w http.ResponseWriter, r *http.Request) *models.APPError {
		start := time.Now()

		inner.ServeHTTP(w, r)

		go log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
		return nil
	})
}
