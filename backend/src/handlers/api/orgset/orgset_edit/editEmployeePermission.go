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

func getEditEmployeePermissionPostValues(r *http.Request) (int64, int64, string, string, *conf.ApiResponse){
	token := strings.TrimSpace(r.PostFormValue("token"))
	employee_id, err := strconv.ParseInt(strings.TrimSpace(
		strings.ToLower(r.PostFormValue("employee_id"))), 10, 64)
	if err != nil {
		return 0, 0, "", "", conf.ErrIdIsNotINT
	}
	value := strings.TrimSpace(strings.ToLower(r.PostFormValue("value")))
	category_id, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("category_id")), 10, 64)
	if err != nil {
		return 0, 0, "", "", conf.ErrCategoryIdNotINT
	}
	return employee_id, category_id, value, token, nil
}

func EditEmployeePermissionHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_POST(w)
	if r.Method == http.MethodPost {
		employee_id, category_id, value, token, apiErr := getEditEmployeePermissionPostValues(r)
		if apiErr != nil {
			apiErr.Print(w)
		} else {
			employees.EditEmployeePermission(token, employee_id, category_id, value, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleEditEmployeePermission(router *mux.Router)  {
	router.HandleFunc("/orgset.employee.permission.edit", EditEmployeePermissionHandler)
}

