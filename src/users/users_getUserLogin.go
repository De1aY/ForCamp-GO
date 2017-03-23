package users

import (
	"net/http"
	"forcamp/conf"
	"encoding/json"
	"fmt"
	"database/sql"
)

func GetUserLogin(Token string, ResponseWriter http.ResponseWriter) bool{
	Query, err := Connection.Query("SELECT login FROM sessions WHERE token=?", Token)
	if err!= nil{
		return conf.PrintError(conf.ErrDatabaseQueryFailed, ResponseWriter)
	}
	Login := getUserLoginFromQuery(Query, ResponseWriter)
	Resp := Success_GetUserLogin{200, "success", Login}
	Response, _ := json.Marshal(Resp)
	fmt.Fprintf(ResponseWriter, string(Response))
	return true
}

func getUserLoginFromQuery(rows *sql.Rows, ResponseWriter http.ResponseWriter) (login string){
	for rows.Next(){
		defer rows.Close()
		err := rows.Scan(&login)
		if err != nil {
			conf.PrintError(conf.ErrDatabaseQueryFailed, ResponseWriter)
		}
	}
	return login
}