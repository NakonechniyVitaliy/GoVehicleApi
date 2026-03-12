package vehicle

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/vehicle"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func Delete(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.vehicle.delete"))

		id64, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
		if err != nil {
			log.Error("failed to get vehicle ID", slog.String("error", err.Error()))
			resp.RenderError(w, r, http.StatusBadRequest, "failed to get vehicle ID")
			return
		}
		vehicleID := uint16(id64)

		err = srv.Delete(r.Context(), vehicleID)
		if errors.Is(err, service.ErrVehicleNotFound) {
			resp.RenderError(w, r, http.StatusNotFound, service.ErrVehicleNotFound.Error())
			return
		}
		if err != nil {
			resp.RenderError(w, r, http.StatusInternalServerError, service.ErrGetVehicle.Error())
			return
		}
		render.JSON(w, r, resp.OK())
	}
}
