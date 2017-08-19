package categories

import (
	"net/http"
	"forcamp/src/api/authorization"
	"forcamp/src"
	"forcamp/conf"
	"database/sql"
	"forcamp/src/api/orgset"
)

type Category struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	NegativeMarks string `json:"negative_marks"`
}

type getCategories_Success struct {
	Categories []Category `json:"categories"`
}


func GetCategories(token string, responseWriter http.ResponseWriter) bool {
	if authorization.IsTokenNotEmpty(token, responseWriter) {
		if authorization.IsTokenValid(token, responseWriter) {
			Organization, _, apiErr := orgset.GetUserOrganizationAndIdByToken(token)
			if apiErr != nil {
				return apiErr.Print(responseWriter)
			}
			src.CustomConnection = src.Connect_Custom(Organization)
			rawResp, apiErr := GetCategories_Request()
			if apiErr != nil {
				return apiErr.Print(responseWriter)
			}
			resp := &conf.ApiResponse{200, "success", getCategories_Success{rawResp}}
			resp.Print(responseWriter)
		} else {
			return conf.ErrUserTokenIncorrect.Print(responseWriter)
		}
	}
	return true
}

func GetCategories_Request() ([]Category, *conf.ApiResponse){
	query, err := src.CustomConnection.Query("SELECT * FROM categories")
	if err != nil {
		return nil, conf.ErrDatabaseQueryFailed
	}
	return getCategoriesFromQuery(query)
}

func getCategoriesFromQuery(rows *sql.Rows) ([]Category, *conf.ApiResponse){
	defer rows.Close()
	var categories []Category
	var category Category
	for rows.Next(){
		err := rows.Scan(&category.ID, &category.Name, &category.NegativeMarks)
		if err != nil{
			return nil, conf.ErrDatabaseQueryFailed
		}
		categories = append(categories, Category{category.ID, category.Name, category.NegativeMarks})
	}
	if categories == nil {
		return make([]Category, 0), nil
	}
	return categories, nil
}