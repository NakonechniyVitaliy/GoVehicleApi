package body_style

import (
	"log/slog"
	"net/http"
	"strconv"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	bodyStyle "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/body_style"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func Delete(log *slog.Logger, repository bodyStyle.RepositoryInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.bodyStyle.delete"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		id64, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
		if err != nil {
			log.Error("failed to get body style ID", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get body style ID"))
			return
		}
		bodyStyleID := uint16(id64)
		log.Info("ID retrieved successfully", slog.Any("body style ID", bodyStyleID))

		log.Info("deleting body style")
		err = repository.Delete(r.Context(), bodyStyleID)
		if err != nil {
			log.Error("failed to delete body style", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to delete body style"))
			return
		}

		log.Info("body style deleted", slog.Any("body style ID", bodyStyleID))

		render.JSON(w, r, resp.OK())
	}
}
