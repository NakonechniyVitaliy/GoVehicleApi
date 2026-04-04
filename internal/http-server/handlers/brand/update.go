package brand

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	dtoErrors "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/_errors"
	dto "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/brand"
	response "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/brand"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type UpdateResponse struct {
	Response response.Response
	Brand    *models.Brand
}

// Update godoc
// @Summary      Оновити бренд
// @Tags         brand
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id    path      int           true  "ID бренду"
// @Param        body  body      BrandPayload  true  "Нові дані"
// @Success      200   {object}  UpdateResponse
// @Failure      400   {object}  response.Response
// @Failure      404   {object}  response.Response
// @Failure      500   {object}  response.Response
// @Router       /brand/{id} [put]
func Update(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.brand.update"))

		id64, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
		if err != nil {
			log.Error("failed to get brand ID", slog.String("error", err.Error()))
			response.RenderError(w, r, http.StatusBadRequest, "failed to get brand ID")
			return
		}

		brandID := uint16(id64)
		var req dto.UpdateRequest

		err = render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error(dtoErrors.InvalidJSONorWrongFieldType, slog.String("error", err.Error()))
			response.RenderError(w, r, http.StatusBadRequest, dtoErrors.InvalidJSONorWrongFieldType)
			return
		}

		err = req.Validate()
		if err != nil {
			log.Error("validation error", slog.String("error", err.Error()))
			response.RenderError(w, r, http.StatusBadRequest, err.Error())
			return
		}

		brandFromRequest := req.Brand.ToModel()
		updatedBrand, err := srv.Update(r.Context(), brandFromRequest, brandID)

		if errors.Is(err, service.ErrBrandNotFound) {
			response.RenderError(w, r, http.StatusNotFound, service.ErrBrandNotFound.Error())
			return
		}
		if errors.Is(err, service.ErrUpdateBrand) {
			response.RenderError(w, r, http.StatusInternalServerError, service.ErrUpdateBrand.Error())
			return
		}

		render.JSON(w, r, UpdateResponse{
			Response: response.OK(),
			Brand:    updatedBrand,
		})

	}
}
