package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func noteView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Showing a note"))
}

func noteCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a note... (don't wait for it)"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/notes/view/", noteView)
	mux.HandleFunc("/notes/create/", noteCreate)

	port := ":4000"
	log.Printf("Starting server on %v", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
