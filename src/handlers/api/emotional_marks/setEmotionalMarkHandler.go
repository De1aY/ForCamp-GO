/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package emotional_marks

import (
	"net/http"
	"nullteam.info/wplay/demo/src"
	"nullteam.info/wplay/demo/src/api/emotional_marks"
	"nullteam.info/wplay/demo/conf"
	"strconv"
	"strings"
	"github.com/gorilla/mux"
)

func getRequestData(r *http.Request) (string, int64, *conf.ApiResponse) {
	token := r.PostFormValue("token")
	emotionalMark_value, err := strconv.ParseInt(strings.ToLower(
		strings.TrimSpace(r.PostFormValue("value"))), 10, 64)
	if err != nil {
		return "", 0, conf.ErrEmotionalMarkValueIncorrect
	}
	return token, emotionalMark_value, nil
}

func SetEmotionalMarkHandler(responseWriter http.ResponseWriter, r *http.Request) {
	src.SetHeaders_API_POST(responseWriter)
	if r.Method == http.MethodPost {
		token, emotionalMark_value, apiErr := getRequestData(r); if apiErr != nil {
			apiErr.Print(responseWriter)
		} else {
			emotional_marks.SetEmotionalMark(token, emotionalMark_value, responseWriter)
		}
	} else {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(responseWriter)
	}
}

func HandleSetEmotionalMark(router *mux.Router)  {
	router.HandleFunc("/emotional_mark.set", SetEmotionalMarkHandler)
}
