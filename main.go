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
	http.ListenAndServe(":80", Router)
	http.ListenAndServeTLS(":443", "./conf/tls/apiforcamp.pem", "./conf/tls/apiforcamp_key.pem", APISite)
	http.ListenAndServeTLS(":443", "./conf/tls/forcamp.pem", "./conf/tls/forcamp_key.pem", MainSite)
}
