package reasons

import (
	"forcamp/src/orgset"
	"net/http"
	"forcamp/src"
	"forcamp/conf"
	"log"
)

func DeleteReason(token string, id int64, ResponseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token, ResponseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil {
			return conf.PrintError(APIerr, ResponseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		APIerr = deleteReason_Request(id)
		if APIerr != nil {
			return conf.PrintError(APIerr, ResponseWriter)
		}
		conf.PrintSuccess(conf.RequestSuccess, ResponseWriter)
	}
	return true
}

func deleteReason_Request(id int64) *conf.ApiError{
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