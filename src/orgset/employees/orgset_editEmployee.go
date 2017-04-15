/*
	Copyright: "Null team", 2016 - 2017
	Author: "De1aY"
	Documentation: https://bitbucket.org/lyceumdevelopers/golang/wiki/Home
*/
package employees

import (
	"net/http"
	"forcamp/src/orgset"
	"forcamp/conf"
	"forcamp/src"
	"log"
)

func EditEmployee(token string, employee Employee, ResponseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token, ResponseWriter) {
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil {
			return conf.PrintError(APIerr, ResponseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		if checkEditEmployeeData(employee, ResponseWriter) {
			EmployeeOrganization, APIerr := orgset.GetUserOrganizationByLogin(employee.Login)
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			if EmployeeOrganization != Organization {
				return conf.PrintError(conf.ErrUserNotFound, ResponseWriter)
			}
			APIerr = editEmployee_Request(employee)
			return conf.PrintSuccess(conf.RequestSuccess, ResponseWriter)
		}
	}
	return true
}

func editEmployee_Request(employee Employee) *conf.ApiError{
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
							return conf.PrintError(conf.ErrEmployeeSexIncorrect, w)
						}
					} else {
						return conf.PrintError(conf.ErrEmployeePostEmpty, w)
					}
				} else {
					return conf.PrintError(conf.ErrEmployeeMiddlenameEmpty, w)
				}
			} else {
				return conf.PrintError(conf.ErrEmployeeSurnameEmpty, w)
			}
		} else {
			return conf.PrintError(conf.ErrEmployeeNameEmpty, w)
		}
	} else {
		return conf.PrintError(conf.ErrUserNotFound, w)
	}
}
