package vehicle

import "github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"

type VehicleDTO struct {
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
	ID uint16 `json:"id"`
	*models.Brand
	*models.DriverType
	*models.Gearbox
	*models.BodyStyle
	*models.Category
	Mileage uint32 `json:"mileage"`
	Model   string `json:"model"`
	Price   uint16 `json:"price"`
}
