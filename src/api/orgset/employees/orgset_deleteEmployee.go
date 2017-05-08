/*
	Copyright: "Null team", 2016 - 2017
	Author: "De1aY"
	Documentation: https://bitbucket.org/lyceumdevelopers/golang/wiki/Home
*/
package employees

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src"
	"log"
	"forcamp/src/api/orgset"
)

func DeleteEmployee(token string, login string, ResponseWriter http.ResponseWriter) bool {
	if orgset.CheckUserAccess(token, ResponseWriter) {
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil {
			return conf.PrintError(APIerr, ResponseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		APIerr = deleteEmployee_Request(login)
		if APIerr != nil {
			return conf.PrintError(APIerr, ResponseWriter)
		}
		return conf.PrintSuccess(conf.RequestSuccess, ResponseWriter)
	}
	return true
}

func deleteEmployee_Request(login string) *conf.ApiError {
	APIerr := deleteEmployee_Organization(login)
	if APIerr != nil {
		return APIerr
	}
	APIerr = deleteEmployee_Main(login)
	if APIerr != nil {
		return APIerr
	}
	return nil
}

func deleteEmployee_Main(login string) *conf.ApiError {
	Query, err := src.Connection.Prepare("DELETE FROM users WHERE login=?")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(login)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	Query, err = src.Connection.Prepare("DELETE FROM sessions WHERE login=?")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(login)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func deleteEmployee_Organization(login string) *conf.ApiError {
	Query, err := src.CustomConnection.Prepare("DELETE FROM users WHERE login=? AND access='1'")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	resp, err := Query.Exec(login)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	rowsAffected, err := resp.RowsAffected()
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	if rowsAffected == 0 {
		return conf.ErrUserNotFound
	}
	Query, err = src.CustomConnection.Prepare("DELETE FROM employees WHERE login=?")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(login)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}
