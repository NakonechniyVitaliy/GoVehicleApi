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
	queryBrands := `
	CREATE TABLE IF NOT EXISTS brands (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    marka_id INTEGER UNIQUE NOT NULL,
	    category_id INTEGER NOT NULL,
	    cnt INTEGER NOT NULL,
	    country_id INTEGER NOT NULL,
	    eng TEXT NOT NULL,
	    name TEXT NOT NULL,
	    slang TEXT NOT NULL,
	    value INTEGER NOT NULL
	);`

	queryVehicleTypes := `
	CREATE TABLE IF NOT EXISTS body_styles (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    name TEXT UNIQUE NOT NULL,
		value INTEGER UNIQUE NOT NULL
	);`

	queryVehicleCategories := `
	CREATE TABLE IF NOT EXISTS vehicle_categories (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		value INTEGER UNIQUE NOT NULL
	);`

	_, err := db.Exec(queryBrands)
	if err != nil {
		return err
	}
	_, err = db.Exec(queryVehicleTypes)
	if err != nil {
		return err
	}
	_, err = db.Exec(queryVehicleCategories)
	if err != nil {
		return err
	}
	return nil
}
