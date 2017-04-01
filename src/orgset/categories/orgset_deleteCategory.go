package categories

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src"
	"log"
	"strconv"
	"forcamp/src/orgset"
)

func DeleteCategory(token string, id int64, ResponseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token, ResponseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		src.NewConnection = src.Connect_Custom(Organization)
		APIerr = deleteCategory_Request(id)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		return conf.PrintSuccess(conf.RequestSuccess, ResponseWriter)
	}
	return true
}

func deleteCategory_Request(id int64) *conf.ApiError{
	Query, err := src.NewConnection.Prepare("DELETE FROM categories WHERE id=?")
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

func deleteCategory_Participants(id int64) *conf.ApiError{
	_, err := src.NewConnection.Query("ALTER TABLE participants DROP COLUMN `"+strconv.FormatInt(id, 10)+"`")
	if err != nil{
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func deleteCategory_Employees(id int64) *conf.ApiError{
	_, err := src.NewConnection.Query("ALTER TABLE employees DROP COLUMN `"+strconv.FormatInt(id, 10)+"`")
	if err != nil{
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}