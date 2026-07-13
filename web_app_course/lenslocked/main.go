package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/sergioTheTower/calhoun/web_app_course/lenslocked/controllers"
	"github.com/sergioTheTower/calhoun/web_app_course/lenslocked/views"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))))
	fmt.Println("Starting Server...")
	http.ListenAndServe(":3000", r)
}
