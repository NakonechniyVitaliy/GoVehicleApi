package brand

import (
	"log/slog"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	brandHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/brand"
	brandRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/brand"
	"github.com/go-chi/chi/v5"
)

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
