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

type GetTeams_Success struct {
	Code   int `json:"code"`
	Status string `json:"status"`
	Teams  []Team `json:"teams"`
}

func GetTeams(token string, ResponseWriter http.ResponseWriter) bool {
	Connection := src.Connect()
	defer Connection.Close()
	if authorization.CheckTokenForEmpty(token, ResponseWriter) {
		if authorization.CheckToken(token, Connection, ResponseWriter) {
			Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token, Connection)
			if APIerr != nil{
				return conf.PrintError(APIerr, ResponseWriter)
			}
			NewConnection := src.Connect_Custom(Organization)
			defer NewConnection.Close()
			Resp, APIerr := getTeams_Request(NewConnection)
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

func getTeams_Request(Connection *sql.DB) (GetTeams_Success, *conf.ApiError) {
	Query, err := Connection.Query("SELECT * FROM teams")
	if err != nil {
		log.Print(err)
		return GetTeams_Success{}, conf.ErrDatabaseQueryFailed
	}
	Teams, APIerr := getTeamsFromQuery(Query, Connection)
	if APIerr != nil {
		return GetTeams_Success{}, APIerr
	}
	return GetTeams_Success{200, "success", Teams}, nil
}

func getTeamsFromQuery(rows *sql.Rows, Connection *sql.DB) ([]Team, *conf.ApiError) {
	defer rows.Close()
	var Teams []Team
	var team Team
	for rows.Next() {
		err := rows.Scan(&team.Id, &team.Name)
		if err != nil {
			log.Print(err)
			return []Team{}, conf.ErrDatabaseQueryFailed
		}
		Leader, APIerr := getTeamLeader(team.Id, Connection)
		if APIerr != nil {
			return []Team{}, APIerr
		}
		team.Leader = Leader
		participants, APIerr := getTeamParticipants(team.Id, Connection)
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
	return Teams, nil
}

func getTeamLeader(id int64, Connection *sql.DB) (TeamLeader, *conf.ApiError) {
	Query, err := Connection.Query("SELECT login,name,surname,middlename FROM users WHERE team=? AND access='1' LIMIT 1", id)
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

func getTeamParticipants(id int64, Connection *sql.DB) ([]string, *conf.ApiError) {
	Query, err := Connection.Query("SELECT login FROM users WHERE team=? AND access='0'", id)
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


