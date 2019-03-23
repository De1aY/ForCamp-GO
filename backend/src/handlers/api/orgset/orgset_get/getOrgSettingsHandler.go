package orgset_get

import (
	"net/http"
	"github.com/gorilla/mux"
	"wplay/conf"
	"wplay/src"
	"wplay/src/api/orgset/settings"
	"wplay/src/handlers"
)

func GetOrgSettingsHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_GET(w)
	if r.Method == http.MethodGet {
		settings.GetOrgSettings(handlers.GetToken(r), w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleGetOrgSettings(router *mux.Router)  {
	router.HandleFunc("/orgset.settings.get", GetOrgSettingsHandler)
}
