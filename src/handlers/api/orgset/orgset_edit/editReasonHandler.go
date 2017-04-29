package orgset_edit

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strings"
	"strconv"
	"log"
	"forcamp/src/orgset/reasons"
)

func getEditReasonPostValues(r *http.Request) (string, reasons.Reason, *conf.ApiError){
	Token := r.PostFormValue("token")
	ID, err := strconv.ParseInt(r.PostFormValue("id"), 10, 64)
	if err != nil{
		log.Print(err)
		return "", reasons.Reason{}, conf.ErrIDisNotINT
	}
	CatID, err := strconv.ParseInt(r.PostFormValue("cat_id"), 10, 64)
	if err != nil{
		log.Print(err)
		return "", reasons.Reason{}, conf.ErrIDisNotINT
	}
	Text := strings.ToLower(r.PostFormValue("text"))
	Change, err := strconv.ParseInt(r.PostFormValue("change"), 10, 64)
	if err != nil{
		log.Print(err)
		return "", reasons.Reason{}, conf.ErrIDisNotINT
	}
	return Token, reasons.Reason{Id: ID, Cat_id: CatID, Text: Text, Change: int(Change)}, nil
}

func EditReasonHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		token, reason, APIerr := getEditReasonPostValues(r)
		if APIerr != nil{
			conf.PrintError(APIerr, w)
		} else {
			reasons.EditReason(token, reason, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleEditReason(router *mux.Router)  {
	router.HandleFunc("/orgset.reason.edit", EditReasonHandler)
}