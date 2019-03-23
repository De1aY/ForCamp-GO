package settings

import (
	"net/http"
	"wplay/conf"
	"wplay/src/api/orgset"
	"wplay/src"
	"strconv"
	"strings"
)

func SetOrgSettingValue(token string, setting_name string, setting_value string, responseWriter http.ResponseWriter) bool {
	if orgset.IsUserAdmin(token, responseWriter){
		organizationName, _, apiErr := orgset.GetUserOrganizationAndIdByToken(token); if apiErr != nil {
			return apiErr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(organizationName)
		response := setOrgSettingValue(setting_name, setting_value)
		response.Print(responseWriter)
	}
	return true
}

func setOrgSettingValue(setting_name string, setting_value string) *conf.ApiResponse {
	if len(setting_name) < 1 {
		return conf.ErrOrgSettingNameIncorrect
	}
	if len(setting_value) < 1 {
		return conf.ErrOrgSettingValueIncorrect
	}
	switch setting_name {
		case "participant":
			return setOrgSettingValue_Participant(setting_value)
		case "team":
			return setOrgSettingValue_Team(setting_value)
	    case "organization":
			return setOrgSettingValue_Organization(setting_value)
	    case "period":
			return setOrgSettingValue_Period(setting_value)
		case "self_marks":
			return setOrgSettingValue_SelfMarks(setting_value)
		case "emotional_mark_period":
			return setOrgSettingValue_EmotionalMarkPeriod(setting_value)
		default:
			return conf.ErrOrgSettingNameIncorrect
	}
}

func setOrgSettingValue_Participant(setting_value string) *conf.ApiResponse {
	query, err := src.CustomConnection.Prepare("UPDATE settings SET value=? WHERE name='participant'")
	if err != nil{
		return conf.ErrDatabaseQueryFailed
	}
	defer query.Close()
	_, err = query.Exec(setting_value); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return conf.RequestSuccess
}

func setOrgSettingValue_Team(setting_value string) *conf.ApiResponse {
	query, err := src.CustomConnection.Prepare("UPDATE settings SET value=? WHERE name='team'")
	if err != nil{
		return conf.ErrDatabaseQueryFailed
	}
	defer query.Close()
	_, err = query.Exec(setting_value); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return conf.RequestSuccess
}

func setOrgSettingValue_Organization(setting_value string) *conf.ApiResponse {
	query, err := src.CustomConnection.Prepare("UPDATE settings SET value=? WHERE name='organization'")
	if err != nil{
		return conf.ErrDatabaseQueryFailed
	}
	defer query.Close()
	_, err = query.Exec(setting_value); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return conf.RequestSuccess
}

func setOrgSettingValue_Period(setting_value string) *conf.ApiResponse {
	query, err := src.CustomConnection.Prepare("UPDATE settings SET value=? WHERE name='period'")
	if err != nil{
		return conf.ErrDatabaseQueryFailed
	}
	defer query.Close()
	_, err = query.Exec(setting_value); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return conf.RequestSuccess
}

func setOrgSettingValue_SelfMarks(setting_value string) *conf.ApiResponse {
	setting_value = strings.ToLower(setting_value)
	if setting_value != "true" && setting_value != "false" {
		return conf.ErrOrgSettingValueIncorrect
	}
	query, err := src.CustomConnection.Prepare("UPDATE settings SET value=? WHERE name='self_marks'")
	if err != nil{
		return conf.ErrDatabaseQueryFailed
	}
	defer query.Close()
	_, err = query.Exec(setting_value); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return conf.RequestSuccess
}

func setOrgSettingValue_EmotionalMarkPeriod(setting_value string) *conf.ApiResponse {
	emotionalMarkPeriod, err := strconv.ParseInt(setting_value, 10, 64); if err != nil {
		return conf.ErrOrgSettingValueIncorrect
	}
	if emotionalMarkPeriod < 1 || emotionalMarkPeriod > 24 {
		return conf.ErrOrgSettingValueIncorrect
	}
	query, err := src.CustomConnection.Prepare("UPDATE settings SET value=? WHERE name='emotional_mark_period'")
	if err != nil{
		return conf.ErrDatabaseQueryFailed
	}
	defer query.Close()
	_, err = query.Exec(setting_value); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return conf.RequestSuccess
}