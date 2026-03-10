package autoria

import (
	"net/http"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria"
	"github.com/go-chi/render"
)

func GetDriverTypes(key string) ([]models.DriverType, error) {

	url := autoria.GET_DRIVER_TYPES + key

	resp, err := http.Get(url)
	if err != nil {
		return nil, ErrDriverTypesFetch
	}
	defer resp.Body.Close()

	var driverTypes []models.DriverType
	err = render.DecodeJSON(resp.Body, &driverTypes)
	if err != nil {
		return nil, ErrDecodeDriverTypes
	}

	return driverTypes, nil

}
