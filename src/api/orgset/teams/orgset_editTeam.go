package teams

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/api/orgset"
)

func EditTeam(token string, name string, id int64, responseWriter http.ResponseWriter) bool{
	if checkTeamData(name, responseWriter) && orgset.IsUserAdmin(token, responseWriter){
		organizationName, _, apiErr := orgset.GetUserOrganizationAndIdByToken(token)
		if apiErr != nil{
			return apiErr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(organizationName)
		apiErr = editTeam_Request(name, id)
		if apiErr != nil{
			return apiErr.Print(responseWriter)
		}
		conf.RequestSuccess.Print(responseWriter)
	}
	return true
}

func editTeam_Request(name string, id int64) *conf.ApiResponse{
	query, err := src.CustomConnection.Prepare("UPDATE teams SET name=? WHERE id=?")
	if err != nil{
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(name, id)
	query.Close()
	if err != nil{
		return conf.ErrDatabaseQueryFailed
	}
	return  nil
}