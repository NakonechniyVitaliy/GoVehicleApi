package vehicle

import (
	"log/slog"

	vehicleHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/vehicle"
	vehicleRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle"
	"github.com/go-chi/chi/v5"
)

func SetupVehiclesRoutes(
	router chi.Router,
	log *slog.Logger,
	vehicleRepo vehicleRepo.RepositoryInterface,
) {
	router.Route("/vehicle", func(r chi.Router) {
		r.Post("/", vehicleHandler.New(log, vehicleRepo))
		r.Delete("/{id}", vehicleHandler.Delete(log, vehicleRepo))
		r.Get("/{id}", vehicleHandler.Get(log, vehicleRepo))
		r.Put("/{id}", vehicleHandler.Update(log, vehicleRepo))
		r.Get("/all", vehicleHandler.GetAll(log, vehicleRepo))
	})
}
