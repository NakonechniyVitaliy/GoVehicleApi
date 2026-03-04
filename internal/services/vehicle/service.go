package service

import (
	"context"

	dto "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/vehicle"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"

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
}

func NewService(
	vehicleRepo vehicleRepo.RepositoryInterface,
	brandRepo brandRepo.RepositoryInterface,
	bodyRepo bodyStyleRepo.RepositoryInterface,
	categoryRepo categoryRepo.RepositoryInterface,
	driverRepo driverTypeRepo.RepositoryInterface,
	gearboxRepo gearboxRepo.RepositoryInterface,
) *Service {
	return &Service{
		vehicleRepo,
		brandRepo,
		bodyRepo,
		categoryRepo,
		driverRepo,
		gearboxRepo,
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

func (s Service) GetExpanded(ctx context.Context, vehicle models.Vehicle) (*dto.ExpandedVehicleDTO, error) {
	vBrand, err := s.brandRepo.GetByID(ctx, vehicle.Brand)
	if err != nil {
		return nil, err
	}
	vBodyStyle, err := s.bodyRepo.GetByID(ctx, vehicle.BodyStyle)
	if err != nil {
		return nil, err
	}
	vCategory, err := s.categoryRepo.GetByID(ctx, vehicle.Category)
	if err != nil {
		return nil, err
	}
	vDriverType, err := s.driverRepo.GetByID(ctx, vehicle.DriverType)
	if err != nil {
		return nil, err
	}
	vGearbox, err := s.gearboxRepo.GetByID(ctx, vehicle.Gearbox)
	if err != nil {
		return nil, err
	}

	return &dto.ExpandedVehicleDTO{
		ID:         vehicle.ID,
		Brand:      vBrand,
		DriverType: vDriverType,
		Gearbox:    vGearbox,
		BodyStyle:  vBodyStyle,
		Category:   vCategory,
		Mileage:    vehicle.Mileage,
		Model:      vehicle.Model,
		Price:      vehicle.Price,
	}, nil
}
