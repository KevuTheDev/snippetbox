package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox\n"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Header()
	w.Write([]byte("Creating a new snippet...\n"))
}

func snippetCreateMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/snippet/create" {
		w.Header().Set("Allow", http.MethodPost)
	}
}
