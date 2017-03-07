package folders

import (
	"github.com/gorilla/mux"
	"net/http"
	"forcamp/conf/paths"
)

func getFolder_Scripts() http.Dir{
	ScriptsFolder := http.Dir(paths.FOLDER_SCRIPTS)
	return ScriptsFolder
}

func folderHandler_Scripts() http.Handler{
	ScriptsFolder := getFolder_Scripts()
	ScriptsFileServer := http.FileServer(ScriptsFolder)
	ScriptsHandler := http.StripPrefix("/scripts/", ScriptsFileServer)
	return ScriptsHandler
}

func HandleFolder_Scripts(router *mux.Router){
	ScriptsHandler := folderHandler_Scripts()
	router.PathPrefix("/scripts/").Handler(ScriptsHandler)
}
