package get

import (
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/brand"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type ResponseGetAll struct {
	Response resp.Response
	Brands   []models.Brand
}

func GetAll(log *slog.Logger, repository brand.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.brand.get.Get"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		log.Info("getting brands")

		brands, err := repository.GetAll(r.Context())
		if err != nil {
			log.Error("failed to get brand", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get brand"))
			return
		}

		render.JSON(w, r, ResponseGetAll{
			Response: resp.OK(),
			Brands:   brands,
		})
	}
}
