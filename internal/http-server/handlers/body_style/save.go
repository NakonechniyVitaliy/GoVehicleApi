package body_style

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/_errors"
	dto "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/body_style"
	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	bodyStyle "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/body_style"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage"
	"github.com/go-chi/render"
)

type SaveResponse struct {
	Response  resp.Response
	BodyStyle *models.BodyStyle
}

func New(log *slog.Logger, repository bodyStyle.RepositoryInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		const op = "handlers.bodyStyle.new"
		log = log.With(slog.String("op", op))

		var req dto.SaveRequest
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error(_errors.InvalidJSONorWrongFieldType, slog.String("error", err.Error()))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error(_errors.InvalidJSONorWrongFieldType))
			return
		}
		log.Info("request body decoded", slog.Any("request", req))

		err = req.Validate()
		if err != nil {
			log.Error("validation error", slog.String("error", err.Error()))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error(err.Error()))
			return
		}

		newBS := req.BodyStyle.ToModel()
		log.Info("saving body style", slog.Any("body style", newBS))

		createdBS, err := repository.Create(r.Context(), newBS)
		if errors.Is(err, storage.ErrBodyStyleExists) {
			log.Info("bodyStyle already exists", slog.String("body style", createdBS.Name))
			render.JSON(w, r, resp.Error("body style already exists"))
			return
		}

		if err != nil {
			log.Error("failed to save body style", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to save body style"))
			return
		}
		log.Info("bodyStyle saved", slog.String("body style", createdBS.Name))

		render.JSON(w, r, SaveResponse{
			Response:  resp.OK(),
			BodyStyle: createdBS,
		})
	}
}
