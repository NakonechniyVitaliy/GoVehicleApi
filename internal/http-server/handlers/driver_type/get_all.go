package driver_type

import (
	"log/slog"
	"net/http"

	response "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/driver_type"
	"github.com/go-chi/render"
)

type GetAllResponse struct {
	Response    response.Response
	DriverTypes []models.DriverType
}

// GetAll godoc
// @Summary      Список типів приводу
// @Tags         driver-type
// @Security     BearerAuth
// @Produce      json
// @Success      200  {object}  GetAllResponse
// @Failure      500  {object}  response.Response
// @Router       /driver-type/all [get]
func GetAll(log *slog.Logger, service *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.driver_type.get_all"

		log = log.With(slog.String("op", op))

		log.Info("getting vehicle driver types")

		driverTypes, err := service.GetAll(r.Context())
		if err != nil {
			log.Error("failed to get vehicle driver types", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Failed to get vehicle driver types"))
			return
		}

		render.JSON(w, r, GetAllResponse{
			Response:    response.OK(),
			DriverTypes: driverTypes,
		})
	}
}
