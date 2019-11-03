/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package teams

import (
	"net/http"
	"nullteam.info/wplay/demo/conf"
	"nullteam.info/wplay/demo/src"
	"nullteam.info/wplay/demo/src/api/orgset"
)

type addTeam_Success struct {
	ID int64 `json:"id"`
}

func AddTeam(token string, name string, responseWriter http.ResponseWriter) bool{
	if checkTeamData(name, responseWriter) && orgset.IsUserAdmin(token, responseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndIdByToken(token)
		if APIerr != nil {
			return APIerr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		TeamID, APIerr := addTeam_Request(name)
		if APIerr != nil{
			return APIerr.Print(responseWriter)
		}
		resp := conf.ApiResponse{200, "success", addTeam_Success{TeamID}}
		resp.Print(responseWriter)
	}
	return true
}

func addTeam_Request(name string) (int64, *conf.ApiResponse){
	Query, err := src.CustomConnection.Prepare("INSERT INTO teams(name) VALUES(?)")
	if err != nil{
		return 0, conf.ErrDatabaseQueryFailed
	}
	Resp, err := Query.Exec(name)
	Query.Close()
	if err != nil{
		return 0, conf.ErrDatabaseQueryFailed
	}
	TeamID, err := Resp.LastInsertId()
	if err != nil{
		return 0, conf.ErrDatabaseQueryFailed
	}
	return TeamID, nil
}

func checkTeamData(name string, w http.ResponseWriter) bool{
	if len(name) > 0 {
		return true
	} else {
		return conf.ErrTeamNameEmpty.Print(w)
	}
}
