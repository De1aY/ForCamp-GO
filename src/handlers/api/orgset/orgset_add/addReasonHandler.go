package orgset_add

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strings"
	"strconv"
	"log"
	"forcamp/src/api/orgset/reasons"
)

func getAddReasonPostValues(r *http.Request) (string, reasons.Reason, *conf.ApiError){
	Token := r.PostFormValue("token")
	CatID, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("cat_id")), 10, 64)
	if err != nil{
		log.Print(err)
		return "", reasons.Reason{}, conf.ErrIDisNotINT
	}
	Text := strings.TrimSpace(strings.ToLower(r.PostFormValue("text")))
	Change, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("change")), 10, 64)
	if err != nil{
		log.Print(err)
		return "", reasons.Reason{}, conf.ErrIDisNotINT
	}
	return Token, reasons.Reason{Id: 0, Cat_id: CatID, Text: Text, Change: int(Change)}, nil
}

func AddReasonHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		token, reason, APIerr := getAddReasonPostValues(r)
		if APIerr != nil{
			conf.PrintError(APIerr, w)
		} else {
			reasons.AddReason(token, reason, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleAddReason(router *mux.Router)  {
	router.HandleFunc("/orgset.reason.add", AddReasonHandler)
}

