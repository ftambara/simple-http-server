package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	port := flag.String("port", ":4000", "HTTP network address")

	flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(sanitizedFS{http.Dir("./ui/static/")})
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/notes/view", noteView)
	mux.HandleFunc("/notes/create", noteCreate)

	log.Printf("Starting server on %s", *port)
	err := http.ListenAndServe(*port, mux)
	if err != nil {
		log.Fatal(err)
	}
}

type sanitizedFS struct {
	dir http.Dir
}

// Prevent serving of static files directly
func (sfs sanitizedFS) Open(path string) (http.File, error) {
	f, err := sfs.dir.Open(path)
	// Whatever was asked could truly not be found
	if err != nil {
		return nil, err
	}
	// If it's a directory, check if it has an index.html
	s, err := f.Stat()
	if err != nil {
		return nil, err
	} else if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err = sfs.dir.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil

}
