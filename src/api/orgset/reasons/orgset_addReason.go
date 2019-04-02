/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package reasons

import (
	"wplay/src/api/orgset"
	"net/http"
	"wplay/src"
	"wplay/conf"
)

type addReason_Success struct {
	ID int64 `json:"id"`
}

func AddReason(token string, reason Reason, responseWriter http.ResponseWriter) bool{
	if orgset.IsUserAdmin(token, responseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndIdByToken(token)
		if APIerr != nil {
			return APIerr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		if orgset.IsCategoryExist(reason.Cat_id, responseWriter){
			rawResp, APIerr := addReason_Request(reason)
			if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			resp := conf.ApiResponse{200, "success", rawResp}
			resp.Print(responseWriter)
		}
	}
	return true
}

func addReason_Request(reason Reason) (addReason_Success, *conf.ApiResponse){
	Query, err := src.CustomConnection.Prepare("INSERT INTO reasons(category_id,text,modification) VALUES(?,?,?)")
	if err != nil {
		return addReason_Success{}, conf.ErrDatabaseQueryFailed
	}
	Resp, err := Query.Exec(reason.Cat_id, reason.Text, reason.Change)
	if err != nil {
		return addReason_Success{}, conf.ErrDatabaseQueryFailed
	}
	ID, err := Resp.LastInsertId()
	if err != nil {
		return addReason_Success{}, conf.ErrDatabaseQueryFailed
	}
	return addReason_Success{ID}, nil
}
