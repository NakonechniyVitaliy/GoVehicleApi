package autoria

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria"
	"github.com/go-chi/render"
)

func GetVehicleCategories(key string) ([]models.VehicleCategory, error) {

	url := autoria.GET_VEHICLE_CATEGORIES + key

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetching Error %s", slog.String("error", err.Error()))
	}
	defer resp.Body.Close()

	var vehicleCategories []models.VehicleCategory
	err = render.DecodeJSON(resp.Body, &vehicleCategories)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Autoria response body %s", slog.String("error", err.Error()))
	}

	return vehicleCategories, nil

}
