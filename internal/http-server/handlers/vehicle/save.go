package vehicle

import (
	"errors"
	"log/slog"
	"net/http"

	dtoErrors "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/_errors"
	dto "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/vehicle"
	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/vehicle"
	"github.com/go-chi/render"
)

type SaveResponse struct {
	Response resp.Response
	Vehicle  *models.Vehicle
}

func New(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.vehicle.save"))

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

		mappedVehicle := req.Vehicle.ToModel()
		createdVehicle, err := srv.Save(r.Context(), mappedVehicle)

		if errors.Is(err, service.ErrVehicleExists) {
			resp.RenderError(w, r, http.StatusConflict, service.ErrVehicleExists.Error())
			return
		}
		if errors.Is(err, service.ErrSaveVehicle) {
			resp.RenderError(w, r, http.StatusInternalServerError, service.ErrSaveVehicle.Error())
			return
		}

		render.JSON(w, r, SaveResponse{
			Response: resp.OK(),
			Vehicle:  createdVehicle,
		})
	}
}
