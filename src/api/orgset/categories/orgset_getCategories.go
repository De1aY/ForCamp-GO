package categories

import (
	"net/http"
	"forcamp/src/api/authorization"
	"forcamp/src"
	"forcamp/conf"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"forcamp/src/api/orgset"
)

type Category struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	NegativeMarks string `json:"negative_marks"`
}

type getCategories_Success struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Categories []Category `json:"categories"`
}

func (success *getCategories_Success) toJSON() string {
	resp, _ := json.Marshal(success)
	return string(resp)
}


func GetCategories(token string, ResponseWriter http.ResponseWriter) bool {
	if authorization.CheckTokenForEmpty(token, ResponseWriter) {
		if authorization.CheckToken(token, ResponseWriter) {
			Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			src.CustomConnection = src.Connect_Custom(Organization)
			resp, APIerr := getCategories_Request()
			if APIerr != nil {
				return conf.PrintError(APIerr, ResponseWriter)
			}
			fmt.Fprintf(ResponseWriter, resp.toJSON())
		} else {
			return conf.PrintError(conf.ErrUserTokenIncorrect, ResponseWriter)
		}
	}
	return true
}

func getCategories_Request() (getCategories_Success, *conf.ApiError){
	Query, err := src.CustomConnection.Query("SELECT * FROM categories")
	if err != nil {
		log.Print(err)
		return getCategories_Success{}, conf.ErrDatabaseQueryFailed
	}
	return getCategoriesFromQuery(Query)
}

func getCategoriesFromQuery(rows *sql.Rows) (getCategories_Success, *conf.ApiError){
	defer rows.Close()
	var categories []Category
	var category Category
	for rows.Next(){
		err := rows.Scan(&category.ID, &category.Name, &category.NegativeMarks)
		if err != nil{
			log.Print(err)
			return getCategories_Success{}, conf.ErrDatabaseQueryFailed
		}
		categories = append(categories, Category{category.ID, category.Name, category.NegativeMarks})
	}
	if categories == nil {
		return getCategories_Success{200, "success", make([]Category, 0)}, nil
	}
	return getCategories_Success{200, "success", categories}, nil
}