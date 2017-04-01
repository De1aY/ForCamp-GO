package employees

import (
	"net/http"
	"forcamp/src/authorization"
	"forcamp/src/orgset"
	"forcamp/conf"
	"forcamp/src"
	"log"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
)

type Permission struct {
	Id    int64 `json:"id"`
	Value string `json:"value"`
}

type Employee struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Middlename  string `json:"middlename"`
	Sex         int `json:"sex"`
	Team        int `json:"team"`
	Permissions []Permission `json:"permissions"`
}

type GetEmployees_Success struct {
	Code      int `json:"code"`
	Status    string `json:"status"`
	Employees []Employee `json:"employees"`
}

func GetEmployees(token string, ResponseWriter http.ResponseWriter) bool {
	if authorization.CheckTokenForEmpty(token, ResponseWriter) {
		if authorization.CheckToken(token, ResponseWriter) {
			Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			src.NewConnection = src.Connect_Custom(Organization)
			Resp, APIerr := getEmployees_Request()
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			Response, _ := json.Marshal(Resp)
			fmt.Fprintf(ResponseWriter, string(Response))
		} else {
			return conf.PrintError(conf.ErrUserTokenIncorrect, ResponseWriter)
		}
	}
	return true
}

func getEmployees_Request() (GetEmployees_Success, *conf.ApiError) {
	Query, err := src.NewConnection.Query("SELECT login,name,surname,middlename,sex,team FROM users WHERE access='1'")
	if err != nil {
		log.Print(err)
		return GetEmployees_Success{}, conf.ErrDatabaseQueryFailed
	}
	return getEmployeesFromResponse(Query)
}

func getEmployeesFromResponse(rows *sql.Rows) (GetEmployees_Success, *conf.ApiError) {
	defer rows.Close()
	Permissions, APIerr := getPermissions()
	if APIerr != nil {
		return GetEmployees_Success{}, APIerr
	}
	var (
		employees []Employee
		employee Employee
	)
	for rows.Next() {
		err := rows.Scan(&employee.Login, &employee.Name, &employee.Surname, &employee.Middlename, &employee.Sex, &employee.Team)
		if err != nil {
			log.Print(err)
			return GetEmployees_Success{}, conf.ErrDatabaseQueryFailed
		}
		employee.Permissions = Permissions[employee.Login]
		employees = append(employees, Employee{employee.Login, employee.Name, employee.Surname, employee.Middlename, employee.Sex, employee.Team, employee.Permissions})
	}
	return GetEmployees_Success{200, "success", employees}, nil
}

func getPermissions() (map[string][]Permission, *conf.ApiError) {
	Query, err := src.NewConnection.Query("SELECT * FROM employees")
	if err != nil {
		log.Print(err)
		return make(map[string][]Permission), conf.ErrDatabaseQueryFailed
	}
	CategoriesIDs, err := Query.Columns()
	if err != nil {
		log.Print(err)
		return make(map[string][]Permission), conf.ErrDatabaseQueryFailed
	}
	if len(CategoriesIDs) == 1 {
		return getPermissionsIfNoCategories(Query)
	}
	return getPermissionsIfCategories(Query, CategoriesIDs)
}

func getPermissionsIfNoCategories(rows *sql.Rows) (map[string][]Permission, *conf.ApiError) {
	defer rows.Close()
	var (
		login string
		Permissions = make(map[string][]Permission)
	)
	for rows.Next() {
		err := rows.Scan(&login)
		if err != nil {
			log.Print(err)
			return make(map[string][]Permission), conf.ErrDatabaseQueryFailed
		}
		Permissions[login] = make([]Permission, 0)
	}
	return Permissions, nil
}

func getPermissionsIfCategories(rows *sql.Rows, CategoriesIDs []string) (map[string][]Permission, *conf.ApiError) {
	CategoriesIDs = CategoriesIDs[1:]
	var (
		rawResult = make([][]byte, len(CategoriesIDs) + 1)
		Result = make([]interface{}, len(CategoriesIDs) + 1)
		Permissions = make(map[string][]Permission)
		Values = make([]string, len(CategoriesIDs) + 1)
	)
	for i, _ := range Result {
		Result[i] = &rawResult[i]
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(Result...)
		if err != nil {
			return make(map[string][]Permission), conf.ErrDatabaseQueryFailed
		}
		for i, raw := range rawResult {
			if raw == nil {
				Result[i] = "\\N"
			} else {
				Values[i] = string(raw)
			}
		}
		Login := Values[0]
		Values = Values[1:]
		for i := 0; i < len(Values); i++ {
			id, err := strconv.ParseInt(CategoriesIDs[i], 10, 64)
			if err != nil {
				log.Print(err)
				return make(map[string][]Permission), conf.ErrConvertStringToInt
			}
			Permissions[Login] = append(Permissions[Login], Permission{id, Values[i]})
		}
	}
	return Permissions, nil
}
