package router

import (
	"log/slog"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	deleteBrandHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/brand/delete"
	getBrandHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/brand/get"
	saveBrandHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/brand/save"
	updateBrandHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/brand/update"
	getVehicleCategoryHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/vehicle_category/get"
	updateVehicleCategoryHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/vehicle_category/update"
	deleteVehicleTypeHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/vehicle_type/delete"
	getVehicleTypeHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/vehicle_type/get"
	saveVehicleTypeHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/vehicle_type/save"
	updateVehicleTypeHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/vehicle_type/update"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/brand"
	vehicleCategory "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle_category"
	vehicleType "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle_type"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(
	log *slog.Logger,
	brandRepo brand.Repository,
	vehicleTypeRepo vehicleType.Repository,
	vehicleCategoryRepo vehicleCategory.Repository,
	cfg *config.Config) chi.Router {

	router := chi.NewRouter()

	router.Use(middleware.RequestID) // Присваивание ID каждому запросу
	router.Use(middleware.Logger)    // Логирование
	router.Use(middleware.Recoverer) // Востановление после критикал ошибки
	router.Use(middleware.URLFormat) // Удобное получение параметров из сслыки

	SetupVehicleTypeRoutes(router, log, vehicleTypeRepo, cfg)
	SetupBrandRoutes(router, log, brandRepo, cfg)
	SetupVehicleCategoryRoutes(router, log, vehicleCategoryRepo, cfg)

	return router
}

func SetupVehicleTypeRoutes(
	router chi.Router,
	log *slog.Logger,
	vehicleTypeRepo vehicleType.Repository,
	cfg *config.Config,
) {
	router.Route("/vehicle-type", func(r chi.Router) {
		r.Post("/", saveVehicleTypeHandler.New(log, vehicleTypeRepo))
		r.Delete("/{id}", deleteVehicleTypeHandler.Delete(log, vehicleTypeRepo))
		r.Get("/{id}", getVehicleTypeHandler.Get(log, vehicleTypeRepo))
		r.Get("/all", getVehicleTypeHandler.GetAll(log, vehicleTypeRepo))
		r.Put("/", updateVehicleTypeHandler.Update(log, vehicleTypeRepo))
		r.Put("/refresh", updateVehicleTypeHandler.Refresh(log, vehicleTypeRepo, cfg))
	})
}

func SetupBrandRoutes(
	router chi.Router,
	log *slog.Logger,
	brandRepo brand.Repository,
	cfg *config.Config,
) {
	router.Route("/brand", func(r chi.Router) {
		r.Post("/", saveBrandHandler.New(log, brandRepo))
		r.Delete("/{id}", deleteBrandHandler.Delete(log, brandRepo))
		r.Get("/{id}", getBrandHandler.Get(log, brandRepo))
		r.Get("/all", getBrandHandler.GetAll(log, brandRepo))
		r.Put("/", updateBrandHandler.Update(log, brandRepo))
		r.Put("/refresh", updateBrandHandler.Refresh(log, brandRepo, cfg))
	})
}

func SetupVehicleCategoryRoutes(
	router chi.Router,
	log *slog.Logger,
	vehicleCategoryRepo vehicleCategory.Repository,
	cfg *config.Config,
) {
	router.Route("/vehicle-category", func(r chi.Router) {
		r.Get("/all", getVehicleCategoryHandler.GetAll(log, vehicleCategoryRepo))
		r.Put("/refresh", updateVehicleCategoryHandler.Refresh(log, vehicleCategoryRepo, cfg))
	})
}
