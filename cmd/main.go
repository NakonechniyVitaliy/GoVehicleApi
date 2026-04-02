package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/app"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)
	log.Info("starting server", slog.String("env", cfg.Env))
	log.Debug("Debug messages are enabled")

	application, err := app.New(log, cfg)
	if err != nil {
		log.Error("failed to initialize app", slog.Any("err", err))
		os.Exit(1)
	}

	log.Info("starting server on", slog.String("address", cfg.HTTPServer.Address))
	server := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      application.Router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("failed to start server", slog.Any("err", err))
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error("server forced to shutdown", slog.Any("err", err))
		os.Exit(1)
	}

	log.Info("server stopped")
}
