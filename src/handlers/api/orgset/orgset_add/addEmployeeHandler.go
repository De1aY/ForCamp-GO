package orgset_add

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strings"
	"strconv"
	"log"
	"forcamp/src/api/orgset/employees"
)

func getAddEmployeePostValues(r *http.Request) (employees.Employee, string, *conf.ApiError){
	Token := r.PostFormValue("token")
	Name := strings.TrimSpace(strings.ToLower(r.PostFormValue("name")))
	Surname := strings.TrimSpace(strings.ToLower(r.PostFormValue("surname")))
	Middlename := strings.TrimSpace(strings.ToLower(r.PostFormValue("middlename")))
	Post := strings.TrimSpace(strings.ToLower(r.PostFormValue("post")))
	Sex, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("sex")), 10, 64)
	if err != nil {
		log.Print(err)
		return employees.Employee{}, "", conf.ErrEmployeeSexNotINT
	}
	Team, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("team")), 10, 64)
	if err != nil {
		log.Print(err)
		return employees.Employee{}, "", conf.ErrEmployeeTeamNotINT
	}
	return employees.Employee{"", Name, Surname, Middlename, int(Sex), Team, Post, nil}, Token, nil
}

func AddEmployeeHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		employee, token, APIerr := getAddEmployeePostValues(r)
		if APIerr != nil {
			conf.PrintError(APIerr, w)
		} else {
			employees.AddEmployee(token, employee, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleAddEmployee(router *mux.Router)  {
	router.HandleFunc("/orgset.employee.add", AddEmployeeHandler)
}
