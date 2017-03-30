package orgset

import (
	"forcamp/src"
	"database/sql"
	"net/http"
	"forcamp/src/authorization"
	"forcamp/conf"
	"log"
)

type OrgSettings struct {
	Participant string `json:"participant"`
	Team string `json:"team"`
	Organization string `json:"organization"`
	Period string `json:"period"`
	SelfMarks string `json:"self_marks"`
}

type GetOrgSettings_Success struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Settings OrgSettings `json:"settings"`
}

var Connection = src.Connect()
var NewConnection *sql.DB

func checkUserAccess(token string, ResponseWriter http.ResponseWriter) bool{
	if authorization.CheckTokenForEmpty(token, ResponseWriter) {
		if (authorization.CheckToken(token, ResponseWriter)) {
			Organization, Login, APIerr := getUserOrganizationAndLoginByToken(token)
			if APIerr != nil{
				return conf.PrintError(APIerr, ResponseWriter)
			}
			NewConnection = src.Connect_Custom(Organization)
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

func getUserOrganizationAndLoginByToken(Token string) (string, string, *conf.ApiError){
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