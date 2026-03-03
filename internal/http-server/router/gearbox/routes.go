package gearbox

import (
	"log/slog"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	gearboxHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/gearbox"
	gearboxRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/gearbox"
	"github.com/go-chi/chi/v5"
)

func SetupGearboxRoutes(
	router chi.Router,
	log *slog.Logger,
	gearboxRepo gearboxRepo.RepositoryInterface,
	cfg *config.Config,
) {
	router.Route("/gearbox", func(r chi.Router) {
		r.Get("/all", gearboxHandler.GetAll(log, gearboxRepo))
		r.Put("/refresh", gearboxHandler.Refresh(log, gearboxRepo, cfg))
	})
}
