package brand

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/brand"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type GetResponse struct {
	Response resp.Response
	Brand    *models.Brand
}

func Get(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.brand.get"))

		id64, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
		if err != nil {
			log.Error("failed to get body style ID", slog.String("error", err.Error()))
			resp.RenderError(w, r, http.StatusBadRequest, "failed to get body style ID")
			return
		}

		brandID := uint16(id64)

		Brand, err := srv.GetByID(r.Context(), brandID)
		if errors.Is(err, service.ErrBrandNotFound) {
			resp.RenderError(w, r, http.StatusNotFound, service.ErrBrandNotFound.Error())
			return
		}
		if err != nil {
			resp.RenderError(w, r, http.StatusInternalServerError, service.ErrGetBrand.Error())
			return
		}

		render.JSON(w, r, GetResponse{
			Response: resp.OK(),
			Brand:    Brand,
		})
	}
}
