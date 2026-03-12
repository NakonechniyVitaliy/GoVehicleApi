package brand

import (
	"errors"
	"log/slog"
	"net/http"

	dtoErrors "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/_errors"
	dto "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/brand"
	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/brand"
	"github.com/go-chi/render"
)

type SaveResponse struct {
	Response resp.Response
	Brand    *models.Brand
}

func New(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.brand.save"))

		var req dto.SaveRequest
		err := render.DecodeJSON(r.Body, &req)
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

		mappedBrand := req.Brand.ToModel()
		createdBrand, err := srv.Save(r.Context(), mappedBrand)

		if errors.Is(err, service.ErrBrandExists) {
			resp.RenderError(w, r, http.StatusConflict, service.ErrBrandExists.Error())
			return
		}
		if errors.Is(err, service.ErrSaveBrand) {
			resp.RenderError(w, r, http.StatusInternalServerError, service.ErrSaveBrand.Error())
			return
		}

		render.JSON(w, r, SaveResponse{
			Response: resp.OK(),
			Brand:    createdBrand,
		})
	}
}
