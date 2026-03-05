package service

import (
	"context"
	"log/slog"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	gearboxRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/gearbox"
	gearboxRequests "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/gearboxes"
)

type Service struct {
	repo       gearboxRepo.RepositoryInterface
	log        *slog.Logger
	autoRiaKey string
}

func NewService(repository gearboxRepo.RepositoryInterface, logger *slog.Logger, key string) *Service {
	return &Service{
		repo:       repository,
		log:        logger,
		autoRiaKey: key,
	}
}

func (s Service) Refresh(ctx context.Context) error {

	gearboxes, err := gearboxRequests.GetGearboxes(s.autoRiaKey)
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
