package driver_type

import (
	"log/slog"

	handler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/driver_type"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/driver_type"
	"github.com/go-chi/chi/v5"
)

func SetupDriverTypeRoutes(
	router chi.Router,
	log *slog.Logger,
	service *service.Service,
) {
	router.Route("/driver-type", func(r chi.Router) {
		r.Get("/all", handler.GetAll(log, service))
		r.Put("/refresh", handler.Refresh(log, service))
	})
}
