package orgset_delete

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strconv"
	"log"
	"forcamp/src/orgset/categories"
)

func getDeleteCategoryPostValues(r *http.Request) (int64, string, *conf.ApiError){
	Token := r.PostFormValue("token")
	ID, err := strconv.ParseInt(r.PostFormValue("id"), 10, 64)
	if err != nil{
		log.Print(err)
		return 0, "", conf.ErrIDisNotINT
	}
	return ID, Token, nil
}

func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		ID, token, err := getDeleteCategoryPostValues(r)
		if err != nil{
			 conf.PrintError(err, w)
		} else {
			categories.DeleteCategory(token, ID, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleDeleteCategory(router *mux.Router)  {
	router.HandleFunc("/orgset.category.delete", DeleteCategoryHandler)
}
