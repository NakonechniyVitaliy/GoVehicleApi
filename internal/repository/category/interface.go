package category

import (
	"context"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
)

type RepositoryInterface interface {
	GetAll(ctx context.Context) ([]models.Category, error)
	InsertOrUpdate(ctx context.Context, category models.Category) error
}
