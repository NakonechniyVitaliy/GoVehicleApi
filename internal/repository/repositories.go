package repository

import (
	bodyStyleRepo "github.com/NakonechniyVitalii/GoVehicleApi/internal/repository/body_style"
	brandRepo "github.com/NakonechniyVitalii/GoVehicleApi/internal/repository/brand"
	categoryRepo "github.com/NakonechniyVitalii/GoVehicleApi/internal/repository/category"
	driverTypeRepo "github.com/NakonechniyVitalii/GoVehicleApi/internal/repository/driver_type"
	gearboxRepo "github.com/NakonechniyVitalii/GoVehicleApi/internal/repository/gearbox"
	userRepo "github.com/NakonechniyVitalii/GoVehicleApi/internal/repository/user"
	vehicleRepo "github.com/NakonechniyVitalii/GoVehicleApi/internal/repository/vehicle"
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
