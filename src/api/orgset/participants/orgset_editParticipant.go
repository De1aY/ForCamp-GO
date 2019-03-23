package participants

import (
	"net/http"
	"wplay/src/api/orgset"
	"wplay/conf"
	"wplay/src"
)

func EditParticipant(token string, participant Participant, responseWriter http.ResponseWriter) bool{
	if orgset.IsUserAdmin(token, responseWriter) {
		organizationName, _, apiErr := orgset.GetUserOrganizationAndIdByToken(token); if apiErr != nil {
			return apiErr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(organizationName)
		if isEditParticipantDataValid(participant, responseWriter) {
			participant_organization, apiErr := orgset.GetUserOrganizationByID(participant.ID); if apiErr != nil {
				return apiErr.Print(responseWriter)
			}
			if participant_organization != organizationName {
				return conf.ErrUserNotFound.Print(responseWriter)
			}
			apiErr = editParticipant_Request(participant)
			return conf.RequestSuccess.Print(responseWriter)
		}
	}
	return true
}

func editParticipant_Request(participant Participant) *conf.ApiResponse{
	query, err := src.CustomConnection.Prepare("UPDATE users SET name=?, surname=?, middlename=?, " +
		"team=?, sex=? WHERE id=? AND access='0'")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(participant.Name, participant.Surname, participant.Middlename,
		participant.Team, participant.Sex, participant.ID)
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	query.Close()
	return nil
}

func isEditParticipantDataValid(participant Participant, w http.ResponseWriter) bool {
	if participant.ID > 0 {
		if len(participant.Name) > 0 {
			if len(participant.Surname) > 0 {
				if len(participant.Middlename) > 0 {
					if participant.Sex == 0 || participant.Sex == 1 {
						if orgset.IsTeamExist(participant.Team, w) {
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