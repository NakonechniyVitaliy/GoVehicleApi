package body_style

import (
	"context"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
)

type RepositoryInterface interface {
	Create(ctx context.Context, bodyStyle models.BodyStyle) error
	Delete(ctx context.Context, bodyStyleID uint16) error
	GetByID(ctx context.Context, bodyStyleID uint16) (*models.BodyStyle, error)
	Update(ctx context.Context, bodyStyle models.BodyStyle) error
	GetAll(ctx context.Context) ([]models.BodyStyle, error)
	InsertOrUpdate(ctx context.Context, bodyStyle models.BodyStyle) error
}
