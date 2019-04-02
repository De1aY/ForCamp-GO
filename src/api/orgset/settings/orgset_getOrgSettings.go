/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package settings

import (
	"net/http"
	"wplay/conf"
	"database/sql"
	"wplay/src/api/authorization"
	"wplay/src"
	"wplay/src/api/orgset"
)

type OrgSettings struct {
	Team string `json:"team"`
	Period string `json:"period"`
	Privacy string `json:"privacy"`
	SelfMarks string `json:"self_marks"`
	Participant string `json:"participant"`
	Organization string `json:"organization"`
	EmotionalMarkPeriod string `json:"emotional_mark_period"`
}

type getOrgSettings_Success struct {
	Settings OrgSettings `json:"settings"`
}

func GetOrgSettings(token string, responseWriter http.ResponseWriter) bool {
	if authorization.IsTokenNotEmpty(token, responseWriter){
		if authorization.IsTokenValid(token, responseWriter) {
			Organization, _, apiErr := orgset.GetUserOrganizationAndIdByToken(token)
			if apiErr != nil {
				return apiErr.Print(responseWriter)
			}
			src.CustomConnection = src.Connect_Custom(Organization)
			data, apiErr := GetOrgSettings_Request()
			if apiErr != nil {
				return apiErr.Print(responseWriter)
			}
			resp := conf.ApiResponse{200, "success", getOrgSettings_Success{data}}
			resp.Print(responseWriter)
			return true
		} else {
			return conf.ErrUserTokenIncorrect.Print(responseWriter)
		}
	}
	return false
}

func GetOrgSettings_Request() (OrgSettings, *conf.ApiResponse) {
	rows, err := src.CustomConnection.Query("SELECT * FROM settings")
	if err != nil {
		return OrgSettings{}, conf.ErrDatabaseQueryFailed
	}
	data, apiErr := getOrgSettingFromQuery(rows)
	if apiErr != nil {
		return OrgSettings{}, apiErr
	}
	return data, nil
}

func getOrgSettingFromQuery(rows *sql.Rows) (OrgSettings, *conf.ApiResponse) {
	OrgSettingsRaw := make(map[string]string)
	var key, value string
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&key, &value)
		if err != nil {
			return OrgSettings{}, conf.ErrDatabaseQueryFailed
		}
		OrgSettingsRaw[key] = value
	}
	OrgSettings := OrgSettings{
		Team:                OrgSettingsRaw["team"],
		Period:              OrgSettingsRaw["period"],
		Privacy:             OrgSettingsRaw["privacy"],
		SelfMarks:           OrgSettingsRaw["self_marks"],
		Participant:         OrgSettingsRaw["participant"],
		Organization:        OrgSettingsRaw["organization"],
		EmotionalMarkPeriod: OrgSettingsRaw["emotional_mark_period"],
	}
	return OrgSettings, nil
}
