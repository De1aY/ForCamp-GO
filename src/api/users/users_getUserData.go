package users

import (
	"forcamp/src/api/authorization"
	"forcamp/src"
	"net/http"
	"forcamp/conf"
	"database/sql"
	"forcamp/src/api/orgset"
	"log"
	"strconv"
	"forcamp/src/api/orgset/participants"
	"forcamp/src/api/orgset/employees"
	"forcamp/src/api/orgset/categories"
	"forcamp/src/api/marks"
	"forcamp/src/api/orgset/teams"
)

func GetUserData(Token string, responseWriter http.ResponseWriter, login string) bool {
	if checkData(Token, login, responseWriter) {
		if authorization.CheckToken(Token, responseWriter) {
			Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(Token)
			if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			src.CustomConnection = src.Connect_Custom(Organization)
			UserOrganization, APIerr := orgset.GetUserOrganizationByLogin(login)
			if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			if UserOrganization != Organization {
				return conf.ErrUserNotFound.Print(responseWriter)
			}
			userData, APIerr := GetUserData_Request(login)
			if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			userData.Organization = Organization
			rawResp := getUserData_Success{userData}
			resp := &conf.ApiResponse{200, "success", rawResp}
			resp.Print(responseWriter)

		} else {
			return conf.ErrUserTokenIncorrect.Print(responseWriter)
		}
	}
	return true
}

func GetUserData_Request(login string) (UserData, *conf.ApiResponse) {
	Query, err := src.CustomConnection.Query("SELECT name, surname, middlename, sex, access, avatar, team FROM users WHERE login=?", login)
	if err != nil {
		log.Print(err)
		return UserData{}, conf.ErrDatabaseQueryFailed
	}
	userData, APIerr := getUserDataFromQuery(Query, login)
	if APIerr != nil {
		return UserData{}, APIerr
	}
	return userData, nil
}

func getUserDataFromQuery(rows *sql.Rows, login string) (UserData, *conf.ApiResponse) {
	defer rows.Close()
	var userData UserData
	var teamID int64
	for rows.Next() {
		err := rows.Scan(&userData.Name, &userData.Surname, &userData.Middlename, &userData.Sex, &userData.Access, &userData.Avatar, &teamID)
		if err != nil {
			log.Print(err)
			return UserData{}, conf.ErrDatabaseQueryFailed
		}
	}
	actions, apiErr := marks.GetMarksChanges_Request(login); if apiErr != nil {
		return UserData{}, apiErr
	}
	teamInfo, apiErr := getTeamInfo(teamID); if apiErr != nil {
		return userData, apiErr
	}
	if userData.Access == 0 {
		marks, APIerr := getMarks(login)
		if APIerr != nil {
			return UserData{}, APIerr
		}
		var orgSettings_Participant string
		err := src.CustomConnection.QueryRow("SELECT value FROM settings WHERE name='participant'").Scan(&orgSettings_Participant)
		if err != nil {
			return UserData{}, conf.ErrDatabaseQueryFailed
		}
		return UserData{Name: userData.Name,
			Surname: userData.Surname,
			Middlename: userData.Middlename,
			Sex: userData.Sex,
			Team: teamInfo,
			Access: userData.Access,
			Avatar: userData.Avatar,
			Post: orgSettings_Participant,
			Actions: actions,
			AdditionalData: marks}, nil
	} else {
		permissions, post, APIerr := getPermissionsAndPost(login)
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
			Actions: actions,
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
		participants, apiErr := teams.GetTeamParticipants(teamID); if apiErr != nil {
			return teamInfo, apiErr
		}
		teamInfo.Leader = leader
		teamInfo.Participants = participants
	}
	return teamInfo, nil
}

func getMarks(login string) ([]participants.Mark, *conf.ApiResponse) {
	Query, err := src.CustomConnection.Query("SELECT * FROM participants WHERE login=?", login)
	if err != nil {
		log.Print(err)
		return nil, conf.ErrDatabaseQueryFailed
	}
	CategoriesIDs, err := Query.Columns()
	if err != nil {
		log.Print(err)
		return nil, conf.ErrDatabaseQueryFailed
	}
	if len(CategoriesIDs) == 1 {
		return getMarksIfNoCategories(Query)
	}
	return getMarksIfCategories(Query, CategoriesIDs)
}

func getMarksIfNoCategories(rows *sql.Rows) ([]participants.Mark, *conf.ApiResponse) {
	defer rows.Close()
	var (
		login string
		marks []participants.Mark
	)
	for rows.Next() {
		err := rows.Scan(&login)
		if err != nil {
			log.Print(err)
			return nil, conf.ErrDatabaseQueryFailed
		}
		marks = make([]participants.Mark, 0)
	}
	return marks, nil
}

func getMarksIfCategories(rows *sql.Rows, categoriesIDs []string) ([]participants.Mark, *conf.ApiResponse) {
	categoriesIDs = categoriesIDs[1:]
	var (
		rawResult = make([][]byte, len(categoriesIDs) + 1)
		Result = make([]interface{}, len(categoriesIDs) + 1)
		marks []participants.Mark
		Values = make([]string, len(categoriesIDs) + 1)
	)
	for i, _ := range Result {
		Result[i] = &rawResult[i]
	}
	categoriesList, apiErr := categories.GetCategories_Request(); if apiErr != nil {
		return nil, apiErr
	}
	defer rows.Close()
	for rows.Next() {

		err := rows.Scan(Result...)
		if err != nil {
			log.Print(err)
			return nil, conf.ErrDatabaseQueryFailed
		}
		for i, raw := range rawResult {
			if raw == nil {
				Result[i] = "\\N"
			} else {
				Values[i] = string(raw)
			}
		}
		categoriesValues := Values[1:]
		for i := 0; i < len(categoriesValues); i++ {
			id, err := strconv.ParseInt(categoriesIDs[i], 10, 64)
			if err != nil {
				log.Print(err)
				return nil, conf.ErrConvertStringToInt
			}
			value, err := strconv.ParseInt(categoriesValues[i], 10, 64)
			if err != nil {
				log.Print(err)
				return nil, conf.ErrConvertStringToInt
			}
			marks = append(marks, participants.Mark{id, categoriesList[i].Name, value})
		}
	}
	return marks, nil
}

func getPermissionsAndPost(login string) ([]employees.Permission, string, *conf.ApiResponse) {
	Query, err := src.CustomConnection.Query("SELECT * FROM employees WHERE login=?", login)
	if err != nil {
		log.Print(err)
		return nil, "", conf.ErrDatabaseQueryFailed
	}
	CategoriesIDs, err := Query.Columns()
	if err != nil {
		log.Print(err)
		return nil, "", conf.ErrDatabaseQueryFailed
	}
	if len(CategoriesIDs) == 1 {
		return getPermissionsIfNoCategories(Query)
	}
	return getPermissionsIfCategories(Query, CategoriesIDs)
}

func getPermissionsIfNoCategories(rows *sql.Rows) ([]employees.Permission, string, *conf.ApiResponse) {
	defer rows.Close()
	var (
		login string
		post string
		Permissions []employees.Permission
	)
	for rows.Next() {
		err := rows.Scan(&login, &post)
		if err != nil {
			log.Print(err)
			return nil, "", conf.ErrDatabaseQueryFailed
		}
		Permissions = make([]employees.Permission, 0)
	}
	return Permissions, post, nil
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
			log.Print(err)
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
				log.Print(err)
				return nil, "", conf.ErrConvertStringToInt
			}
			permissions = append(permissions, employees.Permission{id, categoriesList[i].Name, values[i]})
		}
		values = make([]string, len(categoriesIDs) + 2)
	}
	return permissions, post, nil
}
