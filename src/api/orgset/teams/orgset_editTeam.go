package teams

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src"
	"log"
	"forcamp/src/api/orgset"
)

func EditTeam(token string, name string, id int64, responseWriter http.ResponseWriter) bool{
	if checkTeamData(name, responseWriter) && orgset.CheckUserAccess(token, responseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil{
			return APIerr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		APIerr = editTeam_Request(name, id)
		if APIerr != nil{
			return APIerr.Print(responseWriter)
		}
		conf.RequestSuccess.Print(responseWriter)
	}
	return true
}

func editTeam_Request(name string, id int64) *conf.ApiResponse{
	Query, err := src.CustomConnection.Prepare("UPDATE teams SET name=? WHERE id=?")
	if err != nil{
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(name, id)
	Query.Close()
	if err != nil{
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return  nil
}