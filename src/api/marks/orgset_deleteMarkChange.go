package marks

import (
	"net/http"
	"forcamp/src/api/authorization"
	"forcamp/conf"
	"forcamp/src/api/orgset"
	"forcamp/src"
	"log"
	"strconv"
)

func DeleteMarkChange(token string, id int64, responseWriter http.ResponseWriter) bool {
	if authorization.CheckTokenForEmpty(token, responseWriter) {
		if authorization.CheckToken(token, responseWriter) {
			organization, login, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
			if APIerr != nil {
				return conf.PrintError(APIerr, responseWriter)
			}
			src.CustomConnection = src.Connect_Custom(organization)
			checkAdminPerm, APIerr := checkAdminPermissions(login)
			if APIerr != nil {
				return conf.PrintError(APIerr, responseWriter)
			}
			if checkAdminPerm {
				APIerr = deleteMarkChange(id)
				if APIerr != nil {
					return conf.PrintError(APIerr, responseWriter)
				}
			} else {
				APIerr = checkMarkChangeEmployee(login, id)
				if APIerr != nil {
					return conf.PrintError(APIerr, responseWriter)
				}
				APIerr = deleteMarkChange(id)
				if APIerr != nil {
					return conf.PrintError(APIerr, responseWriter)
				}
			}
			return conf.PrintSuccess(conf.RequestSuccess, responseWriter)
		} else {
			return conf.PrintError(conf.ErrUserTokenIncorrect, responseWriter)
		}
	}
	return true
}

func checkAdminPermissions(login string) (bool, *conf.ApiError) {
	var access int
	err := src.CustomConnection.QueryRow("SELECT access FROM users WHERE login=?", login).Scan(&access)
	if err != nil {
		log.Print(err)
		return false, conf.ErrDatabaseQueryFailed
	}
	if access == 2 {
		return true, nil
	} else {
		return false, nil
	}
}

func checkMarkChangeEmployee(employee_login string, id int64) *conf.ApiError {
	var count int
	err := src.CustomConnection.QueryRow("SELECT COUNT(id) FROM marks_changes WHERE employee_login=? AND id=?", employee_login, id).Scan(&count)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	if count == 1 {
		return nil
	} else {
		return conf.ErrInsufficientRights
	}
}

func deleteMarkChange(id int64) *conf.ApiError {
	var (
		reason_id int64
		participant_login string
	)
	err := src.CustomConnection.QueryRow("SELECT reason_id, participant_login FROM marks_changes WHERE id=?", id).Scan(&reason_id, &participant_login)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	reason_change, APIerr := getReasonChange(reason_id)
	if APIerr != nil {
		return APIerr
	}
	category_id, APIerr := getReasonCategoryID(reason_id)
	if APIerr != nil {
		return APIerr
	}
	current_mark, APIerr := getCurrentMarkValue(participant_login, category_id)
	APIerr = updateParticipantMark(participant_login, category_id, current_mark-reason_change)
	if APIerr != nil {
		return APIerr
	}
	query, err := src.CustomConnection.Prepare("DELETE FROM marks_changes WHERE id=?")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(id)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func updateParticipantMark(login string, category_id int64, newMark int) *conf.ApiError {
	query, err := src.CustomConnection.Prepare("UPDATE participants SET `"+strconv.FormatInt(category_id, 10)+"`=? WHERE login=?")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(newMark, login)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func getReasonCategoryID(reason_id int64) (int64, *conf.ApiError) {
	var category_id int64
	err := src.CustomConnection.QueryRow("SELECT cat_id FROM reasons WHERE id=?", reason_id).Scan(&category_id)
	if err != nil {
		log.Print(err)
		return 0, conf.ErrDatabaseQueryFailed
	}
	return category_id, nil
}

