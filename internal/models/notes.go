package models

import (
	"context"
	"time"

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
	Conn *pgxpool.Pool
}

func (m *NoteModel) Insert(title, body string) (int, error) {
	query := `INSERT INTO notes (title, body, created, modified)
	VALUES($1, $2, NOW(), NOW()) RETURNING id;`

	var id int
	err := m.Conn.QueryRow(context.Background(), query, title, body).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *NoteModel) Get(id int) (*Note, error) {
	return nil, nil
}

func (m *NoteModel) Latest(n int) ([]*Note, error) {
	return nil, nil
}
