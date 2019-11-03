/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package orgset_edit

import (
	"net/http"
	"github.com/gorilla/mux"
	"nullteam.info/wplay/demo/conf"
	"nullteam.info/wplay/demo/src"
	"strings"
	"strconv"
	"nullteam.info/wplay/demo/src/api/orgset/categories"
)

func getEditCategoryPostValues(r *http.Request) (categories.Category, string, *conf.ApiResponse){
	token := r.PostFormValue("token")
	category_id, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("category_id")), 10, 64)
	if err != nil{
		return categories.Category{}, "", conf.ErrIdIsNotINT
	}
	Name := strings.TrimSpace(strings.ToLower(r.PostFormValue("name")))
	NegativeMarks := strings.TrimSpace(strings.ToLower(r.PostFormValue("negative_marks")))
	return categories.Category{ID: category_id, Name: Name, NegativeMarks: NegativeMarks}, token, nil
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
