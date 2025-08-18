package save

import (
	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"log/slog"
	"net/http"
)

type Request struct {
	URL string `json:"url" validate:"required,url"`
}

type Responce struct {
	resp.Response
}

type newBrand interface {
	newBrand(brandToSave models.Brand) error
}

func New(log *slog.Logger, brandToSave models.Brand) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
