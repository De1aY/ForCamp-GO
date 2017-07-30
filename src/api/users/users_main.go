package users

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src/api/marks"
	"forcamp/src/api/orgset/teams"
)

type getUserLogin_Success struct {
	Login  string `json:"login"`
}

type getUserData_Success struct {
	Data UserData `json:"data"`
}

type UserData struct {
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Middlename   string `json:"middlename"`
	Team         teams.Team `json:"team"`
	Access       int `json:"access"`
	Avatar       string `json:"avatar"`
	Sex          int `json:"sex"`
	Organization string `json:"organization"`
	Post         string `json:"post"`
	Actions      []marks.MarksChange `json:"actions"`
	AdditionalData interface{} `json:"additional_data"`
}

func checkData(token string, login string, w http.ResponseWriter) bool {
	if len(token) > 0 {
		if len(login) > 0 {
			return true
		} else {
			return conf.ErrUserLoginEmpty.Print(w)
		}
	} else {
		return conf.ErrUserTokenEmpty.Print(w)
	}
}