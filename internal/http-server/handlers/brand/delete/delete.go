package save

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Response struct {
	resp.Response
}

type BrandDeleter interface {
	Delete(ctx context.Context, brandID int) error
}

func Delete(log *slog.Logger, brandDeleter BrandDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.brand.delete.Delete"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		brandID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Error("failed to get brand ID", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get brand ID"))
			return
		}

		log.Info("ID retrieved successfully", slog.Any("brandID", brandID))
		log.Info("deleting brand")

		err = brandDeleter.Delete(r.Context(), brandID)
		if err != nil {
			log.Error("failed to delete brand", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to delete brand"))
			return
		}

		log.Info("brand deleted", slog.Int("brandID", brandID))

		render.JSON(w, r, Response{
			Response: resp.OK(),
		})
	}
}
