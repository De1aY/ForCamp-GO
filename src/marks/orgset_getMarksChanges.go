package marks

import (
	"net/http"
	"forcamp/src/authorization"
	"forcamp/conf"
	"forcamp/src/orgset"
	"forcamp/src"
	"log"
	"encoding/json"
	"fmt"
)

type marksChange struct {
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
	Code int `json:"code"`
	Status string `json:"status"`
	Marks_changes []marksChange `json:"marks_changes"`
}

func (success *getMarksChanges_Success) toJSON() string {
	resp, _ := json.Marshal(success)
	return string(resp)
}


func GetMarksChanges(token string, responseWriter http.ResponseWriter) bool {
	if authorization.CheckTokenForEmpty(token, responseWriter){
		if authorization.CheckToken(token, responseWriter) {
			organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
			if APIerr != nil {
				return conf.PrintError(APIerr, responseWriter)
			}
			src.CustomConnection = src.Connect_Custom(organization)

			response, APIerr := getMarksChanges_Request()
			if APIerr != nil {
				return conf.PrintError(APIerr, responseWriter)
			}
			fmt.Fprintf(responseWriter, response.toJSON())
		} else {
			return conf.PrintError(conf.ErrUserTokenIncorrect, responseWriter)
		}
	}
	return true
}

func getMarksChanges_Request() (getMarksChanges_Success, *conf.ApiError) {
	marksChangesRaw, APIerr := getMarksChangesFromDataTable()
	if APIerr != nil {
		return getMarksChanges_Success{}, APIerr
	}
	marksChanges := make([]marksChange, 0)
	for i := range marksChangesRaw {
		reasonText, reasonChange, APIerr := getReasonText(marksChangesRaw[i].Reason_ID)
		if APIerr != nil {
			return getMarksChanges_Success{}, APIerr
		}
		marksChanges = append(marksChanges, marksChange{ID: marksChangesRaw[i].ID, Employee_login: marksChangesRaw[i].Employee_login, Participant_login: marksChangesRaw[i].Participant_login, Text: reasonText, Time: marksChangesRaw[i].Time, Change: reasonChange})
	}
	return getMarksChanges_Success{200, "success", marksChanges}, nil
}

func getMarksChangesFromDataTable() ([]marksChange_Raw, *conf.ApiError) {
	query, err := src.CustomConnection.Query("SELECT id, employee_login, participant_login, reason_id, time FROM marks_changes")
	defer query.Close()
	if err != nil {
		log.Print(err)
		return make([]marksChange_Raw, 0), conf.ErrDatabaseQueryFailed
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
			return make([]marksChange_Raw, 0), conf.ErrDatabaseQueryFailed
		}
		marksChangesRaw = append(marksChangesRaw, marksChange_Raw{ID: id, Employee_login: employee_login, Participant_login: participant_login, Time: time, Reason_ID: reason_id})
	}
	return marksChangesRaw, nil
}

func getReasonText(reason_id int64) (string, int, *conf.ApiError) {
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
