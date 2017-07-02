/*
	Copyright: "Null team", 2016 - 2017
	Author: "De1aY"
	Documentation: https://bitbucket.org/lyceumdevelopers/golang/wiki/Home
*/
package employees

import (
	"net/http"
	"forcamp/src"
	"forcamp/src/api/orgset"
	"forcamp/conf"
	"log"
	"strconv"
)

func EditEmployeePermission(token string, login string, catId int64, value string, responseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token, responseWriter) {
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil {
			return APIerr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		EmployeeOrganization, APIerr := orgset.GetUserOrganizationByLogin(login)
		if APIerr != nil {
			return APIerr.Print(responseWriter)
		}
		if EmployeeOrganization != Organization {
			return conf.ErrUserNotFound.Print(responseWriter)
		}
		if orgset.CheckCategoryId(catId, responseWriter) && checkPermissionValue(value, responseWriter){
			APIerr = editEmployeePermission_Request(login, catId, value)
			return conf.RequestSuccess.Print(responseWriter)
		}
	}
	return true
}

func editEmployeePermission_Request(login string, catId int64, value string) *conf.ApiResponse{
	Query, err := src.CustomConnection.Prepare("UPDATE employees SET `"+strconv.FormatInt(catId, 10)+"`=? WHERE login=?")
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
		return conf.ErrPermissionValueIncorrect.Print(w)
	}
}
