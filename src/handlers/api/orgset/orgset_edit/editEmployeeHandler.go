package orgset_edit

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

func getEditEmployeePostValues(r *http.Request) (employees.Employee, string, *conf.ApiResponse){
	Token := strings.TrimSpace(r.PostFormValue("token"))
	Login := strings.TrimSpace(strings.ToLower(r.PostFormValue("login")))
	Name := strings.TrimSpace(strings.ToLower(r.PostFormValue("name")))
	Surname := strings.TrimSpace(strings.ToLower(r.PostFormValue("surname")))
	Middlename := strings.TrimSpace(strings.ToLower(r.PostFormValue("middlename")))
	Post := strings.TrimSpace(strings.ToLower(r.PostFormValue("post")))
	Sex, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("sex")), 10, 64)
	if err != nil {
		log.Print(err)
		return employees.Employee{}, "", conf.ErrSexNotINT
	}
	Team, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("team")), 10, 64)
	if err != nil {
		log.Print(err)
		return employees.Employee{}, "", conf.ErrTeamNotINT
	}
	return employees.Employee{Login, Name, Surname, Middlename, int(Sex), Team, Post,nil}, Token, nil
}

func EditEmployeeHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		employee, token, APIerr := getEditEmployeePostValues(r)
		if APIerr != nil {
			APIerr.Print(w)
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
