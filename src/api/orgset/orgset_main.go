/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package orgset

import (
	"wplay/conf"
	"wplay/src"
	"wplay/src/api/authorization"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func IsUserAdmin(token string, responseWriter http.ResponseWriter) bool {
	if authorization.IsTokenNotEmpty(token, responseWriter) {
		if authorization.IsTokenValid(token, responseWriter) {
			organizationName, id, apiErr := GetUserOrganizationAndIdByToken(token)
			if apiErr != nil {
				return apiErr.Print(responseWriter)
			}
			customConnection := src.Connect_Custom(organizationName)
			defer customConnection.Close()
			var access int
			err := customConnection.QueryRow("SELECT access FROM users WHERE id=?", id).Scan(&access)
			if err != nil {
				return conf.ErrDatabaseQueryFailed.Print(responseWriter)
			}
			if access == 2 {
				return true
			} else {
				return conf.ErrInsufficientRights.Print(responseWriter)
			}
		} else {
			return conf.ErrUserTokenIncorrect.Print(responseWriter)
		}
	}
	return false
}

func GetUserOrganizationAndIdByToken(token string) (string, int64, *conf.ApiResponse) {
	id, apiErr := GetUserIdByToken(token)
	if apiErr != nil {
		return "", -1, apiErr
	}
	organizationName, apiErr := GetUserOrganizationByID(id)
	if apiErr != nil {
		return "", -1, apiErr
	}
	return organizationName, id, nil
}

func GetUserOrganizationAndLoginByID(user_id int64) (string, string, *conf.ApiResponse) {
	user_login, apiErr := GetUserLoginByID(user_id)
	if apiErr != nil {
		return "", "", apiErr
	}
	organizationName, apiErr := GetUserOrganizationByID(user_id)
	if apiErr != nil {
		return "", "", apiErr
	}
	return organizationName, user_login, nil
}

func GetUserLoginByID(user_id int64) (string, *conf.ApiResponse) {
	var user_login string
	err := src.Connection.QueryRow("SELECT login FROM users WHERE id=?", user_id).Scan(&user_login)
	if err != nil {
		return user_login, conf.ErrDatabaseQueryFailed
	}
	return user_login, nil
}

func GeneratePassword() (string, string) {
	password := ""
	for len(password) < 7 {
		rand.Seed(time.Now().UnixNano())
		password = strconv.FormatInt(rand.Int63n(1000000000)+rand.Int63n(1000000000)+rand.Int63n(1000000000)+rand.Int63n(100000), 10)
	}
	password = password[0:6]
	return password, authorization.GeneratePasswordHash(password)
}

func IsTeamExist(id int64, w http.ResponseWriter) bool {
	if id != 0 {
		var count int
		err := src.CustomConnection.QueryRow("SELECT COUNT(id) FROM teams WHERE id=?", id).Scan(&count)
		if err != nil {
			return conf.ErrDatabaseQueryFailed.Print(w)
		}
		if count > 0 {
			return true
		} else {
			return conf.ErrTeamIncorrect.Print(w)
		}
	} else {
		return true
	}
}

func IsReasonExist(id int64, category_id int64, w http.ResponseWriter) bool {
	var count int
	err := src.CustomConnection.QueryRow("SELECT COUNT(id) FROM reasons WHERE id=? AND category_id=?", id, category_id).Scan(&count)
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(w)
	}
	if count > 0 {
		return true
	} else {
		return conf.ErrReasonIncorrect.Print(w)
	}
}

func GetUserIdByToken(token string) (int64, *conf.ApiResponse) {
	var id int64
	var login string
	err := src.Connection.QueryRow("SELECT login FROM sessions WHERE token=?", token).Scan(&login)
	if err != nil {
		return id, conf.ErrDatabaseQueryFailed
	}
	loginData := strings.Split(login, "_")
	if len(loginData) > 1 {
		id, err = strconv.ParseInt(loginData[1], 10, 64)
		if err != nil {
			return id, conf.ErrDatabaseQueryFailed
		}
	} else {
		err := src.Connection.QueryRow("SELECT id FROM users WHERE login=?", login).Scan(&id)
		if err != nil {
			return id, conf.ErrDatabaseQueryFailed
		}
	}
	return id, nil
}

func GetUserOrganizationByToken(token string) (string, *conf.ApiResponse) {
	var organizationName, login string
	err := src.Connection.QueryRow("SELECT login FROM sessions WHERE token=?", token).Scan(&login)
	if err != nil {
		return organizationName, conf.ErrDatabaseQueryFailed
	}
	err = src.Connection.QueryRow("SELECT organization FROM users WHERE login=?", login).Scan(&organizationName)
	if err != nil {
		return organizationName, conf.ErrDatabaseQueryFailed
	}
	return organizationName, nil
}

func GetUserOrganizationByID(id int64) (string, *conf.ApiResponse) {
	var organizationName string
	err := src.Connection.QueryRow("SELECT organization FROM users WHERE id=?", id).Scan(&organizationName)
	if err != nil {
		return organizationName, conf.ErrDatabaseQueryFailed
	}
	return organizationName, nil
}

func IsCategoryExist(id int64, w http.ResponseWriter) bool {
	var count int
	err := src.CustomConnection.QueryRow("SELECT COUNT(id) FROM categories WHERE id=?", id).Scan(&count)
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(w)
	}
	if count > 0 {
		return true
	} else {
		return conf.ErrCategoryIdIncorrect.Print(w)
	}
}
