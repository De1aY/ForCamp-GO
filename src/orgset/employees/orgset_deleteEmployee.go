package employees

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src"
	"log"
	"forcamp/src/orgset"
	"database/sql"
)

func DeleteEmployee(token string, login string, ResponseWriter http.ResponseWriter) bool{
	Connection := src.Connect()
	defer Connection.Close()
	if orgset.CheckUserAccess(token, Connection, ResponseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token, Connection)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		NewConnection := src.Connect_Custom(Organization)
		defer NewConnection.Close()
		APIerr = deleteEmployee_Request(login, Connection, NewConnection)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		return conf.PrintSuccess(conf.RequestSuccess, ResponseWriter)
	}
	return true
}

func deleteEmployee_Request(login string, Connection *sql.DB, NewConnection *sql.DB) *conf.ApiError{
	APIerr := deleteEmployee_Organization(login, NewConnection)
	if APIerr != nil{
		return APIerr
	}
	APIerr = deleteEmployee_Main(login, Connection)
	if APIerr != nil{
		return APIerr
	}
	return nil
}

func deleteEmployee_Main(login string, Connection *sql.DB) *conf.ApiError{
	Query, err := Connection.Prepare("DELETE FROM users WHERE login=?")
	if err != nil{
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(login)
	if err != nil{
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	Query, err = Connection.Prepare("DELETE FROM sessions WHERE login=?")
	if err != nil{
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(login)
	if err != nil{
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func deleteEmployee_Organization(login string, Connection *sql.DB) *conf.ApiError{
	Query, err := Connection.Prepare("DELETE FROM users WHERE login=? AND access='1'")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	resp, err := Query.Exec(login)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	rowsAffected, err := resp.RowsAffected()
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	if rowsAffected == 0{
		return conf.ErrUserNotFound
	}
	Query, err = Connection.Prepare("DELETE FROM employees WHERE login=?")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(login)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}
