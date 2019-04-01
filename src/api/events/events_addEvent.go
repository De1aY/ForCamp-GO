package events

import (
	"wplay/src"
	"wplay/conf"
)

func AddEvent(eventType int, participant_id int64, employee_id int64) (int64, *conf.ApiResponse) {
	if participant_id < 0 || employee_id < 0 {
		return 0, conf.ErrIdIncorrect
	}
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
