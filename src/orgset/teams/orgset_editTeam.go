package teams

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src"
	"log"
	"forcamp/src/orgset"
)

func EditTeam(token string, name string, id int64, ResponseWriter http.ResponseWriter) bool{
	if checkTeamData(name, ResponseWriter) && orgset.CheckUserAccess(token, ResponseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		APIerr = editTeam_Request(name, id)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		conf.PrintSuccess(conf.RequestSuccess, ResponseWriter)
	}
	return true
}

func editTeam_Request(name string, id int64) *conf.ApiError{
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