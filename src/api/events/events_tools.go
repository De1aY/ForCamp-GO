package events

import (
	"forcamp/src"
	"forcamp/conf"
)

type Event struct {
	ID             int64       `json:"id"`
	Type           int         `json:"type"`
	Participant_ID int64       `json:"participant_id"`
	Employee_ID    int64       `json:"employee_id"`
	Time           string      `json:"time"`
	EventData      interface{} `json:"event_data"`
}

type Event_Raw struct {
	ID             int64
	Type           int
	Participant_ID int64
	Employee_ID    int64
	Time           string
}

func isUserAdmin(user_id int64) (bool, *conf.ApiResponse) {
	var access int
	err := src.CustomConnection.QueryRow("SELECT access FROM users WHERE id=?", user_id).Scan(&access)
	if err != nil {
		return false, conf.ErrDatabaseQueryFailed
	}
	if access == 2 {
		return true, nil
	} else {
		return false, nil
	}
}

func isEventTypeExist(eventType int) bool {
	for _, b := range conf.EVENT_TYPES {
		if b == eventType {
			return true
		}
	}
	return false
}

func getRawEvent(event_id int64) (Event_Raw, *conf.ApiResponse){
	var rawEvent Event_Raw
	rawEvent.ID = event_id
	err := src.CustomConnection.QueryRow("SELECT type, participant_id, employee_id, time FROM events " +
		"WHERE id=?", event_id).Scan(&rawEvent.Type, &rawEvent.Participant_ID, &rawEvent.Employee_ID, &rawEvent.Time)
	if err != nil {
		return rawEvent, conf.ErrDatabaseQueryFailed
	}
	return rawEvent, nil
}
