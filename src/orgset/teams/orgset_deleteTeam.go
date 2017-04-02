package teams

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src"
	"log"
	"forcamp/src/orgset"
	"database/sql"
)

func DeleteTeam(token string, id int64, ResponseWriter http.ResponseWriter) bool{
	Connection := src.Connect()
	defer Connection.Close()
	if orgset.CheckUserAccess(token, Connection, ResponseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token, Connection)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		NewConnection := src.Connect_Custom(Organization)
		defer NewConnection.Close()
		APIerr = deleteTeam_Request(id, NewConnection)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		return conf.PrintSuccess(conf.RequestSuccess, ResponseWriter)
	}
	return true
}

func deleteTeam_Request(id int64, Connection *sql.DB) *conf.ApiError{
	Query, err := Connection.Prepare("DELETE FROM teams WHERE id=?")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(id)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	APIerr := deleteTeam_Users(id, Connection)
	if APIerr != nil {
		return APIerr
	}
	return nil
}

func deleteTeam_Users(id int64, Connection *sql.DB) *conf.ApiError{
	Query, err := Connection.Prepare("UPDATE users SET team='0' WHERE team=?")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(id)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	return nil
}
