package save

import (
	"errors"
	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
)

type Request struct {
	Brand models.Brand
}

type Responce struct {
	resp.Response
}

type BrandSaver interface {
	SaveBrand(brand models.Brand) error
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

		err = brandSaver.SaveBrand(req.Brand)
		if errors.Is("brand already exists", slog.String("brand", req.Brand) {}


	}
}
