package orgset

import (
	"forcamp/src"
	"database/sql"
	"net/http"
	"forcamp/src/authorization"
	"forcamp/conf"
	"log"
	"strconv"
	"math/rand"
	"time"
)

func CheckUserAccess(token string, Connection *sql.DB, ResponseWriter http.ResponseWriter) bool{
	if authorization.CheckTokenForEmpty(token, ResponseWriter) {
		if (authorization.CheckToken(token, Connection, ResponseWriter)) {
			Organization, Login, APIerr := GetUserOrganizationAndLoginByToken(token, Connection)
			if APIerr != nil{
				return conf.PrintError(APIerr, ResponseWriter)
			}
			NewConnection := src.Connect_Custom(Organization)
			defer NewConnection.Close()
			Query, err := NewConnection.Query("SELECT access FROM users WHERE login=?", Login)
			if err != nil {
				log.Print(err)
				return conf.PrintError(conf.ErrDatabaseQueryFailed, ResponseWriter)
			}
			return checkAccessFromQuery(Query, ResponseWriter)
		} else {
			return conf.PrintError(conf.ErrUserTokenIncorrect, ResponseWriter)
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
			return conf.PrintError(conf.ErrDatabaseQueryFailed, w)
		}
		if access == 2{
			return true
		} else {
			return conf.PrintError(conf.ErrInsufficientRights, w)
		}
	}
	return true
}

func GetUserOrganizationAndLoginByToken(Token string, Connection *sql.DB) (string, string, *conf.ApiError){
	Query, err := Connection.Query("SELECT login FROM sessions WHERE token=?", Token)
	if err!= nil{
		log.Print(err)
		return "", "", conf.ErrDatabaseQueryFailed
	}
	Login, APIerr := getUserLoginFromQuery(Query)
	if APIerr != nil {
		return "", "", APIerr
	}
	Query, err = Connection.Query("SELECT organization FROM users WHERE login=?", Login)
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

func getUserOrganizationFromQuery(rows *sql.Rows) (string, *conf.ApiError){
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

func getUserLoginFromQuery(rows *sql.Rows) (string, *conf.ApiError){
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

func getTeamsIDs(Connection *sql.DB) (map[int64]bool, *conf.ApiError){
	Query, err := Connection.Query("SELECT id FROM teams")
	if err != nil {
		log.Print(err)
		return make(map[int64]bool), conf.ErrDatabaseQueryFailed
	}
	defer Query.Close()
	IDs := make(map[int64]bool)
	var id int64
	for Query.Next(){
		err = Query.Scan(&id)
		if err != nil {
			log.Print(err)
			return make(map[int64]bool), conf.ErrDatabaseQueryFailed
		}
		IDs[id] = true
	}
	return IDs, nil
}

func CheckTeamID(id int64, w http.ResponseWriter, Connection *sql.DB) bool{
	TeamsIDs, APIerr := getTeamsIDs(Connection)
	if id != 0 {
		if APIerr != nil {
			return conf.PrintError(APIerr, w)
		}
		if TeamsIDs[id] {
			return true
		} else {
			return conf.PrintError(conf.ErrParticipantTeamIncorrect, w)
		}
	} else {
		return true
	}
}

func GetUserOrganizationByLogin(login string, Connection *sql.DB) (string, *conf.ApiError){
	Query, err := Connection.Query("SELECT organization FROM users WHERE login=?", login)
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

func CheckCategoryId(id int64, w http.ResponseWriter, connection *sql.DB) bool{
	Query, err := connection.Query("SELECT id FROM categories")
	if err != nil{
		log.Print(err)
		return conf.PrintError(conf.ErrDatabaseQueryFailed, w)
	}
	defer Query.Close()
	CatIDs := make(map[int64]bool)
	var catId int64
	for Query.Next(){
		err = Query.Scan(&catId)
		if err != nil {
			log.Print(err)
			return conf.PrintError(conf.ErrDatabaseQueryFailed, w)
		}
		CatIDs[catId] = true
	}
	if CatIDs[id]{
		return true
	} else {
		return conf.PrintError(conf.ErrCategoryIdIncorrect, w)
	}
}