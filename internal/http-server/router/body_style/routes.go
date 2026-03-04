package body_style

import (
	"log/slog"

	handler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/body_style"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/body_style"

	"github.com/go-chi/chi/v5"
)

func SetupBodyStyleRoutes(
	router chi.Router,
	log *slog.Logger,
	service *service.Service,
) {
	router.Route("/body-style", func(r chi.Router) {
		r.Post("/", handler.New(log, service))
		r.Delete("/{id}", handler.Delete(log, service))
		r.Get("/{id}", handler.Get(log, *service))
		r.Get("/all", handler.GetAll(log, *service))
		r.Put("/{id}", handler.Update(log, *service))
		r.Put("/refresh", handler.Refresh(log, *service))
	})
}
