package body_style

import (
	"errors"
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	bodyStyle "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/body_style"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type SaveRequest struct {
	BodyStyle models.BodyStyle
}

func New(log *slog.Logger, repository bodyStyle.RepositoryInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.bodyStyle.new"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req SaveRequest

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

		newBodyStyle := models.BodyStyle{
			Name:  req.BodyStyle.Name,
			Value: req.BodyStyle.Value,
		}

		log.Info("saving body style", slog.Any("body style", newBodyStyle))

		err = repository.Create(r.Context(), newBodyStyle)
		if errors.Is(err, storage.ErrBodyStyleExists) {
			log.Info("bodyStyle already exists", slog.String("body style", req.BodyStyle.Name))
			render.JSON(w, r, resp.Error("body style already exists"))
			return
		}

		if err != nil {
			log.Error("failed to save body style", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to save body style"))
			return
		}

		log.Info("bodyStyle saved", slog.String("body style", req.BodyStyle.Name))

		render.JSON(w, r, resp.OK())
	}
}
