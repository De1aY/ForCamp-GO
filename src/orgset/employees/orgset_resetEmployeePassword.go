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

type resetEmployeePassword_Success struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Password string `json:"password"`
}

func (success *resetEmployeePassword_Success) toJSON() string {
	resp, _ := json.Marshal(success)
	return string(resp)
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
		resp, APIerr := resetEmployeePassword_Request(login)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		fmt.Fprintf(ResponseWriter, resp.toJSON())
	}
	return true
}

func resetEmployeePassword_Request(login string) (resetEmployeePassword_Success, *conf.ApiError){
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
	return resetEmployeePassword_Success{200, "success", Password}, nil
}
