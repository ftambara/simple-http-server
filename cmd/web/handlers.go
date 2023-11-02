package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func printServerError(w http.ResponseWriter, err error) {
	log.Print(err.Error())
	http.Error(w, "Internal Server Error", 500)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		printServerError(w, err)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		printServerError(w, err)
		return
	}
}

func noteView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"id": "%d ", "body": "Building an HTTP server with Go!"}`, id)
}

func noteCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Creating a note... (don't wait for it)"))
}
