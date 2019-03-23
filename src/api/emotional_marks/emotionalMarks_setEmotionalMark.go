package emotional_marks

import (
	"net/http"
	"wplay/src/api/authorization"
	"wplay/src"
	"wplay/conf"
	"wplay/src/api/orgset"
)

func SetEmotionalMark(token string, emotional_mark int64, responseWriter http.ResponseWriter) bool {
	if authorization.IsTokenValid(token, responseWriter) {
		organizationName, user_id, apiErr := orgset.GetUserOrganizationAndIdByToken(token); if apiErr != nil {
			return apiErr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(organizationName)
		if isParticipant(user_id, responseWriter) {
			if emotional_mark > 0 && emotional_mark < 6 {
				response := setEmotionalMark(user_id, emotional_mark); if response != nil {
					return response.Print(responseWriter)
				} else {
					return conf.RequestSuccess.Print(responseWriter)
				}
			} else {
				return conf.ErrEmotionalMarkValueIncorrect.Print(responseWriter)
			}
		}
	}
	return true
}

func isParticipant(user_id int64, responseWriter http.ResponseWriter) bool {
	var access int
	err := src.CustomConnection.QueryRow("SELECT access FROM users WHERE id=?", user_id).Scan(&access)
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(responseWriter)
	}
	if access == 0 {
		return true
	} else {
		return conf.ErrInsufficientRights.Print(responseWriter)
	}
}

func setEmotionalMark(participant_id int64, emotional_mark int64) *conf.ApiResponse {
	event_id, apiErr := addEvent(conf.EVENT_TYPE_EMOTIONAL_MARK, participant_id, -1); if apiErr != nil {
		return apiErr
	}
	query, err := src.CustomConnection.Prepare("INSERT INTO emotional_marks(id,participant_id,mark) VALUES(?,?,?)")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(event_id, participant_id, emotional_mark)
	return nil
}

func addEvent(eventType int, participant_id int64, employee_id int64) (int64, *conf.ApiResponse) {
	if !isEventTypeExist(eventType) {
		return 0, conf.ErrEventTypeIncorrect
	}
	req, err := src.CustomConnection.Prepare("INSERT INTO events(type, participant_id, employee_id) VALUES(?,?,?)"); if err != nil {
		return 0, conf.ErrDatabaseQueryFailed
	}
	result, err := req.Exec(eventType, participant_id, employee_id); if err != nil {
		return 0, conf.ErrDatabaseQueryFailed
	}
	eventID, err := result.LastInsertId(); if err != nil {
		return 0, conf.ErrDatabaseQueryFailed
	}
	return eventID, nil
}

func isEventTypeExist(eventType int) bool {
	for _, b := range conf.EVENT_TYPES {
		if b == eventType {
			return true
		}
	}
	return false
}