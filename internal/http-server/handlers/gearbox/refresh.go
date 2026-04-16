package gearbox

import (
	"log/slog"
	"net/http"

	response "github.com/NakonechniyVitalii/GoVehicleApi/internal/lib/api/response"
	service "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/gearbox"

	"github.com/go-chi/render"
)

type Response struct {
	response.Response
}

// Refresh godoc
// @Summary      Синхронізувати коробки передач з Autoria
// @Tags         gearbox
// @Security     BearerAuth
// @Produce      json
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /gearbox/refresh [put]
func Refresh(log *slog.Logger, service *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.gearbox.Refresh"))

		err := service.Fetch(r.Context())
		if err != nil {
			log.Error("failed to update gearboxes", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Failed to update gearboxes"))
			return
		}

		render.JSON(w, r, response.OK())
	}
}
