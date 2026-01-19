package storage

import (
	"context"
	"errors"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
)

var (
	ErrBrandExists = errors.New("URL exists")
)

type Storage interface {
	NewBrand(brand models.Brand, ctx context.Context) error
	RefreshBrands() error
	DeleteBrand(brandID int, ctx context.Context) error
}
