package brand

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	response "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/brand"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type GetResponse struct {
	Response response.Response
	Brand    *models.Brand
}

// Get godoc
// @Summary      Отримати бренд за ID
// @Tags         brand
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "ID бренду"
// @Success      200  {object}  GetResponse
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /brand/{id} [get]
func Get(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.brand.get"))

		id64, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
		if err != nil {
			log.Error("failed to get body style ID", slog.String("error", err.Error()))
			response.RenderError(w, r, http.StatusBadRequest, "failed to get body style ID")
			return
		}

		brandID := uint16(id64)

		Brand, err := srv.GetByID(r.Context(), brandID)
		if errors.Is(err, service.ErrBrandNotFound) {
			response.RenderError(w, r, http.StatusNotFound, service.ErrBrandNotFound.Error())
			return
		}
		if err != nil {
			response.RenderError(w, r, http.StatusInternalServerError, service.ErrGetBrand.Error())
			return
		}

		render.JSON(w, r, GetResponse{
			Response: response.OK(),
			Brand:    Brand,
		})
	}
}
