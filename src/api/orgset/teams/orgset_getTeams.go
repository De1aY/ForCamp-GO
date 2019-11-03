/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package teams

import (
	"net/http"
	"nullteam.info/wplay/demo/src/api/authorization"
	"nullteam.info/wplay/demo/conf"
	"nullteam.info/wplay/demo/src"
	"database/sql"
	"nullteam.info/wplay/demo/src/api/orgset"
)

type TeamLeader struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Middlename string `json:"middlename"`
	ID         string `json:"id"`
}

type Team struct {
	Id     int64 `json:"id"`
	Name   string `json:"name"`
	Leader TeamLeader `json:"leader"`
	Participants  []int64 `json:"participants"`
}

type getTeams_Success struct {
	Teams  []Team `json:"teams"`
}

func GetTeams(token string, responseWriter http.ResponseWriter) bool {
	if authorization.IsTokenNotEmpty(token, responseWriter) {
		if authorization.IsTokenValid(token, responseWriter) {
			organizationName, _, apiErr := orgset.GetUserOrganizationAndIdByToken(token); if apiErr != nil{
				return apiErr.Print(responseWriter)
			}
			src.CustomConnection = src.Connect_Custom(organizationName)
			rawResp, apiErr := getTeams(); if apiErr != nil {
				return apiErr.Print(responseWriter)
			}
			resp := conf.ApiResponse{200, "success", rawResp}
			resp.Print(responseWriter)
		} else {
			return conf.ErrUserTokenIncorrect.Print(responseWriter)
		}
	}
	return true
}

func getTeams() (getTeams_Success, *conf.ApiResponse) {
	rows, err := src.CustomConnection.Query("SELECT * FROM teams"); if err != nil {
		return getTeams_Success{}, conf.ErrDatabaseQueryFailed
	}
	teams, apiErr := getTeamsFromQuery(rows); if apiErr != nil {
		return getTeams_Success{}, apiErr
	}
	return getTeams_Success{teams}, nil
}

func getTeamsFromQuery(rows *sql.Rows) ([]Team, *conf.ApiResponse) {
	defer rows.Close()
	var teams []Team
	var team Team
	for rows.Next() {
		err := rows.Scan(&team.Id, &team.Name); if err != nil {
			return teams, conf.ErrDatabaseQueryFailed
		}
		teamLeader, apiErr := GetTeamLeader(team.Id); if apiErr != nil {
			return teams, apiErr
		}
		team.Leader = teamLeader
		participants, apiErr := GetTeamParticipants(team.Id); if apiErr != nil {
			return teams, apiErr
		}
		team.Participants = participants
		teams = append(teams, team)
	}
	if teams == nil {
		return make([]Team, 0), nil
	}
	return teams, nil
}

func GetTeamLeader(id int64) (TeamLeader, *conf.ApiResponse) {
	rows, err := src.CustomConnection.Query("SELECT id,name,surname,middlename FROM users " +
		"WHERE team=? AND access='1' LIMIT 1", id)
	if err != nil {
		return TeamLeader{}, conf.ErrDatabaseQueryFailed
	}
	defer rows.Close()
	var teamLeader TeamLeader
	for rows.Next() {
		err = rows.Scan(&teamLeader.ID, &teamLeader.Name, &teamLeader.Surname, &teamLeader.Middlename); if err != nil {
			return TeamLeader{}, conf.ErrDatabaseQueryFailed
		}
	}
	return teamLeader, nil
}

func GetTeamParticipants(id int64) ([]int64, *conf.ApiResponse) {
	rows, err := src.CustomConnection.Query("SELECT id FROM users WHERE team=? AND access='0'", id)
	if err != nil {
		return nil, conf.ErrDatabaseQueryFailed
	}
	defer rows.Close()
	var participant_id int64
	var participantIDs []int64
	for rows.Next() {
		err = rows.Scan(&participant_id)
		if err != nil {
			return nil, conf.ErrDatabaseQueryFailed
		}
		participantIDs = append(participantIDs, participant_id)
	}
	if participantIDs == nil {
		participantIDs = make([]int64, 0)
	}
	return participantIDs, nil
}


