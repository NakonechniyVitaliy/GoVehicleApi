package brand

import (
	"log/slog"
	"net/http"

	response "github.com/NakonechniyVitalii/GoVehicleApi/internal/lib/api/response"
	service "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/brand"
	"github.com/go-chi/render"
)

// Refresh godoc
// @Summary      Синхронізувати бренди з Autoria
// @Tags         brand
// @Security     BearerAuth
// @Produce      json
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /brand/refresh [put]
func Refresh(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.brand.refresh"))

		err := srv.Fetch(r.Context())
		if err != nil {
			response.RenderError(w, r, http.StatusInternalServerError, service.ErrRefreshBrands.Error())
			return
		}

		render.JSON(w, r, response.OK())
	}
}
