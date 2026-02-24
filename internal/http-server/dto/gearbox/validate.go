package gearbox

import (
	"fmt"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/_errors"
)

type SaveRequest struct {
	GearboxDTO
}

func (r SaveRequest) Validate() error {
	g := r.GearboxDTO

	if g.Name == nil || g.Value == nil {
		return fmt.Errorf(_errors.AllFieldsAreRequired)
	}
	return nil
}

type UpdateRequest struct {
	GearboxDTO
}

func (r UpdateRequest) Validate() error {
	g := r.GearboxDTO

	if g.Name == nil || g.Value == nil {
		return fmt.Errorf(_errors.AllFieldsAreRequired)
	}
	return nil
}
