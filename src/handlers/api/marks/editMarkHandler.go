package marks

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"strconv"
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/api/marks"
)

func getEditMarkPostValues(r *http.Request) (string, string, int64, int64, *conf.ApiResponse){
	token := strings.TrimSpace(r.PostFormValue("token"))
	category_id, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("category_id")), 10, 64)
	if err != nil {
		return "", "", 0, 0, conf.ErrCategoryIdNotINT
	}
	reason_id, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("reason_id")), 10, 64)
	if err != nil {
		return "", "", 0, 0, conf.ErrReasonIncorrect
	}
	participant_login := strings.TrimSpace(strings.ToLower(r.PostFormValue("login")))
	return token, participant_login, category_id, reason_id, nil
}

func editMarkHandler(w http.ResponseWriter, r *http.Request) {
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		token, participant_login, category_id, reason_id, APIerr := getEditMarkPostValues(r)
		if APIerr != nil {
			APIerr.Print(w)
		} else {
			marks.EditMark(token, participant_login, category_id, reason_id, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleEditMark(router *mux.Router) {
	router.HandleFunc("/mark.edit", editMarkHandler)
}