/*
	Copyright: "Null team", 2016 - 2017
	Author: "De1aY"
	Documentation: https://bitbucket.org/lyceumdevelopers/golang/wiki/Home
*/
package main

import (
	"github.com/gorilla/mux"
	"forcamp/src/handlers"
	"forcamp/src"
	"forcamp/conf"
	"net/http"
	"forcamp/src/handlers/api/orgset/orgset_get"
	"forcamp/src/handlers/api/users/users_get"
	"forcamp/src/handlers/api/orgset/orgset_add"
	"forcamp/src/handlers/api/orgset/orgset_edit"
	"forcamp/src/handlers/api/orgset/orgset_delete"
	"forcamp/src/handlers/api/marks"
	"forcamp/src/handlers/api/apanel"
)

func main() {
	// Domains routing
	Router := mux.NewRouter()
	WWWSite := Router.Host(conf.WWW_MAIN_SITE_DOMAIN).Subrouter()
	MainSite := Router.Host(conf.MAIN_SITE_DOMAIN).Subrouter()
	APISite := Router.Host(conf.API_SITE_DOMAIN).Subrouter()

	// Handlers: API site
	// Authorization
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
	orgset_get.HandleGetReasons(APISite)
	orgset_get.HandleGetEmployeesExcel(APISite)
	// OrgSet: ADD
	orgset_add.HandleAddTeam(APISite)
	orgset_add.HandleAddCategory(APISite)
	orgset_add.HandleAddParticipant(APISite)
	orgset_add.HandleAddEmployee(APISite)
	orgset_add.HandleAddReason(APISite)
	// OrgSet: EDIT
	orgset_edit.HandleSetOrgSettingValue(APISite)
	orgset_edit.HandleEditCategory(APISite)
	orgset_edit.HandleEditTeam(APISite)
	orgset_edit.HandleResetParticipantPassword(APISite)
	orgset_edit.HandleEditParticipant(APISite)
	orgset_edit.HandleResetEmployeePassword(APISite)
	orgset_edit.HandleEditEmployee(APISite)
	orgset_edit.HandleEditEmployeePermission(APISite)
	orgset_edit.HandleEditReason(APISite)
	// OrgSet: DELETE
	orgset_delete.HandleDeleteCategory(APISite)
	orgset_delete.HandleDeleteTeam(APISite)
	orgset_delete.HandleDeleteParticipant(APISite)
	orgset_delete.HandleDeleteEmployee(APISite)
	orgset_delete.HandleDeleteReason(APISite)
	// Marks
	marks.HandleEditMark(APISite)
	marks.HandleGetMarksChanges(APISite)
	marks.HandleDeleteMarkChange(APISite)
	// Apanel
	apanel.HandleAddOrganization(APISite)

	// Handlers: Main site
	handlers.HandleFolder_MainSite(WWWSite)
	handlers.HandleFolder_MainSite(MainSite)
	handlers.HandleExit(WWWSite)
	handlers.HandleExit(MainSite)

	// Database: "forcamp"
	src.Connection = src.Connect()

	// Server
	go handlers.HandleTLS(Router)
	http.ListenAndServe(conf.SERVER_PORT, Router)
}
