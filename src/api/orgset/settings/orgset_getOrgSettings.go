package settings

import (
	"net/http"
	"forcamp/conf"
	"database/sql"
	"forcamp/src/api/authorization"
	"forcamp/src"
	"log"
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
			Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
			if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			src.CustomConnection = src.Connect_Custom(Organization)
			Query, err := src.CustomConnection.Query("SELECT * FROM settings")
			if err != nil {
				log.Print(err)
				return conf.ErrDatabaseQueryFailed.Print(responseWriter)
			}
			Data, APIerr := getOrgSettingFromQuery(Query)
			if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			resp := conf.ApiResponse{200, "success", getOrgSettings_Success{Data}}
			resp.Print(responseWriter)
			return true
		} else {
			return conf.ErrUserTokenIncorrect.Print(responseWriter)
		}
	}
	return false
}

func getOrgSettingFromQuery(rows *sql.Rows) (OrgSettings, *conf.ApiResponse) {
	OrgSettingsRaw := make(map[string]string)
	var key, value string
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&key, &value)
		if err != nil {
			log.Print(err)
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