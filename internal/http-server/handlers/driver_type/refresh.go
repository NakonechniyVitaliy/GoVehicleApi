package driver_type

import (
	"log/slog"
	"net/http"

	response "github.com/NakonechniyVitalii/GoVehicleApi/internal/lib/api/response"
	service "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/driver_type"
	"github.com/go-chi/render"
)

type Response struct {
	response.Response
}

// Refresh godoc
// @Summary      Синхронізувати типи приводу з Autoria
// @Tags         driver-type
// @Security     BearerAuth
// @Produce      json
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /driver-type/refresh [put]
func Refresh(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.driver_type.refresh"

		log = log.With(slog.String("op", op))

		err := srv.Fetch(r.Context())
		if err != nil {
			log.Error("failed to update driver types", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Failed to update driver types"))
			return
		}

		render.JSON(w, r, response.OK())
	}
}
