package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	deleteBrandHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/brand/delete"
	getBrandHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/brand/get"
	saveBrandHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/brand/save"
	updateBrandHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/brand/update"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/brand"
	vehicleType "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle_type"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage/mongo"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	deleteVehicleTypeHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/vehicle_type/delete"
	getVehicleTypeHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/vehicle_type/get"
	saveVehicleTypeHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/vehicle_type/save"
	updateVehicleTypeHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/vehicle_type/update"
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

	brandRepo, vehicleTypeRepo, err := setupRepositories(Storage)
	if err != nil {
		log.Error("failed to setup repositories", slog.Any("err", err))
	}

	log.Info("Database successfully enabled", slog.String("database", cfg.DataBase))

	router := setupRouter()

	router.Route("/brand", func(r chi.Router) {
		r.Post("/", saveBrandHandler.New(log, brandRepo))
		r.Delete("/{id}", deleteBrandHandler.Delete(log, brandRepo))
		r.Get("/{id}", getBrandHandler.Get(log, brandRepo))
		r.Get("/all", getBrandHandler.GetAll(log, brandRepo))
		r.Put("/", updateBrandHandler.Update(log, brandRepo))
	})
	router.Route("/vehicle-type", func(r chi.Router) {
		r.Post("/", saveVehicleTypeHandler.New(log, vehicleTypeRepo))
		r.Delete("/{id}", deleteVehicleTypeHandler.Delete(log, vehicleTypeRepo))
		r.Get("/{id}", getVehicleTypeHandler.Get(log, vehicleTypeRepo))
		r.Get("/all", getVehicleTypeHandler.GetAll(log, vehicleTypeRepo))
		r.Put("/", updateVehicleTypeHandler.Update(log, vehicleTypeRepo))
	})

	log.Info("starting server on", slog.String("adreess", cfg.Address))

	server := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
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

func setupRouter() chi.Router {
	router := chi.NewRouter()

	// Присваивание ID каждому запросу
	router.Use(middleware.RequestID)
	// Логирование
	router.Use(middleware.Logger)
	// Востановление после критикал ошибки
	router.Use(middleware.Recoverer)
	// Удобное получение параметров из сслыки
	router.Use(middleware.URLFormat)

	return router
}

// REFRESH BRANDS SQLITE
//_ = Storage
//brands, _ := requests.GetBrands(cfg.AutoriaKey)
//err = Storage.RefreshBrands(brands)
//if err != nil {
//	fmt.Println("Error refreshing brands", err)
//}
