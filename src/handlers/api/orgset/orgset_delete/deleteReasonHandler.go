package orgset_delete

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strconv"
	"log"
	"forcamp/src/orgset/reasons"
	"strings"
)

func getDeleteReasonPostValues(r *http.Request) (string, int64, *conf.ApiError){
	Token := strings.TrimSpace(r.PostFormValue("token"))
	ID, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("id")), 10, 64)
	if err != nil{
		log.Print(err)
		return "", 0, conf.ErrIDisNotINT
	}

	return Token, ID, nil
}

func DeleteReasonHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		token, id, APIerr := getDeleteReasonPostValues(r)
		if APIerr != nil{
			conf.PrintError(APIerr, w)
		} else {
			reasons.DeleteReason(token, id, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleDeleteReason(router *mux.Router)  {
	router.HandleFunc("/orgset.reason.delete", DeleteReasonHandler)
}


