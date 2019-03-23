package handlers

import (
	"wplay/src/api/authorization"
	"net/http"
	"github.com/gorilla/mux"
	"wplay/conf"
	"wplay/src"
	"strings"
)

func getAuthorizationData(r *http.Request) authorization.AuthInf{
	UserLogin := strings.TrimSpace(r.PostFormValue("login"))
	UserPassword := strings.TrimSpace(r.PostFormValue("password"))
	authInf := authorization.AuthInf{}
	authInf.Login = UserLogin
	authInf.Password = UserPassword
	return authInf
}

func LoginAndPasswordAuthHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_POST(w)
	if r.Method == http.MethodPost {
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
