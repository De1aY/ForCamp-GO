package reasons

import (
	"forcamp/src"
	"forcamp/src/api/authorization"
	"net/http"
	"forcamp/conf"
	"forcamp/src/api/orgset"
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

type getReasons_Success struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Reasons []Reason `json:"reasons"`
}

func (success *getReasons_Success) toJSON() string {
	resp, _ := json.Marshal(success)
	return string(resp)
}

func GetReasons(token string, ResponseWriter http.ResponseWriter) bool{
	if authorization.CheckTokenForEmpty(token, ResponseWriter) {
		if authorization.CheckToken(token, ResponseWriter) {
			Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			src.CustomConnection = src.Connect_Custom(Organization)
			resp, APIerr := getReasons_Request()
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			fmt.Fprintf(ResponseWriter, resp.toJSON())
		} else {
			return conf.PrintError(conf.ErrUserTokenIncorrect, ResponseWriter)
		}
	}
	return true
}

func getReasons_Request() (getReasons_Success, *conf.ApiError){
	Query, err := src.CustomConnection.Query("SELECT id,cat_id,text,modification FROM reasons")
	if err != nil {
		log.Print(err)
		return getReasons_Success{}, conf.ErrDatabaseQueryFailed
	}
	Reasons, APIerr := getReasonsFromQuery(Query)
	if APIerr != nil {
		return getReasons_Success{}, APIerr
	}
	return getReasons_Success{200, "success", Reasons}, nil
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
	if reasons == nil {
		return make([]Reason, 0), nil
	}
	return reasons, nil
}