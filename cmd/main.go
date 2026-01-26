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
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/brand"
	vehicleType "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle_type"
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

	brandRepo, vehicleTypeRepo, err := setupRepositories(Storage)
	if err != nil {
		log.Error("failed to setup repositories", slog.Any("err", err))
	}
	log.Info("repositories successfully setup", slog.String("database", cfg.DataBase))

	appRouter := router.SetupRouter(log, brandRepo, vehicleTypeRepo)

	log.Info("starting server on", slog.String("adreess", cfg.Address))

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

func setupRepositories(Storage storage.Storage) (brand.Repository, vehicleType.Repository, error) {

	switch Storage.GetName() {

	case mongoDB:
		// type assertion - get object from interface
		mongoStorage := Storage.(*mongo.MongoStorage)
		brandRepo := brand.NewMongo(mongoStorage.DB)
		vehicleTypeRepo := vehicleType.NewMongo(mongoStorage.DB)
		return brandRepo, vehicleTypeRepo, nil

	case Sqlite:
		sqliteStorage := Storage.(*sqlite.SqliteStorage)
		brandRepo := brand.NewSqlite(sqliteStorage.DB)
		vehicleTypeRepo := vehicleType.NewSqlite(sqliteStorage.DB)
		return brandRepo, vehicleTypeRepo, nil

	default:
		return nil, nil, nil
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

// REFRESH BRANDS SQLITE
//_ = Storage
//brands, _ := requests.GetBrands(cfg.AutoriaKey)
//err = Storage.RefreshBrands(brands)
//if err != nil {
//	fmt.Println("Error refreshing brands", err)
//}
