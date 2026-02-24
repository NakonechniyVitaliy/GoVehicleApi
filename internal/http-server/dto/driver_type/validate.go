package driver_type

import (
	"fmt"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/_errors"
)

type SaveRequest struct {
	DriverTypeDTO
}

func (r SaveRequest) Validate() error {
	c := r.DriverTypeDTO

	if c.Name == nil || c.Value == nil {
		return fmt.Errorf(_errors.AllFieldsAreRequired)
	}
	return nil
}

type UpdateRequest struct {
	DriverTypeDTO
}

func (r UpdateRequest) Validate() error {
	c := r.DriverTypeDTO

	if c.Name == nil || c.Value == nil {
		return fmt.Errorf(_errors.AllFieldsAreRequired)
	}
	return nil
}
