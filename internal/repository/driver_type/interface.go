package driver_type

import (
	"context"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
)

type RepositoryInterface interface {
	GetByID(ctx context.Context, driverTypeID uint16) (*models.DriverType, error)
	GetAll(ctx context.Context) ([]models.DriverType, error)
	InsertOrUpdate(ctx context.Context, driverType models.DriverType) error
}
