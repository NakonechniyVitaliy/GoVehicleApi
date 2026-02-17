package body_style

import (
	"log/slog"
	"net/http"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	bodyStyleRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/body_style"
	bodyStyleService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/service/body_style"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func Refresh(log *slog.Logger, repository bodyStyleRepo.Repository, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.bodyStyle.refresh"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		err := bodyStyleService.RefreshBodyStyles(r.Context(), cfg, repository)
		if err != nil {
			log.Error("failed to update body styles", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to update body style"))
			return
		}

		render.JSON(w, r, resp.OK())
	}
}
