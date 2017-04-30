package users_get

import (
	"forcamp/src/users"
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/handlers"
	"strings"
)

func getLogin(r *http.Request) string{
	Login := strings.TrimSpace(r.FormValue("login"))
	return Login
}

func GetUserDataHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodGet {
		users.GetUserData(handlers.GetToken(r), w, getLogin(r))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleGetUserData(router *mux.Router)  {
	router.HandleFunc("/user.data.get", GetUserDataHandler)
}
