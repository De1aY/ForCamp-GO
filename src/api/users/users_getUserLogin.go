package users

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src/api/authorization"
	"forcamp/src"
)

func GetUserLogin(token string, responseWriter http.ResponseWriter) bool{
	if authorization.CheckTokenForEmpty(token, responseWriter) {
		if authorization.CheckToken(token, responseWriter) {
			login, apiErr := GetUserLogin_Request(token)
			if apiErr != nil {
				return apiErr.Print(responseWriter)
			}
			rawResp := getUserLogin_Success{login}
			resp := &conf.ApiResponse{200, "success", rawResp}
			resp.Print(responseWriter)
			return true
		} else {
			return conf.ErrUserTokenIncorrect.Print(responseWriter)
		}
	}
	return false
}

func GetUserLogin_Request (token string) (string, *conf.ApiResponse) {
	var login string
	err := src.Connection.QueryRow("SELECT login FROM sessions WHERE token=?", token).Scan(&login)
	if err != nil {
		return "", conf.ErrDatabaseQueryFailed
	}
	return login, nil
}
