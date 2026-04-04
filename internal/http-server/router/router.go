package router

import (
	"log/slog"

	myMiddleware "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/middleware"
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
	httpSwagger "github.com/swaggo/http-swagger"
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

	// Swagger UI (публічний)
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	// Публичные роуты (без аутентификации)
	user.SetupUserRoutes(router, log, services)

	// Защищённые роуты (требуют JWT токен)
	router.Group(func(r chi.Router) {
		r.Use(myMiddleware.JWTAuth(log, services.JWT))

		body_style.SetupBodyStyleRoutes(r, log, services)
		driver_type.SetupDriverTypeRoutes(r, log, services)
		brand.SetupBrandRoutes(r, log, services)
		category.SetupCategoryRoutes(r, log, services)
		gearbox.SetupGearboxRoutes(r, log, services)
		vehicle.SetupVehiclesRoutes(r, log, services)
	})

	return router
}
