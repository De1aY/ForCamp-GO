package handlers

import (
	"forcamp/src/users"
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
)


func GetUserLoginHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodGet {
		users.GetUserLogin(getToken(r), w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleGetUserLoginByToken(router *mux.Router)  {
	router.HandleFunc("/user.login.get", GetUserLoginHandler)
}
