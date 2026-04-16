package vehicle

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	response "github.com/NakonechniyVitalii/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/vehicle"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type GetResponse struct {
	Response response.Response
	Vehicle  *models.Vehicle
}

// Get godoc
// @Summary      Отримати транспортний засіб за ID
// @Tags         vehicle
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "ID транспортного засобу"
// @Success      200  {object}  GetResponse
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /vehicle/{id} [get]
func Get(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.vehicle.get"))

		id64, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
		if err != nil {
			log.Error("failed to get vehicle ID", slog.String("error", err.Error()))
			response.RenderError(w, r, http.StatusBadRequest, "failed to get vehicle ID")
			return
		}

		vehicleID := uint16(id64)

		Vehicle, err := srv.GetByID(r.Context(), vehicleID)
		if errors.Is(err, service.ErrVehicleNotFound) {
			response.RenderError(w, r, http.StatusNotFound, service.ErrVehicleNotFound.Error())
			return
		}
		if err != nil {
			response.RenderError(w, r, http.StatusInternalServerError, service.ErrGetVehicle.Error())
			return
		}

		render.JSON(w, r, GetResponse{
			Response: response.OK(),
			Vehicle:  Vehicle,
		})
	}
}
