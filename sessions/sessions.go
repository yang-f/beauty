package sessions

import (
	"github.com/yang-f/beauty/db"
	"github.com/yang-f/beauty/models"
	"github.com/yang-f/beauty/utils/token"
	"net/http"
	"strings"
)

func CurrentUser(r *http.Request) *models.User {
	cookie, err := r.Cookie("token")
	if err != nil || cookie.Value == "" {
		return models.User{}
	}
	key, err := token.Valid(cookie.Value)
	if err != nil {
		return models.User{}
	}
	if !strings.Contains(key, "|") {
		return models.User{}
	}
	keys := strings.Split(key, "|")
	rows, res, err := db.QueryNonLogging("select * from user where user_id = '%v' and user_pass = '%v'", keys[0], keys[1])

	if err != nil {
		return models.User{}
	}

	if len(rows) == 0 {
		return models.User{}
	}
	row := rows[0]
	user := models.User{
		User_id:   row.Int(res.Map("user_id")),
		User_name: row.Str(res.Map("user_name")),
		User_type: row.Str(res.Map("user_type")),
		Add_time:  row.Str(res.Map("add_time"))}

	return &user
}
