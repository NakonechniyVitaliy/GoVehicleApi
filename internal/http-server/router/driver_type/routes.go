package driver_type

import (
	"log/slog"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	driverTypeHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/driver_type"
	driverTypeRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/driver_type"
	"github.com/go-chi/chi/v5"
)

func SetupDriverTypeRoutes(
	router chi.Router,
	log *slog.Logger,
	driverTypeRepo driverTypeRepo.RepositoryInterface,
	cfg *config.Config,
) {
	router.Route("/driver-type", func(r chi.Router) {
		r.Get("/all", driverTypeHandler.GetAll(log, driverTypeRepo))
		r.Put("/refresh", driverTypeHandler.Refresh(log, driverTypeRepo, cfg))
	})
}
