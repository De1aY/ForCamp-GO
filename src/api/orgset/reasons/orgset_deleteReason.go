package reasons

import (
	"forcamp/src/api/orgset"
	"net/http"
	"forcamp/src"
	"forcamp/conf"
	"log"
)

func DeleteReason(token string, id int64, responseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token, responseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil {
			return APIerr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		APIerr = deleteReason_Request(id)
		if APIerr != nil {
			return APIerr.Print(responseWriter)
		}
		conf.RequestSuccess.Print(responseWriter)
	}
	return true
}

func deleteReason_Request(id int64) *conf.ApiResponse{
	Query, err := src.CustomConnection.Prepare("DELETE FROM reasons WHERE id=?")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(id)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}