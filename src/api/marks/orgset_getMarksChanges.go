package marks

import (
	"net/http"
	"forcamp/src/api/authorization"
	"forcamp/conf"
	"forcamp/src/api/orgset"
	"forcamp/src"
	"database/sql"
	"forcamp/src/api/orgset/teams"
)

type MarksChange struct {
	ID int64 `json:"id"`
	Employee userData `json:"employee"`
	Participant userData `json:"participant"`
	Text string `json:"text"`
	Change int64 `json:"change"`
	Time string `json:"time"`
}

type marksChange_Raw struct {
	ID int64 `json:"id"`
	Employee_login string
	Participant_login string
	Reason_ID int64
	Time string
}

type userData struct {
	Login		   string `json:"login"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Middlename     string `json:"middlename"`
	Team           teams.Team `json:"team"`
	Access         int `json:"access"`
	Avatar         string `json:"avatar"`
	Sex            int `json:"sex"`
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
	marksChangesRaw, apiErr := getMarksChangesFromDataTable(login)
	if apiErr != nil {
		return nil, apiErr
	}
	marksChanges := make([]MarksChange, 0)
	for i := range marksChangesRaw {
		reasonText, reasonChange, apiErr := getReasonText(marksChangesRaw[i].Reason_ID)
		if apiErr != nil {
			return nil, apiErr
		}
		employeeData, apiErr := getUserData(marksChangesRaw[i].Employee_login)
		participantData, apiErr := getUserData(marksChangesRaw[i].Participant_login)
		marksChanges = append(marksChanges, MarksChange{ID: marksChangesRaw[i].ID,
			Employee: employeeData,
			Participant: participantData,
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
			return nil, conf.ErrDatabaseQueryFailed
		}
		marksChangesRaw = append(marksChangesRaw, marksChange_Raw{ID: id, Employee_login: employee_login, Participant_login: participant_login, Time: time, Reason_ID: reason_id})
	}
	return marksChangesRaw, nil
}

func getReasonText(reason_id int64) (string, int64, *conf.ApiResponse) {
	var (
		text string
		change int64
	)
	err := src.CustomConnection.QueryRow("SELECT text, modification FROM reasons WHERE id=?", reason_id).Scan(&text, &change)
	if err != nil {
		return "", 0, conf.ErrDatabaseQueryFailed
	}
	return text, change, nil
}

func getUserData(login string) (userData, *conf.ApiResponse) {
	var data userData
	var teamID int64
	err := src.CustomConnection.QueryRow("SELECT name, surname, middlename, sex, access, avatar, team FROM users " +
		"WHERE login=?", login).Scan(&data.Name, &data.Surname, &data.Middlename, &data.Sex, &data.Access, &data.Avatar, &teamID)
	if err != nil {
		return data, conf.ErrDatabaseQueryFailed
	}
	teamData, apiErr := getTeamInfo(teamID); if apiErr != nil {
		return data, apiErr
	}
	data.Login = login
	data.Team = teamData
	return data, nil
}

func getTeamInfo(teamID int64) (teams.Team, *conf.ApiResponse){
	var teamInfo teams.Team
	rows, err := src.CustomConnection.Query("SELECT * FROM teams WHERE id=?", teamID); if err != nil {
		return teamInfo, conf.ErrDatabaseQueryFailed
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&teamInfo.Id, &teamInfo.Name); if err != nil {
			return teamInfo, conf.ErrDatabaseQueryFailed
		}
		leader, apiErr := teams.GetTeamLeader(teamID); if apiErr != nil {
			return teamInfo, apiErr
		}
		participantsData, apiErr := teams.GetTeamParticipants(teamID); if apiErr != nil {
			return teamInfo, apiErr
		}
		teamInfo.Leader = leader
		teamInfo.Participants = participantsData
	}
	return teamInfo, nil
}
