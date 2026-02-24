package category

import (
	"log/slog"
	"net/http"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	CategoryRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/category"
	CategoryService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/service/category"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func Refresh(log *slog.Logger, repository CategoryRepo.RepositoryInterface, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.category.refresh"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		err := CategoryService.RefreshCategories(r.Context(), cfg, repository)
		if err != nil {
			log.Error("failed to update categories", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to update category"))
			return
		}

		render.JSON(w, r, resp.OK())
	}
}
