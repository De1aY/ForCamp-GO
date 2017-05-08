package orgset_get

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/handlers"
	"forcamp/src/api/orgset"
	"io/ioutil"
	"time"
	"bytes"
)


func GetParticipantsExcelHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		Token := handlers.GetToken(r)
		if orgset.CheckUserAccess(Token, w){
			Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(Token);
			if APIerr != nil {
				src.SetHeaders_API(w)
				conf.PrintError(APIerr, w)
			} else {
				file, err := ioutil.ReadFile(conf.FOLDER_PARTICIPANTS+"/"+Organization+".xlsx")
				if err != nil {
					src.SetHeaders_API(w)
					conf.PrintError(conf.ErrOpenExcelFile, w)
				} else {
					src.SetHeaders_API_Download(w, "участники.xlsx", r)
					http.ServeContent(w, r, "участники.xlsx", time.Now(), bytes.NewReader(file))
				}
			}
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleGetParticipantsExcel(router *mux.Router)  {
	router.HandleFunc("/orgset.participants.password.get", GetParticipantsExcelHandler)
}
