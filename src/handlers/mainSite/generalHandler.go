package mainSite

import (
	"net/http"
	"forcamp/src"
	"forcamp/conf"
	"forcamp/src/tools"
	"fmt"
)

func GeneralHandler(w http.ResponseWriter, r *http.Request){
	if r.TLS != nil {
		src.SetHeaders_Main(w)
		token, err := r.Cookie("token");
		if err != nil {
			http.Redirect(w, r, "https://forcamp.ga/exit", http.StatusTemporaryRedirect)
		} else {
			if tools.CheckToken(token.Value) {
				http.ServeFile(w, r, conf.FILE_GENERAL)
			} else {
				http.Redirect(w, r, "https://forcamp.ga/exit", http.StatusTemporaryRedirect)
			}
		}
	} else {
		http.Redirect(w, r, "https://"+r.Host+r.URL.Path, http.StatusTemporaryRedirect)
	}
}