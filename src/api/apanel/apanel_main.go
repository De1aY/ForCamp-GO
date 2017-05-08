package apanel

import (
	"encoding/json"
	"forcamp/conf"
	"forcamp/src"
	"log"
)

type createOrganization_Success struct {
	Code int `json:"code"`
	Status string `json:"status"`
	AdminLogin string `json:"admin_login"`
	AdminPassword string `json:"admin_password"`
}

func (success *createOrganization_Success) toJSON() string {
	resp, _ := json.Marshal(success)
	return string(resp)
}

func getLoginByToken(token string) (string, *conf.ApiError) {
	if len(token) == 0{
		return "", conf.ErrUserTokenEmpty
	}
	var login string
	err := src.Connection.QueryRow("SELECT login FROM sessions WHERE token=?", token).Scan(&login)
	if err != nil {
		log.Print(err)
		return "", conf.ErrDatabaseQueryFailed
	}
	if len(login) == 0 {
		return "", conf.ErrUserTokenIncorrect
	}
	return login, nil
}

func checkPermissions(token string) *conf.ApiError{
	login, APIerr := getLoginByToken(token)
	if APIerr != nil {
		return APIerr
	}
	var admin_status bool
	err := src.Connection.QueryRow("SELECT admin FROM users WHERE login=?", login).Scan(&admin_status)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	if admin_status {
		return nil
	} else {
		return conf.ErrInsufficientRights
	}
}