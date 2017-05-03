package decorates

import (
	"github.com/yang-f/beauty/utils"
	"github.com/yang-f/beauty/utils/log"
	"github.com/yang-f/beauty/utils/token"
	"net/http"
)

func Auth(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil || cookie.Value == "" {
			utils.Response(w, "token not found.", "AUTH_FAILED", 403)
			return
		}
		user_id, err := token.Valid(cookie.Value)
		if err != nil {
			utils.Response(w, "bad token.", "AUTH_FAILED", 403)
			return
		}
		log.Printf("user_id:%v", user_id)
		inner.ServeHTTP(w, r)
	})
}
