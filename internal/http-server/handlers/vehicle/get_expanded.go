package vehicle

import (
	"log/slog"
	"net/http"
	"strconv"

	dto "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/vehicle"
	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/vehicle"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type GetExpandedResponse struct {
	Response resp.Response
	Vehicle  *dto.ExpandedVehicleDTO
}

func GetExpanded(log *slog.Logger, service *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.vehicle.get_expanded"

		log = log.With(slog.String("op", op))

		id64, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
		if err != nil {
			log.Error("failed to get vehicle ID", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get vehicle ID"))
			return
		}
		vehicleID := uint16(id64)
		log.Info("ID retrieved successfully", slog.Any("vehicleID", vehicleID))

		log.Info("getting expanded vehicle")
		expandedVehicle, err := service.GetExpanded(r.Context(), vehicleID)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, resp.Error("Failed to get vehicle"))
			return
		}

		render.JSON(w, r, GetExpandedResponse{
			Response: resp.OK(),
			Vehicle:  expandedVehicle,
		})
	}
}
