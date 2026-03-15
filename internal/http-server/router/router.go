package router

import (
	"log/slog"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/router/body_style"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/router/brand"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/router/category"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/router/driver_type"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/router/gearbox"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/router/user"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/router/vehicle"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(
	log *slog.Logger,
	services services.Container,
) chi.Router {

	router := chi.NewRouter()

	router.Use(middleware.RequestID) // Присваивание ID каждому запросу
	router.Use(middleware.Logger)    // Логирование
	router.Use(middleware.Recoverer) // Востановление после критикал ошибки
	router.Use(middleware.URLFormat) // Удобное получение параметров из сслыки

	body_style.SetupBodyStyleRoutes(router, log, services)
	driver_type.SetupDriverTypeRoutes(router, log, services)
	brand.SetupBrandRoutes(router, log, services)
	category.SetupCategoryRoutes(router, log, services)
	gearbox.SetupGearboxRoutes(router, log, services)
	vehicle.SetupVehiclesRoutes(router, log, services)
	user.SetupUserRoutes(router, log, services)

	return router
}
