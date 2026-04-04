package category

import (
	"errors"
	"log/slog"
	"net/http"

	response "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/category"
	"github.com/go-chi/render"
)

type GetAllResponse struct {
	Response   response.Response
	Categories []models.Category
}

// GetAll godoc
// @Summary      Список категорій
// @Tags         category
// @Security     BearerAuth
// @Produce      json
// @Success      200  {object}  GetAllResponse
// @Failure      500  {object}  response.Response
// @Router       /category/all [get]
func GetAll(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.category.get_all"))

		categories, err := srv.GetAll(r.Context())

		if errors.Is(err, service.ErrGetCategories) {
			response.RenderError(w, r, http.StatusInternalServerError, service.ErrGetCategories.Error())
			return
		}
		render.JSON(w, r, GetAllResponse{
			Response:   response.OK(),
			Categories: categories,
		})
	}
}
