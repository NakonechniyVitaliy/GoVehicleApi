package driver_type

import (
	"context"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
)

type Repository interface {
	GetAll(ctx context.Context) ([]models.Gearbox, error)
	InsertOrUpdate(ctx context.Context, gearbox models.Gearbox) error
}
