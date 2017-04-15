/*
	Copyright: "Null team", 2016 - 2017
	Author: "De1aY"
	Documentation: https://bitbucket.org/lyceumdevelopers/golang/wiki/Home
*/
package employees

import (
	"net/http"
	"forcamp/src/orgset"
	"forcamp/conf"
	"forcamp/src"
	"log"
	"encoding/json"
	"fmt"
)

type ResetEmployeePassword_Success struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Password string `json:"password"`
}

func ResetEmployeePassword(token string, login string, ResponseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token, ResponseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		EmployeeOrganization, APIerr := orgset.GetUserOrganizationByLogin(login)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		if EmployeeOrganization != Organization{
			return conf.PrintError(conf.ErrUserNotFound, ResponseWriter)
		}
		Resp, APIerr := resetEmployeePassword_Request(login)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		Response, _ := json.Marshal(Resp)
		fmt.Fprintf(ResponseWriter, string(Response))
	}
	return true
}

func resetEmployeePassword_Request(login string) (ResetEmployeePassword_Success, *conf.ApiError){
	Password, Hash := orgset.GeneratePassword()
	Query, err := src.Connection.Prepare("UPDATE users SET password=? WHERE login=?")
	if err != nil {
		log.Print(err)
		return ResetEmployeePassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(Hash, login)
	if err != nil {
		log.Print(err)
		return ResetEmployeePassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	Query, err = src.Connection.Prepare("DELETE FROM sessions WHERE login=?")
	if err != nil {
		log.Print(err)
		return ResetEmployeePassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(login)
	if err != nil {
		log.Print(err)
		return ResetEmployeePassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	return ResetEmployeePassword_Success{200, "success", Password}, nil
}
