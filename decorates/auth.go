package decorates

import (
	"github.com/yang-f/beauty/db"
	"github.com/yang-f/beauty/utils"
	"github.com/yang-f/beauty/utils/log"
	"github.com/yang-f/beauty/utils/token"
	"net/http"
	"strings"
)

func Auth(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil || cookie.Value == "" {
			utils.Response(w, "token not found.", "AUTH_FAILED", 403)
			return
		}
		key, err := token.Valid(cookie.Value)
		if err != nil {
			utils.Response(w, "bad token.", "AUTH_FAILED", 403)
			return
		}
		if !strings.Contains(key, "|") {
			utils.Response(w, "user not found.", "NOT_FOUND", 404)
			return
		}
		keys := strings.Split(key, "|")
		rows, _, err := db.QueryNonLogging("select * from user where user_id = '%v' and user_pass = '%v'", keys[0], keys[1])
		if err != nil {
			utils.Response(w, "can not connect database.", "DB_ERROR", 500)
			return
		}
		if len(rows) == 0 {
			utils.Response(w, "user not found.", "NOT_FOUND", 404)
			return
		}
		log.Printf("user_id:%v", keys[0])
		inner.ServeHTTP(w, r)
	})
}
