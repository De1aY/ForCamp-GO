package participants

import (
	"net/http"
	"forcamp/src/api/orgset"
	"forcamp/conf"
	"forcamp/src"
	"log"
)

func EditParticipant(token string, participant Participant, responseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token, responseWriter) {
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil {
			return APIerr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		if checkEditParticipantData(participant, responseWriter) {
			ParticipantOrganization, APIerr := orgset.GetUserOrganizationByLogin(participant.Login)
			if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			if ParticipantOrganization != Organization {
				return conf.ErrUserNotFound.Print(responseWriter)
			}
			APIerr = editParticipant_Request(participant)
			return conf.RequestSuccess.Print(responseWriter)
		}
	}
	return true
}

func editParticipant_Request(participant Participant) *conf.ApiResponse{
	Query, err := src.CustomConnection.Prepare("UPDATE users SET name=?, surname=?, middlename=?, team=?, sex=? WHERE login=? AND access='0'")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(participant.Name, participant.Surname, participant.Middlename, participant.Team, participant.Sex, participant.Login)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	return nil
}

func checkEditParticipantData(participant Participant, w http.ResponseWriter) bool {
	if len(participant.Login) > 0 {
		if len(participant.Name) > 0 {
			if len(participant.Surname) > 0 {
				if len(participant.Middlename) > 0 {
					if participant.Sex == 0 || participant.Sex == 1 {
						if orgset.CheckTeamID(participant.Team, w) {
							return true
						} else {
							return false
						}
					} else {
						return conf.ErrSexIncorrect.Print(w)
					}
				} else {
					return conf.ErrMiddlenameEmpty.Print(w)
				}
			} else {
				return conf.ErrSurnameEmpty.Print(w)
			}
		} else {
			return conf.ErrNameEmpty.Print(w)
		}
	} else {
		return conf.ErrUserNotFound.Print(w)
	}
}