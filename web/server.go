package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Handle API routes
	api := r.PathPrefix("/api/").Subrouter()
	api.HandleFunc("/student", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "From the API")
	})

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Serve index page on all unhandled routes
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	fmt.Println("http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
