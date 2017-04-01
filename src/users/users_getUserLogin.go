package users

import (
	"net/http"
	"forcamp/conf"
	"encoding/json"
	"fmt"
	"database/sql"
	"forcamp/src/authorization"
	"forcamp/src"
)

func GetUserLogin(Token string, ResponseWriter http.ResponseWriter) bool{
	if authorization.CheckTokenForEmpty(Token, ResponseWriter) {
		if authorization.CheckToken(Token, ResponseWriter) {
			Query, err := src.Connection.Query("SELECT login FROM sessions WHERE token=?", Token)
			if err != nil {
				return conf.PrintError(conf.ErrDatabaseQueryFailed, ResponseWriter)
			}
			Login, APIerr := getUserLoginFromQuery(Query)
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			Resp := Success_GetUserLogin{200, "success", Login}
			Response, _ := json.Marshal(Resp)
			fmt.Fprintf(ResponseWriter, string(Response))
			return true
		} else {
			return conf.PrintError(conf.ErrUserTokenIncorrect, ResponseWriter)
		}
	}
	return false
}

func getUserLoginFromQuery(rows *sql.Rows) (string, *conf.ApiError){
	var login string
	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&login)
		if err != nil {
			return "", conf.ErrDatabaseQueryFailed
		}
	}
	return login, nil
}