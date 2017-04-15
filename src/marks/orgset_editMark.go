/*
	Copyright: "Null team", 2016 - 2017
	Author: "De1aY"
	Documentation: https://bitbucket.org/lyceumdevelopers/golang/wiki/Home
*/
package marks

import (
	"forcamp/src"
	"net/http"
	"forcamp/src/authorization"
	"forcamp/conf"
	"forcamp/src/orgset"
	"log"
	"strconv"
)

func EditMark(token string, participant_login string, category_id int64, reason_id int64, responseWriter http.ResponseWriter) bool {
	if authorization.CheckTokenForEmpty(token, responseWriter) {
		if authorization.CheckToken(token, responseWriter) {
			organization, employee_login, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
			if APIerr != nil {
				return conf.PrintError(APIerr, responseWriter)
			}
			src.CustomConnection = src.Connect_Custom(organization)
			if checkData(token, participant_login, category_id, reason_id, responseWriter) {
				APIerr = editMark_Request(participant_login, employee_login, category_id, reason_id)
				if APIerr != nil {
					return conf.PrintError(APIerr, responseWriter)
				}
				return conf.PrintSuccess(conf.RequestSuccess, responseWriter)
			}
		} else {
			return conf.PrintError(conf.ErrUserTokenIncorrect, responseWriter)
		}
	}
	return true
}

func editMark_Request(participant_login string, employee_login string, category_id int64, reason_id int64) *conf.ApiError{
	change, APIerr := getReasonChange(reason_id)
	if APIerr != nil {
		return APIerr
	}
	return editParticipantMark(participant_login, employee_login, category_id, reason_id, change)
}

func editParticipantMark(participant_login string, employee_login string, category_id int64, reason_id int64, change int) *conf.ApiError {
	currentMark, APIerr := getCurrentMarkValue(participant_login, category_id)
	if APIerr != nil {
		return APIerr
	}
	Query, err := src.CustomConnection.Prepare("UPDATE participants SET `"+strconv.FormatInt(category_id, 10)+"`=? WHERE login=?")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(currentMark+change, participant_login)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return logMarkChange(participant_login, employee_login, reason_id)
}

func logMarkChange(participant_login string, employee_login string, reason_id int64) *conf.ApiError {
	Query, err := src.CustomConnection.Prepare("INSERT INTO marks_changes(reason_id, employee_login, participant_login) VALUES (?,?,?)")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(reason_id, employee_login, participant_login)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func getCurrentMarkValue(participant_login string, category_id int64) (int, *conf.ApiError) {
	var value int
	err := src.CustomConnection.QueryRow("SELECT `"+strconv.FormatInt(category_id, 10)+"` FROM participants WHERE login=?", participant_login).Scan(&value)
	if err != nil {
		log.Print(err)
		return 0, conf.ErrDatabaseQueryFailed
	}
	return value, nil
}

func getReasonChange(reason_id int64) (int, *conf.ApiError) {
	var change int
	err := src.CustomConnection.QueryRow("SELECT modification FROM reasons WHERE id=?", reason_id).Scan(&change)
	if err != nil {
		log.Print(err)
		return 0, conf.ErrDatabaseQueryFailed
	}
	return change, nil
}

func checkData(token string, participant_login string, category_id int64, reason_id int64, w http.ResponseWriter) bool {
	if checkUserAccess(token, w){
		if orgset.CheckCategoryId(category_id, w){
			if orgset.CheckReasonID(reason_id, category_id, w){
				if checkParticipantLogin(participant_login, w) {
					return true
				}
			}
		}
	}
	return false
}

func checkUserAccess(token string, w http.ResponseWriter) bool {
	organization, employee_login, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
	if APIerr != nil {
		return conf.PrintError(APIerr, w)
	}
	src.CustomConnection = src.Connect_Custom(organization)
	APIerr = checkUserAccess_Request(employee_login)
	if APIerr != nil {
		return conf.PrintError(APIerr, w)
	}
	return true
}

func checkUserAccess_Request(employee_login string) *conf.ApiError {
	var access int
	err := src.CustomConnection.QueryRow("SELECT access FROM users WHERE login=? LIMIT 1", employee_login).Scan(&access)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	if access != 0 {
		return nil
	} else {
		return conf.ErrInsufficientRights
	}
}

func checkParticipantLogin(participant_login string, w http.ResponseWriter) bool {
	var access int
	err := src.CustomConnection.QueryRow("SELECT access FROM users WHERE login=?", participant_login).Scan(&access)
	if err != nil {
		log.Print(err)
		return conf.PrintError(conf.ErrDatabaseQueryFailed, w)
	}
	if access == 0 {
		return true
	} else {
		return conf.PrintError(conf.ErrParticipantLoginIncorrect, w)
	}
}