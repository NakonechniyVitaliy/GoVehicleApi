package storage

import (
	"context"
	"errors"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
)

var (
	ErrBrandExists   = errors.New("URL exists")
	ErrBrandNotFound = errors.New("brand not found")
)

type Storage interface {
	NewBrand(ctx context.Context, brand models.Brand) error
	DeleteBrand(ctx context.Context, brandID int) error
	GetBrand(ctx context.Context, brandID int) (*models.Brand, error)
	UpdateBrand(ctx context.Context, brand models.Brand) error
	GetAllBrands(ctx context.Context) ([]models.Brand, error)
	RefreshBrands() error
}
