package gearbox

import (
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/gearbox"
	"github.com/go-chi/render"
)

type GetAllResponse struct {
	Response  resp.Response
	Gearboxes []models.Gearbox
}

func GetAll(log *slog.Logger, service *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.gearbox.get_all"

		log = log.With(slog.String("op", op))

		log.Info("getting vehicle gearboxes")

		gearboxes, err := service.GetAll(r.Context())
		if err != nil {
			log.Error("failed to get vehicle gearboxes", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get vehicle gearboxes"))
			return
		}

		render.JSON(w, r, GetAllResponse{
			Response:  resp.OK(),
			Gearboxes: gearboxes,
		})
	}
}
