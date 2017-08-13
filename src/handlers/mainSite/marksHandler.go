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
	"forcamp/src/api/orgset/categories"
)

type marksTemplateData struct {
	Token string
	Login string
	UserData users.UserData
	OrgSettings settings.OrgSettings
	Categories []categories.Category
	// Flags
	IsAdmin bool
	IsEmployee bool
}

var marksTemplateFuncMap = template.FuncMap{
	"stringToBoolean": tools.StringToBoolean,
	"toTitleCase": tools.ToTitleCase,
}

func MarksHandler(w http.ResponseWriter, r *http.Request) {
	if r.TLS != nil {
		src.SetHeaders_Main(w)
		token, err := r.Cookie("token");
		token.Value, err = url.QueryUnescape(token.Value)
		if err == nil && tools.CheckToken(token.Value) {
			marksHTML, err := template.New(conf.FILE_MARKS).Funcs(marksTemplateFuncMap).ParseFiles(conf.FILE_MARKS); if err != nil {
				w.WriteHeader(http.StatusInternalServerError);
			}
			mtd, apiErr := getMarksTemplateData(token.Value); if apiErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			if mtd.UserData.Access > 1 {
				marksHTML.ExecuteTemplate(w, "marks", mtd)
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

func getMarksTemplateData(token string) (marksTemplateData, *conf.ApiResponse) {
	var mtd marksTemplateData
	mtd.Token = token
	organization, login, apiErr := orgset.GetUserOrganizationAndLoginByToken(token); if apiErr != nil {
		return mtd, apiErr
	}
	src.CustomConnection = src.Connect_Custom(organization)
	mtd.Login = login
	apiErr = mtd.GetUserData(); if apiErr != nil {
		return mtd, apiErr
	}
	apiErr = mtd.GetOrgSettings(); if apiErr != nil {
		return mtd, apiErr
	}
	apiErr = mtd.GetCategories(); if apiErr != nil {
		return mtd, apiErr
	}
	mtd.SetFlags()
	return mtd, nil
}

func (mtd *marksTemplateData) GetUserData() *conf.ApiResponse {
	userData, apiErr := users.GetUserData_Request(mtd.Login)
	if apiErr != nil {
		return apiErr
	}
	mtd.UserData = userData
	return nil
}

func (mtd *marksTemplateData) GetOrgSettings() *conf.ApiResponse {
	orgSettings, apiErr := settings.GetOrgSettings_Request()
	if apiErr != nil {
		return apiErr
	}
	mtd.OrgSettings = orgSettings
	return nil
}

func (mtd *marksTemplateData) GetCategories() *conf.ApiResponse {
	categories, apiErr := categories.GetCategories_Request()
	if apiErr != nil {
		return apiErr
	}
	mtd.Categories = categories
	return nil
}

// Flags

func (mtd *marksTemplateData) SetFlags() {
	mtd.setFlag_IsAdmin();
	mtd.setFlag_IsEmployee();
}

func (mtd *marksTemplateData) setFlag_IsAdmin(){
	if mtd.UserData.Access == 2 {
		mtd.IsAdmin = true
		mtd.IsEmployee = true
	}
}

func (mtd *marksTemplateData) setFlag_IsEmployee() {
	if mtd.UserData.Access == 1 {
		mtd.IsEmployee = true
	}
}