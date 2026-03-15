package repository

import (
	bodyStyleRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/body_style"
	brandRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/brand"
	categoryRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/category"
	driverTypeRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/driver_type"
	gearboxRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/gearbox"
	userRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/user"
	vehicleRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle"
)

type Repositories struct {
	Brand      brandRepo.RepositoryInterface
	BodyStyle  bodyStyleRepo.RepositoryInterface
	Category   categoryRepo.RepositoryInterface
	DriverType driverTypeRepo.RepositoryInterface
	Gearbox    gearboxRepo.RepositoryInterface
	Vehicle    vehicleRepo.RepositoryInterface
	User       userRepo.RepositoryInterface
}
