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

	serviceContainer := setupServices(repos, cfg.AutoriaKey)

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

type Repositories struct {
	Brand      brandRepo.RepositoryInterface
	BodyStyle  bodyStyleRepo.RepositoryInterface
	Category   categoryRepo.RepositoryInterface
	DriverType driverTypeRepo.RepositoryInterface
	Gearbox    gearboxRepo.RepositoryInterface
	Vehicle    vehicleRepo.RepositoryInterface
}

func setupRepositories(Storage storage.Storage) (*Repositories, error) {

	switch Storage.GetName() {
	case consts.MongoDB:
		// type assertion - get object from interface
		mongoStorage := Storage.(*mongo.MongoStorage)
		return &Repositories{
			brandRepo.NewMongoBrandRepo(mongoStorage.DB),
			bodyStyleRepo.NewMongoBodyStyleRepo(mongoStorage.DB),
			categoryRepo.NewMongoCategoryRepo(mongoStorage.DB),
			driverTypeRepo.NewMongoDriverTypeRepo(mongoStorage.DB),
			gearboxRepo.NewMongoGearboxRepo(mongoStorage.DB),
			vehicleRepo.NewMongoVehicleRepo(mongoStorage.DB),
		}, nil

	case consts.SqLite:
		// type assertion - get object from interface
		sqliteStorage := Storage.(*sqlite.SqliteStorage)
		return &Repositories{
			brandRepo.NewSqliteBrandRepo(sqliteStorage.DB),
			bodyStyleRepo.NewSqliteBodyStyleRepo(sqliteStorage.DB),
			categoryRepo.NewSqliteCategoryRepo(sqliteStorage.DB),
			driverTypeRepo.NewSqliteDriverTypeRepo(sqliteStorage.DB),
			gearboxRepo.NewSqliteGearboxRepo(sqliteStorage.DB),
			vehicleRepo.NewSqliteVehicleRepo(sqliteStorage.DB),
		}, nil

	default:
		return nil, fmt.Errorf("failed to setup %s repositories", Storage.GetName())
	}
}

func setupServices(repos *Repositories, autoRiaKey string) services.Container {

	return services.Container{
		Brand:      brandService.NewService(repos.Brand, autoRiaKey),
		BodyStyle:  bodyService.NewService(repos.BodyStyle, autoRiaKey),
		Category:   categoryService.NewService(repos.Category, autoRiaKey),
		DriverType: driverService.NewService(repos.DriverType, autoRiaKey),
		Gearbox:    gearboxService.NewService(repos.Gearbox, autoRiaKey),
		Vehicle: vehicleService.NewService(
			repos.Vehicle,
			repos.Brand,
			repos.BodyStyle,
			repos.Category,
			repos.DriverType,
			repos.Gearbox,
		),
	}
}
