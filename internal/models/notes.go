package models

import (
	"database/sql"
	"time"
)

type Note struct {
	ID       int
	Title    string
	Body     string
	Created  time.Time
	Modified time.Time
}

type NoteModel struct {
	DB *sql.DB
}

func (m *NoteModel) Insert(title, body string) (int, error) {
	query := `INSERT INTO notes (title, body, created, modified)
	    VALUES($1, $2, NOW(), NOW()) RETURNING id;`

	var id int
	err := m.DB.QueryRow(query, title, body).Scan(&id)

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

	err := m.DB.QueryRow(query, id).
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
	rows, err := m.DB.Query(query, n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var notes []*Note
	for i := 0; rows.Next(); i++ {
		var note Note
		err := rows.Scan(
			&note.ID,
			&note.Title,
			&note.Body,
			&note.Created,
			&note.Modified,
		)
		if err != nil {
			return nil, err
		}
		notes = append(notes, &note)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return notes, nil

}
