package main

import (
	"html/template"
	"net/http"
)


func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	
	_ = app.render(w,r,"home.page.gohtml",&TemplateData{})

}

type TemplateData struct {
	IP string
	// any is any alias to empty interface
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter,r *http.Request,t string,data *TemplateData) error {

	//  parse the template from disk
	parsedTemplate,err := template.ParseFiles("./templates/"+t)
	if err != nil {
		http.Error(w,"bad request",http.StatusBadRequest)
		return err
	}

	// execute the template, passing it data, if any
	err = parsedTemplate.Execute(w,data)
	if err != nil {
		return err
	}

	// parse the template
	return nil

}