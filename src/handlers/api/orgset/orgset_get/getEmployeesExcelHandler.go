/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package orgset_get

import (
	"net/http"
	"github.com/gorilla/mux"
	"wplay/conf"
	"wplay/src"
	"wplay/src/handlers"
	"wplay/src/api/orgset"
	"io/ioutil"
	"time"
	"bytes"
)


func GetEmployeesExcelHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		token := handlers.GetToken(r)
		if orgset.IsUserAdmin(token, w){
			organizationName, _, apiErr := orgset.GetUserOrganizationAndIdByToken(token);
			if apiErr != nil {
				src.SetHeaders_API_GET(w)
				apiErr.Print(w)
			} else {
				file, err := ioutil.ReadFile(conf.FOLDER_EMPLOYEES+"/"+ organizationName +".xlsx")
				if err != nil {
					src.SetHeaders_API_GET(w)
					conf.ErrOpenExcelFile.Print(w)
				} else {
					src.SetHeaders_API_Download(w, "сотрудники.xlsx", r)
					http.ServeContent(w, r, "сотрудники.xlsx", time.Now(), bytes.NewReader(file))
				}
			}
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleGetEmployeesExcel(router *mux.Router)  {
	router.HandleFunc("/orgset.employees.password.get", GetEmployeesExcelHandler)
}
