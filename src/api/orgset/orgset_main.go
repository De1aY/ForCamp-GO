/*
	Copyright: "Null team", 2016 - 2017
	Author: "De1aY"
	Documentation: https://bitbucket.org/lyceumdevelopers/golang/wiki/Home
*/
package orgset

import (
	"forcamp/src"
	"database/sql"
	"net/http"
	"forcamp/src/api/authorization"
	"forcamp/conf"
	"log"
	"strconv"
	"math/rand"
	"time"
)

func CheckUserAccess(token string, responseWriter http.ResponseWriter) bool{
	if authorization.CheckTokenForEmpty(token, responseWriter) {
		if (authorization.CheckToken(token, responseWriter)) {
			Organization, Login, APIerr := GetUserOrganizationAndLoginByToken(token)
			if APIerr != nil{
				return APIerr.Print(responseWriter)
			}
			CustomConnection := src.Connect_Custom(Organization)
			defer CustomConnection.Close()
			Query, err := CustomConnection.Query("SELECT access FROM users WHERE login=?", Login)
			if err != nil {
				log.Print(err)
				return conf.ErrDatabaseQueryFailed.Print(responseWriter)
			}
			return checkAccessFromQuery(Query, responseWriter)
		} else {
			return conf.ErrUserTokenIncorrect.Print(responseWriter)
		}
	}
	return false
}

func checkAccessFromQuery(rows *sql.Rows, w http.ResponseWriter) bool{
	defer rows.Close()
	for rows.Next(){
		var access int
		err := rows.Scan(&access)
		if err != nil{
			return conf.ErrDatabaseQueryFailed.Print(w)
		}
		if access == 2{
			return true
		} else {
			return conf.ErrInsufficientRights.Print(w)
		}
	}
	return true
}

func GetUserOrganizationAndLoginByToken(Token string) (string, string, *conf.ApiResponse){
	Query, err := src.Connection.Query("SELECT login FROM sessions WHERE token=?", Token)
	if err!= nil{
		log.Print(err)
		return "", "", conf.ErrDatabaseQueryFailed
	}
	Login, APIerr := getUserLoginFromQuery(Query)
	if APIerr != nil {
		return "", "", APIerr
	}
	Query, err = src.Connection.Query("SELECT organization FROM users WHERE login=?", Login)
	if err != nil {
		log.Print(err)
		return "", "", conf.ErrDatabaseQueryFailed
	}
	Organization, APIerr := getUserOrganizationFromQuery(Query)
	if APIerr != nil {
		return "", "", APIerr
	}
	return Organization, Login, nil
}

func getUserOrganizationFromQuery(rows *sql.Rows) (string, *conf.ApiResponse){
	defer rows.Close()
	var organization string
	for rows.Next(){
		err := rows.Scan(&organization)
		if err != nil {
			log.Print(err)
			return "", conf.ErrDatabaseQueryFailed
		}
	}
	return organization, nil
}

func getUserLoginFromQuery(rows *sql.Rows) (string, *conf.ApiResponse){
	var login string
	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&login)
		if err != nil {
			log.Print(err)
			return "", conf.ErrDatabaseQueryFailed
		}
	}
	return login, nil
}

func GeneratePassword() (string, string){
	password := ""
	for len(password) < 7{
		rand.Seed(time.Now().UnixNano())
		password = strconv.FormatInt(rand.Int63n(1000000000)+rand.Int63n(1000000000)+rand.Int63n(1000000000)+rand.Int63n(100000), 10)
	}
	password = password[0:6]
	return password, authorization.GeneratePasswordHash(password)
}

func CheckTeamID(id int64, w http.ResponseWriter) bool{
	if id != 0 {
		var count int
		err := src.CustomConnection.QueryRow("SELECT COUNT(id) FROM teams WHERE id=?", id).Scan(&count)
		if err != nil {
			log.Print(err)
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

func CheckReasonID(id int64, category_id int64, w http.ResponseWriter) bool {
	var count int
	err := src.CustomConnection.QueryRow("SELECT COUNT(id) FROM reasons WHERE id=? AND cat_id=?", id, category_id).Scan(&count)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed.Print(w)
	}
	if count > 0 {
		return true
	} else {
		return conf.ErrReasonIncorrect.Print(w)
	}
}

func GetUserOrganizationByLogin(login string) (string, *conf.ApiResponse){
	Query, err := src.Connection.Query("SELECT organization FROM users WHERE login=?", login)
	if err != nil {
		log.Print(err)
		return "", conf.ErrDatabaseQueryFailed
	}
	defer Query.Close()
	var organization string
	for Query.Next(){
		err := Query.Scan(&organization)
		if err != nil {
			log.Print(err)
			return "", conf.ErrDatabaseQueryFailed
		}
	}
	return organization, nil
}

func CheckCategoryId(id int64, w http.ResponseWriter) bool{
	var count int
	err := src.CustomConnection.QueryRow("SELECT COUNT(id) FROM categories WHERE id=?", id).Scan(&count)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed.Print(w)
	}
	if count > 0 {
		return true
	} else {
		return conf.ErrCategoryIdIncorrect.Print(w)
	}
}