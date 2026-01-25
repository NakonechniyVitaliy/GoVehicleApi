package update

import (
	"context"
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
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

type VehicleTypeUpdater interface {
	Update(ctx context.Context, vehicleType models.VehicleType) error
}

func Update(log *slog.Logger, vehicleTypeUpdater VehicleTypeUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.vehicleType.update.Update"

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

		err = vehicleTypeUpdater.Update(r.Context(), vehicleType)
		if err != nil {
			log.Error("failed to update vehicleType", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to update vehicleType"))
			return
		}

		render.JSON(w, r, Response{
			Response: resp.OK(),
		})

	}
}
