package body_style

import (
	"log/slog"
	"net/http"
	"strconv"

	dto "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/body_style"
	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	bodyStyle "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/body_style"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type UpdateResponse struct {
	Response  resp.Response
	BodyStyle *models.BodyStyle
}

func Update(log *slog.Logger, repository bodyStyle.RepositoryInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.bodyStyle.update"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		id64, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
		if err != nil {
			log.Error("failed to get brand ID", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get brand ID"))
			return
		}
		bodyStyleID := uint16(id64)
		log.Info("ID retrieved successfully", slog.Any("bodyStyleID", bodyStyleID))

		var req dto.UpdateRequest
		err = render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("invalid JSON or wrong field types", slog.String("error", err.Error()))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error("invalid JSON or wrong field types"))
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

		bodyStyleFromRequest := req.BodyStyle.ToModel()

		updatedBodyStyle, err := repository.Update(r.Context(), bodyStyleFromRequest, bodyStyleID)
		if err != nil {
			log.Error("failed to update bodyStyle", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to update bodyStyle"))
			return
		}

		render.JSON(w, r, UpdateResponse{
			Response:  resp.OK(),
			BodyStyle: updatedBodyStyle,
		})
	}
}
