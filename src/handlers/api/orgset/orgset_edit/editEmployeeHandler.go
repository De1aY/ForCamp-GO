package orgset_edit

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strings"
	"strconv"
	"log"
	"forcamp/src/orgset/employees"
)

func getEditEmployeePostValues(r *http.Request) (employees.Employee, string, *conf.ApiError){
	Token := strings.TrimSpace(r.PostFormValue("token"))
	Login := strings.TrimSpace(strings.ToLower(r.PostFormValue("login")))
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
	return employees.Employee{Login, Name, Surname, Middlename, int(Sex), Team, Post,nil}, Token, nil
}

func EditEmployeeHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		employee, token, APIerr := getEditEmployeePostValues(r)
		if APIerr != nil {
			conf.PrintError(APIerr, w)
		} else {
			employees.EditEmployee(token, employee, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleEditEmployee(router *mux.Router)  {
	router.HandleFunc("/orgset.employee.edit", EditEmployeeHandler)
}
