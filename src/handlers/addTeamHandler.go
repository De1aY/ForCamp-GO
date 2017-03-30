package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/orgset"
	"strings"
)

func getAddTeamPostValues(r *http.Request) (string, string){
	Token := r.PostFormValue("token")
	Name := strings.ToLower(r.PostFormValue("name"))
	return Name, Token
}

func AddTeamHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		Name, token := getAddTeamPostValues(r)
		orgset.AddTeam(token, Name, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleAddTeam(router *mux.Router)  {
	router.HandleFunc("/orgset.team.add", AddTeamHandler)
}
