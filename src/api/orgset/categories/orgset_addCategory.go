package categories

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src"
	"strconv"
	"log"
	"forcamp/src/api/orgset"
)

type addCategory_Success struct {
	ID int64 `json:"id"`
}


func AddCategory(token string, category Category, responseWriter http.ResponseWriter) bool{
	if orgset.CheckUserAccess(token, responseWriter) && checkCategoryData(category, responseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil{
			return APIerr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		CatID, APIerr := addCategory_Request(category)
		if APIerr != nil{
			return APIerr.Print(responseWriter)
		}
		resp := conf.ApiResponse{200, "success", addCategory_Success{CatID}}
		resp.Print(responseWriter)
	}
	return true
}

func addCategory_Request(category Category) (int64, *conf.ApiResponse){
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

func addCategory_Participants(CatID int64) *conf.ApiResponse{
	_, err := src.CustomConnection.Query("ALTER TABLE participants ADD `"+strconv.FormatInt(CatID, 10)+"` INT NOT NULL DEFAULT '0'")
	if err != nil{
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func addCategory_Employees(CatID int64) *conf.ApiResponse{
	_, err := src.CustomConnection.Query("ALTER TABLE employees ADD `"+strconv.FormatInt(CatID, 10)+"` ENUM('true','false') NOT NULL DEFAULT 'true'")
	if err != nil{
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func checkCategoryData(category Category, responseWriter http.ResponseWriter) bool{
	if len(category.Name) > 0 {
		if len(category.NegativeMarks) > 0 {
			if category.NegativeMarks != "false" && category.NegativeMarks != "true" {
				return conf.ErrCategoryNegativeMarksIncorrect.Print(responseWriter)
			} else {
				return true
			}
		} else {
			return conf.ErrCategoryNegativeMarksEmpty.Print(responseWriter)
		}
	} else {
		return conf.ErrCategoryNameEmpty.Print(responseWriter)
	}
}