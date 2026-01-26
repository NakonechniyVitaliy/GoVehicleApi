package brands

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/go-chi/render"
)

const (
	GET_BRANDS = "https://developers.ria.com/auto/new/marks?category_id=1&api_key="
)

func GetBrands(key string) ([]models.Brand, error) {

	url := GET_BRANDS + key

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
