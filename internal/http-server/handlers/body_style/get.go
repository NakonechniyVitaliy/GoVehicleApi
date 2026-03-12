package body_style

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/body_style"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type GetResponse struct {
	Response  resp.Response
	BodyStyle *models.BodyStyle
}

func Get(log *slog.Logger, srv service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.body_style.get"))

		id64, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
		if err != nil {
			log.Error("failed to get body style ID", slog.String("error", err.Error()))
			resp.RenderError(w, r, http.StatusBadRequest, "failed to get body style ID")
			return
		}

		bodyStyleID := uint16(id64)

		BodyStyle, err := srv.GetByID(r.Context(), bodyStyleID)
		if errors.Is(err, service.ErrBodyStyleNotFound) {
			resp.RenderError(w, r, http.StatusNotFound, service.ErrBodyStyleNotFound.Error())
			return
		}
		if err != nil {
			resp.RenderError(w, r, http.StatusInternalServerError, service.ErrGetBodyStyle.Error())
			return
		}

		render.JSON(w, r, GetResponse{
			Response:  resp.OK(),
			BodyStyle: BodyStyle,
		})
	}
}
