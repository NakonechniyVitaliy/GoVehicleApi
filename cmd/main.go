package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/router"
	bodyStyleRep "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/body_style"
	brandRep "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/brand"
	driverTypeRep "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/driver_type"
	gearboxRep "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/gearbox"
	vehicleCategoryRep "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle_category"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage/mongo"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envProd  = "prod"
	envDev   = "dev"
	mongoDB  = "mongo"
	Sqlite   = "sqlite"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log.Info("starting server", slog.String("env", cfg.Env))
	log.Debug("Debug messages are enabled")

	Storage, err := setupDataBase(cfg.DataBase, cfg.StoragePath)
	if err != nil {
		log.Error("failed to setup database", slog.Any("err", err))
		os.Exit(1)
	}
	log.Info("Database successfully enabled", slog.String("database", cfg.DataBase))

	brandRepo,
		bodyStyleRepo,
		vehicleCategoryRepo,
		driverTypeRepo,
		gearboxRepo,
		err := setupRepositories(Storage)

	if err != nil {
		log.Error("failed to setup repositories", slog.Any("err", err))
	}
	log.Info("repositories successfully setup", slog.String("database", cfg.DataBase))

	appRouter := router.SetupRouter(log, brandRepo, bodyStyleRepo, vehicleCategoryRepo, driverTypeRepo, gearboxRepo, cfg)

	log.Info("starting server on", slog.String("address", cfg.Address))
	server := &http.Server{
		Addr:         cfg.Address,
		Handler:      appRouter,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Error("server stoped")

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}

func setupRepositories(Storage storage.Storage) (
	brandRep.Repository,
	bodyStyleRep.Repository,
	vehicleCategoryRep.Repository,
	driverTypeRep.Repository,
	gearboxRep.Repository,
	error,
) {
	switch Storage.GetName() {

	case mongoDB:
		// type assertion - get object from interface
		mongoStorage := Storage.(*mongo.MongoStorage)
		brandRepo := brandRep.NewMongo(mongoStorage.DB)
		bodyStyleRepo := bodyStyleRep.NewMongo(mongoStorage.DB)
		vehicleCategoryRepo := vehicleCategoryRep.NewMongo(mongoStorage.DB)
		driverTypeRepo := driverTypeRep.NewMongo(mongoStorage.DB)
		gearboxRepo := gearboxRep.NewMongo(mongoStorage.DB)

		return brandRepo, bodyStyleRepo, vehicleCategoryRepo, driverTypeRepo, gearboxRepo, nil

	case Sqlite:
		sqliteStorage := Storage.(*sqlite.SqliteStorage)
		brandRepo := brandRep.NewSqlite(sqliteStorage.DB)
		bodyStyleRepo := bodyStyleRep.NewSqlite(sqliteStorage.DB)
		vehicleCategoryRepo := vehicleCategoryRep.NewSqlite(sqliteStorage.DB)
		driverTypeRepo := driverTypeRep.NewSqlite(sqliteStorage.DB)
		gearboxRepo := gearboxRep.NewSqlite(sqliteStorage.DB)

		return brandRepo, bodyStyleRepo, vehicleCategoryRepo, driverTypeRepo, gearboxRepo, nil

	default:
		return nil, nil, nil, nil, nil, nil
	}

}

func setupDataBase(db string, storagePath string) (storage.Storage, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	switch db {
	case mongoDB:
		return mongo.New(ctx)

	case Sqlite:
		return sqlite.New(storagePath)

	default:
		return nil, fmt.Errorf("unknown database type: %s", db)
	}
}
