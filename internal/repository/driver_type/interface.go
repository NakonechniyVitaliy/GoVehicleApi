package driver_type

import (
	"context"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
)

type Repository interface {
	GetAll(ctx context.Context) ([]models.DriverType, error)
	InsertOrUpdate(ctx context.Context, driverType models.DriverType) error
}
