package vehicle_type

import (
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	vehicleType "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle_type"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type GetAllResponse struct {
	Response     resp.Response
	VehicleTypes []models.VehicleType
}

func GetAll(log *slog.Logger, repository vehicleType.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.VehicleType.get.Get"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		log.Info("getting VehicleTypes")

		VehicleTypes, err := repository.GetAll(r.Context())
		if err != nil {
			log.Error("failed to get vehicle type", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get vehicle type"))
			return
		}

		render.JSON(w, r, GetAllResponse{
			Response:     resp.OK(),
			VehicleTypes: VehicleTypes,
		})
	}
}
