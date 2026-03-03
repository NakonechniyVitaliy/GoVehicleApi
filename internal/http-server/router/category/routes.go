package category

import (
	"log/slog"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	categoryHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/category"
	categoryRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/category"
	"github.com/go-chi/chi/v5"
)

func SetupCategoryRoutes(
	router chi.Router,
	log *slog.Logger,
	categoryRepo categoryRepo.RepositoryInterface,
	cfg *config.Config,
) {
	router.Route("/category", func(r chi.Router) {
		r.Get("/all", categoryHandler.GetAll(log, categoryRepo))
		r.Put("/refresh", categoryHandler.Refresh(log, categoryRepo, cfg))
	})
}
