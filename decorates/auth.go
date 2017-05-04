package decorates

import (
	"github.com/yang-f/beauty/db"
	"github.com/yang-f/beauty/models"
	"github.com/yang-f/beauty/utils/log"
	"github.com/yang-f/beauty/utils/token"
	"net/http"
	"strings"
)

func Auth(inner Handler) Handler {
	return Handler(func(w http.ResponseWriter, r *http.Request) *models.APPError {
		cookie, err := r.Cookie("token")
		if err != nil || cookie.Value == "" {
			return &models.APPError{err, "token not found.", "AUTH_FAILED", 403}
		}
		key, err := token.Valid(cookie.Value)
		if err != nil {
			return &models.APPError{err, "bad token.", "AUTH_FAILED", 403}
		}
		if !strings.Contains(key, "|") {
			return &models.APPError{err, "user not found.", "NOT_FOUND", 404}
		}
		keys := strings.Split(key, "|")
		rows, _, err := db.QueryNonLogging("select * from user where user_id = '%v' and user_pass = '%v'", keys[0], keys[1])
		if err != nil {
			return &models.APPError{err, "can not connect database.", "DB_ERROR", 500}
		}
		if len(rows) == 0 {
			return &models.APPError{err, "user not found.", "NOT_FOUND", 404}
		}
		go log.Printf("user_id:%v", keys[0])
		inner.ServeHTTP(w, r)
		return nil
	})
}
