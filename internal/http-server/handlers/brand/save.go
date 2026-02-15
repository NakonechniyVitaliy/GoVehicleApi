package brand

import (
	"errors"
	"log/slog"
	"net/http"

	dto "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/brand"
	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/brand"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type SaveResponse struct {
	Response resp.Response
	Brand    *models.Brand
}

func New(log *slog.Logger, repository brand.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.brand.save.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req dto.SaveRequest
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("invalid JSON or wrong field types", slog.String("error", err.Error()))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("invalid JSON or wrong field types"))
			return
		}
		log.Info("request body decoded", slog.Any("request", req))

		err = req.Validate()
		if err != nil {
			log.Error("validation error", slog.String("error", err.Error()))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error(err.Error()))
			return
		}

		newBrand := req.Brand.ToModel()
		log.Info("saving brand", slog.Any("brand", newBrand))

		createdBrand, err := repository.Create(r.Context(), newBrand)
		if errors.Is(err, storage.ErrBrandExists) {
			log.Info("brand already exists", slog.String("brand", createdBrand.Name))
			render.JSON(w, r, resp.Error("Brand already exists"))
			return
		}

		if err != nil {
			log.Error("failed to save brand", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to save brand"))
			return
		}

		log.Info("brand saved", slog.String("brand", createdBrand.Name))

		render.JSON(w, r, SaveResponse{
			Response: resp.OK(),
			Brand:    createdBrand,
		})
	}
}
