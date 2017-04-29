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

type resetParticipantPassword_Success struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Password string `json:"password"`
}

func (success *resetParticipantPassword_Success) toJSON() string {
	resp, _ := json.Marshal(success)
	return string(resp)
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
		resp, APIerr := resetParticipantPassword_Request(login)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		fmt.Fprintf(ResponseWriter, resp.toJSON())
	}
	return true
}

func resetParticipantPassword_Request(login string) (resetParticipantPassword_Success, *conf.ApiError){
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
	return resetParticipantPassword_Success{200, "success", Password}, nil
}