package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))

}

func snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id wildcard from the request using r.PathValue() // and try to convert it to an integer using the strconv.Atoi() function. If // it can't be converted to an integer, or the value is less than 1, we
	// return a 404 page not found response.
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// Use the fmt.Sprintf() function to interpolate the id value with a // message, then write it as the HTTP response.
	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))

}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))

}

// Add a snippetCreatePost handler function.
func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Save a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView) // Add the {id} wildcard to the URL
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	// Create the new route, which is restricted to POST requests only.
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server in :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
