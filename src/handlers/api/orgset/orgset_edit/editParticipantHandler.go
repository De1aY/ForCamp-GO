package orgset_edit

import (
	"net/http"
	"github.com/gorilla/mux"
	"forcamp/conf"
	"forcamp/src"
	"strings"
	"forcamp/src/orgset/participants"
	"strconv"
	"log"
)

func getEditParticipantPostValues(r *http.Request) (participants.Participant, string, *conf.ApiError){
	Token := r.PostFormValue("token")
	Login := strings.ToLower(r.PostFormValue("login"))
	Name := strings.ToLower(r.PostFormValue("name"))
	Surname := strings.ToLower(r.PostFormValue("surname"))
	Middlename := strings.ToLower(r.PostFormValue("middlename"))
	Sex, err := strconv.ParseInt(r.PostFormValue("sex"), 10, 64)
	if err != nil {
		log.Print(err)
		return participants.Participant{}, "", conf.ErrParticipantSexNotINT
	}
	Team, err := strconv.ParseInt(r.PostFormValue("team"), 10, 64)
	if err != nil {
		log.Print(err)
		return participants.Participant{}, "", conf.ErrParticipantTeamNotINT
	}
	return participants.Participant{Login, Name, Surname, Middlename, int(Sex), Team, nil}, Token, nil
}

func EditParticipantHandler(w http.ResponseWriter, r *http.Request){
	src.SetHeaders_API(w)
	if r.Method == http.MethodPost {
		participant, token, APIerr := getEditParticipantPostValues(r)
		if APIerr != nil {
			conf.PrintError(APIerr, w)
		} else {
			participants.EditParticipant(token, participant, w)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		conf.PrintError(conf.ErrMethodNotAllowed,  w)
	}
}

func HandleEditParticipant(router *mux.Router)  {
	router.HandleFunc("/orgset.participant.edit", EditParticipantHandler)
}