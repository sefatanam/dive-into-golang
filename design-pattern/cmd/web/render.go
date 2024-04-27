package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type TemplateData struct {
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, templateName string, templateData *TemplateData) {
	var tmpl *template.Template

	// if we are using template cache, try to get the template from our map and stored in the receiver

	if app.config.useCache {
		if templateFromMap, ok := app.templateMap[templateName]; ok {
			tmpl = templateFromMap
		}
	}

	// if not found in cache
	if tmpl == nil {
		newTemplate, err := app.buildTemplateFromDisk(templateName)
		if err != nil {
			log.Println("Error building template:", err)
			return
		}

		log.Println("builing template from disk success")
		tmpl = newTemplate
	}

	if templateData == nil {
		templateData = &TemplateData{}
	}

	if err := tmpl.ExecuteTemplate(w, templateName, templateData); err != nil {
		log.Println("Error occured:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *application) buildTemplateFromDisk(templateName string) (*template.Template, error) {

	templatesSlice := []string{
		"./templates/base.layout.gohtml",
		"./templates/partials/header.partial.gohtml",
		"./templates/partials/footer.partial.gohtml",
		fmt.Sprintf("./templates/%s", templateName),
	}

	tmpl, err := template.ParseFiles(templatesSlice...)
	if err != nil {
		return nil, err
	}

	app.templateMap[templateName] = tmpl

	return tmpl, nil
}
