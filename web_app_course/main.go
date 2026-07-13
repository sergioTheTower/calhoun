package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcomeeeeeee to The O.C Bitch! </h1>")
}

func main() {
	fmt.Println("hello world!")
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting server....")
	http.ListenAndServe(":8080", nil)
}
