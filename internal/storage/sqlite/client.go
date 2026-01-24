package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteStorage struct {
	DB *sql.DB
}

func New(path string) (*SqliteStorage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := migrate(db); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &SqliteStorage{DB: db}, nil
}

func (s *SqliteStorage) CloseDB() error {
	return s.DB.Close()
}

func (s *SqliteStorage) GetName() string {
	return "sqlite"
}

func migrate(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS brand (
	    category_id INTEGER,
	    cnt INTEGER,
	    country_id INTEGER,
	    eng TEXT,
	    marka_id INTEGER PRIMARY KEY,
	    name TEXT,
	    slang TEXT,
	    value INTEGER
	);`
	_, err := db.Exec(query)
	return err
}
