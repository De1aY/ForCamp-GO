package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func defaultHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello!")
}

func main(){
	Router := mux.NewRouter()
	MainSite := Router.Host("forcamp.ga").Subrouter()
	APISite := Router.Host("api.forcamp.ga").Subrouter()
	APISite.HandleFunc("/", defaultHandler)
	MainSite.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":80", Router)
}
