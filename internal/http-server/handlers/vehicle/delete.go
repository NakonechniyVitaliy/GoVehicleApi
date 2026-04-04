package vehicle

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	response "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/vehicle"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// Delete godoc
// @Summary      Видалити транспортний засіб
// @Tags         vehicle
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "ID транспортного засобу"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /vehicle/{id} [delete]
func Delete(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.vehicle.delete"))

		id64, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
		if err != nil {
			log.Error("failed to get vehicle ID", slog.String("error", err.Error()))
			response.RenderError(w, r, http.StatusBadRequest, "failed to get vehicle ID")
			return
		}
		vehicleID := uint16(id64)

		err = srv.Delete(r.Context(), vehicleID)
		if errors.Is(err, service.ErrVehicleNotFound) {
			response.RenderError(w, r, http.StatusNotFound, service.ErrVehicleNotFound.Error())
			return
		}
		if err != nil {
			response.RenderError(w, r, http.StatusInternalServerError, service.ErrGetVehicle.Error())
			return
		}
		render.JSON(w, r, response.OK())
	}
}
