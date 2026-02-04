package router

import (
	"log/slog"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	brandHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/brand"
	driverTypeHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/driver_type"
	gearboxHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/gearbox"
	vehicleCategoryHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/vehicle_category"
	vehicleTypeHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/vehicle_type"
	brandRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/brand"
	driverTypeRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/driver_type"
	gearboxRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/gearbox"
	vehicleCategoryRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle_category"
	vehicleTypeRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle_type"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle_type"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(
	log *slog.Logger,
	brandRepo brandRepo.Repository,
	vehicleTypeRepo vehicleType.Repository,
	vehicleCategoryRepo vehicleCategoryRepo.Repository,
	driverTypeRepo driverTypeRepo.Repository,
	gearboxRepo gearboxRepo.Repository,
	cfg *config.Config,
) chi.Router {

	router := chi.NewRouter()

	router.Use(middleware.RequestID) // Присваивание ID каждому запросу
	router.Use(middleware.Logger)    // Логирование
	router.Use(middleware.Recoverer) // Востановление после критикал ошибки
	router.Use(middleware.URLFormat) // Удобное получение параметров из сслыки

	SetupVehicleTypeRoutes(router, log, vehicleTypeRepo, cfg)
	SetupDriverTypeRoutes(router, log, driverTypeRepo, cfg)
	SetupBrandRoutes(router, log, brandRepo, cfg)
	SetupVehicleCategoryRoutes(router, log, vehicleCategoryRepo, cfg)
	SetupGearboxRoutes(router, log, gearboxRepo, cfg)
	return router
}

func SetupVehicleTypeRoutes(
	router chi.Router,
	log *slog.Logger,
	vehicleTypeRepo vehicleTypeRepo.Repository,
	cfg *config.Config,
) {
	router.Route("/vehicle-type", func(r chi.Router) {
		r.Post("/", vehicleTypeHandler.New(log, vehicleTypeRepo))
		r.Delete("/{id}", vehicleTypeHandler.Delete(log, vehicleTypeRepo))
		r.Get("/{id}", vehicleTypeHandler.Get(log, vehicleTypeRepo))
		r.Get("/all", vehicleTypeHandler.GetAll(log, vehicleTypeRepo))
		r.Put("/", vehicleTypeHandler.Update(log, vehicleTypeRepo))
		r.Put("/refresh", vehicleTypeHandler.Refresh(log, vehicleTypeRepo, cfg))
	})
}

func SetupBrandRoutes(
	router chi.Router,
	log *slog.Logger,
	brandRepo brandRepo.Repository,
	cfg *config.Config,
) {
	router.Route("/brand", func(r chi.Router) {
		r.Post("/", brandHandler.New(log, brandRepo))
		r.Delete("/{id}", brandHandler.Delete(log, brandRepo))
		r.Get("/{id}", brandHandler.Get(log, brandRepo))
		r.Get("/all", brandHandler.GetAll(log, brandRepo))
		r.Put("/", brandHandler.Update(log, brandRepo))
		r.Put("/refresh", brandHandler.Refresh(log, brandRepo, cfg))
	})
}

func SetupVehicleCategoryRoutes(
	router chi.Router,
	log *slog.Logger,
	vehicleCategoryRepo vehicleCategoryRepo.Repository,
	cfg *config.Config,
) {
	router.Route("/vehicle-category", func(r chi.Router) {
		r.Get("/all", vehicleCategoryHandler.GetAll(log, vehicleCategoryRepo))
		r.Put("/refresh", vehicleCategoryHandler.Refresh(log, vehicleCategoryRepo, cfg))
	})
}

func SetupDriverTypeRoutes(
	router chi.Router,
	log *slog.Logger,
	driverTypeRepo driverTypeRepo.Repository,
	cfg *config.Config,
) {
	router.Route("/driver-type", func(r chi.Router) {
		r.Get("/all", driverTypeHandler.GetAll(log, driverTypeRepo))
		r.Put("/refresh", driverTypeHandler.Refresh(log, driverTypeRepo, cfg))
	})
}

func SetupGearboxRoutes(
	router chi.Router,
	log *slog.Logger,
	gearboxRepo gearboxRepo.Repository,
	cfg *config.Config,
) {
	router.Route("/gearbox", func(r chi.Router) {
		r.Get("/all", gearboxHandler.GetAll(log, gearboxRepo))
		r.Put("/refresh", gearboxHandler.Refresh(log, gearboxRepo, cfg))
	})
}
