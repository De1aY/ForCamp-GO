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
)

type addCategory_Success struct {
	Code int `json:"code"`
	Status string `json:"status"`
	ID int64 `json:"id"`
}

func (success *addCategory_Success) toJSON() string {
	resp, _ := json.Marshal(success)
	return string(resp)
}


func AddCategory(token string, category Category, ResponseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token, ResponseWriter) && checkCategoryData(category, ResponseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		CatID, APIerr := addCategory_Request(category)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		resp := addCategory_Success{200, "success", CatID}
		fmt.Fprintf(ResponseWriter, resp.toJSON())
	}
	return true
}

func addCategory_Request(category Category) (int64, *conf.ApiError){
	Query, err := src.CustomConnection.Prepare("INSERT INTO categories(name, negative_marks) VALUES(?, ?)")
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
	APIerr := addCategory_Participants(CatID)
	if APIerr != nil {
		return 0, APIerr
	}
	APIerr = addCategory_Employees(CatID)
	if APIerr != nil {
		return 0, APIerr
	}
	return CatID, nil
}

func addCategory_Participants(CatID int64) *conf.ApiError{
	_, err := src.CustomConnection.Query("ALTER TABLE participants ADD `"+strconv.FormatInt(CatID, 10)+"` INT NOT NULL DEFAULT '0'")
	if err != nil{
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func addCategory_Employees(CatID int64) *conf.ApiError{
	_, err := src.CustomConnection.Query("ALTER TABLE employees ADD `"+strconv.FormatInt(CatID, 10)+"` ENUM('true','false') NOT NULL DEFAULT 'true'")
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