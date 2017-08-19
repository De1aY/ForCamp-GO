package orgset_add

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strings"
	"strconv"
	"forcamp/src/api/orgset/employees"
)

func getAddEmployeePostValues(r *http.Request) (employees.Employee, string, *conf.ApiResponse){
	Token := r.PostFormValue("token")
	Name := strings.TrimSpace(strings.ToLower(r.PostFormValue("name")))
	Surname := strings.TrimSpace(strings.ToLower(r.PostFormValue("surname")))
	Middlename := strings.TrimSpace(strings.ToLower(r.PostFormValue("middlename")))
	Post := strings.TrimSpace(strings.ToLower(r.PostFormValue("post")))
	Sex, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("sex")), 10, 64)
	if err != nil {
		return employees.Employee{}, "", conf.ErrSexNotINT
	}
	Team, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("team")), 10, 64)
	if err != nil {
		return employees.Employee{}, "", conf.ErrTeamNotINT
	}
	return employees.Employee{0, Name, Surname, Middlename, int(Sex), Team, Post, nil}, Token, nil
}

func AddEmployeeHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_POST(w)
	if r.Method == http.MethodPost {
		employee, token, APIerr := getAddEmployeePostValues(r)
		if APIerr != nil {
			APIerr.Print(w)
		} else {
			employees.AddEmployee(token, employee, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleAddEmployee(router *mux.Router)  {
	router.HandleFunc("/orgset.employee.add", AddEmployeeHandler)
}
