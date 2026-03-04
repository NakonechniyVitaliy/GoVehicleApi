package category

import (
	"log/slog"

	categoryHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/category"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/category"
	"github.com/go-chi/chi/v5"
)

func SetupCategoryRoutes(
	router chi.Router,
	log *slog.Logger,
	service *service.Service,
) {
	router.Route("/category", func(r chi.Router) {
		r.Get("/all", categoryHandler.GetAll(log, service))
		r.Put("/refresh", categoryHandler.Refresh(log, service))
	})
}
