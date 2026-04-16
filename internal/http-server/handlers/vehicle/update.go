package vehicle

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	dtoErrors "github.com/NakonechniyVitalii/GoVehicleApi/internal/http-server/dto/_errors"
	dto "github.com/NakonechniyVitalii/GoVehicleApi/internal/http-server/dto/vehicle"
	response "github.com/NakonechniyVitalii/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/vehicle"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type UpdateResponse struct {
	Response response.Response
	Vehicle  *models.Vehicle
}

// Update godoc
// @Summary      Оновити транспортний засіб
// @Tags         vehicle
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id    path      int             true  "ID транспортного засобу"
// @Param        body  body      VehiclePayload  true  "Нові дані"
// @Success      200   {object}  UpdateResponse
// @Failure      400   {object}  response.Response
// @Failure      404   {object}  response.Response
// @Failure      500   {object}  response.Response
// @Router       /vehicle/{id} [put]
func Update(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.vehicle.update"))

		id64, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
		if err != nil {
			log.Error("failed to get vehicle ID", slog.String("error", err.Error()))
			response.RenderError(w, r, http.StatusBadRequest, "failed to get vehicle ID")
			return
		}

		vehicleID := uint16(id64)
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

		bodyStyleFromRequest := req.Vehicle.ToModel()
		updatedVehicle, err := srv.Update(r.Context(), bodyStyleFromRequest, vehicleID)

		if errors.Is(err, service.ErrVehicleNotFound) {
			response.RenderError(w, r, http.StatusNotFound, service.ErrVehicleNotFound.Error())
			return
		}
		if errors.Is(err, service.ErrUpdateVehicle) {
			response.RenderError(w, r, http.StatusInternalServerError, service.ErrUpdateVehicle.Error())
			return
		}

		render.JSON(w, r, UpdateResponse{
			Response: response.OK(),
			Vehicle:  updatedVehicle,
		})
	}
}
