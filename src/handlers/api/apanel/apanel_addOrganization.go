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
	organizationName := strings.TrimSpace(strings.ToLower(r.PostFormValue("orgname")))
	return organizationName, Token
}

func addOrganizationHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_POST(w)
	if r.Method == http.MethodPost {
		organizationName, token := getAddOrganizationPostValues(r)
		apanel.CreateOrganization(token, organizationName, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleAddOrganization(router *mux.Router)  {
	router.HandleFunc("/apanel.organization.add", addOrganizationHandler)
}
