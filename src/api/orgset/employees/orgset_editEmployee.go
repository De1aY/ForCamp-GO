/*
	Copyright: "Null team", 2016 - 2017
	Author: "De1aY"
	Documentation: https://bitbucket.org/lyceumdevelopers/golang/wiki/Home
*/
package employees

import (
	"net/http"
	"forcamp/src/api/orgset"
	"forcamp/conf"
	"forcamp/src"
	"log"
)

func EditEmployee(token string, employee Employee, responseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token, responseWriter) {
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil {
			return APIerr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		if checkEditEmployeeData(employee, responseWriter) {
			EmployeeOrganization, APIerr := orgset.GetUserOrganizationByLogin(employee.Login)
			if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			if EmployeeOrganization != Organization {
				return conf.ErrUserNotFound.Print(responseWriter)
			}
			APIerr = editEmployee_Request(employee)
			return conf.RequestSuccess.Print(responseWriter)
		}
	}
	return true
}

func editEmployee_Request(employee Employee) *conf.ApiResponse{
	Query, err := src.CustomConnection.Prepare("UPDATE users SET team='0' WHERE team=? AND access='1'")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(employee.Team)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	Query, err = src.CustomConnection.Prepare("UPDATE users SET name=?, surname=?, middlename=?, team=?, sex=? WHERE login=? AND access='1'")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(employee.Name, employee.Surname, employee.Middlename, employee.Team, employee.Sex, employee.Login)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	Query, err = src.CustomConnection.Prepare("UPDATE employees SET post=? WHERE login=?")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(employee.Post, employee.Login)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func checkEditEmployeeData(employee Employee, w http.ResponseWriter) bool {
	if len(employee.Login) > 0 {
		if len(employee.Name) > 0 {
			if len(employee.Surname) > 0 {
				if len(employee.Middlename) > 0 {
					if len(employee.Post) > 0 {
						if employee.Sex == 0 || employee.Sex == 1 {
							if orgset.CheckTeamID(employee.Team, w) {
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
