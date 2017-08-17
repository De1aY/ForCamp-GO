package users_get

import (
	"forcamp/src/api/users"
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/handlers"
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
