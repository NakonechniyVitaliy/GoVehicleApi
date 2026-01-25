package get

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Request struct {
	VehicleTypeID int
}

type ResponseGet struct {
	Response    resp.Response
	VehicleType *models.VehicleType
}

type ResponseGetAll struct {
	Response     resp.Response
	VehicleTypes []models.VehicleType
}

type VehicleTypeGetter interface {
	GetByID(ctx context.Context, VehicleTypeID int) (*models.VehicleType, error)
	GetAll(context.Context) ([]models.VehicleType, error)
}

func Get(log *slog.Logger, VehicleTypeGetter VehicleTypeGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.VehicleType.get.Get"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		VehicleTypeID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Error("failed to get vehicle type ID", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get vehicle type ID"))
			return
		}

		log.Info("ID retrieved successfully", slog.Any("Vehicle type ID", VehicleTypeID))
		log.Info("getting VehicleType")

		VehicleType, err := VehicleTypeGetter.GetByID(r.Context(), VehicleTypeID)
		if err != nil {
			log.Error("failed to get vehicle type", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get vehicle type"))
			return
		}

		render.JSON(w, r, ResponseGet{
			Response:    resp.OK(),
			VehicleType: VehicleType,
		})
	}
}

func GetAll(log *slog.Logger, VehicleTypeGetter VehicleTypeGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.VehicleType.get.Get"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		log.Info("getting VehicleTypes")

		VehicleTypes, err := VehicleTypeGetter.GetAll(r.Context())
		if err != nil {
			log.Error("failed to get vehicle type", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get vehicle type"))
			return
		}

		render.JSON(w, r, ResponseGetAll{
			Response:     resp.OK(),
			VehicleTypes: VehicleTypes,
		})
	}
}
