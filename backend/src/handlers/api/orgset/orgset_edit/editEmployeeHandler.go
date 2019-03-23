package orgset_edit

import (
	"net/http"
	"github.com/gorilla/mux"
	"wplay/conf"
	"wplay/src"
	"strings"
	"strconv"
	"wplay/src/api/orgset/employees"
)

func getEditEmployeePostValues(r *http.Request) (employees.Employee, string, *conf.ApiResponse){
	token := strings.TrimSpace(r.PostFormValue("token"))
	employee_id, err := strconv.ParseInt(strings.TrimSpace(
		strings.ToLower(r.PostFormValue("employee_id"))), 10, 64)
	if err != nil {
		return employees.Employee{}, "", conf.ErrIdIsNotINT
	}
	name := strings.TrimSpace(strings.ToLower(r.PostFormValue("name")))
	surname := strings.TrimSpace(strings.ToLower(r.PostFormValue("surname")))
	middlename := strings.TrimSpace(strings.ToLower(r.PostFormValue("middlename")))
	post := strings.TrimSpace(strings.ToLower(r.PostFormValue("post")))
	sex, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("sex")), 10, 64); if err != nil {
		return employees.Employee{}, "", conf.ErrSexNotINT
	}
	team, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("team")), 10, 64); if err != nil {
		return employees.Employee{}, "", conf.ErrTeamNotINT
	}
	return employees.Employee{employee_id, name, surname, middlename, int(sex), team, post, nil}, token, nil
}

func EditEmployeeHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_POST(w)
	if r.Method == http.MethodPost {
		employee, token, apiErr := getEditEmployeePostValues(r)
		if apiErr != nil {
			apiErr.Print(w)
		} else {
			employees.EditEmployee(token, employee, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleEditEmployee(router *mux.Router)  {
	router.HandleFunc("/orgset.employee.edit", EditEmployeeHandler)
}
