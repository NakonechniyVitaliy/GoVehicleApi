package body_style

import (
	"errors"
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/body_style"
	"github.com/go-chi/render"
)

type GetAllResponse struct {
	Response   resp.Response
	BodyStyles []models.BodyStyle
}

func GetAll(log *slog.Logger, svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.body_style.get_all"))

		bodyStyles, err := svc.GetAll(r.Context())

		if errors.Is(err, service.ErrGetBodyStyles) {
			resp.RenderError(w, r, http.StatusInternalServerError, service.ErrGetBodyStyles.Error())
			return
		}
		render.JSON(w, r, GetAllResponse{
			Response:   resp.OK(),
			BodyStyles: bodyStyles,
		})
	}
}
