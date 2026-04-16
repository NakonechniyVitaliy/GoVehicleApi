package user

import (
	"context"
	"errors"
	"log/slog"

	dto "github.com/NakonechniyVitalii/GoVehicleApi/internal/http-server/dto/user"
	repoErrors "github.com/NakonechniyVitalii/GoVehicleApi/internal/repository/_errors"
	userRepo "github.com/NakonechniyVitalii/GoVehicleApi/internal/repository/user"
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/services/helper"
	jwtService "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo userRepo.RepositoryInterface
	log  *slog.Logger
	key  string
}

func NewService(repository userRepo.RepositoryInterface, logger *slog.Logger, secretJwtKey string) *Service {
	return &Service{
		repo: repository,
		log:  logger,
		key:  secretJwtKey,
	}
}

func (s Service) SignUp(ctx context.Context, userData dto.SignUpDTO) error {
	log := s.log.With(slog.String("op", "services.user.save"))

	hash, err := bcrypt.GenerateFromPassword([]byte(*userData.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(ErrSaveUser.Error(), slog.String("error", err.Error()))
		return ErrSaveUser
	}

	userData.Password = helper.PtrString(string(hash))
	user := userData.ToModel()

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

func (s Service) SignIn(ctx context.Context, signData dto.SignInDTO, jwtService *jwtService.Service) (string, error) {
	log := s.log.With(slog.String("op", "services.user.sign_in"))

	user, err := s.repo.GetByLogin(ctx, helper.DerefString(signData.Login))

	if errors.Is(err, repoErrors.ErrUserNotFound) {
		log.Error(ErrUserNotFound.Error(), slog.String("error", err.Error()))
		return "", ErrIncorrectLoginOrPass
	}
	if err != nil {
		log.Error(ErrSignIn.Error(), slog.String("error", err.Error()))
		return "", ErrSignIn
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(helper.DerefString(signData.Password)))
	if err != nil {
		log.Error(ErrComparePass.Error(), slog.String("error", err.Error()))
		return "", ErrIncorrectLoginOrPass
	}

	token, err := jwtService.GenerateToken(user.ID)
	if err != nil {
		log.Error(ErrSignIn.Error(), slog.String("error", err.Error()))
		return "", ErrSignIn
	}

	return token, nil
}
