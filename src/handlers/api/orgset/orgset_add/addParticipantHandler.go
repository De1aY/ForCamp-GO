package orgset_add

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strings"
	"forcamp/src/api/orgset/participants"
	"strconv"
	"log"
)

func getAddParticipantPostValues(r *http.Request) (participants.Participant, string, *conf.ApiResponse){
	Token := r.PostFormValue("token")
	Name := strings.TrimSpace(strings.ToLower(r.PostFormValue("name")))
	Surname := strings.TrimSpace(strings.ToLower(r.PostFormValue("surname")))
	Middlename := strings.TrimSpace(strings.ToLower(r.PostFormValue("middlename")))
	Sex, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("sex")), 10, 64)
	if err != nil {
		log.Print(err)
		return participants.Participant{}, "", conf.ErrSexNotINT
	}
	Team, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("team")), 10, 64)
	if err != nil {
		log.Print(err)
		return participants.Participant{}, "", conf.ErrTeamNotINT
	}
	return participants.Participant{"", Name, Surname, Middlename, int(Sex), Team, nil}, Token, nil
}

func AddParticipantHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
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
