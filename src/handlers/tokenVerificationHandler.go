package handlers

import (
	"forcamp/src/api/authorization"
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strings"
)

// Parse 'Token' from 'GET' data
func GetToken(r *http.Request) string{
	Token := strings.TrimSpace(strings.ToLower(r.FormValue("token")))
	return Token
}

func TokenVerificationHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		authorization.VerifyToken(GetToken(r), w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleTokenVerification(router *mux.Router)  {
	router.HandleFunc("/token.verify", TokenVerificationHandler)
}
