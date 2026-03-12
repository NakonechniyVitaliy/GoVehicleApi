package vehicle

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	dtoErrors "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/_errors"
	dto "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/vehicle"
	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/vehicle"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type UpdateResponse struct {
	Response resp.Response
	Vehicle  *models.Vehicle
}

func Update(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.vehicle.update"))

		id64, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
		if err != nil {
			log.Error("failed to get vehicle ID", slog.String("error", err.Error()))
			resp.RenderError(w, r, http.StatusBadRequest, "failed to get vehicle ID")
			return
		}

		vehicleID := uint16(id64)
		var req dto.UpdateRequest

		err = render.DecodeJSON(r.Body, &req)
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

		bodyStyleFromRequest := req.Vehicle.ToModel()
		updatedVehicle, err := srv.Update(r.Context(), bodyStyleFromRequest, vehicleID)

		if errors.Is(err, service.ErrVehicleNotFound) {
			resp.RenderError(w, r, http.StatusNotFound, service.ErrVehicleNotFound.Error())
			return
		}
		if errors.Is(err, service.ErrUpdateVehicle) {
			resp.RenderError(w, r, http.StatusInternalServerError, service.ErrUpdateVehicle.Error())
			return
		}

		render.JSON(w, r, UpdateResponse{
			Response: resp.OK(),
			Vehicle:  updatedVehicle,
		})
	}
}
