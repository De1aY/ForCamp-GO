/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package orgset_edit

import (
	"net/http"
	"github.com/gorilla/mux"
	"wplay/conf"
	"wplay/src"
	"strings"
	"wplay/src/api/orgset/settings"
)

func getSetOrgSettingValuePostValues(r *http.Request) (string, string, string){
	token := strings.TrimSpace(r.PostFormValue("token"))
	setting_name := strings.ToLower(strings.TrimSpace(r.PostFormValue("setting_name")))
	setting_value := strings.ToLower(strings.TrimSpace(r.PostFormValue("setting_value")))
	return token, setting_name, setting_value
}

func SetOrgSettingValueHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_POST(w)
	if r.Method == http.MethodPost {
		token, setting_name, setting_value := getSetOrgSettingValuePostValues(r)
		settings.SetOrgSettingValue(token, setting_name, setting_value, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleSetOrgSettingsValue(router *mux.Router)  {
	router.HandleFunc("/orgset.setting.edit", SetOrgSettingValueHandler)
}
