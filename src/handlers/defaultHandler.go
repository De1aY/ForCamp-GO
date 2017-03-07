package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request){
	TargetURL := "https://forcamp.ga/index.html"
	http.Redirect(w, r , TargetURL, http.StatusTemporaryRedirect)
}

func HandleDefault(router *mux.Router){
	router.HandleFunc("/", defaultHandler)
}