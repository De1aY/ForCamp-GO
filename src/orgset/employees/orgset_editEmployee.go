package employees

import (
	"net/http"
	"forcamp/src/orgset"
	"forcamp/conf"
	"forcamp/src"
	"log"
	"database/sql"
)

func EditEmployee(token string, employee Employee, ResponseWriter http.ResponseWriter) bool{
	Connection := src.Connect()
	defer Connection.Close()
	if orgset.CheckUserAccess(token, Connection, ResponseWriter) {
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token, Connection)
		if APIerr != nil {
			return conf.PrintError(APIerr, ResponseWriter)
		}
		NewConnection := src.Connect_Custom(Organization)
		defer NewConnection.Close()
		if checkEditEmployeeData(employee, ResponseWriter, NewConnection) {
			EmployeeOrganization, APIerr := orgset.GetUserOrganizationByLogin(employee.Login, Connection)
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			if EmployeeOrganization != Organization {
				return conf.PrintError(conf.ErrUserNotFound, ResponseWriter)
			}
			APIerr = editEmployee_Request(employee, NewConnection)
			return conf.PrintSuccess(conf.RequestSuccess, ResponseWriter)
		}
	}
	return true
}

func editEmployee_Request(employee Employee, Connection *sql.DB) *conf.ApiError{
	Query, err := Connection.Prepare("UPDATE users SET team='0' WHERE team=? AND access='1'")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(employee.Team)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	Query, err = Connection.Prepare("UPDATE users SET name=?, surname=?, middlename=?, team=?, sex=? WHERE login=? AND access='1'")
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
	Query, err = Connection.Prepare("UPDATE employees SET post=? WHERE login=?")
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

func checkEditEmployeeData(employee Employee, w http.ResponseWriter, Connection *sql.DB) bool {
	if len(employee.Login) > 0 {
		if len(employee.Name) > 0 {
			if len(employee.Surname) > 0 {
				if len(employee.Middlename) > 0 {
					if len(employee.Post) > 0 {
						if employee.Sex == 0 || employee.Sex == 1 {
							if orgset.CheckTeamID(employee.Team, w, Connection) {
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
