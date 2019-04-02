/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package users

import (
	"wplay/src/api/orgset/teams"
	"wplay/src/api/events"
)

type getUserLogin_Success struct {
	ID int64 `json:"id"`
}

type getUserData_Success struct {
	Data UserData `json:"data"`
}

type UserData struct {
	Name           string         `json:"name"`
	Surname        string         `json:"surname"`
	Middlename     string         `json:"middlename"`
	Team           teams.Team     `json:"team"`
	Access         int            `json:"access"`
	Avatar         string         `json:"avatar"`
	Sex            int 			  `json:"sex"`
	Organization   string         `json:"organization"`
	Post           string         `json:"post"`
	Events         []events.Event `json:"events"`
	AdditionalData interface{}    `json:"additional_data"`
}
