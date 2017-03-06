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
	WWWSite := Router.Host("www.forcamp.ga").Subrouter()
	MainSite := Router.Host("forcamp.ga").Subrouter()
	APISite := Router.Host("api.forcamp.ga").Subrouter()
	MainSite.HandleFunc("/", defaultHandler)
	APISite.HandleFunc("/", defaultHandler)
	WWWSite.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":80", Router)
	http.ListenAndServeTLS(":443", "./conf/tls/apiforcamp.pem", "./conf/tls/apiforcamp_key.pem", APISite)
	http.ListenAndServeTLS(":443", "./conf/tls/forcamp.pem", "./conf/tls/forcamp_key.pem", MainSite)
	http.ListenAndServeTLS(":443", "./conf/tls/wwwforcamp.pem", "./conf/tls/wwwforcamp_key.pem", WWWSite)
}
