package brand

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/brand"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type GetResponse struct {
	Response resp.Response
	Brand    *models.Brand
}

func Get(log *slog.Logger, service *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.brand.get"

		log = log.With(slog.String("op", op))

		id64, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
		if err != nil {
			log.Error("failed to get brand ID", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get brand ID"))
			return
		}
		brandID := uint16(id64)
		log.Info("ID retrieved successfully", slog.Any("brandID", brandID))

		log.Info("getting brand")
		requestedBrand, err := service.GetByID(r.Context(), brandID)
		if err != nil {
			log.Error("failed to get brand", slog.String("error", err.Error()))

			if errors.Is(err, storage.ErrBrandNotFound) {
				render.Status(r, http.StatusNotFound)
				render.JSON(w, r, resp.Error("brand not found"))
				return
			}

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, resp.Error("Failed to get brand"))
			return
		}

		render.JSON(w, r, GetResponse{
			Response: resp.OK(),
			Brand:    requestedBrand,
		})
	}
}
