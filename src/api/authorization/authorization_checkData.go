package authorization

import (
	"net/http"
	"forcamp/conf"
	"html"
	"forcamp/src"
)

func checkAuthorizationData(inf AuthInf, responseWriter http.ResponseWriter) bool {
	if checkUserLogin(inf.Login, responseWriter) && checkUserPassword(inf.Password, responseWriter) {
		if checkAuthorizationData_Request(inf, responseWriter){
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func checkUserLogin(login string, responseWriter http.ResponseWriter) bool {
	if len(login) > 0 {
		return true
	} else {
		return conf.ErrUserLoginEmpty.Print(responseWriter)
	}
}

func checkUserPassword(password string, responseWriter http.ResponseWriter) bool {
	if len(password) > 0{
		return true
	} else {
		return conf.ErrUserPasswordEmpty.Print(responseWriter)
	}
}

// select ID by Login and Password
func checkAuthorizationData_Request(authInf AuthInf, responseWriter http.ResponseWriter) bool {
	authInf.Login = html.EscapeString(authInf.Login)
	authInf.Password = GeneratePasswordHash(authInf.Password)
	Query, err := src.Connection.Query("SELECT COUNT(id) as count FROM users WHERE login=? AND password=?", authInf.Login, authInf.Password)
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(responseWriter)
	}
	if getCountVal(Query, responseWriter) > 0 {
		return true
	} else {
		return conf.ErrAuthDataIncorrect.Print(responseWriter)
	}
}