/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package apanel

import (
	"nullteam.info/wplay/demo/conf"
	"nullteam.info/wplay/demo/src"
	"nullteam.info/wplay/demo/src/api/orgset"
)

type createOrganization_Success struct {
	AdminLogin    string `json:"admin_login"`
	AdminPassword string `json:"admin_password"`
}

func isUserAdmin(token string) *conf.ApiResponse {
	user_id, apiErr := orgset.GetUserIdByToken(token)
	if apiErr != nil {
		return apiErr
	}
	var admin_status bool
	err := src.Connection.QueryRow("SELECT admin FROM users WHERE id=?", user_id).Scan(&admin_status)
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	if admin_status {
		return nil
	} else {
		return conf.ErrInsufficientRights
	}
}
