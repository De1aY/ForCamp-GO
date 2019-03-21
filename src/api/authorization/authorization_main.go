package authorization

import (
	"encoding/json"
	"forcamp/conf"
	"forcamp/src"
	"net/http"
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
	query, err := src.Connection.Prepare("INSERT INTO sessions (login, token) VALUES (?,?)")
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(responseWriter)
	}
	defer query.Close()
	token := getToken(login, responseWriter)
	_, err = query.Exec(login, token)
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(responseWriter)
	}
	return printToken(token, responseWriter)
}

func getToken(login string, responseWriter http.ResponseWriter) string {
	for true {
		Token := generateTokenHash(login)
		if IsTokenValid(Token, responseWriter) {
			continue
		} else {
			return Token
		}
	}
	return ""
}

// True - Token is exist, False - NO
func IsTokenValid(token string, responseWriter http.ResponseWriter) bool {
	var count int
	err := src.Connection.QueryRow("SELECT COUNT(login) as count FROM sessions WHERE token=?", token).Scan(&count)
	if err != nil {
		return false
	}
	return count != 0
}

func VerifyToken(token string, responseWriter http.ResponseWriter) bool {
	if len(token) > 0 {
		if IsTokenValid(token, responseWriter) {
			adminStatus, apiErr := isUserAdmin(token)
			if apiErr != nil {
				return apiErr.Print(responseWriter)
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

func isUserAdmin(token string) (bool, *conf.ApiResponse) {
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
	var count int
	err := src.Connection.QueryRow("SELECT COUNT(token) as count FROM sessions WHERE login=?", login).Scan(&count)
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(responseWriter)
	}
	if count > 4 {
		return deleteOldestSession(login, responseWriter)
	} else {
		return true
	}
}

func deleteOldestSession(login string, responseWriter http.ResponseWriter) bool {
	query, err := src.Connection.Prepare("DELETE FROM sessions WHERE login=? LIMIT 1"); if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(responseWriter)
	}
	defer query.Close()
	_, err = query.Exec(login)
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(responseWriter)
	} else {
		return true
	}
}
