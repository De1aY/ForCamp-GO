package orgset_add

import (
	"net/http"
	"github.com/gorilla/mux"
	"wplay/conf"
	"wplay/src"
	"strings"
	"wplay/src/api/orgset/participants"
	"strconv"
)

func getAddParticipantPostValues(r *http.Request) (participants.Participant, string, *conf.ApiResponse){
	Token := r.PostFormValue("token")
	Name := strings.TrimSpace(strings.ToLower(r.PostFormValue("name")))
	Surname := strings.TrimSpace(strings.ToLower(r.PostFormValue("surname")))
	Middlename := strings.TrimSpace(strings.ToLower(r.PostFormValue("middlename")))
	Sex, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("sex")), 10, 64)
	if err != nil {
		return participants.Participant{}, "", conf.ErrSexNotINT
	}
	Team, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("team")), 10, 64)
	if err != nil {
		return participants.Participant{}, "", conf.ErrTeamNotINT
	}
	return participants.Participant{0, Name, Surname, Middlename, int(Sex), Team, nil}, Token, nil
}

func AddParticipantHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_POST(w)
	if r.Method == http.MethodPost {
		participant, token, APIerr := getAddParticipantPostValues(r)
		if APIerr != nil {
			APIerr.Print(w)
		} else {
			participants.AddParticipant(token, participant, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleAddParticipant(router *mux.Router)  {
	router.HandleFunc("/orgset.participant.add", AddParticipantHandler)
}
