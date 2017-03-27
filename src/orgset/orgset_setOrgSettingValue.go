package orgset

import (
	"net/http"
	"forcamp/conf"
)

func SetOrgSettingValue(token string, name string, value string, ResponseWriter http.ResponseWriter){
	if checkUserAccess(token, ResponseWriter){
		if checkSetOrgSettingValueData(token, name, value, ResponseWriter){
			if setOrgSettingValue_Request(name, value, ResponseWriter){
				conf.PrintSuccess(conf.RequestSuccess, ResponseWriter)
			}
		}
	} else {
		conf.PrintError(conf.ErrUserTokenIncorrect, ResponseWriter)
	}
}

func checkSetOrgSettingValueData(token string, name string, value string, w http.ResponseWriter) bool{
	if len(token) == 0 {
		return conf.PrintError(conf.ErrUserTokenEmpty, w)
	}
	if len(name) == 0 {
		return conf.PrintError(conf.ErrOrgSettingNameEmpty, w)
	}
	if len(value) == 0{
		return conf.PrintError(conf.ErrOrgSettingValueEmpty, w)
	}
	return true
}

func setOrgSettingValue_Request(name string, value string, w http.ResponseWriter) bool{
	Query, err := NewConnection.Prepare("UPDATE settings SET value=? WHERE name=?")
	if err != nil{
		return conf.PrintError(conf.ErrDatabaseQueryFailed, w)
	}
	defer Query.Close()
	_, err = Query.Exec(value, name)
	if err != nil {
		return conf.PrintError(conf.ErrOrgSettingNameIncorrect, w)
	}
	return true
}