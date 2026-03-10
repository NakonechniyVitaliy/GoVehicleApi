package service

import (
	"context"
	"log/slog"

	repoErrors "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/_errors"

	bsDTO "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/body_style"
	bDTO "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/brand"
	cDTO "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/category"
	dDTO "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/driver_type"
	gDTO "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/gearbox"
	vDTO "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/vehicle"

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
}

func NewService(repos *repos.Repositories, logger *slog.Logger) *Service {
	return &Service{
		repos.Vehicle,
		repos.Brand,
		repos.BodyStyle,
		repos.Category,
		repos.DriverType,
		repos.Gearbox,
		logger,
	}
}

func (s Service) GetByID(ctx context.Context, id uint16) (*models.Vehicle, error) {
	vehicle, err := s.vehicleRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return vehicle, nil
}

func (s Service) GetAll(ctx context.Context) ([]models.Vehicle, error) {
	vehicles, err := s.vehicleRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return vehicles, nil
}

func (s Service) Save(ctx context.Context, vehicle models.Vehicle) (*models.Vehicle, error) {
	savedVehicle, err := s.vehicleRepo.Create(ctx, vehicle)
	if err != nil {
		return nil, err
	}
	return savedVehicle, nil
}

func (s Service) Update(ctx context.Context, vehicle models.Vehicle, id uint16) (*models.Vehicle, error) {
	updatedVehicle, err := s.vehicleRepo.Update(ctx, vehicle, id)
	if err != nil {
		return nil, err
	}
	return updatedVehicle, nil
}

func (s Service) Delete(ctx context.Context, id uint16) error {
	err := s.vehicleRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) GetExpanded(ctx context.Context, id uint16) (*vDTO.ExpandedVehicleDTO, error) {

	rawVehicle, err := s.GetByID(ctx, id)
	if err != nil {
		s.log.Error("failed to get raw vehicle", slog.String("error", err.Error()))
		return nil, repoErrors.ErrVehicleNotFound
	}

	vBrand, err := s.brandRepo.GetByID(ctx, rawVehicle.Brand)
	if err != nil {
		return nil, repoErrors.ErrBrandNotFound
	}
	vBodyStyle, err := s.bodyRepo.GetByID(ctx, rawVehicle.BodyStyle)
	if err != nil {
		return nil, repoErrors.ErrBodyStyleNotFound
	}
	vCategory, err := s.categoryRepo.GetByID(ctx, rawVehicle.Category)
	if err != nil {
		return nil, repoErrors.ErrCategoryNotFound
	}
	vDriverType, err := s.driverRepo.GetByID(ctx, rawVehicle.DriverType)
	if err != nil {
		return nil, repoErrors.ErrDriverTypeNotFound
	}
	vGearbox, err := s.gearboxRepo.GetByID(ctx, rawVehicle.Gearbox)
	if err != nil {
		return nil, repoErrors.ErrGearboxNotFound
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
