package orgset_edit

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strings"
	"strconv"
	"forcamp/src/api/orgset/categories"
)

func getEditCategoryPostValues(r *http.Request) (categories.Category, string, *conf.ApiResponse){
	Token := r.PostFormValue("token")
	ID, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("id")), 10, 64)
	if err != nil{
		return categories.Category{}, "", conf.ErrIdIsNotINT
	}
	Name := strings.TrimSpace(strings.ToLower(r.PostFormValue("name")))
	NegativeMarks := strings.TrimSpace(strings.ToLower(r.PostFormValue("negative_marks")))
	return categories.Category{ID: ID, Name: Name, NegativeMarks: NegativeMarks}, Token, nil
}

func EditCategoryHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_POST(w)
	if r.Method == http.MethodPost {
		category, token, APIerr := getEditCategoryPostValues(r)
		if APIerr != nil{
			APIerr.Print(w)
		} else {
			categories.EditCategory(token, category, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleEditCategory(router *mux.Router)  {
	router.HandleFunc("/orgset.category.edit", EditCategoryHandler)
}
