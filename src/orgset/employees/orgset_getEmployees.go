/*
	Copyright: "Null team", 2016 - 2017
	Author: "De1aY"
	Documentation: https://bitbucket.org/lyceumdevelopers/golang/wiki/Home
*/
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
	Team        int64 `json:"team"`
	Post        string `json:"post"`
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
			src.CustomConnection = src.Connect_Custom(Organization)
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
	Query, err := src.CustomConnection.Query("SELECT login,name,surname,middlename,sex,team FROM users WHERE access='1'")
	if err != nil {
		log.Print(err)
		return GetEmployees_Success{}, conf.ErrDatabaseQueryFailed
	}
	return getEmployeesFromResponse(Query)
}

func getEmployeesFromResponse(rows *sql.Rows) (GetEmployees_Success, *conf.ApiError) {
	defer rows.Close()
	Permissions, Posts, APIerr := getPermissionsAndPosts()
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
		employee.Post = Posts[employee.Login]
		employees = append(employees, Employee{employee.Login, employee.Name, employee.Surname, employee.Middlename, employee.Sex, employee.Team, employee.Post, employee.Permissions})
	}
	if employees == nil {
		return GetEmployees_Success{200, "success", make([]Employee, 0)}, nil
	}
	return GetEmployees_Success{200, "success", employees}, nil
}

func getPermissionsAndPosts() (map[string][]Permission, map[string]string, *conf.ApiError) {
	Query, err := src.CustomConnection.Query("SELECT * FROM employees")
	if err != nil {
		log.Print(err)
		return make(map[string][]Permission), make(map[string]string), conf.ErrDatabaseQueryFailed
	}
	CategoriesIDs, err := Query.Columns()
	if err != nil {
		log.Print(err)
		return make(map[string][]Permission), make(map[string]string), conf.ErrDatabaseQueryFailed
	}
	if len(CategoriesIDs) == 1 {
		return getPermissionsIfNoCategories(Query)
	}
	return getPermissionsIfCategories(Query, CategoriesIDs)
}

func getPermissionsIfNoCategories(rows *sql.Rows) (map[string][]Permission, map[string]string, *conf.ApiError) {
	defer rows.Close()
	var (
		login string
		post string
		Permissions = make(map[string][]Permission)
		Posts = make(map[string]string)
	)
	for rows.Next() {
		err := rows.Scan(&login, &post)
		if err != nil {
			log.Print(err)
			return make(map[string][]Permission), make(map[string]string), conf.ErrDatabaseQueryFailed
		}
		Permissions[login] = make([]Permission, 0)
		Posts[login] = post
	}
	return Permissions, Posts, nil
}

func getPermissionsIfCategories(rows *sql.Rows, CategoriesIDs []string) (map[string][]Permission, map[string]string,*conf.ApiError) {
	CategoriesIDs = CategoriesIDs[2:]
	var (
		rawResult = make([][]byte, len(CategoriesIDs) + 2)
		Result = make([]interface{}, len(CategoriesIDs) + 2)
		Permissions = make(map[string][]Permission)
		Values = make([]string, len(CategoriesIDs) + 2)
		Posts = make(map[string]string)
	)
	for i, _ := range Result {
		Result[i] = &rawResult[i]
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(Result...)
		if err != nil {
			return make(map[string][]Permission), make(map[string]string),conf.ErrDatabaseQueryFailed
		}
		for i, raw := range rawResult {
			if raw == nil {
				Result[i] = "\\N"
			} else {
				Values[i] = string(raw)
			}
		}
		Login := Values[0]
		Post := Values[1]
		Posts[Login] = Post
		Values = Values[2:]
		for i := 0; i < len(Values); i++ {
			id, err := strconv.ParseInt(CategoriesIDs[i], 10, 64)
			if err != nil {
				log.Print(err)
				return make(map[string][]Permission), make(map[string]string),conf.ErrConvertStringToInt
			}
			Permissions[Login] = append(Permissions[Login], Permission{id, Values[i]})
		}
		Values = make([]string, len(CategoriesIDs) + 2)
	}
	return Permissions, Posts, nil
}
