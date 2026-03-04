package vehicle

import (
	"errors"
	"log/slog"
	"net/http"

	dto "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/vehicle"
	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/vehicle"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage"
	"github.com/go-chi/render"
)

type SaveResponse struct {
	Response resp.Response
	Vehicle  *models.Vehicle
}

func New(log *slog.Logger, service *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.vehicle.save"

		log = log.With(slog.String("op", op))

		var req dto.SaveRequest
		err := render.DecodeJSON(r.Body, &req)
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

		newVehicle := req.Vehicle.ToModel()
		log.Info("saving vehicle", slog.Any("vehicle", newVehicle))

		createdVehicle, err := service.Save(r.Context(), newVehicle)
		if errors.Is(err, storage.ErrVehicleExists) {
			log.Info("vehicle already exists", slog.String("vehicle", createdVehicle.Model))
			render.JSON(w, r, resp.Error("Vehicle already exists"))
			return
		}

		if err != nil {
			log.Error("failed to save vehicle", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to save vehicle"))
			return
		}

		log.Info("vehicle saved", slog.String("vehicle", createdVehicle.Model))

		render.JSON(w, r, SaveResponse{
			Response: resp.OK(),
			Vehicle:  createdVehicle,
		})
	}
}
