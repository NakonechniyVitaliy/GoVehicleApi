package get

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Request struct {
	brandID int
}

type ResponseGet struct {
	Response resp.Response
	Brand    *models.Brand
}

type ResponseGetAll struct {
	Response resp.Response
	Brands   []models.Brand
}

type BrandGetter interface {
	GetByID(ctx context.Context, brandID int) (*models.Brand, error)
	GetAll(context.Context) ([]models.Brand, error)
}

func Get(log *slog.Logger, brandGetter BrandGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.brand.get.Get"

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
		log.Info("getting brand")

		brand, err := brandGetter.GetByID(r.Context(), brandID)
		if err != nil {
			log.Error("failed to get brand", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get brand"))
			return
		}

		render.JSON(w, r, ResponseGet{
			Response: resp.OK(),
			Brand:    brand,
		})
	}
}

func GetAll(log *slog.Logger, brandGetter BrandGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.brand.get.Get"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		log.Info("getting brands")

		brands, err := brandGetter.GetAll(r.Context())
		if err != nil {
			log.Error("failed to get brand", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get brand"))
			return
		}

		render.JSON(w, r, ResponseGetAll{
			Response: resp.OK(),
			Brands:   brands,
		})
	}
}
