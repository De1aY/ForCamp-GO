package orgset_edit

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strings"
	"strconv"
	"log"
	"forcamp/src/orgset/teams"
)

func getEditTeamPostValues(r *http.Request) (string, string, int64, *conf.ApiError){
	Token := r.PostFormValue("token")
	ID, err := strconv.ParseInt(r.PostFormValue("id"), 10, 64)
	if err != nil{
		log.Print(err)
		return "", "", 0, conf.ErrIDisNotINT
	}
	Name := strings.ToLower(r.PostFormValue("name"))
	return Token, Name, ID, nil
}

func EditTeamHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		token, name, id, APIerr := getEditTeamPostValues(r)
		if APIerr != nil{
			conf.PrintError(APIerr, w)
		} else {
			teams.EditTeam(token, name, id, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleEditTeam(router *mux.Router)  {
	router.HandleFunc("/orgset.team.edit", EditTeamHandler)
}