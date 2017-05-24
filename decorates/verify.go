package decorates

import (
	"bytes"
	"facemark/models"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"regexp"
)

func (inner Handler) Verify() Handler {
	return Handler(func(w http.ResponseWriter, r *http.Request) *models.APPError {
		sqlInjection, err := regexp.Compile(`(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`)
		if err != nil {
			return &models.APPError{err, "insecure params.", "INSECURE_PARAMES", 403}
		}
		result := []byte{}

		if r.Body != nil {
			result, err = ioutil.ReadAll(r.Body)
			r.ParseForm()
			r.Body = ioutil.NopCloser(bytes.NewBuffer(result))
		}
		if err != nil {
			return &models.APPError{err, "insecure params.", "INSECURE_PARAMES", 403}
		}
		if sqlInjection.MatchString(string(result)) {
			return &models.APPError{err, "insecure params.", "INSECURE_PARAMES", 403}
		}
		if sqlInjection.MatchString(fmt.Sprintf("%v", mux.Vars(r))) {
			return &models.APPError{err, "insecure params.", "INSECURE_PARAMES", 403}
		}
		if sqlInjection.MatchString(fmt.Sprintf("%v", r.Form)) {
			return &models.APPError{err, "insecure params.", "INSECURE_PARAMES", 403}
		}

		inner.ServeHTTP(w, r)
		return nil
	})
}
