package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello World!"))
}

func noteView(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/notes/view/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"body": "Building an HTTP server with Go!"}`))
}

func noteCreate(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/notes/create/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
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
