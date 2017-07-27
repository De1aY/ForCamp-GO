package marks

import (
	"net/http"
	"forcamp/src/api/authorization"
	"forcamp/conf"
	"forcamp/src/api/orgset"
	"forcamp/src"
	"log"
	"database/sql"
)

type MarksChange struct {
	ID int64 `json:"id"`
	Employee_login string `json:"employee_login"`
	Participant_login string `json:"participant_login"`
	Text string `json:"text"`
	Change int `json:"change"`
	Time string `json:"time"`
}

type marksChange_Raw struct {
	ID int64 `json:"id"`
	Employee_login string
	Participant_login string
	Reason_ID int64
	Time string
}

type getMarksChanges_Success struct {
	Marks_changes []MarksChange `json:"marks_changes"`
}

func GetMarksChanges(token string, responseWriter http.ResponseWriter) bool {
	if authorization.CheckTokenForEmpty(token, responseWriter){
		if authorization.CheckToken(token, responseWriter) {
			organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
			if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			src.CustomConnection = src.Connect_Custom(organization)
			rawResp, APIerr := GetMarksChanges_Request("")
			if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			resp := &conf.ApiResponse{200, "success", getMarksChanges_Success{rawResp}}
			resp.Print(responseWriter)
		} else {
			return conf.ErrUserTokenIncorrect.Print(responseWriter)
		}
	}
	return true
}

func GetMarksChanges_Request(login string) ([]MarksChange, *conf.ApiResponse) {
	marksChangesRaw, APIerr := getMarksChangesFromDataTable(login)
	if APIerr != nil {
		return nil, APIerr
	}
	marksChanges := make([]MarksChange, 0)
	for i := range marksChangesRaw {
		reasonText, reasonChange, APIerr := getReasonText(marksChangesRaw[i].Reason_ID)
		if APIerr != nil {
			return nil, APIerr
		}
		marksChanges = append(marksChanges, MarksChange{ID: marksChangesRaw[i].ID,
			Employee_login: marksChangesRaw[i].Employee_login,
			Participant_login: marksChangesRaw[i].Participant_login,
			Text: reasonText,
			Time: marksChangesRaw[i].Time,
			Change: reasonChange})
	}
	return marksChanges, nil
}

func getMarksChangesFromDataTable(login string) ([]marksChange_Raw, *conf.ApiResponse) {
	var query *sql.Rows
	var err error
	if len(login) > 2 {
		query, err = src.CustomConnection.Query("SELECT id, employee_login, participant_login, reason_id, time FROM marks_changes WHERE employee_login=? OR participant_login=?", login, login)
	} else {
		query, err = src.CustomConnection.Query("SELECT id, employee_login, participant_login, reason_id, time FROM marks_changes")
	}
	defer query.Close()
	if err != nil {
		log.Print(err)
		return nil, conf.ErrDatabaseQueryFailed
	}
	marksChangesRaw := make([]marksChange_Raw, 0)
	var (
		id int64
		employee_login string
		participant_login string
		reason_id int64
		time string
	)
	for query.Next() {
		err := query.Scan(&id, &employee_login, &participant_login, &reason_id, &time)
		if err != nil {
			log.Print(err)
			return nil, conf.ErrDatabaseQueryFailed
		}
		marksChangesRaw = append(marksChangesRaw, marksChange_Raw{ID: id, Employee_login: employee_login, Participant_login: participant_login, Time: time, Reason_ID: reason_id})
	}
	return marksChangesRaw, nil
}

func getReasonText(reason_id int64) (string, int, *conf.ApiResponse) {
	var (
		text string
		change int
	)
	err := src.CustomConnection.QueryRow("SELECT text, modification FROM reasons WHERE id=?", reason_id).Scan(&text, &change)
	if err != nil {
		log.Print(err)
		return "", 0, conf.ErrDatabaseQueryFailed
	}
	return text, change, nil
}
