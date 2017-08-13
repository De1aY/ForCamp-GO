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
	Token := strings.TrimSpace(r.PostFormValue("token"))
	Login := strings.TrimSpace(strings.ToLower(r.PostFormValue("login")))
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
	return participants.Participant{Login, Name, Surname, Middlename, int(Sex), Team, nil}, Token, nil
}

func EditParticipantHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		participant, token, APIerr := getEditParticipantPostValues(r)
		if APIerr != nil {
			APIerr.Print(w)
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
