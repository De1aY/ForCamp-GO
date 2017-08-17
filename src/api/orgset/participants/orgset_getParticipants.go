package participants

import (
	"net/http"
	"forcamp/src/api/authorization"
	"forcamp/src/api/orgset"
	"forcamp/conf"
	"forcamp/src"
	"database/sql"
	"strconv"
	"forcamp/src/api/orgset/categories"
)

type Mark struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

type Participant struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Middlename string `json:"middlename"`
	Sex        int    `json:"sex"`
	Team       int64  `json:"team"`
	Marks      []Mark `json:"marks"`
}

type getParticipants_Success struct {
	Participants []Participant `json:"participants"`
}

func GetParticipants(token string, responseWriter http.ResponseWriter) bool {
	if authorization.IsTokenNotEmpty(token, responseWriter) {
		if authorization.IsTokenValid(token, responseWriter) {
			organizationName, _, apiErr := orgset.GetUserOrganizationAndIdByToken(token); if apiErr != nil {
				return apiErr.Print(responseWriter)
			}
			src.CustomConnection = src.Connect_Custom(organizationName)
			rawResp, apiErr := getParticipants(); if apiErr != nil {
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

func getParticipants() (getParticipants_Success, *conf.ApiResponse) {
	rows, err := src.CustomConnection.Query("SELECT id,name,surname,middlename,sex,team " +
		"FROM users WHERE access='0'")
	if err != nil {
		return getParticipants_Success{}, conf.ErrDatabaseQueryFailed
	}
	return getParticipantsFromResponse(rows)
}

func getParticipantsFromResponse(rows *sql.Rows) (getParticipants_Success, *conf.ApiResponse) {
	defer rows.Close()
	marks, apiErr := getMarks(); if apiErr != nil {
		return getParticipants_Success{}, apiErr
	}
	var (
		participants []Participant
		participant  Participant
	)
	for rows.Next() {
		err := rows.Scan(&participant.ID, &participant.Name, &participant.Surname,
			&participant.Middlename, &participant.Sex, &participant.Team)
		if err != nil {
			return getParticipants_Success{}, conf.ErrDatabaseQueryFailed
		}
		participant.Marks = marks[participant.ID]
		participants = append(participants, Participant{participant.ID, participant.Name,
			participant.Surname, participant.Middlename, participant.Sex,
			participant.Team, participant.Marks})
	}
	if participants == nil {
		return getParticipants_Success{make([]Participant, 0)}, nil
	}
	return getParticipants_Success{participants}, nil
}

func getMarks() (map[int64][]Mark, *conf.ApiResponse) {
	rows, err := src.CustomConnection.Query("SELECT * FROM participants")
	if err != nil {
		return make(map[int64][]Mark), conf.ErrDatabaseQueryFailed
	}
	CategoriesIDs, err := rows.Columns()
	if err != nil {
		return make(map[int64][]Mark), conf.ErrDatabaseQueryFailed
	}
	if len(CategoriesIDs) == 1 {
		return getMarksIfNoCategories(rows)
	}
	return getMarksIfCategories(rows, CategoriesIDs)
}

func getMarksIfNoCategories(rows *sql.Rows) (map[int64][]Mark, *conf.ApiResponse) {
	defer rows.Close()
	var (
		participant_id int64
		marks = make(map[int64][]Mark)
	)
	for rows.Next() {
		err := rows.Scan(&participant_id)
		if err != nil {
			return marks, conf.ErrDatabaseQueryFailed
		}
		marks[participant_id] = make([]Mark, 0)
	}
	return marks, nil
}

func getMarksIfCategories(rows *sql.Rows, categoriesIDs []string) (map[int64][]Mark, *conf.ApiResponse) {
	categoriesIDs = categoriesIDs[1:]
	var (
		rawResult = make([][]byte, len(categoriesIDs)+1)
		result    = make([]interface{}, len(categoriesIDs)+1)
		marks     = make(map[int64][]Mark)
		values    = make([]string, len(categoriesIDs)+1)
	)
	for i, _ := range result {
		result[i] = &rawResult[i]
	}
	categoriesList, apiErr := categories.GetCategories_Request();
	if apiErr != nil {
		return marks, apiErr
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(result...)
		if err != nil {
			return marks, conf.ErrDatabaseQueryFailed
		}
		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				values[i] = string(raw)
			}
		}
		participant_id, err := strconv.ParseInt(values[0], 10, 64); if err != nil {
			return marks, conf.ErrConvertStringToInt
		}
		categoriesValues := values[1:]
		for i := 0; i < len(categoriesValues); i++ {
			id, err := strconv.ParseInt(categoriesIDs[i], 10, 64); if err != nil {
				return marks, conf.ErrConvertStringToInt
			}
			value, err := strconv.ParseInt(categoriesValues[i], 10, 64); if err != nil {
				return marks, conf.ErrConvertStringToInt
			}
			marks[participant_id] = append(marks[participant_id], Mark{id, categoriesList[i].Name, value})
		}
	}
	return marks, nil
}
