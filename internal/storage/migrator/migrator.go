package migrator

import (
	"errors"
	"fmt"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	consts "github.com/NakonechniyVitaliy/GoVehicleApi/internal/constants"
	// Lib for migrations
	"github.com/golang-migrate/migrate/v4"
	// Driver for migration SQL3
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	// Driver for getting migrations from files
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(cfg *config.Config) error {
	const op = "storage.sqlite.new"

	var migrationsPath string
	switch cfg.DataBase {
	case consts.SqLite:
		migrationsPath = "./internal/migrations/sqlite"
	case consts.MongoDB:
		migrationsPath = "./internal/migrations/mongo"
	}

	fmt.Println("file://"+migrationsPath, fmt.Sprintf("sqlite3://%s", cfg.StoragePath))

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("sqlite3://%s", cfg.StoragePath),
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
