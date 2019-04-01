package marks

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"strconv"
	"wplay/conf"
	"wplay/src"
	"wplay/src/api/marks"
)

func getEditMarkPostValues(r *http.Request) (string, int64, int64, int64, *conf.ApiResponse){
	token := strings.TrimSpace(r.PostFormValue("token"))
	category_id, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("category_id")), 10, 64)
	if err != nil {
		return "", 0, 0, 0, conf.ErrCategoryIdNotINT
	}
	reason_id, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("reason_id")), 10, 64)
	if err != nil {
		return "", 0, 0, 0, conf.ErrReasonIncorrect
	}
	participant_id, err := strconv.ParseInt(strings.TrimSpace(
		strings.ToLower(r.PostFormValue("participant_id"))), 10, 64)
	if err != nil {
		return "", 0, 0, 0, conf.ErrReasonIncorrect
	}
	return token, participant_id, category_id, reason_id, nil
}

func editMarkHandler(w http.ResponseWriter, r *http.Request) {
	src.SetHeaders_API_POST(w)
	if r.Method == http.MethodPost {
		token, participant_id, category_id, reason_id, apiErr := getEditMarkPostValues(r)
		if apiErr != nil {
			apiErr.Print(w)
		} else {
			marks.EditMark(token, participant_id, category_id, reason_id, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleEditMark(router *mux.Router) {
	router.HandleFunc("/mark.edit", editMarkHandler)
}
