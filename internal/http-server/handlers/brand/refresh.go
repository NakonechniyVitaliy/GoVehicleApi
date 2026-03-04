package brand

import (
	"errors"
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/brand"
	"github.com/go-chi/render"
)

func Refresh(log *slog.Logger, service *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.brand.refresh"

		log = log.With(slog.String("op", op))

		err := service.Refresh(r.Context())
		if err != nil {
			log.Error("failed to update brands", slog.String("error", err.Error()))

			if errors.Is(err, requests.ErrAutoRiaBrands) {
				render.JSON(w, r, resp.Error("Failed to update brand, autoRia error"))
				return
			}

			render.JSON(w, r, resp.Error("Failed to update brand"))
			return
		}

		render.JSON(w, r, resp.OK())
	}
}
