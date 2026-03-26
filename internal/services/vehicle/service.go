package vehicle

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/cache"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/cache_key"
	repoErrors "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/_errors"
	"github.com/redis/go-redis/v9"

	bsDTO "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/body_style"
	bDTO "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/brand"
	cDTO "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/category"
	dDTO "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/driver_type"
	gDTO "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/gearbox"
	vDTO "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/vehicle"
	fDTO "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/vehicle/filter"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/helper"

	repos "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository"

	bodyStyleRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/body_style"
	brandRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/brand"
	categoryRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/category"
	driverTypeRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/driver_type"
	gearboxRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/gearbox"
	vehicleRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle"
)

type Service struct {
	vehicleRepo  vehicleRepo.RepositoryInterface
	brandRepo    brandRepo.RepositoryInterface
	bodyRepo     bodyStyleRepo.RepositoryInterface
	categoryRepo categoryRepo.RepositoryInterface
	driverRepo   driverTypeRepo.RepositoryInterface
	gearboxRepo  gearboxRepo.RepositoryInterface
	log          *slog.Logger
	cache        *cache.AppCache
}

func NewService(repos *repos.Repositories, logger *slog.Logger, cache *cache.AppCache) *Service {
	return &Service{
		repos.Vehicle,
		repos.Brand,
		repos.BodyStyle,
		repos.Category,
		repos.DriverType,
		repos.Gearbox,
		logger,
		cache,
	}
}

func (s *Service) GetByID(ctx context.Context, id uint16) (*models.Vehicle, error) {
	log := s.log.With(slog.String("op", "services.vehicle.get_by_id"))

	cacheKey := cache_key.VehicleByID(id)

	var cachedVehicle models.Vehicle
	err := s.cache.Get(ctx, cacheKey, &cachedVehicle)
	if err == nil {
		return &cachedVehicle, nil
	}
	if !errors.Is(err, redis.Nil) {
		log.Error("redis get error", slog.String("error", err.Error()))
	}

	vehicle, err := s.vehicleRepo.GetByID(ctx, id)
	if errors.Is(err, repoErrors.ErrVehicleNotFound) {
		log.Error(ErrVehicleNotFound.Error(), slog.String("error", err.Error()))
		return nil, ErrVehicleNotFound
	}
	if err != nil {
		log.Error(ErrGetVehicle.Error(), slog.String("error", err.Error()))
		return nil, err
	}

	err = s.cache.Set(ctx, cacheKey, vehicle)
	if err != nil {
		log.Error("redis set error", slog.String("error", err.Error()))
	}

	return vehicle, nil
}

func (s Service) GetList(ctx context.Context, f fDTO.Filter) ([]models.Vehicle, error) {
	log := s.log.With(slog.String("op", "handlers.vehicle.get_all"))

	if f.Limit == 0 && f.Page == 0 {
		vehicles, err := s.vehicleRepo.GetAll(ctx)
		if err != nil {
			log.Error(ErrGetVehicles.Error(), slog.String("error", err.Error()))
			return nil, ErrGetVehicles
		}
		return vehicles, nil
	}

	return s.GetByPage(ctx, f)

}

func (s Service) GetByPage(ctx context.Context, f fDTO.Filter) ([]models.Vehicle, error) {
	log := s.log.With(slog.String("op", "handlers.vehicle.get_by_page"))

	vehicles, err := s.vehicleRepo.GetByPage(ctx, f)
	if err != nil {
		log.Error(ErrGetVehicles.Error(), slog.String("error", err.Error()))
		return nil, ErrGetVehicles
	}
	return vehicles, nil
}

func (s Service) Save(ctx context.Context, vehicle models.Vehicle) (*models.Vehicle, error) {
	log := s.log.With(slog.String("op", "services.vehicle.save"))

	savedVehicle, err := s.vehicleRepo.Create(ctx, vehicle)
	if errors.Is(err, repoErrors.ErrVehicleExists) {
		log.Error(ErrVehicleExists.Error(), slog.String("error", err.Error()))
		return nil, ErrVehicleExists
	}
	if err != nil {
		log.Error(ErrSaveVehicle.Error(), slog.String("error", err.Error()))
		return nil, ErrSaveVehicle
	}

	return savedVehicle, nil
}

func (s Service) Update(ctx context.Context, vehicle models.Vehicle, id uint16) (*models.Vehicle, error) {
	log := s.log.With(slog.String("op", "services.vehicle.update"))

	updatedVehicle, err := s.vehicleRepo.Update(ctx, vehicle, id)
	if errors.Is(err, repoErrors.ErrVehicleNotFound) {
		log.Error(ErrVehicleNotFound.Error(), slog.String("error", err.Error()))
		return nil, ErrVehicleNotFound
	}

	if err != nil {
		log.Error(ErrUpdateVehicle.Error(), slog.String("error", err.Error()))
		return nil, ErrUpdateVehicle
	}

	redisKey := fmt.Sprintf("vehicle:%d", id)
	if err := s.cache.Delete(ctx, redisKey); err != nil {
		log.Error("redis delete error", slog.String("error", err.Error()))
	}

	return updatedVehicle, nil
}

func (s Service) Delete(ctx context.Context, id uint16) error {
	log := s.log.With(slog.String("op", "services.vehicle.delete"))

	err := s.vehicleRepo.Delete(ctx, id)
	if errors.Is(err, repoErrors.ErrVehicleNotFound) {
		log.Error(ErrVehicleNotFound.Error(), slog.String("error", err.Error()))
		return ErrVehicleNotFound
	}
	if err != nil {
		log.Error("failed to delete vehicle", slog.String("error", err.Error()))
		return ErrGetVehicle
	}

	redisKey := fmt.Sprintf("vehicle:%d", id)
	if err := s.cache.Delete(ctx, redisKey); err != nil {
		log.Error("redis delete error", slog.String("error", err.Error()))
	}

	return nil
}

func (s Service) GetExpanded(ctx context.Context, id uint16) (*vDTO.ExpandedVehicleDTO, error) {

	rawVehicle, err := s.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	vBrand, err := s.brandRepo.GetByID(ctx, rawVehicle.Brand)
	if err != nil {
		return nil, err
	}
	vBodyStyle, err := s.bodyRepo.GetByID(ctx, rawVehicle.BodyStyle)
	if err != nil {
		return nil, err
	}
	vCategory, err := s.categoryRepo.GetByID(ctx, rawVehicle.Category)
	if err != nil {
		return nil, err
	}
	vDriverType, err := s.driverRepo.GetByID(ctx, rawVehicle.DriverType)
	if err != nil {
		return nil, err
	}
	vGearbox, err := s.gearboxRepo.GetByID(ctx, rawVehicle.Gearbox)
	if err != nil {
		return nil, err
	}

	return &vDTO.ExpandedVehicleDTO{
		ID: rawVehicle.ID,
		Brand: bDTO.CompressedBrandDTO{
			ID:   vBrand.ID,
			Name: vBrand.Name,
		},
		DriverType: dDTO.DriverTypeDTO{
			ID:   helper.PtrUint16(vDriverType.ID),
			Name: helper.PtrString(vDriverType.Name),
		},
		Gearbox: gDTO.GearboxDTO{
			ID:   helper.PtrUint16(vGearbox.ID),
			Name: helper.PtrString(vGearbox.Name),
		},
		BodyStyle: bsDTO.BodyStyleDTO{
			ID:   helper.PtrUint16(vBodyStyle.ID),
			Name: helper.PtrString(vBodyStyle.Name),
		},
		Category: cDTO.CategoryDTO{
			ID:   helper.PtrUint16(vCategory.ID),
			Name: helper.PtrString(vCategory.Name),
		},
		Mileage: rawVehicle.Mileage,
		Model:   rawVehicle.Model,
		Price:   rawVehicle.Price,
	}, nil
}
