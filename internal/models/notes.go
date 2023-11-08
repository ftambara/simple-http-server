package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Note struct {
	ID       int
	Title    string
	Body     string
	Created  time.Time
	Modified time.Time
}

type NoteModel struct {
	DB *pgxpool.Pool
}

func (m *NoteModel) Insert(title, body string) (int, error) {
	query := `INSERT INTO notes (title, body, created, modified)
	VALUES($1, $2, NOW(), NOW()) RETURNING id;`

	var id int
	err := m.DB.QueryRow(context.Background(), query, title, body).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *NoteModel) Get(id int) (*Note, error) {
	query := `SELECT id, title, body, created, modified
	FROM notes
	WHERE id=$1`

	n := &Note{}

	err := m.DB.QueryRow(context.Background(), query, id).
		Scan(&n.ID, &n.Title, &n.Body, &n.Created, &n.Modified)
	if err != nil {
		return nil, err
	}
	return n, nil
}

func (m *NoteModel) Latest(n int) ([]*Note, error) {
	query := `SELECT id, title, body, created, modified
	    FROM notes
	    ORDER BY modified DESC
	    LIMIT $1`

	// It's ok to ignore the error here, it will be available
	// after rows are closed
	rows, _ := m.DB.Query(context.Background(), query, n)
	notes, err := pgx.CollectRows(rows, pgx.RowToStructByName[Note])
	if err != nil {
		return nil, err
	}
	notesP := make([]*Note, len(notes))
	for i := 0; i < len(notes); i++ {
		notesP[i] = &notes[i]
	}
	return notesP, nil

}
