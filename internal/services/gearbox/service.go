package gearbox

import (
	"context"
	"errors"
	"log/slog"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	gearboxRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/gearbox"
	requests "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/gearboxes"
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
	log := s.log.With(slog.String("op", "services.gearbox.refresh"))

	gearboxes, err := requests.GetGearboxes(s.autoRiaKey)
	if err != nil {
		switch {
		case errors.Is(err, requests.ErrGearboxesFetch):
			log.Error(ErrGearboxesFetch.Error(), slog.String("error", err.Error()))

		case errors.Is(err, requests.ErrDecodeGearboxes):
			log.Error(ErrDecodeGearboxes.Error(), slog.String("error", err.Error()))
		}
		return ErrRefreshGearboxes
	}

	for _, oneGearbox := range gearboxes {
		err = s.repo.InsertOrUpdate(ctx, oneGearbox)
		if err != nil {
			return ErrRefreshGearboxes
		}
	}
	return nil
}

func (s Service) GetAll(ctx context.Context) ([]models.Gearbox, error) {
	log := s.log.With(slog.String("op", "services.gearbox.get_all"))

	gearboxes, err := s.repo.GetAll(ctx)
	if err != nil {
		log.Error(ErrGetGearboxes.Error(), slog.String("error", err.Error()))
		return nil, ErrGetGearboxes
	}
	return gearboxes, nil
}
