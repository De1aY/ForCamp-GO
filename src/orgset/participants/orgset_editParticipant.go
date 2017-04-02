package participants

import (
	"net/http"
	"forcamp/src/orgset"
	"forcamp/conf"
	"forcamp/src"
	"log"
	"database/sql"
)

func EditParticipant(token string, participant Participant, ResponseWriter http.ResponseWriter) bool{
	Connection := src.Connect()
	defer Connection.Close()
	if orgset.CheckUserAccess(token, Connection, ResponseWriter) {
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token, Connection)
		if APIerr != nil {
			return conf.PrintError(APIerr, ResponseWriter)
		}
		NewConnection := src.Connect_Custom(Organization)
		defer NewConnection.Close()
		if checkEditParticipantData(participant, ResponseWriter, NewConnection) {
			ParticipantOrganization, APIerr := orgset.GetUserOrganizationByLogin(participant.Login, Connection)
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			if ParticipantOrganization != Organization {
				return conf.PrintError(conf.ErrUserNotFound, ResponseWriter)
			}
			APIerr = editParticipant_Request(participant, NewConnection)
			return conf.PrintSuccess(conf.RequestSuccess, ResponseWriter)
		}
	}
	return true
}

func editParticipant_Request(participant Participant, Connection *sql.DB) *conf.ApiError{
	Query, err := Connection.Prepare("UPDATE users SET name=?, surname=?, middlename=?, team=?, sex=? WHERE login=? AND access='0'")
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

func checkEditParticipantData(participant Participant, w http.ResponseWriter, Connection *sql.DB) bool {
	if len(participant.Login) > 0 {
		if len(participant.Name) > 0 {
			if len(participant.Surname) > 0 {
				if len(participant.Middlename) > 0 {
					if participant.Sex == 0 || participant.Sex == 1 {
						if orgset.CheckTeamID(participant.Team, w, Connection) {
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