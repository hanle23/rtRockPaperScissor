package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates = func() *template.Template {
	files, err := getTemplateNames()
	if err != nil {
		log.Fatal("Failed to get template files:", err)
	}
	fmt.Println(files)
	return template.Must(template.ParseFiles(files...))
}()

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getTemplateNames() ([]string, error) {
	f, err := getRootPath()
	if err != nil {
		return nil, err
	}
	fp := f + "/views/templates"
	entries, err := getEntriesFromPath(fp)
	if err != nil {
		return nil, err
	}
	var templateList []string
	for _, e := range entries {
		templateList = append(templateList, fp+"/"+e.Name())
	}
	return templateList, nil
}
