package settings

import (
	"net/http"
	"forcamp/conf"
	"log"
	"forcamp/src/orgset"
	"forcamp/src"
)

func SetOrgSettingValue(token string, name string, value string, ResponseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token, ResponseWriter){
		if checkSetOrgSettingValueData(token, name, value, ResponseWriter){
			Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			src.CustomConnection = src.Connect_Custom(Organization)
			if setOrgSettingValue_Request(name, value, ResponseWriter){
				conf.PrintSuccess(conf.RequestSuccess, ResponseWriter)
			}
		}
	}
	return true
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
	Query, err := src.CustomConnection.Prepare("UPDATE settings SET value=? WHERE name=?")
	if err != nil{
		log.Print(err)
		return conf.PrintError(conf.ErrDatabaseQueryFailed, w)
	}
	defer Query.Close()
	_, err = Query.Exec(value, name)
	if err != nil {
		log.Print(err)
		return conf.PrintError(conf.ErrOrgSettingNameIncorrect, w)
	}
	return true
}