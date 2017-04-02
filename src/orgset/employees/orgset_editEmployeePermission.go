package employees

import (
	"net/http"
	"forcamp/src"
	"forcamp/src/orgset"
	"forcamp/conf"
	"database/sql"
	"log"
	"strconv"
)

func EditEmployeePermission(token string, login string, catId int64, value string, ResponseWriter http.ResponseWriter) bool{
	Connection := src.Connect()
	defer Connection.Close()
	if orgset.CheckUserAccess(token, Connection, ResponseWriter) {
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token, Connection)
		if APIerr != nil {
			return conf.PrintError(APIerr, ResponseWriter)
		}
		NewConnection := src.Connect_Custom(Organization)
		defer NewConnection.Close()
		EmployeeOrganization, APIerr := orgset.GetUserOrganizationByLogin(login, Connection)
		if APIerr != nil {
			return conf.PrintError(APIerr, ResponseWriter)
		}
		if EmployeeOrganization != Organization {
			return conf.PrintError(conf.ErrUserNotFound, ResponseWriter)
		}
		if orgset.CheckCategoryId(catId, ResponseWriter, NewConnection) && checkPermissionValue(value, ResponseWriter){
			APIerr = editEmployeePermission_Request(login, catId, value, NewConnection)
			return conf.PrintSuccess(conf.RequestSuccess, ResponseWriter)
		}
	}
	return true
}

func editEmployeePermission_Request(login string, catId int64, value string, connection *sql.DB) *conf.ApiError{
	Query, err := connection.Prepare("UPDATE employees SET `"+strconv.FormatInt(catId, 10)+"`=? WHERE login=?")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(value, login)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func checkPermissionValue(value string, w http.ResponseWriter) bool{
	if value == "false" || value == "true"{
		return true
	} else {
		return conf.PrintError(conf.ErrPermissionValueIncorrect, w)
	}
}
