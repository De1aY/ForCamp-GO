/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package apanel

import (
	"wplay/conf"
	"wplay/src"
	"wplay/src/api/apanel"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func getAddOrganizationPostValues(r *http.Request) (string, string, string, string, string) {
	token := strings.TrimSpace(r.PostFormValue("token"))
	organizationName := strings.TrimSpace(strings.ToLower(r.PostFormValue("orgname")))
	name := strings.TrimSpace(strings.ToLower(r.PostFormValue("name")))
	surname := strings.TrimSpace(strings.ToLower(r.PostFormValue("surname")))
	middlename := strings.TrimSpace(strings.ToLower(r.PostFormValue("middlename")))
	return organizationName, name, surname, middlename, token
}

func addOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	src.SetHeaders_API_POST(w)
	if r.Method == http.MethodPost {
		organizationName, name, surname, middlename, token := getAddOrganizationPostValues(r)
		apanel.CreateOrganization(token, organizationName, name, surname, middlename, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleAddOrganization(router *mux.Router) {
	router.HandleFunc("/apanel.organization.add", addOrganizationHandler)
}
