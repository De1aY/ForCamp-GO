package teams

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src"
	"encoding/json"
	"fmt"
	"log"
	"forcamp/src/api/orgset"
)

type addTeam_Success struct {
	Code int `json:"code"`
	Status string `json:"status"`
	ID int64 `json:"id"`
}

func (success *addTeam_Success) toJSON() string {
	resp, _ := json.Marshal(success)
	return string(resp)
}

func AddTeam(token string, name string, ResponseWriter http.ResponseWriter) bool{
	if checkTeamData(name, ResponseWriter) && orgset.CheckUserAccess(token, ResponseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil {
			return conf.PrintError(APIerr, ResponseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		TeamID, APIerr := addTeam_Request(name)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		resp := addTeam_Success{200, "success", TeamID}
		fmt.Fprintf(ResponseWriter, resp.toJSON())
	}
	return true
}

func addTeam_Request(name string) (int64, *conf.ApiError){
	Query, err := src.CustomConnection.Prepare("INSERT INTO teams(name) VALUES(?)")
	if err != nil{
		log.Print(err)
		return 0, conf.ErrDatabaseQueryFailed
	}
	Resp, err := Query.Exec(name)
	Query.Close()
	if err != nil{
		log.Print(err)
		return 0, conf.ErrDatabaseQueryFailed
	}
	TeamID, err := Resp.LastInsertId()
	if err != nil{
		log.Print(err)
		return 0, conf.ErrDatabaseQueryFailed
	}
	return TeamID, nil
}

func checkTeamData(name string, w http.ResponseWriter) bool{
	if len(name) > 0 {
		return true
	} else {
		return conf.PrintError(conf.ErrCategoryNameEmpty, w)
	}
}