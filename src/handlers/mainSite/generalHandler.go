package mainSite

import (
	"net/http"
	"forcamp/src"
	"forcamp/conf"
	"forcamp/src/tools"
	"text/template"
)

type generalTemplateData struct {
	Token string

}

var gtd generalTemplateData

func GeneralHandler(w http.ResponseWriter, r *http.Request) {
	if r.TLS != nil {
		src.SetHeaders_Main(w)
		token, err := r.Cookie("token");
		if err == nil && tools.CheckToken(token.Value) {
			generalHtml, err := template.ParseFiles(conf.FILE_GENERAL); if err != nil {
				w.WriteHeader(http.StatusInternalServerError);
			}
			gtd.Token = token.Value;
			generalHtml.ExecuteTemplate(w, "index", gtd);
		} else {
			http.Redirect(w, r, "https://forcamp.ga/exit", http.StatusTemporaryRedirect)
		}
	} else {
		http.Redirect(w, r, "https://" + r.Host + r.URL.Path, http.StatusTemporaryRedirect)
	}
}