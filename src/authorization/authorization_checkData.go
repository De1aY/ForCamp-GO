package authorization

import (
	"net/http"
	"forcamp/conf"
	"html"
	"database/sql"
)

func checkAuthorizationData(inf AuthInf, Connection *sql.DB, ResponseWriter http.ResponseWriter) bool {
	if checkUserLogin(inf.Login, ResponseWriter) && checkUserPassword(inf.Password, ResponseWriter) {
		if checkAuthorizationData_Request(inf, Connection, ResponseWriter){
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func checkUserLogin(login string, ResponseWriter http.ResponseWriter) bool {
	if len(login) > 0 {
		return true
	} else {
		return conf.PrintError(conf.ErrUserLoginEmpty, ResponseWriter)
	}
}

func checkUserPassword(password string, ResponseWriter http.ResponseWriter) bool {
	if len(password) > 0{
		return true
	} else {
		return conf.PrintError(conf.ErrUserPasswordEmpty, ResponseWriter)
	}
}

// select ID by Login and Password
func checkAuthorizationData_Request(authInf AuthInf, Connection *sql.DB, ResponseWriter http.ResponseWriter) bool {
	authInf.Login = html.EscapeString(authInf.Login)
	authInf.Password = GeneratePasswordHash(authInf.Password)
	Query, err := Connection.Query("SELECT COUNT(id) as count FROM users WHERE login=? AND password=?", authInf.Login, authInf.Password)
	if err != nil {
		return conf.PrintError(conf.ErrDatabaseQueryFailed, ResponseWriter)
	}
	if getCountVal(Query, ResponseWriter) > 0 {
		return true
	} else {
		return conf.PrintError(conf.ErrAuthDataIncorrect, ResponseWriter)
	}
}