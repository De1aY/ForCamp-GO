package teams

import (
	"net/http"
	"forcamp/src/authorization"
	"log"
	"forcamp/conf"
	"forcamp/src"
	"database/sql"
	"encoding/json"
	"fmt"
	"forcamp/src/orgset"
)

type TeamLeader struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Middlename string `json:"middlename"`
	Login      string `json:"login"`
}

type Team struct {
	Id     int64 `json:"id"`
	Name   string `json:"name"`
	Leader TeamLeader `json:"leader"`
	Participants  []string `json:"participants"`
}

type getTeams_Success struct {
	Code   int `json:"code"`
	Status string `json:"status"`
	Teams  []Team `json:"teams"`
}

func (success *getTeams_Success) toJSON() string {
	resp, _ := json.Marshal(success)
	return string(resp)
}

// =====================================================

func GetTeams(token string, ResponseWriter http.ResponseWriter) bool {
	if authorization.CheckTokenForEmpty(token, ResponseWriter) {
		if authorization.CheckToken(token, ResponseWriter) {
			Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
			if APIerr != nil{
				return conf.PrintError(APIerr, ResponseWriter)
			}
			src.CustomConnection = src.Connect_Custom(Organization)
			resp, APIerr := getTeams_Request()
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

func getTeams_Request() (getTeams_Success, *conf.ApiError) {
	Query, err := src.CustomConnection.Query("SELECT * FROM teams")
	if err != nil {
		log.Print(err)
		return getTeams_Success{}, conf.ErrDatabaseQueryFailed
	}
	Teams, APIerr := getTeamsFromQuery(Query)
	if APIerr != nil {
		return getTeams_Success{}, APIerr
	}
	return getTeams_Success{200, "success", Teams}, nil
}

func getTeamsFromQuery(rows *sql.Rows) ([]Team, *conf.ApiError) {
	defer rows.Close()
	var Teams []Team
	var team Team
	for rows.Next() {
		err := rows.Scan(&team.Id, &team.Name)
		if err != nil {
			log.Print(err)
			return []Team{}, conf.ErrDatabaseQueryFailed
		}
		Leader, APIerr := getTeamLeader(team.Id)
		if APIerr != nil {
			return []Team{}, APIerr
		}
		team.Leader = Leader
		participants, APIerr := getTeamParticipants(team.Id)
		if APIerr != nil {
			return []Team{}, APIerr
		}
		if participants == nil {
			team.Participants = make([]string, 0)
		} else {
			team.Participants = participants
		}
		Teams = append(Teams, team)
	}
	if Teams == nil {
		return make([]Team, 0), nil
	}
	return Teams, nil
}

func getTeamLeader(id int64) (TeamLeader, *conf.ApiError) {
	Query, err := src.CustomConnection.Query("SELECT login,name,surname,middlename FROM users WHERE team=? AND access='1' LIMIT 1", id)
	if err != nil {
		log.Print(err)
		return TeamLeader{}, conf.ErrDatabaseQueryFailed
	}
	defer Query.Close()
	var Leader TeamLeader
	for Query.Next() {
		err = Query.Scan(&Leader.Login, &Leader.Name, &Leader.Surname, &Leader.Middlename)
		if err != nil {
			log.Print(err)
			return TeamLeader{}, conf.ErrDatabaseQueryFailed
		}
	}
	return Leader, nil
}

func getTeamParticipants(id int64) ([]string, *conf.ApiError) {
	Query, err := src.CustomConnection.Query("SELECT login FROM users WHERE team=? AND access='0'", id)
	if err != nil {
		log.Print(err)
		return nil, conf.ErrDatabaseQueryFailed
	}
	defer Query.Close()
	var login string
	var logins []string
	for Query.Next() {
		err = Query.Scan(&login)
		if err != nil {
			log.Print(err)
			return nil, conf.ErrDatabaseQueryFailed
		}
		logins = append(logins, login)
	}
	return logins, nil
}


