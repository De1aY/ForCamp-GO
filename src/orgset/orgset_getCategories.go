package orgset

import (
	"net/http"
	"forcamp/src/authorization"
	"forcamp/src"
	"forcamp/conf"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

type Category struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	NegativeMarks string `json:"negative_marks"`
}

type GetCategories_Success struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Categories []Category `json:"categories"`
}

func GetCategories(token string, ResponseWriter http.ResponseWriter) bool{
	if authorization.CheckTokenForEmpty(token, ResponseWriter) && authorization.CheckToken(token, ResponseWriter){
		Organization, _, err := getUserOrganizationAndLoginByToken(token)
		if err != nil{
			log.Print(err)
			return conf.PrintError(err, ResponseWriter)
		}
		NewConnection = src.Connect_Custom(Organization)
		Resp, err := getCategories_Request()
		if err != nil{
			log.Print(err)
			return conf.PrintError(err, ResponseWriter)
		}
		Response, _ := json.Marshal(Resp)
		fmt.Fprintf(ResponseWriter, string(Response))
	}
	return true
}

func getCategories_Request() (GetCategories_Success, *conf.ApiError){
	Query, err := NewConnection.Query("SELECT * FROM categories")
	if err != nil {
		log.Print(err)
		return GetCategories_Success{}, conf.ErrDatabaseQueryFailed
	}
	return getCategoriesFromQuery(Query)
}

func getCategoriesFromQuery(rows *sql.Rows) (GetCategories_Success, *conf.ApiError){
	defer rows.Close()
	var categories []Category
	var category Category
	for rows.Next(){
		err := rows.Scan(&category.ID, &category.Name, &category.NegativeMarks)
		if err != nil{
			log.Print(err)
			return GetCategories_Success{}, conf.ErrDatabaseQueryFailed
		}
		categories = append(categories, Category{category.ID, category.Name, category.NegativeMarks})
	}
	return GetCategories_Success{200, "success", categories}, nil
}