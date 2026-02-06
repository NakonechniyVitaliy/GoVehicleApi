package brand

import (
	"context"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
)

type Repository interface {
	Create(ctx context.Context, brand models.Brand) (*models.Brand, error)
	Delete(ctx context.Context, brandID uint16) error
	GetByID(ctx context.Context, brandID uint16) (*models.Brand, error)
	Update(ctx context.Context, brand models.Brand) error
	GetAll(ctx context.Context) ([]models.Brand, error)
	InsertOrUpdate(ctx context.Context, brand models.Brand) error
}
