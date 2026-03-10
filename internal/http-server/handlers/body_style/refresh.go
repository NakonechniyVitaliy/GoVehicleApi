package body_style

import (
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/body_style"
	"github.com/go-chi/render"
)

func Refresh(log *slog.Logger, srv service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.body_style.refresh"))

		err := srv.Refresh(r.Context())
		if err != nil {
			resp.RenderError(w, r, http.StatusInternalServerError, service.ErrRefreshBodyStyles.Error())
			return
		}

		render.JSON(w, r, resp.OK())
	}
}
