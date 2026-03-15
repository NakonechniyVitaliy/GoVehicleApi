package gearbox

import (
	"log/slog"

	gearboxHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/gearbox"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/services"
	"github.com/go-chi/chi/v5"
)

func SetupGearboxRoutes(
	router chi.Router,
	log *slog.Logger,
	services services.Container,
) {
	router.Route("/gearbox", func(r chi.Router) {
		r.Get("/all", gearboxHandler.GetAll(log, services.Gearbox))
		r.Put("/refresh", gearboxHandler.Refresh(log, services.Gearbox))
	})
}
