package orgset_get

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/handlers"
	"forcamp/src/api/orgset/reasons"
)


func GetReasonsHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodGet {
		reasons.GetReasons(handlers.GetToken(r), w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleGetReasons(router *mux.Router)  {
	router.HandleFunc("/orgset.reasons.get", GetReasonsHandler)
}
