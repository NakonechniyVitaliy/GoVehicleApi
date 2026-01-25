package save

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
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

type VehicleTypeSaver interface {
	Create(ctx context.Context, vehicleType models.VehicleType) error
}

func New(log *slog.Logger, vehicleTypeSaver VehicleTypeSaver) http.HandlerFunc {
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

		vehicleType := models.VehicleType{
			Ablative:   req.VehicleType.Ablative,
			CategoryID: req.VehicleType.CategoryID,
			Name:       req.VehicleType.Name,
			Plural:     req.VehicleType.Plural,
			Rewrite:    req.VehicleType.Rewrite,
			Singular:   req.VehicleType.Singular,
		}

		log.Info("saving vehicle type", slog.Any("vehicle type", vehicleType))

		err = vehicleTypeSaver.Create(r.Context(), vehicleType)
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
