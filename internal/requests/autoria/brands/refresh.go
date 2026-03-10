package brands

import (
	"net/http"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria"
	"github.com/go-chi/render"
)

func GetBrands(key string) ([]models.Brand, error) {

	url := autoria.GET_BRANDS + key

	resp, err := http.Get(url)
	if err != nil {
		return nil, ErrBrandsFetch
	}
	defer resp.Body.Close()

	var brands []models.Brand
	err = render.DecodeJSON(resp.Body, &brands)
	if err != nil {
		return nil, ErrDecodeBrands
	}

	return brands, nil

}
