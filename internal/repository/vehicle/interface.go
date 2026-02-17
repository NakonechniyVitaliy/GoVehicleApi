package vehicle

import (
	"context"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
)

type Repository interface {
	Create(ctx context.Context, vehicle models.Vehicle) (*models.Vehicle, error)
	Delete(ctx context.Context, vehicleID uint16) error
	GetByID(ctx context.Context, vehicleID uint16) (*models.Vehicle, error)
	Update(ctx context.Context, vehicle models.Vehicle, vehicleID uint16) (*models.Vehicle, error)
	GetAll(ctx context.Context) ([]models.Vehicle, error)
}
