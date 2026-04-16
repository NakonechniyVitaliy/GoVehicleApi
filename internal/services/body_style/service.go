package body_style

import (
	"context"
	"errors"
	"log/slog"

	"github.com/NakonechniyVitalii/GoVehicleApi/internal/models"
	repoErrors "github.com/NakonechniyVitalii/GoVehicleApi/internal/repository/_errors"
	repo "github.com/NakonechniyVitalii/GoVehicleApi/internal/repository/body_style"
	requests "github.com/NakonechniyVitalii/GoVehicleApi/internal/requests/autoria/body_styles"
)

type Service struct {
	repo       repo.RepositoryInterface
	log        *slog.Logger
	autoRiaKey string
}

func NewService(repository repo.RepositoryInterface, logger *slog.Logger, key string) *Service {
	return &Service{
		repo:       repository,
		log:        logger,
		autoRiaKey: key,
	}
}

func (s Service) Fetch(ctx context.Context) error {
	log := s.log.With(slog.String("op", "services.body_style.fetch"))

	bStyles, err := requests.GetBodyStyles(s.autoRiaKey)
	if err != nil {
		switch {
		case errors.Is(err, requests.ErrBodyStylesFetch):
			log.Error(ErrBodyStylesFetch.Error(), slog.String("error", err.Error()))

		case errors.Is(err, requests.ErrDecodeBodyStyles):
			log.Error(ErrDecodeBodyStyles.Error(), slog.String("error", err.Error()))
		}
		return ErrRefreshBodyStyles
	}

	err = s.InsertOrUpdate(ctx, bStyles)
	if err != nil {
		return ErrRefreshBodyStyles
	}

	return nil
}

func (s Service) InsertOrUpdate(ctx context.Context, bStyles []models.BodyStyle) error {
	log := s.log.With(slog.String("op", "services.body_style.insert_or_update"))

	for _, oneBstyle := range bStyles {
		err := s.repo.InsertOrUpdate(ctx, oneBstyle)
		if err != nil {
			log.Error(ErrRefreshBodyStyles.Error(), slog.String("error", err.Error()))
			return err
		}
	}
	return nil
}

func (s Service) GetByID(ctx context.Context, id uint16) (*models.BodyStyle, error) {
	log := s.log.With(slog.String("op", "services.body_style.get_by_id"))

	bodyStyle, err := s.repo.GetByID(ctx, id)

	if errors.Is(err, repoErrors.ErrBodyStyleNotFound) {
		log.Error(ErrBodyStyleNotFound.Error(), slog.String("error", err.Error()))
		return nil, ErrBodyStyleNotFound
	}

	if err != nil {
		log.Error(ErrGetBodyStyle.Error(), slog.String("error", err.Error()))
		return nil, err
	}
	return bodyStyle, nil
}

func (s Service) GetAll(ctx context.Context) ([]models.BodyStyle, error) {
	log := s.log.With(slog.String("op", "handlers.body_style.get_all"))

	bodyStyles, err := s.repo.GetAll(ctx)
	if err != nil {
		log.Error(ErrGetBodyStyles.Error(), slog.String("error", err.Error()))
		return nil, ErrGetBodyStyles
	}
	return bodyStyles, nil
}

func (s Service) Save(ctx context.Context, bodyStyle models.BodyStyle) (*models.BodyStyle, error) {
	log := s.log.With(slog.String("op", "services.body_style.save"))

	savedBodyStyle, err := s.repo.Create(ctx, bodyStyle)
	if errors.Is(err, repoErrors.ErrBodyStyleExists) {
		log.Error(ErrBodyStyleExists.Error(), slog.String("error", err.Error()))
		return nil, ErrBodyStyleExists
	}
	if err != nil {
		log.Error(ErrSaveBodyStyle.Error(), slog.String("error", err.Error()))
		return nil, ErrSaveBodyStyle
	}

	return savedBodyStyle, nil
}

func (s Service) Update(ctx context.Context, bodyStyle models.BodyStyle, id uint16) (*models.BodyStyle, error) {
	log := s.log.With(slog.String("op", "services.body_style.update"))

	updatedBodyStyle, err := s.repo.Update(ctx, bodyStyle, id)

	if errors.Is(err, repoErrors.ErrBodyStyleNotFound) {
		log.Error(ErrBodyStyleNotFound.Error(), slog.String("error", err.Error()))
		return nil, ErrBodyStyleNotFound
	}

	if err != nil {
		log.Error(ErrUpdateBodyStyle.Error(), slog.String("error", err.Error()))
		return nil, ErrUpdateBodyStyle
	}

	return updatedBodyStyle, nil
}

func (s Service) Delete(ctx context.Context, id uint16) error {
	log := s.log.With(slog.String("op", "services.body_style.delete"))

	err := s.repo.Delete(ctx, id)

	if errors.Is(err, repoErrors.ErrBodyStyleNotFound) {
		log.Error(ErrBodyStyleNotFound.Error(), slog.String("error", err.Error()))
		return ErrBodyStyleNotFound
	}

	if err != nil {
		log.Error("failed to delete body style", slog.String("error", err.Error()))
		return ErrGetBodyStyle
	}

	return nil
}
