package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/orgset"
	"strconv"
	"log"
)

func getDeleteTeamPostValues(r *http.Request) (int64, string, *conf.ApiError){
	Token := r.PostFormValue("token")
	ID, err := strconv.ParseInt(r.PostFormValue("id"), 10, 64)
	if err != nil{
		log.Print(err)
		return 0, "", conf.ErrIDisNotINT
	}
	return ID, Token, nil
}

func DeleteTeamHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		ID, token, err := getDeleteTeamPostValues(r)
		if err != nil{
			conf.PrintError(err, w)
		} else {
			orgset.DeleteTeam(token, ID, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleDeleteTeam(router *mux.Router)  {
	router.HandleFunc("/orgset.team.delete", DeleteTeamHandler)
}
