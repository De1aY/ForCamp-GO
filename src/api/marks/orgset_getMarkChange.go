package marks

import (
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/api/orgset/teams"
)

type MarkChange struct {
	Employee      userData `json:"employee"`
	Participant   userData `json:"participant"`
	Text          string   `json:"text"`
	Initial_Value int64    `json:"initial_value"`
	Final_Value   int64    `json:"final_value"`
	Change        int64    `json:"change"`
	Category_ID   int64    `json:"category_id"`
}

type markChange_Raw struct {
	Employee_ID    int64
	Participant_ID int64
	Category_ID    int64
	Text           string
	Initial_Value  int64
	Final_Value    int64
}

type userData struct {
	ID         int64      `json:"id"`
	Name       string     `json:"name"`
	Surname    string     `json:"surname"`
	Middlename string     `json:"middlename"`
	Team       teams.Team `json:"team"`
	Access     int        `json:"access"`
	Avatar     string     `json:"avatar"`
	Sex        int        `json:"sex"`
}

type getMarksChanges_Success struct {
	Marks_changes []MarkChange `json:"marks_changes"`
}

func GetMarkChange(eventID int64) (MarkChange, *conf.ApiResponse) {
	markChangeRaw, apiErr := getRawMarkChange(eventID)
	if apiErr != nil {
		return MarkChange{}, apiErr
	}
	employeeData, apiErr := getUserData(markChangeRaw.Employee_ID);
	if apiErr != nil {
		return MarkChange{}, apiErr
	}
	participantData, apiErr := getUserData(markChangeRaw.Participant_ID);
	if apiErr != nil {
		return MarkChange{}, apiErr
	}
	markChange := MarkChange{
		Employee:      employeeData,
		Participant:   participantData,
		Text:          markChangeRaw.Text,
		Initial_Value: markChangeRaw.Initial_Value,
		Final_Value:   markChangeRaw.Final_Value,
		Change:        markChangeRaw.Final_Value - markChangeRaw.Initial_Value,
		Category_ID:   markChangeRaw.Category_ID,
	}
	return markChange, nil
}

func getRawMarkChange(eventId int64) (markChange_Raw, *conf.ApiResponse) {
	var markChangeRaw markChange_Raw
	if eventId > 0 {
		err := src.CustomConnection.QueryRow("SELECT employee_id, category_id, participant_id, text, "+
			"initial_value, final_value FROM marks_changes WHERE id=?", eventId).Scan(&markChangeRaw.Employee_ID,
			&markChangeRaw.Category_ID, &markChangeRaw.Participant_ID,
			&markChangeRaw.Text, markChangeRaw.Initial_Value, markChangeRaw.Final_Value)
		if err != nil {
			return markChangeRaw, conf.ErrDatabaseQueryFailed
		}
	} else {
		return markChangeRaw, conf.ErrIdIncorrect
	}
	return markChangeRaw, nil
}

func getUserData(id int64) (userData, *conf.ApiResponse) {
	var data userData
	var teamID int64
	err := src.CustomConnection.QueryRow("SELECT name, surname, middlename, sex, access, avatar, team FROM users "+
		"WHERE id=?", id).Scan(&data.Name, &data.Surname, &data.Middlename, &data.Sex, &data.Access, &data.Avatar, &teamID)
	if err != nil {
		return data, conf.ErrDatabaseQueryFailed
	}
	teamData, apiErr := getTeamInfo(teamID);
	if apiErr != nil {
		return data, apiErr
	}
	data.ID = id
	data.Team = teamData
	return data, nil
}

func getTeamInfo(teamID int64) (teams.Team, *conf.ApiResponse) {
	var teamInfo teams.Team
	rows, err := src.CustomConnection.Query("SELECT * FROM teams WHERE id=?", teamID);
	if err != nil {
		return teamInfo, conf.ErrDatabaseQueryFailed
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&teamInfo.Id, &teamInfo.Name);
		if err != nil {
			return teamInfo, conf.ErrDatabaseQueryFailed
		}
		leader, apiErr := teams.GetTeamLeader(teamID);
		if apiErr != nil {
			return teamInfo, apiErr
		}
		participantsData, apiErr := teams.GetTeamParticipants(teamID);
		if apiErr != nil {
			return teamInfo, apiErr
		}
		teamInfo.Leader = leader
		teamInfo.Participants = participantsData
	}
	return teamInfo, nil
}
