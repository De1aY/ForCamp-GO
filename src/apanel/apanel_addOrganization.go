package apanel

import (
	"net/http"
	"forcamp/conf"
	"forcamp/src/orgset"
	"forcamp/src"
	"log"
	"strconv"
	"fmt"
	"github.com/tealeg/xlsx"
)

func checkAddOrganizationData(token string, orgname string) *conf.ApiError {
	if len(orgname) == 0 {
		return conf.ErrOrganizationNameEmpty
	} else {
		return checkPermissions(token)
	}
}

func CreateOrganization(token string, orgname string, responseWriter http.ResponseWriter) bool {
	APIerr := checkAddOrganizationData(token, orgname)
	if APIerr != nil {
		return conf.PrintError(APIerr, responseWriter)
	}
	resp, APIerr := createOrganization_Request(orgname)
	fmt.Fprintf(responseWriter, resp.toJSON())
	return true
}

func createOrganization_Request(orgname string) (createOrganization_Success, *conf.ApiError) {
	login, password, APIerr := createOrganizationAdminAccount(orgname)
	if APIerr != nil {
		return createOrganization_Success{}, APIerr
	}
	APIerr = createOrganizationDB(orgname)
	if APIerr != nil {
		return createOrganization_Success{}, APIerr
	}
	APIerr = createOrganizationDBTables(orgname, login)
	if APIerr != nil {
		return createOrganization_Success{}, APIerr
	}
	APIerr = createOrganizationExcelFiles(orgname)
	if APIerr != nil {
		return createOrganization_Success{}, APIerr
	}
	return createOrganization_Success{200, "success", login, password}, nil
}

func createOrganizationAdminAccount(orgname string) (string, string, *conf.ApiError){
	password, hash := orgset.GeneratePassword()
	query, err := src.Connection.Prepare("INSERT INTO users(password,organization) VALUES(?,?)")
	if err != nil {
		log.Print(err)
		return "", "", conf.ErrDatabaseQueryFailed
	}
	resp, err := query.Exec(hash, orgname)
	user_id, err := resp.LastInsertId()
	if err != nil {
		log.Print(err)
		return "", "", conf.ErrDatabaseQueryFailed
	}
	query.Close()
	query, err = src.Connection.Prepare("UPDATE users SET login=? WHERE id=?")
	if err != nil {
		log.Print(err)
		return "", "", conf.ErrDatabaseQueryFailed
	}
	user_login := "orgadmin_"+strconv.FormatInt(user_id, 10)
	_, err = query.Exec(user_login, user_id)
	if err != nil {
		log.Print(err)
		return "", "", conf.ErrDatabaseQueryFailed
	}
	query.Close()
	return user_login, password, nil
}

func createOrganizationExcelFiles(orgname string) *conf.ApiError {
	APIerr := createOrganizationExcelFile_Employees(orgname)
	if APIerr != nil {
		return APIerr
	}
	APIerr = createOrganizationExcelFile_Participant(orgname)
	if APIerr != nil {
		return APIerr
	}
	return nil
}

func createOrganizationExcelFile_Participant(orgname string) *conf.ApiError {
	excelFilePath := conf.FOLDER_PARTICIPANTS+"/"+orgname+".xlsx"
	xlFile := xlsx.NewFile()
	sheet, err := xlFile.AddSheet("участники")
	if err != nil {
		log.Print(err)
		return conf.ErrCreateSheetOnExcelFile
	}
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "Фамилия"
	cell = row.AddCell()
	cell.Value = "Имя"
	cell = row.AddCell()
	cell.Value = "Отчество"
	cell = row.AddCell()
	cell.Value = "Команда"
	cell = row.AddCell()
	cell.Value = "Логин"
	cell = row.AddCell()
	cell.Value = "Пароль"
	err = xlFile.Save(excelFilePath)
	if err != nil {
		log.Print(err)
		return conf.ErrSaveExcelFile
	}
	return nil
}

func createOrganizationExcelFile_Employees(orgname string) *conf.ApiError {
	excelFilePath := conf.FOLDER_EMPLOYEES+"/"+orgname+".xlsx"
	xlFile := xlsx.NewFile()
	sheet, err := xlFile.AddSheet("сотрудники")
	if err != nil {
		log.Print(err)
		return conf.ErrCreateSheetOnExcelFile
	}
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "Фамилия"
	cell = row.AddCell()
	cell.Value = "Имя"
	cell = row.AddCell()
	cell.Value = "Отчество"
	cell = row.AddCell()
	cell.Value = "Команда"
	cell = row.AddCell()
	cell.Value = "Логин"
	cell = row.AddCell()
	cell.Value = "Пароль"
	err = xlFile.Save(excelFilePath)
	if err != nil {
		log.Print(err)
		return conf.ErrSaveExcelFile
	}
	return nil
}

func createOrganizationDB(orgname string) *conf.ApiError {
	connection := src.Connect_Admin()
	defer connection.Close()
	_, err := connection.Exec("CREATE DATABASE IF NOT EXISTS "+orgname)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func createOrganizationDBTables(orgname string, user_login string) *conf.ApiError{
	src.CustomConnection = src.Connect_Custom(orgname)
	APIerr := createOrganizationDBTable_Users(user_login)
	if APIerr != nil {
		return APIerr
	}
	APIerr = createOrganizationDBTable_Employees(user_login)
	if APIerr != nil {
		return APIerr
	}
	APIerr = createOrganizationDBTable_Participants()
	if APIerr != nil {
		return APIerr
	}
	APIerr = createOrganizationDBTable_Settings()
	if APIerr != nil {
		return APIerr
	}
	APIerr = createOrganizationDBTable_Teams()
	if APIerr != nil {
		return APIerr
	}
	APIerr = createOrganizationDBTable_Categories()
	if APIerr != nil {
		return APIerr
	}
	APIerr = createOrganizationDBTable_MarksChanges()
	if APIerr != nil {
		return APIerr
	}
	APIerr = createOrganizationDBTable_Reasons()
	if APIerr != nil {
		return APIerr
	}
	return nil
}

func createOrganizationDBTable_Users(user_login string) *conf.ApiError{
	query, err := src.CustomConnection.Prepare("CREATE TABLE IF NOT EXISTS users(" +
		"login TINYTEXT," +
		"name TINYTEXT," +
		"surname TINYTEXT," +
		"middlename TINYTEXT," +
		"sex SMALLINT(6)," +
		"team TINYTEXT," +
		"avatar TINYTEXT," +
		"access SMALLINT(6)" +
		")")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec()
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	query.Close()
	query, err = src.CustomConnection.Prepare("INSERT INTO users(login,name,surname,middlename,sex,team,avatar,access) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(user_login, "admin", "admin", "admin", 0, 0, "default.jpg", 2)
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	query.Close()
	return nil
}

func createOrganizationDBTable_Employees(user_login string) *conf.ApiError{
	query, err := src.CustomConnection.Prepare("CREATE TABLE IF NOT EXISTS employees(" +
		"login TINYTEXT," +
		"post TINYTEXT" +
		")")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec()
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	query.Close()
	query, err = src.CustomConnection.Prepare("INSERT INTO employees(login,post) VALUES(?,?)")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(user_login, "администрация")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	query.Close()
	return nil
}

func createOrganizationDBTable_Participants() *conf.ApiError{
	query, err := src.CustomConnection.Prepare("CREATE TABLE IF NOT EXISTS participants(" +
		"login TINYTEXT" +
		")")
	defer query.Close()
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec()
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func createOrganizationDBTable_Settings() *conf.ApiError{
	query, err := src.CustomConnection.Prepare("CREATE TABLE IF NOT EXISTS settings(" +
		"name TINYTEXT," +
		"value TINYTEXT" +
		")")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec()
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	query.Close()
	query, err = src.CustomConnection.Prepare("INSERT INTO settings(name,value) VALUES(?,?)")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec("team", "команда")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	query.Close()
	query, err = src.CustomConnection.Prepare("INSERT INTO settings(name,value) VALUES(?,?)")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec("self_marks", "true")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	query.Close()
	query, err = src.CustomConnection.Prepare("INSERT INTO settings(name,value) VALUES(?,?)")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec("organization", "организация")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	query.Close()
	query, err = src.CustomConnection.Prepare("INSERT INTO settings(name,value) VALUES(?,?)")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec("period", "1 смена")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	query.Close()
	query, err = src.CustomConnection.Prepare("INSERT INTO settings(name,value) VALUES(?,?)")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec("participant", "участник")
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	query.Close()
	return nil
}

func createOrganizationDBTable_Teams() *conf.ApiError{
	query, err := src.CustomConnection.Prepare("CREATE TABLE teams(" +
		"id INT(11) NOT NULL AUTO_INCREMENT," +
		"name TINYTEXT," +
		"PRIMARY KEY (id)" +
		")")
	defer query.Close()
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec()
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func createOrganizationDBTable_Categories() *conf.ApiError{
	query, err := src.CustomConnection.Prepare("CREATE TABLE IF NOT EXISTS categories(" +
		"id INT(11) NOT NULL AUTO_INCREMENT," +
		"name TINYTEXT," +
		"negative_marks ENUM('true','false') DEFAULT 'true'," +
		"PRIMARY KEY (id)" +
		")")
	defer query.Close()
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec()
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func createOrganizationDBTable_Reasons() *conf.ApiError{
	query, err := src.CustomConnection.Prepare("CREATE TABLE IF NOT EXISTS reasons(" +
		"id INT(11) NOT NULL AUTO_INCREMENT," +
		"cat_id INT(11)," +
		"text TINYTEXT," +
		"modification INT(11)," +
		"PRIMARY KEY (id)" +
		")")
	defer query.Close()
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec()
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

func createOrganizationDBTable_MarksChanges() *conf.ApiError{
	query, err := src.CustomConnection.Prepare("CREATE TABLE IF NOT EXISTS marks_changes(" +
		"id INT(11) NOT NULL AUTO_INCREMENT," +
		"reason_id INT(11)," +
		"employee_login TINYTEXT," +
		"participant_login TINYTEXT," +
		"time TIMESTAMP DEFAULT CURRENT_TIMESTAMP," +
		"PRIMARY KEY (id)" +
		")")
	defer query.Close()
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec()
	if err != nil {
		log.Print(err)
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}

