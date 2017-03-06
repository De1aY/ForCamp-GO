package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"crypto/tls"
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
	TlsConfig := &tls.Config{}
	TlsConfig.Certificates = make([]tls.Certificate, 3)
	TlsConfig.Certificates[0], _ = tls.LoadX509KeyPair("./conf/tls/apiforcamp.pem", "./conf/tls/apiforcamp_key.pem")
	TlsConfig.Certificates[1], _ = tls.LoadX509KeyPair("./conf/tls/forcamp.pem", "./conf/tls/forcamp_key.pem")
	TlsConfig.Certificates[2], _ = tls.LoadX509KeyPair("./conf/tls/wwwforcamp.pem", "./conf/tls/wwwforcamp_key.pem")
	TlsConfig.BuildNameToCertificate()
	Server := http.Server{
		Handler: Router,
		TLSConfig: TlsConfig,
	}
	TLSListener, _ := tls.Listen("tcp", ":443", TlsConfig)
	http.ListenAndServe(":80", Router)
	Server.Serve(TLSListener)
}
