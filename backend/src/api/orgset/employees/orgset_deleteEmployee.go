package employees

import (
	"net/http"
	"wplay/conf"
	"wplay/src"
	"wplay/src/api/orgset"
)

func DeleteEmployee(token string, employee_id int64, responseWriter http.ResponseWriter) bool {
	if orgset.IsUserAdmin(token, responseWriter) {
		organizationName, _, apiErr := orgset.GetUserOrganizationAndIdByToken(token); if apiErr != nil {
			return apiErr.Print(responseWriter)
		}
		employee_organization, employee_login, apiErr := orgset.GetUserOrganizationAndLoginByID(employee_id)
		if apiErr != nil{
			return apiErr.Print(responseWriter)
		}
		if employee_organization != organizationName {
			return conf.ErrUserNotFound.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(organizationName)
		apiErr = deleteEmployee(employee_id, employee_login)
		if apiErr != nil {
			return apiErr.Print(responseWriter)
		}
		return conf.RequestSuccess.Print(responseWriter)
	}
	return true
}

func deleteEmployee(employee_id int64, employee_login string) *conf.ApiResponse {
	apiErr := deleteEmployee_Organization(employee_id)
	if apiErr != nil {
		return apiErr
	}
	apiErr = deleteEmployee_Main(employee_id, employee_login)
	if apiErr != nil {
		return apiErr
	}
	return nil
}

func deleteEmployee_Main(employee_id int64, employee_login string) *conf.ApiResponse {
	query, err := src.Connection.Prepare("DELETE FROM users WHERE id=?"); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(employee_id); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	query, err = src.Connection.Prepare("DELETE FROM sessions WHERE login=?"); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(employee_login); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func deleteEmployee_Organization(employee_id int64) *conf.ApiResponse {
	query, err := src.CustomConnection.Prepare("DELETE FROM users WHERE id=? AND access='1'"); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	resp, err := query.Exec(employee_id); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	query.Close()
	rowsAffected, err := resp.RowsAffected(); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	if rowsAffected == 0 {
		return conf.ErrUserNotFound
	}
	query, err = src.CustomConnection.Prepare("DELETE FROM employees WHERE id=?"); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(employee_id); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}
