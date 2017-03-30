package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/orgset"
	"strings"
)

func getAddCategoryPostValues(r *http.Request) (orgset.Category, string){
	Token := r.PostFormValue("token")
	Name := strings.ToLower(r.PostFormValue("name"))
	NegativeMarks := strings.ToLower(r.PostFormValue("negative_marks"))
	return orgset.Category{ID: 0, Name: Name, NegativeMarks: NegativeMarks}, Token
}

func AddCategoryHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		category, token := getAddCategoryPostValues(r)
		orgset.AddCategory(token, category, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleAddCategory(router *mux.Router)  {
	router.HandleFunc("/orgset.category.add", AddCategoryHandler)
}
