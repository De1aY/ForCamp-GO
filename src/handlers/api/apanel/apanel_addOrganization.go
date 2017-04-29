package apanel

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strings"
	"forcamp/src/apanel"
)

func getAddOrganizationPostValues(r *http.Request) (string, string){
	Token := r.PostFormValue("token")
	Orgname := strings.ToLower(r.PostFormValue("orgname"))
	return Orgname, Token
}

func addOrganizationHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		orgname, token := getAddOrganizationPostValues(r)
		apanel.CreateOrganization(token, orgname, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleAddOrganization(router *mux.Router)  {
	router.HandleFunc("/apanel.organization.add", addOrganizationHandler)
}
