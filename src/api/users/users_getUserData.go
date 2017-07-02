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
			userData, APIerr := getUserData_Request(login)
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

func getUserData_Request(login string) (UserData, *conf.ApiResponse) {
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
	for rows.Next() {
		err := rows.Scan(&userData.Name, &userData.Surname, &userData.Middlename, &userData.Sex, &userData.Access, &userData.Avatar, &userData.Team)
		if err != nil {
			log.Print(err)
			return UserData{}, conf.ErrDatabaseQueryFailed
		}
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
			Team: userData.Team,
			Access: userData.Access,
			Avatar: userData.Avatar,
			Post: orgSettings_Participant,
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
			Team: userData.Team,
			Access: userData.Access,
			Avatar: userData.Avatar,
			Post: post,
			AdditionalData: permissions}, nil
	}
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

func getMarksIfCategories(rows *sql.Rows, CategoriesIDs []string) ([]participants.Mark, *conf.ApiResponse) {
	CategoriesIDs = CategoriesIDs[1:]
	var (
		rawResult = make([][]byte, len(CategoriesIDs) + 1)
		Result = make([]interface{}, len(CategoriesIDs) + 1)
		Marks []participants.Mark
		Values = make([]string, len(CategoriesIDs) + 1)
	)
	for i, _ := range Result {
		Result[i] = &rawResult[i]
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
		CategoriesValues := Values[1:]
		for i := 0; i < len(CategoriesValues); i++ {
			id, err := strconv.ParseInt(CategoriesIDs[i], 10, 64)
			if err != nil {
				log.Print(err)
				return nil, conf.ErrConvertStringToInt
			}
			value, err := strconv.ParseInt(CategoriesValues[i], 10, 64)
			if err != nil {
				log.Print(err)
				return nil, conf.ErrConvertStringToInt
			}
			Marks = append(Marks, participants.Mark{id, value})
		}
	}
	return Marks, nil
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

func getPermissionsIfCategories(rows *sql.Rows, CategoriesIDs []string) ([]employees.Permission, string, *conf.ApiResponse) {
	CategoriesIDs = CategoriesIDs[2:]
	var (
		rawResult = make([][]byte, len(CategoriesIDs) + 2)
		Result = make([]interface{}, len(CategoriesIDs) + 2)
		Permissions []employees.Permission
		Values = make([]string, len(CategoriesIDs) + 2)
		Post string
	)
	for i, _ := range Result {
		Result[i] = &rawResult[i]
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(Result...)
		if err != nil {
			log.Print(err)
			return nil, "", conf.ErrDatabaseQueryFailed
		}
		for i, raw := range rawResult {
			if raw == nil {
				Result[i] = "\\N"
			} else {
				Values[i] = string(raw)
			}
		}
		Post = Values[1]
		Values = Values[2:]
		for i := 0; i < len(Values); i++ {
			id, err := strconv.ParseInt(CategoriesIDs[i], 10, 64)
			if err != nil {
				log.Print(err)
				return nil, "", conf.ErrConvertStringToInt
			}
			Permissions = append(Permissions, employees.Permission{id, Values[i]})
		}
		Values = make([]string, len(CategoriesIDs) + 2)
	}
	return Permissions, Post, nil
}
