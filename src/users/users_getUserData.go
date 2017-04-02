package users

import (
	"forcamp/src/authorization"
	"forcamp/src"
	"encoding/json"
	"fmt"
	"net/http"
	"forcamp/conf"
	"database/sql"
	"forcamp/src/orgset"
)

func GetUserData(Token string, ResponseWriter http.ResponseWriter, login string) bool{
	if checkData(Token, login, ResponseWriter) {
		Connection := src.Connect()
		defer Connection.Close()
		if authorization.CheckToken(Token,Connection, ResponseWriter){
			Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(Token, Connection)
			if APIerr != nil{
				return conf.PrintError(APIerr, ResponseWriter)
			}
			NewConnection := src.Connect_Custom(Organization)
			defer NewConnection.Close()
			userData, err := getUserData_Request(login, NewConnection)
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

func getUserData_Request(login string, Connection *sql.DB) (UserData, *conf.ApiError){
	Query, err := Connection.Query("SELECT name, surname, middlename, sex, access, avatar, team FROM users WHERE login=?", login)
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
