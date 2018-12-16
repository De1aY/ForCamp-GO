package handlers

import (
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/handlers/mainSite"
	"net/http"

	"github.com/gorilla/mux"
)

func getFolderScripts() http.Dir {
	ScriptsFolder := http.Dir(conf.FOLDER_SCRIPTS)
	return ScriptsFolder
}

func getFolderCSS() http.Dir {
	CSSFolder := http.Dir(conf.FOLDER_CSS)
	return CSSFolder
}

func getFolderImages() http.Dir {
	ImagesFolder := http.Dir(conf.FOLDER_IMAGES)
	return ImagesFolder
}

func folderHandlerScripts() http.Handler {
	ScriptsFolder := getFolderScripts()
	ScriptsFileServer := http.FileServer(ScriptsFolder)
	ScriptsHandler := http.StripPrefix("/js", ScriptsFileServer)
	return ScriptsHandler
}

func folderHandlerCSS() http.Handler {
	CSSFolder := getFolderCSS()
	CSSFileServer := http.FileServer(CSSFolder)
	CSSHandler := http.StripPrefix("/css", CSSFileServer)
	return CSSHandler
}

func folderHandlerImages() http.Handler {
	ImagesFolder := getFolderImages()
	ImagesFolderServer := http.FileServer(ImagesFolder)
	ImagesHandler := http.StripPrefix("/img", ImagesFolderServer)
	return ImagesHandler
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.TLS != nil {
		src.SetHeaders_Main(w)
		http.ServeFile(w, r, conf.FILE_INDEX)
	} else {
		http.Redirect(w, r, "https://"+r.Host+r.URL.Path, http.StatusTemporaryRedirect)
	}
}

func HandleMainSite(router *mux.Router) {
	// Folders
	ScriptsHandler := folderHandlerScripts()
	CSSHandler := folderHandlerCSS()
	ImagesHandler := folderHandlerImages()
	router.PathPrefix("/js").Handler(ScriptsHandler)
	router.PathPrefix("/css").Handler(CSSHandler)
	router.PathPrefix("/img").Handler(ImagesHandler)
	// Pages
	router.HandleFunc("/", mainSite.IndexHandler)
}
