package settings

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src/api/orgset"
	"forcamp/src"
)

func SetOrgSettingsValue(token string, orgSet OrgSettings, responseWriter http.ResponseWriter) bool {
	if orgset.IsUserAdmin(token, responseWriter){
		if checkSetOrgSettingsValueData(token, orgSet, responseWriter){
			Organization, _, APIerr := orgset.GetUserOrganizationAndIdByToken(token)
			if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			src.CustomConnection = src.Connect_Custom(Organization)
			if setOrgSettingValue_Request(orgSet, responseWriter){
				conf.RequestSuccess.Print(responseWriter)
			}
		}
	}
	return true
}

func checkSetOrgSettingsValueData(token string, orgSet OrgSettings, w http.ResponseWriter) bool {
	if len(token) == 0 {
		return conf.ErrUserTokenEmpty.Print(w)
	}
	if len(orgSet.Team) == 0 || len(orgSet.Organization) == 0 || len(orgSet.Period) == 0 || len(orgSet.Participant) == 0 {
		return conf.ErrOrgSettingValueEmpty.Print(w)
	}
	if orgSet.SelfMarks != "true" && orgSet.SelfMarks != "false" {
		return conf.ErrSelfMarksIncorrect.Print(w)
	}
	return true
}

func setOrgSettingValue_Request(orgSet OrgSettings, w http.ResponseWriter) bool {
	if setOrgSettingValue_Request_Team(orgSet, w) &&
		setOrgSettingValue_Request_Participant(orgSet, w) &&
		setOrgSettingValue_Request_Organization(orgSet, w) &&
		setOrgSettingValue_Request_Period(orgSet, w) &&
		setOrgSettingValue_Request_SelfMarks(orgSet, w) {
		return true
	} else {
		return false
	}
}

func setOrgSettingValue_Request_Team(orgSet OrgSettings, w http.ResponseWriter) bool {
	Query, err := src.CustomConnection.Prepare("UPDATE settings SET value=? WHERE name='team'")
	if err != nil{
		return conf.ErrDatabaseQueryFailed.Print(w)
	}
	defer Query.Close()
	_, err = Query.Exec(orgSet.Team)
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(w)
	}
	return true
}

func setOrgSettingValue_Request_Participant(orgSet OrgSettings, w http.ResponseWriter) bool {
	Query, err := src.CustomConnection.Prepare("UPDATE settings SET value=? WHERE name='participant'")
	if err != nil{
		return conf.ErrDatabaseQueryFailed.Print(w)
	}
	defer Query.Close()
	_, err = Query.Exec(orgSet.Participant)
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(w)
	}
	return true
}

func setOrgSettingValue_Request_Period(orgSet OrgSettings, w http.ResponseWriter) bool {
	Query, err := src.CustomConnection.Prepare("UPDATE settings SET value=? WHERE name='period'")
	if err != nil{
		return conf.ErrDatabaseQueryFailed.Print(w)
	}
	defer Query.Close()
	_, err = Query.Exec(orgSet.Period)
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(w)
	}
	return true
}

func setOrgSettingValue_Request_Organization(orgSet OrgSettings, w http.ResponseWriter) bool {
	Query, err := src.CustomConnection.Prepare("UPDATE settings SET value=? WHERE name='organization'")
	if err != nil{
		return conf.ErrDatabaseQueryFailed.Print(w)
	}
	defer Query.Close()
	_, err = Query.Exec(orgSet.Organization)
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(w)
	}
	return true
}

func setOrgSettingValue_Request_SelfMarks(orgSet OrgSettings, w http.ResponseWriter) bool {
	Query, err := src.CustomConnection.Prepare("UPDATE settings SET value=? WHERE name='self_marks'")
	if err != nil{
		return conf.ErrDatabaseQueryFailed.Print(w)
	}
	defer Query.Close()
	_, err = Query.Exec(orgSet.SelfMarks)
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(w)
	}
	return true
}