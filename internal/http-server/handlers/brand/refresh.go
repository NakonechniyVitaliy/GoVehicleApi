package brand

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/brand"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests"
	brandService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/service/brand"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func Refresh(log *slog.Logger, repository brand.Repository, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.brand.update.Refresh"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		err := brandService.RefreshBrands(r.Context(), cfg, repository)
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
