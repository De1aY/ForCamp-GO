package events

import (
	"wplay/src/api/authorization"
	"net/http"
	"wplay/src"
	"wplay/conf"
	"database/sql"
	"wplay/src/api/marks"
	"wplay/src/api/orgset"
	"wplay/src/api/emotional_marks"
)

type getEvents_Success struct {
	Events []Event `json:"events"`
}

func GetEvents(token string, user_id int64, rowsPerPage int64,
	page int64, ascending bool, event_type int, responseWriter http.ResponseWriter) bool {
	if authorization.IsTokenValid(token, responseWriter) {
		organizationName, apiErr := orgset.GetUserOrganizationByToken(token); if apiErr != nil {
			return apiErr.Print(responseWriter)
		}
		if user_id > 0 {
			user_organization, apiErr := orgset.GetUserOrganizationByID(user_id); if apiErr != nil {
				return apiErr.Print(responseWriter)
			}
			if organizationName != user_organization {
				return conf.ErrUserNotFound.Print(responseWriter)
			}
		}
		var limit, offset int64
		if rowsPerPage > 0 {
			limit = rowsPerPage
			if page < 0 {
				page = 0
			}
			offset = rowsPerPage * (page - 1)
		} else {
			limit = 1000
			offset = 0
		}
		if !isEventTypeExist(event_type) {
			return conf.ErrEventTypeIncorrect.Print(responseWriter)
		}
		events, apiErr := GetEvents_Request(user_id, limit, offset, ascending, event_type); if apiErr != nil {
			return apiErr.Print(responseWriter)
		}
		response := &conf.ApiResponse{200, "success", getEvents_Success{events}}
		return response.Print(responseWriter)
	}
	return true
}

func GetEvents_Request(user_id int64, limit int64, offset int64,
	ascending bool, event_type int) ([]Event, *conf.ApiResponse) {
	var rows *sql.Rows
	var err error
	if user_id > 0 {
		if event_type > 0 {
			if ascending {
				rows, err = src.CustomConnection.Query("SELECT id,type,participant_id,employee_id,time FROM events "+
					"WHERE (participant_id=? OR employee_id=?) AND type=? " +
					"ORDER BY time ASC LIMIT ? OFFSET ?", user_id, user_id, event_type, limit, offset)
			} else {
				rows, err = src.CustomConnection.Query("SELECT id,type,participant_id,employee_id,time FROM events "+
					"WHERE (participant_id=? OR employee_id=?) AND type=? " +
					"ORDER BY time DESC LIMIT ? OFFSET ?", user_id, user_id, event_type, limit, offset)
			}
		} else {
			if ascending {
				rows, err = src.CustomConnection.Query("SELECT id,type,participant_id,employee_id,time FROM events "+
					"WHERE participant_id=? OR employee_id=? ORDER BY time ASC LIMIT ? OFFSET ?", user_id, user_id, limit, offset)
			} else {
				rows, err = src.CustomConnection.Query("SELECT id,type,participant_id,employee_id,time FROM events "+
					"WHERE participant_id=? OR employee_id=? ORDER BY time DESC LIMIT ? OFFSET ?", user_id, user_id, limit, offset)
			}
		}
	} else {
		if event_type > 0 {
			if ascending {
				rows, err = src.CustomConnection.Query("SELECT id,type,participant_id,employee_id,time FROM events "+
					"WHERE type=? ORDER BY time ASC LIMIT ? OFFSET ?", event_type, limit, offset)
			} else {
				rows, err = src.CustomConnection.Query("SELECT id,type,participant_id,employee_id,time FROM events "+
					"WHERE type=? ORDER BY time DESC LIMIT ? OFFSET ?", event_type, limit, offset)
			}
		} else {
			if ascending {
				rows, err = src.CustomConnection.Query("SELECT id,type,participant_id,employee_id,time FROM events "+
					"ORDER BY time ASC LIMIT ? OFFSET ?", limit, offset)
			} else {
				rows, err = src.CustomConnection.Query("SELECT id,type,participant_id,employee_id,time FROM events "+
					"ORDER BY time DESC LIMIT ? OFFSET ?", limit, offset)
			}
		}
	}
	if err != nil {
		return nil, conf.ErrDatabaseQueryFailed
	}
	var rawEvent Event_Raw
	var events []Event
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&rawEvent.ID, &rawEvent.Type, &rawEvent.Participant_ID, &rawEvent.Employee_ID, &rawEvent.Time)
		if err != nil {
			return nil, conf.ErrDatabaseQueryFailed
		}
		eventData, apiErr := getEventData(rawEvent);if apiErr != nil {
			return nil, apiErr
		}
		events = append(events, Event{rawEvent.ID,
			rawEvent.Type,
			rawEvent.Participant_ID,
			rawEvent.Employee_ID,
			rawEvent.Time,
			eventData})
	}
	if events == nil {
		events = make([]Event, 0)
	}
	return events, nil
}

func getEventData(rawEvent Event_Raw) (interface{}, *conf.ApiResponse) {
	switch rawEvent.Type {
		case conf.EVENT_TYPE_MARK_CHANGE:
			return marks.GetMarkChange(rawEvent.ID)
		case conf.EVENT_TYPE_EMOTIONAL_MARK:
			return emotional_marks.GetEmotionalMark(rawEvent.ID, rawEvent.Time)
		default:
			return nil, conf.ErrEventTypeIncorrect
	}
	return nil, conf.ErrEventTypeIncorrect
}
