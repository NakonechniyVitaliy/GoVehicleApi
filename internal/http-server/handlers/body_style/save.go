package body_style

import (
	"errors"
	"log/slog"
	"net/http"

	dtoErrors "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/_errors"
	dto "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/body_style"
	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/body_style"
	"github.com/go-chi/render"
)

type SaveResponse struct {
	Response  resp.Response
	BodyStyle *models.BodyStyle
}

func New(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.body_style.new"))

		var req dto.SaveRequest
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error(dtoErrors.InvalidJSONorWrongFieldType, slog.String("error", err.Error()))
			resp.RenderError(w, r, http.StatusBadRequest, dtoErrors.InvalidJSONorWrongFieldType)
			return
		}
		err = req.Validate()
		if err != nil {
			log.Error("validation error", slog.String("error", err.Error()))
			resp.RenderError(w, r, http.StatusBadRequest, err.Error())
			return
		}

		mappedBodyStyle := req.BodyStyle.ToModel()
		createdBS, err := srv.Save(r.Context(), mappedBodyStyle)

		if errors.Is(err, service.ErrBodyStyleExists) {
			resp.RenderError(w, r, http.StatusConflict, service.ErrBodyStyleExists.Error())
			return
		}
		if errors.Is(err, service.ErrSaveBodyStyle) {
			resp.RenderError(w, r, http.StatusInternalServerError, service.ErrSaveBodyStyle.Error())
			return
		}

		render.JSON(w, r, SaveResponse{
			Response:  resp.OK(),
			BodyStyle: createdBS,
		})
	}
}
