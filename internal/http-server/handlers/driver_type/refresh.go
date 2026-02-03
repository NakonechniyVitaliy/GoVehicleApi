package driver_type

import (
	"log/slog"
	"net/http"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	driverTypeRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/driver_type"
	driverTypeService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/service/driver_type"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Response struct {
	resp.Response
}

func Refresh(log *slog.Logger, repository driverTypeRepo.Repository, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.driverType.update.Refresh"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		err := driverTypeService.RefreshDriverTypes(r.Context(), cfg, repository)
		if err != nil {
			log.Error("failed to update driver types", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to update driver types"))
			return
		}

		render.JSON(w, r, resp.OK())
	}
}
