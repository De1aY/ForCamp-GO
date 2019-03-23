package orgset_add

import (
	"net/http"
	"github.com/gorilla/mux"
	"wplay/conf"
	"wplay/src"
	"strings"
	"strconv"
	"wplay/src/api/orgset/reasons"
)

func getAddReasonPostValues(r *http.Request) (string, reasons.Reason, *conf.ApiResponse){
	Token := r.PostFormValue("token")
	CatID, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("category_id")), 10, 64)
	if err != nil{
		return "", reasons.Reason{}, conf.ErrIdIsNotINT
	}
	Text := strings.TrimSpace(r.PostFormValue("text"))
	Change, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("change")), 10, 64)
	if err != nil{
		return "", reasons.Reason{}, conf.ErrIdIsNotINT
	}
	return Token, reasons.Reason{Id: 0, Cat_id: CatID, Text: Text, Change: Change}, nil
}

func AddReasonHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_POST(w)
	if r.Method == http.MethodPost {
		token, reason, apiErr := getAddReasonPostValues(r)
		if apiErr != nil{
			apiErr.Print(w)
		} else {
			reasons.AddReason(token, reason, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleAddReason(router *mux.Router)  {
	router.HandleFunc("/orgset.reason.add", AddReasonHandler)
}

