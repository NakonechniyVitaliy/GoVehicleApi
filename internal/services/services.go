package services

import (
	bodyStyleService "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/body_style"
	brandService "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/brand"
	categoryService "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/category"
	driverTypeService "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/driver_type"
	gearboxService "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/gearbox"
	jwtService "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/jwt"
	userService "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/user"
	vehicleService "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/vehicle"
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
