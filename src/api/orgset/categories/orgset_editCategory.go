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

func EditCategory(token string, category Category, responseWriter http.ResponseWriter) bool{
	if orgset.IsUserAdmin(token, responseWriter) && checkCategoryData(category, responseWriter){
		organizationName, _, apiErr := orgset.GetUserOrganizationAndIdByToken(token)
		if apiErr != nil{
			return apiErr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(organizationName)
		apiErr = editCategory_Request(category)
		if apiErr != nil{
			return apiErr.Print(responseWriter)
		}
		conf.RequestSuccess.Print(responseWriter)
	}
	return true
}

func editCategory_Request(category Category) *conf.ApiResponse{
	query, err := src.CustomConnection.Prepare("UPDATE categories SET name=?, negative_marks=? WHERE id=?")
	if err != nil{
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(category.Name, category.NegativeMarks, strconv.FormatInt(category.ID, 10))
	query.Close()
	if err != nil{
		return conf.ErrDatabaseQueryFailed
	}
	return  nil
}
