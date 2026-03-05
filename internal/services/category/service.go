package service

import (
	"context"
	"log/slog"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	categoryRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/category"
	categoryRequests "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/categories"
)

type Service struct {
	repo       categoryRepo.RepositoryInterface
	log        *slog.Logger
	autoRiaKey string
}

func NewService(repository categoryRepo.RepositoryInterface, logger *slog.Logger, key string) *Service {
	return &Service{
		repo:       repository,
		log:        logger,
		autoRiaKey: key,
	}
}

func (s Service) Refresh(ctx context.Context) error {

	categories, err := categoryRequests.GetCategories(s.autoRiaKey)
	if err != nil {
		return err
	}
	for _, oneCategory := range categories {
		err = s.repo.InsertOrUpdate(ctx, oneCategory)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s Service) GetAll(ctx context.Context) ([]models.Category, error) {
	categories, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
