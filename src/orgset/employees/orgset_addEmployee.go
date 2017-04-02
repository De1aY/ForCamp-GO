package employees

import (
	"net/http"
	"forcamp/src/orgset"
	"forcamp/conf"
	"forcamp/src"
	"log"
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/tealeg/xlsx"
	"database/sql"
)

type AddEmployee_Success struct{
	Code int `json:"code"`
	Status string `json:"status"`
	Login string `json:"login"`
}

func AddEmployee(token string, employee Employee, ResponseWriter http.ResponseWriter) bool {
	Connection := src.Connect()
	defer Connection.Close()
	if orgset.CheckUserAccess(token, Connection, ResponseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token, Connection)
		if APIerr != nil {
			return conf.PrintError(APIerr, ResponseWriter)
		}
		NewConnection := src.Connect_Custom(Organization)
		defer NewConnection.Close()
		if checkAddEmployeeData(employee, ResponseWriter, NewConnection) {
			Resp, APIerr := addEmployeeRequest(employee, Organization, Connection, NewConnection)
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			Response, _ := json.Marshal(Resp)
			fmt.Fprintf(ResponseWriter, string(Response))
		}
	}
	return true
}

func addEmployeeRequest(employee Employee, organization string, Connection *sql.DB, NewConnection *sql.DB) (AddEmployee_Success, *conf.ApiError){
	Password, Hash := orgset.GeneratePassword()
	login, APIerr := addEmployee_Main(organization, Connection, Hash)
	if APIerr != nil {
		return AddEmployee_Success{}, APIerr
	}
	employee.Login = login
	APIerr = addEmployee_Organization(employee, NewConnection)
	if APIerr != nil {
		return AddEmployee_Success{}, APIerr
	}
	APIerr = addEmployee_Excel(employee, organization, Password, NewConnection)
	if APIerr != nil {
		return AddEmployee_Success{}, APIerr
	}
	return AddEmployee_Success{200, "success", login}, nil
}

func addEmployee_Main(organization string, Connection *sql.DB, hash string) (string, *conf.ApiError){
	Query, err := Connection.Prepare("INSERT INTO users(password,organization) VALUES(?,?)")
	if err != nil {
		log.Print(err)
		return "", conf.ErrDatabaseQueryFailed
	}
	Resp, err := Query.Exec(hash, organization)
	if err != nil {
		log.Print(err)
		return "", conf.ErrDatabaseQueryFailed
	}
	ID, err := Resp.LastInsertId()
	if err != nil {
		log.Print(err)
		return "", conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	Query, err = Connection.Prepare("UPDATE users SET login=? WHERE id=?")
	if err != nil {
		log.Print(err)
		return "", conf.ErrDatabaseQueryFailed
	}
	login := "employee_"+strconv.FormatInt(ID, 10)
	_, err = Query.Exec(login, ID)
	if err != nil {
		log.Print(err)
		return "", conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	return login, nil
}

func addEmployee_Organization(employee Employee, Connection *sql.DB) *conf.ApiError{
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
	Query, err = Connection.Prepare("INSERT INTO users VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(employee.Login, employee.Name, employee.Surname, employee.Middlename, employee.Team, 1, employee.Sex, "default.jpg")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	Query, err = Connection.Prepare("INSERT INTO employees(login, post) VALUES(?, ?)")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(employee.Login, employee.Post)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	return nil
}

func addEmployee_Excel(employee Employee, organization string, password string, Connection *sql.DB) *conf.ApiError{
	teamName, APIerr := getTeamNameById(employee.Team, Connection)
	if APIerr != nil {
		return APIerr
	}
	excelFilePath := conf.FOLDER_EMPLOYEES+"/"+organization+".xlsx"
	xlFile, err := xlsx.OpenFile(excelFilePath)
	if err != nil {
		log.Print(err)
		return conf.ErrOpenExcelFile
	}
	sheet := xlFile.Sheets[0]
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = employee.Name
	cell = row.AddCell()
	cell.Value = employee.Surname
	cell = row.AddCell()
	cell.Value = employee.Middlename
	cell = row.AddCell()
	cell.Value = teamName
	cell = row.AddCell()
	cell.Value = employee.Login
	cell = row.AddCell()
	cell.Value = password
	err = xlFile.Save(excelFilePath)
	if err != nil {
		log.Print(err)
		return conf.ErrSaveExcelFile
	}
	return nil
}

func getTeamNameById(id int64, Connection *sql.DB) (string, *conf.ApiError){
	if id == 0{
		return "отуствует", nil
	} else {
		Query, err := Connection.Query("SELECT name FROM teams WHERE id=?", id)
		if err != nil {
			log.Print(err)
			return "", conf.ErrDatabaseQueryFailed
		}
		defer Query.Close()
		var name string
		for Query.Next(){
			err = Query.Scan(&name)
			if err != nil {
				log.Print(err)
				return "", conf.ErrDatabaseQueryFailed
			}
		}
		return name, nil
	}

}

func checkAddEmployeeData(employee Employee, w http.ResponseWriter, Connection *sql.DB) bool {
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
}
