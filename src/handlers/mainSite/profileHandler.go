package mainSite

import (
	"forcamp/src/api/users"
	"forcamp/src/api/orgset/settings"
	"net/http"
	"forcamp/src"
	"net/url"
	"forcamp/src/tools"
	"html/template"
	"forcamp/conf"
	"forcamp/src/api/orgset"
	"strings"
)

type profileTemplateData struct {
	Token string
	Login string
	RequestLogin string
	UserData users.UserData
	RequestData users.UserData
	OrgSettings settings.OrgSettings
	// Flags
	IsRequestAdmin bool
	IsRequestEmployee bool
	IsAdmin bool
	IsEmployee bool
	IsOwner bool
}

var profileTemplateFuncMap = template.FuncMap{
	"stringToBoolean": tools.StringToBoolean,
	"toTitleCase": tools.ToTitleCase,
	"isNegative": tools.IsNegative,
}

func ProfileHandler(responseWriter http.ResponseWriter, r *http.Request) {
	if r.TLS != nil {
		src.SetHeaders_Main(responseWriter)
		token, err := r.Cookie("token")
		token.Value, err = url.QueryUnescape(token.Value)
		if err == nil && tools.CheckToken(token.Value) {
			profileHTML, err := template.New(conf.FILE_PROFILE).Funcs(profileTemplateFuncMap).ParseFiles(conf.FILE_PROFILE); if err != nil {
				responseWriter.WriteHeader(http.StatusInternalServerError)
				return
			}
			ptd, apiErr := getProfileTemplateData(token.Value, r);
			if apiErr != nil {
				if apiErr.Code == 618 {
					http.Redirect(responseWriter, r, "https://forcamp.ga/404", http.StatusTemporaryRedirect)
				} else {
					responseWriter.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
			profileHTML.ExecuteTemplate(responseWriter, "profile", ptd)
		} else {
			http.Redirect(responseWriter, r, "https://forcamp.ga/exit", http.StatusTemporaryRedirect)
		}
	} else {
		http.Redirect(responseWriter, r, "https://" + r.Host + r.URL.Path, http.StatusTemporaryRedirect)
	}
}

func getProfileTemplateData(token string, r *http.Request) (profileTemplateData, *conf.ApiResponse) {
	var ptd profileTemplateData
	ptd.Token = token
	organization, login, apiErr := orgset.GetUserOrganizationAndLoginByToken(token); if apiErr != nil {
		return ptd, apiErr
	}
	src.CustomConnection = src.Connect_Custom(organization)
	ptd.Login = login
	requestLogin := strings.ToLower(strings.TrimSpace(r.FormValue("login")))
	if len(requestLogin) == 0 {
		ptd.RequestLogin = ptd.Login
	} else {
		ptd.RequestLogin = requestLogin
	}
	apiErr = ptd.GetOrgSettings(); if apiErr != nil {
		return ptd, apiErr
	}
	apiErr = ptd.GetUserData(); if apiErr != nil {
		return ptd, apiErr
	}
	apiErr = ptd.GetRequestData(); if apiErr != nil {
		return ptd, apiErr
	}
	ptd.SetFlags()
	return ptd, nil
}

func (ptd *profileTemplateData) GetUserData() *conf.ApiResponse {
	userData, apiErr := users.GetUserData_Request(ptd.Login)
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
	if ptd.Login != ptd.RequestLogin {
		requestOrganization, apiErr := orgset.GetUserOrganizationByLogin(ptd.RequestLogin)
		userOrganization, apiErr := orgset.GetUserOrganizationByLogin(ptd.Login)
		if requestOrganization != userOrganization {
			return conf.ErrUserNotFound
		}
		requestData, apiErr := users.GetUserData_Request(ptd.RequestLogin)
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
	if ptd.Login == ptd.RequestLogin {
		ptd.IsOwner = true
	}
}