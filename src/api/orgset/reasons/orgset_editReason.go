package reasons

import (
	"forcamp/src/api/orgset"
	"net/http"
	"forcamp/src"
	"forcamp/conf"
)

func EditReason(token string, reason Reason, responseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token, responseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil {
			return APIerr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		if orgset.CheckCategoryId(reason.Cat_id, responseWriter){
			APIerr = editReason_Request(reason)
			if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			conf.RequestSuccess.Print(responseWriter)
		}
	}
	return true
}

func editReason_Request(reason Reason) *conf.ApiResponse{
	Query, err := src.CustomConnection.Prepare("UPDATE reasons SET text=?, modification=?, cat_id=? WHERE id=?")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(reason.Text, reason.Change, reason.Cat_id, reason.Id)
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}