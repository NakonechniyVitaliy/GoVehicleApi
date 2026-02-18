package vehicle_category

import (
	"context"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
)

type RepositoryInterface interface {
	GetAll(ctx context.Context) ([]models.VehicleCategory, error)
	InsertOrUpdate(ctx context.Context, vehicleCategory models.VehicleCategory) error
}
