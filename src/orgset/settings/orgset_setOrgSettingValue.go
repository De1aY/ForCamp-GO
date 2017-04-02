package settings

import (
	"net/http"
	"forcamp/conf"
	"log"
	"forcamp/src/orgset"
	"forcamp/src"
	"database/sql"
)

func SetOrgSettingValue(token string, name string, value string, ResponseWriter http.ResponseWriter) bool{
	Connection := src.Connect()
	defer Connection.Close()
	if orgset.CheckUserAccess(token, Connection, ResponseWriter){
		if checkSetOrgSettingValueData(token, name, value, ResponseWriter){
			Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token, Connection)
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			NewConnection := src.Connect_Custom(Organization)
			defer NewConnection.Close()
			if setOrgSettingValue_Request(name, value, ResponseWriter, NewConnection){
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

func setOrgSettingValue_Request(name string, value string, w http.ResponseWriter, Connection *sql.DB) bool{
	Query, err := Connection.Prepare("UPDATE settings SET value=? WHERE name=?")
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