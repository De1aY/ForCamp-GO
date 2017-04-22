package handlers

import (
	"net/http"
	"forcamp/src"
	"forcamp/conf"
	"github.com/gorilla/mux"
)

func ExitHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_Main(w)
	if r.TLS != nil {
		if r.Method == http.MethodGet{
			w.Header().Set("Set-Cookie", "token=deleted; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT")
			http.Redirect(w, r, "https://forcamp.ga", http.StatusTemporaryRedirect)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			conf.PrintError(conf.ErrMethodNotAllowed, w)
		}
	} else {
		http.Redirect(w, r, "https://forcamp.ga/exit", http.StatusTemporaryRedirect)
	}
}

func HandleExit(router *mux.Router){
	router.HandleFunc("/exit", ExitHandler)
}