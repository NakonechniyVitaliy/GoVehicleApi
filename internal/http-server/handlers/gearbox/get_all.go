package gearbox

import (
	"log/slog"
	"net/http"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	Gearbox "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/gearbox"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type GetAllResponse struct {
	Response  resp.Response
	Gearboxes []models.Gearbox
}

func GetAll(log *slog.Logger, repository Gearbox.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.Gearbox.GetAll"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		log.Info("getting vehicle gearboxes")

		gearboxes, err := repository.GetAll(r.Context())
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
