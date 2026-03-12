package brand

import (
	"errors"
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/brand"
	"github.com/go-chi/render"
)

type GetAllResponse struct {
	Response resp.Response
	Brands   []models.Brand
}

func GetAll(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.brand.get"))

		brands, err := srv.GetAll(r.Context())

		if errors.Is(err, service.ErrGetBrands) {
			resp.RenderError(w, r, http.StatusInternalServerError, service.ErrGetBrands.Error())
			return
		}
		render.JSON(w, r, GetAllResponse{
			Response: resp.OK(),
			Brands:   brands,
		})
	}
}
