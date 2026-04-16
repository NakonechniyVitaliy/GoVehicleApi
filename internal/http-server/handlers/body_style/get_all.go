package body_style

import (
	"errors"
	"log/slog"
	"net/http"

	response "github.com/NakonechniyVitalii/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/body_style"
	"github.com/go-chi/render"
)

type GetAllResponse struct {
	Response   response.Response
	BodyStyles []models.BodyStyle
}

// GetAll godoc
// @Summary      Список стилів кузова
// @Tags         body-style
// @Security     BearerAuth
// @Produce      json
// @Success      200  {object}  GetAllResponse
// @Failure      500  {object}  response.Response
// @Router       /body-style/all [get]
func GetAll(log *slog.Logger, svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.body_style.get_all"))

		bodyStyles, err := svc.GetAll(r.Context())

		if errors.Is(err, service.ErrGetBodyStyles) {
			response.RenderError(w, r, http.StatusInternalServerError, service.ErrGetBodyStyles.Error())
			return
		}
		render.JSON(w, r, GetAllResponse{
			Response:   response.OK(),
			BodyStyles: bodyStyles,
		})
	}
}
