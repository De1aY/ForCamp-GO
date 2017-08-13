package categories

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src"
	"strconv"
	"forcamp/src/api/orgset"
)

func EditCategory(token string, category Category, responseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token, responseWriter) && checkCategoryData(category, responseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil{
			return APIerr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		APIerr = editCategory_Request(category)
		if APIerr != nil{
			return APIerr.Print(responseWriter)
		}
		conf.RequestSuccess.Print(responseWriter)
	}
	return true
}

func editCategory_Request(category Category) *conf.ApiResponse{
	Query, err := src.CustomConnection.Prepare("UPDATE categories SET name=?, negative_marks=? WHERE id=?")
	if err != nil{
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(category.Name, category.NegativeMarks, strconv.FormatInt(category.ID, 10))
	Query.Close()
	if err != nil{
		return conf.ErrDatabaseQueryFailed
	}
	return  nil
}