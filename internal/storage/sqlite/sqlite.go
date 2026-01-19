package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage"
	_ "github.com/mattn/go-sqlite3" // init sqlite3 driver
)

type SqliteStorage struct {
	db *sql.DB
}

func (s *SqliteStorage) GetBrand(ctx context.Context, brandID int) (*models.Brand, error) {
	const op = "storage.brand.GetBrand"

	const query = `
		SELECT
			category_id,
			cnt,
			country_id,
			eng,
			marka_id,
			name,
			slang,
			value
		FROM brand
		WHERE marka_id = ?
	`
	var brand models.Brand
	err := s.db.QueryRowContext(ctx, query, brandID).Scan(
		&brand.Category,
		&brand.Count,
		&brand.Country,
		&brand.EngName,
		&brand.Marka,
		&brand.Name,
		&brand.Slang,
		&brand.Value,
	)
	if err != nil {
		fmt.Errorf("%s: Error to return brand %w", op, err)
		return nil, err
	}

	return &brand, nil
}

func (s *SqliteStorage) DeleteBrand(ctx context.Context, brandID int) error {
	const op = "storage.brand.DeleteBrand"

	res, err := s.db.ExecContext(
		ctx,
		"DELETE FROM brand WHERE marka_id = ?",
		brandID,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if affected == 0 {
		return storage.ErrBrandNotFound
	}

	return nil
}

func (s *SqliteStorage) NewBrand(ctx context.Context, brand models.Brand) error {
	const op = "storage.brand.NewBrand"

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

	_, err = stmt.Exec(brand.Category, brand.Count, brand.Country, brand.EngName, brand.Marka, brand.Name, brand.Slang, brand.Value)
	if err != nil {
		tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		fmt.Errorf("%s: commit: %w", op, err)
	}

	return nil
}

func (sqlt *SqliteStorage) RefreshBrands() error {
	return nil
}

func New(storagePath string) (*SqliteStorage, error) {
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

	createIndex := `CREATE INDEX IF NOT EXISTS marka_id ON brand (marka_id);`

	if _, err := db.Exec(createTable); err != nil {
		return nil, fmt.Errorf("%s: create table: %w", op, err)
	}

	if _, err := db.Exec(createIndex); err != nil {
		return nil, fmt.Errorf("%s: create index: %w", op, err)
	}

	return &SqliteStorage{db: db}, nil
}

//stmt, err := tx.Prepare("INSERT OR IGNORE INTO brand (category_id, cnt, country_id, eng, marka_id, name, slang, value) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
//if err != nil {
//	tx.Rollback()
//	fmt.Errorf("%s: prepare: %w", op, err)
//}
//defer stmt.Close()

//func (s *Storage) RefreshBrands(brands []models.Brand) error {
//	const op = "storage.brand.RefreshBrands"
//
//	tx, err := s.db.Begin()
//	if err != nil {
//		fmt.Errorf("%s: begin tx: %w", op, err)
//	}
//
//	stmt, err := tx.Prepare("INSERT OR IGNORE INTO brand (autoria_id, name, category_id) VALUES (?, ?, ?)")
//	if err != nil {
//		tx.Rollback()
//		fmt.Errorf("%s: prepare: %w", op, err)
//	}
//	defer stmt.Close()
//
//	for _, brand := range brands {
//		if _, err := stmt.Exec(brand.Marka, brand.Name, brand.Category); err != nil {
//			tx.Rollback()
//			fmt.Errorf("%s: exec: %w", op, err)
//		}
//	}
//
//	if err := tx.Commit(); err != nil {
//		fmt.Errorf("%s: commit: %w", op, err)
//	}
//	return nil
//}
//	for _, brand := range brands {
//		if _, err := stmt.Exec(brand.Category, brand.Count, brand.Country, brand.EngName, brand.Marka, brand.Name, brand.Slang, brand.Value); err != nil {
//			tx.Rollback()
//			fmt.Errorf("%s: exec: %w", op, err)
//		}
//	}
//
//	if err := tx.Commit(); err != nil {
//		fmt.Errorf("%s: commit: %w", op, err)
//	}
//	return nil
// }

//func (s *Storage) GetBrands() ([]models.Brand, error) {
//	const op = "storage.brand.GetBrands"
//
//	rows, err := s.db.Query("SELECT category_id, cnt, country_id, eng, marka_id, name, slang, value FROM brand")
//	if err != nil {
//		fmt.Errorf("%s: exec: %w", op, err)
//	}
//	defer rows.Close()
//
//	var brands []models.Brand
//	for rows.Next() {
//		var b models.Brand
//		if err := rows.Scan(
//			&b.Category, &b.Count, &b.Country, &b.EngName, &b.Marka, &b.Name, &b.Slang, &b.Value); err != nil {
//			return nil, fmt.Errorf("%s: scan: %w", op, err)
//		}
//		brands = append(brands, b)
//	}
//
//	return brands, nil
//
//}
//
