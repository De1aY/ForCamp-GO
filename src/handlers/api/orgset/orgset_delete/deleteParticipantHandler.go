package orgset_delete

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strings"
	"forcamp/src/orgset/participants"
)

func getDeleteParticipantPostValues(r *http.Request) (string, string){
	Token := r.PostFormValue("token")
	Login := strings.ToLower(r.PostFormValue("login"))
	return Login, Token
}

func DeleteParticipantHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		login, token := getDeleteParticipantPostValues(r)
		participants.DeleteParticipant(token, login, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleDeleteParticipant(router *mux.Router)  {
	router.HandleFunc("/orgset.participant.delete", DeleteParticipantHandler)
}
