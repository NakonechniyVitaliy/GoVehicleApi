package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage/migrator"
	_ "github.com/mattn/go-sqlite3"
	// Driver for getting migrations from files
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type SqliteStorage struct {
	DB *sql.DB
}

func New(cfg *config.Config) (*SqliteStorage, error) {
	const op = "storage.sqlite.new"

	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := migrator.Run(cfg); err != nil {
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
