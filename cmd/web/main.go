package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5/pgxpool"

	"ftambara/simple-http-server/internal/models"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger

	notes *models.NoteModel
}

func main() {
	port := *flag.String("port", ":4000", "HTTP network address")

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	infoLog.Printf("Connecting to database")
	db, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := application{
		errorLog: errorLog,
		infoLog:  infoLog,
		notes:    &models.NoteModel{DB: db},
	}

	srv := &http.Server{
		Addr:     port,
		ErrorLog: errorLog,
		Handler:  app.serveMux(),
	}

	infoLog.Printf("Starting server on %s", port)
	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
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
