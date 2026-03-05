package service

import (
	"context"
	"log/slog"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	driverTypeRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/driver_type"
	driverTypeRequests "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/driver_types"
)

type Service struct {
	repo       driverTypeRepo.RepositoryInterface
	log        *slog.Logger
	autoRiaKey string
}

func NewService(repository driverTypeRepo.RepositoryInterface, logger *slog.Logger, key string) *Service {
	return &Service{
		repo:       repository,
		log:        logger,
		autoRiaKey: key,
	}
}

func (s Service) Refresh(ctx context.Context) error {

	driverTypes, err := driverTypeRequests.GetDriverTypes(s.autoRiaKey)
	if err != nil {
		return err
	}
	for _, oneDriverType := range driverTypes {
		err = s.repo.InsertOrUpdate(ctx, oneDriverType)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s Service) GetAll(ctx context.Context) ([]models.DriverType, error) {
	driverTypes, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return driverTypes, nil
}
