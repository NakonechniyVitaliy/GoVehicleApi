package services

import bodyStyleService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/body_style"
import brandService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/brand"
import categoryService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/category"
import driverTypeService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/driver_type"
import gearboxService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/gearbox"
import vehicleService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/vehicle"

type Container struct {
	Brand      *brandService.Service
	BodyStyle  *bodyStyleService.Service
	Category   *categoryService.Service
	DriverType *driverTypeService.Service
	Gearbox    *gearboxService.Service
	Vehicle    *vehicleService.Service
}
