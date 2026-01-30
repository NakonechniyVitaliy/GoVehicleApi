package get

import (
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	vehicleCategory "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle_category"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Response struct {
	Response          resp.Response
	VehicleCategories []models.VehicleCategory
}

func GetAll(log *slog.Logger, repository vehicleCategory.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.vehicleCategory.get.GetAll"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		log.Info("getting vehicle categories")

		vehicleCategories, err := repository.GetAll(r.Context())
		if err != nil {
			log.Error("failed to get vehicle categories", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get vehicle categories"))
			return
		}

		render.JSON(w, r, Response{
			Response:          resp.OK(),
			VehicleCategories: vehicleCategories,
		})
	}
}
