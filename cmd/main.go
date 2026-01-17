package main

import (
	"encoding/json"
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage/mongo"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/brand/save"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage/sqlite"
	"github.com/NakonechniyVitaliy/GoVehicleApi/requests"
	//"github.com/go-chi/chi/v5"
	//"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"os"
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
	_ = Storage

	// REFRESH BRANDS SQLITE
	//_ = Storage
	//brands, _ := requests.GetBrands(cfg.AutoriaKey)
	//err = Storage.RefreshBrands(brands)
	//if err != nil {
	//	fmt.Println("Error refreshing brands", err)
	//}

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
