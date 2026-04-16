package category

import (
	"context"

	"github.com/NakonechniyVitalii/GoVehicleApi/internal/models"
)

type RepositoryInterface interface {
	GetByID(ctx context.Context, categoryID uint16) (*models.Category, error)
	GetAll(ctx context.Context) ([]models.Category, error)
	InsertOrUpdate(ctx context.Context, category models.Category) error
}
