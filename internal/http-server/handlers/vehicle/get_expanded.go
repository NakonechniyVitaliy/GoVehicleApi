package vehicle

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	dto "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/vehicle"
	response "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/vehicle"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type GetExpandedResponse struct {
	Response response.Response
	Vehicle  *dto.ExpandedVehicleDTO
}

// GetExpanded godoc
// @Summary      Отримати розгорнутий транспортний засіб за ID
// @Tags         vehicle
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "ID транспортного засобу"
// @Success      200  {object}  GetExpandedSwaggerResponse
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /vehicle/expanded/{id} [get]
func GetExpanded(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.vehicle.get_expanded"))

		id64, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
		if err != nil {
			log.Error("failed to get vehicle ID", slog.String("error", err.Error()))
			response.RenderError(w, r, http.StatusBadRequest, "failed to get vehicle ID")
			return
		}
		vehicleID := uint16(id64)

		expandedVehicle, err := srv.GetExpanded(r.Context(), vehicleID)
		if errors.Is(err, service.ErrVehicleNotFound) {
			response.RenderError(w, r, http.StatusNotFound, service.ErrVehicleNotFound.Error())
			return
		}
		if err != nil {
			response.RenderError(w, r, http.StatusInternalServerError, service.ErrGetVehicle.Error())
			return
		}

		render.JSON(w, r, GetExpandedResponse{
			Response: response.OK(),
			Vehicle:  expandedVehicle,
		})
	}
}
