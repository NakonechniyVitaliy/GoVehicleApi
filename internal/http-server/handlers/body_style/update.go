package body_style

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	dtoErrors "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/_errors"
	dto "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/body_style"
	response "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/body_style"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type UpdateResponse struct {
	Response  response.Response
	BodyStyle *models.BodyStyle
}

// Update godoc
// @Summary      Оновити стиль кузова
// @Tags         body-style
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id    path      int               true  "ID стилю кузова"
// @Param        body  body      BodyStylePayload  true  "Нові дані"
// @Success      200   {object}  UpdateResponse
// @Failure      400   {object}  response.Response
// @Failure      404   {object}  response.Response
// @Failure      500   {object}  response.Response
// @Router       /body-style/{id} [put]
func Update(log *slog.Logger, srv service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log = log.With(slog.String("op", "handlers.body_style.update"))

		id64, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
		if err != nil {
			log.Error("failed to get brand ID", slog.String("error", err.Error()))
			response.RenderError(w, r, http.StatusBadRequest, "failed to get body style ID")
			return
		}

		bodyStyleID := uint16(id64)
		var req dto.UpdateRequest

		err = render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error(dtoErrors.InvalidJSONorWrongFieldType, slog.String("error", err.Error()))
			response.RenderError(w, r, http.StatusBadRequest, dtoErrors.InvalidJSONorWrongFieldType)
			return
		}

		err = req.Validate()
		if err != nil {
			log.Error("validation error", slog.String("error", err.Error()))
			response.RenderError(w, r, http.StatusBadRequest, err.Error())
			return
		}

		bodyStyleFromRequest := req.BodyStyle.ToModel()
		updatedBodyStyle, err := srv.Update(r.Context(), bodyStyleFromRequest, bodyStyleID)

		if errors.Is(err, service.ErrBodyStyleNotFound) {
			response.RenderError(w, r, http.StatusNotFound, service.ErrBodyStyleNotFound.Error())
			return
		}
		if errors.Is(err, service.ErrUpdateBodyStyle) {
			response.RenderError(w, r, http.StatusInternalServerError, service.ErrUpdateBodyStyle.Error())
			return
		}

		render.JSON(w, r, UpdateResponse{
			Response:  response.OK(),
			BodyStyle: updatedBodyStyle,
		})
	}
}
