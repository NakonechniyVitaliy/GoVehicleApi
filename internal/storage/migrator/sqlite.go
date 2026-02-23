package migrator

import (
	"errors"
	"fmt"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
)

type SQLiteMigrator struct {
	cfg *config.Config
}

func NewSQLiteMigrator(cfg *config.Config) *SQLiteMigrator {
	return &SQLiteMigrator{
		cfg,
	}
}

func (sqlMig SQLiteMigrator) Up() error {
	const op = "storage.sqlite.migrator.up"

	migrationsPath := "./internal/migrations/sqlite"

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("sqlite3://%s", sqlMig.cfg.StoragePath),
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")

			return nil
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	fmt.Println("migrations applied")
	return nil
}
