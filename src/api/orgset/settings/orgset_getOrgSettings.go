package settings

import (
	"net/http"
	"forcamp/conf"
	"database/sql"
	"forcamp/src/api/authorization"
	"forcamp/src"
	"forcamp/src/api/orgset"
)

type OrgSettings struct {
	Participant string `json:"participant"`
	Team string `json:"team"`
	Organization string `json:"organization"`
	Period string `json:"period"`
	SelfMarks string `json:"self_marks"`
}

type getOrgSettings_Success struct {
	Settings OrgSettings `json:"settings"`
}

func GetOrgSettings(token string, responseWriter http.ResponseWriter) bool {
	if authorization.CheckTokenForEmpty(token, responseWriter){
		if authorization.CheckToken(token, responseWriter) {
			Organization, _, apiErr := orgset.GetUserOrganizationAndLoginByToken(token)
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
	query, err := src.CustomConnection.Query("SELECT * FROM settings")
	if err != nil {
		return OrgSettings{}, conf.ErrDatabaseQueryFailed
	}
	data, APIerr := getOrgSettingFromQuery(query)
	if APIerr != nil {
		return OrgSettings{}, APIerr
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
	OrgSettings := OrgSettings{Organization: OrgSettingsRaw["organization"],
		Team: OrgSettingsRaw["team"],
		Participant: OrgSettingsRaw["participant"],
		Period: OrgSettingsRaw["period"],
		SelfMarks: OrgSettingsRaw["self_marks"]}
	return OrgSettings, nil
}