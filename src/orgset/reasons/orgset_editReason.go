package reasons

import (
	"forcamp/src/orgset"
	"net/http"
	"forcamp/src"
	"forcamp/conf"
	"log"
)

func EditReason(token string, reason Reason, ResponseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token, ResponseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil {
			return conf.PrintError(APIerr, ResponseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		if orgset.CheckCategoryId(reason.Cat_id, ResponseWriter){
			APIerr = editReason_Request(reason)
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			conf.PrintSuccess(conf.RequestSuccess, ResponseWriter)
		}
	}
	return true
}

func editReason_Request(reason Reason) *conf.ApiError{
	Query, err := src.CustomConnection.Prepare("UPDATE reasons SET text=?, modification=?, cat_id=? WHERE id=?")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(reason.Text, reason.Change, reason.Cat_id, reason.Id)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}