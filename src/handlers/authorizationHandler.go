/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package handlers

import (
	"wplay/src/api/authorization"
	"net/http"
	"github.com/gorilla/mux"
	"wplay/conf"
	"wplay/src"
	"strings"
)

// Parse 'GET' data to AuthInf
func getAuthorizationData(r *http.Request) authorization.AuthInf{
	UserLogin := strings.TrimSpace(r.FormValue("login"))
	UserPassword := strings.TrimSpace(r.FormValue("password"))
	authInf := authorization.AuthInf{}
	authInf.Login = UserLogin
	authInf.Password = UserPassword
	return authInf
}

func LoginAndPasswordAuthHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_GET(w)
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		authInf := getAuthorizationData(r)
		authorization.Authorize(authInf, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleAuthorizationByLoginAndPassword(router *mux.Router)  {
	router.HandleFunc("/token.get", LoginAndPasswordAuthHandler)
}
