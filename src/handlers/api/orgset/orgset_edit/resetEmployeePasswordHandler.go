/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package orgset_edit

import (
	"net/http"
	"github.com/gorilla/mux"
	"nullteam.info/wplay/demo/conf"
	"nullteam.info/wplay/demo/src"
	"strings"
	"nullteam.info/wplay/demo/src/api/orgset/employees"
	"strconv"
)

func getResetEmployeePasswordPostValues(r *http.Request) (string, int64, *conf.ApiResponse){
	token := strings.TrimSpace(r.PostFormValue("token"))
	employee_id, err := strconv.ParseInt(strings.TrimSpace(
		strings.ToLower(r.PostFormValue("employee_id"))), 10, 64)
	if err != nil {
		return "", 0, conf.ErrIdIsNotINT
	}
	return token, employee_id, nil
}

func ResetEmployeePasswordHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_POST(w)
	if r.Method == http.MethodPost {
		token, employee_id, apiErr := getResetEmployeePasswordPostValues(r); if apiErr != nil {
			apiErr.Print(w)
		} else {
			employees.ResetEmployeePassword(token, employee_id, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleResetEmployeePassword(router *mux.Router)  {
	router.HandleFunc("/orgset.employee.password.reset", ResetEmployeePasswordHandler)
}
