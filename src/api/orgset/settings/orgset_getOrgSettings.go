package settings

import (
	"net/http"
	"forcamp/conf"
	"database/sql"
	"encoding/json"
	"fmt"
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
	Code int `json:"code"`
	Status string `json:"status"`
	Settings OrgSettings `json:"settings"`
}

func (success *getOrgSettings_Success) toJSON() string {
	resp, _ := json.Marshal(success)
	return string(resp)
}

func GetOrgSettings(token string, ResponseWriter http.ResponseWriter) bool {
	if authorization.CheckTokenForEmpty(token, ResponseWriter){
		if authorization.CheckToken(token, ResponseWriter) {
			Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			src.CustomConnection = src.Connect_Custom(Organization)
			Query, err := src.CustomConnection.Query("SELECT * FROM settings")
			if err != nil {
				log.Print(err)
				return conf.PrintError(conf.ErrDatabaseQueryFailed, ResponseWriter)
			}
			Data, APIerr := getOrgSettingFromQuery(Query)
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			resp := getOrgSettings_Success{200, "success", Data}
			fmt.Fprintf(ResponseWriter, resp.toJSON())
			return true
		} else {
			return conf.PrintError(conf.ErrUserTokenIncorrect, ResponseWriter)
		}
	}
	return false
}

func getOrgSettingFromQuery(rows *sql.Rows) (OrgSettings, *conf.ApiError) {
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