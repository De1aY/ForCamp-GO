package categories

import (
	"net/http"
	"forcamp/src/authorization"
	"forcamp/src"
	"forcamp/conf"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"forcamp/src/orgset"
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

func GetCategories(token string, ResponseWriter http.ResponseWriter) bool {
	Connection := src.Connect()
	defer Connection.Close()
	if authorization.CheckTokenForEmpty(token, ResponseWriter) {
		if authorization.CheckToken(token, Connection, ResponseWriter) {
			Organization, _, err := orgset.GetUserOrganizationAndLoginByToken(token, Connection)
			if err != nil {
				log.Print(err)
				return conf.PrintError(err, ResponseWriter)
			}
			NewConnection := src.Connect_Custom(Organization)
			defer NewConnection.Close()
			Resp, err := getCategories_Request(NewConnection)
			if err != nil {
				log.Print(err)
				return conf.PrintError(err, ResponseWriter)
			}
			Response, _ := json.Marshal(Resp)
			fmt.Fprintf(ResponseWriter, string(Response))
		} else {
			return conf.PrintError(conf.ErrUserTokenIncorrect, ResponseWriter)
		}
	}
	return true
}

func getCategories_Request(Connection *sql.DB) (GetCategories_Success, *conf.ApiError){
	Query, err := Connection.Query("SELECT * FROM categories")
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