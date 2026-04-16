package brand

import (
	"log/slog"

	brandHandler "github.com/NakonechniyVitalii/GoVehicleApi/internal/http-server/handlers/brand"
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/services"
	"github.com/go-chi/chi/v5"
)

func SetupBrandRoutes(
	router chi.Router,
	log *slog.Logger,
	services services.Container,
) {
	router.Route("/brand", func(r chi.Router) {
		r.Post("/", brandHandler.New(log, services.Brand))
		r.Delete("/{id}", brandHandler.Delete(log, services.Brand))
		r.Get("/{id}", brandHandler.Get(log, services.Brand))
		r.Get("/all", brandHandler.GetAll(log, services.Brand))
		r.Put("/{id}", brandHandler.Update(log, services.Brand))
		r.Put("/refresh", brandHandler.Refresh(log, services.Brand))
	})
}
