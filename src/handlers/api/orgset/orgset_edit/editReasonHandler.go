package orgset_edit

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

func getEditReasonPostValues(r *http.Request) (string, reasons.Reason, *conf.ApiResponse){
	Token := strings.TrimSpace(r.PostFormValue("token"))
	ID, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("id")), 10, 64)
	if err != nil{
		log.Print(err)
		return "", reasons.Reason{}, conf.ErrIDisNotINT
	}
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
	return Token, reasons.Reason{Id: ID, Cat_id: CatID, Text: Text, Change: int(Change)}, nil
}

func EditReasonHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		token, reason, APIerr := getEditReasonPostValues(r)
		if APIerr != nil{
			APIerr.Print(w)
		} else {
			reasons.EditReason(token, reason, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleEditReason(router *mux.Router)  {
	router.HandleFunc("/orgset.reason.edit", EditReasonHandler)
}
