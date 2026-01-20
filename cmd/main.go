package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	deleteHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/brand/delete"
	getHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/brand/get"
	saveHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/brand/save"
	updateHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/brand/update"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage/mongo"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	envLocal = "local"
	envProd  = "prod"
	envDev   = "dev"
	mongoDB  = "mongo"
	SQLite   = "sqlite"
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

	router := setupRouter()

	router.Route("/brand", func(r chi.Router) {
		r.Post("/", saveHandler.New(log, Storage))
		r.Delete("/{id}", deleteHandler.Delete(log, Storage))
		r.Get("/{id}", getHandler.Get(log, Storage))
		r.Put("/", updateHandler.Update(log, Storage))
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

func setupDataBase(db string, storagePath string) (storage.Storage, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	switch db {
	case mongoDB:
		return mongo.New(ctx)

	case SQLite:
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
