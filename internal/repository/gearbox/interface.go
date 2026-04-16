package gearbox

import (
	"context"

	"github.com/NakonechniyVitalii/GoVehicleApi/internal/models"
)

type RepositoryInterface interface {
	GetByID(ctx context.Context, gearboxID uint16) (*models.Gearbox, error)
	GetAll(ctx context.Context) ([]models.Gearbox, error)
	InsertOrUpdate(ctx context.Context, gearbox models.Gearbox) error
}
