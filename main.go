package main

import (
	"github.com/gorilla/mux"
	"forcamp/src/handlers"
	"forcamp/conf"
	"net/http"
	"forcamp/src/handlers/api/orgset/orgset_get"
	"forcamp/src/handlers/api/users/users_get"
	"forcamp/src/handlers/api/orgset/orgset_add"
	"forcamp/src/handlers/api/orgset/orgset_edit"
	"forcamp/src/handlers/api/orgset/orgset_delete"
)

func main() {
	Router := mux.NewRouter()
	WWWSite := Router.Host(conf.WWW_MAIN_SITE_DOMAIN).Subrouter()
	MainSite := Router.Host(conf.MAIN_SITE_DOMAIN).Subrouter()
	APISite := Router.Host(conf.API_SITE_DOMAIN).Subrouter()

	//API site
	handlers.HandleAuthorizationByLoginAndPassword(APISite)
	handlers.HandleTokenVerification(APISite)
	// Users: GET
	users_get.HandleGetUserLoginByToken(APISite)
	users_get.HandleGetUserData(APISite)
	// OrgSet: GET
	orgset_get.HandleGetTeams(APISite)
	orgset_get.HandleGetOrgSettings(APISite)
	orgset_get.HandleGetCategories(APISite)
	orgset_get.HandleGetParticipants(APISite)
	orgset_get.HandleGetEmployees(APISite)
	orgset_get.HandleGetParticipantsExcel(APISite)
	// OrgSet: ADD
	orgset_add.HandleAddTeam(APISite)
	orgset_add.HandleAddCategory(APISite)
	orgset_add.HandleAddParticipant(APISite)
	orgset_add.HandleAddEmployee(APISite)
	// OrgSet: EDIT
	orgset_edit.HandleSetOrgSettingValue(APISite)
	orgset_edit.HandleEditCategory(APISite)
	orgset_edit.HandleEditTeam(APISite)
	orgset_edit.HandleResetParticipantPassword(APISite)
	orgset_edit.HandleEditParticipant(APISite)
	orgset_edit.HandleResetEmployeePassword(APISite)
	orgset_edit.HandleEditEmployee(APISite)
	orgset_edit.HandleEditEmployeePermission(APISite)
	// OrgSet: DELETE
	orgset_delete.HandleDeleteCategory(APISite)
	orgset_delete.HandleDeleteTeam(APISite)
	orgset_delete.HandleDeleteParticipant(APISite)
	orgset_delete.HandleDeleteEmployee(APISite)

	//Main site
	handlers.HandleFolder_MainSite(WWWSite)
	handlers.HandleFolder_MainSite(MainSite)

	//Server
	go http.ListenAndServe(conf.SERVER_PORT, Router)
	handlers.HandleTLS(Router)
}
