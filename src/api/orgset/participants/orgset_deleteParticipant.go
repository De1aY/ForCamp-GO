package participants

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/api/orgset"
)

func DeleteParticipant(token string, login string, responseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token, responseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil{
			return APIerr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		APIerr = deleteParticipant_Request(login)
		if APIerr != nil{
			return APIerr.Print(responseWriter)
		}
		return conf.RequestSuccess.Print(responseWriter)
	}
	return true
}

func deleteParticipant_Request(login string) *conf.ApiResponse{
	APIerr := deleteParticipant_Organization(login)
	if APIerr != nil{
		return APIerr
	}
	APIerr = deleteParticipant_Main(login)
	if APIerr != nil{
		return APIerr
	}
	return nil
}

func deleteParticipant_Main(login string) *conf.ApiResponse{
	Query, err := src.Connection.Prepare("DELETE FROM users WHERE login=?")
	if err != nil{
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(login)
	if err != nil{
		return conf.ErrDatabaseQueryFailed
	}
	Query, err = src.Connection.Prepare("DELETE FROM sessions WHERE login=?")
	if err != nil{
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(login)
	if err != nil{
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func deleteParticipant_Organization(login string) *conf.ApiResponse{
	Query, err := src.CustomConnection.Prepare("DELETE FROM users WHERE login=? AND access='0'")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	resp, err := Query.Exec(login)
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	rowsAffected, err := resp.RowsAffected()
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	if rowsAffected == 0{
		return conf.ErrUserNotFound
	}
	Query, err = src.CustomConnection.Prepare("DELETE FROM participants WHERE login=?")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(login)
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}