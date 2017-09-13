package users

import (
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/api/authorization"
	"forcamp/src/api/orgset"
	"net/http"
)

func ChangeUserPassword(token string, oldPassword string, newPassword string,
	responseWriter http.ResponseWriter) bool {
	if authorization.IsTokenValid(token, responseWriter) {
		user_id, apiErr := orgset.GetUserIdByToken(token)
		if apiErr != nil {
			return apiErr.Print(responseWriter)
		}
		response := changeUserPassword(user_id, token, oldPassword, newPassword)
		return response.Print(responseWriter)
	}
	return true
}

func changeUserPassword(user_id int64, token string, oldPassword string, newPassword string) *conf.ApiResponse {
	oldPasswordHash := authorization.GeneratePasswordHash(oldPassword)
	var currentPasswordHash string
	err := src.Connection.QueryRow("SELECT password FROM users WHERE id=?", user_id).Scan(&currentPasswordHash)
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	if oldPasswordHash != currentPasswordHash {
		return conf.ErrCurrentPasswordIsWrong
	}
	request, err := src.Connection.Prepare("UPDATE users SET password=? WHERE id=?")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	newPasswordHash := authorization.GeneratePasswordHash(newPassword)
	_, err = request.Exec(&newPasswordHash, &user_id)
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	user_login, apiErr := orgset.GetUserLoginByID(user_id)
	if apiErr != nil {
		return conf.RequestSuccess
	}
	request, err = src.Connection.Prepare("DELETE FROM sessions WHERE login=? AND token!=?")
	if err != nil {
		return conf.RequestSuccess
	}
	_, err = request.Exec(&user_login, &token)
	return conf.RequestSuccess
}
