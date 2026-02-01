package vehicleType

import (
	"context"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
)

type Repository interface {
	Create(ctx context.Context, vehicleType models.VehicleType) error
	Delete(ctx context.Context, vehicleTypeID int) error
	GetByID(ctx context.Context, vehicleTypeID int) (*models.VehicleType, error)
	Update(ctx context.Context, vehicleType models.VehicleType) error
	GetAll(ctx context.Context) ([]models.VehicleType, error)
	InsertOrUpdate(ctx context.Context, vehicleType models.VehicleType) error
}
