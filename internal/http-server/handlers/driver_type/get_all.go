package driver_type

import (
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/driver_type"
	"github.com/go-chi/render"
)

type GetAllResponse struct {
	Response    resp.Response
	DriverTypes []models.DriverType
}

func GetAll(log *slog.Logger, service *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.driver_type.get_all"

		log = log.With(slog.String("op", op))

		log.Info("getting vehicle driver types")

		driverTypes, err := service.GetAll(r.Context())
		if err != nil {
			log.Error("failed to get vehicle driver types", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get vehicle driver types"))
			return
		}

		render.JSON(w, r, GetAllResponse{
			Response:    resp.OK(),
			DriverTypes: driverTypes,
		})
	}
}
