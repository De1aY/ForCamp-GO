package participants

import (
	"net/http"
	"wplay/src/api/orgset"
	"wplay/conf"
	"wplay/src"
)

type resetParticipantPassword_Success struct {
	Password string `json:"password"`
}

func ResetParticipantPassword(token string, participant_id int64, responseWriter http.ResponseWriter) bool{
	if orgset.IsUserAdmin(token, responseWriter){
		organizationName, _, apiErr := orgset.GetUserOrganizationAndIdByToken(token); if apiErr != nil{
			return apiErr.Print(responseWriter)
		}
		participant_organization, participant_login, apiErr := orgset.GetUserOrganizationAndLoginByID(participant_id)
		if apiErr != nil{
			return apiErr.Print(responseWriter)
		}
		if participant_organization != organizationName {
			return conf.ErrUserNotFound.Print(responseWriter)
		}
		rawResp, apiErr := resetParticipantPassword(participant_id, participant_login); if apiErr != nil{
			return apiErr.Print(responseWriter)
		}
		resp := conf.ApiResponse{200, "success", rawResp}
		resp.Print(responseWriter)
	}
	return true
}

func resetParticipantPassword(participant_id int64, participant_login string) (resetParticipantPassword_Success, *conf.ApiResponse){
	password, hash := orgset.GeneratePassword()
	query, err := src.Connection.Prepare("UPDATE users SET password=? WHERE id=?"); if err != nil {
		return resetParticipantPassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(hash, participant_id); if err != nil {
		return resetParticipantPassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	query.Close()
	query, err = src.Connection.Prepare("DELETE FROM sessions WHERE login=?"); if err != nil {
		return resetParticipantPassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(participant_login); if err != nil {
		return resetParticipantPassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	return resetParticipantPassword_Success{password}, nil
}
