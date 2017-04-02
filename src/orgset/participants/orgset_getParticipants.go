package participants

import (
	"net/http"
	"forcamp/src/authorization"
	"forcamp/src/orgset"
	"forcamp/conf"
	"forcamp/src"
	"log"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
)

type Mark struct {
	Id    int64 `json:"id"`
	Value int64 `json:"value"`
}

type Participant struct {
	Login      string `json:"login"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Middlename string `json:"middlename"`
	Sex        int `json:"sex"`
	Team       int64 `json:"team"`
	Marks      []Mark `json:"marks"`
}

type GetParticipants_Success struct {
	Code         int `json:"code"`
	Status       string `json:"status"`
	Participants []Participant `json:"participants"`
}

func GetParticipants(token string, ResponseWriter http.ResponseWriter) bool {
	Connection := src.Connect()
	defer Connection.Close()
	if authorization.CheckTokenForEmpty(token, ResponseWriter) {
		if authorization.CheckToken(token, Connection, ResponseWriter){
			Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token, Connection)
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			NewConnection := src.Connect_Custom(Organization)
			defer NewConnection.Close()
			Resp, APIerr := getParticipants_Request(NewConnection)
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

func getParticipants_Request(Connection *sql.DB) (GetParticipants_Success, *conf.ApiError) {
	Query, err := Connection.Query("SELECT login,name,surname,middlename,sex,team FROM users WHERE access='0'")
	if err != nil {
		log.Print(err)
		return GetParticipants_Success{}, conf.ErrDatabaseQueryFailed
	}
	return getParticipantsFromResponse(Query, Connection)
}

func getParticipantsFromResponse(rows *sql.Rows, Connection *sql.DB) (GetParticipants_Success, *conf.ApiError) {
	defer rows.Close()
	marks, APIerr := getMarks(Connection)
	if APIerr != nil {
		return GetParticipants_Success{}, APIerr
	}
	var (
		participants []Participant
		participant Participant
	)
	for rows.Next() {
		err := rows.Scan(&participant.Login, &participant.Name, &participant.Surname, &participant.Middlename, &participant.Sex, &participant.Team)
		if err != nil {
			log.Print(err)
			return GetParticipants_Success{}, conf.ErrDatabaseQueryFailed
		}
		participant.Marks = marks[participant.Login]
		participants = append(participants, Participant{participant.Login, participant.Name, participant.Surname, participant.Middlename, participant.Sex, participant.Team, participant.Marks})
	}
	return GetParticipants_Success{200, "success", participants}, nil
}

func getMarks(Connection *sql.DB) (map[string][]Mark, *conf.ApiError) {
	Query, err := Connection.Query("SELECT * FROM participants")
	if err != nil {
		log.Print(err)
		return make(map[string][]Mark), conf.ErrDatabaseQueryFailed
	}
	CategoriesIDs, err := Query.Columns()
	if err != nil {
		log.Print(err)
		return make(map[string][]Mark), conf.ErrDatabaseQueryFailed
	}
	if len(CategoriesIDs) == 1 {
		return getMarksIfNoCategories(Query)
	}
	return getMarksIfCategories(Query, CategoriesIDs)
}

func getMarksIfNoCategories(rows *sql.Rows) (map[string][]Mark, *conf.ApiError) {
	defer rows.Close()
	var (
		login string
		marks = make(map[string][]Mark)
	)
	for rows.Next() {
		err := rows.Scan(&login)
		if err != nil {
			log.Print(err)
			return make(map[string][]Mark), conf.ErrDatabaseQueryFailed
		}
		marks[login] = make([]Mark, 0)
	}
	return marks, nil
}

func getMarksIfCategories(rows *sql.Rows, CategoriesIDs []string) (map[string][]Mark, *conf.ApiError) {
	CategoriesIDs = CategoriesIDs[1:]
	var (
		rawResult = make([][]byte, len(CategoriesIDs) + 1)
		Result = make([]interface{}, len(CategoriesIDs) + 1)
		Marks = make(map[string][]Mark)
		Values = make([]string, len(CategoriesIDs) + 1)
	)
	for i, _ := range Result {
		Result[i] = &rawResult[i]
	}
	defer rows.Close()
	for rows.Next() {

		err := rows.Scan(Result...)
		if err != nil {
			return make(map[string][]Mark), conf.ErrDatabaseQueryFailed
		}
		for i, raw := range rawResult {
			if raw == nil {
				Result[i] = "\\N"
			} else {
				Values[i] = string(raw)
			}
		}
		Login := Values[0]
		CategoriesValues := Values[1:]
		for i := 0; i < len(CategoriesValues); i++ {
			id, err := strconv.ParseInt(CategoriesIDs[i], 10, 64)
			if err != nil {
				log.Print(err)
				return make(map[string][]Mark), conf.ErrConvertStringToInt
			}
			value, err := strconv.ParseInt(CategoriesValues[i], 10, 64)
			if err != nil {
				log.Print(err)
				return make(map[string][]Mark), conf.ErrConvertStringToInt
			}
			Marks[Login] = append(Marks[Login], Mark{id, value})
		}
	}
	return Marks, nil
}