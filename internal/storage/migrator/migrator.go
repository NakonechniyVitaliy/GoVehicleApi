package migrator

import (
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	consts "github.com/NakonechniyVitaliy/GoVehicleApi/internal/constants"
)

type MigratorInterface interface {
	Up() error
}

func Run(cfg *config.Config) error {
	var m MigratorInterface

	switch cfg.DataBase {
	case consts.SqLite:
		m = NewSQLiteMigrator(cfg)
	case consts.MongoDB:
		m = NewMongoMigrator(cfg)
	}

	return m.Up()
}
