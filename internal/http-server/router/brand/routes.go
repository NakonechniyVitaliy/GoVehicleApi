package brand

import (
	"log/slog"

	brandHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/brand"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/brand"
	"github.com/go-chi/chi/v5"
)

func SetupBrandRoutes(
	router chi.Router,
	log *slog.Logger,
	service *service.Service,
) {
	router.Route("/brand", func(r chi.Router) {
		r.Post("/", brandHandler.New(log, service))
		r.Delete("/{id}", brandHandler.Delete(log, service))
		r.Get("/{id}", brandHandler.Get(log, service))
		r.Get("/all", brandHandler.GetAll(log, service))
		r.Put("/{id}", brandHandler.Update(log, service))
		r.Put("/refresh", brandHandler.Refresh(log, service))
	})
}
