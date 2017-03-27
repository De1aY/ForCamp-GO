package orgset

import (
	"net/http"
	"forcamp/conf"
	"database/sql"
	"encoding/json"
	"fmt"
	"forcamp/src/authorization"
	"forcamp/src"
)

func GetOrgSettings(token string, ResponseWriter http.ResponseWriter) bool {
	if authorization.CheckTokenForEmpty(token, ResponseWriter) && authorization.CheckToken(token, ResponseWriter) {
		Organization, _ := getUserOrganizationAndLoginByToken(token, ResponseWriter)
		NewConnection = src.Connect_Custom(Organization)
		Query, err := NewConnection.Query("SELECT * FROM settings")
		if err != nil {
			return conf.PrintError(conf.ErrDatabaseQueryFailed, ResponseWriter)
		}
		Data := getOrgSettingFromQuery(Query, ResponseWriter)
		Resp := GetOrgSettings_Success{200, "success", Data}
		Response, _ := json.Marshal(Resp)
		fmt.Fprintf(ResponseWriter, string(Response))
		return true
	}
	return false
}

func getOrgSettingFromQuery(rows *sql.Rows, w http.ResponseWriter) OrgSettings {
	OrgSettingsRaw := make(map[string]string)
	var key, value string
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&key, &value)
		if err != nil {
			conf.PrintError(conf.ErrDatabaseQueryFailed, w)
		}
		OrgSettingsRaw[key] = value
	}
	OrgSettings := OrgSettings{Organization: OrgSettingsRaw["organization"],
		Team: OrgSettingsRaw["team"],
		Participant: OrgSettingsRaw["participant"],
		Period: OrgSettingsRaw["period"],
		SelfMarks: OrgSettingsRaw["self_marks"]}
	return OrgSettings
}