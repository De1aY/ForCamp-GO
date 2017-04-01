package users

import (
	"net/http"
	"forcamp/conf"
)

type Success_GetUserLogin struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Login string `json:"login"`
}

type Success_GetUserData struct {
	Code int `json:"code"`
	Status string `json:"status"`
	UserData UserData `json:"data"`
}

type UserData struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
	Middlename string `json:"middlename"`
	Team string `json:"team"`
	Access int `json:"access"`
	Avatar string `json:"avatar"`
	Sex int `json:"sex"`
	Organization string `json:"organization"`
}

func checkData(token string, login string, w http.ResponseWriter) bool{
	if len(token) > 0{
		if len(login) > 0{
			return true
		} else {
			return conf.PrintError(conf.ErrUserLoginEmpty, w)
		}
	} else {
		return conf.PrintError(conf.ErrUserTokenEmpty, w)
	}
}