package mainSite

import (
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/api/orgset"
	"forcamp/src/api/orgset/categories"
	"forcamp/src/api/orgset/settings"
	"forcamp/src/api/users"
	"forcamp/src/tools"
	"html/template"
	"net/http"
	"net/url"
)

type generalTemplateData struct {
	Token       string
	UserID      int64
	UserData    users.UserData
	OrgSettings settings.OrgSettings
	Categories  []categories.Category
	// Flags
	IsAdmin    bool
	IsEmployee bool
	IsOwner    bool
}

var generalTemplateFuncMap = template.FuncMap{
	"stringToBoolean": tools.StringToBoolean,
	"toTitleCase":     tools.ToTitleCase,
	"isNegative":      tools.IsNegative,
	"timestampToDate": tools.TimestampToDate,
}

func GeneralHandler(responseWriter http.ResponseWriter, r *http.Request) {
	if r.TLS != nil {
		src.SetHeaders_Main(responseWriter)
		token, err := r.Cookie("token")
		if err != nil {
			http.Redirect(responseWriter, r, "https://forcamp.ga/exit", http.StatusTemporaryRedirect)
		}
		token.Value, err = url.QueryUnescape(token.Value)
		if err == nil && tools.CheckToken(token.Value) {
			generalHTML, err := template.New(conf.FILE_GENERAL).Funcs(generalTemplateFuncMap).ParseFiles(conf.FILE_GENERAL)
			if err != nil {
				responseWriter.WriteHeader(http.StatusInternalServerError)
				return
			}
			gtd, apiErr := getGeneralTemplateData(token.Value, r)
			if apiErr != nil {
				responseWriter.WriteHeader(http.StatusInternalServerError)
				return
			}
			generalHTML.ExecuteTemplate(responseWriter, "general", gtd)
		} else {
			http.Redirect(responseWriter, r, "https://forcamp.ga/exit", http.StatusTemporaryRedirect)
		}
	} else {
		http.Redirect(responseWriter, r, "https://"+r.Host+r.URL.Path, http.StatusTemporaryRedirect)
	}
}

func getGeneralTemplateData(token string, r *http.Request) (generalTemplateData, *conf.ApiResponse) {
	var gtd generalTemplateData
	gtd.Token = token
	organization, login, apiErr := orgset.GetUserOrganizationAndIdByToken(token)
	if apiErr != nil {
		return gtd, apiErr
	}
	src.CustomConnection = src.Connect_Custom(organization)
	gtd.UserID = login
	apiErr = gtd.GetOrgSettings()
	if apiErr != nil {
		return gtd, apiErr
	}
	apiErr = gtd.GetUserData()
	if apiErr != nil {
		return gtd, apiErr
	}
	apiErr = gtd.GetCategories()
	if apiErr != nil {
		return gtd, apiErr
	}
	gtd.SetFlags()
	return gtd, nil
}

func (gtd *generalTemplateData) GetUserData() *conf.ApiResponse {
	userData, apiErr := users.GetUserData_Request(gtd.UserID)
	if apiErr != nil {
		return apiErr
	}
	gtd.UserData = userData
	return nil
}

func (gtd *generalTemplateData) GetOrgSettings() *conf.ApiResponse {
	orgSettings, apiErr := settings.GetOrgSettings_Request()
	if apiErr != nil {
		return apiErr
	}
	gtd.OrgSettings = orgSettings
	return nil
}

func (gtd *generalTemplateData) GetCategories() *conf.ApiResponse {
	categories, apiErr := categories.GetCategories_Request()
	if apiErr != nil {
		return apiErr
	}
	gtd.Categories = categories
	return nil
}

func (gtd *generalTemplateData) SetFlags() {
	gtd.setFlag_IsEmployee()
	gtd.setFlag_IsAdmin()
}

func (gtd *generalTemplateData) setFlag_IsAdmin() {
	if gtd.UserData.Access == 2 {
		gtd.IsAdmin = true
		gtd.IsEmployee = true
	}
}

func (gtd *generalTemplateData) setFlag_IsEmployee() {
	if gtd.UserData.Access == 1 {
		gtd.IsEmployee = true
	}
}
