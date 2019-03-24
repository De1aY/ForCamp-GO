package orgset_delete

import (
	"net/http"
	"github.com/gorilla/mux"
	"wplay/conf"
	"wplay/src"
	"strings"
	"wplay/src/api/orgset/employees"
	"strconv"
)

func getDeleteEmployeePostValues(r *http.Request) (string, int64, *conf.ApiResponse){
	token := strings.TrimSpace(r.PostFormValue("token"))
	employee_id, err := strconv.ParseInt(strings.ToLower(
		strings.TrimSpace(r.PostFormValue("employee_id"))), 10, 64)
	if err != nil {
		return "", 0, conf.ErrIdIsNotINT
	}
	return token, employee_id, nil
}

func DeleteEmployeeHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_POST(w)
	if r.Method == http.MethodPost {
		token, employee_id, apiErr := getDeleteEmployeePostValues(r); if apiErr != nil {
			apiErr.Print(w)
		} else {
			employees.DeleteEmployee(token, employee_id, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleDeleteEmployee(router *mux.Router)  {
	router.HandleFunc("/orgset.employee.delete", DeleteEmployeeHandler)
}
