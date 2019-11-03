/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package handlers

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"nullteam.info/wplay/demo/conf"
)

// Read TLS certificates to TlsConfig
func getTlsConfig() *tls.Config {
	TlsConfig := &tls.Config{}
	TlsConfig.Certificates = make([]tls.Certificate, 2)
	TlsConfig.Certificates[0], _ = tls.LoadX509KeyPair(conf.API_SITE_TLS, conf.API_SITE_TLS_KEY)
	TlsConfig.Certificates[1], _ = tls.LoadX509KeyPair(conf.MAIN_SITE_TLS, conf.MAIN_SITE_TLS_KEY)
	// TlsConfig.Certificates[2], _ = tls.LoadX509KeyPair(conf.WWW_MAIN_SITE_TLS, conf.WWW_MAIN_SITE_TLS_KEY)
	TlsConfig.BuildNameToCertificate()
	return TlsConfig
}

func HandleTLS(router *mux.Router) {
	TlsConfig := getTlsConfig()
	Server := http.Server{
		TLSConfig:    TlsConfig,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	TLSListener, _ := tls.Listen("tcp", conf.TLS_PORT, TlsConfig)
	Server.Serve(TLSListener)
}
