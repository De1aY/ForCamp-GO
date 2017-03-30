package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/orgset"
	"strings"
	"strconv"
	"log"
)

func getEditCategoryPostValues(r *http.Request) (orgset.Category, string, *conf.ApiError){
	Token := r.PostFormValue("token")
	ID, err := strconv.ParseInt(r.PostFormValue("id"), 10, 64)
	if err != nil{
		log.Print(err)
		return orgset.Category{}, "", conf.ErrIDisNotINT
	}
	Name := strings.ToLower(r.PostFormValue("name"))
	NegativeMarks := strings.ToLower(r.PostFormValue("negative_marks"))
	return orgset.Category{ID: ID, Name: Name, NegativeMarks: NegativeMarks}, Token, nil
}

func EditCategoryHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		category, token, APIerr := getEditCategoryPostValues(r)
		if APIerr != nil{
			conf.PrintError(APIerr, w)
		} else {
			orgset.EditCategory(token, category, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleEditCategory(router *mux.Router)  {
	router.HandleFunc("/orgset.category.edit", EditCategoryHandler)
}
