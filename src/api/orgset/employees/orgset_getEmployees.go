package employees

import (
	"net/http"
	"forcamp/src/api/authorization"
	"forcamp/src/api/orgset"
	"forcamp/conf"
	"forcamp/src"
	"database/sql"
	"strconv"
	"forcamp/src/api/orgset/categories"
)

type Permission struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
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

type getEmployees_Success struct {
	Employees []Employee `json:"employees"`
}

func GetEmployees(token string, responseWriter http.ResponseWriter) bool {
	if authorization.CheckTokenForEmpty(token, responseWriter) {
		if authorization.CheckToken(token, responseWriter) {
			Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
			if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			src.CustomConnection = src.Connect_Custom(Organization)
			rawResp, APIerr := getEmployees_Request()
			if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			resp := conf.ApiResponse{200, "success", rawResp}
			resp.Print(responseWriter)
		} else {
			return conf.ErrUserTokenIncorrect.Print(responseWriter)
		}
	}
	return true
}

func getEmployees_Request() (getEmployees_Success, *conf.ApiResponse) {
	Query, err := src.CustomConnection.Query("SELECT login,name,surname,middlename,sex,team FROM users WHERE access='1'")
	if err != nil {
		return getEmployees_Success{}, conf.ErrDatabaseQueryFailed
	}
	return getEmployeesFromResponse(Query)
}

func getEmployeesFromResponse(rows *sql.Rows) (getEmployees_Success, *conf.ApiResponse) {
	defer rows.Close()
	Permissions, Posts, APIerr := getPermissionsAndPosts()
	if APIerr != nil {
		return getEmployees_Success{}, APIerr
	}
	var (
		employees []Employee
		employee Employee
	)
	for rows.Next() {
		err := rows.Scan(&employee.Login, &employee.Name, &employee.Surname, &employee.Middlename, &employee.Sex, &employee.Team)
		if err != nil {
			return getEmployees_Success{}, conf.ErrDatabaseQueryFailed
		}
		employee.Permissions = Permissions[employee.Login]
		employee.Post = Posts[employee.Login]
		employees = append(employees, Employee{employee.Login, employee.Name, employee.Surname, employee.Middlename, employee.Sex, employee.Team, employee.Post, employee.Permissions})
	}
	if employees == nil {
		return getEmployees_Success{make([]Employee, 0)}, nil
	}
	return getEmployees_Success{employees}, nil
}

func getPermissionsAndPosts() (map[string][]Permission, map[string]string, *conf.ApiResponse) {
	Query, err := src.CustomConnection.Query("SELECT * FROM employees")
	if err != nil {
		return make(map[string][]Permission), make(map[string]string), conf.ErrDatabaseQueryFailed
	}
	CategoriesIDs, err := Query.Columns()
	if err != nil {
		return make(map[string][]Permission), make(map[string]string), conf.ErrDatabaseQueryFailed
	}
	if len(CategoriesIDs) == 1 {
		return getPermissionsIfNoCategories(Query)
	}
	return getPermissionsIfCategories(Query, CategoriesIDs)
}

func getPermissionsIfNoCategories(rows *sql.Rows) (map[string][]Permission, map[string]string, *conf.ApiResponse) {
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
			return make(map[string][]Permission), make(map[string]string), conf.ErrDatabaseQueryFailed
		}
		Permissions[login] = make([]Permission, 0)
		Posts[login] = post
	}
	return Permissions, Posts, nil
}

func getPermissionsIfCategories(rows *sql.Rows, categoriesIDs []string) (map[string][]Permission, map[string]string,*conf.ApiResponse) {
	categoriesIDs = categoriesIDs[2:]
	var (
		rawResult = make([][]byte, len(categoriesIDs) + 2)
		result = make([]interface{}, len(categoriesIDs) + 2)
		permissions = make(map[string][]Permission)
		values = make([]string, len(categoriesIDs) + 2)
		posts = make(map[string]string)
	)
	for i, _ := range result {
		result[i] = &rawResult[i]
	}
	categoriesList, apiErr := categories.GetCategories_Request(); if apiErr != nil {
		return make(map[string][]Permission), make(map[string]string), apiErr
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(result...)
		if err != nil {
			return make(map[string][]Permission), make(map[string]string), conf.ErrDatabaseQueryFailed
		}
		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				values[i] = string(raw)
			}
		}
		Login := values[0]
		Post := values[1]
		posts[Login] = Post
		values = values[2:]
		for i := 0; i < len(values); i++ {
			id, err := strconv.ParseInt(categoriesIDs[i], 10, 64)
			if err != nil {
				return make(map[string][]Permission), make(map[string]string),conf.ErrConvertStringToInt
			}
			permissions[Login] = append(permissions[Login], Permission{id, categoriesList[i].Name, values[i]})
		}
		values = make([]string, len(categoriesIDs) + 2)
	}
	return permissions, posts, nil
}
