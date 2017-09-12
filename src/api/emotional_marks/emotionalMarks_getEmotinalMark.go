package emotional_marks

import (
	"forcamp/conf"
	"forcamp/src"
	"time"
	"forcamp/src/api/orgset/settings"
	"strconv"
	"forcamp/src/api/orgset/teams"
)


type EmotionalMark struct {
	IsNew bool   `json:"is_new"`
	Value int64   `json:"value"`
	ParticipantData userData `json:"participant"`
}

type EmotionalMark_Raw struct {
	Participant_ID int64
	Mark int64
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

func GetEmotionalMark(event_id int64, event_time string) (EmotionalMark, *conf.ApiResponse) {
	var emotionalMark EmotionalMark
	var rawEmotionalMark EmotionalMark_Raw
	err := src.CustomConnection.QueryRow("SELECT participant_id, mark " +
		"FROM emotional_marks WHERE id=?", event_id).Scan(&rawEmotionalMark.Participant_ID, &rawEmotionalMark.Mark)
	if err != nil {
		return emotionalMark, conf.ErrDatabaseQueryFailed
	}
	timeStamp, err := time.Parse("2006-01-02 15:04:05.999", event_time); if err != nil {
		return emotionalMark, conf.ErrDatabaseQueryFailed
	}
	duration := timeStamp.Sub(time.Now())
	organizationSettings, apiErr := settings.GetOrgSettings_Request(); if apiErr != nil {
		return emotionalMark, apiErr
	}
	emotionalMarkPeriod, err := strconv.ParseInt(organizationSettings.EmotionalMarkPeriod, 10, 64)
	if err != nil {
		return emotionalMark, conf.ErrDatabaseQueryFailed
	}
	if duration > time.Hour * time.Duration(emotionalMarkPeriod) * time.Duration(-1) {
		emotionalMark.IsNew = true
	} else {
		emotionalMark.IsNew = false
	}
	participantData, apiErr := getUserData(rawEmotionalMark.Participant_ID); if apiErr != nil {
		return emotionalMark, apiErr
	}
	emotionalMark.Value = rawEmotionalMark.Mark
	emotionalMark.ParticipantData = participantData
	return emotionalMark, nil
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
