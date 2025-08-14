package main

import (
	"encoding/json"
	"fmt"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage/sqlite"
	"github.com/NakonechniyVitaliy/GoVehicleApi/requests"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
	envDev   = "dev"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log.Info("starting server", slog.String("env", cfg.Env))
	log.Debug("Debug messages are enabled")

	Storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("Error creating storage", err)
		os.Exit(1)
	}

	_ = Storage

	//err = http.ListenAndServe("localhost:8080", nil)
	//if err != nil {
	//	log.Error("Error starting server", err)
	//} else {
	//	log.Info("Server started")
	//}

	brands, _ := requests.GetBrands(cfg.AutoriaKey)
	err = Storage.RefreshBrands(brands)
	if err != nil {
		fmt.Println("Error refreshing brands", err)
	} else {
		log.Info("brands refreshed")
	}

	brandsFromDB, err := Storage.GetBrands()
	if err != nil {
		log.Error("Error getting brands from storage", err)
	}

	data, _ := json.MarshalIndent(brandsFromDB, "", "  ")
	fmt.Println(string(data))

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
