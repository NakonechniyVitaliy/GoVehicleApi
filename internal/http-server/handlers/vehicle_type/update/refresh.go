package update

import (
	"log/slog"
	"net/http"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	vehicleTypeRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle_type"
	vehicleTypeService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/service/vehicle_type"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func Refresh(log *slog.Logger, repository vehicleTypeRepo.Repository, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.vehicleCategory.update.Refresh"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		err := vehicleTypeService.RefreshVehicleTypes(r.Context(), cfg, repository)
		if err != nil {
			log.Error("failed to update vehicleCategorys", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to update vehicleCategory"))
			return
		}

		render.JSON(w, r, Response{
			Response: resp.OK(),
		})
	}
}
