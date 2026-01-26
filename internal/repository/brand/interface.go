package brand

import (
	"context"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
)

type Repository interface {
	Create(ctx context.Context, brand models.Brand) error
	Delete(ctx context.Context, brandID int) error
	GetByID(ctx context.Context, brandID int) (*models.Brand, error)
	Update(ctx context.Context, brand models.Brand) error
	GetAll(ctx context.Context) ([]models.Brand, error)
	InsertOrUpdate(ctx context.Context, brands []models.Brand) error
}
