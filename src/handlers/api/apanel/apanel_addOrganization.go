package apanel

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strings"
	"forcamp/src/api/apanel"
)

func getAddOrganizationPostValues(r *http.Request) (string, string){
	Token := strings.TrimSpace(r.PostFormValue("token"))
	Orgname := strings.TrimSpace(strings.ToLower(r.PostFormValue("orgname")))
	return Orgname, Token
}

func addOrganizationHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		orgname, token := getAddOrganizationPostValues(r)
		apanel.CreateOrganization(token, orgname, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleAddOrganization(router *mux.Router)  {
	router.HandleFunc("/apanel.organization.add", addOrganizationHandler)
}
