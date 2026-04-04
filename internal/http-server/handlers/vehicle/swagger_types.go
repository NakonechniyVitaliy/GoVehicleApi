package vehicle

import response "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"

// VehiclePayload is the request body for POST/PUT /vehicle
type VehiclePayload struct {
	Vehicle struct {
		Brand      *uint16 `json:"brand"`
		DriverType *uint16 `json:"driver_type"`
		Gearbox    *uint16 `json:"gearbox"`
		BodyStyle  *uint16 `json:"body_style"`
		Category   *uint16 `json:"category"`
		Mileage    *uint32 `json:"mileage"`
		Model      *string `json:"model"`
		Price      *uint32 `json:"price"`
	} `json:"Vehicle"`
}

type swaggerCompressedBrand struct {
	ID   uint16 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type swaggerBodyStyle struct {
	ID    *uint16 `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Value *uint16 `json:"value,omitempty"`
}

type swaggerCategory struct {
	ID    *uint16 `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Value *uint16 `json:"value,omitempty"`
}

type swaggerDriverType struct {
	ID    *uint16 `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Value *uint16 `json:"value,omitempty"`
}

type swaggerGearbox struct {
	ID    *uint16 `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Value *uint16 `json:"value,omitempty"`
}

// ExpandedVehicleSwagger mirrors dto.ExpandedVehicleDTO for swagger docs.
type ExpandedVehicleSwagger struct {
	ID         uint16                 `json:"id"`
	Brand      swaggerCompressedBrand `json:"brand"`
	DriverType swaggerDriverType      `json:"driver_type"`
	Gearbox    swaggerGearbox         `json:"gearbox"`
	BodyStyle  swaggerBodyStyle       `json:"body_style"`
	Category   swaggerCategory        `json:"category"`
	Mileage    uint32                 `json:"mileage"`
	Model      string                 `json:"model"`
	Price      uint32                 `json:"price"`
}

// GetExpandedSwaggerResponse is used only for swagger @Success annotation.
type GetExpandedSwaggerResponse struct {
	Response response.Response
	Vehicle  *ExpandedVehicleSwagger
}
