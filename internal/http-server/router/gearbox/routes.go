package gearbox

import (
	"log/slog"

	gearboxHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/gearbox"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/gearbox"
	"github.com/go-chi/chi/v5"
)

func SetupGearboxRoutes(
	router chi.Router,
	log *slog.Logger,
	service *service.Service,
) {
	router.Route("/gearbox", func(r chi.Router) {
		r.Get("/all", gearboxHandler.GetAll(log, service))
		r.Put("/refresh", gearboxHandler.Refresh(log, service))
	})
}
