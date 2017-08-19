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
	"strconv"
	"github.com/tealeg/xlsx"
)

type addEmployee_Success struct{
	ID int64 `json:"id"`
}


func AddEmployee(token string, employee Employee, responseWriter http.ResponseWriter) bool {
	if orgset.IsUserAdmin(token, responseWriter){
		organizationName, _, apiErr := orgset.GetUserOrganizationAndIdByToken(token); if apiErr != nil {
			return apiErr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(organizationName)
		if isAddEmployeeDataValid(employee, responseWriter) {
			rawResp, APIerr := addEmployeeRequest(employee, organizationName); if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			resp := conf.ApiResponse{200, "success", rawResp}
			resp.Print(responseWriter)
		}
	}
	return true
}

func addEmployeeRequest(employee Employee, organization string) (addEmployee_Success, *conf.ApiResponse){
	password, hash := orgset.GeneratePassword()
	employee_id, employee_login, apiErr := addEmployee_Main(organization, hash); if apiErr != nil {
		return addEmployee_Success{}, apiErr
	}
	employee.ID = employee_id
	apiErr = addEmployee_Organization(employee); if apiErr != nil {
		return addEmployee_Success{}, apiErr
	}
	apiErr = addEmployee_Excel(employee, employee_login, organization, password); if apiErr != nil {
		return addEmployee_Success{}, apiErr
	}
	return addEmployee_Success{employee_id}, nil
}

func addEmployee_Main(organization string, hash string) (int64, string, *conf.ApiResponse){
	query, err := src.Connection.Prepare("INSERT INTO users(password,organization) VALUES(?,?)"); if err != nil {
		return 0, "", conf.ErrDatabaseQueryFailed
	}
	resp, err := query.Exec(hash, organization); if err != nil {
		return 0, "", conf.ErrDatabaseQueryFailed
	}
	employee_id, err := resp.LastInsertId(); if err != nil {
		return 0, "", conf.ErrDatabaseQueryFailed
	}
	query.Close()
	query, err = src.Connection.Prepare("UPDATE users SET login=? WHERE id=?"); if err != nil {
		return 0, "", conf.ErrDatabaseQueryFailed
	}
	employee_login := "employee_"+strconv.FormatInt(employee_id, 10)
	_, err = query.Exec(employee_login, employee_id); if err != nil {
		return 0, "", conf.ErrDatabaseQueryFailed
	}
	query.Close()
	return employee_id, employee_login, nil
}

func addEmployee_Organization(employee Employee) *conf.ApiResponse{
	query, err := src.CustomConnection.Prepare("UPDATE users SET team='0' WHERE team=? AND access='1'")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(employee.Team); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	query, err = src.CustomConnection.Prepare("INSERT INTO users(id,name,surname," +
		"middlename,team,access,sex,avatar) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(employee.ID, employee.Name, employee.Surname,
		employee.Middlename, employee.Team, 1, employee.Sex, "default.jpg")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	query.Close()
	query, err = src.CustomConnection.Prepare("INSERT INTO employees(id, post) VALUES(?, ?)"); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(employee.ID, employee.Post); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	query.Close()
	return nil
}

func addEmployee_Excel(employee Employee, employee_login string, organization string, password string) *conf.ApiResponse{
	teamName, apiErr := getTeamNameById(employee.Team); if apiErr != nil {
		return apiErr
	}
	excelFilePath := conf.FOLDER_EMPLOYEES + "/" + organization + ".xlsx"
	xlFile, err := xlsx.OpenFile(excelFilePath); if err != nil {
		return conf.ErrOpenExcelFile
	}
	sheet := xlFile.Sheets[0]
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = employee.Surname
	cell = row.AddCell()
	cell.Value = employee.Name
	cell = row.AddCell()
	cell.Value = employee.Middlename
	cell = row.AddCell()
	cell.Value = teamName
	cell = row.AddCell()
	cell.Value = employee_login
	cell = row.AddCell()
	cell.Value = password
	err = xlFile.Save(excelFilePath); if err != nil {
		return conf.ErrSaveExcelFile
	}
	return nil
}

func getTeamNameById(id int64) (string, *conf.ApiResponse){
	if id == 0{
		return "отуствует", nil
	} else {
		query, err := src.CustomConnection.Query("SELECT name FROM teams WHERE id=?", id); if err != nil {
			return "", conf.ErrDatabaseQueryFailed
		}
		defer query.Close()
		var name string
		for query.Next(){
			err = query.Scan(&name); if err != nil {
				return "", conf.ErrDatabaseQueryFailed
			}
		}
		return name, nil
	}

}

func isAddEmployeeDataValid(employee Employee, w http.ResponseWriter) bool {
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
}
