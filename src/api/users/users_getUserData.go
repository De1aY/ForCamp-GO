package users

import (
	"forcamp/src/api/authorization"
	"forcamp/src"
	"net/http"
	"forcamp/conf"
	"database/sql"
	"forcamp/src/api/orgset"
	"strconv"
	"forcamp/src/api/orgset/participants"
	"forcamp/src/api/orgset/employees"
	"forcamp/src/api/orgset/categories"
	"forcamp/src/api/orgset/teams"
	"forcamp/src/api/orgset/events"
)

func GetUserData(token string, responseWriter http.ResponseWriter, user_id int64) bool {
	if authorization.IsTokenValid(token, responseWriter) {
		organizationName, userLogin, apiErr := orgset.GetUserOrganizationAndIdByToken(token)
		if apiErr != nil {
			return apiErr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(organizationName)
		if user_id != -1 {
			user_organization, apiErr := orgset.GetUserOrganizationByID(user_id)
			if apiErr != nil {
				return apiErr.Print(responseWriter)
			}
			if user_organization != organizationName {
				return conf.ErrUserNotFound.Print(responseWriter)
			}
		} else {
			user_id = userLogin
		}
		userData, apiErr := GetUserData_Request(user_id)
		if apiErr != nil {
			return apiErr.Print(responseWriter)
		}
		userData.Organization = organizationName
		rawResp := getUserData_Success{userData}
		resp := &conf.ApiResponse{200, "success", rawResp}
		resp.Print(responseWriter)

	} else {
		return conf.ErrUserTokenIncorrect.Print(responseWriter)
	}
	return true
}

func GetUserData_Request(user_id int64) (UserData, *conf.ApiResponse) {
	rows, err := src.CustomConnection.Query("SELECT name, surname, middlename, sex, access, avatar, team " +
		"FROM users WHERE id=?", user_id)
	if err != nil {
		return UserData{}, conf.ErrDatabaseQueryFailed
	}
	userData, apiErr := getUserDataFromQuery(rows, user_id)
	if apiErr != nil {
		return UserData{}, apiErr
	}
	return userData, nil
}

func getUserDataFromQuery(rows *sql.Rows, user_id int64) (UserData, *conf.ApiResponse) {
	defer rows.Close()
	var userData UserData
	var teamID int64
	for rows.Next() {
		err := rows.Scan(&userData.Name, &userData.Surname, &userData.Middlename, &userData.Sex, &userData.Access, &userData.Avatar, &teamID)
		if err != nil {
			return userData, conf.ErrDatabaseQueryFailed
		}
	}
	user_events, apiErr := events.GetEvents_Request(user_id, 10, 0, false, -1); if apiErr != nil {
		return userData, apiErr
	}
	teamInfo, apiErr := getTeamInfo(teamID); if apiErr != nil {
		return userData, apiErr
	}
	if userData.Access == 0 {
		marksData, apiErr := getMarks(user_id)
		if apiErr != nil {
			return UserData{}, apiErr
		}
		var orgSettings_Participant string
		err := src.CustomConnection.QueryRow("SELECT value FROM settings WHERE name='participant'").Scan(&orgSettings_Participant)
		if err != nil {
			return UserData{}, conf.ErrDatabaseQueryFailed
		}
		return UserData{Name: userData.Name,
			Surname:          userData.Surname,
			Middlename:       userData.Middlename,
			Sex:              userData.Sex,
			Team:             teamInfo,
			Access:           userData.Access,
			Avatar:           userData.Avatar,
			Post:             orgSettings_Participant,
			Events:           user_events,
			AdditionalData:   marksData}, nil
	} else {
		permissions, post, APIerr := getPermissionsAndPost(user_id)
		if APIerr != nil {
			return UserData{}, APIerr
		}
		return UserData{Name: userData.Name,
			Surname: userData.Surname,
			Middlename: userData.Middlename,
			Sex: userData.Sex,
			Team: teamInfo,
			Access: userData.Access,
			Avatar: userData.Avatar,
			Post: post,
			Events: user_events,
			AdditionalData: permissions}, nil
	}
}

func getTeamInfo(teamID int64) (teams.Team, *conf.ApiResponse){
	var teamInfo teams.Team
	rows, err := src.CustomConnection.Query("SELECT * FROM teams WHERE id=?", teamID); if err != nil {
		return teamInfo, conf.ErrDatabaseQueryFailed
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&teamInfo.Id, &teamInfo.Name); if err != nil {
			return teamInfo, conf.ErrDatabaseQueryFailed
		}
		leader, apiErr := teams.GetTeamLeader(teamID); if apiErr != nil {
			return teamInfo, apiErr
		}
		participantsData, apiErr := teams.GetTeamParticipants(teamID); if apiErr != nil {
			return teamInfo, apiErr
		}
		teamInfo.Leader = leader
		teamInfo.Participants = participantsData
	}
	return teamInfo, nil
}

func getMarks(id int64) ([]participants.Mark, *conf.ApiResponse) {
	query, err := src.CustomConnection.Query("SELECT * FROM participants WHERE id=?", id)
	if err != nil {
		return nil, conf.ErrDatabaseQueryFailed
	}
	categoriesIDs, err := query.Columns()
	if err != nil {
		return nil, conf.ErrDatabaseQueryFailed
	}
	if len(categoriesIDs) == 1 {
		return getMarksIfNoCategories(query)
	}
	return getMarksIfCategories(query, categoriesIDs)
}

func getMarksIfNoCategories(rows *sql.Rows) ([]participants.Mark, *conf.ApiResponse) {
	defer rows.Close()
	var (
		id     int64
		marksData []participants.Mark
	)
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return nil, conf.ErrDatabaseQueryFailed
		}
		marksData = make([]participants.Mark, 0)
	}
	return marksData, nil
}

func getMarksIfCategories(rows *sql.Rows, categoriesIDs []string) ([]participants.Mark, *conf.ApiResponse) {
	categoriesIDs = categoriesIDs[1:]
	var (
		marksData  []participants.Mark
		rawResult  = make([][]byte, len(categoriesIDs) + 1)
		result     = make([]interface{}, len(categoriesIDs) + 1)
		values     = make([]string, len(categoriesIDs) + 1)
	)
	for i, _ := range result {
		result[i] = &rawResult[i]
	}
	categoriesList, apiErr := categories.GetCategories_Request(); if apiErr != nil {
		return nil, apiErr
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(result...)
		if err != nil {
			return nil, conf.ErrDatabaseQueryFailed
		}
		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				values[i] = string(raw)
			}
		}
		categoriesValues := values[1:]
		for i := 0; i < len(categoriesValues); i++ {
			id, err := strconv.ParseInt(categoriesIDs[i], 10, 64)
			if err != nil {
				return nil, conf.ErrConvertStringToInt
			}
			value, err := strconv.ParseInt(categoriesValues[i], 10, 64)
			if err != nil {
				return nil, conf.ErrConvertStringToInt
			}
			marksData = append(marksData, participants.Mark{id, categoriesList[i].Name, value})
		}
	}
	return marksData, nil
}

func getPermissionsAndPost(id int64) ([]employees.Permission, string, *conf.ApiResponse) {
	query, err := src.CustomConnection.Query("SELECT * FROM employees WHERE id=?", id)
	if err != nil {
		return nil, "", conf.ErrDatabaseQueryFailed
	}
	categoriesIDs, err := query.Columns()
	if err != nil {
		return nil, "", conf.ErrDatabaseQueryFailed
	}
	if len(categoriesIDs) == 1 {
		return getPermissionsIfNoCategories(query)
	}
	return getPermissionsIfCategories(query, categoriesIDs)
}

func getPermissionsIfNoCategories(rows *sql.Rows) ([]employees.Permission, string, *conf.ApiResponse) {
	defer rows.Close()
	var (
		login       string
		post        string
		permissions []employees.Permission
	)
	for rows.Next() {
		err := rows.Scan(&login, &post)
		if err != nil {
			return nil, "", conf.ErrDatabaseQueryFailed
		}
		permissions = make([]employees.Permission, 0)
	}
	return permissions, post, nil
}

func getPermissionsIfCategories(rows *sql.Rows, categoriesIDs []string) ([]employees.Permission, string, *conf.ApiResponse) {
	categoriesIDs = categoriesIDs[2:]
	var (
		rawResult = make([][]byte, len(categoriesIDs) + 2)
		result = make([]interface{}, len(categoriesIDs) + 2)
		permissions []employees.Permission
		values = make([]string, len(categoriesIDs) + 2)
		post string
	)
	for i, _ := range result {
		result[i] = &rawResult[i]
	}
	categoriesList, apiErr := categories.GetCategories_Request(); if apiErr != nil {
		return nil, "", apiErr
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(result...)
		if err != nil {
			return nil, "", conf.ErrDatabaseQueryFailed
		}
		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				values[i] = string(raw)
			}
		}
		post = values[1]
		values = values[2:]
		for i := 0; i < len(values); i++ {
			id, err := strconv.ParseInt(categoriesIDs[i], 10, 64)
			if err != nil {
				return nil, "", conf.ErrConvertStringToInt
			}
			permissions = append(permissions, employees.Permission{id, categoriesList[i].Name, values[i]})
		}
		values = make([]string, len(categoriesIDs) + 2)
	}
	return permissions, post, nil
}
