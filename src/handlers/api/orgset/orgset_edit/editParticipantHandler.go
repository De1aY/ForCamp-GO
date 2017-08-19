package orgset_edit

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strings"
	"forcamp/src/api/orgset/participants"
	"strconv"
)

func getEditParticipantPostValues(r *http.Request) (participants.Participant, string, *conf.ApiResponse){
	token := strings.TrimSpace(r.PostFormValue("token"))
	participant_id, err := strconv.ParseInt(strings.TrimSpace(
		strings.ToLower(r.PostFormValue("participant_id"))), 10, 64)
	if err != nil {
		return participants.Participant{}, "", conf.ErrIdIsNotINT
	}
	name := strings.TrimSpace(strings.ToLower(r.PostFormValue("name")))
	surname := strings.TrimSpace(strings.ToLower(r.PostFormValue("surname")))
	middlename := strings.TrimSpace(strings.ToLower(r.PostFormValue("middlename")))
	sex, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("sex")), 10, 64)
	if err != nil {
		return participants.Participant{}, "", conf.ErrSexNotINT
	}
	team, err := strconv.ParseInt(strings.TrimSpace(r.PostFormValue("team")), 10, 64)
	if err != nil {
		return participants.Participant{}, "", conf.ErrTeamNotINT
	}
	return participants.Participant{participant_id, name, surname,
		middlename, int(sex), team, nil}, token, nil
}

func EditParticipantHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API_POST(w)
	if r.Method == http.MethodPost {
		participant, token, apiErr := getEditParticipantPostValues(r)
		if apiErr != nil {
			apiErr.Print(w)
		} else {
			participants.EditParticipant(token, participant, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.ErrMethodNotAllowed.Print(w)
	}
}

func HandleEditParticipant(router *mux.Router)  {
	router.HandleFunc("/orgset.participant.edit", EditParticipantHandler)
}
