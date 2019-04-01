package participants

import (
	"net/http"
	"wplay/conf"
	"wplay/src"
	"wplay/src/api/orgset"
)

func DeleteParticipant(token string, participant_id int64, responseWriter http.ResponseWriter) bool{
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
		src.CustomConnection = src.Connect_Custom(organizationName)
		apiErr = deleteParticipant(participant_id, participant_login)
		if apiErr != nil{
			return apiErr.Print(responseWriter)
		}
		return conf.RequestSuccess.Print(responseWriter)
	}
	return true
}

func deleteParticipant(participant_id int64, participant_login string) *conf.ApiResponse{
	apiErr := deleteParticipant_Organization(participant_id); if apiErr != nil{
		return apiErr
	}
	apiErr = deleteParticipant_Main(participant_id, participant_login); if apiErr != nil{
		return apiErr
	}
	return nil
}

func deleteParticipant_Main(participant_id int64, participant_login string) *conf.ApiResponse{
	query, err := src.Connection.Prepare("DELETE FROM users WHERE id=?"); if err != nil{
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(participant_id); if err != nil{
		return conf.ErrDatabaseQueryFailed
	}
	query, err = src.Connection.Prepare("DELETE FROM sessions WHERE login=?"); if err != nil{
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(participant_login); if err != nil{
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func deleteParticipant_Organization(participant_id int64) *conf.ApiResponse{
	query, err := src.CustomConnection.Prepare("DELETE FROM users WHERE id=? AND access='0'"); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	resp, err := query.Exec(participant_id); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	query.Close()
	rowsAffected, err := resp.RowsAffected(); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	if rowsAffected == 0{
		return conf.ErrUserNotFound
	}
	query, err = src.CustomConnection.Prepare("DELETE FROM participants WHERE id=?"); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(participant_id); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}
