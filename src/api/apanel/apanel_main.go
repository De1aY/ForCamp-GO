package apanel

import (
	"forcamp/conf"
	"forcamp/src"
)

type createOrganization_Success struct {
	AdminLogin string `json:"admin_login"`
	AdminPassword string `json:"admin_password"`
}

func getLoginByToken(token string) (string, *conf.ApiResponse) {
	if len(token) == 0{
		return "", conf.ErrUserTokenEmpty
	}
	var login string
	err := src.Connection.QueryRow("SELECT login FROM sessions WHERE token=?", token).Scan(&login)
	if err != nil {
		return "", conf.ErrDatabaseQueryFailed
	}
	if len(login) == 0 {
		return "", conf.ErrUserTokenIncorrect
	}
	return login, nil
}

func checkPermissions(token string) *conf.ApiResponse{
	login, APIerr := getLoginByToken(token)
	if APIerr != nil {
		return APIerr
	}
	var admin_status bool
	err := src.Connection.QueryRow("SELECT admin FROM users WHERE login=?", login).Scan(&admin_status)
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	if admin_status {
		return nil
	} else {
		return conf.ErrInsufficientRights
	}
}