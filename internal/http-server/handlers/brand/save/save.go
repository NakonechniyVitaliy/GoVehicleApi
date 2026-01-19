package save

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	Brand models.Brand
}

type Response struct {
	resp.Response
}

type BrandSaver interface {
	NewBrand(ctx context.Context, brand models.Brand) error
}

func New(log *slog.Logger, brandSaver BrandSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.brand.save.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

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

		//brand := models.Brand{1, 2, 220, "TestEngName", 777, "Test", "TEst", 10}

		brand := models.Brand{
			Category: req.Brand.Category,
			Count:    req.Brand.Count,
			Country:  req.Brand.Country,
			EngName:  req.Brand.EngName,
			Marka:    req.Brand.Marka,
			Name:     req.Brand.Name,
			Slang:    req.Brand.Slang,
			Value:    req.Brand.Value,
		}

		log.Info("saving brand", slog.Any("brand", brand))

		err = brandSaver.NewBrand(r.Context(), brand)
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

		render.JSON(w, r, Response{
			Response: resp.OK(),
		})
	}
}
