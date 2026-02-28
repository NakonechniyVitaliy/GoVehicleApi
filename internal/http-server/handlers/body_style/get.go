package body_style

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	bodyStyle "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/body_style"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type GetResponse struct {
	Response  resp.Response
	BodyStyle *models.BodyStyle
}

func Get(log *slog.Logger, repository bodyStyle.RepositoryInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		const op = "handlers.body_style.get"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		id64, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
		if err != nil {
			log.Error("failed to get body style ID", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get body style ID"))
			return
		}
		BodyStyleID := uint16(id64)
		log.Info("ID retrieved successfully", slog.Any("body style ID", BodyStyleID))

		log.Info("getting BodyStyle")
		BodyStyle, err := repository.GetByID(r.Context(), BodyStyleID)
		if err != nil {
			log.Error("failed to get body style", slog.String("error", err.Error()))

			if errors.Is(err, storage.ErrBodyStyleNotFound) {
				render.Status(r, http.StatusNotFound)
				render.JSON(w, r, resp.Error("body style not found"))
				return
			}

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, resp.Error("Failed to get body style"))
			return
		}

		render.JSON(w, r, GetResponse{
			Response:  resp.OK(),
			BodyStyle: BodyStyle,
		})
	}
}
