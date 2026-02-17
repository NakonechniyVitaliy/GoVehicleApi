package body_style

import (
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	bodyStyle "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/body_style"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type UpdateRequest struct {
	BodyStyle models.BodyStyle
}

func Update(log *slog.Logger, repository bodyStyle.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.bodyStyle.update.Update"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req UpdateRequest
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

		updatedBodyStyle := models.BodyStyle{
			ID:    req.BodyStyle.ID,
			Name:  req.BodyStyle.Name,
			Value: req.BodyStyle.Value,
		}

		err = repository.Update(r.Context(), updatedBodyStyle)
		if err != nil {
			log.Error("failed to update body style", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to update body style"))
			return
		}

		render.JSON(w, r, resp.OK())

	}
}
