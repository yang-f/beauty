package decorates

import (
	"github.com/yang-f/beauty/models"
	"github.com/yang-f/beauty/utils"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request) *models.APPError

func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil {
		utils.Response(w, e.Message, e.Code, e.Status)
	}
}
