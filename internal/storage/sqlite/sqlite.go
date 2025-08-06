package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // init sqlite3 driver
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS brand (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	autoria_id INTEGER UNIQUE,
    	name TEXT,
    	category_id INTEGER
	);`

	createIndex := `
	CREATE INDEX IF NOT EXISTS autoria_id ON brand (autoria_id);
	`

	if _, err := db.Exec(createTable); err != nil {
		return nil, fmt.Errorf("%s: create table: %w", op, err)
	}

	if _, err := db.Exec(createIndex); err != nil {
		return nil, fmt.Errorf("%s: create index: %w", op, err)
	}

	return &Storage{db: db}, nil

}
