package templates

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"forcamp/conf/paths"
)

func getTemplate_Index() *template.Template{
	TemplateIndex, err := template.ParseFiles(paths.TEMPLATE_INDEX)
	if err != nil {
		log.Fatal(err)
	}
	return TemplateIndex
}

func templateHandler_Index(w http.ResponseWriter, r *http.Request){
	TemplateIndex := getTemplate_Index()
	TemplateIndex.ExecuteTemplate(w, "index", nil)
}

func HandleTemplate_Index(router *mux.Router){
	router.HandleFunc("/index.html", templateHandler_Index)
}