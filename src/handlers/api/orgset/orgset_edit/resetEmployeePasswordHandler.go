package orgset_edit

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strings"
	"forcamp/src/api/orgset/employees"
)

func getResetEmployeePasswordPostValues(r *http.Request) (string, string){
	Token := strings.TrimSpace(r.PostFormValue("token"))
	Login := strings.TrimSpace(strings.ToLower(r.PostFormValue("login")))
	return Login, Token
}

func ResetEmployeePasswordHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		login, token := getResetEmployeePasswordPostValues(r)
		employees.ResetEmployeePassword(token, login, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleResetEmployeePassword(router *mux.Router)  {
	router.HandleFunc("/orgset.employee.password.reset", ResetEmployeePasswordHandler)
}
