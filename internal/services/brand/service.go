package brand

import (
	"context"
	"errors"
	"log/slog"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	repoErrors "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/_errors"
	brandRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/brand"
	requests "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/brands"
)

type Service struct {
	repo       brandRepo.RepositoryInterface
	log        *slog.Logger
	autoRiaKey string
}

func NewService(repository brandRepo.RepositoryInterface, logger *slog.Logger, key string) *Service {
	return &Service{
		repo:       repository,
		log:        logger,
		autoRiaKey: key,
	}
}

func (s Service) Refresh(ctx context.Context) error {
	log := s.log.With(slog.String("op", "services.brand.refresh"))

	brands, err := requests.GetBrands(s.autoRiaKey)
	if err != nil {
		switch {
		case errors.Is(err, requests.ErrBrandsFetch):
			log.Error(ErrBrandsFetch.Error(), slog.String("error", err.Error()))

		case errors.Is(err, requests.ErrDecodeBrands):
			log.Error(ErrDecodeBrands.Error(), slog.String("error", err.Error()))
		}
		return ErrRefreshBrands
	}

	for _, oneBrand := range brands {
		err = s.repo.InsertOrUpdate(ctx, oneBrand)
		if err != nil {
			return ErrRefreshBrands
		}
	}
	return nil
}

func (s Service) GetByID(ctx context.Context, id uint16) (*models.Brand, error) {
	log := s.log.With(slog.String("op", "services.brand.get_by_id"))

	brand, err := s.repo.GetByID(ctx, id)

	if errors.Is(err, repoErrors.ErrBrandNotFound) {
		log.Error(ErrBrandNotFound.Error(), slog.String("error", err.Error()))
		return nil, ErrBrandNotFound
	}

	if err != nil {
		log.Error(ErrGetBrand.Error(), slog.String("error", err.Error()))
		return nil, err
	}
	return brand, nil
}

func (s Service) GetAll(ctx context.Context) ([]models.Brand, error) {
	log := s.log.With(slog.String("op", "handlers.brand.get_all"))

	brands, err := s.repo.GetAll(ctx)
	if err != nil {
		log.Error(ErrGetBrands.Error(), slog.String("error", err.Error()))
		return nil, ErrGetBrands
	}
	return brands, nil
}

func (s Service) Save(ctx context.Context, brand models.Brand) (*models.Brand, error) {
	log := s.log.With(slog.String("op", "services.brand.save"))

	savedBrand, err := s.repo.Create(ctx, brand)
	if errors.Is(err, repoErrors.ErrBrandExists) {
		log.Error(ErrBrandExists.Error(), slog.String("error", err.Error()))
		return nil, ErrBrandExists
	}
	if err != nil {
		log.Error(ErrSaveBrand.Error(), slog.String("error", err.Error()))
		return nil, ErrSaveBrand
	}

	return savedBrand, nil
}

func (s Service) Update(ctx context.Context, brand models.Brand, id uint16) (*models.Brand, error) {
	log := s.log.With(slog.String("op", "services.brand.update"))

	updatedBrand, err := s.repo.Update(ctx, brand, id)

	if errors.Is(err, repoErrors.ErrBrandNotFound) {
		log.Error(ErrBrandNotFound.Error(), slog.String("error", err.Error()))
		return nil, ErrBrandNotFound
	}

	if err != nil {
		log.Error(ErrUpdateBrand.Error(), slog.String("error", err.Error()))
		return nil, ErrUpdateBrand
	}

	return updatedBrand, nil
}

func (s Service) Delete(ctx context.Context, id uint16) error {
	log := s.log.With(slog.String("op", "services.brand.delete"))

	err := s.repo.Delete(ctx, id)

	if errors.Is(err, repoErrors.ErrBrandNotFound) {
		log.Error(ErrBrandNotFound.Error(), slog.String("error", err.Error()))
		return ErrBrandNotFound
	}

	if err != nil {
		log.Error("failed to delete brand", slog.String("error", err.Error()))
		return ErrGetBrand
	}

	return nil
}
