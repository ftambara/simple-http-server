package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger

	// Methods defined in ./handlers.go and ./helpers.go
}

func main() {
	port := flag.String("port", ":4000", "HTTP network address")

	flag.Parse()

	app := application{
		log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	srv := &http.Server{
		Addr:     *port,
		ErrorLog: app.errorLog,
		Handler:  app.serveMux(),
	}

	app.infoLog.Printf("Starting server on %s", *port)
	err := srv.ListenAndServe()
	if err != nil {
		app.errorLog.Fatal(err)
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
