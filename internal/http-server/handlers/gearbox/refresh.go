package gearbox

import (
	"log/slog"
	"net/http"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	gearboxRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/gearbox"
	gearboxService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/service/gearbox"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Response struct {
	resp.Response
}

func Refresh(log *slog.Logger, repository gearboxRepo.RepositoryInterface, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.gearbox.Refresh"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		err := gearboxService.RefreshGearboxes(r.Context(), cfg, repository)
		if err != nil {
			log.Error("failed to update gearboxes", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to update gearboxes"))
			return
		}

		render.JSON(w, r, resp.OK())
	}
}
