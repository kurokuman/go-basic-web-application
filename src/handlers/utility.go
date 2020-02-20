package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Println("Error parsing the template: ", err)
	}
	t.Execute(w, templateData)
}
