package category

import (
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	category "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/category"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type GetAllResponse struct {
	Response   resp.Response
	Categories []models.Category
}

func GetAll(log *slog.Logger, repository category.RepositoryInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.category.GetAll"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		log.Info("getting categories")

		categories, err := repository.GetAll(r.Context())
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
