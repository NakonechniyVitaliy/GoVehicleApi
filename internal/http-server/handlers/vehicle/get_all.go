package vehicle

import (
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type GetAllResponse struct {
	Response resp.Response
	Vehicles []models.Vehicle
}

func GetAll(log *slog.Logger, repository vehicle.RepositoryInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.vehicle.get"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		log.Info("getting vehicles")

		vehicles, err := repository.GetAll(r.Context())
		if err != nil {
			log.Error("failed to get vehicle", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get vehicle"))
			return
		}

		render.JSON(w, r, GetAllResponse{
			Response: resp.OK(),
			Vehicles: vehicles,
		})
	}
}
