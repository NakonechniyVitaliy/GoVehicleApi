package body_style

import (
	"log/slog"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	bodyStyleHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/body_style"
	bodyStyleRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/body_style"
	"github.com/go-chi/chi/v5"
)

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
		r.Put("/{id}", bodyStyleHandler.Update(log, bodyStyleRepo))
		r.Put("/refresh", bodyStyleHandler.Refresh(log, bodyStyleRepo, cfg))
	})
}
