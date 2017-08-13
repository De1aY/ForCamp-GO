package marks

import (
	"forcamp/src"
	"net/http"
	"forcamp/src/api/authorization"
	"forcamp/conf"
	"forcamp/src/api/orgset"
	"strconv"
)

func EditMark(token string, participant_login string, category_id int64, reason_id int64, responseWriter http.ResponseWriter) bool {
	if authorization.CheckTokenForEmpty(token, responseWriter) {
		if authorization.CheckToken(token, responseWriter) {
			organization, employee_login, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
			if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			src.CustomConnection = src.Connect_Custom(organization)

			if checkData(token, participant_login, category_id, reason_id, responseWriter) {
				APIerr = editMark_Request(participant_login, employee_login, category_id, reason_id)
				if APIerr != nil {
					return APIerr.Print(responseWriter)
				}
				return conf.RequestSuccess.Print(responseWriter)
			}
		} else {
			return conf.ErrUserTokenIncorrect.Print(responseWriter)
		}
	}
	return true
}

func editMark_Request(participant_login string, employee_login string, category_id int64, reason_id int64) *conf.ApiResponse{
	change, APIerr := getReasonChange(reason_id)
	if APIerr != nil {
		return APIerr
	}
	return editParticipantMark(participant_login, employee_login, category_id, reason_id, change)
}

func editParticipantMark(participant_login string, employee_login string, category_id int64, reason_id int64, change int64) *conf.ApiResponse {
	currentMark, APIerr := getCurrentMarkValue(participant_login, category_id)
	if APIerr != nil {
		return APIerr
	}
	Query, err := src.CustomConnection.Prepare("UPDATE participants SET `"+strconv.FormatInt(category_id, 10)+"`=? WHERE login=?")
	defer Query.Close()
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(currentMark + change, participant_login)
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return logMarkChange(participant_login, employee_login, reason_id)
}

func logMarkChange(participant_login string, employee_login string, reason_id int64) *conf.ApiResponse {
	Query, err := src.CustomConnection.Prepare("INSERT INTO marks_changes(reason_id, employee_login, participant_login) VALUES (?,?,?)")
	defer Query.Close()
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(reason_id, employee_login, participant_login)
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func getCurrentMarkValue(participant_login string, category_id int64) (int64, *conf.ApiResponse) {
	var value int64
	err := src.CustomConnection.QueryRow("SELECT `"+strconv.FormatInt(category_id, 10)+"` FROM participants WHERE login=?", participant_login).Scan(&value)
	if err != nil {
		return 0, conf.ErrDatabaseQueryFailed
	}
	return value, nil
}

func getReasonChange(reason_id int64) (int64, *conf.ApiResponse) {
	var change int64
	err := src.CustomConnection.QueryRow("SELECT modification FROM reasons WHERE id=?", reason_id).Scan(&change)
	if err != nil {
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
	_, employee_login, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
	if APIerr != nil {
		return APIerr.Print(w)
	}
	APIerr = checkUserAccess_Request(employee_login)
	if APIerr != nil {
		return APIerr.Print(w)
	}
	return true
}

func checkUserAccess_Request(employee_login string) *conf.ApiResponse {
	var access int
	err := src.CustomConnection.QueryRow("SELECT access FROM users WHERE login=? LIMIT 1", employee_login).Scan(&access)
	if err != nil {
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
		return conf.ErrDatabaseQueryFailed.Print(w)
	}
	if access == 0 {
		return true
	} else {
		return conf.ErrLoginIncorrect.Print(w)
	}
}