/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package mainSite

import (
	"wplay/conf"
	"wplay/src"
	"wplay/src/api/orgset"
	"wplay/src/api/orgset/settings"
	"wplay/src/api/users"
	"wplay/src/tools"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type profileTemplateData struct {
	Token         string
	UserID        int64
	RequestUserID int64
	UserData      users.UserData
	RequestData   users.UserData
	OrgSettings   settings.OrgSettings
	// Flags
	IsRequestAdmin    bool
	IsRequestEmployee bool
	IsAdmin           bool
	IsEmployee        bool
	IsOwner           bool
}

var profileTemplateFuncMap = template.FuncMap{
	"stringToBoolean": tools.StringToBoolean,
	"toTitleCase":     tools.ToTitleCase,
	"isNegative":      tools.IsNegative,
	"timestampToDate": tools.TimestampToDate,
}

func ProfileHandler(responseWriter http.ResponseWriter, r *http.Request) {
	if r.TLS != nil {
		src.SetHeaders_Main(responseWriter)
		token, err := r.Cookie("token")
		if err != nil {
			http.Redirect(responseWriter, r, "https://" + conf.MAIN_SITE_DOMAIN + "/exit", http.StatusTemporaryRedirect)
			return
		}
		token.Value, err = url.QueryUnescape(token.Value)
		if err == nil && tools.CheckToken(token.Value) {
			profileHTML, err := template.New(conf.FILE_PROFILE).Funcs(profileTemplateFuncMap).ParseFiles(conf.FILE_PROFILE)
			if err != nil {
				responseWriter.WriteHeader(http.StatusInternalServerError)
				return
			}
			ptd, apiErr := getProfileTemplateData(token.Value, r)
			if apiErr != nil {
				if apiErr.Code == 618 {
					http.Redirect(responseWriter, r, "https://" + conf.MAIN_SITE_DOMAIN + "/404", http.StatusTemporaryRedirect)
					return
				} else {
					responseWriter.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
			profileHTML.ExecuteTemplate(responseWriter, "profile", ptd)
		} else {
			http.Redirect(responseWriter, r, "https://" + conf.MAIN_SITE_DOMAIN + "/exit", http.StatusTemporaryRedirect)
		}
	} else {
		http.Redirect(responseWriter, r, "https://"+r.Host+r.URL.Path, http.StatusTemporaryRedirect)
	}
}

func getProfileTemplateData(token string, r *http.Request) (profileTemplateData, *conf.ApiResponse) {
	var ptd profileTemplateData
	ptd.Token = token
	organization, login, apiErr := orgset.GetUserOrganizationAndIdByToken(token)
	if apiErr != nil {
		return ptd, apiErr
	}
	src.CustomConnection = src.Connect_Custom(organization)
	ptd.UserID = login
	rawRequestUserID := strings.ToLower(strings.TrimSpace(r.FormValue("id")))
	var requestUserID int64
	var err error
	if len(rawRequestUserID) < 1 {
		requestUserID = 0
	} else {
		requestUserID, err = strconv.ParseInt(rawRequestUserID, 10, 64)
		if err != nil {
			return ptd, conf.ErrIdIsNotINT
		}
	}
	if requestUserID < 1 {
		ptd.RequestUserID = ptd.UserID
	} else {
		ptd.RequestUserID = requestUserID
	}
	apiErr = ptd.GetOrgSettings()
	if apiErr != nil {
		return ptd, apiErr
	}
	apiErr = ptd.GetUserData()
	if apiErr != nil {
		return ptd, apiErr
	}
	apiErr = ptd.GetRequestData()
	if apiErr != nil {
		return ptd, apiErr
	}
	ptd.SetFlags()
	return ptd, nil
}

func (ptd *profileTemplateData) GetUserData() *conf.ApiResponse {
	userData, apiErr := users.GetUserData_Request(ptd.UserID)
	if apiErr != nil {
		return apiErr
	}
	ptd.UserData = userData
	return nil
}

func (ptd *profileTemplateData) GetOrgSettings() *conf.ApiResponse {
	orgSettings, apiErr := settings.GetOrgSettings_Request()
	if apiErr != nil {
		return apiErr
	}
	ptd.OrgSettings = orgSettings
	return nil
}

func (ptd *profileTemplateData) GetRequestData() *conf.ApiResponse {
	if ptd.UserID != ptd.RequestUserID {
		requestOrganization, apiErr := orgset.GetUserOrganizationByID(ptd.RequestUserID)
		userOrganization, apiErr := orgset.GetUserOrganizationByID(ptd.UserID)
		if requestOrganization != userOrganization {
			return conf.ErrUserNotFound
		}
		requestData, apiErr := users.GetUserData_Request(ptd.RequestUserID)
		if apiErr != nil {
			return apiErr
		}
		ptd.RequestData = requestData
	} else {
		ptd.RequestData = ptd.UserData
	}
	return nil
}

func (ptd *profileTemplateData) SetFlags() {
	ptd.setFlag_IsOwner()
	ptd.setFlag_IsEmployee()
	ptd.setFlag_IsAdmin()
}

func (ptd *profileTemplateData) setFlag_IsAdmin() {
	if ptd.UserData.Access == 2 {
		ptd.IsAdmin = true
		ptd.IsEmployee = true
	}
	if ptd.RequestData.Access == 2 {
		ptd.IsRequestAdmin = true
		ptd.IsRequestEmployee = true
	}
}

func (ptd *profileTemplateData) setFlag_IsEmployee() {
	if ptd.UserData.Access == 1 {
		ptd.IsEmployee = true
	}
	if ptd.RequestData.Access == 1 {
		ptd.IsRequestEmployee = true
	}
}

func (ptd *profileTemplateData) setFlag_IsOwner() {
	if ptd.UserID == ptd.RequestUserID {
		ptd.IsOwner = true
	}
}
