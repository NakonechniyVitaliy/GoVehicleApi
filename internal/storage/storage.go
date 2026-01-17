package storage

import (
	"errors"
	"net/http"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
)

var (
	ErrBrandExists = errors.New("URL exists")
)

type Storage interface {
	NewBrand(brand models.Brand, r *http.Request) error
	RefreshBrands() error
}
