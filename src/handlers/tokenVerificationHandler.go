/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package handlers

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
	"nullteam.info/wplay/demo/conf"
	"nullteam.info/wplay/demo/src"
	"nullteam.info/wplay/demo/src/api/authorization"
)

// Parse 'Token' from 'GET' data
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
