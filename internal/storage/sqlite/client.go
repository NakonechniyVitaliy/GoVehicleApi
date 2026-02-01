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
	    marka_id INTEGER,
	    category_id INTEGER,
	    cnt INTEGER,
	    country_id INTEGER,
	    eng TEXT,
	    name TEXT,
	    slang TEXT,
	    value INTEGER 
	);`

	queryVehicleTypes := `
	CREATE TABLE IF NOT EXISTS vehicle_types (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    ablative TEXT,
	    category_id INTEGER,
	    name TEXT UNIQUE,
	    plural TEXT,
	    rewrite TEXT,
		singular TEXT
	);`

	queryVehicleCategories := `
	CREATE TABLE IF NOT EXISTS vehicle_categories (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		value INTEGER UNIQUE 
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
