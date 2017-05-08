package participants

import (
	"net/http"
	"forcamp/src/api/authorization"
	"forcamp/src/api/orgset"
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

type getParticipants_Success struct {
	Code         int `json:"code"`
	Status       string `json:"status"`
	Participants []Participant `json:"participants"`
}

func (success *getParticipants_Success) toJSON() string {
	resp, _ := json.Marshal(success)
	return string(resp)
}


func GetParticipants(token string, ResponseWriter http.ResponseWriter) bool {
	if authorization.CheckTokenForEmpty(token, ResponseWriter) {
		if authorization.CheckToken(token, ResponseWriter) {
			Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			src.CustomConnection = src.Connect_Custom(Organization)
			resp, APIerr := getParticipants_Request()
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

func getParticipants_Request() (getParticipants_Success, *conf.ApiError) {
	Query, err := src.CustomConnection.Query("SELECT login,name,surname,middlename,sex,team FROM users WHERE access='0'")
	if err != nil {
		log.Print(err)
		return getParticipants_Success{}, conf.ErrDatabaseQueryFailed
	}
	return getParticipantsFromResponse(Query)
}

func getParticipantsFromResponse(rows *sql.Rows) (getParticipants_Success, *conf.ApiError) {
	defer rows.Close()
	marks, APIerr := getMarks()
	if APIerr != nil {
		return getParticipants_Success{}, APIerr
	}
	var (
		participants []Participant
		participant Participant
	)
	for rows.Next() {
		err := rows.Scan(&participant.Login, &participant.Name, &participant.Surname, &participant.Middlename, &participant.Sex, &participant.Team)
		if err != nil {
			log.Print(err)
			return getParticipants_Success{}, conf.ErrDatabaseQueryFailed
		}
		participant.Marks = marks[participant.Login]
		participants = append(participants, Participant{participant.Login, participant.Name, participant.Surname, participant.Middlename, participant.Sex, participant.Team, participant.Marks})
	}
	if participants == nil {
		return getParticipants_Success{200, "success", make([]Participant, 0)}, nil
	}
	return getParticipants_Success{200, "success", participants}, nil
}

func getMarks() (map[string][]Mark, *conf.ApiError) {
	Query, err := src.CustomConnection.Query("SELECT * FROM participants")
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