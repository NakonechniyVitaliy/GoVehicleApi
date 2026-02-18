package vehicle

import (
	"fmt"
)

type SaveRequest struct {
	Vehicle VehicleDTO
}

func (r SaveRequest) Validate() error {
	b := r.Vehicle

	if b.Brand == nil ||
		b.DriverType == nil ||
		b.Gearbox == nil ||
		b.BodyStyle == nil ||
		b.Category == nil ||
		b.Mileage == nil ||
		b.Model == nil ||
		b.Price == nil {
		return fmt.Errorf("all fields are required")
	}
	return nil
}

type UpdateRequest struct {
	Vehicle VehicleDTO
}

func (r UpdateRequest) Validate() error {
	b := r.Vehicle

	if b.Brand == nil ||
		b.DriverType == nil ||
		b.Gearbox == nil ||
		b.BodyStyle == nil ||
		b.Category == nil ||
		b.Mileage == nil ||
		b.Model == nil ||
		b.Price == nil {
		return fmt.Errorf("all fields are required")
	}
	return nil
}
