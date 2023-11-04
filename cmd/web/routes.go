package main

import "net/http"

func (app *application) serveMux() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(sanitizedFS{http.Dir("./ui/static/")})
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/notes/view", app.noteView)
	mux.HandleFunc("/notes/create", app.noteCreate)
	return mux
}
