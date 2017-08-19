package users_get

import (
	"forcamp/src/api/users"
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/handlers"
	"strings"
	"strconv"
)

func getUserID(r *http.Request) (int64, *conf.ApiResponse) {
	var user_id int64
	var err error
	rawUser_id := strings.TrimSpace(r.FormValue("user_id"))
	if len(rawUser_id) < 1 {
		user_id = -1
	} else {
		user_id, err = strconv.ParseInt(rawUser_id, 10, 64)
		if err != nil {
			return 0, conf.ErrIdIsNotINT
		}
	}
	return user_id, nil
}

func GetUserDataHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_GET(w)
	if r.Method == http.MethodGet {
		user_id, apiErr := getUserID(r); if apiErr != nil {
			apiErr.Print(w)
		} else {
			users.GetUserData(handlers.GetToken(r), w, user_id)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleGetUserData(router *mux.Router)  {
	router.HandleFunc("/user.data.get", GetUserDataHandler)
}
