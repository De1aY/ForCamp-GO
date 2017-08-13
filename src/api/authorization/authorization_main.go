package authorization

import (
	"forcamp/src"
	"forcamp/conf"
	"net/http"
	"encoding/json"
)

type checkToken_Success struct {
	AdminStatus bool `json:"admin_status"`
}

// Function converts struct to JSON
func (success *checkToken_Success) toJSON() string {
	resp, _ := json.Marshal(success)
	return string(resp)
}

func Authorize(inf AuthInf, responseWriter http.ResponseWriter) {
	if checkAuthorizationData(inf, responseWriter) {
		if checkCurrentSessionsVal(inf.Login, responseWriter) {
			setUserToken(inf.Login, responseWriter)
		}
	}
}

func setUserToken(login string, responseWriter http.ResponseWriter) bool {
	Query, err := src.Connection.Prepare("INSERT INTO sessions (login, token) VALUES (?,?)")
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(responseWriter)
	}
	defer Query.Close()
	Token := getToken(login, responseWriter)
	_, err = Query.Exec(login, Token)
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(responseWriter)
	}
	return printToken(Token, responseWriter)
}

func getToken(login string, responseWriter http.ResponseWriter) string {
	for true {
		Token := generateTokenHash(login)
		if CheckToken(Token, responseWriter){
			continue
		} else {
			return Token
		}
	}
	return ""
}

// True - Token is exist, False - NO
func CheckToken(token string, responseWriter http.ResponseWriter) bool {
	Query, err := src.Connection.Query("SELECT COUNT(login) as count FROM sessions WHERE token=?", token)
	if err != nil {
		return false
	}
	Count := getCountVal(Query, responseWriter)
	if Count != 0 {
		return true
	} else {
		return false
	}
}

func VerifyToken(token string, responseWriter http.ResponseWriter) bool{
	if len(token) > 0 {
		if CheckToken(token, responseWriter) {
			adminStatus, APIerr := checkAdminStatus(token)
			if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			resp := &conf.ApiResponse{200, "success", checkToken_Success{adminStatus}}
			resp.Print(responseWriter)
			return true
		} else {
			return conf.ErrUserTokenIncorrect.Print(responseWriter)
		}
	} else {
		return conf.ErrUserTokenEmpty.Print(responseWriter)
	}
}

func checkAdminStatus(token string) (bool, *conf.ApiResponse) {
	var login string
	err := src.Connection.QueryRow("SELECT login FROM sessions WHERE token=?", token).Scan(&login)
	if err != nil {
		return false, conf.ErrDatabaseQueryFailed
	}
	var adminStatus bool
	err = src.Connection.QueryRow("SELECT admin FROM users WHERE login=?", login).Scan(&adminStatus)
	if err != nil {
		return false, conf.ErrDatabaseQueryFailed
	}
	return adminStatus, nil
}

func checkCurrentSessionsVal(login string, responseWriter http.ResponseWriter) bool {
	Query, err := src.Connection.Query("SELECT COUNT(token) as count FROM sessions WHERE login=?", login)
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(responseWriter)
	}
	if getCountVal(Query, responseWriter) > 4 {
		return deleteOldestSession(login, responseWriter)
	} else {
		return true
	}
}

func deleteOldestSession(login string, responseWriter http.ResponseWriter) bool {
	Query, err := src.Connection.Prepare("DELETE FROM sessions WHERE login=? LIMIT 1")
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(responseWriter)
	}
	defer Query.Close()
	_, err = Query.Exec(login)
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(responseWriter)
	} else {
		return true
	}
}