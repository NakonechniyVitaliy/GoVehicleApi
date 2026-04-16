package brand

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	response "github.com/NakonechniyVitalii/GoVehicleApi/internal/lib/api/response"
	service "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/brand"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// Delete godoc
// @Summary      Видалити бренд
// @Tags         brand
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "ID бренду"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /brand/{id} [delete]
func Delete(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.brand.delete"))

		id64, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
		if err != nil {
			log.Error("failed to get brand ID", slog.String("error", err.Error()))
			response.RenderError(w, r, http.StatusBadRequest, "failed to get body style ID")
			return
		}
		brandID := uint16(id64)

		err = srv.Delete(r.Context(), brandID)
		if errors.Is(err, service.ErrBrandNotFound) {
			response.RenderError(w, r, http.StatusNotFound, service.ErrBrandNotFound.Error())
			return
		}
		if err != nil {
			response.RenderError(w, r, http.StatusInternalServerError, service.ErrGetBrand.Error())
			return
		}
		render.JSON(w, r, response.OK())
	}
}
