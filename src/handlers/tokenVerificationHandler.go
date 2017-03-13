package handlers

import (
	"forcamp/src/authorization"
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
)

// Parse 'Token' from 'GET' data
func getToken(r *http.Request) string{
	Token := r.FormValue("token")
	return Token
}

func TokenVerificationHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		authorization.VerifyToken(getToken(r), w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleTokenVerification(router *mux.Router)  {
	router.HandleFunc("/token", TokenVerificationHandler)
}
