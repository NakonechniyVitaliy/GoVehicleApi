package vehicle

import (
	bodyStyleDTO "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/body_style"
	brandDTO "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/brand"
	categoryDTO "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/category"
	driverDTO "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/driver_type"
	gearboxDTO "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/gearbox"
)

type VehicleDTO struct {
	ID         *uint16 `json:"id,omitempty"`
	Brand      *uint16 `json:"brand"`
	DriverType *uint16 `json:"driver_type"`
	Gearbox    *uint16 `json:"gearbox"`
	BodyStyle  *uint16 `json:"body_style"`
	Category   *uint16 `json:"category"`
	Mileage    *uint32 `json:"mileage"`
	Model      *string `json:"model"`
	Price      *uint16 `json:"price"`
}

type ExpandedVehicleDTO struct {
	ID         uint16                      `json:"id"`
	Brand      brandDTO.CompressedBrandDTO `json:"brand"`
	DriverType driverDTO.DriverTypeDTO     `json:"driver_type"`
	Gearbox    gearboxDTO.GearboxDTO       `json:"gearbox"`
	BodyStyle  bodyStyleDTO.BodyStyleDTO   `json:"body_style"`
	Category   categoryDTO.CategoryDTO     `json:"category"`
	Mileage    uint32                      `json:"mileage"`
	Model      string                      `json:"model"`
	Price      uint16                      `json:"price"`
}
