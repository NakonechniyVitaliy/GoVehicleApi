package user

import (
	"context"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
)

type RepositoryInterface interface {
	GetByLogin(ctx context.Context, login string) (*models.User, error)
	Create(ctx context.Context, user models.User) error
}
