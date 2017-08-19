package orgset_get

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/api/orgset/categories"
	"forcamp/src/handlers"
)

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_GET(w)
	if r.Method == http.MethodGet {
		categories.GetCategories(handlers.GetToken(r), w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleGetCategories(router *mux.Router)  {
	router.HandleFunc("/orgset.categories.get", GetCategoriesHandler)
}
