package driver_type

import (
	"log/slog"

	handler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/driver_type"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/services"
	"github.com/go-chi/chi/v5"
)

func SetupDriverTypeRoutes(
	router chi.Router,
	log *slog.Logger,
	services services.Container,
) {
	router.Route("/driver-type", func(r chi.Router) {
		r.Get("/all", handler.GetAll(log, services.DriverType))
		r.Put("/refresh", handler.Refresh(log, services.DriverType))
	})
}
