package orgset_get

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/api/orgset/settings"
	"forcamp/src/handlers"
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
