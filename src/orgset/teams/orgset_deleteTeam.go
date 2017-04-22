package teams

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src"
	"log"
	"forcamp/src/orgset"
)

func DeleteTeam(token string, id int64, ResponseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token, ResponseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		APIerr = deleteTeam_Request(id)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		return conf.PrintSuccess(conf.RequestSuccess, ResponseWriter)
	}
	return true
}

func deleteTeam_Request(id int64) *conf.ApiError{
	Query, err := src.CustomConnection.Prepare("DELETE FROM teams WHERE id=?")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(id)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	APIerr := deleteTeam_Users(id)
	if APIerr != nil {
		return APIerr
	}
	return nil
}

func deleteTeam_Users(id int64) *conf.ApiError{
	Query, err := src.CustomConnection.Prepare("UPDATE users SET team='0' WHERE team=?")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(id)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	return nil
}
