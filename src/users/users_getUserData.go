package users

import (
	"forcamp/src/authorization"
	"forcamp/src"
	"encoding/json"
	"fmt"
	"net/http"
	"forcamp/conf"
	"database/sql"
)

func GetUserData(Token string, ResponseWriter http.ResponseWriter, login string){
	if checkData(Token, login, ResponseWriter) && authorization.CheckToken(Token, ResponseWriter){
		NewConnection = src.Connect_Custom(getUserOrganizationByToken(Token, ResponseWriter))
		userData := getUserData_Request(login, ResponseWriter)
		Resp := Success_GetUserData{200, "success", userData}
		Response, _ := json.Marshal(Resp)
		fmt.Fprintf(ResponseWriter, string(Response))
	}
}

func getUserOrganizationByToken(Token string, ResponseWriter http.ResponseWriter) string{
	Query, err := Connection.Query("SELECT login FROM sessions WHERE token=?", Token)
	if err!= nil{
		conf.PrintError(conf.ErrDatabaseQueryFailed, ResponseWriter)
	}
	Login := getUserLoginFromQuery(Query, ResponseWriter)
	Query, err = Connection.Query("SELECT organization FROM users WHERE login=?", Login)
	if err != nil {
		conf.PrintError(conf.ErrDatabaseQueryFailed, ResponseWriter)
	}
	return getUserOrganizationFromQuery(Query, ResponseWriter)
}

func getUserOrganizationFromQuery(rows *sql.Rows, ResponseWriter http.ResponseWriter) (organization string){
	for rows.Next(){
		defer rows.Close()
		err := rows.Scan(&organization)
		if err != nil {
			conf.PrintError(conf.ErrDatabaseQueryFailed, ResponseWriter)
		}
	}
	return organization
}

func getUserData_Request(login string, ResponseWriter http.ResponseWriter) (userData UserData){
	Query, err := NewConnection.Query("SELECT name, surname, middlename, sex, access, avatar, team FROM users WHERE login=?", login)
	if err != nil {
		conf.PrintError(conf.ErrDatabaseQueryFailed, ResponseWriter)
	}
	userData = getUserDataFromQuery(Query, ResponseWriter)
	return userData
}

func getUserDataFromQuery(rows *sql.Rows, ResponseWriter http.ResponseWriter) (userData UserData){
	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&userData.Name, &userData.Surname, &userData.Middlename, &userData.Sex, &userData.Access, &userData.Avatar, &userData.Team)
		if err != nil {
			conf.PrintError(conf.ErrDatabaseQueryFailed, ResponseWriter)
		}
	}
	return userData
}
