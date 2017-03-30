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

func GetUserData(Token string, ResponseWriter http.ResponseWriter, login string) bool{
	if checkData(Token, login, ResponseWriter) {
		if authorization.CheckToken(Token, ResponseWriter){
			Organization, err := getUserOrganizationByToken(Token)
			if err != nil {
				return conf.PrintError(err, ResponseWriter)
			}
			NewConnection = src.Connect_Custom(Organization)
			userData, err := getUserData_Request(login)
			if err != nil {
				return conf.PrintError(err, ResponseWriter)
			}
			userData.Organization = Organization
			Resp := Success_GetUserData{200, "success", userData}
			Response, _ := json.Marshal(Resp)
			fmt.Fprintf(ResponseWriter, string(Response))
		} else {
			return conf.PrintError(conf.ErrUserTokenIncorrect, ResponseWriter)
		}
	}
	return true
}

func getUserOrganizationByToken(Token string) (string, *conf.ApiError){
	Query, err := Connection.Query("SELECT login FROM sessions WHERE token=?", Token)
	if err!= nil{
		return "", conf.ErrDatabaseQueryFailed
	}
	Login, APIerr := getUserLoginFromQuery(Query)
	if APIerr != nil {
		return "", APIerr
	}
	Query, err = Connection.Query("SELECT organization FROM users WHERE login=?", Login)
	if err != nil {
		return "", conf.ErrDatabaseQueryFailed
	}
	Organization, APIerr := getUserOrganizationFromQuery(Query)
	if APIerr != nil {
		return "", APIerr
	}
	return Organization, nil
}

func getUserOrganizationFromQuery(rows *sql.Rows) (string, *conf.ApiError){
	var organization string
	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&organization)
		if err != nil {
			return "", conf.ErrDatabaseQueryFailed
		}
	}
	return organization, nil
}

func getUserData_Request(login string) (UserData, *conf.ApiError){
	Query, err := NewConnection.Query("SELECT name, surname, middlename, sex, access, avatar, team FROM users WHERE login=?", login)
	if err != nil {
		return UserData{}, conf.ErrDatabaseQueryFailed
	}
	var userData UserData
	userData, APIerr := getUserDataFromQuery(Query)
	if APIerr != nil {
		return UserData{}, APIerr
	}
	return userData, nil
}

func getUserDataFromQuery(rows *sql.Rows) (UserData, *conf.ApiError){
	defer rows.Close()
	var userData UserData
	for rows.Next(){
		err := rows.Scan(&userData.Name, &userData.Surname, &userData.Middlename, &userData.Sex, &userData.Access, &userData.Avatar, &userData.Team)
		if err != nil {
			return UserData{}, conf.ErrDatabaseQueryFailed
		}
	}
	return userData, nil
}
