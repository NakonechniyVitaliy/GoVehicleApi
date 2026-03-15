package category

import (
	"log/slog"

	categoryHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/category"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/services"
	"github.com/go-chi/chi/v5"
)

func SetupCategoryRoutes(
	router chi.Router,
	log *slog.Logger,
	services services.Container,
) {
	router.Route("/category", func(r chi.Router) {
		r.Get("/all", categoryHandler.GetAll(log, services.Category))
		r.Put("/refresh", categoryHandler.Refresh(log, services.Category))
	})
}
