package mainSite

import (
	"net/http"
	"forcamp/src"
	"html/template"
	"forcamp/conf"
)

type indexTemplateData struct {
	LoggedIn bool
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.TLS != nil {
		src.SetHeaders_Main(w)
		indexHtml, err := template.ParseFiles(conf.FILE_INDEX); if err != nil {
			w.WriteHeader(http.StatusInternalServerError);
		}
		itd, apiErr := getIndexTemplateData(r); if apiErr != nil && apiErr.Code != 603 {
			w.WriteHeader(http.StatusInternalServerError)
		}
		indexHtml.ExecuteTemplate(w, "index", itd);
	} else {
		http.Redirect(w, r, "https://" + r.Host + r.URL.Path, http.StatusTemporaryRedirect)
	}
}

func getIndexTemplateData(r *http.Request) (indexTemplateData, *conf.ApiResponse) {
	var itd indexTemplateData
	token, err := r.Cookie("token");
	if err != nil {
		itd.LoggedIn = false
		return itd, conf.ErrUserTokenEmpty
	}
	if len(token.Value) > 0 {
		itd.LoggedIn = true
	} else {
		itd.LoggedIn = false
	}
	return itd, nil
}
