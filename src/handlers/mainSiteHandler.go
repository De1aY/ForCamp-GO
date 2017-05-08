package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/handlers/mainSite"
)

func getFolder_Scripts() http.Dir {
	ScriptsFolder := http.Dir(conf.FOLDER_SCRIPTS)
	return ScriptsFolder
}

func getFolder_Css() http.Dir {
	CssFolder := http.Dir(conf.FOLDER_CSS)
	return CssFolder
}

func getFolder_Images() http.Dir {
	ImagesFolder := http.Dir(conf.FOLDER_IMAGES)
	return ImagesFolder
}

func getFolder_Node_Modules() http.Dir {
	NodeModulesFolder := http.Dir(conf.FOLDER_NODE_MODULES)
	return NodeModulesFolder
}

func getFolder_Media() http.Dir {
	MediaFolder := http.Dir(conf.FOLDER_MEDIA)
	return MediaFolder
}

func folderHandler_Scripts() http.Handler {
	ScriptsFolder := getFolder_Scripts()
	ScriptsFileServer := http.FileServer(ScriptsFolder)
	ScriptsHandler := http.StripPrefix("/scripts", ScriptsFileServer)
	return ScriptsHandler
}

func folderHandler_css() http.Handler {
	CssFolder := getFolder_Css()
	CssFileServer := http.FileServer(CssFolder)
	CssHandler := http.StripPrefix("/css", CssFileServer)
	return CssHandler
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

func folderHandler_Images() http.Handler {
	ImagesFolder := getFolder_Images()
	ImagesFolderServer := http.FileServer(ImagesFolder)
	ImagesHandler := http.StripPrefix("/images", ImagesFolderServer)
	return ImagesHandler
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	if r.TLS != nil {
		src.SetHeaders_Main(w)
		http.ServeFile(w, r, conf.FILE_INDEX)
	} else {
		http.Redirect(w, r, "https://"+r.Host+r.URL.Path, http.StatusTemporaryRedirect)
	}
}

func generalHandler(w http.ResponseWriter, r *http.Request){
	if r.TLS != nil {
		src.SetHeaders_Main(w)
		http.ServeFile(w, r, conf.FILE_GENERAL)
	} else {
		http.Redirect(w, r, "https://"+r.Host+r.URL.Path, http.StatusTemporaryRedirect)
	}
}

func orgSetHandler(w http.ResponseWriter, r *http.Request){
	if r.TLS != nil {
		src.SetHeaders_Main(w)
		http.ServeFile(w, r, conf.FILE_ORGSET)
	} else {
		http.Redirect(w, r, "https://"+r.Host+r.URL.Path, http.StatusTemporaryRedirect)
	}
}

func HandleFolder_MainSite(router *mux.Router) {
	// Folders
	ScriptsHandler := folderHandler_Scripts()
	CssHandler := folderHandler_css()
	MediaHandler := folderHandler_Media()
	NodeModulesHandler := folderHandler_Node_Modules()
	ImagesHandler := folderHandler_Images()
	router.PathPrefix("/scripts").Handler(ScriptsHandler)
	router.PathPrefix("/css").Handler(CssHandler)
	router.PathPrefix("/node_modules").Handler(NodeModulesHandler)
	router.PathPrefix("/media").Handler(MediaHandler)
	router.PathPrefix("/images").Handler(ImagesHandler)
	// Pages
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/main", indexHandler)
	router.HandleFunc("/orgset", orgSetHandler)
	router.HandleFunc("/marks", indexHandler)
	router.HandleFunc("/general", mainSite.GeneralHandler)
	router.HandleFunc("/profile", indexHandler)
	router.HandleFunc("/team", indexHandler)
	router.HandleFunc("/achievements", indexHandler)
	router.HandleFunc("/profile/{login}", indexHandler)
	router.HandleFunc("/apanel", indexHandler)
}
