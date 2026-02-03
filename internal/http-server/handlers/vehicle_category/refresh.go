package vehicle_category

import (
	"log/slog"
	"net/http"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	vehicleCategoryRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle_category"
	vehicleCategoryService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/service/vehicle_category"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func Refresh(log *slog.Logger, repository vehicleCategoryRepo.Repository, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.vehicleCategory.update.Refresh"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		err := vehicleCategoryService.RefreshVehicleCategories(r.Context(), cfg, repository)
		if err != nil {
			log.Error("failed to update vehicleCategorys", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to update vehicleCategory"))
			return
		}

		render.JSON(w, r, resp.OK())
	}
}
