package vehicle

import (
	"log/slog"
	"net/http"
	"strconv"

	dto "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/vehicle"
	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type UpdateResponse struct {
	Response resp.Response
	Vehicle  *models.Vehicle
}

func Update(log *slog.Logger, repository vehicle.RepositoryInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.vehicle.update"

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

		var req dto.UpdateRequest
		err = render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("invalid JSON or wrong field types", slog.String("error", err.Error()))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("invalid JSON or wrong field types"))
			return
		}
		log.Info("request body decoded", slog.Any("request", req))

		err = req.Validate()
		if err != nil {
			log.Error("validation error", slog.String("error", err.Error()))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error(err.Error()))
			return
		}

		vehicleFromRequest := req.Vehicle.ToModel()

		updatedVehicle, err := repository.Update(r.Context(), vehicleFromRequest, vehicleID)
		if err != nil {
			log.Error("failed to update vehicle", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to update vehicle"))
			return
		}

		render.JSON(w, r, UpdateResponse{
			Response: resp.OK(),
			Vehicle:  updatedVehicle,
		})

	}
}
