package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type snippetModel struct {
	DB *sql.DB
}

func (m *snippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

func (m *snippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

func (m *snippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
