package driver_type

import (
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	driverTypeRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/driver_type"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type GetAllResponse struct {
	Response    resp.Response
	DriverTypes []models.DriverType
}

func GetAll(log *slog.Logger, repo driverTypeRepo.RepositoryInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.driverType.get.GetAll"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		log.Info("getting vehicle driver types")

		driverTypes, err := repo.GetAll(r.Context())
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
