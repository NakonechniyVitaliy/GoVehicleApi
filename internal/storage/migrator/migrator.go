package storage

import (
	"errors"
	"fmt"

	consts "github.com/NakonechniyVitaliy/GoVehicleApi/internal/constants"
	// Lib for migrations
	"github.com/golang-migrate/migrate/v4"
	// Driver for migration SQL3
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	// Driver for getting migrations from files
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main(storagePath string) error {
	const op = "storage.sqlite.new"


	var migrationsPath string
	switch storagePath {
		case consts.SqLite:
			&migrationsPath = "./migrations/sqlite"
		case consts.MongoDB:
			&migrationsPath := "./migrations/mongo"
		}

	if cfg.MigrationsPath == "" {
		return fmt.Errorf("%s: %w", op, "migrations path is required!")
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("sqlite3://%s", storagePath),
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
}
