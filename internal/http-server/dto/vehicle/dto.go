package vehicle

type VehicleDTO struct {
	Brand      *uint16 `json:"brand"`
	DriverType *uint16 `json:"driverType"`
	Gearbox    *uint16 `json:"gearbox"`
	BodyStyle  *uint16 `json:"bodyStyle"`
	Category   *uint16 `json:"category"`
	Mileage    *uint32 `json:"mileage"`
	Model      *string `json:"model"`
	Price      *uint16 `json:"price"`
}
