package category

import (
	"context"
	"errors"
	"log/slog"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	categoryRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/category"
	requests "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/categories"
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

	log := s.log.With(slog.String("op", "services.category.refresh"))

	categories, err := requests.GetCategories(s.autoRiaKey)
	if err != nil {
		switch {
		case errors.Is(err, requests.ErrCategoriesFetch):
			log.Error(ErrCategoriesFetch.Error(), slog.String("error", err.Error()))

		case errors.Is(err, requests.ErrDecodeCategories):
			log.Error(ErrDecodeCategories.Error(), slog.String("error", err.Error()))
		}
		return ErrRefreshCategories

	}
	for _, oneCategory := range categories {
		err = s.repo.InsertOrUpdate(ctx, oneCategory)
		if err != nil {
			return ErrRefreshCategories
		}
	}
	return nil
}

func (s Service) GetAll(ctx context.Context) ([]models.Category, error) {

	log := s.log.With(slog.String("op", "services.category.get_all"))

	categories, err := s.repo.GetAll(ctx)
	if err != nil {
		log.Error(ErrGetCategories.Error(), slog.String("error", err.Error()))
		return nil, ErrGetCategories
	}
	return categories, nil
}
