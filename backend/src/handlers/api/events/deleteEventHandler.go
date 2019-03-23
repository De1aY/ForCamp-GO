package events

import (
	"net/http"
	"github.com/gorilla/mux"
	"wplay/conf"
	"wplay/src"
	"strings"
	"strconv"
	"wplay/src/api/events"
)

func getDeleteEventPostValues(r *http.Request) (string, int64, *conf.ApiResponse){
	token := strings.TrimSpace(r.PostFormValue("token"))
	event_id, err := strconv.ParseInt(strings.ToLower(
		strings.TrimSpace(r.PostFormValue("event_id"))), 10, 64)
	if err != nil {
		return "", 0, conf.ErrIdIsNotINT
	}
	return token, event_id, nil
}

func DeleteEventHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_POST(w)
	if r.Method == http.MethodPost {
		token, event_id, apiErr := getDeleteEventPostValues(r); if apiErr != nil {
			apiErr.Print(w)
		} else {
			events.DeleteEvent(token, event_id, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleDeleteEvent(router *mux.Router)  {
	router.HandleFunc("/event.delete", DeleteEventHandler)
}
