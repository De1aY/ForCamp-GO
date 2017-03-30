package orgset

import (
	"net/http"
	"forcamp/src/authorization"
	"log"
	"forcamp/conf"
	"forcamp/src"
	"database/sql"
	"encoding/json"
	"fmt"
)

type TeamLeader struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
	Middlename string `json:"middlename"`
	Login string `json:"login"`
}

type Team struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Leader TeamLeader `json:"leader"`
	Count int `json:"count"`
}

type GetTeams_Success struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Teams []Team `json:"teams"`
}

func GetTeams(token string, ResponseWriter http.ResponseWriter) bool{
	if authorization.CheckTokenForEmpty(token, ResponseWriter) && authorization.CheckToken(token, ResponseWriter){
		Organization, _, err := getUserOrganizationAndLoginByToken(token)
		if err != nil{
			log.Print(err)
			return conf.PrintError(err, ResponseWriter)
		}
		NewConnection = src.Connect_Custom(Organization)
		Resp, APIerr := getTeams_Request()
		if APIerr != nil {
			return conf.PrintError(APIerr, ResponseWriter)
		}
		Response, _ := json.Marshal(Resp)
		fmt.Fprintf(ResponseWriter, string(Response))
	}
	return true
}

func getTeams_Request() (GetTeams_Success, *conf.ApiError){
	Query, err := NewConnection.Query("SELECT * FROM teams")
	if err != nil {
		log.Print(err)
		return GetTeams_Success{}, conf.ErrDatabaseQueryFailed
	}
	Teams, APIerr := getTeamsFromQuery(Query)
	if APIerr != nil {
		return  GetTeams_Success{}, APIerr
	}
	return GetTeams_Success{200, "success", Teams}, nil
}

func getTeamsFromQuery(rows *sql.Rows) ([]Team, *conf.ApiError){
	defer rows.Close()
	var Teams []Team
	var team Team
	for rows.Next(){
		err := rows.Scan(&team.Id, &team.Name)
		if err != nil{
			log.Print(err)
			return []Team{}, conf.ErrDatabaseQueryFailed
		}
		Leader , APIerr := getTeamLeader(team.Id)
		if APIerr != nil {
			return []Team{}, APIerr
		}
		team.Leader = Leader
		Count, APIerr := getTeamCount(team.Id)
		if APIerr != nil {
			return []Team{}, APIerr
		}
		team.Count = Count
		Teams = append(Teams, team)
	}
	return Teams, nil
}

func getTeamLeader(id int64) (TeamLeader, *conf.ApiError){
	Query, err := NewConnection.Query("SELECT login,name,surname,middlename FROM users WHERE team=? AND access='2' LIMIT 1", id)
	if err != nil {
		log.Print(err)
		return TeamLeader{}, conf.ErrDatabaseQueryFailed
	}
	defer Query.Close()
	var Leader TeamLeader
	for Query.Next(){
		err = Query.Scan(&Leader.Login, &Leader.Name, &Leader.Surname, &Leader.Middlename)
		if err != nil {
			log.Print(err)
			return TeamLeader{}, conf.ErrDatabaseQueryFailed
		}
	}
	return Leader, nil
}

func getTeamCount(id int64) (int, *conf.ApiError){
	Query, err := NewConnection.Query("SELECT COUNT(login) as count FROM users WHERE team=? AND access='0'", id)
	if err != nil {
		log.Print(err)
		return 0, conf.ErrDatabaseQueryFailed
	}
	defer Query.Close()
	var count int
	for Query.Next(){
		err = Query.Scan(&count)
		if err != nil{
			log.Print(err)
			return 0, conf.ErrDatabaseQueryFailed
		}
	}
	return count, nil
}


