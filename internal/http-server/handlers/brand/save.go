package brand

import (
	"errors"
	"log/slog"
	"net/http"

	dtoErrors "github.com/NakonechniyVitalii/GoVehicleApi/internal/http-server/dto/_errors"
	dto "github.com/NakonechniyVitalii/GoVehicleApi/internal/http-server/dto/brand"
	response "github.com/NakonechniyVitalii/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/brand"
	"github.com/go-chi/render"
)

type SaveResponse struct {
	Response response.Response
	Brand    *models.Brand
}

// New godoc
// @Summary      Створити бренд
// @Tags         brand
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body  body      BrandPayload  true  "Дані бренду"
// @Success      200   {object}  SaveResponse
// @Failure      400   {object}  response.Response
// @Failure      409   {object}  response.Response
// @Failure      500   {object}  response.Response
// @Router       /brand/ [post]
func New(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.brand.save"))

		var req dto.SaveRequest
		err := render.DecodeJSON(r.Body, &req)
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

		mappedBrand := req.Brand.ToModel()
		createdBrand, err := srv.Save(r.Context(), mappedBrand)

		if errors.Is(err, service.ErrBrandExists) {
			response.RenderError(w, r, http.StatusConflict, service.ErrBrandExists.Error())
			return
		}
		if errors.Is(err, service.ErrSaveBrand) {
			response.RenderError(w, r, http.StatusInternalServerError, service.ErrSaveBrand.Error())
			return
		}

		render.JSON(w, r, SaveResponse{
			Response: response.OK(),
			Brand:    createdBrand,
		})
	}
}
