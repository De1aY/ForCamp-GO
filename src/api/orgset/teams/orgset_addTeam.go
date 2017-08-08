package teams

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src"
	"log"
	"forcamp/src/api/orgset"
)

type addTeam_Success struct {
	ID int64 `json:"id"`
}

func AddTeam(token string, name string, responseWriter http.ResponseWriter) bool{
	if checkTeamData(name, responseWriter) && orgset.CheckUserAccess(token, responseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
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
		return conf.ErrTeamNameEmpty.Print(w)
	}
}