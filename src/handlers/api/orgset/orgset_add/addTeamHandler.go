package orgset_add

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strings"
	"forcamp/src/orgset/teams"
)

func getAddTeamPostValues(r *http.Request) (string, string){
	Token := r.PostFormValue("token")
	Name := strings.TrimSpace(strings.ToLower(r.PostFormValue("name")))
	return Name, Token
}

func AddTeamHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		Name, token := getAddTeamPostValues(r)
		teams.AddTeam(token, Name, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleAddTeam(router *mux.Router)  {
	router.HandleFunc("/orgset.team.add", AddTeamHandler)
}
