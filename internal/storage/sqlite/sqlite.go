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
	    category_id INTEGER,
	    cnt INTEGER,
	    country_id INTEGER,
	    eng STRING,
    	marka_id INTEGER PRIMARY KEY,
    	name STRING,
		slang STRING,
    	value INTEGER
	);`

	createIndex := `
	CREATE INDEX IF NOT EXISTS marka_id ON brand (marka_id);
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

	stmt, err := tx.Prepare("INSERT OR IGNORE INTO brand (category_id, cnt, country_id, eng, marka_id, name, slang, value) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		fmt.Errorf("%s: prepare: %w", op, err)
	}
	defer stmt.Close()

	for _, brand := range brands {
		if _, err := stmt.Exec(brand.Category, brand.Count, brand.Country, brand.EngName, brand.Marka, brand.Name, brand.Slang, brand.Value); err != nil {
			tx.Rollback()
			fmt.Errorf("%s: exec: %w", op, err)
		}
	}

	if err := tx.Commit(); err != nil {
		fmt.Errorf("%s: commit: %w", op, err)
	}
	return nil
}

func (s *Storage) GetBrands() ([]models.Brand, error) {
	const op = "storage.brand.GetBrands"

	rows, err := s.db.Query("SELECT category_id, cnt, country_id, eng, marka_id, name, slang, value FROM brand")
	if err != nil {
		fmt.Errorf("%s: exec: %w", op, err)
	}
	defer rows.Close()

	var brands []models.Brand
	for rows.Next() {
		var b models.Brand
		if err := rows.Scan(
			&b.Category, &b.Count, &b.Country, &b.EngName, &b.Marka, &b.Name, &b.Slang, &b.Value); err != nil {
			return nil, fmt.Errorf("%s: scan: %w", op, err)
		}
		brands = append(brands, b)
	}

	return brands, nil

}
