package update

import (
	"context"
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
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

type BrandUpdater interface {
	UpdateBrand(ctx context.Context, brand models.Brand) error
}

func Update(log *slog.Logger, brandUpdater BrandUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.brand.update.Update"

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

		err = brandUpdater.UpdateBrand(r.Context(), brand)
		if err != nil {
			log.Error("failed to update brand", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to update brand"))
			return
		}

		render.JSON(w, r, Response{
			Response: resp.OK(),
		})

	}
}
