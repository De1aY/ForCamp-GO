package orgset_edit

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strings"
	"forcamp/src/api/orgset/settings"
)

func getSetOrgSettingsValuePostValues(r *http.Request) (string, settings.OrgSettings){
	token := strings.TrimSpace(r.PostFormValue("token"))
	var orgSet settings.OrgSettings
	orgSet.Participant = strings.TrimSpace(strings.ToLower(r.PostFormValue("participant")))
	orgSet.Period = strings.TrimSpace(strings.ToLower(r.PostFormValue("period")))
	orgSet.Team = strings.TrimSpace(strings.ToLower(r.PostFormValue("team")))
	orgSet.Organization = strings.TrimSpace(strings.ToLower(r.PostFormValue("organization")))
	orgSet.SelfMarks = strings.TrimSpace(strings.ToLower(r.PostFormValue("self_marks")))
	return token, orgSet
}

func SetOrgSettingsValueHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		token, orgSet := getSetOrgSettingsValuePostValues(r)
		settings.SetOrgSettingsValue(token, orgSet, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleSetOrgSettingsValue(router *mux.Router)  {
	router.HandleFunc("/orgset.settings.edit", SetOrgSettingsValueHandler)
}
