package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func HandleHTTP(router *mux.Router){
	Server := http.Server{
		Handler: router,
	}
	Server.ListenAndServe()
}
