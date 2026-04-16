package logger

import (
	"log/slog"
	"os"

	consts "github.com/NakonechniyVitalii/GoVehicleApi/internal/constants"
)

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case consts.EnvLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case consts.EnvDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case consts.EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
