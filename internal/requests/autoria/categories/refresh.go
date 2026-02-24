package autoria

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria"
	"github.com/go-chi/render"
)

func GetCategories(key string) ([]models.Category, error) {

	url := autoria.GET_CATEGORIES + key

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetching Error %s", slog.String("error", err.Error()))
	}
	defer resp.Body.Close()

	var categories []models.Category
	err = render.DecodeJSON(resp.Body, &categories)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Autoria response body %s", slog.String("error", err.Error()))
	}

	return categories, nil

}
