package participants

import (
	"net/http"
	"forcamp/src/orgset"
	"forcamp/conf"
	"forcamp/src"
	"log"
	"encoding/json"
	"fmt"
)

type ResetParticipantPassword_Success struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Password string `json:"password"`
}

func ResetParticipantPassword(token string, login string, ResponseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token, ResponseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		ParticipantOrganization, APIerr := orgset.GetUserOrganizationByLogin(login)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		if ParticipantOrganization != Organization{
			return conf.PrintError(conf.ErrUserNotFound, ResponseWriter)
		}
		Resp, APIerr := resetParticipantPassword_Request(login)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		Response, _ := json.Marshal(Resp)
		fmt.Fprintf(ResponseWriter, string(Response))
	}
	return true
}

func resetParticipantPassword_Request(login string) (ResetParticipantPassword_Success, *conf.ApiError){
	Password, Hash := orgset.GeneratePassword()
	Query, err := src.Connection.Prepare("UPDATE users SET password=? WHERE login=?")
	if err != nil {
		log.Print(err)
		return ResetParticipantPassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(Hash, login)
	if err != nil {
		log.Print(err)
		return ResetParticipantPassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	Query, err = src.Connection.Prepare("DELETE FROM sessions WHERE login=?")
	if err != nil {
		log.Print(err)
		return ResetParticipantPassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(login)
	if err != nil {
		log.Print(err)
		return ResetParticipantPassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	return ResetParticipantPassword_Success{200, "success", Password}, nil
}