package body_style

import (
	"log/slog"

	handler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/body_style"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/services"

	"github.com/go-chi/chi/v5"
)

func SetupBodyStyleRoutes(
	router chi.Router,
	log *slog.Logger,
	services services.Container,
) {
	router.Route("/body-style", func(r chi.Router) {
		r.Post("/", handler.New(log, services.BodyStyle))
		r.Delete("/{id}", handler.Delete(log, services.BodyStyle))
		r.Get("/{id}", handler.Get(log, *services.BodyStyle))
		r.Get("/all", handler.GetAll(log, *services.BodyStyle))
		r.Put("/{id}", handler.Update(log, *services.BodyStyle))
		r.Put("/refresh", handler.Refresh(log, *services.BodyStyle))
	})
}
