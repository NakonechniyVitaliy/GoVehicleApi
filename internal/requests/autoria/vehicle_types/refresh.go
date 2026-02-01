package autoria

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria"
	"github.com/go-chi/render"
)

func GetVehicleTypes(key string) ([]models.VehicleType, error) {

	url := autoria.GET_VEHICLE_TYPES + key

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetching Error %s", slog.String("error", err.Error()))
	}
	defer resp.Body.Close()

	var vehicleTypes []models.VehicleType
	err = render.DecodeJSON(resp.Body, &vehicleTypes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Autoria response body %s", slog.String("error", err.Error()))
	}

	return vehicleTypes, nil

}
