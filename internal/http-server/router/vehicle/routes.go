package vehicle

import (
	"log/slog"

	vehicleHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/vehicle"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/vehicle"
	"github.com/go-chi/chi/v5"
)

func SetupVehiclesRoutes(
	router chi.Router,
	log *slog.Logger,
	service *service.Service,

) {
	router.Route("/vehicle", func(r chi.Router) {
		r.Post("/", vehicleHandler.New(log, service))
		r.Delete("/{id}", vehicleHandler.Delete(log, service))
		r.Get("/{id}", vehicleHandler.Get(log, service))
		r.Put("/{id}", vehicleHandler.Update(log, service))
		r.Get("/all", vehicleHandler.GetAll(log, service))
		r.Get("/expanded/{id}", vehicleHandler.GetExpanded(log, service))
	})
}
