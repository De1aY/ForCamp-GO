package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"forcamp/conf"
)

func getFolder_MainSite() http.Dir {
	MainSiteFolder := http.Dir(conf.FOLDER_MAIN_SITE)
	return MainSiteFolder
}

func folderHandler_MainSite() http.Handler {
	MainSiteFolder := getFolder_MainSite()
	MainSiteFileServer := http.FileServer(MainSiteFolder)
	MainSiteHandler := http.StripPrefix("/", MainSiteFileServer)
	return MainSiteHandler
}

func HandleFolder_MainSite(router *mux.Router) {
	MainSiteHandler := folderHandler_MainSite()
	router.PathPrefix("/").Handler(MainSiteHandler)
}
