package events

import (
	"net/http"
	"github.com/gorilla/mux"
	"wplay/conf"
	"wplay/src"
	"wplay/src/api/events"
	"strconv"
	"strings"
	"wplay/src/tools"
	"wplay/src/api/orgset"
)

func getRequestData(r *http.Request) (string, int64, int64, int64, int, bool, *conf.ApiResponse) {
	token := r.FormValue("token")
	var (
		user_id int64
		rowsPerPage int64
		event_type int64
		page int64
		err error
		apiErr *conf.ApiResponse
	)
	rawUserID := strings.TrimSpace(r.FormValue("user_id"))
	if len(rawUserID) < 1 {
		user_id, apiErr = orgset.GetUserIdByToken(token); if apiErr != nil {
			return "", 0, 0, 0, 0, false, apiErr
		}
	} else {
		user_id, err = strconv.ParseInt(rawUserID, 10, 64)
		if err != nil {
			return "", 0, 0, 0, 0, false, conf.ErrIdIsNotINT
		}
	}
	rawAscending := strings.TrimSpace(r.FormValue("ascending"))
	ascending := tools.StringToBoolean(rawAscending)
	rawRowsPerPage := strings.TrimSpace(r.FormValue("rows_per_page"))
	if len(rawRowsPerPage) < 1 {
		rowsPerPage = 10
	} else {
		rowsPerPage, err =  strconv.ParseInt(rawRowsPerPage, 10, 64); if err != nil {
			return "", 0, 0, 0, 0, false, conf.ErrConvertStringToInt
		}
	}
	rawPage := strings.TrimSpace(r.FormValue("page"))
	if len(rawPage) < 1 {
		page = 1
	} else {
		page, err = strconv.ParseInt(rawPage, 10, 64); if err != nil {
			return "", 0, 0, 0, 0, false, conf.ErrConvertStringToInt
		}
	}
	rawEventType := strings.TrimSpace(r.FormValue("event_type"))
	if len(rawEventType) < 1 {
		event_type = -1
	} else {
		event_type, err = strconv.ParseInt(rawEventType, 10, 64); if err != nil {
			return "", 0, 0, 0, 0, false, conf.ErrConvertStringToInt
		}
	}
	return token, user_id, rowsPerPage, page, int(event_type), ascending, nil
}

func GetEventsHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_GET(w)
	if r.Method == http.MethodGet {
		token, user_id, rowsPerPage, page, event_type, ascending, apiErr := getRequestData(r); if apiErr != nil {
			apiErr.Print(w)
		} else {
			events.GetEvents(token, user_id, rowsPerPage, page, ascending, event_type, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleGetEvents(router *mux.Router)  {
	router.HandleFunc("/events.get", GetEventsHandler)
}
