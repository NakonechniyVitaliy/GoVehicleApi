package service

import (
	"context"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	categoryRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/category"
	categoryRequests "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/categories"
)

type Service struct {
	repo       categoryRepo.RepositoryInterface
	autoRiaKey string
}

func NewService(repository categoryRepo.RepositoryInterface, key string) *Service {
	return &Service{
		repo:       repository,
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
