package category

import (
	"errors"
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/category"
	"github.com/go-chi/render"
)

type GetAllResponse struct {
	Response   resp.Response
	Categories []models.Category
}

func GetAll(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.category.get_all"))

		categories, err := srv.GetAll(r.Context())

		if errors.Is(err, service.ErrGetCategories) {
			resp.RenderError(w, r, http.StatusInternalServerError, service.ErrGetCategories.Error())
			return
		}
		render.JSON(w, r, GetAllResponse{
			Response:   resp.OK(),
			Categories: categories,
		})
	}
}
