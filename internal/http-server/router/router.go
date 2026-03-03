package router

import (
	"log/slog"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/router/body_style"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/router/brand"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/router/category"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/router/driver_type"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/router/gearbox"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/router/vehicle"
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

	body_style.SetupBodyStyleRoutes(router, log, bodyStyleRepo, cfg)
	driver_type.SetupDriverTypeRoutes(router, log, driverTypeRepo, cfg)
	brand.SetupBrandRoutes(router, log, brandRepo, cfg)
	category.SetupCategoryRoutes(router, log, categoryRepo, cfg)
	gearbox.SetupGearboxRoutes(router, log, gearboxRepo, cfg)
	vehicle.SetupVehiclesRoutes(router, log, vehicleRepo)
	return router
}
