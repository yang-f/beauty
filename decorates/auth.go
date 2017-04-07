package decorates

import (
	"net/http"
	"github.com/yang-f/beauty/controllers"
	"github.com/yang-f/beauty/db"
	"github.com/yang-f/beauty/utils/token"
	"github.com/yang-f/beauty/utils"
	"github.com/yang-f/beauty/utils/log"
)

func Auth(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil || cookie.Value == ""{
			utils.Response(w, "token not found.", "AUTH_FAILED",403)
			return;
		}
		user_id, err := token.Valid(cookie.Value)
		if err != nil {
			utils.Response(w, "bad token.", "AUTH_FAILED",403)
			return;
		}
		if(controllers.CurrentUser.User_id != user_id){
			rows, res, err := db.Query("select * from user where user_id= '%d'", user_id)

			if err != nil {
				utils.Response(w, "can not connect database.", "DB_ERROR",500)
				return
			}

			if len(rows) == 0 {
				utils.Response(w, "user not found.", "NOT_FOUND",404)
				return
			}
			row := rows[0]
			user := controllers.User{
				User_id:row.Int(res.Map("user_id")), 
				User_name:row.Str(res.Map("user_name")),
				User_type:row.Str(res.Map("user_type")),
				Add_time:row.Str(res.Map("add_time"))}

			controllers.CurrentUser = user
		}
		log.Printf("user_id:%v",controllers.CurrentUser.User_id)
		inner.ServeHTTP(w, r)
	})
}
