package main

import (
	"github.com/gorilla/mux"
	"forcamp/src/handlers"
	"forcamp/conf"
)

func main() {
	Router := mux.NewRouter()
	WWWSite := Router.Host(conf.WWW_MAIN_SITE_DOMAIN).Subrouter()
	MainSite := Router.Host(conf.MAIN_SITE_DOMAIN).Subrouter()
	//APISite := Router.Host(conf.API_SITE_DOMAIN).Subrouter()
	handlers.HandleFolder_MainSite(WWWSite)
	handlers.HandleFolder_MainSite(MainSite)
	handlers.HandleTLS(Router)
}
