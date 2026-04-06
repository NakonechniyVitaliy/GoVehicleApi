package vehicle

import (
	"errors"
	"log/slog"
	"net/http"

	fDTO "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/vehicle/filter"
	response "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/vehicle"
	"github.com/go-chi/render"
)

type GetAllResponse struct {
	Response response.Response
	Vehicles []models.Vehicle
}

// GetList godoc
// @Summary      Список транспортних засобів
// @Tags         vehicle
// @Security     BearerAuth
// @Produce      json
// @Param        page   query     int  false  "Номер сторінки"
// @Param        limit  query     int  false  "Кількість на сторінці"
// @Success      200    {object}  GetAllResponse
// @Failure      400    {object}  response.Response
// @Failure      500    {object}  response.Response
// @Router       /vehicle/ [get]
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
			response.RenderError(w, r, http.StatusBadRequest, err.Error())
			return
		}

		vehicles, err := srv.GetList(r.Context(), *filter)
		if errors.Is(err, service.ErrGetVehicles) {
			response.RenderError(w, r, http.StatusInternalServerError, service.ErrGetVehicles.Error())
			return
		}
		render.JSON(w, r, GetAllResponse{
			Response: response.OK(),
			Vehicles: vehicles,
		})
	}
}
