package orgset

import (
	"net/http"
	"forcamp/conf"
	"database/sql"
	"encoding/json"
	"fmt"
	"forcamp/src/authorization"
	"forcamp/src"
	"log"
)

func GetOrgSettings(token string, ResponseWriter http.ResponseWriter) bool {
	if authorization.CheckTokenForEmpty(token, ResponseWriter) && authorization.CheckToken(token, ResponseWriter) {
		Organization, _, APIerr := getUserOrganizationAndLoginByToken(token)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		NewConnection = src.Connect_Custom(Organization)
		Query, err := NewConnection.Query("SELECT * FROM settings")
		if err != nil {
			log.Print(err)
			return conf.PrintError(conf.ErrDatabaseQueryFailed, ResponseWriter)
		}
		Data, APIerr := getOrgSettingFromQuery(Query)
		if APIerr != nil {
			return conf.PrintError(APIerr, ResponseWriter)
		}
		Resp := GetOrgSettings_Success{200, "success", Data}
		Response, _ := json.Marshal(Resp)
		fmt.Fprintf(ResponseWriter, string(Response))
		return true
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