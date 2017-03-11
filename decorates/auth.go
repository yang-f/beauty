package decorates

import (
	"net/http"
	"github.com/yang-f/beauty/utils/token"
    "github.com/yang-f/beauty/utils/log"
    "github.com/yang-f/beauty/utils"
)

func Auth(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	    cookie, err := r.Cookie("token")
	    if err != nil || cookie.Value == ""{
	        utils.Response(w, "token not found.", "AUTH_FAILED",403)
	        return;
	    }
	    user_id, err := token.Valid(cookie.Value)
	    log.Printf("user_id:%v",user_id)
		inner.ServeHTTP(w, r)
	})
}
