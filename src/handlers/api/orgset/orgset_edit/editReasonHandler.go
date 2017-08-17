package orgset_edit

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strings"
	"strconv"
	"forcamp/src/api/orgset/reasons"
)

func getEditReasonPostValues(r *http.Request) (string, reasons.Reason, *conf.ApiResponse){
	Token := strings.TrimSpace(r.PostFormValue("token"))
	ID, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("reason_id")), 10, 64)
	if err != nil{
		return "", reasons.Reason{}, conf.ErrIdIsNotINT
	}
	CatID, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("category_id")), 10, 64)
	if err != nil{
		return "", reasons.Reason{}, conf.ErrIdIsNotINT
	}
	Text := strings.TrimSpace(r.PostFormValue("text"))
	Change, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("change")), 10, 64)
	if err != nil{
		return "", reasons.Reason{}, conf.ErrIdIsNotINT
	}
	return Token, reasons.Reason{Id: ID, Cat_id: CatID, Text: Text, Change: Change}, nil
}

func EditReasonHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_POST(w)
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
