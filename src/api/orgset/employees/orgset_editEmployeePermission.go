/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package employees

import (
	"net/http"
	"nullteam.info/wplay/demo/src"
	"nullteam.info/wplay/demo/src/api/orgset"
	"nullteam.info/wplay/demo/conf"
	"strconv"
)

func EditEmployeePermission(token string, employee_id int64, category_id int64,
	value string, responseWriter http.ResponseWriter) bool{
	if orgset.IsUserAdmin(token, responseWriter) {
		organizationName, _, apiErr := orgset.GetUserOrganizationAndIdByToken(token); if apiErr != nil {
			return apiErr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(organizationName)
		employee_organization, apiErr := orgset.GetUserOrganizationByID(employee_id); if apiErr != nil {
			return apiErr.Print(responseWriter)
		}
		if employee_organization != organizationName {
			return conf.ErrUserNotFound.Print(responseWriter)
		}
		if orgset.IsCategoryExist(category_id, responseWriter) && isPermissionValueCorrect(value, responseWriter){
			apiErr = editEmployeePermission(employee_id, category_id, value)
			return conf.RequestSuccess.Print(responseWriter)
		}
	}
	return true
}

func editEmployeePermission(employee_id int64, category_id int64, permission_value string) *conf.ApiResponse{
	query, err := src.CustomConnection.Prepare("UPDATE employees SET `"+strconv.FormatInt(category_id, 10)+
		"`=? WHERE id=?")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(permission_value, employee_id); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func isPermissionValueCorrect(value string, w http.ResponseWriter) bool{
	if value == "false" || value == "true"{
		return true
	} else {
		return conf.ErrPermissionValueIncorrect.Print(w)
	}
}
