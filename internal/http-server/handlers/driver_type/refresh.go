package driver_type

import (
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/driver_type"
	"github.com/go-chi/render"
)

type Response struct {
	resp.Response
}

func Refresh(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.driver_type.refresh"

		log = log.With(slog.String("op", op))

		err := srv.Fetch(r.Context())
		if err != nil {
			log.Error("failed to update driver types", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to update driver types"))
			return
		}

		render.JSON(w, r, resp.OK())
	}
}
