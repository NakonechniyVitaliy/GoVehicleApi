package services

import (
	bodyStyleService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/body_style"
	brandService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/brand"
	categoryService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/category"
	driverTypeService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/driver_type"
	gearboxService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/gearbox"
	jwtService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/jwt"
	userService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/user"
	vehicleService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/vehicle"
)

type Container struct {
	Brand      *brandService.Service
	BodyStyle  *bodyStyleService.Service
	Category   *categoryService.Service
	DriverType *driverTypeService.Service
	Gearbox    *gearboxService.Service
	Vehicle    *vehicleService.Service
	User       *userService.Service
	JWT        *jwtService.Service
}
