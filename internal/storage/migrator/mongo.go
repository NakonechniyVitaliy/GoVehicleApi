package migrator

import (
	"errors"
	"fmt"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	// Driver for getting migrations from files
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type MongoMigrator struct {
	cfg *config.Config
}

func NewMongoMigrator(cfg *config.Config) *MongoMigrator {
	return &MongoMigrator{
		cfg,
	}
}

func (sqlMig MongoMigrator) Up() error {
	const op = "storage.mongo.migrator.up"

	migrationsPath := "./internal/migrations/mongo"

	m, err := migrate.New(
		"file://"+migrationsPath,
		sqlMig.cfg.MongoURI,
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
