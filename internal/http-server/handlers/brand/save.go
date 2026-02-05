package brand

import (
	"errors"
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/brand"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type SaveRequest struct {
	Brand models.Brand
}

func New(log *slog.Logger, repository brand.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.brand.save.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req SaveRequest

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

		newBrand := models.Brand{
			CategoryID: req.Brand.CategoryID,
			Count:      req.Brand.Count,
			CountryID:  req.Brand.CountryID,
			EngName:    req.Brand.EngName,
			MarkaID:    req.Brand.MarkaID,
			Name:       req.Brand.Name,
			Slang:      req.Brand.Slang,
			Value:      req.Brand.Value,
		}

		log.Info("saving brand", slog.Any("brand", newBrand))

		err = repository.Create(r.Context(), newBrand)
		if errors.Is(err, storage.ErrBrandExists) {
			log.Info("brand already exists", slog.String("brand", req.Brand.Name))
			render.JSON(w, r, resp.Error("Brand already exists"))
			return
		}

		if err != nil {
			log.Error("failed to save brand", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to save brand"))
			return
		}

		log.Info("brand saved", slog.String("brand", req.Brand.Name))

		render.JSON(w, r, resp.OK())
	}
}
