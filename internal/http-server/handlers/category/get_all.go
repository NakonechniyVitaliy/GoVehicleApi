package category

import (
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/category"
	"github.com/go-chi/render"
)

type GetAllResponse struct {
	Response   resp.Response
	Categories []models.Category
}

func GetAll(log *slog.Logger, service *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.category.get_all"

		log = log.With(slog.String("op", op))

		log.Info("getting categories")

		categories, err := service.GetAll(r.Context())
		if err != nil {
			log.Error("failed to get categories", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get categories"))
			return
		}

		render.JSON(w, r, GetAllResponse{
			Response:   resp.OK(),
			Categories: categories,
		})
	}
}
