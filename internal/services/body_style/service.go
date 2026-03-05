package service

import (
	"context"
	"log/slog"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	bodyStyleRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/body_style"
	bodyStyleRequests "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/body_styles"
)

type Service struct {
	repo       bodyStyleRepo.RepositoryInterface
	log        *slog.Logger
	autoRiaKey string
}

func NewService(repository bodyStyleRepo.RepositoryInterface, logger *slog.Logger, key string) *Service {
	return &Service{
		repo:       repository,
		log:        logger,
		autoRiaKey: key,
	}
}

func (s Service) Refresh(ctx context.Context) error {
	bodyStyles, err := bodyStyleRequests.GetBodyStyles(s.autoRiaKey)
	if err != nil {
		return err
	}
	for _, oneCategory := range bodyStyles {
		err = s.repo.InsertOrUpdate(ctx, oneCategory)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s Service) GetByID(ctx context.Context, id uint16) (*models.BodyStyle, error) {
	bodyStyle, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return bodyStyle, nil
}

func (s Service) GetAll(ctx context.Context) ([]models.BodyStyle, error) {
	bodyStyles, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return bodyStyles, nil
}

func (s Service) Save(ctx context.Context, bodyStyle models.BodyStyle) (*models.BodyStyle, error) {
	savedBodyStyle, err := s.repo.Create(ctx, bodyStyle)
	if err != nil {
		return nil, err
	}
	return savedBodyStyle, nil
}

func (s Service) Update(ctx context.Context, bodyStyle models.BodyStyle, id uint16) (*models.BodyStyle, error) {
	updatedBodyStyle, err := s.repo.Update(ctx, bodyStyle, id)
	if err != nil {
		return nil, err
	}
	return updatedBodyStyle, nil
}

func (s Service) Delete(ctx context.Context, id uint16) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
