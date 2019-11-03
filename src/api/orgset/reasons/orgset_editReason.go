/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package reasons

import (
	"nullteam.info/wplay/demo/src/api/orgset"
	"net/http"
	"nullteam.info/wplay/demo/src"
	"nullteam.info/wplay/demo/conf"
)

func EditReason(token string, reason Reason, responseWriter http.ResponseWriter) bool{
	if orgset.IsUserAdmin(token, responseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndIdByToken(token)
		if APIerr != nil {
			return APIerr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		if orgset.IsCategoryExist(reason.Cat_id, responseWriter){
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
	Query, err := src.CustomConnection.Prepare("UPDATE reasons SET text=?, modification=?, category_id=? WHERE id=?")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(reason.Text, reason.Change, reason.Cat_id, reason.Id)
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}
