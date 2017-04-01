package participants

import (
	"net/http"
	"forcamp/src/orgset"
	"forcamp/conf"
	"forcamp/src"
	"log"
)

func EditParticipant(token string, participant Participant, ResponseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token, ResponseWriter) && checkEditParticipantData(participant, ResponseWriter) {
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil {
			return conf.PrintError(APIerr, ResponseWriter)
		}
		src.NewConnection = src.Connect_Custom(Organization)
		ParticipantOrganization, APIerr := orgset.GetUserOrganizationByLogin(participant.Login)
		if APIerr != nil {
			return conf.PrintError(APIerr, ResponseWriter)
		}
		if ParticipantOrganization != Organization{
			return conf.PrintError(conf.ErrUserNotFound, ResponseWriter)
		}
		APIerr = editParticipant_Request(participant)
		return conf.PrintSuccess(conf.RequestSuccess, ResponseWriter)
	}
	return true
}

func editParticipant_Request(participant Participant) *conf.ApiError{
	Query, err := src.NewConnection.Prepare("UPDATE users SET name=?, surname=?, middlename=?, team=?, sex=? WHERE login=? AND access='0'")
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
						return conf.PrintError(conf.ErrParticipantSexIncorrect, w)
					}
				} else {
					return conf.PrintError(conf.ErrParticipantMiddlenameEmpty, w)
				}
			} else {
				return conf.PrintError(conf.ErrParticipantSurnameEmpty, w)
			}
		} else {
			return conf.PrintError(conf.ErrParticipantNameEmpty, w)
		}
	} else {
		return conf.PrintError(conf.ErrUserNotFound, w)
	}
}