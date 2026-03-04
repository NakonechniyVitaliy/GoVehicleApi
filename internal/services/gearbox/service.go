package service

import (
	"context"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	gearboxRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/gearbox"
	gearboxRequests "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/gearboxes"
)

type Service struct {
	repo       gearboxRepo.RepositoryInterface
	autoRiaKey string
}

func NewService(repository gearboxRepo.RepositoryInterface, key string) *Service {
	return &Service{
		repo:       repository,
		autoRiaKey: key,
	}
}

func (s Service) Refresh(ctx context.Context, autoRiaKey string) error {

	gearboxes, err := gearboxRequests.GetGearboxes(autoRiaKey)
	if err != nil {
		return err
	}
	for _, oneGearbox := range gearboxes {
		err = s.repo.InsertOrUpdate(ctx, oneGearbox)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s Service) GetAll(ctx context.Context) ([]models.Gearbox, error) {
	gearboxes, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return gearboxes, nil
}
