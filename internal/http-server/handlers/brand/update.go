package brand

import (
	"log/slog"
	"net/http"
	"strconv"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/brand"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type UpdateResponse struct {
	Response resp.Response
	Brand    *models.Brand
}
type UpdateRequest struct {
	Brand models.Brand
}

func Update(log *slog.Logger, repository brand.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.brand.update.Update"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		id64, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
		if err != nil {
			log.Error("failed to get brand ID", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get brand ID"))
			return
		}
		brandID := uint16(id64)
		log.Info("ID retrieved successfully", slog.Any("brandID", brandID))

		var req UpdateRequest
		err = render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to decode request"))
			return
		}
		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			log.Error("invalid request", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Invalid request"))
			return
		}

		updatedBrand, err := repository.Update(r.Context(), req.Brand, brandID)
		if err != nil {
			log.Error("failed to update brand", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to update brand"))
			return
		}

		render.JSON(w, r, UpdateResponse{
			Response: resp.OK(),
			Brand:    updatedBrand,
		})

	}
}
