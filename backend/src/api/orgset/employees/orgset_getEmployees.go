package employees

import (
	"net/http"
	"wplay/src/api/authorization"
	"wplay/src/api/orgset"
	"wplay/conf"
	"wplay/src"
	"database/sql"
	"strconv"
	"wplay/src/api/orgset/categories"
)

type Permission struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Employee struct {
	ID          int64        `json:"id"`
	Name        string       `json:"name"`
	Surname     string       `json:"surname"`
	Middlename  string       `json:"middlename"`
	Sex         int          `json:"sex"`
	Team        int64        `json:"team"`
	Post        string       `json:"post"`
	Permissions []Permission `json:"permissions"`
}

type getEmployees_Success struct {
	Employees []Employee `json:"employees"`
}

func GetEmployees(token string, responseWriter http.ResponseWriter) bool {
	if authorization.IsTokenNotEmpty(token, responseWriter) {
		if authorization.IsTokenValid(token, responseWriter) {
			organizationName, _, apiErr := orgset.GetUserOrganizationAndIdByToken(token)
			if apiErr != nil {
				return apiErr.Print(responseWriter)
			}
			src.CustomConnection = src.Connect_Custom(organizationName)
			rawResp, apiErr := getEmployees_Request()
			if apiErr != nil {
				return apiErr.Print(responseWriter)
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
	query, err := src.CustomConnection.Query("SELECT id,name,surname,middlename," +
		"sex,team FROM users WHERE access='1'")
	if err != nil {
		return getEmployees_Success{}, conf.ErrDatabaseQueryFailed
	}
	return getEmployeesFromResponse(query)
}

func getEmployeesFromResponse(rows *sql.Rows) (getEmployees_Success, *conf.ApiResponse) {
	defer rows.Close()
	permissions, posts, apiErr := getPermissionsAndPosts()
	if apiErr != nil {
		return getEmployees_Success{}, apiErr
	}
	var (
		employees []Employee
		employee Employee
	)
	for rows.Next() {
		err := rows.Scan(&employee.ID, &employee.Name, &employee.Surname,
			&employee.Middlename, &employee.Sex, &employee.Team)
		if err != nil {
			return getEmployees_Success{}, conf.ErrDatabaseQueryFailed
		}
		employee.Permissions = permissions[employee.ID]
		employee.Post = posts[employee.ID]
		employees = append(employees, Employee{employee.ID, employee.Name, employee.Surname,
			employee.Middlename, employee.Sex, employee.Team,
			employee.Post, employee.Permissions})
	}
	if employees == nil {
		return getEmployees_Success{make([]Employee, 0)}, nil
	}
	return getEmployees_Success{employees}, nil
}

func getPermissionsAndPosts() (map[int64][]Permission, map[int64]string, *conf.ApiResponse) {
	Query, err := src.CustomConnection.Query("SELECT * FROM employees")
	if err != nil {
		return make(map[int64][]Permission), make(map[int64]string), conf.ErrDatabaseQueryFailed
	}
	CategoriesIDs, err := Query.Columns()
	if err != nil {
		return make(map[int64][]Permission), make(map[int64]string), conf.ErrDatabaseQueryFailed
	}
	if len(CategoriesIDs) == 1 {
		return getPermissionsIfNoCategories(Query)
	}
	return getPermissionsIfCategories(Query, CategoriesIDs)
}

func getPermissionsIfNoCategories(rows *sql.Rows) (map[int64][]Permission, map[int64]string, *conf.ApiResponse) {
	defer rows.Close()
	var (
		employee_id int64
		post        string
		permissions = make(map[int64][]Permission)
		posts       = make(map[int64]string)
	)
	for rows.Next() {
		err := rows.Scan(&employee_id, &post); if err != nil {
			return permissions, posts, conf.ErrDatabaseQueryFailed
		}
		permissions[employee_id] = make([]Permission, 0)
		posts[employee_id] = post
	}
	return permissions, posts, nil
}

func getPermissionsIfCategories(rows *sql.Rows, categoriesIDs []string) (map[int64][]Permission,
	map[int64]string,*conf.ApiResponse) {
	categoriesIDs = categoriesIDs[2:]
	var (
		rawResult = make([][]byte, len(categoriesIDs) + 2)
		result = make([]interface{}, len(categoriesIDs) + 2)
		permissions = make(map[int64][]Permission)
		values = make([]string, len(categoriesIDs) + 2)
		posts = make(map[int64]string)
	)
	for i, _ := range result {
		result[i] = &rawResult[i]
	}
	categoriesList, apiErr := categories.GetCategories_Request(); if apiErr != nil {
		return permissions, posts, apiErr
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(result...); if err != nil {
			return permissions, posts, conf.ErrDatabaseQueryFailed
		}
		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				values[i] = string(raw)
			}
		}
		employee_id, err := strconv.ParseInt(values[0], 10, 64); if err != nil {
			return permissions, posts, conf.ErrConvertStringToInt
		}
		employee_post := values[1]
		posts[employee_id] = employee_post
		values = values[2:]
		for i := 0; i < len(values); i++ {
			id, err := strconv.ParseInt(categoriesIDs[i], 10, 64); if err != nil {
				return permissions, posts, conf.ErrConvertStringToInt
			}
			permissions[employee_id] = append(permissions[employee_id],
				Permission{id, categoriesList[i].Name, values[i]})
		}
		values = make([]string, len(categoriesIDs) + 2)
	}
	return permissions, posts, nil
}
