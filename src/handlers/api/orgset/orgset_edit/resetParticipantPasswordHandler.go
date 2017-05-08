package orgset_edit

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strings"
	"forcamp/src/api/orgset/participants"
)

func getResetParticipantPasswordPostValues(r *http.Request) (string, string){
	Token := strings.TrimSpace(r.PostFormValue("token"))
	Login := strings.TrimSpace(strings.ToLower(r.PostFormValue("login")))
	return Login, Token
}

func ResetParticipantPasswordHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		login, token := getResetParticipantPasswordPostValues(r)
		participants.ResetParticipantPassword(token, login, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleResetParticipantPassword(router *mux.Router)  {
	router.HandleFunc("/orgset.participant.password.reset", ResetParticipantPasswordHandler)
}
