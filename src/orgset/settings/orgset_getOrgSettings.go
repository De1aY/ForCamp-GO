package settings

import (
	"net/http"
	"forcamp/conf"
	"database/sql"
	"encoding/json"
	"fmt"
	"forcamp/src/authorization"
	"forcamp/src"
	"log"
	"forcamp/src/orgset"
)

type OrgSettings struct {
	Participant string `json:"participant"`
	Team string `json:"team"`
	Organization string `json:"organization"`
	Period string `json:"period"`
	SelfMarks string `json:"self_marks"`
}

type GetOrgSettings_Success struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Settings OrgSettings `json:"settings"`
}

func GetOrgSettings(token string, ResponseWriter http.ResponseWriter) bool {
	Connection := src.Connect()
	defer Connection.Close()
	if authorization.CheckTokenForEmpty(token, ResponseWriter){
		if authorization.CheckToken(token, Connection, ResponseWriter) {
			Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token, Connection)
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			NewConnection := src.Connect_Custom(Organization)
			defer NewConnection.Close()
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