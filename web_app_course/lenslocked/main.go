package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/calhoun/web_app_course/app/controllers"
	"github.com/calhoun/web_app_course/app/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))))
	fmt.Println("Starting Server...")
	http.ListenAndServe(":3000", r)
}
