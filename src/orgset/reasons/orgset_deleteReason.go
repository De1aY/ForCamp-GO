package reasons

import (
	"forcamp/src/orgset"
	"net/http"
	"forcamp/src"
	"forcamp/conf"
	"database/sql"
	"log"
)

func DeleteReason(token string, id int64, ResponseWriter http.ResponseWriter) bool{
	connection := src.Connect()
	defer connection.Close()
	if orgset.CheckUserAccess(token, connection, ResponseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token, connection)
		if APIerr != nil {
			return conf.PrintError(APIerr, ResponseWriter)
		}
		NewConnection := src.Connect_Custom(Organization)
		defer NewConnection.Close()
		APIerr = deleteReason_Request(NewConnection, id)
		if APIerr != nil {
			return conf.PrintError(APIerr, ResponseWriter)
		}
		conf.PrintSuccess(conf.RequestSuccess, ResponseWriter)
	}
	return true
}

func deleteReason_Request(connection *sql.DB, id int64) *conf.ApiError{
	Query, err := connection.Prepare("DELETE FROM reasons WHERE id=?")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(id)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}