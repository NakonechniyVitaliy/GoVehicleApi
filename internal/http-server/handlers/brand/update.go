package brand

import (
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/brand"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

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

		var req UpdateRequest
		err := render.DecodeJSON(r.Body, &req)
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

		updatedBrand := models.Brand{
			ID:         req.Brand.ID,
			CategoryID: req.Brand.CategoryID,
			Count:      req.Brand.Count,
			CountryID:  req.Brand.CountryID,
			EngName:    req.Brand.EngName,
			MarkaID:    req.Brand.MarkaID,
			Name:       req.Brand.Name,
			Slang:      req.Brand.Slang,
			Value:      req.Brand.Value,
		}

		err = repository.Update(r.Context(), updatedBrand)
		if err != nil {
			log.Error("failed to update brand", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to update brand"))
			return
		}

		render.JSON(w, r, resp.OK())

	}
}
