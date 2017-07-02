/*
	Copyright: "NullTeam", 2016 - 2017
	Author: "De1aY"
	Documentation: https://bitbucket.org/lyceumdevelopers/openqr_golang/wiki/Home
*/
package conf

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type ApiResponse struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Message interface{} `json:"message"`
}

type ErrorMessage struct {
	Ru string `json:"ru"`
	En string `json:"en"`
}

func (response *ApiResponse) toJSON() string {
	resp, _ := json.Marshal(response)
	return string(resp)
}

func (response *ApiResponse) Print(writer http.ResponseWriter) bool {
	fmt.Fprintf(writer, response.toJSON())
	return false
}
// 200
var RequestSuccess = &ApiResponse{200, "success", nil}
// 400
var ErrMethodNotAllowed = &ApiResponse{400, "error", ErrorMessage{"Метод запрещён", "Metod not allowed"}}
var ErrInsufficientRights = &ApiResponse{401, "error", ErrorMessage{"Недостаточно прав", "Insufficient rights"}}
// 500
var ErrDatabaseQueryFailed = &ApiResponse{501, "error", ErrorMessage{"Ошибка соединения с базой данных", "Database connection failed"}}
var ErrConvertStringToInt = &ApiResponse{502, "error", ErrorMessage{"Невозможно перевести строку в число", "Cannot convert string to int"}}
var ErrOpenExcelFile = &ApiResponse{503, "error", ErrorMessage{"Ошибка открытия файла", "Cannot open file"}}
var ErrSaveExcelFile = &ApiResponse{504, "error", ErrorMessage{"Ошибка сохранения файла", "Cannot save file"}}
var ErrCreateExcelFile = &ApiResponse{505, "error", ErrorMessage{"Ошибка создания файла", "Cannot create file"}}
var ErrCreateSheetOnExcelFile = &ApiResponse{506, "error", ErrorMessage{"Ошибка создания листа Excel", "Cannot create excel sheet"}}
// 600
var ErrUserPasswordEmpty = &ApiResponse{601, "error", ErrorMessage{"Пароль отсутствует", "Password is empty"}}
var ErrUserLoginEmpty = &ApiResponse{602, "error", ErrorMessage{"Логин отсутствует", "Login is empty"}}
var ErrUserTokenEmpty = &ApiResponse{603, "error", ErrorMessage{"Токен отсутствует", "Token is empty"}}
var ErrAuthDataIncorrect = &ApiResponse{604, "error", ErrorMessage{"Неправильный логин или пароль", "Login or password is wrong"}}
var ErrUserTokenIncorrect = &ApiResponse{605, "error", ErrorMessage{"Неверный токен", "Token is invalid"}}
var ErrOrgSettingNameEmpty = &ApiResponse{606, "error", ErrorMessage{"Название настройки отсутствует", "Setting name is empty"}}
var ErrOrgSettingValueEmpty = &ApiResponse{606, "error", ErrorMessage{"Значение настройки отсутствует", "Setting value is empty"}}
var ErrOrgSettingNameIncorrect = &ApiResponse{607, "error", ErrorMessage{"Неверное название настройки", "Setting name is incorrect"}}
var ErrCategoryNameEmpty = &ApiResponse{608, "error", ErrorMessage{"Название категории отсутствует", "Category name is empty"}}
var ErrCategoryNegativeMarksEmpty = &ApiResponse{609, "error", ErrorMessage{"Категория 'отрицательные оценки' отсутствует", "Category 'Negative marks' is empty"}}
var ErrCategoryNegativeMarksIncorrect = &ApiResponse{610, "error", ErrorMessage{"Неверное значение категории 'отрицательные оценки'", "Category 'Negative marks' is incorrect"}}
var ErrIDisNotINT = &ApiResponse{611, "error", ErrorMessage{"ID должен быть числом", "ID must be a number"}}
var ErrNameEmpty = &ApiResponse{612, "error", ErrorMessage{"Имя отсутсвует", "Name is empty"}}
var ErrSurnameEmpty = &ApiResponse{613, "error", ErrorMessage{"Фамилия отсутсвует", "Surname is empty"}}
var ErrMiddlenameEmpty = &ApiResponse{614, "error", ErrorMessage{"Отчество отсутсвует", "Middlename is empty"}}
var ErrSexNotINT = &ApiResponse{615, "error", ErrorMessage{"Пол должен быть числом", "Sex must be a number"}}
var ErrTeamNotINT = &ApiResponse{616, "error", ErrorMessage{"Команда должна быть числом", "Team must be a number"}}
var ErrSexIncorrect = &ApiResponse{617, "error", ErrorMessage{"Некорректный пол", "Sex is incorrect"}}
var ErrUserNotFound = &ApiResponse{618, "error", ErrorMessage{"Пользователь не найден", "User not found"}}
var ErrTeamIncorrect = &ApiResponse{619, "error", ErrorMessage{"Некорректная команда", "Team is incorrect"}}
var ErrPostEmpty = &ApiResponse{620, "error", ErrorMessage{"Должность отсутствует", "Post is empty"}}
var ErrCategoryIdIncorrect = &ApiResponse{621, "error", ErrorMessage{"Неверный ID категории", "Category ID is incorrect"}}
var ErrPermissionValueIncorrect = &ApiResponse{622, "error", ErrorMessage{"Разрешение должно быть boolean", "Permission must be a boolean"}}
var ErrCategoryIdNotINT = &ApiResponse{623, "error", ErrorMessage{"ID категории должно быть числом", "Category id must be a number"}}
var ErrReasonIncorrect = &ApiResponse{624, "error", ErrorMessage{"Некорректная причина", "Reason is incorrect"}}
var ErrLoginIncorrect = &ApiResponse{625, "error", ErrorMessage{"Неверный логин", "Partcipant login incorrect"}}
var ErrOrganizationNameEmpty = &ApiResponse{626, "error", ErrorMessage{"Название организации отсутствует", "Oranization name is empty"}}
