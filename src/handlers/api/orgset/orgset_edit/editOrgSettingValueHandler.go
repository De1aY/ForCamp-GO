package orgset_edit

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strings"
	"forcamp/src/orgset/settings"
)

func getSetOrgSettingValuePostValues(r *http.Request) (string, string, string){
	Token := r.PostFormValue("token")
	Name := strings.ToLower(r.PostFormValue("name"))
	Value := strings.ToLower(r.PostFormValue("value"))
	return Token, Name, Value
}

func SetOrgSettingValueHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		Token, Name, Value := getSetOrgSettingValuePostValues(r)
		settings.SetOrgSettingValue(Token, Name, Value, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleSetOrgSettingValue(router *mux.Router)  {
	router.HandleFunc("/orgset.setting.edit", SetOrgSettingValueHandler)
}
