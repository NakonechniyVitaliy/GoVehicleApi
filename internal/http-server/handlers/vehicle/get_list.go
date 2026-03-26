package vehicle

import (
	"errors"
	"log/slog"
	"net/http"

	fDTO "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/vehicle/filter"
	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/vehicle"
	"github.com/go-chi/render"
)

type GetAllResponse struct {
	Response resp.Response
	Vehicles []models.Vehicle
}

func GetList(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log = log.With(slog.String("op", "handlers.vehicle.get_list"))

		page := r.URL.Query().Get("page")
		limit := r.URL.Query().Get("limit")

		filterDTO := fDTO.FilterDTO{
			Page:  page,
			Limit: limit,
		}

		filter, err := filterDTO.ValidateAndToModel()
		if err != nil {
			resp.RenderError(w, r, http.StatusBadRequest, err.Error())
		}

		vehicles, err := srv.GetList(r.Context(), *filter)
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
