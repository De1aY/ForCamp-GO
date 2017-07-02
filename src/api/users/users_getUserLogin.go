package users

import (
	"net/http"
	"forcamp/conf"
	"database/sql"
	"forcamp/src/api/authorization"
	"forcamp/src"
)

func GetUserLogin(Token string, responseWriter http.ResponseWriter) bool{
	if authorization.CheckTokenForEmpty(Token, responseWriter) {
		if authorization.CheckToken(Token, responseWriter) {
			Query, err := src.Connection.Query("SELECT login FROM sessions WHERE token=?", Token)
			if err != nil {
				return conf.ErrDatabaseQueryFailed.Print(responseWriter)
			}
			Login, APIerr := getUserLoginFromQuery(Query)
			if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			rawResp := getUserLogin_Success{Login}
			resp := &conf.ApiResponse{200, "success", rawResp}
			resp.Print(responseWriter)
			return true
		} else {
			return conf.ErrUserTokenIncorrect.Print(responseWriter)
		}
	}
	return false
}

func getUserLoginFromQuery(rows *sql.Rows) (string, *conf.ApiResponse){
	var login string
	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&login)
		if err != nil {
			return "", conf.ErrDatabaseQueryFailed
		}
	}
	return login, nil
}