package sqlite

import (
	"database/sql"
	"fmt"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
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

func (s *Storage) RefreshBrands(brands []models.Brand) error {
	const op = "storage.brand.RefreshBrands"

	tx, err := s.db.Begin()
	if err != nil {
		fmt.Errorf("%s: begin tx: %w", op, err)
	}

	stmt, err := tx.Prepare("INSERT OR IGNORE INTO brand (autoria_id, name, category_id) VALUES (?, ?, ?)")
	if err != nil {
		tx.Rollback()
		fmt.Errorf("%s: prepare: %w", op, err)
	}
	defer stmt.Close()

	for _, brand := range brands {
		if _, err := stmt.Exec(brand.Marka, brand.Name, brand.Category); err != nil {
			tx.Rollback()
			fmt.Errorf("%s: exec: %w", op, err)
		}
	}

	if err := tx.Commit(); err != nil {
		fmt.Errorf("%s: commit: %w", op, err)
	}
	return nil
}
