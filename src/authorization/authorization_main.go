package authorization

import (
	"forcamp/src"
	"forcamp/conf"
	"net/http"
	"database/sql"
	"log"
)


func Authorize(inf AuthInf, ResponseWriter http.ResponseWriter) {
	Connection := src.Connect()
	defer Connection.Close()
	if checkAuthorizationData(inf, Connection, ResponseWriter) {
		if checkCurrentSessionsVal(inf.Login, Connection, ResponseWriter) {
			setUserToken(inf.Login, Connection, ResponseWriter)
		}
	}
}

func setUserToken(login string, Connection *sql.DB, ResponseWriter http.ResponseWriter) bool {
	Query, err := Connection.Prepare("INSERT INTO sessions (login, token) VALUES (?,?)")
	if err != nil {
		return conf.PrintError(conf.ErrDatabaseQueryFailed, ResponseWriter)
	}
	defer Query.Close()
	Token := getToken(login, Connection, ResponseWriter)
	_, err = Query.Exec(login, Token)
	if err != nil {
		return conf.PrintError(conf.ErrDatabaseQueryFailed, ResponseWriter)
	}
	return printToken(Token, ResponseWriter)
}

func getToken(login string, Connection *sql.DB, ResponseWriter http.ResponseWriter) string {
	for true {
		Token := generateTokenHash(login)
		if CheckToken(Token, Connection, ResponseWriter){
			continue
		} else {
			return Token
		}
	}
	return ""
}

// True - Token is exist, False - NO
func CheckToken(token string, Connection *sql.DB, ResponseWriter http.ResponseWriter) bool {
	Query, err := Connection.Query("SELECT COUNT(login) as count FROM sessions WHERE token=?", token)
	if err != nil {
		log.Print(err)
		return false
	}
	Count := getCountVal(Query, ResponseWriter)
	if Count != 0 {
		return true
	} else {
		return false
	}
}

func VerifyToken(token string, Connection *sql.DB, ResponseWriter http.ResponseWriter) bool{
	if len(token) > 0 {
		if CheckToken(token, Connection, ResponseWriter) {
			return conf.PrintSuccess(conf.RequestSuccess, ResponseWriter)
		} else {
			return conf.PrintError(conf.ErrUserTokenIncorrect, ResponseWriter)
		}
	} else {
		return conf.PrintError(conf.ErrUserTokenEmpty, ResponseWriter)
	}
}

func checkCurrentSessionsVal(login string, Connection *sql.DB, ResponseWriter http.ResponseWriter) bool {
	Query, err := Connection.Query("SELECT COUNT(token) as count FROM sessions WHERE login=?", login)
	if err != nil {
		return conf.PrintError(conf.ErrDatabaseQueryFailed, ResponseWriter)
	}
	if getCountVal(Query, ResponseWriter) > 4 {
		return deleteOldestSession(login, Connection, ResponseWriter)
	} else {
		return true
	}
}

func deleteOldestSession(login string, Connection *sql.DB, ResponseWriter http.ResponseWriter) bool {
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