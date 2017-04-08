package reasons

import (
	"forcamp/src"
	"forcamp/src/authorization"
	"net/http"
	"forcamp/conf"
	"forcamp/src/orgset"
	"database/sql"
	"log"
	"encoding/json"
	"fmt"
)

type Reason struct {
	Id int64 `json:"id"`
	Cat_id int64 `json:"cat_id"`
	Text string `json:"text"`
	Change int `json:"change"`
}

type GetReasons_Success struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Reasons []Reason `json:"reasons"`
}

func GetReasons(token string, ResponseWriter http.ResponseWriter) bool{
	Connection := src.Connect()
	defer Connection.Close()
	if authorization.CheckTokenForEmpty(token, ResponseWriter) {
		if authorization.CheckToken(token, Connection, ResponseWriter) {
			Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token, Connection)
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			NewConnection:= src.Connect_Custom(Organization)
			defer NewConnection.Close()
			Resp, APIerr := getReasons_Request(NewConnection)
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			Response, _ := json.Marshal(Resp)
			fmt.Fprintf(ResponseWriter, string(Response))
		} else {
			return conf.PrintError(conf.ErrUserTokenIncorrect, ResponseWriter)
		}
	}
	return true
}

func getReasons_Request(connection *sql.DB) (GetReasons_Success, *conf.ApiError){
	Query, err := connection.Query("SELECT id,cat_id,text,modification FROM reasons")
	if err != nil {
		log.Print(err)
		return GetReasons_Success{}, conf.ErrDatabaseQueryFailed
	}
	Reasons, APIerr := getReasonsFromQuery(Query)
	if APIerr != nil {
		return GetReasons_Success{}, APIerr
	}
	return GetReasons_Success{200, "success", Reasons}, nil
}

func getReasonsFromQuery(rows *sql.Rows) ([]Reason, *conf.ApiError){
	defer rows.Close()
	var (
		reason Reason
		reasons []Reason
	)
	for rows.Next(){
		err := rows.Scan(&reason.Id, &reason.Cat_id, &reason.Text, &reason.Change)
		if err != nil {
			log.Print(err)
			return nil, conf.ErrDatabaseQueryFailed
		}
		reasons = append(reasons, reason)
	}
	return reasons, nil
}