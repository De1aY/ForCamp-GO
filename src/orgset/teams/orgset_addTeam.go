package teams

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src"
	"encoding/json"
	"fmt"
	"log"
	"forcamp/src/orgset"
)

type AddTeam_Success struct {
	Code int `json:"code"`
	Status string `json:"status"`
	ID int64 `json:"id"`
}

func AddTeam(token string, name string, ResponseWriter http.ResponseWriter) bool{
	if checkTeamData(name, ResponseWriter) && orgset.CheckUserAccess(token, ResponseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		src.NewConnection = src.Connect_Custom(Organization)
		TeamID, APIerr := addTeam_Request(name)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		Resp := AddTeam_Success{200, "success", TeamID}
		Response, _ := json.Marshal(Resp)
		fmt.Fprintf(ResponseWriter, string(Response))
	}
	return true
}

func addTeam_Request(name string) (int64, *conf.ApiError){
	Query, err := src.NewConnection.Prepare("INSERT INTO teams(name) VALUES(?)")
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