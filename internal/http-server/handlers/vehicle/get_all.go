package vehicle

import (
	"errors"
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/vehicle"
	"github.com/go-chi/render"
)

type GetAllResponse struct {
	Response resp.Response
	Vehicles []models.Vehicle
}

func GetAll(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.vehicle.get"))

		vehicles, err := srv.GetAll(r.Context())

		if errors.Is(err, service.ErrGetVehicles) {
			resp.RenderError(w, r, http.StatusInternalServerError, service.ErrGetVehicles.Error())
			return
		}
		render.JSON(w, r, GetAllResponse{
			Response: resp.OK(),
			Vehicles: vehicles,
		})
	}
}
