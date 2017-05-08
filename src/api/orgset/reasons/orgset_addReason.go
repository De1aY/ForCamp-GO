package reasons

import (
	"forcamp/src/api/orgset"
	"net/http"
	"forcamp/src"
	"forcamp/conf"
	"log"
	"encoding/json"
	"fmt"
)

type addReason_Success struct {
	Code int `json:"code"`
	Status string `json:"status"`
	ID int64 `json:"id"`
}

func (success *addReason_Success) toJSON() string {
	resp, _ := json.Marshal(success)
	return string(resp)
}

func AddReason(token string, reason Reason, ResponseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token, ResponseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil {
			return conf.PrintError(APIerr, ResponseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		if orgset.CheckCategoryId(reason.Cat_id, ResponseWriter){
			resp, APIerr := addReason_Request(reason)
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			fmt.Fprintf(ResponseWriter, resp.toJSON())
		}
	}
	return true
}

func addReason_Request(reason Reason) (addReason_Success, *conf.ApiError){
	Query, err := src.CustomConnection.Prepare("INSERT INTO reasons(cat_id,text,modification) VALUES(?,?,?)")
	if err != nil {
		log.Print(err)
		return addReason_Success{}, conf.ErrDatabaseQueryFailed
	}
	Resp, err := Query.Exec(reason.Cat_id, reason.Text, reason.Change)
	if err != nil {
		log.Print(err)
		return addReason_Success{}, conf.ErrDatabaseQueryFailed
	}
	ID, err := Resp.LastInsertId()
	if err != nil {
		log.Print(err)
		return addReason_Success{}, conf.ErrDatabaseQueryFailed
	}
	return addReason_Success{200, "success", ID}, nil
}