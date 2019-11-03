/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package employees

import (
	"net/http"
	"nullteam.info/wplay/demo/src/api/orgset"
	"nullteam.info/wplay/demo/conf"
	"nullteam.info/wplay/demo/src"
)

type resetEmployeePassword_Success struct {
	Password string `json:"password"`
}


func ResetEmployeePassword(token string, employee_id int64, responseWriter http.ResponseWriter) bool{
	if orgset.IsUserAdmin(token, responseWriter){
		organizationName, _, apiErr := orgset.GetUserOrganizationAndIdByToken(token)
		if apiErr != nil{
			return apiErr.Print(responseWriter)
		}
		employee_organization, employee_login, apiErr := orgset.GetUserOrganizationAndLoginByID(employee_id)
		if apiErr != nil{
			return apiErr.Print(responseWriter)
		}
		if employee_organization != organizationName {
			return conf.ErrUserNotFound.Print(responseWriter)
		}
		rawResp, apiErr := resetEmployeePassword(employee_id, employee_login)
		if apiErr != nil{
			return apiErr.Print(responseWriter)
		}
		resp := conf.ApiResponse{200, "success", rawResp}
		resp.Print(responseWriter)
	}
	return true
}

func resetEmployeePassword(employee_id int64, employee_login string) (resetEmployeePassword_Success, *conf.ApiResponse){
	password, hash := orgset.GeneratePassword()
	query, err := src.Connection.Prepare("UPDATE users SET password=? WHERE id=?")
	if err != nil {
		return resetEmployeePassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(hash, employee_id)
	if err != nil {
		return resetEmployeePassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	query.Close()
	query, err = src.Connection.Prepare("DELETE FROM sessions WHERE login=?")
	if err != nil {
		return resetEmployeePassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(employee_login)
	if err != nil {
		return resetEmployeePassword_Success{}, conf.ErrDatabaseQueryFailed
	}
	return resetEmployeePassword_Success{password}, nil
}
