/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"nullteam.info/wplay/demo/conf"
	"nullteam.info/wplay/demo/src"
)

func ExitHandler(w http.ResponseWriter, r *http.Request) {
	src.SetHeaders_Main(w)
	if r.TLS != nil {
		if r.Method == http.MethodGet {
			w.Header().Set("Set-Cookie", "token=deleted; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT")
			http.Redirect(w, r, "https://"+conf.MAIN_SITE_DOMAIN, http.StatusTemporaryRedirect)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			conf.ErrMethodNotAllowed.Print(w)
		}
	} else {
		http.Redirect(w, r, "https://"+conf.MAIN_SITE_DOMAIN+"/exit", http.StatusTemporaryRedirect)
	}
}

func HandleExit(router *mux.Router) {
	router.HandleFunc("/exit", ExitHandler)
}
