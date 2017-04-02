package employees

import (
	"net/http"
	"forcamp/src/orgset"
	"forcamp/conf"
	"forcamp/src"
	"log"
	"encoding/json"
	"fmt"
	"database/sql"
)

type ResetEmployeePassword_Success struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Password string `json:"password"`
}

func ResetEmployeePassword(token string, login string, ResponseWriter http.ResponseWriter) bool{
	Connection := src.Connect()
	defer Connection.Close()
	if orgset.CheckUserAccess(token, Connection, ResponseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token, Connection)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		EmployeeOrganization, APIerr := orgset.GetUserOrganizationByLogin(login, Connection)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		if EmployeeOrganization != Organization{
			return conf.PrintError(conf.ErrUserNotFound, ResponseWriter)
		}
		Resp, APIerr := resetEmployeePassword_Request(login, Connection)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		Response, _ := json.Marshal(Resp)
		fmt.Fprintf(ResponseWriter, string(Response))
	}
	return true
}

func resetEmployeePassword_Request(login string, Connection *sql.DB) (ResetEmployeePassword_Success, *conf.ApiError){
	Password, Hash := orgset.GeneratePassword()
	Query, err := Connection.Prepare("UPDATE users SET password=? WHERE login=?")
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
	Query, err = Connection.Prepare("DELETE FROM sessions WHERE login=?")
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
