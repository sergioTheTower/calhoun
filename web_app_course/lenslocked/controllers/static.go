// Package controllers will store all the glue logic between models and views.
package controllers

import (
	"net/http"

	"github.com/calhoun/web_app_course/app/views"
)

// StaticHandler will return a http.Handlerfunc.
func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
