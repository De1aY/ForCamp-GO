package categories

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src"
	"log"
	"strconv"
	"forcamp/src/api/orgset"
)

func DeleteCategory(token string, id int64, responseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token ,responseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil{
			return APIerr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		APIerr = deleteCategory_Request(id)
		if APIerr != nil{
			return APIerr.Print(responseWriter)
		}
		return conf.RequestSuccess.Print(responseWriter)
	}
	return true
}

func deleteCategory_Request(id int64) *conf.ApiResponse{
	Query, err := src.CustomConnection.Prepare("DELETE FROM categories WHERE id=?")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(id)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	APIerr := deleteCategory_Participants(id)
	if APIerr != nil{
		return APIerr
	}
	APIerr = deleteCategory_Employees(id)
	if APIerr != nil{
		return APIerr
	}
	return nil
}

func deleteCategory_Participants(id int64) *conf.ApiResponse{
	_, err := src.CustomConnection.Query("ALTER TABLE participants DROP COLUMN `"+strconv.FormatInt(id, 10)+"`")
	if err != nil{
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func deleteCategory_Employees(id int64) *conf.ApiResponse{
	_, err := src.CustomConnection.Query("ALTER TABLE employees DROP COLUMN `"+strconv.FormatInt(id, 10)+"`")
	if err != nil{
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}