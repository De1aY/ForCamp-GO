/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package users_edit

import (
	"wplay/conf"
	"wplay/src"
	"wplay/src/api/users"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func getChangeUserAvatarPostValues(r *http.Request) string {
	token := strings.TrimSpace(r.PostFormValue("token"))
	return token
}

func changeUserAvatar(w http.ResponseWriter, r *http.Request) {
	src.SetHeaders_API_POST(w)
	if r.Method == http.MethodPost {
		token := getChangeUserAvatarPostValues(r)
		users.ChangeAvatar(token, r, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleChangeUserAvatar(router *mux.Router) {
	router.HandleFunc("/user.avatar.edit", changeUserAvatar)
}
