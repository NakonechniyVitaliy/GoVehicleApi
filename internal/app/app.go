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
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage/mongo"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage/sqlite"

	bodyStyleRep "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/body_style"
	brandRep "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/brand"
	categoryRep "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/category"
	driverTypeRep "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/driver_type"
	gearboxRep "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/gearbox"
	vehicleRep "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle"
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

	appRouter := router.SetupRouter(
		log,
		repos.Brand,
		repos.BodyStyle,
		repos.Category,
		repos.DriverType,
		repos.Gearbox,
		repos.Vehicle,
		cfg,
	)

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
	Brand      brandRep.RepositoryInterface
	BodyStyle  bodyStyleRep.RepositoryInterface
	Category   categoryRep.RepositoryInterface
	DriverType driverTypeRep.RepositoryInterface
	Gearbox    gearboxRep.RepositoryInterface
	Vehicle    vehicleRep.RepositoryInterface
}

func setupRepositories(Storage storage.Storage) (*Repositories, error) {

	switch Storage.GetName() {
	case consts.MongoDB:
		// type assertion - get object from interface
		mongoStorage := Storage.(*mongo.MongoStorage)
		return &Repositories{
			brandRep.NewMongoBrandRepo(mongoStorage.DB),
			bodyStyleRep.NewMongoBodyStyleRepo(mongoStorage.DB),
			categoryRep.NewMongoCategoryRepo(mongoStorage.DB),
			driverTypeRep.NewMongoDriverTypeRepo(mongoStorage.DB),
			gearboxRep.NewMongoGearboxRepo(mongoStorage.DB),
			vehicleRep.NewMongoVehicleRepo(mongoStorage.DB),
		}, nil

	case consts.SqLite:
		// type assertion - get object from interface
		sqliteStorage := Storage.(*sqlite.SqliteStorage)
		return &Repositories{
			brandRep.NewSqliteBrandRepo(sqliteStorage.DB),
			bodyStyleRep.NewSqliteBodyStyleRepo(sqliteStorage.DB),
			categoryRep.NewSqliteCategoryRepo(sqliteStorage.DB),
			driverTypeRep.NewSqliteDriverTypeRepo(sqliteStorage.DB),
			gearboxRep.NewSqliteGearboxRepo(sqliteStorage.DB),
			vehicleRep.NewSqliteVehicleRepo(sqliteStorage.DB),
		}, nil

	default:
		return nil, fmt.Errorf("failed to setup %s repositories", Storage.GetName())
	}

}
