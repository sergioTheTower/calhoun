// Package views will handle all the view logic.
package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	htmlTpl *template.Template
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

// Parse will take a filepath and try to open up the template.
func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("failed to parse template err %w", err)
	}
	return Template{htmlTpl: tpl}, nil
}

// Execute will set the http header and execute the template with the given data.
func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := t.htmlTpl.Execute(w, data); err != nil {
		log.Printf("Execute template: %v", err)
		http.Error(w, "Error executing the template", http.StatusInternalServerError)
	}
}
