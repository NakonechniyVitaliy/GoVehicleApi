package vehicle_type

import (
	"log/slog"
	"net/http"
	"strconv"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	vehicleType "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle_type"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func Delete(log *slog.Logger, repository vehicleType.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.vehicleType.delete.Delete"

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
		vehicleTypeID := uint16(id64)
		log.Info("ID retrieved successfully", slog.Any("vehicle type ID", vehicleTypeID))

		log.Info("deleting vehicle type")
		err = repository.Delete(r.Context(), vehicleTypeID)
		if err != nil {
			log.Error("failed to delete vehicle type", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to delete vehicle type"))
			return
		}

		log.Info("vehicle type deleted", slog.Any("vehicle type ID", vehicleTypeID))

		render.JSON(w, r, resp.OK())
	}
}
