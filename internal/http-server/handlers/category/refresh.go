package category

import (
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/category"
	"github.com/go-chi/render"
)

func Refresh(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.category.refresh"))

		err := srv.Refresh(r.Context())
		if err != nil {
			resp.RenderError(w, r, http.StatusInternalServerError, service.ErrRefreshCategories.Error())
			return
		}

		render.JSON(w, r, resp.OK())
	}
}
