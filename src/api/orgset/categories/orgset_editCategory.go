package categories

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src"
	"strconv"
	"log"
	"forcamp/src/api/orgset"
)

func EditCategory(token string, category Category, ResponseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token, ResponseWriter) && checkCategoryData(category, ResponseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		APIerr = editCategory_Request(category)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		conf.PrintSuccess(conf.RequestSuccess, ResponseWriter)
	}
	return true
}

func editCategory_Request(category Category) *conf.ApiError{
	Query, err := src.CustomConnection.Prepare("UPDATE categories SET name=?, negative_marks=? WHERE id=?")
	if err != nil{
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(category.Name, category.NegativeMarks, strconv.FormatInt(category.ID, 10))
	Query.Close()
	if err != nil{
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return  nil
}