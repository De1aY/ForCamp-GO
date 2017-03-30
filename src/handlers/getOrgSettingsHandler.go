package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/orgset"
)


func GetOrgSettingsHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodGet {
		orgset.GetOrgSettings(getToken(r), w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleGetOrgSettings(router *mux.Router)  {
	router.HandleFunc("/orgset.settings.get", GetOrgSettingsHandler)
}
