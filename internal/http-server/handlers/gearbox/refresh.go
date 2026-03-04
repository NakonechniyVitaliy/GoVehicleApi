package gearbox

import (
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/gearbox"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Response struct {
	resp.Response
}

func Refresh(log *slog.Logger, service *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.gearbox.Refresh"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		err := service.Refresh(r.Context())
		if err != nil {
			log.Error("failed to update gearboxes", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to update gearboxes"))
			return
		}

		render.JSON(w, r, resp.OK())
	}
}
