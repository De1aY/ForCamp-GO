package categories

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src"
	"strconv"
	"encoding/json"
	"fmt"
	"log"
	"forcamp/src/orgset"
	"database/sql"
)

type AddCategory_Success struct {
	Code int `json:"code"`
	Status string `json:"status"`
	ID int64 `json:"id"`
}

func AddCategory(token string, category Category, ResponseWriter http.ResponseWriter) bool{
	Connection := src.Connect()
	defer Connection.Close()
	if orgset.CheckUserAccess(token, Connection, ResponseWriter) && checkCategoryData(category, ResponseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token, Connection)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		NewConnection := src.Connect_Custom(Organization)
		defer NewConnection.Close()
		CatID, APIerr := addCategory_Request(category, NewConnection)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		Resp := AddCategory_Success{200, "success", CatID}
		Response, _ := json.Marshal(Resp)
		fmt.Fprintf(ResponseWriter, string(Response))
	}
	return true
}

func addCategory_Request(category Category, Connection *sql.DB) (int64, *conf.ApiError){
	Query, err := Connection.Prepare("INSERT INTO categories(name, negative_marks) VALUES(?, ?)")
	if err != nil{
		log.Print(err)
		return 0, conf.ErrDatabaseQueryFailed
	}
	Resp, err := Query.Exec(category.Name, category.NegativeMarks)
	Query.Close()
	if err != nil{
		log.Print(err)
		return 0, conf.ErrDatabaseQueryFailed
	}
	CatID, err := Resp.LastInsertId()
	if err != nil{
		log.Print(err)
		return 0, conf.ErrDatabaseQueryFailed
	}
	APIerr := addCategory_Participants(CatID, Connection)
	if APIerr != nil {
		return 0, APIerr
	}
	APIerr = addCategory_Employees(CatID, Connection)
	if APIerr != nil {
		return 0, APIerr
	}
	return CatID, nil
}

func addCategory_Participants(CatID int64, Connection *sql.DB) *conf.ApiError{
	_, err := Connection.Query("ALTER TABLE participants ADD `"+strconv.FormatInt(CatID, 10)+"` INT NOT NULL DEFAULT '0'")
	if err != nil{
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func addCategory_Employees(CatID int64, Connection *sql.DB) *conf.ApiError{
	_, err := Connection.Query("ALTER TABLE employees ADD `"+strconv.FormatInt(CatID, 10)+"` ENUM('true','false') NOT NULL DEFAULT 'true'")
	if err != nil{
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func checkCategoryData(category Category, w http.ResponseWriter) bool{
	if len(category.Name) > 0 {
		if len(category.NegativeMarks) > 0 {
			if category.NegativeMarks != "false" && category.NegativeMarks != "true" {
				return conf.PrintError(conf.ErrCategoryNegativeMarksIncorrect, w)
			} else {
				return true
			}
		} else {
			return conf.PrintError(conf.ErrCategoryNegativeMarksEmpty, w)
		}
	} else {
		return conf.PrintError(conf.ErrCategoryNameEmpty, w)
	}
}