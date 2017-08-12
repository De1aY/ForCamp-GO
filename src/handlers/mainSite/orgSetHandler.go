package mainSite

import (
	"net/http"
	"forcamp/src"
	"forcamp/conf"
	"forcamp/src/tools"
	"text/template"
	"forcamp/src/api/users"
	"forcamp/src/api/orgset"
	"forcamp/src/api/orgset/settings"
	"net/url"
)

type orgSetTemplateData struct {
	Token string
	Login string
	UserData users.UserData
	OrgSettings settings.OrgSettings
	// Flags
	IsAdmin bool
	IsEmployee bool
}

var orgSetTemplateFuncMap = template.FuncMap{
	"stringToBoolean": tools.StringToBoolean,
	"toTitleCase": tools.ToTitleCase,
}

func OrgSetHandler(w http.ResponseWriter, r *http.Request) {
	if r.TLS != nil {
		src.SetHeaders_Main(w)
		token, err := r.Cookie("token");
		token.Value, err = url.QueryUnescape(token.Value)
		if err == nil && tools.CheckToken(token.Value) {
			orgSetHTML, err := template.New(conf.FILE_ORGSET).Funcs(orgSetTemplateFuncMap).ParseFiles(conf.FILE_ORGSET); if err != nil {
				w.WriteHeader(http.StatusInternalServerError);
			}
			ostd, apiErr := getOrgSetTemplateData(token.Value); if apiErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			if ostd.UserData.Access == 2 {
				orgSetHTML.ExecuteTemplate(w, "orgset", ostd)
			} else {
				http.Redirect(w, r, "https://forcamp.ga/general", http.StatusTemporaryRedirect)
			}
		} else {
			http.Redirect(w, r, "https://forcamp.ga/exit", http.StatusTemporaryRedirect)
		}
	} else {
		http.Redirect(w, r, "https://" + r.Host + r.URL.Path, http.StatusTemporaryRedirect)
	}
}

func getOrgSetTemplateData (token string) (orgSetTemplateData, *conf.ApiResponse) {
	var ostd orgSetTemplateData
	ostd.Token = token
	organization, login, apiErr := orgset.GetUserOrganizationAndLoginByToken(token); if apiErr != nil {
		return ostd, apiErr
	}
	src.CustomConnection = src.Connect_Custom(organization)
	ostd.Login = login
	apiErr = ostd.GetUserData(); if apiErr != nil {
		return ostd, apiErr
	}
	apiErr = ostd.GetOrgSettings(); if apiErr != nil {
		return ostd, apiErr
	}
	ostd.SetFlags()
	return ostd, nil
}

func (ostd *orgSetTemplateData) GetUserData() *conf.ApiResponse {
	userData, apiErr := users.GetUserData_Request(ostd.Login)
	if apiErr != nil {
		return apiErr
	}
	ostd.UserData = userData
	return nil
}

func (ostd *orgSetTemplateData) GetOrgSettings() *conf.ApiResponse {
	orgSettings, apiErr := settings.GetOrgSettings_Request()
	if apiErr != nil {
		return apiErr
	}
	ostd.OrgSettings = orgSettings
	return nil
}

// Flags

func (ostd *orgSetTemplateData) SetFlags() {
	ostd.setFlag_IsAdmin();
	ostd.setFlag_IsEmployee();
}

func (ostd *orgSetTemplateData) setFlag_IsAdmin(){
	if ostd.UserData.Access == 2 {
		ostd.IsAdmin = true
		ostd.IsEmployee = true
	}
}

func (ostd *orgSetTemplateData) setFlag_IsEmployee() {
	if ostd.UserData.Access == 1 {
		ostd.IsEmployee = true
	}
}