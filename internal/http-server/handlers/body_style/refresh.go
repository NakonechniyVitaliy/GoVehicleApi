package body_style

import (
	"log/slog"
	"net/http"

	response "github.com/NakonechniyVitalii/GoVehicleApi/internal/lib/api/response"
	service "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/body_style"
	"github.com/go-chi/render"
)

// Refresh godoc
// @Summary      Синхронізувати стилі кузова з Autoria
// @Tags         body-style
// @Security     BearerAuth
// @Produce      json
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /body-style/refresh [put]
func Refresh(log *slog.Logger, srv service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.body_style.refresh"))

		err := srv.Fetch(r.Context())
		if err != nil {
			response.RenderError(w, r, http.StatusInternalServerError, service.ErrRefreshBodyStyles.Error())
			return
		}

		render.JSON(w, r, response.OK())
	}
}
