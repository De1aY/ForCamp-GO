package main

import (
	"github.com/gorilla/mux"
	"forcamp/src/handlers"
	"forcamp/conf"
	"forcamp/src/handlers/templates"
	"forcamp/src/handlers/folders"
)

func main() {
	Router := mux.NewRouter()

	WWWSite := Router.Host(conf.WWW_MAIN_SITE_DOMAIN).Subrouter()
	MainSite := Router.Host(conf.MAIN_SITE_DOMAIN).Subrouter()
	APISite := Router.Host(conf.API_SITE_DOMAIN).Subrouter()
	handlers.HandleDefault(WWWSite)
	handlers.HandleDefault(MainSite)
	handlers.HandleDefault(APISite)
	templates.HandleTemplate_Index(WWWSite)
	templates.HandleTemplate_Index(MainSite)
	folders.HandleFolder_CSS(WWWSite)
	folders.HandleFolder_Fonts(WWWSite)
	folders.HandleFolder_Scripts(WWWSite)
	folders.HandleFolder_CSS(MainSite)
	folders.HandleFolder_Fonts(MainSite)
	folders.HandleFolder_Scripts(MainSite)
	handlers.HandleHTTP(Router)
	handlers.HandleTLS(Router)
}
