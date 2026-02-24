package body_style

import (
	"fmt"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/_errors"
)

type SaveRequest struct {
	BodyStyleDTO
}

func (r SaveRequest) Validate() error {
	bs := r.BodyStyleDTO

	if bs.Name == nil || bs.Value == nil {
		return fmt.Errorf(_errors.AllFieldsAreRequired)
	}
	return nil
}

type UpdateRequest struct {
	BodyStyleDTO
}

func (r UpdateRequest) Validate() error {
	bs := r.BodyStyleDTO

	if bs.Name == nil || bs.Value == nil {
		return fmt.Errorf(_errors.AllFieldsAreRequired)
	}
	return nil
}
