package brands

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria"
	"github.com/go-chi/render"
)

func GetBrands(key string) ([]models.Brand, error) {

	url := autoria.GET_BRANDS + key

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetching Error %s", slog.String("error", err.Error()))
	}
	defer resp.Body.Close()

	var brands []models.Brand
	err = render.DecodeJSON(resp.Body, &brands)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Autoria response body %s", slog.String("error", err.Error()))
	}

	return brands, nil

}
