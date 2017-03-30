package main

import (
	"github.com/gorilla/mux"
	"forcamp/src/handlers"
	"forcamp/conf"
	"net/http"
)

func main() {
	Router := mux.NewRouter()
	WWWSite := Router.Host(conf.WWW_MAIN_SITE_DOMAIN).Subrouter()
	MainSite := Router.Host(conf.MAIN_SITE_DOMAIN).Subrouter()
	APISite := Router.Host(conf.API_SITE_DOMAIN).Subrouter()

	//API
	handlers.HandleAuthorizationByLoginAndPassword(APISite)
	handlers.HandleTokenVerification(APISite)
	handlers.HandleGetUserLoginByToken(APISite)
	handlers.HandleGetUserData(APISite)
	handlers.HandleSetOrgSettingValue(APISite)
	handlers.HandleGetOrgSettings(APISite)
	handlers.HandleGetCategories(APISite)
	handlers.HandleAddCategory(APISite)
	handlers.HandleDeleteCategory(APISite)
	handlers.HandleEditCategory(APISite)
	handlers.HandleGetTeams(APISite)
	handlers.HandleAddTeam(APISite)
	handlers.HandleDeleteTeam(APISite)
	handlers.HandleTeamCategory(APISite)

	//Main
	handlers.HandleFolder_MainSite(WWWSite)
	handlers.HandleFolder_MainSite(MainSite)
	go http.ListenAndServe(conf.SERVER_PORT, Router)
	handlers.HandleTLS(Router)
}
