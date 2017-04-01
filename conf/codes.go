package conf

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type ApiError struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Message string `json:"message"`
}

type LoginSuccess struct {
	Code int `json:"code"`
	Token string `json:"token"`
	Status string `json:"status"`
}

type Success struct {
	Code int `json:"code"`
	Status string `json:"status"`
}

func (err *ApiError) Error() string{
	return err.Message
}

func NewApiError(err error) *ApiError{
	return &ApiError{0, "ERROR", err.Error()}
}

func PrintError(err *ApiError, w http.ResponseWriter) bool{
	Response, _ := json.Marshal(err)
	fmt.Fprintf(w, string(Response))
	return false
}

func PrintSuccess(success *Success, w http.ResponseWriter) bool{
	Response, _ := json.Marshal(success)
	fmt.Fprintf(w, string(Response))
	return true
}

// 200
var RequestSuccess = &Success{200, "success"}
// 400
var ErrMethodNotAllowed = &ApiError{400, "ERROR", "Method not allowed"}
var ErrInsufficientRights = &ApiError{401, "ERROR", "Insufficient rights"}
// 500
var ErrDatabaseQueryFailed = &ApiError{501, "ERROR", "Query failed"}
var ErrConvertStringToInt = &ApiError{502, "ERROR", "Cannot convert string to int"}
var ErrOpenExcelFile = &ApiError{503, "ERROR", "Cannot open participants file"}
var ErrSaveExcelFile = &ApiError{504, "ERROR", "Cannot save participants file"}
// 600
var ErrUserPasswordEmpty = &ApiError{601, "ERROR", "Password is empty"}
var ErrUserLoginEmpty = &ApiError{602, "ERROR", "Login is empty"}
var ErrUserTokenEmpty = &ApiError{603, "ERROR", "Token is empty"}
var ErrAuthDataIncorrect = &ApiError{604, "ERROR", "Login or password is wrong"}
var ErrUserTokenIncorrect = &ApiError{605, "ERROR", "Token is invalid"}
var ErrOrgSettingNameEmpty = &ApiError{606, "ERROR", "Setting name is empty"}
var ErrOrgSettingValueEmpty = &ApiError{606, "ERROR", "Setting value is empty"}
var ErrOrgSettingNameIncorrect = &ApiError{607, "ERROR", "Setting name is incorrect"}
var ErrCategoryNameEmpty = &ApiError{608, "ERROR", "Category name is empty"}
var ErrCategoryNegativeMarksEmpty = &ApiError{609, "ERROR", "Category 'Negative marks' is empty"}
var ErrCategoryNegativeMarksIncorrect = &ApiError{610, "ERROR", "Category 'Negative marks' is incorrect"}
var ErrIDisNotINT = &ApiError{611, "ERROR", "ID must be a number"}
var ErrParticipantNameEmpty = &ApiError{612, "ERROR", "Name is empty"}
var ErrParticipantSurnameEmpty = &ApiError{613, "ERROR", "Surname is empty"}
var ErrParticipantMiddlenameEmpty = &ApiError{614, "ERROR", "Middlename is empty"}
var ErrParticipantSexNotINT = &ApiError{615, "ERROR", "Sex must be a number"}
var ErrParticipantTeamNotINT = &ApiError{616, "ERROR", "Team must be a number"}
var ErrParticipantSexIncorrect = &ApiError{617, "ERROR", "Sex is incorrect"}
var ErrParticipantTeamIncorrect = &ApiError{618, "ERROR", "Team is incorrect"}
var ErrUserNotFound = &ApiError{619, "ERROR", "User not found"}
var ErrEmployeeNameEmpty = &ApiError{620, "ERROR", "Name is empty"}
var ErrEmployeeSurnameEmpty = &ApiError{621, "ERROR", "Surname is empty"}
var ErrEmployeeMiddlenameEmpty = &ApiError{622, "ERROR", "Middlename is empty"}
var ErrEmployeeSexNotINT = &ApiError{623, "ERROR", "Sex must be a number"}
var ErrEmployeeTeamNotINT = &ApiError{624, "ERROR", "Team must be a number"}
var ErrEmployeeSexIncorrect = &ApiError{625, "ERROR", "Sex is incorrect"}
var ErrEmployeeTeamIncorrect = &ApiError{626, "ERROR", "Team is incorrect"}
var ErrEmployeePostEmpty = &ApiError{627, "ERROR", "Post is empty"}
