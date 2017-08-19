package emotional_marks

import (
	"forcamp/conf"
	"forcamp/src"
	"time"
	"forcamp/src/api/orgset/settings"
	"strconv"
)

type EmotionalMark struct {
	IsNew bool   `json:"is_new"`
	Value int64   `json:"value"`
}

type EmotionalMark_Raw struct {
	Participant_ID int64
	Mark int64
}

func GetEmotionalMark(event_id int64, event_time string) (EmotionalMark, *conf.ApiResponse) {
	var emotionalMark EmotionalMark
	var rawEmotionalMark EmotionalMark_Raw
	err := src.CustomConnection.QueryRow("SELECT participant_id, mark " +
		"FROM emotional_marks WHERE id=?", event_id).Scan(&rawEmotionalMark.Participant_ID, &rawEmotionalMark.Mark)
	if err != nil {
		return emotionalMark, conf.ErrDatabaseQueryFailed
	}
	timeStamp, err := time.Parse("2006-01-02 15:04:05.999", event_time); if err != nil {
		return emotionalMark, conf.ErrDatabaseQueryFailed
	}
	duration := timeStamp.Sub(time.Now())
	organizationSettings, apiErr := settings.GetOrgSettings_Request(); if apiErr != nil {
		return emotionalMark, apiErr
	}
	emotionalMarkPeriod, err := strconv.ParseInt(organizationSettings.EmotionalMarkPeriod, 10, 64)
	if err != nil {
		return emotionalMark, conf.ErrDatabaseQueryFailed
	}
	if duration > time.Hour * time.Duration(emotionalMarkPeriod) * time.Duration(-1) {
		emotionalMark.IsNew = true
	} else {
		emotionalMark.IsNew = false
	}
	emotionalMark.Value = rawEmotionalMark.Mark
	return emotionalMark, nil
}
