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
	//"forcamp/src/api/orgset/categories"
)

type orgSetTemplateData struct {
	Token string
	Login string
	UserData users.UserData
	OrgSettings settings.OrgSettings
	//Categories []categories.Category
	// Flags
	IsAdmin bool
	IsEmployee bool
}

var orgSetTemplateFuncMap = template.FuncMap{
	"stringToBoolean": tools.StringToBoolean,
	"toTitleCase": tools.ToTitleCase,
};

func OrgSetHandler(w http.ResponseWriter, r *http.Request) {
	if r.TLS != nil {
		src.SetHeaders_Main(w)
		token, err := r.Cookie("token");
		if err == nil && tools.CheckToken(token.Value) {
			orgSetHTML, err := template.New(conf.FILE_ORGSET).Funcs(orgSetTemplateFuncMap).ParseFiles(conf.FILE_ORGSET); if err != nil {
				w.WriteHeader(http.StatusInternalServerError);
			}
			ostd, apiErr := getOrgSetTemplateData(token.Value); if apiErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			orgSetHTML.ExecuteTemplate(w, "orgset", ostd);
		} else {
			http.Redirect(w, r, "https://forcamp.ga/exit", http.StatusTemporaryRedirect)
		}
	} else {
		http.Redirect(w, r, "https://" + r.Host + r.URL.Path, http.StatusTemporaryRedirect)
	}
}

func getOrgSetTemplateData (token string) (orgSetTemplateData, *conf.ApiResponse) {
	var ostd orgSetTemplateData
	Organization, _, apiErr := orgset.GetUserOrganizationAndLoginByToken(token); if apiErr != nil {
		return ostd, apiErr
	}
	src.CustomConnection = src.Connect_Custom(Organization)
	apiErr = getOrgSetTemplateData_Login(token, &ostd); if apiErr != nil {
		return ostd, apiErr
	}
	apiErr = getOrgSetTemplateData_UserData(&ostd); if apiErr != nil {
		return ostd, apiErr
	}
	apiErr = getOrgSetTemplateData_OrgSettings(&ostd); if apiErr != nil {
		return ostd, apiErr
	}
	//apiErr = getOrgSetTemplateData_Categories(&ostd); if apiErr != nil {
	//	return ostd, apiErr
	//}
	getOrgSetTemplateData_SetFlags(&ostd)
	return ostd, nil
}

func getOrgSetTemplateData_Login (token string, ostd *orgSetTemplateData) *conf.ApiResponse {
	login, apiErr := users.GetUserLogin_Request(token)
	if apiErr != nil {
		return apiErr
	}
	ostd.Login = login
	return nil
}

func getOrgSetTemplateData_UserData (ostd *orgSetTemplateData) *conf.ApiResponse {
	userData, apiErr := users.GetUserData_Request(ostd.Login)
	if apiErr != nil {
		return apiErr
	}
	ostd.UserData = userData
	return nil
}

func getOrgSetTemplateData_OrgSettings (ostd *orgSetTemplateData) *conf.ApiResponse {
	orgSettings, apiErr := settings.GetOrgSettings_Request()
	if apiErr != nil {
		return apiErr
	}
	ostd.OrgSettings = orgSettings
	return nil
}

/*
func getOrgSetTemplateData_Categories (ostd *orgSetTemplateData) *conf.ApiResponse {
	categories, apiErr := categories.GetCategories_Request()
	if apiErr != nil {
		return apiErr
	}
	ostd.Categories = categories
	return nil
}*/

// Flags

func getOrgSetTemplateData_SetFlags(ostd *orgSetTemplateData) {
	getOrgSetTemplateData_SetFlagIsAdmin(ostd);
	getOrgSetTemplateData_SetFlagIsEmployee(ostd);
}

func getOrgSetTemplateData_SetFlagIsAdmin(ostd *orgSetTemplateData) {
	if ostd.UserData.Access == 2 {
		ostd.IsAdmin = true
		ostd.IsEmployee = true
	}
}

func getOrgSetTemplateData_SetFlagIsEmployee(ostd *orgSetTemplateData) {
	if ostd.UserData.Access == 1 {
		ostd.IsEmployee = true
	}
}