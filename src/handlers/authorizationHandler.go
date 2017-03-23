package handlers

import (
	"forcamp/src/authorization"
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
)

// Parse 'GET' data to AuthInf
func getAuthorizationData(r *http.Request) authorization.AuthInf{
	UserLogin := r.FormValue("login")
	UserPassword := r.FormValue("password")
	authInf := authorization.AuthInf{}
	authInf.Login = UserLogin
	authInf.Password = UserPassword
	return authInf
}

func LoginAndPasswordAuthHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		authInf := getAuthorizationData(r)
		authorization.Authorize(authInf, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleAuthorizationByLoginAndPassword(router *mux.Router)  {
	router.HandleFunc("/token.get", LoginAndPasswordAuthHandler)
}
