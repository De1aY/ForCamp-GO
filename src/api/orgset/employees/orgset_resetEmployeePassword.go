/*
	Copyright: "Null team", 2016 - 2017
	Author: "De1aY"
	Documentation: https://bitbucket.org/lyceumdevelopers/golang/wiki/Home
*/
package employees

import (
	"net/http"
	"forcamp/src/api/orgset"
	"forcamp/conf"
	"forcamp/src"
	"log"
)

type resetEmployeePassword_Success struct {
	Password string `json:"password"`
}


func ResetEmployeePassword(token string, login string, responseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token, responseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil{
			return APIerr.Print(responseWriter)
		}
		EmployeeOrganization, APIerr := orgset.GetUserOrganizationByLogin(login)
		if APIerr != nil{
			return APIerr.Print(responseWriter)
		}
		if EmployeeOrganization != Organization{
			return conf.ErrUserNotFound.Print(responseWriter)
		}
		rawResp, APIerr := resetEmployeePassword_Request(login)
		if APIerr != nil{
			return APIerr.Print(responseWriter)
		}
		resp := conf.ApiResponse{200, "success", rawResp}
		resp.Print(responseWriter)
	}
	return true
}

func resetEmployeePassword_Request(login string) (resetEmployeePassword_Success, *conf.ApiResponse){
	Password, Hash := orgset.GeneratePassword()
	Query, err := src.Connection.Prepare("UPDATE users SET password=? WHERE login=?")
	if err != nil {
		log.Print(err)
		return resetEmployeePassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(Hash, login)
	if err != nil {
		log.Print(err)
		return resetEmployeePassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	Query, err = src.Connection.Prepare("DELETE FROM sessions WHERE login=?")
	if err != nil {
		log.Print(err)
		return resetEmployeePassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(login)
	if err != nil {
		log.Print(err)
		return resetEmployeePassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	return resetEmployeePassword_Success{Password}, nil
}
