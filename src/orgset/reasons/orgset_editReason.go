package reasons

import (
	"forcamp/src/orgset"
	"net/http"
	"forcamp/src"
	"forcamp/conf"
	"database/sql"
	"log"
)

func EditReason(token string, reason Reason, ResponseWriter http.ResponseWriter) bool{
	connection := src.Connect()
	defer connection.Close()
	if orgset.CheckUserAccess(token, connection, ResponseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token, connection)
		if APIerr != nil {
			return conf.PrintError(APIerr, ResponseWriter)
		}
		NewConnection := src.Connect_Custom(Organization)
		defer NewConnection.Close()
		if orgset.CheckCategoryId(reason.Cat_id, ResponseWriter, NewConnection){
			APIerr = editReason_Request(NewConnection, reason)
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			conf.PrintSuccess(conf.RequestSuccess, ResponseWriter)
		}
	}
	return true
}

func editReason_Request(connection *sql.DB, reason Reason) *conf.ApiError{
	Query, err := connection.Prepare("UPDATE reasons SET text=?, modification=?, cat_id=? WHERE id=?")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(reason.Text, reason.Change, reason.Cat_id, reason.Id)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}