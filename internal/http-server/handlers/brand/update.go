package brand

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	dtoErrors "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/_errors"
	dto "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/brand"
	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/brand"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type UpdateResponse struct {
	Response resp.Response
	Brand    *models.Brand
}

func Update(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.brand.update"))

		id64, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
		if err != nil {
			log.Error("failed to get brand ID", slog.String("error", err.Error()))
			resp.RenderError(w, r, http.StatusBadRequest, "failed to get brand ID")
			return
		}

		brandID := uint16(id64)
		var req dto.UpdateRequest

		err = render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error(dtoErrors.InvalidJSONorWrongFieldType, slog.String("error", err.Error()))
			resp.RenderError(w, r, http.StatusBadRequest, dtoErrors.InvalidJSONorWrongFieldType)
			return
		}

		err = req.Validate()
		if err != nil {
			log.Error("validation error", slog.String("error", err.Error()))
			resp.RenderError(w, r, http.StatusBadRequest, err.Error())
			return
		}

		brandFromRequest := req.Brand.ToModel()
		updatedBrand, err := srv.Update(r.Context(), brandFromRequest, brandID)

		if errors.Is(err, service.ErrBrandNotFound) {
			resp.RenderError(w, r, http.StatusNotFound, service.ErrBrandNotFound.Error())
			return
		}
		if errors.Is(err, service.ErrUpdateBrand) {
			resp.RenderError(w, r, http.StatusInternalServerError, service.ErrUpdateBrand.Error())
			return
		}

		render.JSON(w, r, UpdateResponse{
			Response: resp.OK(),
			Brand:    updatedBrand,
		})

	}
}
