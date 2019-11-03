/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package orgset_edit

import (
	"net/http"
	"github.com/gorilla/mux"
	"nullteam.info/wplay/demo/conf"
	"nullteam.info/wplay/demo/src"
	"strings"
	"strconv"
	"nullteam.info/wplay/demo/src/api/orgset/teams"
)

func getEditTeamPostValues(r *http.Request) (string, string, int64, *conf.ApiResponse){
	Token := strings.TrimSpace(r.PostFormValue("token"))
	ID, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("team_id")), 10, 64)
	if err != nil{
		return "", "", 0, conf.ErrIdIsNotINT
	}
	Name := strings.TrimSpace(strings.ToLower(r.PostFormValue("name")))
	return Token, Name, ID, nil
}

func EditTeamHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_POST(w)
	if r.Method == http.MethodPost {
		token, name, id, APIerr := getEditTeamPostValues(r)
		if APIerr != nil{
			APIerr.Print(w)
		} else {
			teams.EditTeam(token, name, id, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleEditTeam(router *mux.Router)  {
	router.HandleFunc("/orgset.team.edit", EditTeamHandler)
}
