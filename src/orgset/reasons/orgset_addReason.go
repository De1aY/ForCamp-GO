package reasons

import (
	"forcamp/src/orgset"
	"net/http"
	"forcamp/src"
	"forcamp/conf"
	"database/sql"
	"log"
	"encoding/json"
	"fmt"
)

type AddReason_Success struct {
	Code int `json:"code"`
	Status string `json:"status"`
	ID int64 `json:"id"`
}

func AddReason(token string, reason Reason, ResponseWriter http.ResponseWriter) bool{
	connection := src.Connect()
	defer connection.Close()
	if orgset.CheckUserAccess(token, connection, ResponseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token, connection)
		if APIerr != nil {
			return conf.PrintError(APIerr, ResponseWriter)
		}
		NewConnection := src.Connect_Custom(Organization)
		defer NewConnection.Close()
		if orgset.CheckCategoryId(reason.Cat_id, ResponseWriter, NewConnection){
			Resp, APIerr := addReason_Request(NewConnection, reason)
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			Response, _ := json.Marshal(Resp)
			fmt.Fprintf(ResponseWriter, string(Response))
		}
	}
	return true
}

func addReason_Request(connection *sql.DB, reason Reason) (AddReason_Success, *conf.ApiError){
	Query, err := connection.Prepare("INSERT INTO reasons(cat_id,text,modification) VALUES(?,?,?)")
	if err != nil {
		log.Print(err)
		return AddReason_Success{}, conf.ErrDatabaseQueryFailed
	}
	Resp, err := Query.Exec(reason.Cat_id, reason.Text, reason.Change)
	if err != nil {
		log.Print(err)
		return AddReason_Success{}, conf.ErrDatabaseQueryFailed
	}
	ID, err := Resp.LastInsertId()
	if err != nil {
		log.Print(err)
		return AddReason_Success{}, conf.ErrDatabaseQueryFailed
	}
	return AddReason_Success{200, "success", ID}, nil
}