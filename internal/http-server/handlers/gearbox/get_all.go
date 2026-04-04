package gearbox

import (
	"log/slog"
	"net/http"

	response "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/gearbox"
	"github.com/go-chi/render"
)

type GetAllResponse struct {
	Response  response.Response
	Gearboxes []models.Gearbox
}

// GetAll godoc
// @Summary      Список коробок передач
// @Tags         gearbox
// @Security     BearerAuth
// @Produce      json
// @Success      200  {object}  GetAllResponse
// @Failure      500  {object}  response.Response
// @Router       /gearbox/all [get]
func GetAll(log *slog.Logger, service *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.gearbox.get_all"

		log = log.With(slog.String("op", op))

		log.Info("getting vehicle gearboxes")

		gearboxes, err := service.GetAll(r.Context())
		if err != nil {
			log.Error("failed to get vehicle gearboxes", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Failed to get vehicle gearboxes"))
			return
		}

		render.JSON(w, r, GetAllResponse{
			Response:  response.OK(),
			Gearboxes: gearboxes,
		})
	}
}
