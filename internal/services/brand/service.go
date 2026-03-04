package service

import (
	"context"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	brandRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/brand"
	brandRequests "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/brands"
)

type Service struct {
	repo       brandRepo.RepositoryInterface
	autoRiaKey string
}

func NewService(repository brandRepo.RepositoryInterface, key string) *Service {
	return &Service{
		repo:       repository,
		autoRiaKey: key,
	}
}

func (s Service) Refresh(ctx context.Context, autoRiaKey string) error {

	brands, err := brandRequests.GetBrands(autoRiaKey)
	if err != nil {
		return err
	}
	for _, oneBrand := range brands {
		err = s.repo.InsertOrUpdate(ctx, oneBrand)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s Service) GetByID(ctx context.Context, id uint16) (*models.Brand, error) {
	brand, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return brand, nil
}

func (s Service) GetAll(ctx context.Context) ([]models.Brand, error) {
	brands, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return brands, nil
}

func (s Service) Save(ctx context.Context, brand models.Brand) (*models.Brand, error) {
	savedBrand, err := s.repo.Create(ctx, brand)
	if err != nil {
		return nil, err
	}
	return savedBrand, nil
}

func (s Service) Update(ctx context.Context, brand models.Brand, id uint16) (*models.Brand, error) {
	updatedBrand, err := s.repo.Update(ctx, brand, id)
	if err != nil {
		return nil, err
	}
	return updatedBrand, nil
}

func (s Service) Delete(ctx context.Context, id uint16) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
