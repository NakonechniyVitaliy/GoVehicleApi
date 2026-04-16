package vehicle

import (
	"errors"
	"log/slog"
	"net/http"

	dtoErrors "github.com/NakonechniyVitalii/GoVehicleApi/internal/http-server/dto/_errors"
	dto "github.com/NakonechniyVitalii/GoVehicleApi/internal/http-server/dto/vehicle"
	response "github.com/NakonechniyVitalii/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/vehicle"
	"github.com/go-chi/render"
)

type SaveResponse struct {
	Response response.Response
	Vehicle  *models.Vehicle
}

// New godoc
// @Summary      Створити транспортний засіб
// @Tags         vehicle
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body  body      VehiclePayload  true  "Дані транспортного засобу"
// @Success      200   {object}  SaveResponse
// @Failure      400   {object}  response.Response
// @Failure      409   {object}  response.Response
// @Failure      500   {object}  response.Response
// @Router       /vehicle/ [post]
func New(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.vehicle.save"))

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

		mappedVehicle := req.Vehicle.ToModel()
		createdVehicle, err := srv.Save(r.Context(), mappedVehicle)

		if errors.Is(err, service.ErrVehicleExists) {
			response.RenderError(w, r, http.StatusConflict, service.ErrVehicleExists.Error())
			return
		}
		if errors.Is(err, service.ErrSaveVehicle) {
			response.RenderError(w, r, http.StatusInternalServerError, service.ErrSaveVehicle.Error())
			return
		}

		render.JSON(w, r, SaveResponse{
			Response: response.OK(),
			Vehicle:  createdVehicle,
		})
	}
}
