/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package marks

import (
	"wplay/conf"
	"wplay/src"
	"strconv"
)

func DeleteMarkChange(eventID int64, user_id int64, isAdmin bool) *conf.ApiResponse {
	if isAdmin {
		return deleteMarkChange(eventID)
	} else {
		access := isEmployee(user_id); if access != nil {
			return access
		}
		return deleteMarkChange(eventID)
	}
}

func isEmployee(user_id int64) *conf.ApiResponse {
	var access int
	err := src.CustomConnection.QueryRow("SELECT access FROM users WHERE id=?", user_id).Scan(&access)
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	if access == 0 {
		return conf.ErrInsufficientRights
	} else {
		return nil
	}
}

func deleteMarkChange(eventID int64) *conf.ApiResponse {
	apiErr := undoMarkChange(eventID); if apiErr != nil {
		return apiErr
	}
	query, err := src.CustomConnection.Prepare("DELETE FROM marks_changes WHERE id=?");
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(eventID);
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func undoMarkChange(eventID int64) *conf.ApiResponse {
	markChange, apiErr := getRawMarkChange(eventID); if apiErr != nil {
		return apiErr
	}
	currentMark, apiErr := getCurrentMarkValue(markChange.Participant_ID, markChange.Category_ID); if apiErr != nil {
		return apiErr
	}
	negativeMarks, apiErr := isNegativeMarksAllowed(markChange.Category_ID); if apiErr != nil {
		return apiErr
	}
	change := (markChange.Final_Value - markChange.Initial_Value) * -1
	var newMark int64
	if !negativeMarks && currentMark + change < 0 {
		newMark = 0
	} else {
		newMark = currentMark + change
	}
	query, err := src.CustomConnection.Prepare("UPDATE participants SET `"+
		strconv.FormatInt(markChange.Category_ID, 10)+"`=? WHERE id=?"); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(newMark, markChange.Participant_ID); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

