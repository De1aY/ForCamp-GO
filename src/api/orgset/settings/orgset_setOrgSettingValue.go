package settings

import (
	"net/http"
	"forcamp/conf"
	"log"
	"forcamp/src/api/orgset"
	"forcamp/src"
)

func SetOrgSettingValue(token string, name string, value string, responseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token, responseWriter){
		if checkSetOrgSettingValueData(token, name, value, responseWriter){
			Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
			if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			src.CustomConnection = src.Connect_Custom(Organization)
			if setOrgSettingValue_Request(name, value, responseWriter){
				conf.RequestSuccess.Print(responseWriter)
			}
		}
	}
	return true
}

func checkSetOrgSettingValueData(token string, name string, value string, w http.ResponseWriter) bool{
	if len(token) == 0 {
		return conf.ErrUserTokenEmpty.Print(w)
	}
	if len(name) == 0 {
		return conf.ErrOrgSettingNameEmpty.Print(w)
	}
	if len(value) == 0{
		return conf.ErrOrgSettingValueEmpty.Print(w)
	}
	return true
}

func setOrgSettingValue_Request(name string, value string, w http.ResponseWriter) bool{
	Query, err := src.CustomConnection.Prepare("UPDATE settings SET value=? WHERE name=?")
	if err != nil{
		log.Print(err)
		return conf.ErrDatabaseQueryFailed.Print(w)
	}
	defer Query.Close()
	_, err = Query.Exec(value, name)
	if err != nil {
		log.Print(err)
		return conf.ErrOrgSettingNameIncorrect.Print(w)
	}
	return true
}