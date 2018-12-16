package mainSite

import (
	"forcamp/conf"
	"forcamp/src"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.TLS != nil {
		src.SetHeaders_Main(w)
		http.ServeFile(w, r, conf.FILE_INDEX)
	} else {
		http.Redirect(w, r, "https://"+r.Host+r.URL.Path, http.StatusTemporaryRedirect)
	}
}
