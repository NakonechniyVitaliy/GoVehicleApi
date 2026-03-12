package driver_type

import (
	"context"
	"errors"
	"log/slog"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	driverTypeRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/driver_type"
	requests "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/driver_types"
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
	log := s.log.With(slog.String("op", "services.driver_type.refresh"))

	driverTypes, err := requests.GetDriverTypes(s.autoRiaKey)
	if err != nil {
		switch {
		case errors.Is(err, requests.ErrDriverTypesFetch):
			log.Error(ErrDriverTypesFetch.Error(), slog.String("error", err.Error()))

		case errors.Is(err, requests.ErrDecodeDriverTypes):
			log.Error(ErrDecodeDriverTypes.Error(), slog.String("error", err.Error()))
		}
		return ErrRefreshDriverTypes

	}
	for _, oneDriverType := range driverTypes {
		err = s.repo.InsertOrUpdate(ctx, oneDriverType)
		if err != nil {
			return ErrRefreshDriverTypes
		}
	}
	return nil
}

func (s Service) GetAll(ctx context.Context) ([]models.DriverType, error) {
	log := s.log.With(slog.String("op", "services.driver_type.get_all"))

	driverTypes, err := s.repo.GetAll(ctx)
	if err != nil {
		log.Error(ErrGetDriverTypes.Error(), slog.String("error", err.Error()))
		return nil, ErrGetDriverTypes
	}
	return driverTypes, nil
}
