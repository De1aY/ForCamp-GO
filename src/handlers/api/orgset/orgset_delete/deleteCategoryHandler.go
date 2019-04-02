/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package orgset_delete

import (
	"net/http"
	"github.com/gorilla/mux"
	"wplay/conf"
	"wplay/src"
	"strconv"
	"wplay/src/api/orgset/categories"
	"strings"
)

func getDeleteCategoryPostValues(r *http.Request) (int64, string, *conf.ApiResponse){
	Token := strings.TrimSpace(r.PostFormValue("token"))
	ID, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("category_id")), 10, 64)
	if err != nil{
		return 0, "", conf.ErrIdIsNotINT
	}
	return ID, Token, nil
}

func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_POST(w)
	if r.Method == http.MethodPost {
		ID, token, err := getDeleteCategoryPostValues(r)
		if err != nil{
			 err.Print(w)
		} else {
			categories.DeleteCategory(token, ID, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleDeleteCategory(router *mux.Router)  {
	router.HandleFunc("/orgset.category.delete", DeleteCategoryHandler)
}
