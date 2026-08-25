// Package controllers will store all the glue logic between models and views.
package controllers

import (
	"net/http"

	"github.com/sergioTheTower/calhoun/web_app_course/lenslocked/views"
)

type Static struct {
	Template views.Template
}

func (static Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	static.Template.Execute(w, nil)
}

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
