package migrator

import (
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/config"
	consts "github.com/NakonechniyVitalii/GoVehicleApi/internal/constants"
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
