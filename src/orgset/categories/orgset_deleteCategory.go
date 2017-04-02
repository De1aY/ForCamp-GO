package categories

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src"
	"log"
	"strconv"
	"forcamp/src/orgset"
	"database/sql"
)

func DeleteCategory(token string, id int64, ResponseWriter http.ResponseWriter) bool{
	Connection := src.Connect()
	defer Connection.Close()
	if orgset.CheckUserAccess(token, Connection ,ResponseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token, Connection)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		NewConnection := src.Connect_Custom(Organization)
		defer NewConnection.Close()
		APIerr = deleteCategory_Request(id, NewConnection)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		return conf.PrintSuccess(conf.RequestSuccess, ResponseWriter)
	}
	return true
}

func deleteCategory_Request(id int64, Connection *sql.DB) *conf.ApiError{
	Query, err := Connection.Prepare("DELETE FROM categories WHERE id=?")
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
	APIerr := deleteCategory_Participants(id, Connection)
	if APIerr != nil{
		return APIerr
	}
	APIerr = deleteCategory_Employees(id, Connection)
	if APIerr != nil{
		return APIerr
	}
	return nil
}

func deleteCategory_Participants(id int64, Connection *sql.DB) *conf.ApiError{
	_, err := Connection.Query("ALTER TABLE participants DROP COLUMN `"+strconv.FormatInt(id, 10)+"`")
	if err != nil{
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func deleteCategory_Employees(id int64, Connection *sql.DB) *conf.ApiError{
	_, err := Connection.Query("ALTER TABLE employees DROP COLUMN `"+strconv.FormatInt(id, 10)+"`")
	if err != nil{
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}