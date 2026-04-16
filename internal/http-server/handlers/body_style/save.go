package body_style

import (
	"errors"
	"log/slog"
	"net/http"

	dtoErrors "github.com/NakonechniyVitalii/GoVehicleApi/internal/http-server/dto/_errors"
	dto "github.com/NakonechniyVitalii/GoVehicleApi/internal/http-server/dto/body_style"
	response "github.com/NakonechniyVitalii/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/body_style"
	"github.com/go-chi/render"
)

type SaveResponse struct {
	Response  response.Response
	BodyStyle *models.BodyStyle
}

// New godoc
// @Summary      Створити стиль кузова
// @Tags         body-style
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body  body      BodyStylePayload  true  "Дані стилю кузова"
// @Success      200   {object}  SaveResponse
// @Failure      400   {object}  response.Response
// @Failure      409   {object}  response.Response
// @Failure      500   {object}  response.Response
// @Router       /body-style/ [post]
func New(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.body_style.new"))

		var req dto.SaveRequest
		err := render.DecodeJSON(r.Body, &req)
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

		mappedBodyStyle := req.BodyStyle.ToModel()
		createdBS, err := srv.Save(r.Context(), mappedBodyStyle)

		if errors.Is(err, service.ErrBodyStyleExists) {
			response.RenderError(w, r, http.StatusConflict, service.ErrBodyStyleExists.Error())
			return
		}
		if errors.Is(err, service.ErrSaveBodyStyle) {
			response.RenderError(w, r, http.StatusInternalServerError, service.ErrSaveBodyStyle.Error())
			return
		}

		render.JSON(w, r, SaveResponse{
			Response:  response.OK(),
			BodyStyle: createdBS,
		})
	}
}
