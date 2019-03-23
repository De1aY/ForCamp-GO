package authorization

import (
	"net/http"
	"wplay/conf"
	"html"
	"wplay/src"
)

func checkAuthorizationData(inf AuthInf, responseWriter http.ResponseWriter) bool {
	if checkUserLogin(inf.Login, responseWriter) && checkUserPassword(inf.Password, responseWriter) {
		if isAuthorizationDataCorrect(inf, responseWriter){
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
func isAuthorizationDataCorrect(authInf AuthInf, responseWriter http.ResponseWriter) bool {
	authInf.Login = html.EscapeString(authInf.Login)
	authInf.Password = GeneratePasswordHash(authInf.Password)
	var count int
	err := src.Connection.QueryRow("SELECT COUNT(id) as count FROM users WHERE " +
		"login=? AND password=?", authInf.Login, authInf.Password).Scan(&count)
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(responseWriter)
	}
	if count > 0 {
		return true
	} else {
		return conf.ErrAuthDataIncorrect.Print(responseWriter)
	}
}