package users_edit

import (
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/api/users"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func getChangeUserPasswordPostValues(r *http.Request) (string, string, string) {
	token := strings.TrimSpace(r.PostFormValue("token"))
	oldPassword := strings.TrimSpace(r.PostFormValue("password_current"))
	newPassword := strings.TrimSpace(r.PostFormValue("password_new"))
	return token, oldPassword, newPassword
}

func changeUserPassword(w http.ResponseWriter, r *http.Request) {
	src.SetHeaders_API_POST(w)
	if r.Method == http.MethodPost {
		token, oldPassword, newPassword := getChangeUserPasswordPostValues(r)
		users.ChangeUserPassword(token, oldPassword, newPassword, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleChangeUserPassword(router *mux.Router) {
	router.HandleFunc("/user.password.edit", changeUserPassword)
}
