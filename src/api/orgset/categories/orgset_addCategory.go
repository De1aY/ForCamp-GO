/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package categories

import (
	"net/http"
	"wplay/conf"
	"wplay/src"
	"strconv"
	"wplay/src/api/orgset"
)

type addCategory_Success struct {
	ID int64 `json:"id"`
}


func AddCategory(token string, category Category, responseWriter http.ResponseWriter) bool{
	if orgset.IsUserAdmin(token, responseWriter) && checkCategoryData(category, responseWriter){
		Organization, _, apiErr := orgset.GetUserOrganizationAndIdByToken(token)
		if apiErr != nil{
			return apiErr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		CatID, apiErr := addCategory_Request(category)
		if apiErr != nil{
			return apiErr.Print(responseWriter)
		}
		resp := conf.ApiResponse{200, "success", addCategory_Success{CatID}}
		resp.Print(responseWriter)
	}
	return true
}

func addCategory_Request(category Category) (int64, *conf.ApiResponse){
	query, err := src.CustomConnection.Prepare("INSERT INTO categories(name, negative_marks) VALUES(?, ?)")
	if err != nil{
		return 0, conf.ErrDatabaseQueryFailed
	}
	Resp, err := query.Exec(category.Name, category.NegativeMarks)
	query.Close()
	if err != nil{
		return 0, conf.ErrDatabaseQueryFailed
	}
	CatID, err := Resp.LastInsertId()
	if err != nil{
		return 0, conf.ErrDatabaseQueryFailed
	}
	apiErr := addCategory_Participants(CatID)
	if apiErr != nil {
		return 0, apiErr
	}
	apiErr = addCategory_Employees(CatID)
	if apiErr != nil {
		return 0, apiErr
	}
	return CatID, nil
}

func addCategory_Participants(CatID int64) *conf.ApiResponse{
	_, err := src.CustomConnection.Query("ALTER TABLE participants ADD `"+strconv.FormatInt(CatID, 10)+"` INT NOT NULL DEFAULT '0'")
	if err != nil{
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func addCategory_Employees(CatID int64) *conf.ApiResponse{
	_, err := src.CustomConnection.Query("ALTER TABLE employees ADD `"+strconv.FormatInt(CatID, 10)+"` ENUM('true','false') NOT NULL DEFAULT 'true'")
	if err != nil{
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
