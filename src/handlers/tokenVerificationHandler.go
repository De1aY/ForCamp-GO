package handlers

import (
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/api/authorization"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

func GetToken(r *http.Request) string {
	Token, _ := url.QueryUnescape(strings.TrimSpace(strings.ToLower(r.FormValue("token"))))
	return Token
}

func TokenVerificationHandler(w http.ResponseWriter, r *http.Request) {
	src.SetHeaders_API_GET(w)
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		authorization.VerifyToken(GetToken(r), w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleTokenVerification(router *mux.Router) {
	router.HandleFunc("/token.verify", TokenVerificationHandler)
}
