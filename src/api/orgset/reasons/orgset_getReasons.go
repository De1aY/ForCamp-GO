package reasons

import (
	"forcamp/src"
	"forcamp/src/api/authorization"
	"net/http"
	"forcamp/conf"
	"forcamp/src/api/orgset"
	"database/sql"
	"log"
)

type Reason struct {
	Id int64 `json:"id"`
	Cat_id int64 `json:"cat_id"`
	Text string `json:"text"`
	Change int `json:"change"`
}

type getReasons_Success struct {
	Reasons []Reason `json:"reasons"`
}

func GetReasons(token string, responseWriter http.ResponseWriter) bool{
	if authorization.CheckTokenForEmpty(token, responseWriter) {
		if authorization.CheckToken(token, responseWriter) {
			Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
			if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			src.CustomConnection = src.Connect_Custom(Organization)
			rawResp, APIerr := getReasons_Request()
			if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			resp := conf.ApiResponse{200, "success", rawResp}
			resp.Print(responseWriter)
		} else {
			return conf.ErrUserTokenIncorrect.Print(responseWriter)
		}
	}
	return true
}

func getReasons_Request() (getReasons_Success, *conf.ApiResponse){
	Query, err := src.CustomConnection.Query("SELECT id,cat_id,text,modification FROM reasons")
	if err != nil {
		log.Print(err)
		return getReasons_Success{}, conf.ErrDatabaseQueryFailed
	}
	Reasons, APIerr := getReasonsFromQuery(Query)
	if APIerr != nil {
		return getReasons_Success{}, APIerr
	}
	return getReasons_Success{Reasons}, nil
}

func getReasonsFromQuery(rows *sql.Rows) ([]Reason, *conf.ApiResponse){
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