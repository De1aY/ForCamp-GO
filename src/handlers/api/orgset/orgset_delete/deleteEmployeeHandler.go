package orgset_delete

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strings"
	"forcamp/src/api/orgset/employees"
)

func getDeleteEmployeePostValues(r *http.Request) (string, string){
	Token := strings.TrimSpace(r.PostFormValue("token"))
	Login := strings.ToLower(strings.TrimSpace(r.PostFormValue("login")))
	return Login, Token
}

func DeleteEmployeeHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		login, token := getDeleteEmployeePostValues(r)
		employees.DeleteEmployee(token, login, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleDeleteEmployee(router *mux.Router)  {
	router.HandleFunc("/orgset.employee.delete", DeleteEmployeeHandler)
}
