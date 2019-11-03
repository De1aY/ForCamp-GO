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

func EditEmployee(token string, employee Employee, responseWriter http.ResponseWriter) bool{
	if orgset.IsUserAdmin(token, responseWriter) {
		organizationName, _, apiErr := orgset.GetUserOrganizationAndIdByToken(token); if apiErr != nil {
			return apiErr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(organizationName)
		if isEditEmployeeDataCorrect(employee, responseWriter) {
			employeeOrganization, apiErr := orgset.GetUserOrganizationByID(employee.ID); if apiErr != nil {
				return apiErr.Print(responseWriter)
			}
			if employeeOrganization != organizationName {
				return conf.ErrUserNotFound.Print(responseWriter)
			}
			apiErr = editEmployee(employee)
			return conf.RequestSuccess.Print(responseWriter)
		}
	}
	return true
}

func editEmployee(employee Employee) *conf.ApiResponse{
	query, err := src.CustomConnection.Prepare("UPDATE users SET team='0' WHERE team=? AND access='1'")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(employee.Team); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	query, err = src.CustomConnection.Prepare("UPDATE users SET name=?, surname=?, middlename=?, " +
		"team=?, sex=? WHERE id=? AND access='1'")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(employee.Name, employee.Surname, employee.Middlename, employee.Team, employee.Sex, employee.ID)
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	query.Close()
	query, err = src.CustomConnection.Prepare("UPDATE employees SET post=? WHERE id=?"); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(employee.Post, employee.ID)
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func isEditEmployeeDataCorrect(employee Employee, w http.ResponseWriter) bool {
	if employee.ID > 0 {
		if len(employee.Name) > 0 {
			if len(employee.Surname) > 0 {
				if len(employee.Middlename) > 0 {
					if len(employee.Post) > 0 {
						if employee.Sex == 0 || employee.Sex == 1 {
							if orgset.IsTeamExist(employee.Team, w) {
								return true
							} else {
								return false
							}
						} else {
							return conf.ErrSexIncorrect.Print(w)
						}
					} else {
						return conf.ErrPostEmpty.Print(w)
					}
				} else {
					return conf.ErrMiddlenameEmpty.Print(w)
				}
			} else {
				return conf.ErrSurnameEmpty.Print(w)
			}
		} else {
			return conf.ErrNameEmpty.Print(w)
		}
	} else {
		return conf.ErrUserNotFound.Print(w)
	}
}
