package user

import (
	"context"
	"errors"
	"log/slog"

	dto "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/user"
	repoErrors "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/_errors"
	userRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/user"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/helper"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo       userRepo.RepositoryInterface
	log        *slog.Logger
	autoRiaKey string
}

func NewService(repository userRepo.RepositoryInterface, logger *slog.Logger, key string) *Service {
	return &Service{
		repo:       repository,
		log:        logger,
		autoRiaKey: key,
	}
}

func (s Service) SignUp(ctx context.Context, userData dto.SignUpDTO) error {
	log := s.log.With(slog.String("op", "services.user.save"))

	hash, err := bcrypt.GenerateFromPassword([]byte(*userData.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(ErrSaveUser.Error(), slog.String("error", err.Error()))
		return ErrSaveUser
	}

	user := userData.ToModel(hash)

	err = s.repo.Create(ctx, user)
	if errors.Is(err, repoErrors.ErrUserExists) {
		log.Error(ErrUserExists.Error(), slog.String("error", err.Error()))
		return ErrUserExists
	}
	if err != nil {
		log.Error(ErrSaveUser.Error(), slog.String("error", err.Error()))
		return ErrSaveUser
	}

	return nil
}

func (s Service) SignIn(ctx context.Context, signData dto.SignInDTO) (tokenJWT *string, err error) {
	log := s.log.With(slog.String("op", "services.user.sign_in"))

	user, err := s.repo.GetByLogin(ctx, helper.DerefString(signData.Login))

	if errors.Is(err, repoErrors.ErrUserNotFound) {
		log.Error(ErrUserNotFound.Error(), slog.String("error", err.Error()))
		return nil, ErrIncorrectLoginOrPass
	}
	if err != nil {
		log.Error(ErrSignIn.Error(), slog.String("error", err.Error()))
		return nil, ErrSignIn
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(helper.DerefString(signData.Password)))
	if err != nil {
		log.Error(ErrComparePass.Error(), slog.String("error", err.Error()))
		return nil, ErrIncorrectLoginOrPass
	}

	token := "fake-jwt-token"

	return &token, nil
}
