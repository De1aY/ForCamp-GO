/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package events

import (
	"net/http"
	"wplay/src/api/authorization"
	"wplay/src/api/orgset"
	"wplay/src"
	"wplay/conf"
	"wplay/src/api/marks"
)

func DeleteEvent(token string, event_id int64, responseWriter http.ResponseWriter) bool {
	if authorization.IsTokenNotEmpty(token, responseWriter) {
		if authorization.IsTokenValid(token, responseWriter) {
			organizationName, user_id, apiErr := orgset.GetUserOrganizationAndIdByToken(token); if apiErr != nil {
				return apiErr.Print(responseWriter)
			}
			src.CustomConnection = src.Connect_Custom(organizationName)
			if event_id < 0 {
				return conf.ErrIdIncorrect.Print(responseWriter)
			}
			rawEvent, apiErr := getRawEvent(event_id); if apiErr != nil {
				return apiErr.Print(responseWriter)
			}
			isAdmin, apiErr := isUserAdmin(user_id); if apiErr != nil {
				return apiErr.Print(responseWriter)
			}
			if rawEvent.Employee_ID == user_id || rawEvent.Participant_ID == user_id || isAdmin {
				apiErr = deleteEvent(rawEvent, user_id, isAdmin)
				if apiErr != nil {
					return apiErr.Print(responseWriter)
				} else {
					return conf.RequestSuccess.Print(responseWriter)
				}
			} else {
				return conf.ErrInsufficientRights.Print(responseWriter)
			}
		}
	}
	return false
}

func deleteEvent(rawEvent Event_Raw, user_id int64, isAdmin bool) *conf.ApiResponse {
	query, err := src.CustomConnection.Prepare("DELETE FROM events WHERE id=?"); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(rawEvent.ID); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	switch rawEvent.Type {
		case conf.EVENT_TYPE_MARK_CHANGE: {
			return marks.DeleteMarkChange(rawEvent.ID, user_id, isAdmin)
		}
		default:
			return conf.ErrEventTypeIncorrect
	}
	return conf.ErrEventTypeIncorrect
}
