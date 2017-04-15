package orgset_get

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/handlers"
	"forcamp/src/orgset"
	"io/ioutil"
	"time"
	"bytes"
)


func GetEmployeesExcelHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		Token := handlers.GetToken(r)
		if orgset.CheckUserAccess(Token, w){
			Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(Token);
			if APIerr != nil {
				src.SetHeaders_API(w)
				conf.PrintError(APIerr, w)
			} else {
				file, err := ioutil.ReadFile(conf.FOLDER_EMPLOYEES+"/"+Organization+".xlsx")
				if err != nil {
					src.SetHeaders_API(w)
					conf.PrintError(conf.ErrOpenExcelFile, w)
				} else {
					src.SetHeaders_API_Download(w, "сотрудники.xlsx", r)
					http.ServeContent(w, r, "сотрудники.xlsx", time.Now(), bytes.NewReader(file))
				}
			}
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleGetEmployeesExcel(router *mux.Router)  {
	router.HandleFunc("/orgset.employees.password.get", GetEmployeesExcelHandler)
}
