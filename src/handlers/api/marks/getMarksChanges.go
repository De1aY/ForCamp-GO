package marks

import (
	"net/http"
	"forcamp/src"
	"forcamp/conf"
	"forcamp/src/api/marks"
	"github.com/gorilla/mux"
	"strings"
)

func getToken(r *http.Request) string {
	token := strings.TrimSpace(r.FormValue("token"))
	return token
}

func getMarksChangesHandler(w http.ResponseWriter, r *http.Request) {
	src.SetHeaders_API(w)
	if r.Method == http.MethodGet {
		token := getToken(r)
		marks.GetMarksChanges(token, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleGetMarksChanges(router *mux.Router) {
	router.HandleFunc("/marks.changes.get", getMarksChangesHandler)
}