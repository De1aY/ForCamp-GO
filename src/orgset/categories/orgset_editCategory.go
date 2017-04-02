package categories

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src"
	"strconv"
	"log"
	"forcamp/src/orgset"
	"database/sql"
)

func EditCategory(token string, category Category, ResponseWriter http.ResponseWriter) bool{
	Connection := src.Connect()
	defer Connection.Close()
	if orgset.CheckUserAccess(token, Connection, ResponseWriter) && checkCategoryData(category, ResponseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token, Connection)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		NewConnection := src.Connect_Custom(Organization)
		defer NewConnection.Close()
		APIerr = editCategory_Request(category, NewConnection)
		if APIerr != nil{
			return conf.PrintError(APIerr, ResponseWriter)
		}
		conf.PrintSuccess(conf.RequestSuccess, ResponseWriter)
	}
	return true
}

func editCategory_Request(category Category, Connection *sql.DB) *conf.ApiError{
	Query, err := Connection.Prepare("UPDATE categories SET name=?, negative_marks=? WHERE id=?")
	if err != nil{
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(category.Name, category.NegativeMarks, strconv.FormatInt(category.ID, 10))
	Query.Close()
	if err != nil{
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return  nil
}