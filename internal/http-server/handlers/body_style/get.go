package body_style

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	response "github.com/NakonechniyVitalii/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/body_style"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type GetResponse struct {
	Response  response.Response
	BodyStyle *models.BodyStyle
}

// Get godoc
// @Summary      Отримати стиль кузова за ID
// @Tags         body-style
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "ID стилю кузова"
// @Success      200  {object}  GetResponse
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /body-style/{id} [get]
func Get(log *slog.Logger, srv service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.body_style.get"))

		id64, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
		if err != nil {
			log.Error("failed to get body style ID", slog.String("error", err.Error()))
			response.RenderError(w, r, http.StatusBadRequest, "failed to get body style ID")
			return
		}

		bodyStyleID := uint16(id64)

		BodyStyle, err := srv.GetByID(r.Context(), bodyStyleID)
		if errors.Is(err, service.ErrBodyStyleNotFound) {
			response.RenderError(w, r, http.StatusNotFound, service.ErrBodyStyleNotFound.Error())
			return
		}
		if err != nil {
			response.RenderError(w, r, http.StatusInternalServerError, service.ErrGetBodyStyle.Error())
			return
		}

		render.JSON(w, r, GetResponse{
			Response:  response.OK(),
			BodyStyle: BodyStyle,
		})
	}
}
