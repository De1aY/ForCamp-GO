package authorization

import (
	"forcamp/src"
	"forcamp/conf"
	"net/http"
)

// Connection to Database
var Connection = src.Connect()

func Authorize(inf AuthInf, ResponseWriter http.ResponseWriter) {
	if checkAuthorizationData(inf, ResponseWriter) {
		if checkCurrentSessionsVal(inf.Login, ResponseWriter) {
			setUserToken(inf.Login, ResponseWriter)
		}
	}
}

func setUserToken(login string, ResponseWriter http.ResponseWriter) bool {
	Query, err := Connection.Prepare("INSERT INTO sessions (login, token) VALUES (?,?)")
	if err != nil {
		return conf.PrintError(conf.ErrDatabaseQueryFailed, ResponseWriter)
	}
	defer Query.Close()
	Token := getToken(login, ResponseWriter)
	_, err = Query.Exec(login, Token)
	if err != nil {
		return conf.PrintError(conf.ErrDatabaseQueryFailed, ResponseWriter)
	}
	return printToken(Token, ResponseWriter)
}

func getToken(login string, ResponseWriter http.ResponseWriter) string {
	for true {
		Token := generateTokenHash(login)
		if CheckToken(Token, ResponseWriter){
			continue
		} else {
			return Token
		}
	}
	return ""
}

// True - Token is exist, False - NO
func CheckToken(token string, ResponseWriter http.ResponseWriter) bool {
	Query, err := Connection.Query("SELECT COUNT(login) as count FROM sessions WHERE token=?", token)
	if err != nil {
		return conf.PrintError(conf.ErrDatabaseQueryFailed, ResponseWriter)
	}
	Count := getCountVal(Query, ResponseWriter)
	if Count != 0 {
		return true
	} else {
		return false
	}
}

func VerifyToken(token string, ResponseWriter http.ResponseWriter) bool{
	if len(token) > 0 {
		if CheckToken(token, ResponseWriter) {
			return conf.PrintSuccess(conf.RequestSuccess, ResponseWriter)
		} else {
			return conf.PrintError(conf.ErrUserTokenIncorrect, ResponseWriter)
		}
	} else {
		return conf.PrintError(conf.ErrUserTokenEmpty, ResponseWriter)
	}
}

func checkCurrentSessionsVal(login string, ResponseWriter http.ResponseWriter) bool {
	Query, err := Connection.Query("SELECT COUNT(token) as count FROM sessions WHERE login=?", login)
	if err != nil {
		return conf.PrintError(conf.ErrDatabaseQueryFailed, ResponseWriter)
	}
	if getCountVal(Query, ResponseWriter) > 4 {
		return deleteOldestSession(login, ResponseWriter)
	} else {
		return true
	}
}

func deleteOldestSession(login string, ResponseWriter http.ResponseWriter) bool {
	Query, err := Connection.Prepare("DELETE FROM sessions WHERE login=? LIMIT 1")
	if err != nil {
		return conf.PrintError(conf.ErrDatabaseQueryFailed, ResponseWriter)
	}
	defer Query.Close()
	_, err = Query.Exec(login)
	if err != nil {
		return conf.PrintError(conf.ErrDatabaseQueryFailed, ResponseWriter)
	} else {
		return true
	}
}