package folders

import (
	"github.com/gorilla/mux"
	"net/http"
	"forcamp/conf/paths"
)

func getFolder_CSS() http.Dir{
	CSSFolder := http.Dir(paths.FOLDER_CSS)
	return CSSFolder
}

func folderHandler_CSS() http.Handler{
	CSSFolder := getFolder_CSS()
	CSSFileServer := http.FileServer(CSSFolder)
	CSSHandler := http.StripPrefix("/css/", CSSFileServer)
	return CSSHandler
}

func HandleFolder_CSS(router *mux.Router){
	CSSHandler := folderHandler_CSS()
	router.PathPrefix("/css/").Handler(CSSHandler)
}
