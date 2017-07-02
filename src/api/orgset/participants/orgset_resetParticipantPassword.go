package participants

import (
	"net/http"
	"forcamp/src/api/orgset"
	"forcamp/conf"
	"forcamp/src"
	"log"
)

type resetParticipantPassword_Success struct {
	Password string `json:"password"`
}

func ResetParticipantPassword(token string, login string, responseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token, responseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil{
			return APIerr.Print(responseWriter)
		}
		ParticipantOrganization, APIerr := orgset.GetUserOrganizationByLogin(login)
		if APIerr != nil{
			return APIerr.Print(responseWriter)
		}
		if ParticipantOrganization != Organization{
			return conf.ErrUserNotFound.Print(responseWriter)
		}
		rawResp, APIerr := resetParticipantPassword_Request(login)
		if APIerr != nil{
			return APIerr.Print(responseWriter)
		}
		resp := conf.ApiResponse{200, "success", rawResp}
		resp.Print(responseWriter)
	}
	return true
}

func resetParticipantPassword_Request(login string) (resetParticipantPassword_Success, *conf.ApiResponse){
	Password, Hash := orgset.GeneratePassword()
	Query, err := src.Connection.Prepare("UPDATE users SET password=? WHERE login=?")
	if err != nil {
		log.Print(err)
		return resetParticipantPassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(Hash, login)
	if err != nil {
		log.Print(err)
		return resetParticipantPassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	Query, err = src.Connection.Prepare("DELETE FROM sessions WHERE login=?")
	if err != nil {
		log.Print(err)
		return resetParticipantPassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(login)
	if err != nil {
		log.Print(err)
		return resetParticipantPassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	return resetParticipantPassword_Success{Password}, nil
}