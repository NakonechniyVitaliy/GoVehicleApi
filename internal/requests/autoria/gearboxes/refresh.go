package gearboxes

import (
	"net/http"

	"github.com/NakonechniyVitalii/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/requests/autoria"
	"github.com/go-chi/render"
)

func GetGearboxes(key string) ([]models.Gearbox, error) {

	url := autoria.GET_GEARBOXES + key

	resp, err := http.Get(url)
	if err != nil {
		return nil, ErrGearboxesFetch
	}
	defer resp.Body.Close()

	var gearboxes []models.Gearbox
	err = render.DecodeJSON(resp.Body, &gearboxes)
	if err != nil {
		return nil, ErrDecodeGearboxes
	}

	return gearboxes, nil

}
