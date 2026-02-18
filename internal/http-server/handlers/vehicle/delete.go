package vehicle

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func Delete(log *slog.Logger, repository vehicle.RepositoryInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.vehicle.delete"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		id64, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
		if err != nil {
			log.Error("failed to get vehicle ID", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get vehicle ID"))
			return
		}
		vehicleID := uint16(id64)
		log.Info("ID retrieved successfully", slog.Any("vehicleID", vehicleID))

		log.Info("deleting vehicle")
		err = repository.Delete(r.Context(), vehicleID)
		if err != nil {
			log.Error("failed to delete vehicle", slog.String("error", err.Error()))

			if errors.Is(err, storage.ErrVehicleNotFound) {
				render.Status(r, http.StatusNotFound)
				render.JSON(w, r, resp.Error("vehicle not found"))
				return
			}

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, resp.Error("Failed to delete vehicle"))
			return
		}

		log.Info("vehicle deleted", slog.Any("vehicleID", vehicleID))

		render.JSON(w, r, resp.OK())
	}
}
