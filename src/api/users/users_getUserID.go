package users

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src/api/authorization"
	"forcamp/src"
	"strings"
	"strconv"
)

func GetUserID(token string, responseWriter http.ResponseWriter) bool{
	if authorization.IsTokenNotEmpty(token, responseWriter) {
		if authorization.IsTokenValid(token, responseWriter) {
			id, apiErr := GetUserID_Request(token)
			if apiErr != nil {
				return apiErr.Print(responseWriter)
			}
			rawResp := getUserLogin_Success{id}
			resp := &conf.ApiResponse{200, "success", rawResp}
			resp.Print(responseWriter)
			return true
		} else {
			return conf.ErrUserTokenIncorrect.Print(responseWriter)
		}
	}
	return false
}

func GetUserID_Request(token string) (int64, *conf.ApiResponse) {
	var id int64
	var login string
	err := src.Connection.QueryRow("SELECT login FROM sessions WHERE token=?", token).Scan(&login)
	if err != nil {
		return id, conf.ErrDatabaseQueryFailed
	}
	loginData := strings.Split(login, "_")
	id, err = strconv.ParseInt(loginData[1], 10, 64); if err != nil {
		return id, conf.ErrDatabaseQueryFailed
	}
	return id, nil
}
