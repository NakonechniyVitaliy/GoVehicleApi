package body_style

import (
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	bodyStyle "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/body_style"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type GetAllResponse struct {
	Response   resp.Response
	BodyStyles []models.BodyStyle
}

func GetAll(log *slog.Logger, repository bodyStyle.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.BodyStyle.getAll"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		log.Info("getting body styles")

		BodyStyles, err := repository.GetAll(r.Context())
		if err != nil {
			log.Error("failed to get body style", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get body style"))
			return
		}

		render.JSON(w, r, GetAllResponse{
			Response:   resp.OK(),
			BodyStyles: BodyStyles,
		})
	}
}
