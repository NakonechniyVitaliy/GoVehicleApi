package router

import (
	"log/slog"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	bodyStyleHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/body_style"
	brandHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/brand"
	categoryHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/category"
	driverTypeHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/driver_type"
	gearboxHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/gearbox"
	vehicleHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/vehicle"
	bodyStyleRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/body_style"
	brandRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/brand"
	categoryRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/category"
	driverTypeRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/driver_type"
	gearboxRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/gearbox"
	vehicleRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(
	log *slog.Logger,
	brandRepo brandRepo.RepositoryInterface,
	bodyStyleRepo bodyStyleRepo.RepositoryInterface,
	categoryRepo categoryRepo.RepositoryInterface,
	driverTypeRepo driverTypeRepo.RepositoryInterface,
	gearboxRepo gearboxRepo.RepositoryInterface,
	vehicleRepo vehicleRepo.RepositoryInterface,
	cfg *config.Config,
) chi.Router {

	router := chi.NewRouter()

	router.Use(middleware.RequestID) // Присваивание ID каждому запросу
	router.Use(middleware.Logger)    // Логирование
	router.Use(middleware.Recoverer) // Востановление после критикал ошибки
	router.Use(middleware.URLFormat) // Удобное получение параметров из сслыки

	SetupBodyStyleRoutes(router, log, bodyStyleRepo, cfg)
	SetupDriverTypeRoutes(router, log, driverTypeRepo, cfg)
	SetupBrandRoutes(router, log, brandRepo, cfg)
	SetupCategoryRoutes(router, log, categoryRepo, cfg)
	SetupGearboxRoutes(router, log, gearboxRepo, cfg)
	SetupVehiclesRoutes(router, log, vehicleRepo)
	return router
}

func SetupBodyStyleRoutes(
	router chi.Router,
	log *slog.Logger,
	bodyStyleRepo bodyStyleRepo.RepositoryInterface,
	cfg *config.Config,
) {
	router.Route("/body-style", func(r chi.Router) {
		r.Post("/", bodyStyleHandler.New(log, bodyStyleRepo))
		r.Delete("/{id}", bodyStyleHandler.Delete(log, bodyStyleRepo))
		r.Get("/{id}", bodyStyleHandler.Get(log, bodyStyleRepo))
		r.Get("/all", bodyStyleHandler.GetAll(log, bodyStyleRepo))
		r.Put("/", bodyStyleHandler.Update(log, bodyStyleRepo))
		r.Put("/refresh", bodyStyleHandler.Refresh(log, bodyStyleRepo, cfg))
	})
}

func SetupBrandRoutes(
	router chi.Router,
	log *slog.Logger,
	brandRepo brandRepo.RepositoryInterface,
	cfg *config.Config,
) {
	router.Route("/brand", func(r chi.Router) {
		r.Post("/", brandHandler.New(log, brandRepo))
		r.Delete("/{id}", brandHandler.Delete(log, brandRepo))
		r.Get("/{id}", brandHandler.Get(log, brandRepo))
		r.Get("/all", brandHandler.GetAll(log, brandRepo))
		r.Put("/{id}", brandHandler.Update(log, brandRepo))
		r.Put("/refresh", brandHandler.Refresh(log, brandRepo, cfg))
	})
}

func SetupCategoryRoutes(
	router chi.Router,
	log *slog.Logger,
	categoryRepo categoryRepo.RepositoryInterface,
	cfg *config.Config,
) {
	router.Route("/vehicle-category", func(r chi.Router) {
		r.Get("/all", categoryHandler.GetAll(log, categoryRepo))
		r.Put("/refresh", categoryHandler.Refresh(log, categoryRepo, cfg))
	})
}

func SetupDriverTypeRoutes(
	router chi.Router,
	log *slog.Logger,
	driverTypeRepo driverTypeRepo.RepositoryInterface,
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
	gearboxRepo gearboxRepo.RepositoryInterface,
	cfg *config.Config,
) {
	router.Route("/gearbox", func(r chi.Router) {
		r.Get("/all", gearboxHandler.GetAll(log, gearboxRepo))
		r.Put("/refresh", gearboxHandler.Refresh(log, gearboxRepo, cfg))
	})
}

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
