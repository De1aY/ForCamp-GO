package teams

import (
	"net/http"
	"forcamp/src/api/authorization"
	"log"
	"forcamp/conf"
	"forcamp/src"
	"database/sql"
	"forcamp/src/api/orgset"
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
	Teams  []Team `json:"teams"`
}

func GetTeams(token string, responseWriter http.ResponseWriter) bool {
	if authorization.CheckTokenForEmpty(token, responseWriter) {
		if authorization.CheckToken(token, responseWriter) {
			Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
			if APIerr != nil{
				return APIerr.Print(responseWriter)
			}
			src.CustomConnection = src.Connect_Custom(Organization)
			rawResp, APIerr := getTeams_Request()
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

func getTeams_Request() (getTeams_Success, *conf.ApiResponse) {
	Query, err := src.CustomConnection.Query("SELECT * FROM teams")
	if err != nil {
		log.Print(err)
		return getTeams_Success{}, conf.ErrDatabaseQueryFailed
	}
	Teams, APIerr := getTeamsFromQuery(Query)
	if APIerr != nil {
		return getTeams_Success{}, APIerr
	}
	return getTeams_Success{Teams}, nil
}

func getTeamsFromQuery(rows *sql.Rows) ([]Team, *conf.ApiResponse) {
	defer rows.Close()
	var Teams []Team
	var team Team
	for rows.Next() {
		err := rows.Scan(&team.Id, &team.Name)
		if err != nil {
			log.Print(err)
			return Teams, conf.ErrDatabaseQueryFailed
		}
		Leader, APIerr := GetTeamLeader(team.Id)
		if APIerr != nil {
			return Teams, APIerr
		}
		team.Leader = Leader
		participants, APIerr := GetTeamParticipants(team.Id)
		if APIerr != nil {
			return Teams, APIerr
		}
		team.Participants = participants
		Teams = append(Teams, team)
	}
	if Teams == nil {
		return make([]Team, 0), nil
	}
	return Teams, nil
}

func GetTeamLeader(id int64) (TeamLeader, *conf.ApiResponse) {
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

func GetTeamParticipants(id int64) ([]string, *conf.ApiResponse) {
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
	if logins == nil {
		logins = make([]string, 0)
	}
	return logins, nil
}


