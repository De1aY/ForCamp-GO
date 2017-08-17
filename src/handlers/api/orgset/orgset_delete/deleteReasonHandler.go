package orgset_delete

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strconv"
	"forcamp/src/api/orgset/reasons"
	"strings"
)

func getDeleteReasonPostValues(r *http.Request) (string, int64, *conf.ApiResponse){
	Token := strings.TrimSpace(r.PostFormValue("token"))
	ID, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("reason_id")), 10, 64)
	if err != nil{
		return "", 0, conf.ErrIdIsNotINT
	}

	return Token, ID, nil
}

func DeleteReasonHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_POST(w)
	if r.Method == http.MethodPost {
		token, id, APIerr := getDeleteReasonPostValues(r)
		if APIerr != nil{
			APIerr.Print(w)
		} else {
			reasons.DeleteReason(token, id, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleDeleteReason(router *mux.Router)  {
	router.HandleFunc("/orgset.reason.delete", DeleteReasonHandler)
}


