package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"forcamp/conf"
	"forcamp/src"
)

func getFolder_App() http.Dir {
	AppFolder := http.Dir(conf.FOLDER_APP)
	return AppFolder
}

func getFolder_Scss() http.Dir {
	ScssFolder := http.Dir(conf.FOLDER_SCSS)
	return ScssFolder
}

func getFolder_Node_Modules() http.Dir {
	NodeModulesFolder := http.Dir(conf.FOLDER_NODE_MODULES)
	return NodeModulesFolder
}

func getFolder_Media() http.Dir {
	MediaFolder := http.Dir(conf.FOLDER_MEDIA)
	return MediaFolder
}

func folderHandler_App() http.Handler {
	AppFolder := getFolder_App()
	AppFileServer := http.FileServer(AppFolder)
	AppHandler := http.StripPrefix("/app", AppFileServer)
	return AppHandler
}

func folderHandler_Scss() http.Handler {
	ScssFolder := getFolder_Scss()
	ScssFileServer := http.FileServer(ScssFolder)
	ScssHandler := http.StripPrefix("/scss", ScssFileServer)
	return ScssHandler
}

func folderHandler_Node_Modules() http.Handler {
	NodeModulesFolder := getFolder_Node_Modules()
	NodeModulesFileServer := http.FileServer(NodeModulesFolder)
	NodeModulesHandler := http.StripPrefix("/node_modules", NodeModulesFileServer)
	return NodeModulesHandler
}

func folderHandler_Media() http.Handler {
	MediaFolder := getFolder_Media()
	MediaFileServer := http.FileServer(MediaFolder)
	MediaHandler := http.StripPrefix("/media", MediaFileServer)
	return MediaHandler
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	if r.TLS != nil {
		src.SetHeaders_Main(w)
		http.ServeFile(w, r, conf.FILE_INDEX)
	} else {
		http.Redirect(w, r, "https://"+r.Host+r.URL.Path, http.StatusTemporaryRedirect)
	}
}

func systemConfigHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, conf.FILE_SYSTEM_CONFIG)
}

func HandleFolder_MainSite(router *mux.Router) {
	//Folders
	AppHandler := folderHandler_App()
	ScssHandler := folderHandler_Scss()
	MediaHandler := folderHandler_Media()
	NodeModulesHandler := folderHandler_Node_Modules()
	router.PathPrefix("/app").Handler(AppHandler)
	router.PathPrefix("/scss").Handler(ScssHandler)
	router.PathPrefix("/node_modules").Handler(NodeModulesHandler)
	router.PathPrefix("/media").Handler(MediaHandler)
	//Pages
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/main", indexHandler)
	router.HandleFunc("/orgset", indexHandler)
	//SystemConfig
	router.HandleFunc("/systemjs.config.js", systemConfigHandler)
}
