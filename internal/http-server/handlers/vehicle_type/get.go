package vehicle_type

import (
	"log/slog"
	"net/http"
	"strconv"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	vehicleType "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle_type"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type GetResponse struct {
	Response    resp.Response
	VehicleType *models.VehicleType
}

func Get(log *slog.Logger, repository vehicleType.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		const op = "handlers.VehicleType.get.Get"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		id64, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
		if err != nil {
			log.Error("failed to get vehicle type ID", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get vehicle type ID"))
			return
		}
		VehicleTypeID := uint16(id64)
		log.Info("ID retrieved successfully", slog.Any("Vehicle type ID", VehicleTypeID))

		log.Info("getting VehicleType")
		VehicleType, err := repository.GetByID(r.Context(), VehicleTypeID)
		if err != nil {
			log.Error("failed to get vehicle type", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get vehicle type"))
			return
		}

		render.JSON(w, r, GetResponse{
			Response:    resp.OK(),
			VehicleType: VehicleType,
		})
	}
}
