package category

import (
	"log/slog"
	"net/http"

	response "github.com/NakonechniyVitalii/GoVehicleApi/internal/lib/api/response"
	service "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/category"
	"github.com/go-chi/render"
)

// Refresh godoc
// @Summary      Синхронізувати категорії з Autoria
// @Tags         category
// @Security     BearerAuth
// @Produce      json
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /category/refresh [put]
func Refresh(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.category.refresh"))

		err := srv.Fetch(r.Context())
		if err != nil {
			response.RenderError(w, r, http.StatusInternalServerError, service.ErrRefreshCategories.Error())
			return
		}

		render.JSON(w, r, response.OK())
	}
}
