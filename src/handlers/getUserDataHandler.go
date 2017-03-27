package handlers

import (
	"forcamp/src/users"
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
)

func getLogin(r *http.Request) string{
	Login := r.FormValue("login")
	return Login
}

func GetUserDataHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodGet {
		users.GetUserData(getToken(r), w, getLogin(r))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleGetUserData(router *mux.Router)  {
	router.HandleFunc("/user.data.get", GetUserDataHandler)
}
