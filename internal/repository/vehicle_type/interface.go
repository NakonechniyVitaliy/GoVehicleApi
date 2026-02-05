package vehicleType

import (
	"context"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
)

type Repository interface {
	Create(ctx context.Context, vehicleType models.VehicleType) error
	Delete(ctx context.Context, vehicleTypeID uint16) error
	GetByID(ctx context.Context, vehicleTypeID uint16) (*models.VehicleType, error)
	Update(ctx context.Context, vehicleType models.VehicleType) error
	GetAll(ctx context.Context) ([]models.VehicleType, error)
	InsertOrUpdate(ctx context.Context, vehicleType models.VehicleType) error
}
