package orgset_get

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/api/orgset/teams"
	"forcamp/src/handlers"
)

func GetTeamsHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_GET(w)
	if r.Method == http.MethodGet {
		teams.GetTeams(handlers.GetToken(r), w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleGetTeams(router *mux.Router)  {
	router.HandleFunc("/orgset.teams.get", GetTeamsHandler)
}
