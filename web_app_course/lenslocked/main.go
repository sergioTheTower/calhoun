package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sergioTheTower/calhoun/web_app_course/lenslocked/controllers"
	"github.com/sergioTheTower/calhoun/web_app_course/lenslocked/templates"
	"github.com/sergioTheTower/calhoun/web_app_course/lenslocked/views"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml"))))
	r.Get("/faq", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml"))))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Healthy! Starting the server on :3000")
	http.ListenAndServe(":3000", r)
}
