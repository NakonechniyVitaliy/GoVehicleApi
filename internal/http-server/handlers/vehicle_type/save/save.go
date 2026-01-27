package save

import (
	"errors"
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	vehicleType "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle_type"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	VehicleType models.VehicleType
}

type Response struct {
	resp.Response
}

func New(log *slog.Logger, repository vehicleType.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.vehicleType.save.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to decode request"))
			return
		}
		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			log.Error("invalid request", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Invalid request"))
			return
		}

		newVehicleType := models.VehicleType{
			Ablative:   req.VehicleType.Ablative,
			CategoryID: req.VehicleType.CategoryID,
			Name:       req.VehicleType.Name,
			Plural:     req.VehicleType.Plural,
			Rewrite:    req.VehicleType.Rewrite,
			Singular:   req.VehicleType.Singular,
		}

		log.Info("saving vehicle type", slog.Any("vehicle type", newVehicleType))

		err = repository.Create(r.Context(), newVehicleType)
		if errors.Is(err, storage.ErrVehicleTypeExists) {
			log.Info("vehicleType already exists", slog.String("vehicle type", req.VehicleType.Name))
			render.JSON(w, r, resp.Error("vehicle type already exists"))
			return
		}

		if err != nil {
			log.Error("failed to save vehicle type", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to save vehicle type"))
			return
		}

		log.Info("vehicleType saved", slog.String("vehicle type", req.VehicleType.Name))

		render.JSON(w, r, Response{
			Response: resp.OK(),
		})
	}
}
