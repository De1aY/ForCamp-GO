/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package users_get

import (
	"nullteam.info/wplay/demo/src/api/users"
	"net/http"
	"github.com/gorilla/mux"
	"nullteam.info/wplay/demo/conf"
	"nullteam.info/wplay/demo/src"
	"nullteam.info/wplay/demo/src/handlers"
)


func GetUserIdHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_GET(w)
	if r.Method == http.MethodGet {
		users.GetUserID(handlers.GetToken(r), w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleGetUserLoginByToken(router *mux.Router)  {
	router.HandleFunc("/user.id.get", GetUserIdHandler)
}
