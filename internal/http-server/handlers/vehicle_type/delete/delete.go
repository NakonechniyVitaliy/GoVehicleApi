package save

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Response struct {
	resp.Response
}

type VehicleTypeDeleter interface {
	Delete(ctx context.Context, vehicleTypeID int) error
}

func Delete(log *slog.Logger, vehicleTypeDeleter VehicleTypeDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.vehicleType.delete.Delete"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		vehicleTypeID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Error("failed to get vehicle type ID", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get vehicle type ID"))
			return
		}

		log.Info("ID retrieved successfully", slog.Any("vehicle type ID", vehicleTypeID))
		log.Info("deleting vehicle type")

		err = vehicleTypeDeleter.Delete(r.Context(), vehicleTypeID)
		if err != nil {
			log.Error("failed to delete vehicle type", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to delete vehicle type"))
			return
		}

		log.Info("vehicle type deleted", slog.Int("vehicle type ID", vehicleTypeID))

		render.JSON(w, r, Response{
			Response: resp.OK(),
		})
	}
}
