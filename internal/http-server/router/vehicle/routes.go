package vehicle

import (
	"log/slog"

	vehicleHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/vehicle"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/services"
	"github.com/go-chi/chi/v5"
)

func SetupVehiclesRoutes(
	router chi.Router,
	log *slog.Logger,
	services services.Container,
) {

	router.Route("/vehicle", func(r chi.Router) {
		r.Post("/", vehicleHandler.New(log, services.Vehicle))
		r.Delete("/{id}", vehicleHandler.Delete(log, services.Vehicle))
		r.Get("/{id}", vehicleHandler.Get(log, services.Vehicle))
		r.Put("/{id}", vehicleHandler.Update(log, services.Vehicle))
		r.Get("/all", vehicleHandler.GetAll(log, services.Vehicle))
		r.Get("/expanded/{id}", vehicleHandler.GetExpanded(log, services.Vehicle))
	})
}
