package vehicle

import (
	"log/slog"
	"net/http"
	"strconv"

	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/body_style"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/brand"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/category"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/driver_type"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/gearbox"
	vehicleService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/vehicle"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func GetExpanded(
	log *slog.Logger,
	vRepo vehicle.RepositoryInterface,
	bRepo brand.RepositoryInterface,
	bsRepo body_style.RepositoryInterface,
	cRepo category.RepositoryInterface,
	dRepo driver_type.RepositoryInterface,
	gRepo gearbox.RepositoryInterface,
	) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.vehicle.get_expanded"

		log = log.With(slog.String("op", op))

		id64, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
		if err != nil {
			log.Error("failed to get vehicle ID", slog.String("error", err.Error()))
			render.JSON(w, r, resp.Error("Failed to get vehicle ID"))
			return
		}
		vehicleID := uint16(id64)
		log.Info("ID retrieved successfully", slog.Any("vehicleID", vehicleID))


		log.Info("getting expanded vehicle")
		expandedVehicle, err := vehicleService.GetExpanded(r.Context(), vehicleID, vRepo, bRepo,
			bsRepo, cRepo, dRepo, gRepo)


		//requestedVehicle, err := repository.GetByID(r.Context(), vehicleID)
		//if err != nil {
		//	log.Error("failed to get vehicle", slog.String("error", err.Error()))
		//
		//	if errors.Is(err, storage.ErrVehicleNotFound) {
		//		render.Status(r, http.StatusNotFound)
		//		render.JSON(w, r, resp.Error("vehicle not found"))
		//		return
		//	}
		//
		//	render.Status(r, http.StatusInternalServerError)
		//	render.JSON(w, r, resp.Error("Failed to get vehicle"))
		//	return
		}

		render.JSON(w, r, GetResponse{
			Response: resp.OK(),
			Vehicle:  requestedVehicle,
		})
	}
}
