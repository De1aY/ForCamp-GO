package folders

import (
	"github.com/gorilla/mux"
	"net/http"
	"forcamp/conf/paths"
)

func getFolder_Fonts() http.Dir{
	FontsFolder := http.Dir(paths.FOLDER_FONTS)
	return FontsFolder
}

func folderHandler_Fonts() http.Handler{
	FontsFolder := getFolder_Fonts()
	FontsFileServer := http.FileServer(FontsFolder)
	FontsHandler := http.StripPrefix("/fonts/", FontsFileServer)
	return FontsHandler
}

func HandleFolder_Fonts(router *mux.Router){
	FontsHandler := folderHandler_Fonts()
	router.PathPrefix("/fonts/").Handler(FontsHandler)
}
