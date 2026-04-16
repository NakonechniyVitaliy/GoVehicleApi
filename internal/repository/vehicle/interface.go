package vehicle

import (
	"context"

	"github.com/NakonechniyVitalii/GoVehicleApi/internal/http-server/dto/vehicle/filter"
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/models"
)

type RepositoryInterface interface {
	Create(ctx context.Context, vehicle models.Vehicle) (*models.Vehicle, error)
	Delete(ctx context.Context, vehicleID uint16) error
	GetByID(ctx context.Context, vehicleID uint16) (*models.Vehicle, error)
	Update(ctx context.Context, vehicle models.Vehicle, vehicleID uint16) (*models.Vehicle, error)
	GetAll(ctx context.Context) ([]models.Vehicle, error)
	GetByPage(ctx context.Context, f filter.Filter) ([]models.Vehicle, error)
}
