/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package apanel

import (
	"wplay/conf"
	"wplay/src"
	"wplay/src/api/orgset"
	"net/http"
	"strconv"

	"github.com/tealeg/xlsx"
)

func CreateOrganization(token string, orgname string, name string,
	surname string, middlename string, responseWriter http.ResponseWriter) bool {
	isAdmin := isUserAdmin(token)
	if isAdmin != nil {
		return isAdmin.Print(responseWriter)
	}
	rawResp, apiErr := createOrganization(orgname, name, surname, middlename)
	if apiErr != nil {
		return apiErr.Print(responseWriter)
	}
	resp := &conf.ApiResponse{200, "success", rawResp}
	return resp.Print(responseWriter)
}

func createOrganization(orgname string, name string,
	surname string, middlename string) (createOrganization_Success, *conf.ApiResponse) {
	user_id, login, password, APIerr := createOrganizationAdminAccount(orgname)
	if APIerr != nil {
		return createOrganization_Success{}, APIerr
	}
	APIerr = createOrganizationDB(orgname, name, surname, middlename, user_id)
	if APIerr != nil {
		return createOrganization_Success{}, APIerr
	}
	APIerr = createOrganizationExcelFiles(orgname)
	if APIerr != nil {
		return createOrganization_Success{}, APIerr
	}
	return createOrganization_Success{login, password}, nil
}

func createOrganizationAdminAccount(orgname string) (int64, string, string, *conf.ApiResponse) {
	password, hash := orgset.GeneratePassword()
	query, err := src.Connection.Prepare("INSERT INTO users(password,organization) VALUES(?,?)")
	if err != nil {
		return 0, "", "", conf.ErrDatabaseQueryFailed
	}
	resp, err := query.Exec(hash, orgname)
	user_id, err := resp.LastInsertId()
	if err != nil {
		return 0, "", "", conf.ErrDatabaseQueryFailed
	}
	query.Close()
	query, err = src.Connection.Prepare("UPDATE users SET login=? WHERE id=?")
	if err != nil {
		return 0, "", "", conf.ErrDatabaseQueryFailed
	}
	user_login := "orgadmin_" + strconv.FormatInt(user_id, 10)
	_, err = query.Exec(user_login, user_id)
	if err != nil {
		return 0, "", "", conf.ErrDatabaseQueryFailed
	}
	query.Close()
	return user_id, user_login, password, nil
}

func createOrganizationExcelFiles(orgname string) *conf.ApiResponse {
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

func createOrganizationExcelFile_Participant(orgname string) *conf.ApiResponse {
	excelFilePath := conf.FOLDER_PARTICIPANTS + "/" + orgname + ".xlsx"
	xlFile := xlsx.NewFile()
	sheet, err := xlFile.AddSheet("участники")
	if err != nil {
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
		return conf.ErrSaveExcelFile
	}
	return nil
}

func createOrganizationExcelFile_Employees(orgname string) *conf.ApiResponse {
	excelFilePath := conf.FOLDER_EMPLOYEES + "/" + orgname + ".xlsx"
	xlFile := xlsx.NewFile()
	sheet, err := xlFile.AddSheet("сотрудники")
	if err != nil {
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
		return conf.ErrSaveExcelFile
	}
	return nil
}

func createOrganizationDB(orgname string, name string,
	surname string, middlename string, user_id int64) *conf.ApiResponse {
	connection := src.Connect_Admin()
	defer connection.Close()
	_, err := connection.Exec("CREATE DATABASE IF NOT EXISTS " + orgname)
	_, err = connection.Exec("CREATE TABLE " + orgname + ".users LIKE starter.users")
	_, err = connection.Exec("CREATE TABLE " + orgname + ".events LIKE starter.events")
	_, err = connection.Exec("CREATE TABLE " + orgname + ".employees LIKE starter.employees")
	_, err = connection.Exec("CREATE TABLE " + orgname + ".participants LIKE starter.participants")
	_, err = connection.Exec("CREATE TABLE " + orgname + ".reasons LIKE starter.reasons")
	_, err = connection.Exec("CREATE TABLE " + orgname + ".emotional_marks LIKE starter.emotional_marks")
	_, err = connection.Exec("CREATE TABLE " + orgname + ".teams LIKE starter.teams")
	_, err = connection.Exec("CREATE TABLE " + orgname + ".settings LIKE starter.settings")
	_, err = connection.Exec("INSERT INTO "  + orgname + ".settings SELECT * FROM starter.settings")
	_, err = connection.Exec("CREATE TABLE " + orgname + ".marks_changes LIKE starter.marks_changes")
	_, err = connection.Exec("CREATE TABLE " + orgname + ".categories LIKE starter.categories")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	src.CustomConnection = src.Connect_Custom(orgname)
	request, err := src.CustomConnection.Prepare("INSERT INTO users(id,name,surname,middlename,access,sex,avatar) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = request.Exec(&user_id, &name, &surname, &middlename, 2, 0, "default.jpg")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	request, err = src.CustomConnection.Prepare("INSERT INTO employees(id,post) VALUES(?,?)")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = request.Exec(&user_id, "администрация")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return nil
}
