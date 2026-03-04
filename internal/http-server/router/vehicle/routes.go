package vehicle

import (
	"log/slog"

	vehicleHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/vehicle"
	bodyStyleRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/body_style"
	brandRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/brand"
	categoryRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/category"
	driverTypeRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/driver_type"
	gearboxRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/gearbox"
	vehicleRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle"
	"github.com/go-chi/chi/v5"
)

func SetupVehiclesRoutes(
	router chi.Router,
	log *slog.Logger,
	vehicleRepo vehicleRepo.RepositoryInterface,
	brandRepo brandRepo.RepositoryInterface,
	bodyStyleRepo bodyStyleRepo.RepositoryInterface,
	categoryRepo categoryRepo.RepositoryInterface,
	driverTypeRepo driverTypeRepo.RepositoryInterface,
	gearboxRepo gearboxRepo.RepositoryInterface,
) {
	router.Route("/vehicle", func(r chi.Router) {
		r.Post("/", vehicleHandler.New(log, vehicleRepo))
		r.Delete("/{id}", vehicleHandler.Delete(log, vehicleRepo))
		r.Get("/{id}", vehicleHandler.Get(log, vehicleRepo))
		r.Put("/{id}", vehicleHandler.Update(log, vehicleRepo))
		r.Get("/all", vehicleHandler.GetAll(log, vehicleRepo))
		r.Get("/expanded/{id}", vehicleHandler.GetExpanded(
			log,
			vehicleRepo,
			brandRepo,
			bodyStyleRepo,
			categoryRepo,
			driverTypeRepo,
			gearboxRepo,
		))
	})
}
