package category

import (
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/category"
	"github.com/go-chi/render"
)

func Refresh(log *slog.Logger, service *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.category.refresh"

		log = log.With(slog.String("op", op))

		err := service.Refresh(r.Context())
		if err != nil {
			log.Error("failed to update categories", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to update category"))
			return
		}

		render.JSON(w, r, resp.OK())
	}
}
