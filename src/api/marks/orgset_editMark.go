/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package marks

import (
	"wplay/src"
	"net/http"
	"wplay/src/api/authorization"
	"wplay/conf"
	"wplay/src/api/orgset"
	"strconv"
	"wplay/src/api/orgset/reasons"
)

func EditMark(token string, participant_id int64, category_id int64, reason_id int64, responseWriter http.ResponseWriter) bool {
	if authorization.IsTokenNotEmpty(token, responseWriter) {
		if authorization.IsTokenValid(token, responseWriter) {
			organizationName, employee_id, apiErr := orgset.GetUserOrganizationAndIdByToken(token)
			if apiErr != nil {
				return apiErr.Print(responseWriter)
			}
			src.CustomConnection = src.Connect_Custom(organizationName)
			if isEditMarkDataValid(participant_id, category_id, reason_id, employee_id, responseWriter) {
				apiErr = editMark(participant_id, employee_id, category_id, reason_id)
				if apiErr != nil {
					return apiErr.Print(responseWriter)
				}
				return conf.RequestSuccess.Print(responseWriter)
			}
		} else {
			return conf.ErrUserTokenIncorrect.Print(responseWriter)
		}
	}
	return true
}

func editMark(participant_id int64, employee_id int64, category_id int64, reason_id int64) *conf.ApiResponse{
	reasonData, apiErr := getReasonData(reason_id)
	if apiErr != nil {
		return apiErr
	}
	return editParticipantMark(participant_id, employee_id, category_id, reasonData)
}

func editParticipantMark(participant_id int64, employee_id int64,
	category_id int64, reasonData reasons.Reason) *conf.ApiResponse {
	currentMark, apiErr := getCurrentMarkValue(participant_id, category_id)
	negativeMarks, apiErr := isNegativeMarksAllowed(category_id); if apiErr != nil {
		return apiErr
	}
	var newMark int64
	if !negativeMarks && currentMark + reasonData.Change < 0 {
		newMark = 0
	} else {
		newMark = currentMark + reasonData.Change
	}
	if apiErr != nil {
		return apiErr
	}
	query, err := src.CustomConnection.Prepare("UPDATE participants SET `"+
		strconv.FormatInt(category_id, 10)+"`=? WHERE id=?")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	defer query.Close()
	_, err = query.Exec(newMark, participant_id); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return logMarkChange(participant_id, employee_id, category_id, reasonData.Text, currentMark, newMark)
}

func logMarkChange(participant_id int64, employee_id int64,
	category_id int64, text string, initial_value int64, final_value int64) *conf.ApiResponse {
	event_id, apiErr := addEvent(conf.EVENT_TYPE_MARK_CHANGE, participant_id, employee_id); if apiErr != nil {
		return apiErr
	}
	query, err := src.CustomConnection.Prepare("INSERT INTO marks_changes(id, category_id, employee_id, " +
		"participant_id, text, initial_value, final_value) VALUES (?,?,?,?,?,?,?)")
	defer query.Close()
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(event_id, category_id, employee_id, participant_id, text, initial_value, final_value)
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func getCurrentMarkValue(participant_id int64, category_id int64) (int64, *conf.ApiResponse) {
	var value int64
	err := src.CustomConnection.QueryRow("SELECT `"+strconv.FormatInt(category_id, 10)+"` FROM participants WHERE id=?", participant_id).Scan(&value)
	if err != nil {
		return 0, conf.ErrDatabaseQueryFailed
	}
	return value, nil
}

func getReasonData(reason_id int64) (reasons.Reason, *conf.ApiResponse) {
	var reasonData reasons.Reason
	reasonData.Id = reason_id
	err := src.CustomConnection.QueryRow("SELECT category_id, text, modification FROM reasons " +
		"WHERE id=?", reason_id).Scan(&reasonData.Cat_id, &reasonData.Text, &reasonData.Change)
	if err != nil {
		return reasonData, conf.ErrDatabaseQueryFailed
	}
	return reasonData, nil
}

func isEditMarkDataValid(participant_id int64, category_id int64, reason_id int64, employee_id int64, w http.ResponseWriter) bool {
	if !isUserEmployee(employee_id, w){
		return false
	}
	if !isEditingAllowed(employee_id, participant_id, category_id, w) {
		return false
	}
	if !orgset.IsCategoryExist(category_id, w) {
		return false
	}
	if !orgset.IsReasonExist(reason_id, category_id, w) {
		return false
	}
	if !isParticipantExist(participant_id, w) {
		return false
	}
	return true
}

func isEditingAllowed(employee_id int64, participant_id int64, category_id int64, responseWriter http.ResponseWriter) bool {
	var permission string
	err := src.CustomConnection.QueryRow("SELECT `" + strconv.FormatInt(category_id, 10) +
		"` FROM employees WHERE id=?", employee_id).Scan(&permission)
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(responseWriter)
	}
	if permission == "false" {
		return conf.ErrInsufficientRights.Print(responseWriter)
	}
	var participant_team, employee_team int64
	err = src.CustomConnection.QueryRow("SELECT team FROM users WHERE id=?", participant_id).Scan(&participant_team)
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(responseWriter)
	}
	err = src.CustomConnection.QueryRow("SELECT team FROM users WHERE id=?", employee_id).Scan(&employee_team)
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(responseWriter)
	}
	var isSelfMarksAllowed string
	err = src.CustomConnection.QueryRow("SELECT value FROM " +
		"settings WHERE name=?", "self_marks").Scan(&isSelfMarksAllowed)
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(responseWriter)
	}
	if isSelfMarksAllowed == "false" && participant_team == employee_team {
		return conf.ErrInsufficientRights.Print(responseWriter)
	}
	return true
}

func isNegativeMarksAllowed(category_id int64) (bool, *conf.ApiResponse) {
	var negMarks string
	err := src.CustomConnection.QueryRow("SELECT negative_marks FROM categories WHERE id=?", category_id).Scan(&negMarks)
	if err != nil {
		return false, conf.ErrDatabaseQueryFailed
	}
	if negMarks == "true" {
		return true, nil
	} else {
		return false, nil
	}
}

func isUserEmployee(employee_id int64, w http.ResponseWriter) bool {
	apiErr := isUserEmployee_Request(employee_id)
	if apiErr != nil {
		return apiErr.Print(w)
	}
	return true
}

func isUserEmployee_Request(employee_id int64) *conf.ApiResponse {
	var access int
	err := src.CustomConnection.QueryRow("SELECT access FROM users WHERE id=?", employee_id).Scan(&access)
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	if access != 0 {
		return nil
	} else {
		return conf.ErrInsufficientRights
	}
}

func isParticipantExist(participant_id int64, w http.ResponseWriter) bool {
	var access int
	err := src.CustomConnection.QueryRow("SELECT access FROM users WHERE id=?", participant_id).Scan(&access)
	if err != nil {
		return conf.ErrDatabaseQueryFailed.Print(w)
	}
	if access == 0 {
		return true
	} else {
		return conf.ErrIdIncorrect.Print(w)
	}
}

func addEvent(eventType int, participant_id int64, employee_id int64) (int64, *conf.ApiResponse) {
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

func isEventTypeExist(eventType int) bool {
	for _, b := range conf.EVENT_TYPES {
		if b == eventType {
			return true
		}
	}
	return false
}
