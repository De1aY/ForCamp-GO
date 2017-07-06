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
)

type generalTemplateData struct {
	Token string
	Login string
	UserData users.UserData
	OrgSettings settings.OrgSettings
	// Flags
	IsAdmin bool
	IsEmployee bool
}

func GeneralHandler(w http.ResponseWriter, r *http.Request) {
	if r.TLS != nil {
		src.SetHeaders_Main(w)
		token, err := r.Cookie("token");
		if err == nil && tools.CheckToken(token.Value) {
			generalHtml, err := template.ParseFiles(conf.FILE_GENERAL); if err != nil {
				w.WriteHeader(http.StatusInternalServerError);
			}
			gtd, apiErr := getGeneralTemplateData(token.Value); if apiErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			generalHtml.ExecuteTemplate(w, "general", gtd);
		} else {
			http.Redirect(w, r, "https://forcamp.ga/exit", http.StatusTemporaryRedirect)
		}
	} else {
		http.Redirect(w, r, "https://" + r.Host + r.URL.Path, http.StatusTemporaryRedirect)
	}
}

func getGeneralTemplateData (token string) (generalTemplateData, *conf.ApiResponse) {
	var gtd generalTemplateData
	Organization, _, apiErr := orgset.GetUserOrganizationAndLoginByToken(token); if apiErr != nil {
		return gtd, apiErr
	}
	src.CustomConnection = src.Connect_Custom(Organization)
	apiErr = getGeneralTemplateData_Login(token, &gtd); if apiErr != nil {
		return gtd, apiErr
	}
	apiErr = getGeneralTemplateData_UserData(&gtd); if apiErr != nil {
		return gtd, apiErr
	}
	apiErr = getGeneralTemplateData_OrgSettings(&gtd); if apiErr != nil {
		return gtd, apiErr
	}
	getGeneralTemplateData_SetFlags(&gtd)
	return gtd, nil
}

func getGeneralTemplateData_Login (token string, gtd *generalTemplateData) *conf.ApiResponse {
	login, apiErr := users.GetUserLogin_Request(token)
	if apiErr != nil {
		return apiErr
	}
	gtd.Login = login
	return nil
}

func getGeneralTemplateData_UserData (gtd *generalTemplateData) *conf.ApiResponse {
	userData, apiErr := users.GetUserData_Request(gtd.Login)
	if apiErr != nil {
		return apiErr
	}
	gtd.UserData = userData
	return nil
}

func getGeneralTemplateData_OrgSettings (gtd *generalTemplateData) *conf.ApiResponse {
	orgSettings, apiErr := settings.GetOrgSettings_Query()
	if apiErr != nil {
		return apiErr
	}
	gtd.OrgSettings = orgSettings
	return nil
}

// Flags

func getGeneralTemplateData_SetFlags(gtd *generalTemplateData) {
	getGeneralTemplateData_SetFlagIsAdmin(gtd);
	getGeneralTemplateData_SetFlagIsEmployee(gtd);
}

func getGeneralTemplateData_SetFlagIsAdmin(gtd *generalTemplateData) {
	if gtd.UserData.Access == 2 {
		gtd.IsAdmin = true
		gtd.IsEmployee = true
	}
}

func getGeneralTemplateData_SetFlagIsEmployee(gtd *generalTemplateData) {
	if gtd.UserData.Access == 1 {
		gtd.IsEmployee = true
	}
}