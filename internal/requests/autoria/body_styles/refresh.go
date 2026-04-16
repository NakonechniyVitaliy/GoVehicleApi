package body_styles

import (
	"net/http"

	"github.com/NakonechniyVitalii/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/requests/autoria"
	"github.com/go-chi/render"
)

func GetBodyStyles(key string) ([]models.BodyStyle, error) {

	url := autoria.GET_BODY_STYLES + key

	resp, err := http.Get(url)
	if err != nil {
		return nil, ErrBodyStylesFetch
	}
	defer resp.Body.Close()

	var bodyStyles []models.BodyStyle
	err = render.DecodeJSON(resp.Body, &bodyStyles)
	if err != nil {
		return nil, ErrDecodeBodyStyles
	}

	return bodyStyles, nil

}
