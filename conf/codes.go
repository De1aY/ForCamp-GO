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
// 600
var ErrUserPasswordEmpty = &ApiError{601, "ERROR", "Password is empty"}
var ErrUserLoginEmpty = &ApiError{602, "ERROR", "Login is empty"}
var ErrUserTokenEmpty = &ApiError{603, "ERROR", "Token is empty"}
var ErrAuthDataIncorrect = &ApiError{604, "ERROR", "Login or password is wrong"}
var ErrUserTokenIncorrect = &ApiError{605, "ERROR", "Token is invalid"}
var ErrOrgSettingNameEmpty = &ApiError{606, "ERROR", "Setting name is empty"}
var ErrOrgSettingValueEmpty = &ApiError{606, "ERROR", "Setting value is empty"}
var ErrOrgSettingNameIncorrect = &ApiError{607, "ERROR", "Setting name is incorrect"}
