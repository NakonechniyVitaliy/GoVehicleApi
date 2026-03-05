package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	consts "github.com/NakonechniyVitaliy/GoVehicleApi/internal/constants"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/router"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/services"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage/mongo"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage/sqlite"

	repositories "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository"

	bodyStyleRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/body_style"
	brandRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/brand"
	categoryRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/category"
	driverTypeRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/driver_type"
	gearboxRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/gearbox"
	vehicleRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle"

	bodyService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/body_style"
	brandService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/brand"
	categoryService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/category"
	driverService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/driver_type"
	gearboxService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/gearbox"
	vehicleService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/vehicle"
)

type App struct {
	Router  http.Handler
	Storage storage.Storage
}

func New(log *slog.Logger, cfg *config.Config) (*App, error) {

	appStorage, err := setupStorage(cfg)
	if err != nil {
		return nil, err
	}

	repos, err := setupRepositories(appStorage)
	if err != nil {
		return nil, err
	}

	serviceContainer := setupServices(repos, log, cfg.AutoriaKey)

	appRouter := router.SetupRouter(log, serviceContainer)

	return &App{
		Router:  appRouter,
		Storage: appStorage,
	}, nil
}

func setupStorage(cfg *config.Config) (storage.Storage, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	switch cfg.DataBase {
	case consts.MongoDB:
		return mongo.New(ctx, cfg)

	case consts.SqLite:
		return sqlite.New(cfg)

	default:
		return nil, fmt.Errorf("unknown database type: %s", cfg.DataBase)
	}
}

func setupRepositories(Storage storage.Storage) (*repositories.Repositories, error) {

	switch Storage.GetName() {
	case consts.MongoDB:
		// type assertion - get object from interface
		mongoStorage := Storage.(*mongo.MongoStorage)
		return &repositories.Repositories{
			Brand:      brandRepo.NewMongoBrandRepo(mongoStorage.DB),
			BodyStyle:  bodyStyleRepo.NewMongoBodyStyleRepo(mongoStorage.DB),
			Category:   categoryRepo.NewMongoCategoryRepo(mongoStorage.DB),
			DriverType: driverTypeRepo.NewMongoDriverTypeRepo(mongoStorage.DB),
			Gearbox:    gearboxRepo.NewMongoGearboxRepo(mongoStorage.DB),
			Vehicle:    vehicleRepo.NewMongoVehicleRepo(mongoStorage.DB),
		}, nil

	case consts.SqLite:
		// type assertion - get object from interface
		sqliteStorage := Storage.(*sqlite.SqliteStorage)
		return &repositories.Repositories{
			Brand:      brandRepo.NewSqliteBrandRepo(sqliteStorage.DB),
			BodyStyle:  bodyStyleRepo.NewSqliteBodyStyleRepo(sqliteStorage.DB),
			Category:   categoryRepo.NewSqliteCategoryRepo(sqliteStorage.DB),
			DriverType: driverTypeRepo.NewSqliteDriverTypeRepo(sqliteStorage.DB),
			Gearbox:    gearboxRepo.NewSqliteGearboxRepo(sqliteStorage.DB),
			Vehicle:    vehicleRepo.NewSqliteVehicleRepo(sqliteStorage.DB),
		}, nil

	default:
		return nil, fmt.Errorf("failed to setup %s repositories", Storage.GetName())
	}
}

func setupServices(repos *repositories.Repositories, log *slog.Logger, autoRiaKey string) services.Container {

	return services.Container{
		Brand:      brandService.NewService(repos.Brand, log, autoRiaKey),
		BodyStyle:  bodyService.NewService(repos.BodyStyle, log, autoRiaKey),
		Category:   categoryService.NewService(repos.Category, log, autoRiaKey),
		DriverType: driverService.NewService(repos.DriverType, log, autoRiaKey),
		Gearbox:    gearboxService.NewService(repos.Gearbox, log, autoRiaKey),
		Vehicle:    vehicleService.NewService(repos, log),
	}
}
